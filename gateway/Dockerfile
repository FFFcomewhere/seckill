FROM golang:alpine AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build/zero
ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
WORKDIR gateway
RUN go build -ldflags="-s -w" -o /app/main .

FROM alpine
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /build/zero/gateway/bootstrap.yaml /app/bootstrap.yaml
CMD ["./main"]