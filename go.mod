module ajaykumar/grpc-lb

go 1.16

require (
	google.golang.org/grpc v1.39.1
	google.golang.org/protobuf v1.27.1
)

replace (
github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
go.uber.org/atomic => github.com/uber-go/atomic v1.5.0
)
