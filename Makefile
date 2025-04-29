start:
	CYCLOPS_MCP_TRANSPORT=sse go run -tags musl ./cmd/mcp-cyclops

build:
	go build -o bin ./...

group-imports:
	goimports -w .

unit-test:
	go test -v -tags musl,dynamic ./...
