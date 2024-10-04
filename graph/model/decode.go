package model

import (
	"fmt"
	"reflect"

	"github.com/gempages/go-shopify-graphql-model/graph/types"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
)

// Decode decodes the input data into an object of the type specified by GqlTypeNameKey.
func Decode(data map[string]any) (any, error) {
	typeName, ok := data[types.GqlTypeNameKey].(string)
	if !ok {
		return nil, fmt.Errorf("%s field is not supported or not a string in `%s`", types.GqlTypeNameKey, data)
	}

	objType, err := ConcludeObjectType(types.GqlTypeName(typeName))
	if err != nil {
		return nil, fmt.Errorf("concludeObjectType: %w", err)
	}

	objStruct := reflect.New(objType).Interface()
	decoder, err := newDecoder(decodeHook, &objStruct)
	if err != nil {
		return nil, fmt.Errorf("newDecoder: %w", err)
	}

	if err = decoder.Decode(data); err != nil {
		return nil, fmt.Errorf("decoder.Decode: %w", err)
	}

	return objStruct, nil
}

// newDecoder creates a new `mapstructure.Decoder` configured with a custom decode hook and output target.
func newDecoder(decodeHook any, output any) (*mapstructure.Decoder, error) {
	decoderConfig := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			decodeHook,
		),
		Result:  output,
		TagName: "json",
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return nil, fmt.Errorf("create decoder: %w", err)
	}

	return decoder, nil
}

func decodeHook(_ reflect.Type, to reflect.Type, data any) (any, error) {
	switch to {
	case reflect.TypeOf((*AppSubscriptionDiscountValue)(nil)).Elem():
		m, ok := data.(map[string]any)
		if !ok {
			return data, nil
		}

		// Calling Decode recursive to decode children interfaces
		result, err := Decode(m)
		if err != nil {
			return nil, fmt.Errorf("decode Interface %s: %w", to, err)
		}

		return result, nil
	case reflect.TypeOf(decimal.Decimal{}):
		// Handle converts types to decimal.Decimal
		num, err := decodeDecimal(data)
		if err != nil {
			return nil, fmt.Errorf("decode decimal: %w", err)
		}

		return num, nil
	}

	return data, nil
}

// decodeDecimal converts various types to decimal.Decimal.
func decodeDecimal(data any) (*decimal.Decimal, error) {
	switch v := data.(type) {
	case string:
		d, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		return &d, nil
	case float64:
		d := decimal.NewFromFloat(v)
		return &d, nil
	case int:
		d := decimal.NewFromInt(int64(v))
		return &d, nil
	case int64:
		d := decimal.NewFromInt(v)
		return &d, nil
	default:
		return nil, fmt.Errorf("cannot convert %v to decimal.Decimal", data)
	}
}
