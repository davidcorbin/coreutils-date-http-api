package internal

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type DateInput struct {
	date string
}

func getDateString() string {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(fmt.Sprintf("%s", out))
}

func GetDate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"date": getDateString()})
}

func setDate(dateStr string, c *gin.Context) error {
	_, err := time.Parse(time.UnixDate, dateStr)
	if err != nil {
		return err
	}

	out, err := exec.Command("date -s " + dateStr).Output()
	if err != nil {
		return err
	}

	log.Println("System date set to:", out)
	return nil
}

func SetDate(c *gin.Context) {
	var dateInput DateInput
	err := c.ShouldBindJSON(&dateInput)
	if err != nil {
		log.Println(err.Error())
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	err = setDate(dateInput.date, c)
	if err != nil {
		log.Println(err.Error())
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "true"})
}
