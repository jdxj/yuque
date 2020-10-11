package user

import (
	"fmt"

	"github.com/jdxj/yuque/cli/cmd"

	"github.com/urfave/cli/v2"
)

func NewUser() *cli.Command {
	cmd := &cli.Command{
		Name:                   "user",
		Aliases:                nil,
		Usage:                  "user command used for get user information",
		UsageText:              "",
		Description:            "",
		ArgsUsage:              "",
		Category:               "",
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 userAction,
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  userFlags(),
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

func userFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "id",
			Aliases:     nil,
			Usage:       "id or path",
			EnvVars:     nil,
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			DefaultText: "",
			Destination: nil,
			HasBeenSet:  false,
		},
	}
}

func userAction(ctx *cli.Context) error {
	id := ctx.String("id")
	if id == "" {
		us, err := cmd.Yuque.AuthenticatedUser()
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", us)
		return nil
	}

	us, err := cmd.Yuque.IndividualUser(id)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", us)
	return nil
}
