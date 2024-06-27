# HashOn

What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing.


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `HashOnNone`       | none               |
| `HashOnConsumer`   | consumer           |
| `HashOnIP`         | ip                 |
| `HashOnHeader`     | header             |
| `HashOnCookie`     | cookie             |
| `HashOnPath`       | path               |
| `HashOnQueryArg`   | query_arg          |
| `HashOnURICapture` | uri_capture        |