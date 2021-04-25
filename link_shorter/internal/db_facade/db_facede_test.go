package db_facade

import (
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDbFacade(t *testing.T) {
	conf := common.NewConfig()
	dbFacade, err := InitDbFacade(&conf.DataBase)
	if err != nil {
		t.Fatal(err)
	}
	defer dbFacade.Db_.Close()
	_, _ = dbFacade.Db_.Begin()
	if err = dbFacade.Db_.Ping(); err != nil {
		t.Fatal(err)
	}
	full_link := "Test_Fake_link_" + time.Now().String()
	new_link, _ := dbFacade.CreateShortLink(full_link)
	assert.Equal(t, new_link.Link_, full_link)

	old_link, _ := dbFacade.GetLink(model.DefaultLink + new_link.ShortLinkKey_)
	assert.Equal(t, new_link.Link_, old_link.Link_)
}
