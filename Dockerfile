FROM golang:1.19.2-alpine3.16
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . ./
RUN go build  ./main.go
RUN go mod tidy
CMD ["./main"]