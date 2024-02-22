VERSION = `git rev-parse --short HEAD`
BUILDTIME = `date +%FT%T`

%:
	@true



## tidy: 整理现有的依赖
.PHONY: tidy
tidy:
	go mod tidy

## download: go依赖下载
.PHONY: download
download:
	go mod download

## test: 单元测试全部测试代码
.PHONY: test
test:
	go test ./... -cover

## vet: 静态检测全部go代码
.PHONY: vet
vet:
#	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
#	go vet -vettool=`which shadow` ./...
	go vet ./...

## bench: 并发测试
.PHONY: bench
bench:
	go test ./...  -test.bench . -test.benchmem=true

## test: 单元测试全部测试代码
.PHONY: fmt
fmt:
	gofmt -w -l .


## lint: golangci-lint
.PHONY: lint
lint:
	golangci-lint cache clean
	golangci-lint run


## check: fmt lint vet
.PHONY: check
check: fmt lint vet



## serve: 启动服务
.PHONY: serve
serve:
	cd cmd/k8sviewserver && go run main.go --env=local

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo  "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'



