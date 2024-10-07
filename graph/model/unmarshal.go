package model

import (
	"encoding/json"
	"fmt"
)

func (e *MediaEdge) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	result, err := Decode(m, &MediaEdge{})
	if err != nil {
		return fmt.Errorf("decode MediaEdge: %w", err)
	}
	*e = *result.(*MediaEdge)
	return nil
}

func (c *MediaConnection) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	result, err := Decode(m, &MediaConnection{})
	if err != nil {
		return fmt.Errorf("decode MediaConnection: %w", err)
	}
	*c = *result.(*MediaConnection)
	return nil
}

func (s *WebhookSubscription) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	result, err := Decode(m, &WebhookSubscription{})
	if err != nil {
		return fmt.Errorf("decode WebhookSubscription: %w", err)
	}
	*s = *result.(*WebhookSubscription)
	return nil
}

func (dc *DiscountNode) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	result, err := Decode(m, &DiscountNode{})
	if err != nil {
		return fmt.Errorf("decode DiscountNode: %w", err)
	}
	*dc = *result.(*DiscountNode)
	return nil
}

func (fe *FileEdge) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	result, err := Decode(m, &FileEdge{})
	if err != nil {
		return fmt.Errorf("decode FileEdge: %w", err)
	}
	*fe = *result.(*FileEdge)
	return nil
}

func (fp *FileCreatePayload) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	result, err := Decode(m, &FileCreatePayload{})
	if err != nil {
		return fmt.Errorf("decode FileCreatePayload: %w", err)
	}
	*fp = *result.(*FileCreatePayload)
	return nil
}

func (fc *FileConnection) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	result, err := Decode(m, &FileConnection{})
	if err != nil {
		return fmt.Errorf("decode FileConnection: %w", err)
	}
	*fc = *result.(*FileConnection)
	return nil
}

func (d *DiscountAutomaticNode) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	result, err := Decode(m, &DiscountAutomaticNode{})
	if err != nil {
		return fmt.Errorf("decode DiscountAutomaticNode: %w", err)
	}
	*d = *result.(*DiscountAutomaticNode)
	return nil
}

func (ap *AppPlanV2) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	result, err := Decode(m, &AppPlanV2{})
	if err != nil {
		return fmt.Errorf("decode AppPlanV2: %w", err)
	}
	*ap = *result.(*AppPlanV2)
	return nil
}
