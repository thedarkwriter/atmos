package cmd

import (
	e "github.com/cloudposse/atmos/internal/exec"
	u "github.com/cloudposse/atmos/pkg/utils"
	"github.com/spf13/cobra"
)

// workflowCmd executes a workflow
var workflowCmd = &cobra.Command{
	Use:                "workflow",
	Short:              "Execute a workflow",
	Long:               `This command executes a workflow: atmos workflow <name> -f <file>`,
	FParseErrWhitelist: struct{ UnknownFlags bool }{UnknownFlags: false},
	Run: func(cmd *cobra.Command, args []string) {
		err := e.ExecuteWorkflowCmd(cmd, args)
		if err != nil {
			u.LogErrorAndExit(err)
		}
	},
}

func init() {
	workflowCmd.DisableFlagParsing = false
	workflowCmd.PersistentFlags().StringP("file", "f", "", "atmos workflow <name> -f <file>")
	workflowCmd.PersistentFlags().Bool("dry-run", false, "atmos workflow <name> -f <file> --dry-run")
	workflowCmd.PersistentFlags().StringP("stack", "s", "", "atmos workflow <name> -f <file> -s <stack>")
	workflowCmd.PersistentFlags().String("from-step", "", "atmos workflow <name> -f <file> --from-step <step-name>")

	err := workflowCmd.MarkPersistentFlagRequired("file")
	if err != nil {
		u.LogErrorAndExit(err)
	}

	RootCmd.AddCommand(workflowCmd)
}
