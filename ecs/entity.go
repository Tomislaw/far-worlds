package ecs

import (
	"sync/atomic"
)

var (
	idInc uint64
)

// A Entity is simply a set of components with a unique ID attached to it,
// nothing more. It belongs to any amount of Systems, and has a number of
// Components
type Entity struct {
	// Entity ID.
	id             uint64
	componentFlags uint64
	parent         *Entity
	children       []*Entity
	manager        *Manager
}

// Identifier is an interface for anything that implements the basic ID() uint64,
// as the BasicEntity does.  It is useful as more specific interface for an
// entity registry than just the interface{} interface
type Identifier interface {
	ID() uint64
}

// IdentifierSlice implements the sort.Interface, so you can use the
// store entites in slices, and use the P=n*log n lookup for them
type IdentifierSlice []Identifier

// NewEntity creates a new Entity with a new unique identifier. It is safe for
// concurrent use.
func NewEntity(manager *Manager) *Entity {
	return &Entity{id: atomic.AddUint64(&idInc, 1), manager: manager}
}

// NewEntities creates an amount of new entities with a new unique identifiers. It
// is safe for concurrent use, and performs better than NewBasic for large
// numbers of entities.
func NewEntities(manager *Manager, amount int) []*Entity {
	entities := make([]*Entity, amount)

	lastID := atomic.AddUint64(&idInc, uint64(amount))
	for i := 0; i < amount; i++ {
		entities[i] = &Entity{}
		entities[i].id = lastID - uint64(amount) + uint64(i) + 1
		entities[i].manager = manager
	}

	return entities
}

// ID returns the unique identifier of the entity.
func (e Entity) ID() uint64 {
	return e.id
}

// GetEntity returns a Pointer to the BasicEntity itself
// By having this method, All Entities containing a BasicEntity now automatically have a GetEntity Method
// This allows system.Add functions to recieve a single interface
//}
func (e *Entity) GetEntity() *Entity {
	return e
}

// AppendChild appends a child to the BasicEntity
func (e *Entity) AppendChild(child *Entity) {
	child.parent = e
	e.children = append(e.children, child)
}

func (e *Entity) RemoveChild(child *Entity) {
	delete := -1
	for i, v := range e.children {
		if v.ID() == child.ID() {
			delete = i
			break
		}
	}
	if delete >= 0 {
		e.children = append(e.children[:delete], e.children[delete+1:]...)
	}
}

// Children returns the children of the BasicEntity
func (e *Entity) Children() []Entity {
	ret := []Entity{}
	for _, child := range e.children {
		ret = append(ret, *child)
	}
	return ret
}

// Descendents returns the children and their children all the way down the tree.
func (e *Entity) Descendents() []Entity {
	return descendents([]Entity{}, e, e)
}

func (e *Entity) Remove() {
	e.manager.commandBuffer.removeEntity(e)
}

func (e *Entity) Register() *Entity {
	e.manager.commandBuffer.addEntity(e)
	return e
}

func descendents(in []Entity, this, top *Entity) []Entity {
	for _, child := range this.children {
		in = descendents(in, child, top)
	}
	if this.ID() == top.ID() {
		return in
	}
	return append(in, *this)
}

// Parent returns the parent of the BasicEntity
func (e *Entity) Parent() *Entity {
	return e.parent
}

// Len returns the length of the underlying slice
// part of the sort.Interface
func (is IdentifierSlice) Len() int {
	return len(is)
}

// Less will return true if the ID of element at i is less than j;
// part of the sort.Interface
func (is IdentifierSlice) Less(i, j int) bool {
	return is[i].ID() < is[j].ID()
}

// Swap the elements at positions i and j
// part of the sort.Interface
func (is IdentifierSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

// Face is an interface that Entity and entities containing
// a BasicEntity implement.
type Face interface {
	GetEntity() *Entity
}

type entites map[uint64]*Entity
