# Layout

Information about how the tile is placed on the dashboard.

Examples:
- a tile occupying the first half of the top row: `{ "position": { "col": 0, "row": 0 }, size: { "cols": 3, "rows": 1 } }`
- a tile occupying the second half of the top row: `{ "position": { "col": 3, "row": 0 }, size: { "cols": 3, "rows": 1 } }`



## Fields

| Field                                                                                                                                            | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Position`                                                                                                                                       | [components.Position](../../models/components/position.md)                                                                                       | :heavy_check_mark:                                                                                                                               | Position of the tile in the dashboard's grid.  Numbering starts at 0, so a tile in the upper left of the dashboard will be at column 0, row 0.<br/> |
| `Size`                                                                                                                                           | [components.Size](../../models/components/size.md)                                                                                               | :heavy_check_mark:                                                                                                                               | Number of columns and rows the tile occupies.  A dashboard always has 6 columns, but has as many rows as needed to accommodate the given tiles.<br/> |