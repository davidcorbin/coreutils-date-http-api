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
	Date string `json:"date" binding:"required"`
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

	args := []string{"date", "--set", dateStr}
	_, err = exec.Command("sudo", args...).Output()
	if err != nil {
		return err
	}

	log.Println("System date set to:", dateStr)
	return nil
}

func SetDate(c *gin.Context) {
	var dateInput DateInput
	err := c.BindJSON(&dateInput)
	if err != nil {
		log.Println(err.Error())
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	err = setDate(dateInput.Date, c)
	if err != nil {
		log.Println(err.Error())
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "true"})
}
