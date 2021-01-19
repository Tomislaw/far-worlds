package ecs

import (
	"reflect"
	"sync"
)

type Component struct {
	id       uint8
	tag      Tag
	datalock *sync.RWMutex
	data     map[uint64]interface{}
}

type components map[reflect.Type]*Component

func (s components) Len() int {
	return len(s)
}

// AddComponent adds any object to entity
func (entity *Entity) AddComponent(component interface{}) *Entity {
	entity.manager.AddComponent(entity, component)
	return entity
}

func (entity *Entity) GetComponent(componentType reflect.Type) interface{} {
	return entity.manager.GetComponent(entity, componentType)
}

func (entity *Entity) HasComponent(componentType reflect.Type) bool {
	return entity.manager.HasComponent(entity, componentType)
}
