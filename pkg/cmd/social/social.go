package social

import (
	"github.com/bharat-rajani/GoHunter/pkg/cmdutil"
	"github.com/bharat-rajani/GoHunter/pkg/sherlock"
	"github.com/spf13/cobra"
)

func NewSocialCmd(cmdContainer *cmdutil.CMDContainerFactory) *cobra.Command {
	cmd := &cobra.Command{Use: "social [<username>... | -]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return socialRun(cmdContainer,args)
		},
	}
	return cmd
}

func socialRun(cmdContainer *cmdutil.CMDContainerFactory,usernames []string) error{

	// creating a new httpClient from factory
	client,err := cmdContainer.HttpClient()
	if err!=nil{
		return  err
	}
	sherlockObj := sherlock.NewSherLock(usernames,client)
	sherlockObj.Run()
	return nil
}