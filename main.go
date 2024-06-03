package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var config string
	var local string

	cmds := []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install v2ray",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "config",
					Aliases:     []string{"c"},
					Destination: &config,
				},
				&cli.StringFlag{
					Name:        "local",
					Aliases:     []string{"l"},
					Destination: &local,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				err = downloadBinaryFile(local)
				if err != nil {
					return err
				}
				err = installBinaryFile()
				if err != nil {
					return err
				}
				err = installConfig(config)
				if err != nil {
					return err
				}
				return installService()
			},
		},
		{
			Name:  "uninstall",
			Usage: "Remove config,cache and uninstall v2ray",
			Action: func(ctx *cli.Context) (err error) {
				err = uninstallService()
				if err != nil {
					return err
				}
				return uninstallBinaryFile()
			},
		},
		{
			Name:  "update",
			Usage: "Update v2ray",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "config",
					Aliases:     []string{"c"},
					Destination: &config,
				},
				&cli.StringFlag{
					Name:        "local",
					Aliases:     []string{"l"},
					Destination: &local,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				err = updateBinaryFile(local)
				if err != nil {
					return err
				}
				if config != "" {
					err = installConfig(config)
					if err != nil {
						return err
					}
				}
				return updateService()
			},
		},
		{
			Name:  "reload",
			Usage: "Reload service",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "config",
					Aliases:     []string{"c"},
					Destination: &config,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				if config != "" {
					err = installConfig(config)
					if err != nil {
						return err
					}
				}
				return reloadService()
			},
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s", cCtx.App.Version)
	}

	app := &cli.App{
		Usage:    "v2ray quick install tool",
		Version:  "v3.02",
		Commands: cmds,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
