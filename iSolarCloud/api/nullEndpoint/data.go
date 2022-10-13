package nullEndpoint

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"time"
)


const Url = "%URL%"
const Disabled = true

type RequestData struct {
	// DeviceType string `json:"device_type" required:"true"`
}

// IsValid Checks for validity of results data.
func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

// Help provides more info to the user on request JSON fields.
func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

// ResultData holds data returned from the API.
type ResultData struct {
	Dummy string `json:"dummy"`
}

// IsValid Checks for validity of results data.
func (e *ResultData) IsValid() error {
	var err error
	switch {
		case e.Dummy == "":
			break
		default:
			err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	}
	return err
}

// type DecodeResultData ResultData
//
// func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
// }


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", time.Time{})
	}

	return entries
}