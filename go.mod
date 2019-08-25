module award

go 1.12

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.23.0

require (
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	google.golang.org/grpc v0.0.0-00010101000000-000000000000
)
