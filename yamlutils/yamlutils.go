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
func Get(strict bool, file string, path ...string) (interface{}, error) {
	file = "config/" + file + ".yaml"
	if _, err := os.Stat(file); err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		if strict {
			logger.Error(err, "Could not read "+file)
		}
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
			if strict {
				logger.Error(errors.New("yaml object is not map"), "YAML object is not map")
			}
			return nil, err
		}
		parsed, ok = parsedMap[p]
		if !ok {
			if strict {
				logger.Error(errors.New("yaml map does not contain key \""+p+"\""), "YAML map does not contain key \""+p+"\"")
			}
			return nil, err
		}
	}
	return parsed, nil
}

// Set puts a YAML object to the path in file
func Set(toSet interface{}, file string, path ...string) error {
	file = "config/" + file + ".yaml"

	var parsed map[interface{}]interface{}
	if _, err := os.Stat(file); err == nil {
		b, err := ioutil.ReadFile(file)
		if logger.Error(err, "Could not read "+file) {
			return err
		}

		err = yaml.Unmarshal(b, &parsed)
		if logger.Error(err, "Could not parse "+file) {
			return err
		}
	} else {
		parsed = make(map[interface{}]interface{})
	}

	curr := &parsed
	for i := 0; i < len(path)-1; i++ {
		next, ok := (*curr)[path[i]].(map[interface{}]interface{})
		if !ok {
			if (*curr)[path[i]] != nil {
				logger.Warning(errors.New(""), "Overwriting YAML object with map")
			}
			next = make(map[interface{}]interface{})
			(*curr)[path[i]] = next
		}
		curr = &next
	}
	(*curr)[path[len(path)-1]] = toSet

	b, err := yaml.Marshal(parsed)
	if logger.Error(err, "Could not convert data into YAML") {
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)
	if logger.Error(err, "Could not write to "+file) {
		return err
	}

	return nil
}

// GetToken gets the bot token
func GetToken() string {
	token, err := Get(true, "config", "token")

	if err != nil || token == nil {
		token = input("Bot token: ")
		Set(token, "config", "token")
	}

	return token.(string)
}

// GetPrefix gets the bot prefix
func GetPrefix(serverID string) (string, bool) {
	prefix, err := Get(false, "config", serverID, "prefix")

	if prefix == nil {
		prefix, err = Get(false, "config", "global", "prefix")
	}

	if logger.Warning(err, "Couldn't get prefix for server.") {
		return "", false
	}

	p, ok := prefix.(string)
	if !ok {
		return "", false
	}
	return p, true
}

// SetPrefix sets the prefix for the specified server
func SetPrefix(serverID string, prefix string) bool {
	err := Set(prefix, "config", serverID, "prefix")
	if err != nil {
		return false
	}
	return true
}

// GetAdmin gets the admins for the server
func GetAdmin(serverID string) []string {
	admin, err := Get(false, "config", serverID, "admin")

	logger.Warning(err, "Couldn't get admins for server.")

	return admin.([]string)
}

// GetOwnerID gets the owner id set in config.yaml
func GetOwnerID() string {
	owner, err := Get(false, "config", "owner")
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
