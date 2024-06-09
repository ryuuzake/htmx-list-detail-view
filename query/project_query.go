package query

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/ryuuzake/htmx-list-detail-view/model"
)

func ProjectQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&model.Project{})
}

func GetListProject(dao *daos.Dao) ([]*model.Project, error) {
	projects := []*model.Project{}

	err := ProjectQuery(dao).AndWhere(dbx.HashExp{"parent": ""}).OrderBy("created asc").All(&projects)

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func GetListProjectByParentId(dao *daos.Dao, id string) ([]*model.Project, error) {
	projects := []*model.Project{}

	err := ProjectQuery(dao).AndWhere(dbx.HashExp{"parent": id}).OrderBy("created asc").All(&projects)

	if err != nil {
		return nil, err
	}

	return projects, nil
}
