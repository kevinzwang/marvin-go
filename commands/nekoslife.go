package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../logger"
)

func nekoslife(category string) string {
	resp, err := http.Get("https://nekos.life/api/v2/img/" + category)
	if logger.Error(err, "Could not access nekos.life API") {
		return "Problem getting " + category + " image, please try again."
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var parsed map[string]string
	err = json.Unmarshal(body, &parsed)
	if logger.Error(err, "Could not parse JSON") {
		return "Problem parsing JSON, please try again."
	}

	return parsed["url"]
}
