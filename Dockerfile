FROM golang:1.22

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN echo "Files copied"
RUN go build -o /usr/local/bin/app -buildvcs=false ./cmd/... 

RUN mkdir /config
RUN mkdir /data
RUN mkdir /logs

CMD ["app"]
