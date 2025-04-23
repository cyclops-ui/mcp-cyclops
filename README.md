# Cyclops MCP

MCP (Model Context Protocol) server implementation for Cyclops resources

---

Cyclops MCP allows your favorite AI agent to manage your Kubernetes applications. Cyclops MCP servers provide tools for agents to create and update existing applications safely.

This means it can check all of your existing templates and the schema of those templates to create accurate and production-ready applications. Your agent now has much less room to make a misconfiguration since it creates high-level resources (Cyclops Modules) instead of touching every line of your Kubernetes resources (Deployments, Services, and Ingresses).

It allows you to move fast and ensure no uncaught misconfigurations are hitting your production.

**With Cyclops and our MCP, you can now abstract Kubernetes complexity from your developers AND your AI agents**

## Install

### 1. Make sure Cyclops is installed in your Kubernetes cluster

Check our docs on how it install it with a single command - https://cyclops-ui.com/docs/installation/install/manifest

### 2. Download MCP server

You can download the Cyclops MCP server binary with the following command:

```yaml
GOBIN="$HOME/go/bin" go install github.com/cyclops-ui/mcp-cyclops/cmd/mcp-cyclops@latest
```

### 3. Add server configuration

<aside>
⚠️

By default, Cyclops MCP will use the `.kube/config` file to connect to your cluster

</aside>

Configure your MCP Cyclops server:

```json
{
  "mcpServers": {
    "mcp-cyclops": {
      "command": "mcp-cyclops"
    }
  }
}
```

---

## Tools

| `create_module` | Create new Module. Before calling this tool, make sure to call `get_template_schema` to validate values for the given template |
| --- | --- |
| `get_module` | Fetch Module by name |
| `list_modules` | List all Cyclops Modules |
| `update_module` | Update Module by Name. Before calling this tool, make sure to call `get_template_schema` to validate values for the given template |
| `get_template_schema` | Returns JSON schema for the given template. Needs to be checked before calling `create_module` tool |
| `get_template_store`  | Fetch Template Store by Name |
| `list_template_store` | List Template Stores from cluster |