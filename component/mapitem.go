package component

type MapItem struct {
	tileX uint8
	tileY uint8
	tileZ uint8

	chunkX uint8
	chunkY uint8

	mapId uint8
}

type MapItemBlock struct {
	sizeX uint8
	sizeY uint8
	sizeZ uint8
}

type MapItemMovement struct {
	moveX    uint8
	moveY    uint8
	moveZ    uint8
	progress float32
}
