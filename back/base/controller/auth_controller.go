package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"base/config"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// generateState 生成随机的 state 值
func (c *AuthController) generateState() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// GithubAuth 引导用户到 GitHub 授权页面
func (c *AuthController) GithubAuth(ctx *gin.Context) {
	state := c.generateState()
	params := fmt.Sprintf("client_id=%s&redirect_uri=%s&scope=repo user:email&state=%s",
		config.GlobalConfig.Github.ClientID, config.GlobalConfig.Github.RedirectURL, state)
	authURL := fmt.Sprintf("%s?%s", config.GlobalConfig.Github.AuthURL, params)
	ctx.Redirect(http.StatusFound, authURL)
}

// GithubCallback 处理 GitHub 回调
func (c *AuthController) GithubCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	state := ctx.Query("state")

	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	// 请求访问令牌
	reqBody := fmt.Sprintf(
		"client_id=%s&client_secret=%s&code=%s&state=%s",
		config.GlobalConfig.Github.ClientID, config.GlobalConfig.Github.ClientSecret, code, state,
	)

	resp, err := http.Post(config.GlobalConfig.Github.TokenURL, "application/x-www-form-urlencoded", strings.NewReader(reqBody))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	values, _ := url.ParseQuery(string(body))

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": values.Get("access_token"),
		"token_type":   values.Get("token_type"),
		"scope":        values.Get("scope"),
	})
}

// GithubUserInfo 获取 GitHub 用户信息
func (c *AuthController) GithubUserInfo(ctx *gin.Context) {
	authHeader := ctx.GetHeader(config.GlobalConfig.Github.AuthorizationHeader)
	if authHeader == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid authorization header"})
		return
	}

	tokenParts := []string{}
	if len(authHeader) > 6 && authHeader[:7] == "Bearer " {
		tokenParts = append(tokenParts, authHeader[:6], authHeader[7:])
	}

	if len(tokenParts) != 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid authorization header"})
		return
	}

	token := tokenParts[1]
	req, err := http.NewRequest("GET", config.GlobalConfig.Github.UserInfoURL, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set(config.GlobalConfig.Github.AuthorizationHeader, fmt.Sprintf("Bearer %s", token))
	req.Header.Set(config.GlobalConfig.Github.UserAgentHeader, config.GlobalConfig.Github.UserAgentValue)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch user data"})
		return
	}

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userInfo)
}
