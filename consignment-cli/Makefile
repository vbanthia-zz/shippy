build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	docker run --rm \
		-e MICRO_REGISTRY=mdns \
		consignment-cli
