# AwsPrivateDNSResolverResponsePrivateDNSStateMetadata

Metadata describing the backing state of the Private Dns and why it may be in an erroneous state.



## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     | Example                                                                                         |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ReportedStatus`                                                                                | **string*                                                                                       | :heavy_minus_sign:                                                                              | Reported status of the Private Dns from backing infrastructure.                                 | ERROR                                                                                           |
| `Reason`                                                                                        | **string*                                                                                       | :heavy_minus_sign:                                                                              | Reason why the Private Dns may be in an erroneous state, reported from backing infrastructure.<br/> | Failed to create Private Dns due to invalid Cloud Provider configuration.<br/>                  |