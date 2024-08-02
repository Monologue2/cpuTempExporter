# Build 단계
FROM golang:alpine AS builder

ENV GO111MOBULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum main.go ./
RUN go mod download

RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .

# Execute 단계
FROM scratch
COPY --from=builder /dist/main .

EXPOSE 4000
ENTRYPOINT ["/main"]