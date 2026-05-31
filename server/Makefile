.PHONY: run build clean test

TAGS := sonic,avx

# 运行服务
run:
	go run -tags $(TAGS) cmd/server/main.go -config configs/config.dev.yaml

# 编译构建
build:
	go build -tags $(TAGS) -o bin/server cmd/server/main.go

# 清理
clean:
	rm -rf bin/
	rm -rf logs/

# 测试
test:
	go test -tags $(TAGS) ./... -v -cover
