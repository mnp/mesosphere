#
# Build metrics container
#

build:
	(cd metrics && go build)

container:
	(cd metrics && docker build -t metrics .)

test:
	py.test test_client.py
