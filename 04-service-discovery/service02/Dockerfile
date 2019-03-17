FROM golang:1.12.1-stretch as builder
WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service02 .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /work/service02 .
EXPOSE 9002
CMD ["./service02"] 