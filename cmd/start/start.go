package start

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/hosseinmirzapur/goravel-cli/config"
	"github.com/hosseinmirzapur/goravel-cli/utils"
	"github.com/spf13/cobra"
)

var (
	appName = ""
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start New Goravel App",
	Long:  `Create a fresh Goravel application`,
	Run: func(cmd *cobra.Command, args []string) {
		// set project name
		if appName == "." || appName == "" {
			appName = findCurrentWorkDir()
		}

		// clone goravel repo
		utils.Info("cloning...")
		err := cloneGoravelRepo(appName)
		utils.HandleError("start", "clone goravel repository", err)
		utils.Success("cloned successfully")

		// remove .git from cloned repository
		utils.Info("removing .git directory...")
		err = removeVersionControlDir(appName)
		utils.HandleError("start", "remove version control directory", err)
		utils.Success("removed successfully")

		// cd into project
		utils.Info("cd into project...")
		err = os.Chdir(appName)
		utils.HandleError("start", "cd into project", err)
		utils.Success("cd successfully")

		// run go mod tidy to install dependencies
		utils.Info("installing dependencies...")
		err = goModTidy()
		utils.HandleError("start", "run go mod tidy", err)
		utils.Success("installed successfully")

		// copy .env from .env.example
		utils.Info("copy .env from .env.example...")
		err = copyDotEnv()
		utils.HandleError("start", "copy .env from .env.example", err)
		utils.Success("copied successfully")

		// generate app key from artisan console
		utils.Info("generating app key...")
		err = generateAppKey()
		utils.HandleError("start", "generate app key", err)
		utils.Success("generated successfully")

		utils.Success("You project is ready! Create something amazing :)")
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
	command := exec.Command(
		"git",
		"clone",
		config.GetGoravelConfig().GithubRepo,
		fmt.Sprintf("./%s", appName),
	)
	out, err := command.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func removeVersionControlDir(appName string) error {
	command := exec.Command(
		"rm",
		"-rf",
		fmt.Sprintf("./%s/.git", appName),
	)
	out, err := command.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func goModTidy() error {
	command := exec.Command(
		"go",
		"mod",
		"tidy",
	)
	out, err := command.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func copyDotEnv() error {
	command := exec.Command(
		"cp",
		"./.env.example",
		"./.env",
	)
	out, err := command.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func generateAppKey() error {
	command := exec.Command(
		"go",
		"run",
		".",
		"artisan",
		"key:generate",
	)
	out, err := command.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil

}

// Start Point
func init() {
	StartCmd.Flags().StringVarP(&appName, "name", "n", "", "Name of the Project")
}
