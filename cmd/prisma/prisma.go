/*
Copyright © 2023 Hossein Mirzapur: hosseinmirzapur@gmail.com
*/
package prisma

import (
	"github.com/hosseinmirzapur/goravel-cli/cmd/prisma/install"
	"github.com/spf13/cobra"
)

var PrismaCmd = &cobra.Command{
	Use:                "prisma",
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
	},
}

func init() {
	PrismaCmd.AddCommand(install.InstallCmd)
}
