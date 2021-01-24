package main

import (
	"fmt"

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

}
