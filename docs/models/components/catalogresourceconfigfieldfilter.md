# CatalogResourceConfigFieldFilter

Filters on the resource's `config` field. Filters must use dot-notation to identify
the property that will be used to filter the results. For example:





  - `filter[config.control_plane_id]`
  - `filter[config.gateway_service_id][neq]=cedbd134-fc5c-4d44-9e38-ccd2a0c6e0ae`



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Eq`               | **string*          | :heavy_minus_sign: | N/A                |
| `Contains`         | *string*           | :heavy_check_mark: | N/A                |
| `Ocontains`        | *string*           | :heavy_check_mark: | N/A                |
| `Oeq`              | *string*           | :heavy_check_mark: | N/A                |
| `Neq`              | *string*           | :heavy_check_mark: | N/A                |