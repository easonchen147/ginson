FROM golang:alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOFLAGS="-mod=mod" \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /go/cache
COPY go.mod .
RUN go mod download

WORKDIR /go/release
COPY . .
RUN go build -v -o ginson

FROM alpine:latest as app

WORKDIR /app/conf
COPY --from=build /go/release/app.toml .

WORKDIR /app
COPY --from=build /go/release/ginson .

ENV CONFIG_FILE="/app/conf/app.toml"

EXPOSE 8080

ENTRYPOINT ["/app/ginson"]