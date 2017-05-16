package controller

import (
	"app/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func jsonRequest(url string, p interface{}) (*http.Request, error) {
	content, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	req.Header["Content-Type"] = []string{"application/json"}
	req.Header["Accept"] = []string{"application/json"}
	return req, nil
}

func TestLoginGet(t *testing.T) {
	_, err := model.UserCreate("alice@example.com", "secret")
	if err != nil {
		t.Error(err)
	}

	e := echo.New()

	cases := []struct {
		form LoginPostForm
		ok   bool
	}{
		{LoginPostForm{"alice@example.com", "secret"}, true},
		{LoginPostForm{"bob@example.com", "secret"}, false},
	}

	for _, cs := range cases {
		req, err := jsonRequest("/login", cs.form)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()
		c := e.NewContext(req, resp)

		err = LoginPost(c)
		if cs.ok {
			if err != nil {
				t.Fatal("unexpected error")
			}
			if resp.Code != http.StatusOK {
				t.Fatal("unexpected status code")
			}
		} else {
			if err != nil {
				t.Fatal("unexpected error")
			}
			if resp.Code == http.StatusOK {
				t.Fatal("unexpected status code", resp.Code, "with", cs.form)
			}
		}
	}
}
