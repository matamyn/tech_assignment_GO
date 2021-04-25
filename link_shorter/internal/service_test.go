package link_shorter

import (
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLinkShorterService_HandleHello(t *testing.T) {
	s := linkShorterService(common.NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/HELLO", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello")
}
