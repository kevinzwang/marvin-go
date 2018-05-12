package commands

import (
	"strings"
)

// If is a command for creating bot rules
type If struct{}

func (cmd *If) execute(ctx *Context, args []string) {
	// first split it between condition and action
	var condition, action []string
	for i, a := range args {
		if strings.HasSuffix(a, ",") {
			args[i] = args[i][:len(args[i])-1]
			condition = args[:i+1]
			if args[i+1] == "then" {
				i++
			}
			action = args[i+1:]
			break
		}
		if a == "then" {
			condition = args[:i]
			action = args[i+1:]
			break
		}
	}

	// then split by "and" and put into rules
	rule := make([][][]string, 2)
	rule[0] = sliceSplit(condition, "and")
	rule[1] = sliceSplit(action, "and")
	rules = append(rules, rule)

	em := embedRule(rule, len(rules)-1)
	em.Description = "Successfully made new rule!"
	ctx.sendEmbed(&em)

}

func sliceSplit(sl []string, field string) (ret [][]string) {
	prev := 0
	for i, s := range sl {
		if s == field {
			ret = append(ret, sl[prev:i])
			prev = i + 1
		}
	}
	ret = append(ret, sl[prev:])
	return
}

func (cmd *If) description() string { return "Uses fake NLP to make cool custom actions" }

func (cmd *If) category() string { return "misc" }

func (cmd *If) numArgs() (int, int) { return 1, -1 }

func (cmd *If) names() []string { return []string{"if"} }

func (cmd *If) onlyOwner() bool { return true }

func (cmd *If) usage() []string { return []string{"<condition>, then <action>"} }
