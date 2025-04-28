start:
	CYCLOPS_MCP_TRANSPORT=sse go run -tags musl ./cmd/mcp-cyclops

group-imports:
	goimports -w .

unit-test:
	go test -v -tags musl,dynamic ./...
