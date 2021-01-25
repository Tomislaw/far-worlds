package component

import (
	"reflect"

	"github.com/Tomislaw/far-worlds/ecs"
)

type TypeDefinitions struct {
	GUID            reflect.Type
	MapItem         reflect.Type
	MapItemBlock    reflect.Type
	MapItemMovement reflect.Type
}

var Type = TypeDefinitions{
	GUID:            reflect.TypeOf((*GUID)(nil)).Elem(),
	MapItem:         reflect.TypeOf((*MapItem)(nil)).Elem(),
	MapItemBlock:    reflect.TypeOf((*MapItemBlock)(nil)).Elem(),
	MapItemMovement: reflect.TypeOf((*MapItemMovement)(nil)).Elem(),
}

func RegisterComponents(manager *ecs.Manager) {
	manager.RegisterComponents(
		Type.GUID,
		Type.MapItem,
		Type.MapItemBlock,
		Type.MapItemMovement,
	)
}
