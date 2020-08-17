package saycommand

import (
	"github.com/3sixtyied/rpdSix/commands"
	"github.com/ztrue/tracerr"
)

func Initialize() {
	commands.AddCommand(
		commands.Command{
			Run:                         run,
			Names:                       []string{"say"},
			ExpectedPositionalArguments: []string{"toSay"},
		},
	)
}

func run(ctx commands.CommandContext) {
	_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, ctx.Arguments["toSay"])
	if err != nil {
		tracerr.PrintSourceColor(err)
	}
}
