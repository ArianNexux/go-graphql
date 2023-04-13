FROM golang:1.18


WORKDIR /app

COPY . .

EXPOSE 8080

CMD ["tail", "-f /dev/null"]