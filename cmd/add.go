package cmd

import (
	"fmt"
	"os"

	"github.com/dtchanpura/deployment-agent/config"
	"github.com/dtchanpura/deployment-agent/constants"
	"github.com/dtchanpura/deployment-agent/manage"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	name      string
	maxargs   int
	hooks     []string
	preHook   string
	postHook  string
	errorHook string
	workDir   string
	cidr      []string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "For adding a new project",
	Long:  `Projects can be added with this command including the hooks, its IPs etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.ParseFlags(args)
		// var newTokens []config.TokenDetail
		// // project.Tokens = []Token{}
		project := config.NewProject(cidr...)
		project.Name = name
		project.MaxArgs = maxargs
		project.Hooks = config.NewHooks(hooks...)
		project.PreHook = preHook
		project.PostHook = postHook
		project.ErrorHook = errorHook
		project.WorkDir = workDir
		// bts, _ := yaml.Marshal(project)
		if err := project.ValidateProjectConfiguration(); err.Error() == constants.ErrorInvalidConfiguration {
			fmt.Println(err)
			os.Exit(1)
		}
		err := manage.AddProject(cfgFile, project)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// fmt.Println(string(bts[:]))
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}
	addCmd.Flags().StringVar(&name, "name", "", "Name of project.")
	addCmd.Flags().IntVar(&maxargs, "max-args", 0, "Maximum arguments limit for each of the hooks in the project.")
	addCmd.Flags().StringArrayVar(&hooks, "hook", []string{}, "Path to script to be executed on webhook call.")
	addCmd.Flags().StringVar(&preHook, "pre-hook", "", "Path to script to be executed before the event.")
	addCmd.Flags().StringVar(&postHook, "post-hook", "", "Path to script to be executed after the event.")
	addCmd.Flags().StringVar(&errorHook, "error-hook", "", "Path to script to be executed in case of error.")
	addCmd.Flags().StringVar(&workDir, "work-dir", home, "Work directory.")
	addCmd.Flags().StringArrayVar(&cidr, "ip-cidr", []string{"0.0.0.0/0"}, "Whitelist network CIDR which can access the webhook.")
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
