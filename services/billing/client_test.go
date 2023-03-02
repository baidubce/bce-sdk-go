package billing

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/billing/api"
)

var (
	ak     = os.Getenv("baidu_ak")
	sk     = os.Getenv("baidu_sk")
	region = ""
	cli, _ = NewClient(ak, sk, region)
)

func TestClient_GetBalance(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "test",
			want:    "demo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cli
			got, err := c.GetBalance()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetBalance() = %v, want %v", got.CashBalance, tt.want)
			}
		})
	}
}

func TestClient_GetBilling(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "test",
			want:    "demo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cli
			got, err := c.GetBilling(&api.BillingParams{
				Month:       "2023-02",
				ProductType: "postpay",
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetBilling() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("Client.GetBilling() = %v, want %v", got, tt.want)
				fmt.Printf("%+v", got)
			}
		})
	}
}
