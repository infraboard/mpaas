FROM golang:1.20 AS builder

LABEL stage=gobuilder

COPY . /code

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY https://goproxy.cn,direct
# ENV GOPRIVATE="*.gitlab.com"

WORKDIR /code
RUN make build

FROM alpine

WORKDIR /mpaas

COPY --from=builder /code/dist/mpaas /mpaas/mpaas-api
COPY --from=builder /code/etc /mpaas/etc

CMD ["./mpaas-api", "start", "-t", "env"]