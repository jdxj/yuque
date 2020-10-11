package repo

import (
	"fmt"

	"github.com/jdxj/yuque/utils"

	"github.com/jdxj/yuque/client"

	"github.com/urfave/cli/v2"
)

func create() *cli.Command {
	return &cli.Command{
		Name:                   "create",
		Aliases:                nil,
		Usage:                  "create repository",
		UsageText:              "",
		Description:            "",
		ArgsUsage:              "",
		Category:               "",
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 createAction,
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  createFlags(),
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}

func createFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "token",
			Aliases:     nil,
			Usage:       "yuque token",
			EnvVars:     nil,
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			DefaultText: "",
			Destination: nil,
			HasBeenSet:  false,
		},

		&cli.StringFlag{
			Name:        "id",
			Aliases:     nil,
			Usage:       "user id/user path",
			EnvVars:     nil,
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			DefaultText: "",
			Destination: nil,
			HasBeenSet:  false,
		},

		&cli.StringFlag{
			Name:        "name",
			Aliases:     nil,
			Usage:       "repository name",
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

		&cli.StringFlag{
			Name:        "description",
			Aliases:     nil,
			Usage:       "repository description",
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

		&cli.IntFlag{
			Name:        "public",
			Aliases:     nil,
			Usage:       "repository access permission",
			EnvVars:     nil,
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			Value:       client.Private,
			DefaultText: "0",
			Destination: nil,
			HasBeenSet:  false,
		},

		&cli.StringFlag{
			Name:        "type",
			Aliases:     nil,
			Usage:       "type of repository. Book, Design, All",
			EnvVars:     nil,
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			TakesFile:   false,
			Value:       client.Book,
			DefaultText: client.Book,
			Destination: nil,
			HasBeenSet:  false,
		},

		&cli.StringFlag{
			Name:        "position",
			Aliases:     nil,
			Usage:       "position of repository created. user or group",
			EnvVars:     nil,
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			TakesFile:   false,
			Value:       "user",
			DefaultText: "user",
			Destination: nil,
			HasBeenSet:  false,
		},
	}
}

func createAction(ctx *cli.Context) (err error) {
	token := ctx.String("token")
	yuque := client.New(token)

	id := ctx.String("id")
	crp := &client.CreateRepoParams{
		Name:        ctx.String("name"),
		Slug:        utils.GenRandString(6),
		Description: ctx.String("description"),
		Public:      ctx.Int("public"),
		Type:        ctx.String("type"),
	}

	var bds *client.BookDetailSerializer

	position := ctx.String("position")
	switch position {
	case "group":
		bds, err = yuque.CreateGroupRepository(id, crp)
	default:
		bds, err = yuque.CreateUserRepository(id, crp)
	}

	if bds != nil {
		fmt.Printf("%s\n", bds)
	}
	return
}
