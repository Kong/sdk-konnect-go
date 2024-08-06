# TeamMapping

A team assignment is a mapping of an IdP group to a Konnect Team.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Group`                                                  | **string*                                                | :heavy_minus_sign:                                       | The IdP group.                                           | Service Developers                                       |
| `TeamIds`                                                | []*string*                                               | :heavy_minus_sign:                                       | An array of ID's that are mapped to the specified group. |                                                          |