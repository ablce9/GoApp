package http

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"bytes"

	"encoding/json"

	"gitlab.com/zenport.io/go-assignment/engine"
	"gitlab.com/zenport.io/go-assignment/providers/database"
)

var (
	router http.Handler
)

func TestMain(m *testing.M) {
	provider := database.NewProvider()
	// todo: init database (ex: create table, clear previous data, etc.)
	e := engine.NewEngine(provider)

	router = nil // todo: add your router

	code := m.Run()
	provider.Close()
	os.Exit(code)
}

func TestPostKnightBipolelm(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/knight", bytes.NewBuffer([]byte(`{"name":"Bipolelm","strength":10,"weapon_power":20}`)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusCreated)
	}
}

func TestPostKnightElrynd(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/knight", bytes.NewBuffer([]byte(`{"name":"Elrynd","strength":10,"weapon_power":50}`)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusCreated)
	}
}

func TestPostKnightBadData(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/knight", bytes.NewBuffer([]byte(`{"name":"FAILED"}`)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusBadRequest)
	}

	response := map[string]interface{}{}

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if _, has := response["code"]; !has {
		t.Fatal("Response error: Expected code field")
	}

	if _, has := response["message"]; !has {
		t.Fatal("Response error: Expected message field")
	}
}

func TestPostKnightBadType(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/knight", bytes.NewBuffer([]byte(`name:"Bipolelm"`)))
	req.Header.Add("Content-Type", "text/plain")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusBadRequest)
	}
}

func TestGetKnights(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/knight", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusOK)
	}

	var response []map[string]interface{}

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if len(response) != 2 {
		t.Fatal("Response error: Expected 2 knights")
	}

	knight := response[0]

	if _, has := knight["id"]; !has {
		t.Fatal("Response error: Expected id field in knight object")
	}
	if _, has := knight["name"]; !has {
		t.Fatal("Response error: Expected name field in knight object")
	}
	if _, has := knight["strength"]; !has {
		t.Fatal("Response error: Expected strength field in knight object")
	}
	if _, has := knight["weapon_power"]; !has {
		t.Fatal("Response error: Expected weapon_power field in knight object")
	}

	if response[0]["id"].(string) == response[1]["id"].(string) {
		t.Fatal("Response error: Expected not same id for each knights")
	}
}

func TestGetKnightNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/knight/123456789", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusNotFound)
	}

	response := map[string]interface{}{}

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if _, has := response["code"]; !has {
		t.Fatal("Response error: Expected code field")
	}

	if _, has := response["message"]; !has {
		t.Fatal("Response error: Expected message field")
	}

	if response["message"].(string) != "Knight #123456789 not found." {
		t.Fatal("Response error: Expected error message 'Knight #123456789 not found.'")
	}
}
