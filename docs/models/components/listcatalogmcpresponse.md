# ListCatalogMCPResponse

Response schema for a paginated list of MCPs.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Meta`                                                               | [components.PaginatedMeta](../../models/components/paginatedmeta.md) | :heavy_check_mark:                                                   | returns the pagination information                                   |
| `Data`                                                               | [][components.CatalogMCP](../../models/components/catalogmcp.md)     | :heavy_check_mark:                                                   | N/A                                                                  |