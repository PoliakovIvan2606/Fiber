FROM golang:1.24-alpine AS builder
WORKDIR /usr/local/src
RUN apk --no-cache add bash git make gcc gettext musl-dev

#dependencies
COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download

#build
COPY app ./
# RUN go build -o ./bin/app cmd/app/main.go
# RUN go build -o ./bin/migrate cmd/migrator/migrate.go
RUN make build_docker

FROM alpine AS runner
RUN apk --no-cache add make
COPY --from=builder /usr/local/src/bin/app /
COPY --from=builder /usr/local/src/bin/migrate /
COPY app/configs/config.yaml configs/config.yaml
COPY app/migrations ./migrations
COPY app/Makefile Makefile

CMD ["make", "run_build_docker"]
