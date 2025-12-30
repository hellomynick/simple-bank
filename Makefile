gen: proto sqlc

proto:
	buf generate

sqlc:
	sqlc generate

.PHONY: gen proto sqlc