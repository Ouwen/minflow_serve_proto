protoc -I=./tools --go_out=$GOPATH/src ./tools/protobuf/classification.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/protobuf/get_model_metadata.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/protobuf/inference.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/protobuf/input.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/protobuf/model.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/protobuf/predict.proto
protoc -I=./tools --go_out=plugins=grpc:$GOPATH/src ./tools/protobuf/prediction_service.proto

protoc -I=./tools --go_out=$GOPATH/src ./tools/protobuf/regression.proto

protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/example/example.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/example/feature.proto

protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/attr_value.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/function.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/graph.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/node_def.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/op_def.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/resource_handle.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/tensor.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/tensor_shape.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/types.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/framework/versions.proto

protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/protobuf/meta_graph.proto
protoc -I=./tools --go_out=$GOPATH/src ./tools/tensorflow/core/protobuf/saver.proto
