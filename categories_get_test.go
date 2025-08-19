package myob_test

import (
	"context"
	"encoding/json"
	"testing"
)

func TestCategoriesGet(t *testing.T) {
	req := client.NewCategoriesGetRequest()
	req.QueryParams().Top.Set(10)
	resp, err := req.All(context.Background())
	if err != nil {

		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	t.Log(string(b))
}
