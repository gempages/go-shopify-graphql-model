package model

import (
	"fmt"
	"reflect"

	"github.com/gempages/go-shopify-graphql-model/graph/types"
)

// ConcludeObjectType returns the `reflect.Type` corresponding to a given `typeName`, returns error if `typeName` is not supported.
func ConcludeObjectType(typeName types.GqlTypeName) (reflect.Type, error) {
	switch typeName {
	case types.MediaImage:
		return reflect.TypeOf(MediaImage{}), nil
	case types.Video:
		return reflect.TypeOf(Video{}), nil
	case types.Model3d:
		return reflect.TypeOf(Model3d{}), nil
	case types.ExternalVideo:
		return reflect.TypeOf(ExternalVideo{}), nil
	case types.GenericFile:
		return reflect.TypeOf(GenericFile{}), nil
	case types.DiscountAutomaticApp:
		return reflect.TypeOf(DiscountAutomaticApp{}), nil
	case types.DiscountCodeBasic:
		return reflect.TypeOf(DiscountCodeBasic{}), nil
	case types.DiscountCodeBuyXGetY:
		return reflect.TypeOf(DiscountCodeBxgy{}), nil
	case types.DiscountCodeFreeShipping:
		return reflect.TypeOf(DiscountCodeFreeShipping{}), nil
	case types.DiscountAutomaticBasic:
		return reflect.TypeOf(DiscountAutomaticBasic{}), nil
	case types.DiscountAutomaticBxgy:
		return reflect.TypeOf(DiscountAutomaticBxgy{}), nil
	case types.DiscountAutomaticFreeShipping:
		return reflect.TypeOf(DiscountAutomaticFreeShipping{}), nil
	case types.AppRecurringPricing:
		return reflect.TypeOf(AppRecurringPricing{}), nil
	case types.AppUsagePricing:
		return reflect.TypeOf(AppUsagePricing{}), nil
	case types.AppSubscriptionDiscount:
		return reflect.TypeOf(AppSubscriptionDiscount{}), nil
	case types.AppSubscriptionDiscountAmount:
		return reflect.TypeOf(AppSubscriptionDiscountAmount{}), nil
	case types.AppSubscriptionDiscountPercentage:
		return reflect.TypeOf(AppSubscriptionDiscountPercentage{}), nil
	default:
		return nil, fmt.Errorf("`%s` is not implemented", typeName)
	}
}
