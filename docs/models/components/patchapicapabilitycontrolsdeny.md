# PatchAPICapabilityControlsDeny

The capabilities of the mapped API to deny. Omitted keys preserve their current value, an empty array clears the capability, and a non-empty array replaces.



## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `Operations`                                                                                             | [][components.APICapabilityControlsOperation](../../models/components/apicapabilitycontrolsoperation.md) | :heavy_minus_sign:                                                                                       | The API operations that are denied.                                                                      |