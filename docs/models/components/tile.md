# Tile


## Supported Types

### ChartTile

```go
tile := components.CreateTileChart(components.ChartTile{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch tile.Type {
	case components.TileTypeChart:
		// tile.ChartTile is populated
}
```
