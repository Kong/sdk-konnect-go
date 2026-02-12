# Mode

* hide_prefix - the configured prefix is hidden from clients for topics and IDs when reading.


  Created resources are written with the prefix on the backend cluster.
* enforce_prefix - the configured prefix remains visible to clients.


  Created resources must include the prefix or the request will fail.



## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `ModeHidePrefix`    | hide_prefix         |
| `ModeEnforcePrefix` | enforce_prefix      |