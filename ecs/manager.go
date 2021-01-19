package ecs

import (
	"reflect"
	"sort"
	"sync"
)

// Manager contains a bunch of Entities, and a bunch of Systems. It is the
// recommended way to run ecs.

const maxComponents = 64

type Manager struct {
	systems systems

	components     components
	componentslock sync.RWMutex

	sysIn, sysEx map[reflect.Type][]reflect.Type
}

func (w *Manager) AddComponent(entity *Entity, component interface{}) *Manager {
	ctype := w.addComponentType(reflect.TypeOf(component))
	ctype.datalock.Lock()
	ctype.data[entity.id] = component
	ctype.datalock.Unlock()
	entity.componentFlags |= 1 << ctype.id
	return w
}

func (w *Manager) GetComponent(entity *Entity, componentType reflect.Type) interface{} {

	w.componentslock.RLock()

	ok := w.HasComponent(entity, componentType)
	if !ok {
		w.componentslock.RUnlock()
		return nil
	}

	component, ok := w.components[componentType]
	if !ok {
		w.componentslock.RUnlock()
		return nil
	}

	data, ok := component.data[entity.id]
	if !ok {
		w.componentslock.RUnlock()
		return nil
	}
	return data
}

func (w *Manager) HasComponent(entity *Entity, componentType reflect.Type) bool {
	ctype := w.getComponentType(componentType)
	if ctype == nil {
		return false
	}
	id := ctype.id
	var flag uint64 = 1 << id
	return entity.componentFlags&flag == flag
}

func (w *Manager) getComponentType(componentType reflect.Type) *Component {

	w.componentslock.RLock()
	component, ok := w.components[componentType]
	w.componentslock.RUnlock()

	if ok {
		return component
	}
	return nil
}

func (w *Manager) addComponentType(componentType reflect.Type) *Component {

	w.componentslock.Lock()

	component, ok := w.components[componentType]

	if ok {
		w.componentslock.Unlock()
		return component
	}

	if w.components.Len() > maxComponents {
		w.componentslock.Unlock()
		panic("Too many component types!")
	}

	newType := &Component{
		id:       uint8(w.components.Len()),
		datalock: &sync.RWMutex{},
	}

	w.components[componentType] = newType
	w.componentslock.Unlock()
	return component
}

// AddSystem adds the given System to the World, sorted by priority.
func (w *Manager) AddSystem(system System) *Manager {
	if initializer, ok := system.(Initializer); ok {
		initializer.New(w)
	}

	w.systems = append(w.systems, system)
	sort.Sort(w.systems)
	return w
}

// AddSystemInterface adds a system to the world, but also adds a filter that allows
// automatic adding of entities that match the provided in interface, and excludes any
// that match the provided ex interface, even if they also match in. in and ex must be
// pointers to the interface or else this panics.
func (w *Manager) AddSystemInterface(sys SystemAddByInterfacer, in interface{}, ex interface{}) {
	w.AddSystem(sys)

	if w.sysIn == nil {
		w.sysIn = make(map[reflect.Type][]reflect.Type)
	}

	if !reflect.TypeOf(in).AssignableTo(reflect.TypeOf([]interface{}{})) {
		in = []interface{}{in}
	}
	for _, v := range in.([]interface{}) {
		w.sysIn[reflect.TypeOf(sys)] = append(w.sysIn[reflect.TypeOf(sys)], reflect.TypeOf(v).Elem())
	}

	if ex == nil {
		return
	}

	if w.sysEx == nil {
		w.sysEx = make(map[reflect.Type][]reflect.Type)
	}

	if !reflect.TypeOf(ex).AssignableTo(reflect.TypeOf([]interface{}{})) {
		ex = []interface{}{ex}
	}
	for _, v := range ex.([]interface{}) {
		w.sysEx[reflect.TypeOf(sys)] = append(w.sysEx[reflect.TypeOf(sys)], reflect.TypeOf(v).Elem())
	}
}

// AddEntity adds the entity to all systems that have been added via
// AddSystemInterface. If the system was added via AddSystem the entity will not be
// added to it.
func (w *Manager) AddEntity(e Identifier) {
	if w.sysIn == nil {
		w.sysIn = make(map[reflect.Type][]reflect.Type)
	}
	if w.sysEx == nil {
		w.sysEx = make(map[reflect.Type][]reflect.Type)
	}

	search := func(i Identifier, types []reflect.Type) bool {
		for _, t := range types {
			if reflect.TypeOf(i).Implements(t) {
				return true
			}
		}
		return false
	}
	for _, system := range w.systems {
		sys, ok := system.(SystemAddByInterfacer)
		if !ok {
			continue
		}

		if ex, not := w.sysEx[reflect.TypeOf(sys)]; not {
			if search(e, ex) {
				continue
			}
		}
		if in, ok := w.sysIn[reflect.TypeOf(sys)]; ok {
			if search(e, in) {
				sys.AddByInterface(e)
				continue
			}
		}
	}
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

// SortSystems sorts the systems in the world.
func (w *Manager) SortSystems() {
	sort.Sort(w.systems)
}
