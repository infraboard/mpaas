FROM registry.cn-hangzhou.aliyuncs.com/godev/golang:1.20 AS builder

LABEL stage=gobuilder

WORKDIR /src
COPY go.mod .
COPY go.sum .

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY https://goproxy.cn,direct
# ENV GOPRIVATE="*.gitlab.com"

# 下载依赖
RUN go mod download
# 执行构建
RUN make build

FROM registry.cn-hangzhou.aliyuncs.com/godev/alpine:latest

WORKDIR /app
EXPOSE 80

COPY --from=builder /src/dist/mpaas /app/mpaas-api
COPY --from=builder /src/etc /app/etc

CMD ["./mpaas-api", "start", "-t", "env"]