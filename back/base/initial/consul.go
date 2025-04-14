package initial

import (
	"base/config"
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/consul/api"
)

const (
	ServiceName = "hello-service"
	ServicePort = 8888
)

type ConsulConfig struct {
	client *api.Client
}

// InitConsul 初始化Consul客户端并注册服务
func InitConsul() *ConsulConfig {
	// 使用viper获取consul配置
	consulConfig := api.DefaultConfig()
	consulConfig.Address = config.GlobalConfig.Consul.Address
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Failed to create consul client: %v", err)
	}
	return &ConsulConfig{client: consulClient}
}

func (c *ConsulConfig) RegisterService() error {
	ip, err := GetOutboundIP()
	if err != nil {
		return fmt.Errorf("failed to get outbound IP: %v", err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      ServiceName,
		Name:    ServiceName,
		Port:    ServicePort,
		Address: ip.String(),
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/hello", ip.String(), ServicePort),
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	return c.client.Agent().ServiceRegister(registration)
}

func (c *ConsulConfig) DeregisterService() error {
	return c.client.Agent().ServiceDeregister(ServiceName)
}

// GetOutboundIP 获取本机的出口IP
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
