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
	case "GenericFile":
		return reflect.TypeOf(GenericFile{}), nil
	case "DiscountAutomaticApp":
		return reflect.TypeOf(DiscountAutomaticApp{}), nil
	case "DiscountCodeBasic":
		return reflect.TypeOf(DiscountCodeBasic{}), nil
	case "DiscountCodeBuyXGetY":
		return reflect.TypeOf(DiscountCodeBxgy{}), nil
	case "DiscountCodeFreeShipping":
		return reflect.TypeOf(DiscountCodeFreeShipping{}), nil
	case "DiscountAutomaticBasic":
		return reflect.TypeOf(DiscountAutomaticBasic{}), nil
	case "DiscountAutomaticBxgy":
		return reflect.TypeOf(DiscountAutomaticBxgy{}), nil
	case "DiscountAutomaticFreeShipping":
		return reflect.TypeOf(DiscountAutomaticFreeShipping{}), nil
	default:
		return reflect.TypeOf(nil), fmt.Errorf("`%s` not implemented type", typeName)
	}
}
