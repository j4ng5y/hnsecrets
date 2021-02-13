package cli

import (
	"log"

	"github.com/spf13/cobra"
)

// Run simply runs the CLI
func Run() {
	var (
		rootCMD = &cobra.Command{
			Use: "gateway",
			Run: func(ccmd *cobra.Command, args []string) {},
		}
	)

	if err := rootCMD.Execute(); err != nil {
		log.Fatal(err)
	}
}
