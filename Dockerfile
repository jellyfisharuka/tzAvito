FROM golang:1.23
WORKDIR /app
COPY . .

WORKDIR /app/cmd

RUN go build -o main .
ENTRYPOINT [ "./main" ]