# CreateFeatureRequestMeterReference

The meter that the feature is associated with and based on which usage is
calculated. If not specified, the feature is static.


## Fields

| Field                                                                                             | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ID`                                                                                              | `string`                                                                                          | :heavy_check_mark:                                                                                | The ID of the meter to associate with this feature.                                               | 01G65Z755AFWAKHE12NY0CQ9FH                                                                        |
| `Filters`                                                                                         | map[string][metering.QueryFilterStringMapItem](../../models/metering/queryfilterstringmapitem.md) | :heavy_minus_sign:                                                                                | Filters to apply to the dimensions of the meter.                                                  |                                                                                                   |