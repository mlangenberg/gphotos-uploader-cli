package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/gphotosuploader/gphotos-uploader-cli/internal/app"
	"github.com/gphotosuploader/gphotos-uploader-cli/internal/cmd/flags"
)

// AuthCmd holds the required data for the init cmd
type AuthCmd struct {
	*flags.GlobalFlags
}

func NewAuthCmd(globalFlags *flags.GlobalFlags) *cobra.Command {
	cmd := &AuthCmd{GlobalFlags: globalFlags}

	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate with Google Photos to refresh tokens",
		Long:  `Force authentication against Google Photos to refresh tokens.`,
		Args:  cobra.NoArgs,
		RunE:  cmd.Run,
	}

	return authCmd
}

func (cmd *AuthCmd) Run(cobraCmd *cobra.Command, args []string) error {
	ctx := context.Background()
	cli, err := app.Start(ctx, cmd.CfgDir)
	if err != nil {
		return err
	}
	defer func() {
		_ = cli.Stop()
	}()

	cli.Logger.Donef("Successful authentication for account '%s'", cli.Config.Account)
	return nil
}
