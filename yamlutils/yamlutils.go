package yamlutils

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"../logger"
	"gopkg.in/yaml.v2"
)

// Get gets a YAML object that corresponds to path from file
func Get(file string, path ...string) (interface{}, error) {
	file = "config/" + file + ".yaml"
	if _, err := os.Stat(file); err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(file)
	if logger.Error(err, "Could not read "+file) {
		return nil, err
	}

	var parsed interface{}
	err = yaml.Unmarshal(b, &parsed)

	if logger.Error(err, "Could not parse "+file) {
		return nil, err
	}

	for _, p := range path {
		parsedMap, ok := parsed.(map[interface{}]interface{})
		if !ok {
			logger.Error(errors.New("yaml object is not map"), "YAML object is not map")
			return nil, err
		}
		parsed, ok = parsedMap[p]
		if !ok {
			logger.Error(errors.New("yaml map does not contain key \""+p+"\""), "YAML map does not contain key \""+p+"\"")
			return nil, err
		}
	}
	return parsed, nil
}

// Set puts a YAML object to the path in file
func Set(toSet interface{}, file string, path ...string) error {
	file = "config/" + file + ".yaml"

	parsed := new(map[interface{}]interface{})
	if _, err := os.Stat(file); err == nil {
		b, err := ioutil.ReadFile(file)
		if logger.Error(err, "Could not read "+file) {
			return err
		}

		err = yaml.Unmarshal(b, &parsed)
		if logger.Error(err, "Could not parse "+file) {
			return err
		}
	}

	curr := new(interface{})
	*curr = parsed
	prev := new(map[interface{}]interface{})
	for i := 0; i < len(path); i++ {
		currMap, ok := (*curr).(map[interface{}]interface{})
		if !ok {
			logger.Warning(errors.New("overwriting yaml object with map"), "Overwriting YAML object with map")
			currMap = *new(map[interface{}]interface{})
			*curr = currMap
		}

		*prev = (*curr).(map[interface{}]interface{})
		*curr, ok = currMap[path[i]]
		if !ok {
			*curr = new(map[interface{}]interface{})
		}
	}
	*curr = toSet

	b, err := yaml.Marshal(*parsed)
	if logger.Error(err, "Could not convert data into YAML") {
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)
	if logger.Error(err, "Could write to "+file) {
		return err
	}

	return nil
}

// GetToken gets the bot token
func GetToken() string {
	token, err := Get("config", "token")

	if err != nil || token == nil {
		token = input("Bot token: ")
		Set(token, "config", "token")
	}

	return token.(string)
}

// GetPrefix gets the bot prefix
func GetPrefix() string {
	prefix, err := Get("config", "prefix")

	if err != nil || prefix == nil {
		prefix = input("Bot prefix: ")
		Set(prefix, "config", "prefix")
	}

	return prefix.(string)
}

// GetOwnerID gets the bot owner ID
func GetOwnerID() string {
	owner, err := Get("config", "owner")

	if err != nil || owner == nil {
		owner = input("Bot owner ID: ")
		Set(owner, "config", "owner")
	}

	return owner.(string)
}

func input(s string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(s)
	text, err := reader.ReadString('\n')
	logger.Fatal(err, "Unable to read input from stdin")
	return text[:len(text)-1]
}
