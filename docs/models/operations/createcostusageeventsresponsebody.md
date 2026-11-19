# CreateCostUsageEventsResponseBody

Every event was accepted for storage. The body lists each event's `id` and `source`, in request order.


## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `Data`                                                                                                               | [][operations.CreateCostUsageEventsUsageEventsData](../../models/operations/createcostusageeventsusageeventsdata.md) | :heavy_check_mark:                                                                                                   | The `id` and `source` of each accepted event, in request order. Together they identify the event once it's stored.   |