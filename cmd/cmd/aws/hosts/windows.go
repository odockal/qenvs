package hosts

import (
	params "github.com/adrianriobo/qenvs/cmd/cmd/constants"
	qenvsContext "github.com/adrianriobo/qenvs/pkg/manager/context"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/action/windows"
	"github.com/adrianriobo/qenvs/pkg/util/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	cmdWindows     = "windows"
	cmdWindowsDesc = "manage windows dedicated host"

	amiName            string = "ami-name"
	amiNameDesc        string = "name for the custom ami to be used within windows machine. Check README on how to build it"
	amiNameDefault     string = "Windows_Server-2019-English-Full-HyperV-RHQE"
	amiUsername        string = "ami-username"
	amiUsernameDesc    string = "name for de default user on the custom AMI"
	amiUsernameDefault string = "ec2-user"
	amiOwner           string = "ami-owner"
	amiOwnerDesc       string = "alias name for the owner of the custom AMI"
	amiOwnerDefault    string = "self"
	amiLang            string = "ami-lang"
	amiLangDesc        string = "language for the ami possible values (eng, non-eng). This param is used when no ami-name is set and the action uses the default custom ami"
	amiLangDefault     string = "eng"
)

func GetWindowsCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   cmdWindows,
		Short: cmdWindowsDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
	}
	c.AddCommand(getWindowsCreate(), getWindowsDestroy())
	return c
}

func getWindowsCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   params.CreateCmdName,
		Short: params.CreateCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			// Initialize context
			qenvsContext.Init(
				viper.GetString(params.ProjectName),
				viper.GetString(params.BackedURL),
				viper.GetString(params.ConnectionDetailsOutput),
				viper.GetStringMapString(params.Tags))

			// Run create
			if err := windows.Create(
				&windows.Request{
					Prefix:   "main",
					AMIName:  viper.GetString(amiName),
					AMIUser:  viper.GetString(amiUsername),
					AMIOwner: viper.GetString(amiOwner),
					AMILang:  viper.GetString(amiLang),
					Spot:     viper.IsSet(spot),
					Airgap:   viper.IsSet(airgap)}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	flagSet := pflag.NewFlagSet(params.CreateCmdName, pflag.ExitOnError)
	flagSet.StringP(params.ConnectionDetailsOutput, "", "", params.ConnectionDetailsOutputDesc)
	flagSet.StringToStringP(params.Tags, "", nil, params.TagsDesc)
	flagSet.StringP(amiName, "", amiNameDefault, amiNameDesc)
	flagSet.StringP(amiUsername, "", amiUsernameDefault, amiUsernameDesc)
	flagSet.StringP(amiOwner, "", amiOwnerDefault, amiOwnerDesc)
	flagSet.StringP(amiLang, "", amiLangDefault, amiLangDesc)
	flagSet.Bool(airgap, false, airgapDesc)
	flagSet.Bool(spot, false, spotDesc)
	c.PersistentFlags().AddFlagSet(flagSet)
	return c
}

func getWindowsDestroy() *cobra.Command {
	c := &cobra.Command{
		Use:   params.DestroyCmdName,
		Short: params.DestroyCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			qenvsContext.InitBase(
				viper.GetString(params.ProjectName),
				viper.GetString(params.BackedURL))

			if err := windows.Destroy(); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	return c
}
