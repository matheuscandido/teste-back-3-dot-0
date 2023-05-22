FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./main ./cmd/main.go

FROM alpine:latest as runner
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 4000
ENTRYPOINT [ "./main" ]
