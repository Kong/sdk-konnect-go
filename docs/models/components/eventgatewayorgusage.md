# EventGatewayOrgUsage

Usage counts for a single organization.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `OrgID`                                                                                        | `string`                                                                                       | :heavy_check_mark:                                                                             | Organization ID.                                                                               |
| `Counts`                                                                                       | [components.EventGatewayOrgUsageCounts](../../models/components/eventgatewayorgusagecounts.md) | :heavy_check_mark:                                                                             | Map of usage metric name to its integer count.                                                 |