FROM golang:1.15-alpine

ADD ./ /go/src/vaccine-bot-lamda-aws

WORKDIR /go/src/vaccine-bot-lamda-aws/src

RUN go build -o "vaccine-cowin-telegram-bot"

CMD ["./vaccine-cowin-telegram-bot"]