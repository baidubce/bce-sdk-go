package types

import (
	"encoding/json"
	"testing"
)

func TestContainerNetworkConfig_CreateClusterIPPoolFieldsJSON(t *testing.T) {
	ippoolMinAllocateIPs := 10
	bceCustomerMaxIP := 64

	config := ContainerNetworkConfig{
		IPPoolMinAllocateIPs: &ippoolMinAllocateIPs,
		BCECustomerMaxIP:     &bceCustomerMaxIP,
	}

	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("marshal ContainerNetworkConfig failed: %v", err)
	}

	var got struct {
		IPPoolMinAllocateIPs int `json:"ippoolMinAllocateIPs"`
		BCECustomerMaxIP     int `json:"bceCustomerMaxIp"`
	}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal ContainerNetworkConfig failed: %v", err)
	}

	if got.IPPoolMinAllocateIPs != ippoolMinAllocateIPs {
		t.Fatalf("ippoolMinAllocateIPs = %d, want %d", got.IPPoolMinAllocateIPs, ippoolMinAllocateIPs)
	}
	if got.BCECustomerMaxIP != bceCustomerMaxIP {
		t.Fatalf("bceCustomerMaxIp = %d, want %d", got.BCECustomerMaxIP, bceCustomerMaxIP)
	}
}
