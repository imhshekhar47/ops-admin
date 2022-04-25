dependency:
	go install github.com/spf13/cobra/cobra@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


clean:
	if [ -f "./build/bin/ops-admin" ]; then rm "./build/bin/ops-admin" ; fi
	if [ -d "./build/openapi" ]; then rm -rf "./build/openapi" ; fi
	if [ -d "./build/tests" ]; then rm -rf "./build/tests"; fi 
	rm -rf ./build/*.out
	rm -rf "./pb"

generate: clean
	if [ ! -d "./build/openapi/v2" ]; then mkdir -p "./build/openapi/v2"; fi
	if [ ! -d "./build/tests" ]; then mkdir -p "./build/tests" ; fi 
	if [ ! -d "./build/bin" ]; then mkdir -p "./build/bin"  ; fi
	if [ ! -d "./pb" ]; then mkdir -p "./pb" ; fi

	protoc -I="ops-spec/proto" \
        --go_out="pb" --go_opt=paths=source_relative \
        --go-grpc_out="pb" --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out="pb" --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out ./build/openapi/v2 \
        --openapiv2_opt logtostderr=true \
        ops-spec/proto/*.proto

test: generate
	gotestsum --format testname \
		--junitfile "./build/tests/unit-tests.xml" \
		-- -coverprofile=build/coverage.out ./...

build: test
	go build -o ./build/bin/ops-admin

install: build
	go install

run: build
	./build/bin/ops-admin start 


