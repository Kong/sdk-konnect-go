# StateMetadata

Metadata describing the state of the managed cache add-on in the data-plane group.



## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ErrorReason`                                                                                        | **string*                                                                                            | :heavy_minus_sign:                                                                                   | Reason why the managed cache add-on may be in an error state, reported from backing infrastructure.<br/> | Failed to create managed cache add-on due to invalid configuration.<br/>                             |