FROM golang:alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY main.go main.go
COPY handlers/ handlers/

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o=node-down-webhook


FROM scratch AS runner

COPY --from=builder /app/node-down-webhook /node-down-webhook

ENTRYPOINT ["/node-down-webhook"]
CMD []
