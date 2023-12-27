/*
Copyright © 2023 Hossein Mirzapur: hosseinmirzapur@gmail.com
*/
package artisan

import (
	"fmt"

	"github.com/hosseinmirzapur/goravel-cli/utils"
	"github.com/spf13/cobra"
)

var ArtisanCmd = &cobra.Command{
	Use:     "artisan [command]...",
	Short:   "A wrapper around \"go run . artisan\" command",
	Example: "goravel-cli artisan list",
	Run:     runArtisanCommand,
}

func runArtisanCommand(cmd *cobra.Command, args []string) {
	err := utils.GoRunDot(append([]string{"artisan"}, args...))
	utils.Error(
		"artisan",
		fmt.Sprintf("unable to run \"%+v\"", args),
		err,
	)
	fmt.Printf(`
ADDITIONAL COMMANDS:
   env:
     env:encrypt  %s
     env:decrypt  %s
	 
`,
		encryptEnvCmd.Short,
		decryptEnvCmd.Short,
	)
}

func init() {
	ArtisanCmd.AddGroup(envGroup)
	ArtisanCmd.AddCommand(encryptEnvCmd, decryptEnvCmd)
}
