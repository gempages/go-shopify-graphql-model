package model_test

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"

	"github.com/gempages/go-shopify-graphql-model/graph/model"
)

var _ = Describe("UnmarshalJSON", func() {
	It("can unmarshal Media interface in MediaEdge", func() {
		id := "gid://shopify/MediaImage/123"
		alt := "example"
		result := new(model.MediaEdge)
		err := json.Unmarshal([]byte(mediaEdgeJSON), result)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).NotTo(BeNil())
		Expect(result.Cursor).To(Equal("cursor1"))
		Expect(result.Node).To(BeAssignableToTypeOf(&model.MediaImage{}))
		Expect(result.Node.GetID()).To(Equal(id))
		Expect(*(result.Node.GetAlt())).To(Equal(alt))
		img := result.Node.(*model.MediaImage)
		Expect(img.Image.ID).NotTo(BeNil())
		Expect(*img.Image.ID).To(Equal(id))
		Expect(img.UpdatedAt.Format(time.RFC3339)).To(Equal("2021-02-23T21:51:39Z"))
	})

	It("can unmarshal Media interface in MediaConnection", func() {
		result := new(model.MediaConnection)
		err := json.Unmarshal([]byte(mediaConnectionJSON), result)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).NotTo(BeNil())
		Expect(result.Nodes).NotTo(BeEmpty())
		Expect(result.Nodes[0]).To(BeAssignableToTypeOf(&model.MediaImage{}))
		Expect(result.Nodes[1]).To(BeAssignableToTypeOf(&model.Video{}))
		img := result.Nodes[0].(*model.MediaImage)
		Expect(img.ID).To(Equal("gid://shopify/MediaImage/123"))
		video := result.Nodes[1].(*model.Video)
		Expect(video.ID).To(Equal("gid://shopify/Video/123"))
	})

	It("can unmarshal WebhookSubscriptionEndpoint interface", func() {
		result := new(model.WebhookSubscription)
		err := json.Unmarshal([]byte(webhookSubscriptionJSON), result)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).NotTo(BeNil())
		Expect(result.Topic).To(Equal(model.WebhookSubscriptionTopicAppUninstalled))
		Expect(result.Endpoint).To(BeAssignableToTypeOf(&model.WebhookHTTPEndpoint{}))
		endpoint := result.Endpoint.(*model.WebhookHTTPEndpoint)
		Expect(endpoint.CallbackURL).To(Equal("https://example.com/webhooks/app_uninstalled"))
	})

	It("can unmarshal Discount interface", func() {
		result := new(model.DiscountNode)
		err := json.Unmarshal([]byte(discountNodeJSON), result)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).NotTo(BeNil())
		Expect(result.Discount).To(BeAssignableToTypeOf(&model.DiscountCodeBasic{}))
		discount := result.Discount.(*model.DiscountCodeBasic)
		Expect(discount.StartsAt.Format(time.RFC3339)).To(Equal("2024-09-01T08:30:00Z"))
		Expect(discount.TotalSales.Amount).To(Equal(decimal.New(1000, -1)))
		Expect(discount.TotalSales.CurrencyCode).To(Equal(model.CurrencyCodeEur))
	})

	It("can unmarshal File interface", func() {
		result := new(model.FileEdge)
		err := json.Unmarshal([]byte(fileEdgeJSON), result)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).NotTo(BeNil())
		Expect(result.Cursor).To(Equal("cursor1"))
		Expect(result.Node).To(BeAssignableToTypeOf(&model.GenericFile{}))
		Expect(result.Node.GetID()).To(Equal("gid://shopify/File/123"))
		file := result.Node.(*model.GenericFile)
		Expect(file.FileStatus).To(Equal(model.FileStatusFailed))
		Expect(file.UpdatedAt.Format(time.RFC3339)).To(Equal("2021-02-23T21:51:39Z"))
		Expect(file.FileErrors).NotTo(BeEmpty())
		Expect(file.FileErrors[0].Code).To(Equal(model.FileErrorCodeGenericFileInvalidSize))
	})

	It("can unmarshal File interface in FileCreatePayload", func() {
		result := new(model.FileCreatePayload)
		err := json.Unmarshal([]byte(fileCreatePayloadJSON), result)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).NotTo(BeNil())
		Expect(result.Files).NotTo(BeEmpty())
		Expect(result.UserErrors).NotTo(BeEmpty())
		Expect(result.Files[0]).To(BeAssignableToTypeOf(&model.GenericFile{}))
		file := result.Files[0].(*model.GenericFile)
		Expect(file.ID).To(Equal("gid://shopify/File/123"))
		Expect(file.FileStatus).To(Equal(model.FileStatusFailed))
		Expect(result.UserErrors[0].Code).NotTo(BeNil())
		Expect(*result.UserErrors[0].Code).To(Equal(model.FilesErrorCodeInvalidFilename))
	})

	It("can unmarshal PricingDetails interface", func() {
		result := new(model.AppPlanV2)
		err := json.Unmarshal([]byte(appPlanV2JSON), result)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).NotTo(BeNil())
		Expect(result.PricingDetails).NotTo(BeNil())
		Expect(result.PricingDetails).To(BeAssignableToTypeOf(&model.AppRecurringPricing{}))
		pd := result.PricingDetails.(*model.AppRecurringPricing)
		Expect(pd.Price).NotTo(BeNil())
		Expect(pd.Price.Amount).To(Equal(decimal.New(225, -1)))
		Expect(pd.Price.CurrencyCode).To(Equal(model.CurrencyCodeUsd))
	})
})

var (
	mediaEdgeJSON = `{
	"cursor": "cursor1",
	"node": {
		"__typename": "MediaImage",
		"id": "gid://shopify/MediaImage/123",
		"alt": "example",
		"image": {
			"id": "gid://shopify/MediaImage/123"
		},
		"updatedAt": "2021-02-23T21:51:39Z"
	}
}`

	mediaConnectionJSON = `{
	"nodes": [{
		"__typename": "MediaImage",
		"id": "gid://shopify/MediaImage/123",
		"alt": "example"
	}, {
		"__typename": "Video",
		"id": "gid://shopify/Video/123",
		"alt": "example-video"
	}],
	"pageInfo": {
		"hasNextPage": true,
		"hasPreviousPage": false
	}
}`

	webhookSubscriptionJSON = `{
	"topic": "APP_UNINSTALLED",
	"endpoint": {
		"__typename": "WebhookHttpEndpoint",
		"callbackUrl": "https://example.com/webhooks/app_uninstalled"
	}
}`

	discountNodeJSON = `{
	"discount": {
		"__typename": "DiscountCodeBasic",
		"startsAt": "2024-09-01T08:30:00Z",
		"totalSales": {
			"amount": "100.0",
			"currencyCode": "EUR"
		}
	}
}`

	fileEdgeJSON = `{
	"cursor": "cursor1",
	"node": {
		"__typename": "GenericFile",
		"id": "gid://shopify/File/123",
		"fileStatus": "FAILED",
		"mimeType": "text/plain",
		"updatedAt": "2021-02-23T21:51:39Z",
		"fileErrors": [{
			"code": "GENERIC_FILE_INVALID_SIZE",
			"details": "File is too large",
			"message": "The file is too large"
		}]
	}
}`

	fileCreatePayloadJSON = `{
	"files": [{
		"__typename": "GenericFile",
		"id": "gid://shopify/File/123",
		"fileStatus": "FAILED",
		"mimeType": "text/plain",
		"updatedAt": "2021-02-23T21:51:39Z"
	}],
	"userErrors": [{
		"code": "INVALID_FILENAME"
	}]
}`

	appPlanV2JSON = `{
	"pricingDetails": {
		"__typename": "AppRecurringPricing",
		"price": {
			"amount": "22.5",
			"currencyCode": "USD"
		}
	}
}`
)
