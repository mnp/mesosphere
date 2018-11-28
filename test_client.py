#!/usr/bin/env python3

# Integration test for metrics service.  Requires env vars
#   TESTURL    - the service under test

from MetricsClient import MetricsClient
import os
import pytest

@pytest.fixture()
def client():
    return MetricsClient(os.environ['TESTURL'])

def test_node(client):
    (code) = client.post_metrics('/v1/metrics/node/node123/', '{ "timeslice":8888.8, "cpu":3.4, "mem": 5.6 }')
    assert(code == 200)

def test_nodes(client):
    (code) = client.post_metrics("/v1/metrics/nodes/foo234/process/proc567/", '{}')
    assert(code == 200)

def test_analytics_nodes(client):
    (code, json) = client.get_metrics("/v1/analytics/nodes/average/444444.999")
    assert(code == 200)

def test_analytics_procs(client):
    (code, json) = client.get_metrics("/v1/analytics/processes")
    assert(code == 200)

if __name__ == "__main__":
    client = client()
    test_node(client)
    test_nodes(client)
    test_analytics_nodes(client)
    test_analytics_procs(client)
