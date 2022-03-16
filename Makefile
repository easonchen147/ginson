BINARY="ginson"
VERSION=1.0

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY="https://goproxy.cn,direct" GOFLAGS="-mod=mod" go build -v -o ${BINARY}

run:
	@nohup ./${BINARY} > std.out 2>&1 &

docker-build:
	@docker build -t ${BINARY}:${VERSION} .
	@docker rmi $(docker images -f "dangling=true" -q)

docker-run:
	@docker run --name=${BINARY}-${VERSION} -d -p 8080:8080 ${BINARY}:${VERSION}

help:
	@echo "make build 编译程序"
	@echo "make run 运行程序"
	@echo "make docker-build 构建镜像"
	@echo "make docker-run 运行容器"
