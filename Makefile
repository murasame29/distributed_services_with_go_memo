compile:
	protoc StructureDataWithProtobuf/api/v1/*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=.

test:
	go test -race ./...

