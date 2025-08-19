package myob_test

import (
	"encoding/json"
	"testing"
)

func TestGeneralJournalPost(t *testing.T) {
	req := client.NewGeneralJournalPostRequest()
	body := req.NewRequestBody()
	req.SetRequestBody(body)
	b, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))

	//resp, err := req.Do(context.Background())
	//if err != nil {
	//
	//	t.Fatal(err)
	//}

}
