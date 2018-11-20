#
# Metrics container
# 

FROM golang:1.8-alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 9911

CMD ["app"]
