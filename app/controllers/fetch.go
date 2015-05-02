package controllers

import (
	"github.com/etcinit/central/app/applications"
	"github.com/etcinit/central/app/documents"
	"github.com/etcinit/central/database/models"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
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

	pretty.Println(token)

	if err != nil {
		c.JSON(200, gin.H{
			"status": "error",
			"messages": []string{
				"Unknown application",
			},
		})
		return
	}

	documents, err := f.DocumentFinder.FindAllByApplication(application)

	if err != nil {
		c.JSON(200, gin.H{
			"status": "error",
			"messages": []string{
				"Error fetching files",
			},
		})
		return
	}

	files := map[string]string{}
	for _, document := range documents {
		files[document.Alias] = document.Contents
	}

	c.JSON(200, gin.H{
		"status":      "success",
		"application": application,
		"files":       files,
	})

}
