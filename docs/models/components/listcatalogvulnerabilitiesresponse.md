# ListCatalogVulnerabilitiesResponse

A paginated list response for a collection of catalog-wide vulnerabilities.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Meta`                                                                               | [components.PaginatedMeta](../../models/components/paginatedmeta.md)                 | :heavy_check_mark:                                                                   | returns the pagination information                                                   |
| `Data`                                                                               | [][components.CatalogVulnerability](../../models/components/catalogvulnerability.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |