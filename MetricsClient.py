#!/usr/bin/env python3

# Client to communicate with a metrics service
#
# Usage:
#   mc = MetricsClient("http://localhost:1234")
#   mc.get_metrics("/metrics/nodes")
#   mc.post_metrics("/analytics/nodes", {"data": 567})

import requests
import sys

class MetricsClient(object):
    def __init__(self, testurl):
        self.testurl = testurl

    def get_metrics(self, path):
        r = requests.get(self.testurl + path)
        return (r.status_code, {"contents": "TODO"})

    def post_metrics(self, path, payload):
        r = requests.post(self.testurl + path, data=payload)
        return r.status_code

if __name__ == "__main__":
    # This can be a CLI with command line parsing here.
    pass
