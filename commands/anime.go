package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Anime gives info about the queried anime
type Anime struct{}

func (cmd *Anime) execute(ctx *Context, args []string) {
	query := strings.Join(args, " ")
	searchResp, err := http.Get("https://api.jikan.moe/search/anime/" + query)
	if err != nil {
		ctx.send("Problem searching for anime, please try again.")
		return
	}

	defer searchResp.Body.Close()

	searchBody, err := ioutil.ReadAll(searchResp.Body)
	var searchParsed map[string]interface{}
	err = json.Unmarshal(searchBody, &searchParsed)
	if err != nil {
		ctx.send("Problem parsing JSON, please try again.")
		return
	}

	searchResults, ok := searchParsed["result"].([]interface{})
	if !ok {
		ctx.send("Problem parsing JSON, please try again.")
		return
	}

	if len(searchResults) == 0 {
		ctx.reply("No results matched your query.")
		return
	}

	firstResult, ok := searchResults[0].(map[string]interface{})
	if !ok {
		ctx.send("Problem parsing JSON, please try again.")
		return
	}

	resultID, ok := firstResult["id"].(float64)
	if !ok {
		ctx.send("Problem parsing JSON, please try again.")
		return
	}

	resultResp, err := http.Get("https://api.jikan.moe/anime/" + strconv.Itoa(int(resultID)))
	if err != nil {
		ctx.send("Problem finding anime, please try again.")
		return
	}

	defer resultResp.Body.Close()

	resultBody, err := ioutil.ReadAll(resultResp.Body)
	if err != nil {
		ctx.send("Problem parsing JSON, please try again.")
		return
	}

	var resultParsed map[string]interface{}
	err = json.Unmarshal(resultBody, &resultParsed)
	if err != nil {
		ctx.send("Problem parsing JSON, please try again.")
		return
	}

	em := new(discordgo.MessageEmbed)

	em.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: resultParsed["image_url"].(string)}
	title := resultParsed["title"].(string)
	if resultParsed["title_english"] != nil {
		title += " (English: " + resultParsed["title_english"].(string) + ")"
	}
	em.Title = title
	em.URL = resultParsed["link_canonical"].(string)

	desc := resultParsed["synopsis"].(string)
	if len(desc) > 200 {
		desc = desc[:200]
		desc = desc[:strings.LastIndex(desc, " ")]
		desc += " [...]"
	}
	em.Description = desc

	genres := ""
	genreList := resultParsed["genre"].([]interface{})
	for _, g := range genreList {
		genres += "[" + g.(map[string]interface{})["name"].(string) + "](" + g.(map[string]interface{})["url"].(string) + ") | "
	}
	genres = genres[:len(genres)-3]

	em.Fields = []*discordgo.MessageEmbedField{
		&discordgo.MessageEmbedField{Name: "Episodes", Value: strconv.Itoa(int(resultParsed["episodes"].(float64))), Inline: true},
		&discordgo.MessageEmbedField{Name: "Status", Value: resultParsed["status"].(string), Inline: true},
		&discordgo.MessageEmbedField{Name: "Score", Value: strconv.FormatFloat(resultParsed["score"].(float64), 'f', 1, 64), Inline: true},
		&discordgo.MessageEmbedField{Name: "Popularity", Value: strconv.Itoa(int(resultParsed["popularity"].(float64))), Inline: true},
		&discordgo.MessageEmbedField{Name: "Genres", Value: genres, Inline: true},
	}

	em.Footer = &discordgo.MessageEmbedFooter{Text: "Searched using the Jikan API for MyAnimeList.net"}

	ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, em)
}

func (cmd *Anime) description() string { return "gives info about the queried anime" }

func (cmd *Anime) category() string { return "misc" }

func (cmd *Anime) numArgs() (int, int) { return 1, -1 }

func (cmd *Anime) names() []string { return []string{"anime"} }

func (cmd *Anime) onlyOwner() bool { return false }

func (cmd *Anime) usage() []string { return []string{""} }
