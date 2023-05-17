package apisix

import (
	"encoding/json"

	"github.com/infraboard/mcube/tools/pretty"
)

func NewReponseList() *ReponseList {
	return &ReponseList{
		List: []*Reponse{},
	}
}

type ReponseList struct {
	Total int        `json:"total"`
	List  []*Reponse `json:"list"`
}

func (r *ReponseList) String() string {
	return pretty.ToJSON(r)
}

func (r *ReponseList) Values(l Lister) {
	for i := range r.List {
		item := r.List[i]
		l.Add(item.Value)
	}
}

func NewReponse(v any) *Reponse {
	return &Reponse{}
}

type Reponse struct {
	ModifiedIndex int             `json:"modifiedIndex"`
	Key           string          `json:"key"`
	Value         json.RawMessage `json:"value"`
	CreatedIndex  int             `json:"createdIndex"`
}

type Lister interface {
	Add(item json.RawMessage)
}
