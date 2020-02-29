module shippy.com/consignment-cli

go 1.13

require (
	github.com/micro/go-micro v1.18.0
	google.golang.org/grpc v1.25.1
	shippy.com/protoctl v0.0.0-00010101000000-000000000000
)

replace (
	shippy.com/protoctl => ../protoctl
	shippy.com/protoctl/consignment => ../consignment
)
