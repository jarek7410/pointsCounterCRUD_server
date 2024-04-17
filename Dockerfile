FROM golang:1.22.2
LABEL authors="jarek"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLE=0 GOOS=linux go build -o /docker-app

EXPOSE 2137
EXPOSE 5432
CMD ["/docker-app"]