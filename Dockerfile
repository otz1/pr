FROM golang:alpine
RUN apk add --no-cache git

WORKDIR /go/src/github.com/otz1/pr
COPY . .

RUN go get ./...
RUN go install

CMD ["pr"]

EXPOSE 8001