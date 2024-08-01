package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppHandlers(t *testing.T) {
	app := NewApp(8080)
	ts := httptest.NewServer(app.router)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("GET request returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedGet := "get request"
	body := readBody(res)
	if body != expectedGet {
		t.Errorf("GET request returned unexpected body: got %v want %v", body, expectedGet)
	}

	res, err = http.Post(ts.URL, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusCreated {
		t.Errorf("POST request returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	expectedPost := "post request"
	body = readBody(res)
	if body != expectedPost {
		t.Errorf("POST request returned unexpected body: got %v want %v", body, expectedPost)
	}
}

func readBody(res *http.Response) string {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
