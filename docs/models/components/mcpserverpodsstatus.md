# MCPServerPodsStatus


## Fields

| Field                                 | Type                                  | Required                              | Description                           | Example                               |
| ------------------------------------- | ------------------------------------- | ------------------------------------- | ------------------------------------- | ------------------------------------- |
| `Ready`                               | `int64`                               | :heavy_check_mark:                    | Number of pods in the ready state.    | 2                                     |
| `Starting`                            | `int64`                               | :heavy_check_mark:                    | Number of pods in the starting state. | 0                                     |
| `Failing`                             | `int64`                               | :heavy_check_mark:                    | Number of pods in the failing state.  | 0                                     |