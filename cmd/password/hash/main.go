package hash

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/usvc/password"
)

type Flags struct {
	KeyLength    uint32
	OutputFormat string
	Separator    string
}

type Output struct {
	Hash      string `json:"hash"`
	Salt      string `json:"salt"`
	Timestamp string `json:"timestamp"`
}

func GetCommand() *cobra.Command {
	flags := Flags{}
	cmd := &cobra.Command{
		Use: "hash",
		Run: func(command *cobra.Command, args []string) {
			plaintext := strings.Join(args, " ")
			if len(plaintext) == 0 {
				fmt.Println("no password string provided")
				return
			}
			hash, salt, err := password.Hash(plaintext, flags.KeyLength)
			if err != nil {
				panic(err)
			}
			switch flags.OutputFormat {
			case "json":
				output := Output{
					Hash:      hash,
					Salt:      salt,
					Timestamp: time.Now().Format(time.RFC3339),
				}
				json, err := json.Marshal(output)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(json))
			default:
				fmt.Printf("%s%s%s\n", hash, flags.Separator, salt)
			}
		},
	}
	cmd.Flags().StringVarP(&flags.OutputFormat, "output", "o", "text", "specifies the output format (either 'json' or 'text')")
	cmd.Flags().StringVarP(&flags.Separator, "separator", "s", ".", "specifies the separator between hash and salt")
	cmd.Flags().Uint32VarP(&flags.KeyLength, "key-length", "l", uint32(32), "specifies the intended length of the generated hash")
	return cmd
}
