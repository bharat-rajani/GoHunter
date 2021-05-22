package root

import (
	socialCmd "github.com/bharat-rajani/GoHunter/pkg/cmd/social"
	cmdutil2 "github.com/bharat-rajani/GoHunter/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRoot(cmdContainerFactory *cmdutil2.CMDContainerFactory) *cobra.Command {

	cmd := &cobra.Command{
		Use:     "gohunter <command> <subcommand> [flags]",
		Short:   "GoHunter CLI",
		Long:    `Hunt down them mofos.`,
		Example: `gohunter social username1 username2 ....`,
		Annotations: map[string]string{
			"help:feedback": `
				Open an issue on '
			`,
			"help:environment": `
				See 'gh help environment' for the list of supported environment variables.
			`,
		},
	}

	cmd.SetOut(cmdContainerFactory.IOStreams.Out)
	cmd.SetErr(cmdContainerFactory.IOStreams.ErrOut)

	cs := cmdContainerFactory.IOStreams.ColorScheme()

	helpHelper := func(command *cobra.Command, args []string) {
		rootHelpFunc(cs, command, args)
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")
	cmd.SetHelpFunc(helpHelper)
	cmd.SetUsageFunc(rootUsageFunc)
	//cmd.SetFlagErrorFunc(rootFlagErrorFunc)

	cmd.AddCommand(socialCmd.NewSocialCmd(cmdContainerFactory))

	return cmd

}
