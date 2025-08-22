package myob_test

import (
	"context"
	"encoding/json"
	"testing"
)

func TestGeneralJournalsGet(t *testing.T) {
	req := client.NewGeneralJournalsGetRequest()
	req.QueryParams().Top.Set(10)
	resp, err := req.Do(context.Background())
	if err != nil {

		t.Fatal(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	t.Log(string(b))
}
