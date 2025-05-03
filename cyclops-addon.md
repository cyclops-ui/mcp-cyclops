# Cyclops MCP (Model Context Protocol)

Cyclops MCP allows your favorite AI agent to manage your Kubernetes applications. Cyclops MCP servers provide tools for agents to create and update existing applications safely.

This means it can check all of your existing templates and the schema of those templates to create accurate and production-ready applications. Your agent now has much less room to make a misconfiguration since it creates high-level resources (Cyclops Modules) instead of touching every line of your Kubernetes resources (Deployments, Services, and Ingresses).

It allows you to move fast and ensure no uncaught misconfigurations are hitting your production.

**With Cyclops and our MCP, you can now abstract Kubernetes complexity from your developers AND your AI agents**

## Install

1. Click `Install Cyclops MCP server` above which will install the MCP server to your cluster.

2. You can now expose the `cyclops-mcp` service. To test your MCP server, you can port-forward it:

    ```shell
    kubectl port-forward svc/cyclops-mcp -n cyclops 8000:8000
    ```

3. Add your Cyclops MCP server host, or in case you are testing it, the [http://localhost:8000/sse](http://localhost:8000/sse) address where you port-forwarded the MCP service:

    ```json
    {
      "mcpServers": {
        "mcp-cyclops": {
          "url": "http://localhost:8000/sse"
        }
      }
    }
    ```

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
