.PHONY: generate
generate:
	go generate ./...

.PHONY: run
run:
	go run server.go

.PHONY: test
test: gen_json_query
	@curl -s -X POST -H "Content-Type: application/json" -d @./test/query.json localhost:8080/query | yq -P

gen_json_query:
	@echo "{\"query\":\"$$(cat ./test/query.gql)\"}" | tr '\n' ' ' > ./test/query.json
