package repo

import (
	"github.com/urfave/cli/v2"
)

func Repo() *cli.Command {
	return &cli.Command{
		Name:                   "repo",
		Aliases:                nil,
		Usage:                  "repo is used to manipulate the knowledge base",
		UsageText:              "",
		Description:            "repo description",
		ArgsUsage:              "",
		Category:               "",
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		OnUsageError:           nil,
		Subcommands:            repoSubcommands(),
		Flags:                  repoFlags(),
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}

func repoFlags() []cli.Flag {
	return []cli.Flag{}
}

func repoSubcommands() []*cli.Command {
	return []*cli.Command{
		create(),
	}
}
