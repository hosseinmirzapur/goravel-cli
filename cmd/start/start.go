package start

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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
			appName = findProjectName()
		}

		// call the shell script
		command := exec.Command("sh", "goravel-init", appName)

		out, err := command.Output()
		if err != nil {
			log.Println(strings.Repeat("*", 10), "START", strings.Repeat("*", 10))
			log.Println("err: ", err)
		}

		fmt.Println(string(out))
	},
}

func init() {
	StartCmd.Flags().StringVarP(&appName, "name", "n", "", "Name of the Project")
}

func findProjectName() string {
	currDir, err := os.Getwd()
	if err != nil {
		log.Println("START:", "unable to find current workdir")
		os.Exit(1)
	}
	results := strings.SplitAfter(currDir, "/")
	return results[len(results)-1]
}
