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
- [ ] app can retrieve some metrics
- [ ] integ test can push and pull and check sanity
- [ ] app supports aggregations
- [ ] stand up >1 instance 
- [ ] test >1 instance
- [ ] test exercises all the endpoints and checks math
- [ ] math is correct
- [ ] kube helm spec instead of makefile tape
- [ ] centralized logging 

# Demo

```sh
$ make container
...
Successfully tagged metrics:latest

$ 

$ curl  -X POST -d '{ "timeslice":9.9, "cpu":8.8 "mem": 7.7 }'  http://localhost:9911/v1/metrics/node/foo/
{"status_code":400,"message":"Error decoding JSON"}

$ curl -X POST -d '{ "timeslice":2222222.3, "cpu":3.4, "mem": 5.6 }'  http://localhost:9911/v1/metrics/node/foo/

$ curl -X GET  http://localhost:9911/v1/analytics/nodes/average/2.3
{"timeslice":2.3,"cpu":3.4,"mem":5.6}
```

# 
