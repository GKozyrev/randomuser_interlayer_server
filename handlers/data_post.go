package handlers

import (
	"log"
	"net/http"
	"testapi/requests"

	"github.com/labstack/echo/v4"
)

// DataPostRequestModel represents server request structure
type DataPostRequestModel struct {
	FromTime string `json:"from"`
	ToTime   string `json:"to"`
}

// DataPostResponseModel represents server response structure
type DataPostResponseModel struct {
	Status   string `json:"status"`
	FromTime string `json:"from"`
	ToTime   string `json:"to"`
}

// DataPost represents POST /data endpoint
// Expected:
// --> { from:<ISO8601 date (string)>, to: <ISO8061 date (string)> }
// <-- no validation, just seding POST request to "https://reqbin.com/echo/post/json" and returning the incoming params + status: success
func DataPost(c echo.Context) error {

	var requestModel = &DataPostRequestModel{}

	err := c.Bind(requestModel)

	if err != nil {
		log.Print(err)
		return err
	}

	err = requests.ReqbinEchoPost()
	if err != nil {
		log.Print(err)
		return err
	}

	return c.JSON(http.StatusOK, DataPostResponseModel{
		Status:   "Success",
		FromTime: requestModel.FromTime,
		ToTime:   requestModel.ToTime,
	})
}
