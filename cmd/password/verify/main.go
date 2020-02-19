package verify

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/usvc/password"
)

type Flags struct {
	KeyLength    uint32
	OutputFormat string
	Separator    string
}

type Output struct {
	OK bool `json:"ok"`
}

func GetCommand() *cobra.Command {
	flags := Flags{}
	cmd := &cobra.Command{
		Use: "verify",
		Run: func(command *cobra.Command, args []string) {
			plaintext := args[0]
			if len(plaintext) == 0 {
				fmt.Println("no password string provided")
				os.Exit(1)
				return
			}
			encodedHash := args[1]
			if len(encodedHash) == 0 {
				fmt.Println("no hash provided")
				os.Exit(2)
				return
			}
			encodedSalt := args[2]
			if len(encodedSalt) == 0 {
				fmt.Println("no salt provided")
				os.Exit(3)
				return
			}
			verificationFailed := password.Verify(plaintext, encodedHash, encodedSalt)
			switch flags.OutputFormat {
			case "json":
				output := Output{
					OK: verificationFailed != nil,
				}
				json, err := json.Marshal(output)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(json))
			default:
				if verificationFailed != nil {
					fmt.Println("not ok")
				} else {
					fmt.Println("ok")
				}
			}
			if verificationFailed != nil {
				os.Exit(255)
				return
			}
			os.Exit(0)
		},
	}
	cmd.Flags().StringVarP(&flags.OutputFormat, "output", "o", "text", "specifies the output format (either 'json' or 'text')")
	return cmd
}
