package bec

import (
	"testing"
)

// ////////////////////////////////////////////
// node API
// ////////////////////////////////////////////
func TestGetBecAvailableNodeInfoVo(t *testing.T) {
	res, err := CLIENT.GetBecAvailableNodeInfoVo("vm")
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}
