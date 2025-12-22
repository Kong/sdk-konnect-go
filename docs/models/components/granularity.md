# Granularity

Force time grouping into buckets of the specified duration. Only has an effect if `time` is in the `dimensions` list.

The granularity of the result may be coarser than requested.
The exact result granularity will be reported in the response `meta.granularity_ms` field.

If granularity is not specified and `time` is in the `dimensions` list, a default will be chosen based on the time range requested.

Different relative times support different granularities:




  - 1d  => daily
  - 7d  => daily, weekly
  - 3w  => daily, weekly
  - 30d => daily, weekly
  - 3mo => daily, weekly
  - 1y  => daily, weekly

For special time ranges:




  - current_week, previous_week   => daily
  - current_month, previous_month => daily, weekly

For absolute time ranges, daily will be used.



## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `GranularityDaily`  | daily               |
| `GranularityWeekly` | weekly              |