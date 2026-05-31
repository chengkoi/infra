.PHONY: run build swag clean test

# 运行服务
run:
	go run cmd/server/main.go -config configs/config.dev.yaml

# 编译构建
build:
	go build -o bin/server cmd/server/main.go

# 生成Swagger文档
SWAG := $(shell go env GOPATH)/bin/swag

swag:
	$(SWAG) init -g cmd/server/main.go -o docs

# 清理
clean:
	rm -rf bin/
	rm -rf logs/

# 测试
test:
	go test ./... -v -cover