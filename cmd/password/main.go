package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/usvc/go-password/cmd/password/hash"
	"github.com/usvc/go-password/cmd/password/verify"
)

var (
	Commit    string
	Version   string
	Timestamp string
)

func main() {
	cmd := cobra.Command{
		Use:     "password",
		Version: fmt.Sprintf("%s-%s %s", Version, Commit, Timestamp),
		Run: func(command *cobra.Command, args []string) {
			command.Help()
		},
	}
	cmd.AddCommand(hash.GetCommand())
	cmd.AddCommand(verify.GetCommand())
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
