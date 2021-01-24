package system

import "github.com/Tomislaw/far-worlds/ecs"

type MovementSystem struct {
	entites map[uint64]*ecs.Entity
}

func (s *MovementSystem) Priority() int { return 500 }

func (s *MovementSystem) Remove(entity *ecs.Entity) {
	s.entites[entity.ID()] = nil
}

func (s *MovementSystem) Update(dt float32) {

}
