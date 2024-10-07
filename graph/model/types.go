package model

import (
	"fmt"

	"github.com/gempages/go-shopify-graphql-model/graph/types"
)

// concludeObjectType returns a pointer to the corresponding struct of given `typeName`. It returns error if `typeName` is not implemented.
func concludeObjectType(typeName types.GqlTypeName) (any, error) {
	switch typeName {
	case types.MediaImage:
		return &MediaImage{}, nil
	case types.Video:
		return &Video{}, nil
	case types.Model3d:
		return &Model3d{}, nil
	case types.ExternalVideo:
		return &ExternalVideo{}, nil
	case types.GenericFile:
		return &GenericFile{}, nil
	case types.DiscountAutomaticApp:
		return &DiscountAutomaticApp{}, nil
	case types.DiscountCodeBasic:
		return &DiscountCodeBasic{}, nil
	case types.DiscountCodeBuyXGetY:
		return &DiscountCodeBxgy{}, nil
	case types.DiscountCodeFreeShipping:
		return &DiscountCodeFreeShipping{}, nil
	case types.DiscountAutomaticBasic:
		return &DiscountAutomaticBasic{}, nil
	case types.DiscountAutomaticBxgy:
		return &DiscountAutomaticBxgy{}, nil
	case types.DiscountAutomaticFreeShipping:
		return &DiscountAutomaticFreeShipping{}, nil
	case types.AppRecurringPricing:
		return &AppRecurringPricing{}, nil
	case types.AppUsagePricing:
		return &AppUsagePricing{}, nil
	case types.AppSubscriptionDiscount:
		return &AppSubscriptionDiscount{}, nil
	case types.AppSubscriptionDiscountAmount:
		return &AppSubscriptionDiscountAmount{}, nil
	case types.AppSubscriptionDiscountPercentage:
		return &AppSubscriptionDiscountPercentage{}, nil
	case types.WebhookHttpEndpoint:
		return &WebhookHTTPEndpoint{}, nil
	case types.WebhookEventBridgeEndpoint:
		return &WebhookEventBridgeEndpoint{}, nil
	case types.WebhookPubSubEndpoint:
		return &WebhookPubSubEndpoint{}, nil
	default:
		return nil, fmt.Errorf("`%s` is not implemented", typeName)
	}
}
