package annotations

import "github.com/grafana/grafana/pkg/components/simplejson"

type Repository interface {
	Save(item *Item) error
	Find(query *ItemQuery) ([]*Item, error)
}

type ItemQuery struct {
	OrgId   int64    `json:"orgId"`
	From    int64    `json:"from"`
	To      int64    `json:"from"`
	Type    ItemType `json:"type"`
	AlertId int64    `json:"alertId"`

	Limit int64 `json:"alertId"`
}

var repositoryInstance Repository

func GetRepository() Repository {
	return repositoryInstance
}

func SetRepository(rep Repository) {
	repositoryInstance = rep
}

type ItemType string

const (
	AlertType ItemType = "alert"
)

type Item struct {
	Id        int64    `json:"id"`
	OrgId     int64    `json:"orgId"`
	Type      ItemType `json:"type"`
	Title     string   `json:"title"`
	Text      string   `json:"text"`
	Metric    string   `json:"metric"`
	AlertId   int64    `json:"alertId"`
	UserId    int64    `json:"userId"`
	PrevState string   `json:"prevState"`
	NewState  string   `json:"newState"`
	Epoch     int64    `json:"epoch"`

	Data *simplejson.Json `json:"data"`
}