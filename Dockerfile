FROM golang:14.2-alpine

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ENTRYPOINT ["fresh"]