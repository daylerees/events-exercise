# Generate compiled protocol buffers for Go.
# Requires protoc + protoc-gen-go.
generate:
	protoc -I=./proto --go_out=. ./proto/accounts.proto
