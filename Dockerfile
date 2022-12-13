FROM golang:1.19.3-alpine3.16

WORKDIR /app

RUN ls -la

ENV AE_KEY=""

COPY go.mod ./
COPY go.sum ./
RUN #go mod download

COPY . .

CMD ["go", "run", "/app/cmd/api/main.go"]
