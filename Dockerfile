# build stage
FROM golang:1.15.6 as BUILDER
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=BUILDER /app/main .
EXPOSE 8080
CMD ["./main"]