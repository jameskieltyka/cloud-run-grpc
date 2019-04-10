# STAGE 1 BUILD
FROM golang:1.11.4-alpine3.8 AS builder
# Add ca-certs
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

ENV GO111MODULE=on

# Run Build
RUN mkdir /inter
WORKDIR /inter
COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /inter/inter ./internal-grpc/main.go

# STAGE 2 Create Minimal Image
FROM scratch
COPY --from=builder /inter/inter /go/bin/inter
ENTRYPOINT ["/go/bin/inter"]