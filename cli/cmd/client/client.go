package client

import (
	"fmt"

	"github.com/jdxj/yuque/cli/cmd"

	"github.com/urfave/cli/v2"
)

func NewClient() *cli.Command {
	cmd := &cli.Command{
		Name:                   "client",
		Aliases:                nil,
		Usage:                  "query api limit",
		UsageText:              "",
		Description:            "",
		ArgsUsage:              "",
		Category:               "",
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 clientAction,
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  clientFlags(),
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
	return cmd
}

func clientFlags() []cli.Flag {
	return nil
}

func clientAction(ctx *cli.Context) error {
	_, err := cmd.Yuque.AuthenticatedUser()
	if err != nil {
		return err
	}

	fmt.Printf("limit: %d\n", cmd.Yuque.XRateLimitLimit())
	fmt.Printf("remaining: %d\n", cmd.Yuque.XRateLimitRemaining())
	return nil
}
