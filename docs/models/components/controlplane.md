# ControlPlane

Control Plane is the owner for the add-on.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Kind`                                                                           | *string*                                                                         | :heavy_check_mark:                                                               | Type of owner for the add-on.                                                    |                                                                                  |
| `ControlPlaneID`                                                                 | *string*                                                                         | :heavy_check_mark:                                                               | ID of the control-plane that owns this add-on.                                   | 123e4567-e89b-12d3-a456-426614174000                                             |
| `ControlPlaneGeo`                                                                | [components.ControlPlaneGeo](../../models/components/controlplanegeo.md)         | :heavy_check_mark:                                                               | Set of control-plane geos supported for deploying cloud-gateways configurations. |                                                                                  |