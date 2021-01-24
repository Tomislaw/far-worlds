package ecs

import (
	"reflect"
	"sort"
	"sync"
)

// Manager contains a bunch of Entities, and a bunch of Systems. It is the
// recommended way to run ecs.

const maxComponents = 64

// Manager ...
type Manager struct {
	commandBuffer commandBuffer

	systems    systems
	components components
	entites    entites
}

func NewManager() *Manager {
	return &Manager{
		systems:    make([]System, 0),
		components: make(map[reflect.Type]*Component, 64),
		entites:    make(map[uint64]*Entity, 1000),
	}
}

// GetComponentType returns component type
func (w *Manager) GetComponentType(componentType reflect.Type) *Component {

	component, ok := w.components[componentType]

	if ok {
		return component
	}
	return nil
}

// RegisterComponent registers type of struct which can be later added later to entities
// Can't be used async
func (w *Manager) RegisterComponent(componentType reflect.Type) *Manager {

	_, ok := w.components[componentType]

	if ok {
		return w
	}

	if w.components.Len() > maxComponents {
		panic("Too many component types!")
	}

	newType := &Component{
		id:       uint8(w.components.Len()),
		datalock: &sync.RWMutex{},
		data:     make(map[uint64]*interface{}),
	}

	w.components[componentType] = newType
	return w
}

// RegisterComponents todo
func (w *Manager) RegisterComponents(componentTypes ...reflect.Type) *Manager {

	for _, component := range componentTypes {
		w.RegisterComponent(component)
	}
	return w
}

// RegisterSystem todo
func (w *Manager) RegisterSystem(system System) *Manager {
	if initializer, ok := system.(Initializer); ok {
		initializer.New(w)
	}

	w.systems = append(w.systems, system)
	sort.Sort(w.systems)
	return w
}

// RegisterSystems todo
func (w *Manager) RegisterSystems(systems ...System) *Manager {

	for _, system := range systems {
		if initializer, ok := system.(Initializer); ok {
			initializer.New(w)
		}

		w.systems = append(w.systems, system)
	}
	sort.Sort(w.systems)
	return w
}

// Systems returns the list of Systems managed by the World.
func (w *Manager) Systems() []System {
	return w.systems
}

// Update updates each System managed by the World. It is invoked by the engine
// once every frame, with dt being the duration since the previous update.
func (w *Manager) Update(dt float32) {
	w.commandBuffer.resolveComponents(w.components)
	w.commandBuffer.resolveEntites(w.entites)
	for _, system := range w.Systems() {
		system.Update(dt)
	}
}

// RemoveEntity removes the entity across all systems.
func (w *Manager) RemoveEntity(e *Entity) {
	w.commandBuffer.removeEntity(e)
}

func (w *Manager) RemoveEntityWithId(id uint64) {
	entity, ok := w.entites[id]
	if ok {
		w.commandBuffer.removeEntity(entity)
	}
}
