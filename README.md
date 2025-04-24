# Cyclops MCP (Model Context Protocol)

<p align="center" width="100%">
    <img width="75%" src="https://raw.githubusercontent.com/cyclops-ui/cyclops/main/web/static/img/cyclops-simplistic.png">
<p/>

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

> ⚠️ By default, Cyclops MCP will use the `.kube/config` file to connect to your cluster

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

| Tool                  | Description                                                                                                                        |
|-----------------------|------------------------------------------------------------------------------------------------------------------------------------|
| `create_module`       | Create new Module. Before calling this tool, make sure to call `get_template_schema` to validate values for the given template     |
| `get_module`          | Fetch Module by name                                                                                                               |
| `list_modules`        | List all Cyclops Modules                                                                                                           |
| `update_module`       | Update Module by Name. Before calling this tool, make sure to call `get_template_schema` to validate values for the given template |
| `get_template_schema` | Returns JSON schema for the given template. Needs to be checked before calling `create_module` tool                                |
| `get_template_store`  | Fetch Template Store by Name                                                                                                       |
| `list_template_store` | List Template Stores from cluster                                                                                                  |

## Configuration

You can configure Cyclops MCP server via env variables. Below is an example of adding the configuration for specifying the kubeconfig file the Cyclops MCP server should use when managing your Cyclops applications.

```json
{
  "mcpServers": {
    "mcp-cyclops": {
      "command": "mcp-cyclops",
      "env": {
        "KUBECONFIG": "/path/to/your/kubeconfig"
      }
    }
  }
}

```

### Environment variables

Below is the list of environment variables used for configuring your Cyclops MCP server:

| Env var                           | Description                                                                             |
|-----------------------------------|-----------------------------------------------------------------------------------------|
| `KUBECONFIG`                      | Path to kubeconfig file (optional, defaults to in-cluster config or $HOME/.kube/config) |
| `CYCLOPS_KUBE_CONTEXT`            | Kubernetes context to use (optional)                                                    |
| `CYCLOPS_MODULE_NAMESPACE`        | Namespace where modules are stored                                                      |
| `CYCLOPS_HELM_RELEASE_NAMESPACE`  | Namespace for Helm releases                                                             |
| `CYCLOPS_MODULE_TARGET_NAMESPACE` | Target namespace for modules                                                            |
