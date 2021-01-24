package component

import (
	"reflect"

	"github.com/Tomislaw/far-worlds/ecs"
)

func RegisterComponents(manager *ecs.Manager) {
	manager.RegisterComponents(
		reflect.TypeOf((*GUID)(nil)).Elem(),
		reflect.TypeOf((*MapItem)(nil)).Elem(),
		reflect.TypeOf((*MapItemBlock)(nil)).Elem(),
	)
}
