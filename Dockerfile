FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /miniflux-email-client .

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /miniflux-email-client /miniflux-email-client

ENTRYPOINT ["/miniflux-email-client"]
