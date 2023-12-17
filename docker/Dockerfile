FROM golang:1.21.5-alpine3.19

RUN mkdir /opt/helloworld
WORKDIR /opt/helloworld

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o helloworld main.go
RUN mv helloworld /usr/local/bin/

ENTRYPOINT ["helloworld"]