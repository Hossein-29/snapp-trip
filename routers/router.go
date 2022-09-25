package routers

import (
	"example/snapp/databases"
	"example/snapp/models"
	"example/snapp/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRule(c *gin.Context) {
	var rules []models.Rule
	err := c.BindJSON(&rules)
	var ruleReport models.RuleResponse
	if err != nil {
		ruleReport.Status = "FAILED"
		ruleReport.Message = err.Error()
		c.IndentedJSON(http.StatusBadRequest, ruleReport)
		return
	}

	isValid := validations.ValidateRule(rules)

	if !isValid {
		ruleReport.Status = "FAILED"
		ruleReport.Message = err.Error()
		c.IndentedJSON(http.StatusBadRequest, ruleReport)
		return
	}

	ruleReport.Status = "SUCCESS"
	c.IndentedJSON(http.StatusOK, ruleReport)

	databases.CreateRuleTable(rules)

}

func CreateTicket(c *gin.Context) {
	var tickets []models.Ticket
	err := c.BindJSON(&tickets)
	if err != nil {
		return
	}

	//c.IndentedJSON(http.StatusOK, tickets)

	var ticketresponses []models.TicketResponse

	for i := range tickets {
		temp := databases.MatchTicket(tickets[i], c)
		ticketresponses = append(ticketresponses, temp)
	}

	c.IndentedJSON(http.StatusOK, ticketresponses)

}

func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s :)", name)
}
