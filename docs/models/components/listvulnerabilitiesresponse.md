# ListVulnerabilitiesResponse

A paginated list response for a collection of vulnerabilities.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Meta`                                                                 | [components.PaginatedMeta](../../models/components/paginatedmeta.md)   | :heavy_check_mark:                                                     | returns the pagination information                                     |
| `Data`                                                                 | [][components.Vulnerability](../../models/components/vulnerability.md) | :heavy_check_mark:                                                     | N/A                                                                    |