FROM golang:alpine

WORKDIR /opt/chalky-server
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
CMD ["chalky-server"]