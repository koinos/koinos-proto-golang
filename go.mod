module github.com/koinos/koinos-proto-golang/v2

go 1.16

require (
	github.com/btcsuite/btcutil v1.0.2
	github.com/stretchr/testify v1.8.1
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

replace google.golang.org/protobuf => github.com/koinos/protobuf-go v1.27.2-0.20211026185306-2456c83214fe
