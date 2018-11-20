#!/usr/bin/env python3

from MetricsClient import MetricsClient
import pytest

HOSTPORT = "localhost:9911"
PROTO = "http"

@pytest.fixture()
def client():
    return MetricsClient(HOSTPORT, PROTO)

def test_node(client):
    (code, json) = client.get_metrics("/metrics/node")
    assert(code == 200)

def test_nodes(client):
    (code, json) = client.get_metrics("/metrics/nodes")
    assert(code == 200)

def test_analytics_nodes(client):
    code = client.post_metrics("/analytics/nodes", {})
    assert(code == 200)

def test_analytics_procs(client):
    code = client.post_metrics("/analytics/processes", {})
    assert(code == 200)
