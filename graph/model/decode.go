package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"

	"github.com/gempages/go-shopify-graphql-model/graph/types"
)

// Decode decodes the input map into an object of the type specified by GqlTypeNameKey.
// If the target is a non-empty interface, it will decode the data into the concrete type.
func Decode(data map[string]any, target any) (any, error) {
	decoder, err := newDecoder(decodeHook, target)
	if err != nil {
		return nil, fmt.Errorf("newDecoder: %w", err)
	}

	if err = decoder.Decode(data); err != nil {
		return nil, fmt.Errorf("decoder.Decode: %w", err)
	}

	return target, nil
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
	isInterface := to.Name() != "" && to.Kind() == reflect.Interface
	// If the target is a non-empty interface, we need to decode the data into the concrete type
	if m, ok := data.(map[string]any); ok && isInterface {
		typeName, ok := m[types.GqlTypeNameKey].(string)
		if !ok || typeName == "" {
			return nil, fmt.Errorf("must query %s to decode %s", types.GqlTypeNameKey, to.Name())
		}
		objType, err := concludeObjectType(types.GqlTypeName(typeName))
		if err != nil {
			return nil, fmt.Errorf("concludeObjectType: %w", err)
		}
		// Calling Decode recursively to decode children interfaces
		result, err := Decode(m, objType)
		if err != nil {
			return nil, fmt.Errorf("decode to interface %s: %w", to.Name(), err)
		}
		return result, nil
	}

	if to == reflect.TypeOf(decimal.Decimal{}) {
		// Handle converts types to decimal.Decimal
		num, err := decodeDecimal(data)
		if err != nil {
			return nil, fmt.Errorf("decode decimal: %w", err)
		}
		return num, nil
	}

	if to == reflect.TypeOf(time.Time{}) {
		// Handle converts types to time.Time
		t, err := time.Parse(time.RFC3339, data.(string))
		if err != nil {
			return nil, fmt.Errorf("decode time: %w", err)
		}
		return t, nil
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
