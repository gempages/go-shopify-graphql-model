package types

type GqlTypeName string

const GqlTypeNameKey = "__typename"

const (
	MediaImage                        GqlTypeName = "MediaImage"
	Video                             GqlTypeName = "Video"
	Model3d                           GqlTypeName = "Model3d"
	ExternalVideo                     GqlTypeName = "ExternalVideo"
	GenericFile                       GqlTypeName = "GenericFile"
	DiscountAutomaticApp              GqlTypeName = "DiscountAutomaticApp"
	DiscountCodeBasic                 GqlTypeName = "DiscountCodeBasic"
	DiscountCodeBuyXGetY              GqlTypeName = "DiscountCodeBuyXGetY"
	DiscountCodeFreeShipping          GqlTypeName = "DiscountCodeFreeShipping"
	DiscountAutomaticBasic            GqlTypeName = "DiscountAutomaticBasic"
	DiscountAutomaticBxgy             GqlTypeName = "DiscountAutomaticBxgy"
	DiscountAutomaticFreeShipping     GqlTypeName = "DiscountAutomaticFreeShipping"
	AppRecurringPricing               GqlTypeName = "AppRecurringPricing"
	AppUsagePricing                   GqlTypeName = "AppUsagePricing"
	AppSubscriptionDiscount           GqlTypeName = "AppSubscriptionDiscount"
	AppSubscriptionDiscountAmount     GqlTypeName = "AppSubscriptionDiscountAmount"
	AppSubscriptionDiscountPercentage GqlTypeName = "AppSubscriptionDiscountPercentage"
	WebhookHttpEndpoint               GqlTypeName = "WebhookHttpEndpoint"
	WebhookEventBridgeEndpoint        GqlTypeName = "WebhookEventBridgeEndpoint"
	WebhookPubSubEndpoint             GqlTypeName = "WebhookPubSubEndpoint"
)
