package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type GetLocationRequest struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Response struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Label         string  `json:"label"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	Street        string  `json:"street"`
	PostalCode    string  `json:"postal_code"`
	Region        string  `json:"region"`
	RegionCode    string  `json:"region_code"`
	Neighbourhood string  `json:"neighbourhood"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
}

type LocationApiResponse struct {
	Data []Response `json:"data"`
}

func (h *Handler) GetLocation(c *gin.Context) {
	var l GetLocationRequest

	if err := c.BindJSON(&l); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse request. check documentation: https://github.com/gwuah/tinderclone/blob/master/Readme.MD",
		})
		return
	}

	response, err := MakeRequest(fmt.Sprintf("http://api.positionstack.com/v1/reverse?access_key=%s&query=%s,%s&limit=%d",
		os.Getenv("ACCESS_KEY"), l.Latitude, l.Longitude, 1), os.Getenv("PORT"), nil, "GET")

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to make request"})
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to process request"})
		return
	}

	var locationApiResponse LocationApiResponse
	err = json.Unmarshal(body, &locationApiResponse)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to process request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "location received",
		"location": locationApiResponse,
	})

}
