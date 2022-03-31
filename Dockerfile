FROM golang:1.18 as build

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

#support chrome dp to use base image
#FROM chromedp/headless-shell:100.0.4896.46 as app

#support chinese web page to add font to base image
#COPY font/simsun.ttc /usr/share/fonts/ttf-dejavu/simsun.ttf

WORKDIR /app/conf
COPY --from=build /go/release/app.toml .

WORKDIR /app
COPY --from=build /go/release/ginson .

ENV CONFIG_FILE="/app/conf/app.toml"

EXPOSE 8080

ENTRYPOINT ["/app/ginson"]