# Service

Service entities, as the name implies, are abstractions of each of your own upstream services. Examples of Services would be a data transformation microservice, a billing API, etc. The main attribute of a Service is its URL (where Kong should proxy traffic to), which can be set as a single string or by specifying its `protocol`, `host`, `port` and `path` individually. Services are associated to Routes (a Service can have many Routes associated with it). Routes are entry-points in Kong and define rules to match client requests. Once a Route is matched, Kong proxies the request to its associated Service. See the [Proxy Reference][proxy-reference] for a detailed explanation of how Kong proxies traffic.


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `CreatedAt`                                    | *int64*                                        | :heavy_check_mark:                             | Unix epoch when the resource was created.      |
| `ID`                                           | *string*                                       | :heavy_check_mark:                             | N/A                                            |
| `UpdatedAt`                                    | *int64*                                        | :heavy_check_mark:                             | Unix epoch when the resource was last updated. |