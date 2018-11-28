#
# Metrics container for local minikube testing only. Based on circlci's container for
# consistency.
#

FROM circleci/golang:1.8

WORKDIR /go/src
COPY  . /go/src
RUN go get -d -v ./... && \
    go build -o metrics

CMD ["/go/src/metrics"]
