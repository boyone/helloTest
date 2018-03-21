package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloAPI(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/hello/espresso")

	if res.StatusCode != 200 {
		t.Errorf("Expect status code == 200, but got %d", res.StatusCode)
	}

	if str, _ := ioutil.ReadAll(res.Body); string(str) != "Hello, espresso" {
		fmt.Println(string(str))
		t.Errorf("Expect response body == Hello, espresso, but got %v", string(str))
	}
}
