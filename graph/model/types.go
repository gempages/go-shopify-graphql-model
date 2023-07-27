package model

import (
	"fmt"
	"reflect"
	"regexp"
)

var gidRegex *regexp.Regexp

func init() {
	gidRegex = regexp.MustCompile(`^gid://shopify/(\w+)/\d+$`)
}

func concludeObjectType(gid string) (reflect.Type, error) {
	submatches := gidRegex.FindStringSubmatch(gid)
	if len(submatches) != 2 {
		return reflect.TypeOf(nil), fmt.Errorf("malformed gid=`%s`", gid)
	}
	resource := submatches[1]
	switch resource {
	case "MediaImage":
		return reflect.TypeOf(MediaImage{}), nil
	case "Video":
		return reflect.TypeOf(Video{}), nil
	case "Model3d":
		return reflect.TypeOf(Model3d{}), nil
	case "ExternalVideo":
		return reflect.TypeOf(ExternalVideo{}), nil
	default:
		return reflect.TypeOf(nil), fmt.Errorf("`%s` not implemented type", resource)
	}
}
