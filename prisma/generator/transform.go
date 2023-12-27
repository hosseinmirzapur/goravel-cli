package generator

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hosseinmirzapur/goravel-cli/prisma/generator/ast/transform"
)

// Transform builds the AST from the flat DMMF so it can be used properly in templates
func Transform(input *Root) {
	input.AST = transform.New(&input.DMMF)
	if os.Getenv("DEBUG") != "" {
		d, _ := json.MarshalIndent(input.AST, "", "  ")
		fmt.Printf("AST: %s\n", string(d))
	}
}
