# MCPServerSignalV1

Signal notifying the client that the MCP control plane config has changed.



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Type`                                                                               | [components.MCPServerSignalV1Type](../../models/components/mcpserversignalv1type.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `Version`                                                                            | `string`                                                                             | :heavy_check_mark:                                                                   | The current version of the MCP control plane.                                        |
| `Offset`                                                                             | `string`                                                                             | :heavy_check_mark:                                                                   | The offset for the MCP signal.                                                       |