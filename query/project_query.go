package query

import (
	"fmt"

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

func GetListProjectWithChildren(dao *daos.Dao) ([]*model.ProjectList, error) {
	projects := []*model.Project{}

	err := ProjectQuery(dao).OrderBy("created asc").All(&projects)

	if err != nil {
		return nil, err
	}

	projectsList := mapProject(projects)

	return projectsList, nil
}

func mapProject(source []*model.Project) []*model.ProjectList {
	projectMap := make(map[string]*model.ProjectList)

	for _, item := range source {
		projectList := &model.ProjectList{
			Project: *item,
		}
		projectMap[item.Id] = projectList

		if item.Parent != "" {
			parentID := item.Parent
			parentNode, exists := projectMap[parentID]
			if exists {
				parentNode.Children = append(parentNode.Children, projectList)
			} else {
				// Handle case where parent doesn't exist (optional)
				fmt.Printf("Parent node with ID %s not found for child %s\n", parentID, item.Id)
			}
		}
	}

	var result []*model.ProjectList
	for _, project := range projectMap {
		if project.Parent == "" {
			result = append(result, project)
		}
	}

	return result
}

func GetProjectById(dao *daos.Dao, id string) (*model.Project, error) {
	project := &model.Project{}

	err := ProjectQuery(dao).AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(project)

	if err != nil {
		return nil, err
	}

	return project, nil
}
