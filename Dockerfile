FROM golang:1.23.8-alpine as build

WORKDIR /build

COPY go.mod ./
RUN go mod download
RUN go mod tidy

COPY ./ ./

RUN mkdir /build/bin
RUN go build -o /build/bin ./...

FROM alpine:3.20.0

ARG VERSION
ENV CYCLOPS_MCP_VERSION=$VERSION
ENV CYCLOPS_MCP_TRANSPORT=sse

WORKDIR /app

RUN mkdir /app/bin

COPY --from=build /build/bin/mcp-cyclops bin/mcp-cyclops

CMD ["bin/mcp-cyclops"]
