package http_server

import (
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/db_facade"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpServer_GetLink(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/HELLO", nil)

	//todo: Добавить тестовый контур для фасада.
	facade, err := db_facade.InitDbFacade(&common.NewConfig().DataBase)
	if err != nil {
		t.Fatal(err)
	}
	defer facade.Db_.Close()
	s := newHttpServer(facade)
	s.ServeHttp(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
