#用于编译时自动处理命令的文件，编译规则

#声明一个虚拟的目标，表示以下命令都需要执行
.PHONY: all build run gotool clean help

BINARY="bluebell"

#表示执行以下所有命令gotool，build
all: gotool build

#编译的命令
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -0 ${BINARY}

#运行
run:
	@go run ./cmd/main.go conf/config1.yml

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码，并编译生成二进制文件"
	@echo "go build - 编译 Go 代码， 生成二进制文件"
	@echo  "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"

