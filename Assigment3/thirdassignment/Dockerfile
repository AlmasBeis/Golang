FROM golang:1.20.2

WORKDIR /app



COPY ./ ./

#install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

#wait for postgres
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o main .

CMD ["./main"]