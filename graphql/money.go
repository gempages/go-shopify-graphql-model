package graphql

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func MarshalMoney(v decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		b, err := v.MarshalJSON()
		if err != nil {
			logrus.Debugf("marshal Money value: %v", err)
		}
		w.Write(b)
	})
}

func UnmarshalMoney(v interface{}) (decimal.Decimal, error) {
	var (
		d   decimal.Decimal
		err error
	)

	switch v := v.(type) {
	case string:
		d, err = decimal.NewFromString(v)
	case *string:
		if v == nil {
			return decimal.Decimal{}, nil
		}
		d, err = decimal.NewFromString(*v)
	default:
		err = fmt.Errorf("%T is not a string or *string and cannot be unmarshalled into Money", v)
	}

	if err != nil {
		return decimal.Decimal{}, err
	}
	return d, nil
}
