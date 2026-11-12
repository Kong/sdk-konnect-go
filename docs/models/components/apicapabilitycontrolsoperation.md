# APICapabilityControlsOperation

An API operation that is denied.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Path`                                                 | `string`                                               | :heavy_check_mark:                                     | The path of the API operation.                         | /library/{id}/books                                    |
| `Methods`                                              | []`string`                                             | :heavy_check_mark:                                     | The HTTP methods of the API operation that are denied. | [<br/>"GET"<br/>]                                      |