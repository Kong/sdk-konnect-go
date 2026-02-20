# CatalogServiceAPIMapping

Represents an API mapping associated with a service in the catalog.


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `ID`                                      | *string*                                  | :heavy_check_mark:                        | The mapping ID.                           |
| `ServiceID`                               | *string*                                  | :heavy_check_mark:                        | The ID of the service.                    |
| `APIID`                                   | *string*                                  | :heavy_check_mark:                        | The ID of the API.                        |
| `CreatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | Creation timestamp.                       |
| `UpdatedAt`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | Last update timestamp.                    |