#
# Build metrics container
#

build:
	(cd metrics && go build)

container:
	(cd metrics && docker build -t metrics .)

test:
	pytest test_client.py
