FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o confluent-metrics ./cmd

EXPOSE 8090

CMD [ "./confluent-metrics" ]