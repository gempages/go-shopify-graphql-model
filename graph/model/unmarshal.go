package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

func (e *MediaEdge) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	if cursor, ok := m["cursor"].(string); ok {
		e.Cursor = cursor
	}
	if node, ok := m["node"].(map[string]interface{}); ok {
		e.Node, err = decodeMedia(node)
		if err != nil {
			return fmt.Errorf("decode media node: %w", err)
		}
	}
	return nil
}

func (c *MediaConnection) UnmarshalJSON(b []byte) error {
	var (
		m     map[string]interface{}
		mConn struct {
			Edges    []MediaEdge `json:"edges,omitempty"`
			PageInfo *PageInfo   `json:"pageInfo,omitempty"`
		}
	)
	err := json.Unmarshal(b, &mConn)
	if err != nil {
		return err
	}
	c.Edges = mConn.Edges
	c.PageInfo = mConn.PageInfo

	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	if nodes, ok := m["nodes"].([]interface{}); ok {
		c.Nodes = make([]Media, len(nodes))
		for i, n := range nodes {
			if node, ok := n.(map[string]interface{}); ok {
				c.Nodes[i], err = decodeMedia(node)
				if err != nil {
					return fmt.Errorf("decode media node: %w", err)
				}
			} else {
				return fmt.Errorf("expected type map[string]interface{} for Media node, got %T", n)
			}
		}
	}
	return nil
}

func decodeMedia(node map[string]interface{}) (Media, error) {
	if typeName, ok := node["__typename"].(string); ok {
		mediaType, err := concludeObjectType(GqlTypeName(typeName))
		if err != nil {
			return nil, fmt.Errorf("conclude object type: %w", err)
		}
		media := reflect.New(mediaType).Interface()
		err = mapstructure.Decode(node, media)
		if err != nil {
			return nil, fmt.Errorf("decode media node: %w", err)
		}
		return media.(Media), nil
	}
	return nil, fmt.Errorf("must query __typename to decode Media")
}

func (s *WebhookSubscription) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	// mapstructure.Decode can't decode these fields
	createdAt := m["createdAt"]
	updatedAt := m["updatedAt"]
	endpoint := m["endpoint"]
	delete(m, "createdAt")
	delete(m, "updatedAt")
	delete(m, "endpoint")
	err = mapstructure.Decode(m, s)
	if err != nil {
		return fmt.Errorf("decode map: %w", err)
	}
	if c, ok := createdAt.(string); ok {
		s.CreatedAt, _ = time.Parse(time.RFC3339, c)
	}
	if u, ok := updatedAt.(string); ok {
		s.UpdatedAt, _ = time.Parse(time.RFC3339, u)
	}
	if node, ok := endpoint.(map[string]interface{}); ok {
		s.Endpoint, err = decodeWebhookSubscriptionEndpoint(node)
		if err != nil {
			return fmt.Errorf("decode endpoint: %w", err)
		}
	}
	return nil
}

func decodeWebhookSubscriptionEndpoint(node map[string]interface{}) (WebhookSubscriptionEndpoint, error) {
	var endpoint WebhookSubscriptionEndpoint
	if typeName, ok := node["__typename"].(string); ok {
		if typeName == "WebhookHttpEndpoint" {
			endpoint = &WebhookHTTPEndpoint{}
		}
		if typeName == "WebhookEventBridgeEndpoint" {
			endpoint = &WebhookEventBridgeEndpoint{}
		}
		if typeName == "WebhookPubSubEndpoint" {
			endpoint = &WebhookPubSubEndpoint{}
		}
	} else {
		return nil, fmt.Errorf("must query __typename to decode WebhookSubscriptionEndpoint")
	}
	err := mapstructure.Decode(node, endpoint)
	if err != nil {
		return nil, fmt.Errorf("decode webhook subscription endpoint node: %w", err)
	}
	return endpoint, nil
}

func (dc *DiscountNode) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	if discount, ok := m["discount"].(map[string]interface{}); ok {
		disc, err := decodeDiscount(discount)
		if err != nil {
			return fmt.Errorf("decodeDiscount: %w", err)
		}
		dc.Discount = disc.(Discount)
	}
	return nil
}

func (fe *FileEdge) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if cursor, ok := m["cursor"].(string); ok {
		fe.Cursor = cursor
	}
	if node, ok := m["node"].(map[string]interface{}); ok {
		fe.Node, err = decodeFile(node)
		if err != nil {
			return fmt.Errorf("decode file node: %w", err)
		}
	}
	return nil
}

func (fp *FileCreatePayload) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if files, ok := m["files"].([]interface{}); ok {
		fp.Files = make([]File, len(files))
		for i, n := range files {
			if file, ok := n.(map[string]interface{}); ok {
				fp.Files[i], err = decodeFile(file)
				if err != nil {
					return fmt.Errorf("decode file node: %w", err)
				}
			} else {
				return fmt.Errorf("expected type map[string]interface{} for File, got %T", n)
			}
		}
	}

	if userErrors, ok := m["userErrors"].([]interface{}); ok {
		var uErrors []FilesUserError
		bytes, err := json.Marshal(userErrors)
		if err != nil {
			return fmt.Errorf("marshal userErros got %w", err)
		}

		err = json.Unmarshal(bytes, &uErrors)
		if err != nil {
			return fmt.Errorf("unmarshal userErros got %w", err)
		}
		fp.UserErrors = uErrors
	}

	return nil
}

func (fc *FileConnection) UnmarshalJSON(b []byte) error {
	var (
		m     map[string]interface{}
		mConn struct {
			Edges    []FileEdge `json:"edges,omitempty"`
			PageInfo *PageInfo  `json:"pageInfo,omitempty"`
		}
	)

	err := json.Unmarshal(b, &mConn)
	if err != nil {
		return err
	}
	fc.Edges = mConn.Edges
	fc.PageInfo = mConn.PageInfo

	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	if nodes, ok := m["nodes"].([]interface{}); ok {
		fc.Nodes = make([]File, len(nodes))
		for i, n := range nodes {
			if node, ok := n.(map[string]interface{}); ok {
				fc.Nodes[i], err = decodeFile(node)
				if err != nil {
					return fmt.Errorf("decode file node: %w", err)
				}
			} else {
				return fmt.Errorf("expected type map[string]interface{} for File node, got %T", n)
			}
		}
	}
	return nil
}

func (d *DiscountAutomaticNode) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return fmt.Errorf("unmarshal DiscountAutomaticNode: %w", err)
	}

	if id, ok := m["id"].(string); ok {
		d.ID = id
	}
	if metafield, ok := m["metafield"].(map[string]any); ok {
		if d.Metafield == nil {
			d.Metafield = &Metafield{}
		}
		if err := mapstructure.Decode(metafield, d.Metafield); err != nil {
			return fmt.Errorf("decode metafield: %w", err)
		}
	}

	// Unmarshal AutomaticDiscount into a map to access __typename
	if automaticDiscount, ok := m["automaticDiscount"].(map[string]any); ok {
		discount, err := decodeDiscount(automaticDiscount)
		if err != nil {
			return fmt.Errorf("decodeDiscount: %w", err)
		}
		d.AutomaticDiscount = discount.(DiscountAutomatic)
	}

	return nil
}

func decodeDiscount(node map[string]any) (any, error) {
	typeName, ok := node["__typename"].(string)
	if !ok {
		return nil, fmt.Errorf("`__typename` field not found or not a string in `%s`", node)
	}
	discountType, err := concludeObjectType(GqlTypeName(typeName))
	if err != nil {
		return nil, fmt.Errorf("concludeObjectType: %w", err)
	}
	if startsAt, ok := node["startsAt"].(string); ok {
		node["startsAt"] = cast.ToTime(startsAt)
	}
	if endsAt, ok := node["endsAt"]; ok && endsAt != nil {
		node["endsAt"] = cast.ToTime(endsAt)
	}
	discount := reflect.New(discountType).Interface()
	if err := mapstructure.Decode(node, discount); err != nil {
		return nil, fmt.Errorf("mapstructure.Decode AutomaticDiscount: %w", err)
	}
	return discount, nil
}

func decodeFile(node map[string]interface{}) (File, error) {
	if typeName, ok := node["__typename"].(string); ok {
		fileType, err := concludeObjectType(GqlTypeName(typeName))
		if err != nil {
			return nil, fmt.Errorf("conclude object type: %w", err)
		}
		file := reflect.New(fileType).Interface()
		err = mapstructure.Decode(node, file)
		if err != nil {
			return nil, fmt.Errorf("decode file node: %w", err)
		}

		return file.(File), nil
	}
	return nil, fmt.Errorf("must query __typename to decode File")
}

func (p *AppPlanV2) UnmarshalJSON(data []byte) error {
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return fmt.Errorf("unmarshal AppPlanV2: %w", err)
	}

	if pd, ok := m["pricingDetails"].(map[string]any); ok {
		pricingDetails, err := decodePricingDetails(pd)
		if err != nil {
			return fmt.Errorf("decodePricingDetails: %w", err)
		}
		p.PricingDetails = pricingDetails.(AppPricingDetails)
	}
	return nil
}

func decodePricingDetails(data map[string]any) (any, error) {
	typeName, ok := data["__typename"].(string)
	if !ok {
		return nil, fmt.Errorf("`__typename` field is not supported or not a string in `%s`", data)
	}

	// pricingDetails has a nested structure with struct pointers and interfaces, preventing direct use of `mapstructure.Decode`.
	// So its fields must be formated to correct data types before decoding.
	newData := make(map[string]any)
	switch GqlTypeName(typeName) {
	case AppRecurringPricingTypeName:
		newData = formatAppRecurringPricingDetails(data)
	case AppUsagePricingTypeName:
		newData = formatAppUsagePricingDetails(data)
	default:
		return nil, fmt.Errorf("pricingDetails `__typeName` is not supported")
	}

	pricingType, err := concludeObjectType(GqlTypeName(typeName))
	if err != nil {
		return nil, fmt.Errorf("concludeObjectType: %w", err)
	}

	pricingDetails := reflect.New(pricingType).Interface()
	if err = mapstructure.Decode(newData, pricingDetails); err != nil {
		return nil, fmt.Errorf("mapstructure.Decode PricingDetails: %w", err)
	}

	return pricingDetails, nil
}

func formatAppUsagePricingDetails(data map[string]any) map[string]any {
	if balanceUsed, ok := data["balanceUsed"].(map[string]any); ok {
		if amount, ok := balanceUsed["amount"]; ok {
			balanceUsed["amount"] = toDecimal(amount)
		}
	}
	if cappedAmount, ok := data["cappedAmount"].(map[string]any); ok {
		if amount, ok := cappedAmount["amount"]; ok {
			cappedAmount["amount"] = toDecimal(amount)
		}
	}
	if price, ok := data["price"].(map[string]any); ok {
		if amount, ok := price["amount"]; ok {
			price["amount"] = toDecimal(amount)
		}
	}

	return data
}

func formatAppRecurringPricingDetails(data map[string]any) map[string]any {
	if price, ok := data["price"].(map[string]any); ok {
		if amount, ok := price["amount"]; ok {
			price["amount"] = toDecimal(amount)
		}
	}

	if discount, ok := data["discount"].(map[string]any); ok {
		if priceAfterDiscount, ok := discount["priceAfterDiscount"].(map[string]any); ok {
			if amount, ok := priceAfterDiscount["amount"]; ok {
				priceAfterDiscount["amount"] = toDecimal(amount)
			}
		}

		if value, ok := discount["value"].(map[string]any); ok {
			if a, ok := value["amount"].(map[string]any); ok {
				if amount, ok := a["amount"]; ok {
					if currencyCode, ok := a["currencyCode"]; ok {
						discount["value"] = &AppSubscriptionDiscountAmount{
							Amount: &MoneyV2{
								Amount:       toDecimal(amount),
								CurrencyCode: CurrencyCode(cast.ToString(currencyCode)),
							},
						}
					}
				}
			} else if percentage, ok := value["percentage"].(float64); ok {
				discount["value"] = &AppSubscriptionDiscountPercentage{
					Percentage: percentage,
				}
			}
		}
	}

	return data
}

// toDecimal converts a value to float64, then from float64 to Decimal.
func toDecimal(val any) decimal.Decimal {
	return decimal.NewFromFloat(cast.ToFloat64(val))
}
