build:
	protoc -I. --go_out=plugins=micro:$(PWD) \
		proto/consignment/consignment.proto
	docker build -t consignment-service .

run:
	docker run --rm -p 50051:50051 \
		--name consignment-service \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
	consignment-service
