package ecs

import (
	"fmt"
	"reflect"
	"sync"
)

type commandBuffer struct {
	componentsToAdd     sync.Pool // component
	componentsToDestroy sync.Pool // component

	entitesToAdd     sync.Pool // *Entity
	entitesToDestroy sync.Pool // *Entity

}

type component struct {
	reflectType reflect.Type
	entityID    uint64
	data        *interface{}
}

func (b *commandBuffer) removeEntity(e *Entity) {
	b.entitesToDestroy.Put(e)
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

	ctype := reflect.TypeOf(reflect.ValueOf(c).Elem().Interface())
	b.componentsToAdd.Put(component{reflectType: ctype, entityID: e.id, data: c})
}

func (b *commandBuffer) resolveEntites(entites entites) {

	for {
		e := b.entitesToAdd.Get()
		if e == nil {
			break
		}
		entityToAdd := e.(*Entity)
		entites[entityToAdd.id] = entityToAdd
		fmt.Printf("Entity added - id: %v\n", entityToAdd.id)
	}

	for {
		e := b.entitesToDestroy.Get()
		if e == nil {
			break
		}
		entityToRemove := e.(*Entity)
		entites[entityToRemove.id] = nil
		fmt.Printf("Entity removed - id: %v\n", entityToRemove.id)
	}
	return
}

func (b *commandBuffer) resolveComponents(components components) {

	for {
		c := b.componentsToAdd.Get()
		if c == nil {
			break
		}
		componentToAdd := c.(component)
		componentData := components[componentToAdd.reflectType]
		componentData.data[componentToAdd.entityID] = componentToAdd.data
		fmt.Printf("Component added to entity: %v, type: %v\n", componentToAdd.entityID, componentToAdd.reflectType)
		//componentdata.
		//.data[componentToAdd.entityID] = componentToAdd.data
	}

	for {
		c := b.componentsToDestroy.Get()
		if c == nil {
			break
		}
		componentToRemove := c.(component)
		components[componentToRemove.reflectType].data[componentToRemove.entityID] = nil
		fmt.Printf("Component removed - entity: %v, type: %v\n", componentToRemove.entityID, componentToRemove.reflectType)
	}
	return
}
