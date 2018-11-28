# Demonstration project

# Features Implemented

- [ ] https, authn, authz, cert management
- [x] skeleton metrics app (golang) opens a port responds to http, and writes log message
- [x] container for app
- [x] client library (python) can hit the app and get a 200
- [x] integ test uses the client
- [x] container can build in circleci 
- [ ] unit test passes in circleci
- [ ] integrate a cache/persistence service - influxdb?
- [ ] app can persist some state
- [x] app can retrieve some metrics
- [x] integ test can push and pull and check sanity
- [ ] app supports aggregations
- [ ] stand up >1 instance 
- [ ] test >1 instance
- [ ] test exercises all the endpoints and checks math
- [ ] math is correct
- [ ] kube helm spec instead of makefile tape
- [ ] centralized logging 

# Local Demo

```sh
$ make start-minikube build-in-minikube run-in-minikube

$ curl $url
Metrics Help

	POST /v1/metrics/node/{nodename}/
	POST /v1/metrics/nodes/{nodename}/process/{processname}/
	GET /v1/analytics/nodes/average
	GET /v1/analytics/processes/
	GET /v1/analytics/processes/{processname}/

$ curl  -X POST -d '{ "timeslice":9999.9, "cpu":8.8, "mem": 7.7 }' $url/v1/metrics/node/foo/

$ curl  -X POST -d '{ "timeslice":8888.8, "cpu":3.4, "mem": 5.6 }' $url/v1/metrics/node/bar/

$ curl $url/v1/analytics/nodes/average/44444
{"timeslice":44444,"cpu":6.1000004,"mem":6.6499996}

```

