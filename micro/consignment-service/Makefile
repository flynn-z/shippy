build:
	# 不再使用 grpc 插件
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy/micro/consignment-service proto/consignment/consignment.proto

run:
	docker run -p 50051:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \ 
	 -e MICRO_REGISTRY=mdns \ 
	 consignment-service