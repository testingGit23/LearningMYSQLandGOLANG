package new

import (
	"LearningMYSQLandGOLANG/visualBase/pkg/opendb"
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
	opendb.Tmpl.ExecuteTemplate(w, "New", nil)
}
