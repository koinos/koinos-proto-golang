module github.com/koinos/koinos-proto-golang

go 1.16

require (
	github.com/btcsuite/btcutil v1.0.2 // indirect
	github.com/iancoleman/orderedmap v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/protobuf => github.com/koinos/protobuf-go v1.27.2-0.20211016005428-adb3d63afc5e
