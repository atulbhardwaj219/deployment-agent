package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/dtchanpura/deployment-agent/config"
	"github.com/dtchanpura/deployment-agent/constants"
	"github.com/spf13/cobra"
)

// regenerateCmd represents the regenerate command
var regenerateCmd = &cobra.Command{
	Use:   "regenerate [UUID]",
	Short: "Regenerate Token",
	Args:  cobra.ExactArgs(1),
	Long:  `Regenerate token based on changes in name.`,
	Run: func(cmd *cobra.Command, args []string) {
		project, err := regenerate(args[0])
		if err != nil {
			fmt.Println("Error while regenerating: ", err)
			os.Exit(1)
		}
		for i := range project.Tokens {
			fmt.Println(project.Tokens[i].WhitelistedNetwork, ": ", project.GetHash(i))
		}
	},
}

func init() {
	RootCmd.AddCommand(regenerateCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// regenerateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// regenerateCmd.Flags().StringVar(p, name, value, usage)("toggle", "t", false, "Help message for toggle")
}

func regenerate(uuid string) (*config.Project, error) {
	if uuid != "" {
		project, err := config.FindProjectWithUUID(uuid)
		if err != nil {
			// fmt.Println(err)
			fmt.Printf("Project with UUID %s not found.\n", uuid)
			return nil, err
		}
		return &project, nil
	}
	return nil, errors.New(constants.ErrorInvalidUUID)
}
