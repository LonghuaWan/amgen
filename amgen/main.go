package main

import (
	"github.com/chxfantasy/cmgen"
	"gopkg.in/urfave/cli.v2"
	"os"
)

func main() {

	app := &cli.App{
		Name:  "amgen",
		Usage: "code generate for mgo",
		Commands: []*cli.Command{
			//{Name: "interface", Usage: "create model interface go file", Flags: defaultInterfaceFlag(), Action: cmgen.InterfaceAction},
			{Name: "mgo", Usage: "generate golang code", Flags: defaultModelFlag(), Action: cmgen.MgoAction},
		},
		Version: "0.1.0",
	}

	app.Run(os.Args)
}
