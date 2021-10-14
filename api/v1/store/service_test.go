package store

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSetPair(t *testing.T) {
	url := "/store"
	e := echo.New()
	pairRequest := `{"key":"tk1","value":"tv1"}`
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(pairRequest))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, set(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":\"key-value set successfully\"}\n", rec.Body.String())
	}
}

func TestGetPair_WhenKeyExists(t *testing.T) {
	url := "/store"
	e := echo.New()

	// insert dummy data
	pairRequest1 := `{"key":"tk2","value":"tv2"}`
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(pairRequest1))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, set(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":\"key-value set successfully\"}\n", rec.Body.String())
	}

	// test
	req, _ = http.NewRequest(http.MethodGet, url+"?key=tk2", nil)
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":{\"tk2\":\"tv2\"}}\n", rec.Body.String())
	}
}

func TestGetPair_WhenKeyNotExists(t *testing.T) {
	url := "/store"
	e := echo.New()

	// insert dummy data
	pairRequest1 := `{"key":"tk3","value":"tv3"}`
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(pairRequest1))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, set(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":\"key-value set successfully\"}\n", rec.Body.String())
	}

	// test
	req, _ = http.NewRequest(http.MethodGet, url+"?key=tk12", nil)
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":null}\n", rec.Body.String())
	}
}

func TestGetPair_WhenKeyQueryParamNotExists(t *testing.T) {
	url := "/store"
	e := echo.New()

	// insert dummy data
	pairRequest1 := `{"key":"tk4","value":"tv4"}`
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(pairRequest1))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, set(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":\"key-value set successfully\"}\n", rec.Body.String())
	}

	// test
	req, _ = http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":null}\n", rec.Body.String())
	}
}

func TestFlush(t *testing.T) {
	url := "/store"
	e := echo.New()

	// insert dummy data
	pairRequest1 := `{"key":"tk5","value":"tv5"}`
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(pairRequest1))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, set(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":\"key-value set successfully\"}\n", rec.Body.String())
	}

	// insert dummy data
	pairRequest2 := `{"key":"tk6","value":"tv6"}`
	req, _ = http.NewRequest(http.MethodPost, url, strings.NewReader(pairRequest2))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if assert.NoError(t, set(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":\"key-value set successfully\"}\n", rec.Body.String())
	}

	// validation
	req, _ = http.NewRequest(http.MethodGet, url+"?key=tk5", nil)
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if assert.NoError(t, get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":{\"tk5\":\"tv5\"}}\n", rec.Body.String())
	}

	// test
	req, _ = http.NewRequest(http.MethodPut, url+"/flush", nil)
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, flush(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":\"key-value flush successfully\"}\n", rec.Body.String())
	}

	// final validation
	req, _ = http.NewRequest(http.MethodGet, url+"?key=tk5", nil)
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if assert.NoError(t, get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"code\":200,\"error\":false,\"errors\":null,\"result\":null}\n", rec.Body.String())
	}
}
