package world

type World struct {
	maps []Map
}

func (world *World) Init() {

	for id := range world.maps {
		go world.maps[id].StartMainLoop()
	}

}
