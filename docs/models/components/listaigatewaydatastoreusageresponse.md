# ListAIGatewayDatastoreUsageResponse

A paginated list of the policies and models that reference an AI Gateway Datastore.


## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `Data`                                                                                               | [][components.AIGatewayDatastoreUsageEntry](../../models/components/aigatewaydatastoreusageentry.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `Meta`                                                                                               | [components.CursorMetaWithTotal](../../models/components/cursormetawithtotal.md)                     | :heavy_check_mark:                                                                                   | Pagination metadata with exact total. Useful when the collection size is inexpensive to provide.     |