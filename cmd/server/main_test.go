package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"thirsty-jotter/pkg/api/models"

	"github.com/stretchr/testify/assert"
)

func TestGetPlantsRoutes(t *testing.T) {
	router := serveApplication()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/plants/2", nil)

	router.ServeHTTP(w, req)

	mockResponse, err := json.Marshal(models.Plant{Id: "2", Name: "sunflower", Description: "mica", Type: "flower", Image: ""})
	if err != nil {
		t.Fatal(err)
	}

	var actualResponse models.Plant
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Fatal(err)
	}

	actualJSON, err := json.Marshal(actualResponse)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, string(mockResponse), string(actualJSON))

}
