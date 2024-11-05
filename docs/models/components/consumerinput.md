# ConsumerInput

The Consumer object represents a consumer - or a user - of a Service. You can either rely on Kong as the primary datastore, or you can map the consumer list with your database to keep consistency between Kong and your existing primary datastore.


## Fields

| Field                                                                                                                                                                                    | Type                                                                                                                                                                                     | Required                                                                                                                                                                                 | Description                                                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CustomID`                                                                                                                                                                               | **string*                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                       | Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database. You must send either this field or `username` with the request. |
| `ID`                                                                                                                                                                                     | **string*                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                       | N/A                                                                                                                                                                                      |
| `Tags`                                                                                                                                                                                   | []*string*                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                       | An optional set of strings associated with the Consumer for grouping and filtering.                                                                                                      |
| `Username`                                                                                                                                                                               | **string*                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                       | The unique username of the Consumer. You must send either this field or `custom_id` with the request.                                                                                    |