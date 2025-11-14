# CreateCustomDomainRequest

Request schema for creating a custom domain in the global API.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ControlPlaneID`                                                                 | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              | 0949471e-b759-45ba-87ab-ee63fb781388                                             |
| `ControlPlaneGeo`                                                                | [components.ControlPlaneGeo](../../models/components/controlplanegeo.md)         | :heavy_check_mark:                                                               | Set of control-plane geos supported for deploying cloud-gateways configurations. |                                                                                  |
| `Domain`                                                                         | *string*                                                                         | :heavy_check_mark:                                                               | Domain name of the custom domain.                                                | example.com                                                                      |