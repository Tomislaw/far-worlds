package tile

type Tile struct {
	Id         uint16 `json:"id"`
	MaterialID uint8  `json:"material"`
	Name       string `json:"name"`
	Block      bool   `json:"block"`
}
