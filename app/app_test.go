package app_test

import (
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/ringsaturn/hertz-template-test-panic/app"
)

func TestAppHello(t *testing.T) {
	h := app.NewHertz()
	w := ut.PerformRequest(h.Engine, "GET", "/hello", nil)
	resp := w.Result()
	t.Log(string(resp.BodyBytes()))
	assert.DeepEqual(t, http.StatusOK, resp.StatusCode())
}
