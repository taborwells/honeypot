// Copyright (c) 2022 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the AGPLv3 license, a copy of
// which may be found at https://github.com/silverton-io/honeypot/blob/main/LICENSE

package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/silverton-io/honeypot/pkg/response"
)

func TestHealthcheckHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	HealthcheckHandler(c)

	resp := rec.Result()

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(`HealthcheckHandler returned status code %v, want %v`, resp.StatusCode, http.StatusOK)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	marshaledB, _ := json.Marshal(response.Ok)
	equiv := reflect.DeepEqual(b, marshaledB)
	if !equiv {
		t.Fatalf(`HealthcheckHandler returned body %v, want %v`, b, marshaledB)
	}
}
