package world

import "github.com/Tomislaw/far-worlds/world/tile"

const chunkWidth = 64
const chunkHeight = 8

type Chunk struct {
	tiles [chunkWidth][chunkWidth][chunkHeight]uint8 `json:"id"`
}

func (ch Chunk) GetTile(x uint8, y uint8, z uint8) tile.Tile {
	return tile.Atlas.Tiles[ch.tiles[x][y][z]]
}
