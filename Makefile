#
# Build metrics container
#

build:
	(cd metrics && go build)

container:
	(cd metrics && docker build -t metrics .)

test:
	py.test-3 test_client.py
