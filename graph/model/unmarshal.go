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
