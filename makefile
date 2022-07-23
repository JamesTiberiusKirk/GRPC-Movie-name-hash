clean-grpc:
	find . -iname '*.pb.go' -exec rm {} \;

clean-mock:
	find . -iname '*mock.go' -exec rm {} \;

clean: clean-grpc clean-mock

generate:
	go generate -v ./...

test: generate
	go test -v ./...
	
test-only: 
	go test -v ./...

mod:
	go mod tiny
