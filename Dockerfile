#BUILD
FROM golang:1.18 AS builder

#Set working directory, our binary will be here
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

#Compile the binary for linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

#RUN
FROM alpine:3.8

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/server .

EXPOSE 8000

CMD ["./server"]
