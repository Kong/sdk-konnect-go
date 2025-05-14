# StateMetadata

Metadata describing the backing state of the dataplane group and why it may be in an erroneous state.



## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ReportedStatus`                                                                                    | **string*                                                                                           | :heavy_minus_sign:                                                                                  | Reported status of the dataplane group from backing infrastructure.                                 | ERROR                                                                                               |
| `Reason`                                                                                            | **string*                                                                                           | :heavy_minus_sign:                                                                                  | Reason why the dataplane group may be in an erroneous state, reported from backing infrastructure.<br/> | Dataplane group could not be deployed due to insufficient cloud provider compute instances.<br/>    |