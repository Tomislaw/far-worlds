package main

import (
	"fmt"
	"time"

	"github.com/Tomislaw/far-worlds/component"
	"github.com/Tomislaw/far-worlds/ecs"
	"github.com/Tomislaw/far-worlds/system"
	"github.com/Tomislaw/far-worlds/world/tile"
)

func main() {
	fmt.Println(tile.Atlas.String())

	manager := ecs.NewManager()
	component.RegisterComponents(manager)
	system.RegisterSystems(manager)

	entity := ecs.NewEntity(manager).
		AddComponent(component.NewRandomGUID()).
		AddComponent(component.MapItemMovement{}).
		Register()

	manager.Update(0.1)
	entity.AddComponent(component.MapItem{})
	manager.Update(0.1)
	entity.RemoveComponent(component.Type.GUID)
	manager.Update(0.1)
	entity.Remove()
	manager.Update(0.1)
	go add(manager)
	go add(manager)
	go add(manager)
	go add(manager)
	go add(manager)
	go add(manager)
	entity2 := ecs.NewEntity(manager).
		AddComponent(component.NewRandomGUID()).
		AddComponent(component.MapItemMovement{}).
		Register()
	go add(manager)
	go add(manager)
	manager.Update(0.1)
	go add(manager)
	go add(manager)
	go add(manager)
	go add(manager)
	manager.Update(0.1)
	entity2.Remove()
	manager.Update(0.1)
	time.Sleep(2 * time.Second)
	manager.Update(0.1)
}

func add(manager *ecs.Manager) {
	ecs.NewEntity(manager).
		AddComponent(component.NewRandomGUID()).
		AddComponent(component.MapItemMovement{}).
		Register()
}
