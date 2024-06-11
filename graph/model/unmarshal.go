package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
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
		mediaType, err := concludeObjectType(typeName)
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
		if typeName, ok := discount["__typename"].(string); ok {
			discountType, err := concludeObjectType(typeName)
			if err != nil {
				return fmt.Errorf("conclude object type: %w", err)
			}
			disc := reflect.New(discountType).Interface()
			err = mapstructure.Decode(discount, disc)
			if err != nil {
				return fmt.Errorf("decode discount node: %w", err)
			}
			dc.Discount = disc.(Discount)
			return nil
		}
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

func decodeFile(node map[string]interface{}) (File, error) {
	if typeName, ok := node["__typename"].(string); ok {
		fileType, err := concludeObjectType(typeName)
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
