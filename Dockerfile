FROM registry.cn-hangzhou.aliyuncs.com/godev/golang:1.20 AS builder

LABEL stage=gobuilder

COPY . /src

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY https://goproxy.cn,direct
# ENV GOPRIVATE="*.gitlab.com"

WORKDIR /src
RUN make build

FROM alpine

WORKDIR /app
EXPOSE 8080

COPY --from=builder /src/dist/mpaas /app/mpaas-api
COPY --from=builder /src/etc /app/etc

CMD ["./mpaas-api", "start", "-t", "env"]