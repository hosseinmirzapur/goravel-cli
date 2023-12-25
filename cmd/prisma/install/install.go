/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package install

import (
	"os/exec"

	"github.com/hosseinmirzapur/goravel-cli/config"
	"github.com/hosseinmirzapur/goravel-cli/utils"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "installs prisma client for go",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Alert("Installing golang prisma client on machine...", false)
		installGoPrismaClient()
		utils.Success("prisma client installed", true)
		utils.Info("You can access the prisma client by running \"prisma-client-go <COMMMAND>\" ", false)
	},
}

func installGoPrismaClient() {
	utils.Info("looking for go prisma client...", false)
	command := exec.Command(
		"go",
		"install",
		config.GetPrismaConfig().GithubRepo,
	)

	if err := utils.HandleOutput(command); err != nil {
		utils.Error(
			"prisma",
			"unable to get golang's client for prisma",
			err,
		)
	}
}

func init() {

}
