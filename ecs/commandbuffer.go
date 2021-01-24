package ecs

import (
	"reflect"
	"sync"
)

type commandBuffer struct {
	componentsToAdd     sync.Pool // component
	componentsToDestroy sync.Pool // component

	entitesToAdd     sync.Pool // *Entity
	entitesToDestroy sync.Pool // uint64

}

type component struct {
	reflectType reflect.Type
	entityID    uint64
	data        *interface{}
}

func (b *commandBuffer) removeEntity(e *Entity) {
	b.entitesToDestroy.Put(e.id)
	for _, item := range e.GetComponentsTypes() {
		b.removeComponent(e, item)
	}
}

func (b *commandBuffer) addEntity(e *Entity) {
	b.entitesToAdd.Put(e)
}

func (b *commandBuffer) removeComponent(e *Entity, t reflect.Type) {
	b.componentsToDestroy.Put(component{reflectType: t, entityID: e.id})
}

func (b *commandBuffer) addComponent(e *Entity, c *interface{}) {
	b.entitesToAdd.Put(component{reflectType: reflect.TypeOf(c), entityID: e.id, data: c})
}

func (b *commandBuffer) resolveEntites(entites entites) (resolvedEntites entites) {

	for {
		entity := b.entitesToAdd.Get().(*Entity)
		if entity == nil {
			break
		}

		entites = append(resolvedEntites, entity)
	}
	return
}

func (b *commandBuffer) resolveComponents(entites entites) (resolvedEntites entites) {

	for {
		entity := b.entitesToAdd.Get().(*Entity)
		if entity == nil {
			break
		}

		entites = append(resolvedEntites, entity)
	}
	return
}
