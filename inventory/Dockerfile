FROM golang:alpine as deps
RUN apk add --no-cache git ca-certificates
ADD go.* /app/
WORKDIR /app
RUN go mod download

FROM deps as builder
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-extldflags -w -s" -o bin/app main.go


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/app /app/bin/inventory
ENTRYPOINT ["/app/bin/inventory"]