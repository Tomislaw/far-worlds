package utils

import "reflect"

type ReflectTypeGetter interface {
	ReflectType() reflect.Type
}
