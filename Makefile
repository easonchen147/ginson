BINARY="ginson"
VERSION=1.0

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY="https://goproxy.cn,direct" go build -v -o ${BINARY}

run:
	@nohup ./${BINARY} > std.out 2>&1 &

docker-build:
	@docker build -t ${BINARY}:${VERSION} .

docker-run:
	@docker run --name=ginson -d -p 8080:8080 ginson:${VERSION}

help:
	@echo "make build 编译二进制"
	@echo "make run 运行二进制"
	@echo "make docker 打包镜像"
	@echo "make docker 运行镜像"
