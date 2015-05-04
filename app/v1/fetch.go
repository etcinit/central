package v1

import (
	"github.com/etcinit/central/app/applications"
	"github.com/etcinit/central/app/documents"
	"github.com/etcinit/central/app/responses"
	"github.com/etcinit/central/database/models"
	"github.com/gin-gonic/gin"
)

// FetchController provides routes for fetching configuration.
type FetchController struct {
	DocumentFinder    *documents.Finder    `inject:""`
	ApplicationFinder *applications.Finder `inject:""`
}

// Register registers all the routes for this controller
func (f *FetchController) Register(r *gin.RouterGroup) {
	r.GET("/fetch", f.getFetch)
}

func (f *FetchController) getFetch(c *gin.Context) {
	tokenRaw, err := c.Get("token")
	if err != nil {
		panic("Missing token")
	}

	token := tokenRaw.(*models.ApplicationToken)
	application, err := f.ApplicationFinder.FindByToken(token)

	if err != nil {
		responses.SendInvalidInputMessages(c, "Unknown application")
		return
	}

	documents, err := f.DocumentFinder.FindAllByApplication(application)

	if err != nil {
		responses.SendInvalidInputMessages(c, "Unable to find any files")
		return
	}

	files := map[string]string{}
	for _, document := range documents {
		if document.Alias != "" {
			files[document.Alias] = document.Contents
		} else {
			files[document.Name] = document.Contents
		}
	}

	c.JSON(200, gin.H{
		"status":      "success",
		"application": application,
		"files":       files,
	})

}
