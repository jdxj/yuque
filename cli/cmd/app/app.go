package app

import (
	"time"

	"github.com/jdxj/yuque/cli/cmd"
	cmdClient "github.com/jdxj/yuque/cli/cmd/client"
	"github.com/jdxj/yuque/cli/cmd/repo"
	"github.com/jdxj/yuque/cli/cmd/user"
	"github.com/jdxj/yuque/client"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:                   "yuque",
		HelpName:               "",
		Usage:                  "",
		UsageText:              "",
		ArgsUsage:              "",
		Version:                "",
		Description:            "",
		Commands:               commands(),
		Flags:                  appFlags(),
		EnableBashCompletion:   false,
		HideHelp:               false,
		HideHelpCommand:        false,
		HideVersion:            false,
		BashComplete:           nil,
		Before:                 before,
		After:                  nil,
		Action:                 nil,
		CommandNotFound:        nil,
		OnUsageError:           nil,
		Compiled:               time.Time{},
		Authors:                nil,
		Copyright:              "",
		Writer:                 nil,
		ErrWriter:              nil,
		ExitErrHandler:         nil,
		Metadata:               nil,
		ExtraInfo:              nil,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}
	return app
}

func appFlags() []cli.Flag {
	return []cli.Flag{}
}

func before(ctx *cli.Context) error {
	token := ctx.String("token")
	cmd.Yuque = client.New(token)
	return nil
}

func commands() []*cli.Command {
	return []*cli.Command{
		user.NewUser(),
		cmdClient.NewClient(),
		repo.Repo(),
	}
}
