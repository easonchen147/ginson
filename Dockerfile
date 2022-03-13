FROM centos:centos7

WORKDIR /app
COPY ${BINARY} .

WORKDIR /app/conf
COPY app.toml .

ENV CONFIG_FILE="/app/conf/app.toml"

EXPOSE 8080

ENTRYPOINT [ "/app/ginson" ]