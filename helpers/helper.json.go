package helpers

import (
	"encoding/json"

	"github.com/fleetimee/flee/schemes"
	"github.com/sirupsen/logrus"
)

func Strigify(payload interface{}) []byte {
	response, _ := json.Marshal(payload)
	return response
}

func Parse(payload []byte) schemes.SchemeResponses {
	var jsonResponse schemes.SchemeResponses
	err := json.Unmarshal(payload, &jsonResponse)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return jsonResponse
}
