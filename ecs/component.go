package ecs

import (
	"reflect"
	"sync"
)

// Component contains component type and data for each entity of this component type
// Maxiumum component count is 64
type Component struct {
	id       uint8
	datalock *sync.RWMutex
	data     map[uint64]*interface{}
}

type components map[reflect.Type]*Component

func (s components) Len() int {
	return len(s)
}

// AddComponent adds any object to entity
func (entity *Entity) AddComponent(component interface{}) *Entity {
	ctype := entity.manager.GetComponentType(reflect.TypeOf(component))

	if ctype == nil {
		return entity
	}

	ctype.datalock.Lock()
	ctype.data[entity.id] = &component
	entity.componentFlags |= 1 << ctype.id
	ctype.datalock.Unlock()
	return entity
}

// AddComponents adds multiple structs to entity
func (entity *Entity) AddComponents(components ...interface{}) *Entity {

	for _, component := range components {
		entity.AddComponent(component)
	}
	return entity
}

// GetComponent return component assignet to this entity or nil if not found
func (entity *Entity) GetComponent(componentType reflect.Type) interface{} {

	ok := entity.HasComponent(componentType)
	if !ok {
		return nil
	}

	component, ok := entity.manager.components[componentType]
	if !ok {
		return nil
	}

	data, ok := component.data[entity.id]
	if !ok {
		return nil
	}
	return data
}

func (entity *Entity) GetComponents() []*interface{} {

	var components []*interface{}

	for _, component := range entity.manager.components {

		component.datalock.RLock()
		item, ok := component.data[entity.id]

		if !ok {
			component.datalock.RUnlock()
			continue
		}
		components = append(components, item)
	}
	return components
}

func (entity *Entity) GetComponentsTypes() []reflect.Type {

	var components []reflect.Type

	for ctype, component := range entity.manager.components {

		component.datalock.RLock()
		_, ok := component.data[entity.id]

		if !ok {
			component.datalock.RUnlock()
			continue
		}
		components = append(components, ctype)
	}
	return components
}

// HasComponent returns true if contains component
func (entity *Entity) HasComponent(componentType reflect.Type) bool {
	ctype := entity.manager.GetComponentType(componentType)
	if ctype == nil {
		return false
	}
	id := ctype.id
	var flag uint64 = 1 << id
	return entity.componentFlags&flag == flag
}
