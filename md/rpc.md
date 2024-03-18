下载protoc
再GitHub上下载对应版本：
https://github.com/google/protobuf/releases

protoc --go_out=.  test.proto

protoc --go_out=plugins=grpc:. .\test.proto   这样才有    RegisterHelloServer

rpc.zip  里面有  protoc  