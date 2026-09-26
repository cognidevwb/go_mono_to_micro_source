package payments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Gateway calls the external payment provider.
type Gateway struct {
	baseURL string
	client  *http.Client
}

// NewGateway builds a provider client.
func NewGateway(baseURL string) *Gateway {
	return &Gateway{baseURL: baseURL, client: &http.Client{Timeout: 5 * time.Second}}
}

// Charge asks the provider to capture amount for orderID.
func (g *Gateway) Charge(orderID uint, amount float64) error {
	body, _ := json.Marshal(map[string]any{"order": orderID, "amount": amount})
	resp, err := g.client.Post(g.baseURL+"/v1/charges", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("provider declined: %d", resp.StatusCode)
	}
	return nil
}
