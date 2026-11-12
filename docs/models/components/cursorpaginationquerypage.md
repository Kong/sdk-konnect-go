# CursorPaginationQueryPage

Determines which page of the collection to retrieve.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Size`                                                                           | `*int64`                                                                         | :heavy_minus_sign:                                                               | The number of items to include per page.                                         |
| `After`                                                                          | `*string`                                                                        | :heavy_minus_sign:                                                               | Request the next page of data, starting with the item after this parameter.      |
| `Before`                                                                         | `*string`                                                                        | :heavy_minus_sign:                                                               | Request the previous page of data, starting with the item before this parameter. |