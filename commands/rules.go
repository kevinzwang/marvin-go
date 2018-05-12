package commands

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Rules is used to manage rules created with the If command
type Rules struct{}

func (cmd *Rules) execute(ctx *Context, args []string) {
	if len(args) == 0 {
		if len(rules) == 0 {
			ctx.send("Currently no rules set.")
			return
		}

		msg := "__Rules:__\n"
		for i, rule := range rules {
			condition := listifyRulePart(rule[0], "", " and ")
			condition = condition[:len(condition)-5]
			action := listifyRulePart(rule[1], "", " and ")
			action = action[:len(action)-5]
			msg += "\n`" + strconv.Itoa(i) + ". if " + condition + ", then " + action + "`"
		}
		ctx.send(msg)
	} else if len(args) == 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil || i >= len(rules) || i < 0 {
			ctx.wrongUsage("rules")
			return
		}

		em := embedRule(rules[i], i)
		ctx.sendEmbed(&em)
	} else if len(args) == 2 && args[0] == "remove" {
		i, err := strconv.Atoi(args[1])
		if err != nil || i >= len(rules) || i < 0 {
			ctx.wrongUsage("rules")
			return
		}

		rules = append(rules[:i], rules[i+1:]...)
		ctx.reply("Successfully deleted rule #" + args[1])
	} else {
		ctx.wrongUsage("rules")
	}
}

func embedRule(rule [][][]string, num int) (em discordgo.MessageEmbed) {
	em = discordgo.MessageEmbed{
		Title: "Rule #" + strconv.Itoa(num),
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Conditions",
				Value:  listifyRulePart(rule[0], "- `", "`\n"),
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Actions",
				Value:  listifyRulePart(rule[1], "- `", "`\n"),
				Inline: false,
			},
		},
	}
	return
}

func listifyRulePart(pt [][]string, before string, after string) (s string) {
	for _, sl := range pt {
		s += before + strings.Join(sl, " ") + after
	}
	return
}

func (cmd *Rules) description() string { return "used to manage rules created with the If command" }

func (cmd *Rules) category() string { return "admin" }

func (cmd *Rules) numArgs() (int, int) { return 0, 2 }

func (cmd *Rules) names() []string { return []string{"rules"} }

func (cmd *Rules) onlyOwner() bool { return true }

func (cmd *Rules) usage() []string { return []string{"", "<number>", "remove <number>"} }
