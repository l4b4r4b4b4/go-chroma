server:
	@go run server/*.go

gen:
	@protoc \
			--proto_path=protobuf "protobuf/chroma.proto" \
			--go_out=common/genproto/chroma --go_opt=paths=source_relative \
			--go-grpc_out=common/genproto/chroma --go-grpc_opt=paths=source_relative \
			--experimental_allow_proto3_optional
          