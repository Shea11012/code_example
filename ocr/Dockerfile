FROM golang:1.19.2-alpine3.15

# 安装依赖
RUN apk -U --no-cache add ca-certificates g++ musl-dev tesseract-ocr-dev

ENV TESSDATA_PREFIX=/usr/share/tessdata/

# 安装语言包
RUN apk add --no-cache tesseract-ocr-data-chi_sim

# 列出已安装语言包
RUN tesseract --list-langs

# 配置go环境
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace
