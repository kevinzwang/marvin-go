package commands

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Cat shows an image of a cat
type Cat struct{}

func (cmd *Cat) execute(ctx *Context, args []string) {
	resp, err := http.Get("http://thecatapi.com/api/images/get?format=html")
	if err != nil {
		ctx.send("Sorry, could not get cat image :( Try again.")
		return
	}

	defer resp.Body.Close()

	inBytes, err := ioutil.ReadAll(resp.Body)
	body := string(inBytes)
	body = body[strings.Index(body, "src")+3:]
	body = body[strings.Index(body, "http"):]
	body = body[:strings.LastIndex(body, "\"")]
	ctx.send(body)

}

func (cmd *Cat) description() string { return "finds a random picture of a cat" }

func (cmd *Cat) category() string { return "fun" }

func (cmd *Cat) numArgs() (int, int) { return 0, 0 }

func (cmd *Cat) names() []string { return []string{"cat"} }

func (cmd *Cat) onlyOwner() bool { return false }

func (cmd *Cat) usage() []string { return []string{""} }
