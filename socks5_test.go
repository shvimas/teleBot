package teleBot

import (
	"net/http"
	"testing"
)

func check(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

func TestSendViaSOCKS5(t *testing.T) {
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	check(err, t)
	resp, err := SendSOCKS5Request(req, "<enter proxy address with port>", nil, nil)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error("got", resp.Status)
	}
}
