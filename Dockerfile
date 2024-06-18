FROM golang:1.22.4-bullseys as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags "-w -s" -o app


# ----

FROM debian:bullseys-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

FROM golang:1.22.4 as dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD [ "air" ]