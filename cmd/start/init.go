package start

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hosseinmirzapur/goravel-cli/config"
	"github.com/hosseinmirzapur/goravel-cli/utils"
	"github.com/spf13/cobra"
)

var (
	appName = ""
)

var StartCmd = &cobra.Command{
	Use:     "init",
	Short:   "Start New Goravel App",
	Long:    `Create a fresh Goravel application`,
	Example: "goravel-cli start -n <PROJECT_NAME>",
	Run: func(cmd *cobra.Command, args []string) {
		// set project name
		if appName == "." || appName == "" {
			if len(args) > 0 {
				appName = args[0]
			} else {
				appName = findCurrentWorkDir()
			}
		}

		// clone goravel repo
		utils.Info("cloning...", false)
		err := cloneGoravelRepo(appName)
		utils.Error("start", "clone goravel repository", err)
		utils.Success("cloned successfully", true)

		// remove .git from cloned repository
		utils.Info("removing .git directory...", false)
		err = removeVersionControlDir(appName)
		utils.Error("start", "remove version control directory", err)
		utils.Success("removed successfully", true)

		// cd into project
		utils.Info("cd into project...", false)
		err = os.Chdir(appName)
		utils.Error("start", "cd into project", err)
		utils.Success("cd successfully", true)

		// run go mod tidy to install dependencies
		utils.Info("installing dependencies...", false)
		err = goModTidy()
		utils.Error("start", "run go mod tidy", err)
		utils.Success("installed successfully", true)

		// copy .env from .env.example
		utils.Info("copy .env from .env.example...", false)
		err = copyDotEnv()
		utils.Error("start", "copy .env from .env.example", err)
		utils.Success("copied successfully", true)

		// generate app key from artisan console
		utils.Info("generating app key...", false)
		err = generateAppKey()
		utils.Error("start", "generate app key", err)
		utils.Success("generated successfully", true)

		utils.Alert("You're set! Create something amazing :)", false)
	},
}

func findCurrentWorkDir() string {
	currDir, err := os.Getwd()
	if err != nil {
		log.Println("START:", "unable to find current work dir")
		os.Exit(1)
	}
	results := strings.SplitAfter(currDir, "/")
	return results[len(results)-1]
}

func cloneGoravelRepo(appName string) error {
	return utils.RunCommand([]string{
		"git",
		"clone",
		config.GetGoravelConfig().GithubRepo,
		fmt.Sprintf("./%s", appName),
	})
}

func removeVersionControlDir(appName string) error {
	return os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
}

func goModTidy() error {
	return utils.RunCommand([]string{
		"go",
		"mod",
		"tidy",
	})
}

func copyDotEnv() error {
	return utils.Copy(".env.example", ".env")
}

func generateAppKey() error {
	return utils.GoRunDot([]string{
		"artisan",
		"key:generate",
	})
}

// Start Point
func init() {
	StartCmd.Flags().StringVarP(&appName, "name", "n", "", "Name of the Project")
}
