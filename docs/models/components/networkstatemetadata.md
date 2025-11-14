# NetworkStateMetadata

Metadata describing the backing state of the network and why it may be in an erroneous state.



## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ReportedStatus`                                                                            | **string*                                                                                   | :heavy_minus_sign:                                                                          | Reported status of the network from backing infrastructure.                                 | INVALID                                                                                     |
| `Reason`                                                                                    | **string*                                                                                   | :heavy_minus_sign:                                                                          | Reason why the network may be in an erroneous state, reported from backing infrastructure.<br/> | Network could not be deployed due to insufficient cloud provider compute instances.<br/>    |