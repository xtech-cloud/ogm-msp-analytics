.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/analytics/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/analytics/record.proto