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

// GetComponentType returns component type
func (w *Manager) GetComponentType(componentType reflect.Type) *Component {

	component, ok := w.components[componentType]

	if ok {
		return component
	}
	return nil
}

// RegisterComponentsType registers type of struct which can be later added later to entities
// Can't be used async
func (w *Manager) RegisterComponentsType(componentType reflect.Type) *Manager {

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
	}

	w.components[componentType] = newType
	return w
}

// RegisterComponentsTypes todo
func (w *Manager) RegisterComponentsTypes(componentTypes ...reflect.Type) *Manager {

	for _, component := range componentTypes {
		w.RegisterComponentsType(component)
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
	for _, system := range w.Systems() {
		system.Update(dt)
	}
}

// RemoveEntity removes the entity across all systems.
func (w *Manager) RemoveEntity(e Entity) {
	for _, sys := range w.systems {
		sys.Remove(e)
	}
}
