# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protodir=proto

protoc --go_out src \
   --go-grpc_out src \
   --grpc-gateway_out src --grpc-gateway_opt allow_delete_body=true \
   --openapiv2_out src/swagger --openapiv2_opt logtostderr=true --openapiv2_opt allow_delete_body=true \
   -I $protodir $protodir/apiservice.proto

protoc --go_out=src -I $protodir $protodir/const.proto

# protoc --proto_path=proto --swagger_out=logtostderr=true:./swaggerui apiservice.proto
