FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o tages cmd/tages/main.go

FROM gcr.io/distroless/base-debian11

COPY --from=builder app/tages .

EXPOSE 80

CMD ["/tages"]