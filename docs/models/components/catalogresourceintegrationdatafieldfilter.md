# CatalogResourceIntegrationDataFieldFilter

Filters on the resource's `integration_data` field. Filters must use dot-notation to identify
the property that will be used to filter the results. For example:


  - `filter[integration_data.control_plane.id]=760562c3-fedd-43e4-b742-fc1663fef188`
  - `filter[integration_data.gateway_service.name][contains]=legacy`



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Eq`               | **string*          | :heavy_minus_sign: | N/A                |
| `Contains`         | *string*           | :heavy_check_mark: | N/A                |
| `Ocontains`        | *string*           | :heavy_check_mark: | N/A                |
| `Oeq`              | *string*           | :heavy_check_mark: | N/A                |
| `Neq`              | *string*           | :heavy_check_mark: | N/A                |