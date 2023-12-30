/*
Copyright © 2023 Hossein Mirzapur: hosseinmirzapur@gmail.com
*/
package prisma

import (
	"github.com/hosseinmirzapur/goravel-cli/prisma/renderer"
	"github.com/spf13/cobra"
)

var PrismaCmd = &cobra.Command{
	Use:                "prisma [command]",
	Short:              "Prisma Client in Go",
	DisableFlagParsing: true,
	Example: `
goravel-cli prisma generate

# or

goravel-cli prisma migrate dev
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		renderer.Render(args)
	},
}

func init() {
}
