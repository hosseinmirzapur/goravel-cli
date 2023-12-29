package cmd

import (
	"os"

	"github.com/hosseinmirzapur/goravel-cli/utils"
	"github.com/spf13/cobra/doc"
)

func generateDocs() {
	err := doc.GenMarkdownTree(rootCmd, "./")
	if err != nil {
		utils.Error("root", "Failed to generate markdown docs", err)
		os.Exit(1)
	}
}
