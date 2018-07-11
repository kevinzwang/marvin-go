package commands

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"../logger"
	"github.com/bwmarrin/discordgo"
	"github.com/lunny/html2md"
)

// Anime gives info about the queried anime
type Anime struct{}

type query struct {
	Data struct {
		Media queryMedia
	}
	Errors []queryError
}

type queryMedia struct {
	Title struct {
		UserPreferred string
	}
	SiteURL     string
	Description string
	Format      string
	Status      string
	MeanScore   int
	Rankings    []queryMediaRank
	Genres      []string
	CoverImage  struct {
		Medium string
	}
	NextAiringEpisode queryMediaNextAiringEpisode
}

type queryMediaNextAiringEpisode struct {
	TimeUntilAiring int
	Episode         int
}

type queryMediaRank struct {
	Rank    int
	AllTime bool
}

type queryError struct {
	Message   string
	Status    int
	Locations []queryErrorLocation
}

type queryErrorLocation struct {
	Line   int
	Column int
}

func (cmd *Anime) execute(ctx *Context, args []string) {
	queryString := `query GetRelevantAnime ($search: String) {
			Media (search: $search, type: ANIME, sort: [POPULARITY_DESC]) {
				title {
					userPreferred
				}
				siteUrl
				description
				format
				status
				meanScore
				rankings {
					rank
					allTime
				}
				genres
				coverImage {
					medium
				}
				nextAiringEpisode {
					timeUntilAiring
					episode
				}
			}
		}`

	body, err := json.Marshal(map[string]interface{}{
		"query":     queryString,
		"variables": map[string]string{"search": strings.Join(args, " ")},
	})

	if logger.Error(err, "Problem formatting query to JSON") {
		ctx.send("Could not format query, please try again.")
		return
	}

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewReader(body))
	if logger.Error(err, "Could not access Anilist API") {
		ctx.send("Problem accessing Anilist, please try again.")
		return
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if logger.Error(err, "Could not get response") {
		ctx.send("Problem getting response from Anilist, please try again.")
		return
	}

	var q query
	err = json.Unmarshal(respBody, &q)
	if logger.Error(err, "Could not parse JSON") {
		ctx.send("Problem parsing JSON, please try again.")
		return
	}
	if q.Errors != nil {
		ctx.reply("No results matched your query.")
		return
	}

	m := q.Data.Media

	descLen := 256
	if len(m.Description) > descLen {
		m.Description = m.Description[:descLen-3] + "..."
	}

	allTimePop := 0
	for _, p := range m.Rankings {
		if p.AllTime {
			allTimePop = p.Rank
			break
		}
	}
	if allTimePop < 1 && len(m.Rankings) > 0 {
		allTimePop = m.Rankings[0].Rank
	}

	genres := ""
	for i, g := range m.Genres {
		genres += "[" + g + "](https://anilist.co/search/anime?includedGenres=" + url.QueryEscape(g) + ")"
		if i != len(m.Genres)-1 {
			genres += "  |  "
		}
	}

	format := ""
	switch m.Format {
	case "TV_SHORT":
		format = "TV Short"
	case "MOVIE":
		format = "Movie"
	case "SPECIAL":
		format = "Special"
	case "MUSIC":
		format = "Music"
	default:
		format = m.Format
	}

	var status string
	// here we turn the integer for next airing episode to a string for days, hours, and minutes
	if m.NextAiringEpisode != (queryMediaNextAiringEpisode{}) {
		if m.NextAiringEpisode.Episode == 1 {
			status = "Primieres in"
		} else {
			status = "Ep " + strconv.Itoa(m.NextAiringEpisode.Episode) + " in"
		}

		// prevCounted used so that things like 4d 0h 2m will be displayed
		prevCounted := false
		remainder := m.NextAiringEpisode.TimeUntilAiring

		// days
		if remainder/86400 != 0 {
			status += " " + strconv.Itoa(remainder/86400) + "d"
			remainder %= 86400
			prevCounted = true
		}

		// hours
		if remainder/3600 != 0 || prevCounted {
			status += " " + strconv.Itoa(remainder/3600) + "h"
			remainder %= 3600
			prevCounted = true
		}

		status += " " + strconv.Itoa(remainder/60) + "m"
	} else {
		status = strings.Title(strings.ToLower(strings.Replace(m.Status, "_", " ", -1)))
	}

	em := discordgo.MessageEmbed{
		URL:         m.SiteURL,
		Title:       m.Title.UserPreferred,
		Description: html2md.Convert(strings.Replace(strings.Replace(m.Description, "<br>", "", -1), "\n", " ", -1)),
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: m.CoverImage.Medium},
		Color:       0x44b5f0,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{Name: "Format", Value: format, Inline: true},
			&discordgo.MessageEmbedField{Name: "Status", Value: status, Inline: true},
			&discordgo.MessageEmbedField{Name: "Score", Value: strconv.Itoa(m.MeanScore) + "%", Inline: true},
			&discordgo.MessageEmbedField{Name: "Popularity", Value: "#" + strconv.Itoa(allTimePop), Inline: true},
			&discordgo.MessageEmbedField{Name: "Genres", Value: genres, Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{Text: "Fetched from Anilist.co"},
	}
	ctx.sendEmbed(&em)
}

func (cmd *Anime) description() string { return "gives info about the queried anime" }

func (cmd *Anime) category() string { return "weeb" }

func (cmd *Anime) numArgs() (int, int) { return 1, -1 }

func (cmd *Anime) names() []string { return []string{"anime"} }

func (cmd *Anime) onlyOwner() bool { return false }

func (cmd *Anime) usage() []string { return []string{"<Anime Title>"} }
