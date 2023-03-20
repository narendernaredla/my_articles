FROM golang:1.17.7 as builder

WORKDIR /myarticles/

COPY . .

RUN go build -o myarticles /myarticles/cmd/main.go

EXPOSE 8080

CMD ./myarticles