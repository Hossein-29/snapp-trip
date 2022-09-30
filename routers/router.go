package routers

import (
	"example/snapp/databases"
	"example/snapp/models"
	"example/snapp/validations"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRule(c *gin.Context) {
	var rules []models.Rule
	err := c.BindJSON(&rules)
	var ruleReport models.RuleResponse
	if err != nil {
		fmt.Printf("CreateRule: %s", err.Error())
		ruleReport.Status = "FAILED"
		ruleReport.Message = err.Error()

		// for debugging and more readability
		// c.IndentedJSON(http.StatusBadRequest, ruleReport)

		// for best performance
		c.JSON(http.StatusBadRequest, ruleReport)

		return
	}

	isValid := validations.ValidateRule(rules)

	if !isValid {
		ruleReport.Status = "FAILED"
		ruleReport.Message = "UNVALID RULE"

		// for debugging and more readability
		//c.IndentedJSON(http.StatusBadRequest, ruleReport)

		// for best performance
		c.JSON(http.StatusBadRequest, ruleReport)

		return
	}

	ruleReport.Status = "SUCCESS"

	// for testing and more readability
	// c.IndentedJSON(http.StatusOK, ruleReport)

	// for best performance
	c.JSON(http.StatusOK, ruleReport)

	databases.CreateRuleTable(rules)

}

func CreateTicket(c *gin.Context) {
	var tickets []models.Ticket
	err := c.BindJSON(&tickets)
	if err != nil {
		c.String(http.StatusBadRequest, "UNVALID TICKET")
		fmt.Printf("CreateTicket: %s", err.Error())
		return
	}

	//c.IndentedJSON(http.StatusOK, tickets)

	var ticketresponses []models.TicketResponse

	for i := range tickets {
		temp := databases.MatchTicket(tickets[i])
		ticketresponses = append(ticketresponses, temp)
	}

	// for debugging and more readability
	// c.IndentedJSON(http.StatusOK, ticketresponses)

	// for best performance
	c.JSON(http.StatusOK, ticketresponses)
}

func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s :)", name)
}
