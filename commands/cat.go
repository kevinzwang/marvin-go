package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type catQuery struct {
	File string
}

// Cat shows an image of a cat
type Cat struct{}

func (cmd *Cat) execute(ctx *Context, args []string) {
	resp, err := http.Get("http://aws.random.cat/meow")
	if err != nil {
		ctx.send("Sorry, could not get cat image :( Try again.")
		return
	}

	defer resp.Body.Close()

	inBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.send("Sorry, could not read cat image :( Try again.")
		return
	}
	var bodyJSON catQuery
	json.Unmarshal(inBytes, &bodyJSON)
	// fmt.Println(string(inBytes))
	ctx.send(bodyJSON.File)

}

func (cmd *Cat) description() string { return "finds a random picture of a cat" }

func (cmd *Cat) category() string { return "fun" }

func (cmd *Cat) numArgs() (int, int) { return 0, 0 }

func (cmd *Cat) names() []string { return []string{"cat"} }

func (cmd *Cat) onlyOwner() bool { return false }

func (cmd *Cat) usage() []string { return []string{""} }
