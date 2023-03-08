package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"onChainLayer/x/onchainlayer/types"
)

func CmdCreateRegistrand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-registrand [ip] [nickname] [sig] [pk] [msg]",
		Short: "Create a new registrand",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIp := args[0]
			argNickname := args[1]
			argSig := args[2]
			argPk := args[3]
			argMsg := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRegistrand(clientCtx.GetFromAddress().String(), argIp, argNickname, argSig, argPk, argMsg)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateRegistrand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-registrand [id] [ip] [nickname] [sig] [pk] [msg]",
		Short: "Update a registrand",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argIp := args[1]

			argNickname := args[2]

			argSig := args[3]

			argPk := args[4]

			argMsg := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRegistrand(clientCtx.GetFromAddress().String(), id, argIp, argNickname, argSig, argPk, argMsg)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteRegistrand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-registrand [id]",
		Short: "Delete a registrand by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRegistrand(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
