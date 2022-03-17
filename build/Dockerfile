FROM golang:1.17.5-alpine3.14 as builder
LABEL description="sbi-port-manager"

COPY main.go /app/main.go
COPY cmd/ /app/cmd/
COPY internal/ /app/internal/
COPY go.mod /app/go.mod
WORKDIR /app
RUN go mod tidy && go build -o /app/sbi-port-manager /app/

FROM alpine:3.14
COPY --from=builder /app/sbi-port-manager /sbi-port-manager
ENTRYPOINT ["/sbi-port-manager"]

EXPOSE 80
