package exec

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	cfg "github.com/cloudposse/atmos/pkg/config"
	"github.com/cloudposse/atmos/pkg/schema"
)

// ExecuteDescribeComponentCmd executes `describe component` command
func ExecuteDescribeComponentCmd(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("invalid arguments. The command requires one argument `component`")
	}

	_, err := cfg.InitCliConfig(schema.ConfigAndStacksInfo{}, true)
	if err != nil {
		return err
	}

	flags := cmd.Flags()

	stack, err := flags.GetString("stack")
	if err != nil {
		return err
	}

	format, err := flags.GetString("format")
	if err != nil {
		return err
	}

	file, err := flags.GetString("file")
	if err != nil {
		return err
	}

	component := args[0]

	componentSection, err := ExecuteDescribeComponent(component, stack)
	if err != nil {
		return err
	}

	err = printOrWriteToFile(format, file, componentSection)
	if err != nil {
		return err
	}

	return nil
}

// ExecuteDescribeComponent describes component config
func ExecuteDescribeComponent(component string, stack string) (map[string]any, error) {
	var configAndStacksInfo schema.ConfigAndStacksInfo
	configAndStacksInfo.ComponentFromArg = component
	configAndStacksInfo.Stack = stack

	cliConfig, err := cfg.InitCliConfig(configAndStacksInfo, true)
	if err != nil {
		return nil, err
	}

	configAndStacksInfo.ComponentType = "terraform"
	configAndStacksInfo, err = ProcessStacks(cliConfig, configAndStacksInfo, true)
	if err != nil {
		configAndStacksInfo.ComponentType = "helmfile"
		configAndStacksInfo, err = ProcessStacks(cliConfig, configAndStacksInfo, true)
		if err != nil {
			return nil, err
		}
	}

	// Add Atmos component and stack
	configAndStacksInfo.ComponentSection["atmos_component"] = configAndStacksInfo.ComponentFromArg
	configAndStacksInfo.ComponentSection["atmos_stack"] = configAndStacksInfo.StackFromArg

	// If the command-line component does not inherit anything, then the Terraform/Helmfile component is the same as the provided one
	if comp, ok := configAndStacksInfo.ComponentSection["component"].(string); !ok || comp == "" {
		configAndStacksInfo.ComponentSection["component"] = configAndStacksInfo.ComponentFromArg
	}

	// Spacelift stack
	spaceliftStackName, err := BuildSpaceliftStackNameFromComponentConfig(
		cliConfig,
		component,
		stack,
		configAndStacksInfo.ComponentSettingsSection,
		configAndStacksInfo.ComponentVarsSection,
	)

	if err != nil {
		return nil, err
	}

	if spaceliftStackName != "" {
		configAndStacksInfo.ComponentSection["spacelift_stack"] = spaceliftStackName
	}

	// Atlantis project
	atlantisProjectName, err := BuildAtlantisProjectNameFromComponentConfig(
		cliConfig,
		component,
		configAndStacksInfo.ComponentSettingsSection,
		configAndStacksInfo.ComponentVarsSection,
	)

	if err != nil {
		return nil, err
	}

	if atlantisProjectName != "" {
		configAndStacksInfo.ComponentSection["atlantis_project"] = atlantisProjectName
	}

	return configAndStacksInfo.ComponentSection, nil
}
