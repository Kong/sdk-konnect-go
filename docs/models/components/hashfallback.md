# HashFallback

What to use as hashing input if the primary `hash_on` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hash_on` is set to `cookie`.


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `HashFallbackNone`       | none                     |
| `HashFallbackConsumer`   | consumer                 |
| `HashFallbackIP`         | ip                       |
| `HashFallbackHeader`     | header                   |
| `HashFallbackCookie`     | cookie                   |
| `HashFallbackPath`       | path                     |
| `HashFallbackQueryArg`   | query_arg                |
| `HashFallbackURICapture` | uri_capture              |