FROM golang:1.17.7-alpine as builder

WORKDIR /myarticles/

COPY . .

RUN go mod download

RUN go build -o myarticle /myarticles/cmd/main.go

FROM alpine:latest

WORKDIR /myarticles/

COPY --from=builder /myarticles/myarticle /myarticles/myarticle

EXPOSE 8080

CMD ./myarticle