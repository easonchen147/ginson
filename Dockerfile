FROM golang:alpine

ARG BINARY="ginson"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /data/ginson/src

COPY . .

RUN go build -v -o ${BINARY}

WORKDIR /data/ginson/conf
RUN cp /data/ginson/src/app.toml .

WORKDIR /data/ginson
RUN cp /data/ginson/src/${BINARY} .

ENV CONFIG_FILE="/data/ginson/conf/app.toml"

EXPOSE 8080

ENTRYPOINT [ './ginson' ]