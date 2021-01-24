package world

import (
	"time"

	"github.com/Tomislaw/far-worlds/component"
	"github.com/Tomislaw/far-worlds/ecs"
	"github.com/Tomislaw/far-worlds/system"
)

const mapWidth = 8

type Map struct {
	globalChunkManager GlobalChunksManager
	chunk              [mapWidth][mapWidth]Chunk
	manager            ecs.Manager

	mainLoopActive bool

	time int64
}

func (m *Map) GetChunk(x uint16, y uint16) *Chunk {
	return &m.chunk[x][y]
}

type GlobalChunksManager interface {
	GetChunk(x uint16, y uint16) *Chunk
}

func LoadMap() Map {
	m := Map{}

	component.RegisterComponents(&m.manager)
	system.RegisterSystems(&m.manager)

	return m
}

func (m *Map) Update(dt float32) {
	m.manager.Update(dt)
}

func (m *Map) StartMainLoop() {
	m.mainLoopActive = true
	m.time = time.Now().Local().UnixNano()
	go m.mainLoop()

}

func (m *Map) StopMainLoop() {
	m.mainLoopActive = false
}

func (m *Map) mainLoop() {
	for m.mainLoopActive {
		newTime := time.Now().Local().UnixNano()
		dt := newTime - m.time
		m.manager.Update(float32(dt) / 1000000000)
		m.time = newTime
	}
}
