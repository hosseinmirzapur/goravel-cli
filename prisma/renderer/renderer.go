package renderer

import (
	"log"
	"os"

	"github.com/hosseinmirzapur/goravel-cli/prisma/cli"
	"github.com/hosseinmirzapur/goravel-cli/prisma/logger"
	"github.com/hosseinmirzapur/goravel-cli/utils"
)

func Render(args []string) {
	logger.Debug.Printf("invoking command %+v", args)

	switch args[0] {
	case "prefetch":
		// just run prisma -v to trigger the download
		if err := cli.Run([]string{"-v"}, true); err != nil {
			utils.Error("prisma", "invoking prisma command error", err)
		}
		os.Exit(0)
		return
	case "init":
		// override default init flags
		args = append(args, "--generator-provider", ".")
		if err := cli.Run(args, true); err != nil {
			panic(err)
		}
		os.Exit(0)
		return
	}

	// prisma CLI
	if err := cli.Run(args, true); err != nil {
		panic(err)
	}

	if len(args) > 2 {
		if err := invokePrisma(); err != nil {
			log.Printf("error occurred when invoking prisma: %s", err)
			os.Exit(1)
		}
	}

	logger.Debug.Printf("success")
}
