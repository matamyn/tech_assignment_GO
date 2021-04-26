package http_server

import (
	"bytes"
	"encoding/json"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/db_facade"
	pb "github.com/matamyn/tech_assignment_GO/link_shorter/internal/model"
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

	full_link := "Test_Fake_link_" + time.Now().String()

	set_b := &bytes.Buffer{}

	_ = json.NewEncoder(set_b).Encode(struct {
		Url string
	}{Url: full_link})

	set_res := httptest.NewRecorder()
	set_req, _ := http.NewRequest(http.MethodPost, "/SetLink", set_b)
	s.router.ServeHTTP(set_res, set_req)

	get_b := &bytes.Buffer{}
	_ = json.NewEncoder(get_b).Encode(struct {
		ShortUrl string
	}{ShortUrl: pb.DefaultLink + set_res.Body.String()})

	get_rec := httptest.NewRecorder()

	get_req, _ := http.NewRequest(http.MethodPost, "/GetLink", get_b)
	s.router.ServeHTTP(get_rec, get_req)

	assert.Equal(t, get_rec.Body.String(), full_link)
}
