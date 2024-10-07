package model_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gempages/go-shopify-graphql-model/graph/model"
)

var _ = Describe("Decode", func() {
	var (
		data   map[string]any
		result any
		err    error
	)

	When("data has an unknown `__typename`", func() {
		BeforeEach(func() {
			data = map[string]any{
				"__typename": "UnknownType",
			}
		})

		It("returns an error indicating unknown type", func() {
			result, err = model.Decode(data, nil)
			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
		})
	})

	When("data has no `__typename`", func() {
		BeforeEach(func() {
			data = map[string]any{
				"price": map[string]any{
					"amount":       "10.0",
					"currencyCode": "USD",
				},
			}
		})

		It("returns an error", func() {
			result, err = model.Decode(data, nil)
			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
		})
	})

	When("data has invalid decimal values", func() {
		BeforeEach(func() {
			data = map[string]any{
				"__typename": "AppRecurringPricing",
				"price": map[string]any{
					"amount":       "invalid_decimal",
					"currencyCode": "USD",
				},
			}
		})

		It("returns an error during decimal decoding", func() {
			result, err = model.Decode(data, &model.AppRecurringPricing{})
			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
		})
	})

	When("data has __typename AppRecurringPricing with AppSubscriptionDiscountAmount", func() {
		BeforeEach(func() {
			data = map[string]any{
				"__typename": "AppRecurringPricing",
				"price": map[string]any{
					"amount":       "10.0",
					"currencyCode": "USD",
				},
				"discount": map[string]any{
					"value": map[string]any{
						"__typename": "AppSubscriptionDiscountAmount",
						"amount": map[string]any{
							"amount":       "2.1",
							"currencyCode": "USD",
						},
					},
					"priceAfterDiscount": map[string]any{
						"amount":       "7.9",
						"currencyCode": "USD",
					},
				},
			}
		})

		It("decodes the data into an AppRecurringPricing object", func() {
			result, err = model.Decode(data, &model.AppRecurringPricing{})
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeAssignableToTypeOf(&model.AppRecurringPricing{}))

			arp, ok := result.(*model.AppRecurringPricing)
			Expect(ok).To(BeTrue())

			Expect(arp.Price).NotTo(BeNil())
			Expect(arp.Price.Amount.String()).To(Equal("10"))
			Expect(arp.Price.CurrencyCode.String()).To(Equal("USD"))

			discount := arp.Discount
			Expect(discount).To(BeAssignableToTypeOf(&model.AppSubscriptionDiscount{}))
			Expect(discount.PriceAfterDiscount).NotTo(BeNil())
			Expect(discount.PriceAfterDiscount.Amount.String()).To(Equal("7.9"))
			Expect(discount.PriceAfterDiscount.CurrencyCode.String()).To(Equal("USD"))

			discountValue, ok := discount.Value.(*model.AppSubscriptionDiscountAmount)
			Expect(ok).To(BeTrue())
			Expect(discountValue).NotTo(BeNil())
			Expect(discountValue.Amount.Amount.String()).To(Equal("2.1"))
			Expect(discountValue.Amount.CurrencyCode.String()).To(Equal("USD"))
		})
	})

	When("data has __typename AppRecurringPricing with AppSubscriptionDiscountPercentage", func() {
		BeforeEach(func() {
			data = map[string]any{
				"__typename": "AppRecurringPricing",
				"price": map[string]any{
					"amount":       "10.0",
					"currencyCode": "USD",
				},
				"discount": map[string]any{
					"value": map[string]any{
						"__typename": "AppSubscriptionDiscountPercentage",
						"percentage": 0.11,
					},
					"priceAfterDiscount": map[string]any{
						"amount":       "8.9",
						"currencyCode": "USD",
					},
				},
			}
		})

		It("decodes the data into an AppRecurringPricing object", func() {
			result, err = model.Decode(data, &model.AppRecurringPricing{})
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeAssignableToTypeOf(&model.AppRecurringPricing{}))

			arp, ok := result.(*model.AppRecurringPricing)
			Expect(ok).To(BeTrue())

			Expect(arp.Price).NotTo(BeNil())
			Expect(arp.Price.Amount.String()).To(Equal("10"))
			Expect(arp.Price.CurrencyCode.String()).To(Equal("USD"))

			discount := arp.Discount
			Expect(discount).To(BeAssignableToTypeOf(&model.AppSubscriptionDiscount{}))
			Expect(discount.PriceAfterDiscount).NotTo(BeNil())
			Expect(discount.PriceAfterDiscount.Amount.String()).To(Equal("8.9"))
			Expect(discount.PriceAfterDiscount.CurrencyCode.String()).To(Equal("USD"))

			discountValue, ok := discount.Value.(*model.AppSubscriptionDiscountPercentage)
			Expect(ok).To(BeTrue())
			Expect(discountValue).NotTo(BeNil())
			Expect(discountValue.Percentage).To(Equal(0.11))
		})
	})

	When("data has __typename AppUsagePricing", func() {
		BeforeEach(func() {
			data = map[string]any{
				"__typename": "AppUsagePricing",
				"cappedAmount": map[string]any{
					"amount":       "10.0",
					"currencyCode": "USD",
				},
				"balanceUsed": map[string]any{
					"amount":       "0",
					"currencyCode": "USD",
				},
				"terms":    "Some text here",
				"interval": "3",
			}
		})

		It("decodes the data into an AppUsagePricing object", func() {
			result, err = model.Decode(data, &model.AppUsagePricing{})
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeAssignableToTypeOf(&model.AppUsagePricing{}))

			aup, ok := result.(*model.AppUsagePricing)
			Expect(ok).To(BeTrue())

			Expect(aup.CappedAmount).NotTo(BeNil())
			Expect(aup.CappedAmount.Amount.String()).To(Equal("10"))
			Expect(aup.CappedAmount.CurrencyCode.String()).To(Equal("USD"))

			Expect(aup.BalanceUsed).NotTo(BeNil())
			Expect(aup.BalanceUsed.Amount.String()).To(Equal("0"))
			Expect(aup.BalanceUsed.CurrencyCode.String()).To(Equal("USD"))

			Expect(aup.Interval.String()).To(Equal("3"))
			Expect(aup.Terms).To(Equal("Some text here"))
		})
	})
})
