FROM golang:1.23-alpine

WORKDIR /app

# 安装必要的构建依赖
# RUN apk add gcc musl-dev
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
# 复制依赖文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码并构建
COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN GOOS=linux go build -o main .

# 使用多阶段构建，创建最小运行环境
FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/main .
# 暴露服务端口
EXPOSE 8888

CMD ["./main"]