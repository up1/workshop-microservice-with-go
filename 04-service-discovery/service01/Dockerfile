FROM golang:1.12.1-stretch as builder
WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service01 .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /work/service01 .
EXPOSE 9001
CMD ["./service01"] 