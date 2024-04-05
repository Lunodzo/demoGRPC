generate_grpc_code:
	protoc --go_out=invoice --go_opt=paths=source_relative --go-grpc_out=invoice --go-grpc_opt=paths=source_relative invoice.proto 