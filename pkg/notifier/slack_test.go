package notifier

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRocket_Post(t *testing.T) {
	fields := []Field{
		{Name: "name1", Value: "value1"},
		{Name: "name2", Value: "value2"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var payload = SlackPayload{}
		err = json.Unmarshal(b, &payload)

		if payload.Attachments[0].AuthorName != "podinfo.test" {
			t.Fatal("wrong author name")
		}

		if len(payload.Attachments[0].Fields) != len(fields) {
			t.Fatal("wrong field")
		}
	}))
	defer ts.Close()

	rocket, err := NewRocket(ts.URL, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	err = rocket.Post("podinfo", "test", "test", fields, "error")
	if err != nil {
		t.Fatal(err)
	}
}
