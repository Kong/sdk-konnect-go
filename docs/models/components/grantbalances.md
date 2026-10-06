# GrantBalances

The balance of each active grant at the start and at the end of the segment,
keyed by grant ID.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Start`                                                       | map[string]`string`                                           | :heavy_check_mark:                                            | The balance of each active grant at the start of the segment. | {<br/>"01G65Z755AFWAKHE12NY0CQ9FH": "100"<br/>}               |
| `End`                                                         | map[string]`string`                                           | :heavy_check_mark:                                            | The balance of each active grant at the end of the segment.   | {<br/>"01G65Z755AFWAKHE12NY0CQ9FH": "50"<br/>}                |