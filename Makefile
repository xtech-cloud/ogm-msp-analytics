.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/analytics/healthy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/analytics/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/analytics/tracker.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/analytics/generator.proto
