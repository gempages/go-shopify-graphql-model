package model

import (
	"fmt"
	"reflect"
)

func concludeObjectType(typeName string) (reflect.Type, error) {
	switch typeName {
	case "MediaImage":
		return reflect.TypeOf(MediaImage{}), nil
	case "Video":
		return reflect.TypeOf(Video{}), nil
	case "Model3d":
		return reflect.TypeOf(Model3d{}), nil
	case "ExternalVideo":
		return reflect.TypeOf(ExternalVideo{}), nil
	default:
		return reflect.TypeOf(nil), fmt.Errorf("`%s` not implemented type", typeName)
	}
}
