package routers

import (
	"example/snapp/databases"
	"example/snapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRule(c *gin.Context) {
	var rules []models.Rule
	err := c.BindJSON(&rules)
	var ruleReport models.Report
	if err != nil {
		ruleReport.Status = "FAILED"
		ruleReport.Message = err.Error()
		c.IndentedJSON(http.StatusBadRequest, ruleReport)
	} else {
		ruleReport.Status = "SUCCESS"
		c.IndentedJSON(http.StatusOK, ruleReport)
	}

	databases.Db.Create(&rules)

}

func CreateTicket(c *gin.Context) {
	var tickets []models.Ticket
	err := c.BindJSON(&tickets)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, tickets)

}

func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s :)", name)
}
