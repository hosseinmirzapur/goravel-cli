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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		runArtisanCommand(args)
	},
}

func runArtisanCommand(args []string) {
	err := utils.GoRunDot(append([]string{"artisan"}, args...))
	utils.Error(
		"artisan",
		fmt.Sprintf("unable to run \"%+v\"", args),
		err,
	)
}

func init() {

}
