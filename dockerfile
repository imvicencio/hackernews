FROM golang:alpine AS build

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app


COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server .


EXPOSE 8080
CMD ["./server"]