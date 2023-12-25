/*
Copyright © 2023 Hossein Mirzapur: hosseinmirzapur@gmail.com
*/
package prisma

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/hosseinmirzapur/goravel-cli/config"
	"github.com/hosseinmirzapur/goravel-cli/utils"
	"github.com/spf13/cobra"
	"github.com/steebchen/prisma-client-go/binaries"
)

var PrismaCmd = &cobra.Command{
	Use:                "prisma",
	Short:              "Your lovely prisma client but in Go!",
	DisableFlagParsing: true,
	Example: `
goravel-cli prisma generate

# or even:

goravel-cli prisma studio
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		runPrismaWithArgs(args[:])
	},
}

func runPrismaWithArgs(args []string) {
	fmt.Println("Running prisma version:", binaries.PrismaVersion)

	plainText := strings.Join(args, " ")
	command := exec.Command(
		"go",
		"run",
		config.GetPrismaConfig().GithubRepo,
		plainText,
	)

	if err := utils.HandleOutput(command); err != nil {
		utils.Error(
			"prisma",
			fmt.Sprintf("unable to run \"%s\"", plainText),
			err,
		)
	}
}

func init() {

}
