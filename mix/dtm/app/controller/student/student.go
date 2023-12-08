package student

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DetailAction (c *gin.Context) {
	studentId := c.DefaultQuery("id", "0")
	name := c.DefaultQuery("name", "默认")

	c.JSON(200, gin.H{
		"id": studentId,
		"name": name,
	})
}

type UpdateParam struct  {
	Id     string `form:"id" json:"id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

type HeaderParam struct {
	PCD   string    `header:"Pcd"`
}


func UpdateAction (c *gin.Context) {
	var p UpdateParam
	var h HeaderParam
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}


	c.JSON(200, gin.H{
		"id": p.Id,
		"name": p.Name,
		"pcd": h.PCD,
	})
}

func DetailTplAction (c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}