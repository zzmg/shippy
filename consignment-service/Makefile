build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
	docker build -t consignment-service .
run:
	docker run -p 50051:50051 -e MICRO_REGISTRY=etcd -e MICRO_REGISTRY_ADDRESS=127.0.0.1:12379 consignment-service
	
