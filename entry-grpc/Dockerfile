# STAGE 1 BUILD
FROM golang:1.11.4-alpine3.8 as builder
# Add ca-certs
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

ENV GO111MODULE=on

# Run Build
RUN mkdir /entry
WORKDIR /entry
COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /entry/entry ./entry-grpc/main.go

# STAGE 2 Create Minimal Image
FROM scratch
COPY --from=builder /entry/entry /go/bin/entry
ENTRYPOINT ["/go/bin/entry"]