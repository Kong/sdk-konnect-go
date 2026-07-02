# GovernanceQueryRequestFeature

Optional list of feature keys to evaluate access for. If omitted, all features
available in the organization are returned. Providing this list is recommended
to reduce the response size and the load on the backend services.


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `Keys`                                       | []`string`                                   | :heavy_check_mark:                           | List of feature keys to evaluate access for. |