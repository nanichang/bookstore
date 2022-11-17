package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(req *http.Request, res interface{}) {
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), res); err != nil{
			return
		}
	}
}