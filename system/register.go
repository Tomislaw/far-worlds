package system

import "github.com/Tomislaw/far-worlds/ecs"

func RegisterSystems(manager *ecs.Manager) {
	manager.RegisterSystems(&MovementSystem{})
}
