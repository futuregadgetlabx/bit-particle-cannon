FROM golang:1.20-alpine AS builder
LABEL maintainer="cruii <cruii811@gmail.com>" \
  org.label-schema.name="drone-feishu" \
  org.label-schema.vendor="cruii" \
  org.label-schema.schema-version="1.0"

LABEL org.opencontainers.image.source=https://github.com/futuregadgetlabx/drone-feishu
LABEL org.opencontainers.image.description="bit-particle-cannon"
LABEL org.opencontainers.image.licenses=MIT
COPY . /app
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o bpc .

FROM alpine:3.18.4
LABEL maintainer="cruii <cruii811@gmail.com>"
COPY --from=builder /app/bpc /app/bpc
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
ENV GIN_MODE=release
ENV ENV=release
EXPOSE 8081
ENTRYPOINT ["./app/bpc"]