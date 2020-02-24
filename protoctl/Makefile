build:
	protoc -I$(GOPATH)/src --proto_path=$(GOPATH)/src/shippy.com/protoctl/ --go_out=plugins=grpc:$(GOPATH)/src  $(GOPATH)/src/shippy.com/protoctl/consignment/*.proto
build-consignment:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy.com/protoctl consignment/*.proto
build-vessel:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy.com/protoctl vessel/*.proto
