FROM golang:alpine

WORKDIR /work

COPY . .

RUN go build -o fizzbuzz

ENTRYPOINT [ "/work/fizzbuzz", "web" ]