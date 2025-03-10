# ビルド
FROM golang:1.23-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trumpath -ldflags "-w -s" -o app

# デプロイ
FROM debian:bullseye-slim as deploy
RUN apt-get update

# ビルド で -o app の成果物は /app配下に設置されている
COPY --from=deploy-builder /app/app .
CMD ["./app"]

# ローカル
FROM golang:1.23 as dev

WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]