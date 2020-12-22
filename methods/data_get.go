package methods

import (
	"fmt"
	"log"
	"net/http"
	"testapi/requests"

	"github.com/labstack/echo/v4"
)

var err error

// DataGet represents GET /data endpoint
// Expected:
// --> ? results=<number> & from=<ISO8601 date> & to=<ISO8061 date>
// <-- marshalled by "json" field requests.Randomuser.DataContainer
func DataGet(c echo.Context) error {
	var randomuser = requests.Randomuser{}

	err = randomuser.ParseResults(c.QueryParam("results"))
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": fmt.Sprint(err)})
	}
	err = randomuser.ParseFromTime(c.QueryParam("from"))
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": fmt.Sprint(err)})
	}
	err = randomuser.ParseToTime(c.QueryParam("to"))
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": fmt.Sprint(err)})
	}
	err = randomuser.Request()
	if err != nil {
		log.Print(err)
		return err
	}

	response := randomuser.DataContainer

	if err != nil {
		log.Print(err)
		return err
	}

	return c.JSON(http.StatusOK, response)
}
