FROM golang:latest as builder
RUN mkdir /app
ADD Super-market/ /app
WORKDIR /app
RUN go install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 9595
CMD ["go","run","Super-market" ] 
