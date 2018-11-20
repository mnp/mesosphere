#!/usr/bin/env python3

import requests
import sys

class MetricsClient(object):
    def __init__(self, hostport, proto):
        self.proto = proto
        self.hostport = hostport

    def get_metrics(self, path):
        r = requests.get(self.proto + '://' + self.hostport + path)
        return (r.status_code, {"contents": "TODO"})

    def post_metrics(self, path, payload):
        r = requests.post(self.proto + '://' + self.hostport + path, data=payload)
        return r.status_code

if __name__ == "__main__":
    pass
