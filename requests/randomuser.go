package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/m7shapan/njson"
)

// Parameters is a structure to store requested params
type Parameters struct {
	// date range
	FromTime time.Time
	ToTime   time.Time
	// requested amount of users
	Results int32
}

// Randomuser used to store and manage randomuser response and "testapi" server response
type Randomuser struct {
	// structure to map randomuser API response
	// "njson" tag used to unmarshal randomuser response
	// "json" tag used to marshal "testapi" server response
	DataContainer struct {
		Data []struct {
			Gender    string `json:"gender" njson:"gender"`
			FirstName string `json:"first_name" njson:"name.first"`
			LastName  string `json:"last_name" njson:"name.last"`
			Postcode  int32  `json:"postcode" njson:"location.postcode"`
			CreatedAt string `json:"created_at" njson:"registered.date"`
		} `json:"data" njson:"results"`
	}
	// represents request parameters
	Parameters Parameters
}

// getURL builds link (https://randomuser.me/api?...) for get request
// can be extended with other parameters
func (r *Randomuser) getURL() string {
	var url = "https://randomuser.me/api?"
	if r.Parameters.Results > 0 {
		url += fmt.Sprintf("results=%d", r.Parameters.Results)
	}
	return url
}

// checkTimeSpan helps to check if datetime is between "from" and "to" datetimes
// converts ISO 8601
func (r *Randomuser) checkTimeSpan(check string) bool {
	checkTime, err := time.Parse(time.RFC3339Nano, check)
	if err != nil {
		return false
	}
	if !r.Parameters.FromTime.IsZero() && !r.Parameters.ToTime.IsZero() {
		return checkTime.After(r.Parameters.FromTime) && checkTime.Before(r.Parameters.ToTime)
	} else if !r.Parameters.FromTime.IsZero() {
		return checkTime.After(r.Parameters.FromTime)
	} else if !r.Parameters.ToTime.IsZero() {
		return checkTime.Before(r.Parameters.ToTime)
	} else {
		return false
	}
}

// parseJSON parses json (returned from randomuser.me/api) and
// unmarshals it to DataContainer by "njson" field tags
func (r *Randomuser) parseJSON(jsonByte []byte) error {
	err := njson.Unmarshal(jsonByte, &r.DataContainer)
	return err
}

// filter checks if any filtration parameters applied (time range a.t.m.)
// can be extended to filter by oher params
func (r *Randomuser) filter() error {
	if !r.Parameters.FromTime.IsZero() || !r.Parameters.FromTime.IsZero() {
		var randomuser = Randomuser{}
		for _, data := range r.DataContainer.Data {
			if r.checkTimeSpan(data.CreatedAt) {
				randomuser.DataContainer.Data = append(randomuser.DataContainer.Data, data)
			}
		}
		r.DataContainer = randomuser.DataContainer
	}
	return nil
}

// ParseFromTime converts ISO 8601 "from" datetime string to time.Time ojectect
func (r *Randomuser) ParseFromTime(timeString string) error {
	if timeString == "" {
		return nil
	}
	fromTime, err := time.Parse(time.RFC3339Nano, timeString)
	if err != nil {
		return errors.New("message: invalid \"from\" time format, please use ISO 8601 format")
	}
	r.Parameters.FromTime = fromTime
	return nil
}

// ParseToTime converts ISO 8601 "to" datetime string to time.Time ojectect
func (r *Randomuser) ParseToTime(timeString string) error {
	if timeString == "" {
		return nil
	}
	toTime, err := time.Parse(time.RFC3339Nano, timeString)
	if err != nil {
		return errors.New("message: invalid \"to\" time format, please use ISO 8601 format")
	}
	r.Parameters.ToTime = toTime
	return nil
}

// ParseResults converts results (amount of users) parameter to int32 and saves it
func (r *Randomuser) ParseResults(results string) error {
	if results == "" {
		return nil
	}
	intResults, err := strconv.Atoi(results)
	if err != nil {
		return errors.New("message: \"results\" must be a number")
	}
	r.Parameters.Results = int32(intResults)
	return nil
}

// Request sends request to randomuser api
// and calls parseJson and filter
func (r *Randomuser) Request() error {
	response, err := http.Get(r.getURL())
	responseBody, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseBody))
	r.parseJSON(responseBody)
	r.filter()
	defer response.Body.Close()
	return err
}

// JSONResponse returns json from DataContainer
// marshalled by "json" field tags
func (r *Randomuser) JSONResponse() ([]byte, error) {
	jsonByte, err := json.Marshal(&r.DataContainer)
	return jsonByte, err
}
