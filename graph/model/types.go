package model

import (
	"fmt"
	"reflect"
)

type GqlTypeName string

const (
	MediaImageTypeName                    GqlTypeName = "MediaImage"
	VideoTypeName                         GqlTypeName = "Video"
	Model3dTypeName                       GqlTypeName = "Model3d"
	ExternalVideoTypeName                 GqlTypeName = "ExternalVideo"
	GenericFileTypeName                   GqlTypeName = "GenericFile"
	DiscountAutomaticAppTypeName          GqlTypeName = "DiscountAutomaticApp"
	DiscountCodeBasicTypeName             GqlTypeName = "DiscountCodeBasic"
	DiscountCodeBuyXGetYTypeName          GqlTypeName = "DiscountCodeBuyXGetY"
	DiscountCodeFreeShippingTypeName      GqlTypeName = "DiscountCodeFreeShipping"
	DiscountAutomaticBasicTypeName        GqlTypeName = "DiscountAutomaticBasic"
	DiscountAutomaticBxgyTypeName         GqlTypeName = "DiscountAutomaticBxgy"
	DiscountAutomaticFreeShippingTypeName GqlTypeName = "DiscountAutomaticFreeShipping"
	AppRecurringPricingTypeName           GqlTypeName = "AppRecurringPricing"
	AppUsagePricingTypeName               GqlTypeName = "AppUsagePricing"
)

// concludeObjectType returns the `reflect.Type` corresponding to a given `typeName`, returns error if `typeName` is not supported.
func concludeObjectType(typeName GqlTypeName) (reflect.Type, error) {
	switch typeName {
	case MediaImageTypeName:
		return reflect.TypeOf(MediaImage{}), nil
	case VideoTypeName:
		return reflect.TypeOf(Video{}), nil
	case Model3dTypeName:
		return reflect.TypeOf(Model3d{}), nil
	case ExternalVideoTypeName:
		return reflect.TypeOf(ExternalVideo{}), nil
	case GenericFileTypeName:
		return reflect.TypeOf(GenericFile{}), nil
	case DiscountAutomaticAppTypeName:
		return reflect.TypeOf(DiscountAutomaticApp{}), nil
	case DiscountCodeBasicTypeName:
		return reflect.TypeOf(DiscountCodeBasic{}), nil
	case DiscountCodeBuyXGetYTypeName:
		return reflect.TypeOf(DiscountCodeBxgy{}), nil
	case DiscountCodeFreeShippingTypeName:
		return reflect.TypeOf(DiscountCodeFreeShipping{}), nil
	case DiscountAutomaticBasicTypeName:
		return reflect.TypeOf(DiscountAutomaticBasic{}), nil
	case DiscountAutomaticBxgyTypeName:
		return reflect.TypeOf(DiscountAutomaticBxgy{}), nil
	case DiscountAutomaticFreeShippingTypeName:
		return reflect.TypeOf(DiscountAutomaticFreeShipping{}), nil
	case AppRecurringPricingTypeName:
		return reflect.TypeOf(AppRecurringPricing{}), nil
	case AppUsagePricingTypeName:
		return reflect.TypeOf(AppUsagePricing{}), nil
	default:
		return reflect.TypeOf(nil), fmt.Errorf("`%s` not implemented type", typeName)
	}
}
