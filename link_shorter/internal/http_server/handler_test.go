package http_server

import (
	"bytes"
	"encoding/json"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/db_facade"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHttpServer(t *testing.T) {

	//todo: Добавить тестовый контур для фасада.
	facade, err := db_facade.InitDbFacade(&common.NewConfig().DataBase)
	if err != nil {
		t.Fatal(err)
	}
	defer facade.Db_.Close()
	s := newHttpServer(facade)

	full_link := "{\"Link\": \"Test_Fake_link_" + time.Now().String() + "\""

	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(struct {
		Url string
	}{Url: full_link})
	set_rec := httptest.NewRecorder()
	set_req, _ := http.NewRequest(http.MethodPost, "/SetLink", b)
	s.router.ServeHTTP(set_rec, set_req)

	get_rec := httptest.NewRecorder()
	json.NewEncoder(b).Encode(struct {
		ShortUrl string
	}{ShortUrl: model.DefaultLink + set_rec.Body.String()})
	get_req, _ := http.NewRequest(http.MethodPost, "/GetLink", b)
	s.router.ServeHTTP(get_rec, get_req)

	assert.Equal(t, get_rec.Body.String(), full_link)

}
