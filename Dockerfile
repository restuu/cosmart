# golang:1.18.9-alpine3.17
FROM golang@sha256:dc89e8f094fbb90480d9efc2a2950ed34ceee03f95f373fdf6d96ce3fb2d7f5d as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod vendor
RUN ls -la

COPY . .

RUN GOOS=linux go build -mod vendor -o library .
RUN ls -la

# alpine:3.17.0
FROM alpine@sha256:c0d488a800e4127c334ad20d61d7bc21b4097540327217dfab52262adc02380c

WORKDIR /app

COPY --from=builder /app/library .

CMD ["./library", "webserver"]