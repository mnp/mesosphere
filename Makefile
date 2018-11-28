#
# 1. Build metrics container for use in CI:
#    make
#
# 2. Build container for local use
#    make container
#
# 3. Stand up minikube environment and run tests there
#    make test-in-minikube
#
# TODO - rebuild and push for faster iteration, eg
#	kubectl set image deployment $(REPO) *=$(REPO):$(VERSION)

GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
REPO=metrics
VERSION=1

default: build

metrics:
	mkdir -p metrics

build: metrics/metrics

build-native: $(GOFILES)
	go build -o metrics/native-metrics .

metrics/metrics: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o metrics/metrics .

container:
	docker build -t $(REPO) .

test:
	TESTURL=http://localhost:9911 py.test-3 test_client.py

test-minikube: start-minikube build-in-minikube run-in-minikube nuke-minikube

start-minikube:
	sudo -E minikube start --cpus 2 --memory 4096

build-in-minikube:
	@eval $$(minikube docker-env) ;\
	docker build -t $(REPO):$(VERSION) -f Dockerfile .

run-in-minikube:
	eval $$(sudo minikube docker-env)                                          && \
	kubectl run postgresql    --image=circleci/postgres:9.6-alpine --port=5432 && \
	kubectl run hello-metrics --image=$(REPO):$(VERSION) --port=9911           && \
	sleep 5                                                                     ; \
	kubectl expose deployment hello-metrics --type=NodePort
	sleep 120   # TODO: ping loop here
	url=$$(sudo minikube service --url hello-metrics) ;\
	echo service is at $$url ;\
	curl $$url

test-in-minikube:
	url=$$(sudo minikube service --url hello-metrics) ;\
	echo service is at $$url ;\
	TESTURL=url py.test test_client.py

clean-minikube:
	eval $$(sudo minikube docker-env)       ;\
	kubectl delete deployment hello-metrics ;\
	kubectl delete deployment postgresql    ;\
	kubectl delete service hello-metrics

nuke-minikube: clean-minikube
	sudo minikube stop
