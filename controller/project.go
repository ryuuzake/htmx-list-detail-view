package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/ryuuzake/htmx-list-detail-view/query"
	"github.com/ryuuzake/htmx-list-detail-view/utils"
	"github.com/ryuuzake/htmx-list-detail-view/view/project"
)

type ProjectHandler struct{ Handler }

func NewProjectHandler(app *pocketbase.PocketBase) ProjectHandler {
	return ProjectHandler{NewHandler(app)}
}

func (h ProjectHandler) GetProjects(c echo.Context) error {
	projectList, err := query.GetListProjectWithChildren(h.app.Dao())

	if err != nil {
		return apis.NewBadRequestError("", err)
	}

	return utils.Render(project.ProjectView(projectList, nil, false), c)
}

func (h ProjectHandler) GetProjectById(c echo.Context) error {
	id := c.PathParam("id")

	result, err := query.GetProjectById(h.app.Dao(), id)

	if err != nil {
		return apis.NewNotFoundError("Project Not Found", err)
	}

	return utils.Render(project.ProjectDetailView(result), c)
}
