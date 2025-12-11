# ListCatalogServiceDocumentsResponseBody

A paginated list response for a collection of documents targeted by a service.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Meta`                                                                     | [components.PaginatedMeta](../../models/components/paginatedmeta.md)       | :heavy_check_mark:                                                         | returns the pagination information                                         |
| `Data`                                                                     | [][components.ServiceDocument](../../models/components/servicedocument.md) | :heavy_check_mark:                                                         | N/A                                                                        |