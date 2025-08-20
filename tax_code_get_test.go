package myob_test

import (
	"context"
	"encoding/json"
	"testing"
)

func TestTaxCodesGet(t *testing.T) {
	req := client.NewTaxCodesGetRequest()
	req.QueryParams().Top.Set(10)
	resp, err := req.All(context.Background())
	if err != nil {

		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	t.Log(string(b))
}
