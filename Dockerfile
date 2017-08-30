FROM golang:1.8
MAINTAINER Geektimus

RUN go get github.com/gorilla/mux
VOLUME ./:/w
WORKDIR /w

COPY main.go /w

EXPOSE 8080

CMD ["go", "run", "/w/main.go"]
