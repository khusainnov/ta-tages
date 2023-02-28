API_PATH=api/tages
PROTO_OUT_DIR=pkg/tages-api
PROTO_API_DIR=$(API_PATH)

.PHONY: gen
gen: gen-proto generate


.PHONY: gen-proto
gen-proto:
	mkdir -p $(PROTO_OUT_DIR)
	protoc \
		-I $(API_PATH)/v1 \
		-I third_party/googleapis \
		--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT_DIR)  --go-grpc_opt=paths=source_relative \
        ./$(PROTO_API_DIR)/v1/*.proto

test:
	go test ./... -cover -count=1

generate:
	go generate ./...
run:
	go run cmd/tages/main.go
