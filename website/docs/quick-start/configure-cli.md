---
title: Configure CLI
sidebar_position: 4
sidebar_label: Configure CLI
---

In the previous step, we've decided on the following:

- Use a monorepo to configure and provision two Terraform components into three AWS accounts and two AWS regions
- The filesystem layout for the infrastructure monorepo
- To be able to use [Component Remote State](/core-concepts/components/remote-state), we put the `atmos.yaml` CLI config file
  into `/usr/local/etc/atmos/atmos.yaml` folder and set the ENV var `ATMOS_BASE_PATH` to point to the absolute path of the root of the repo

Next step is to configure `atmos.yaml`.

`atmos.yaml` configuration file is used to control the behavior of the `atmos` CLI. The file supports many features that are configured in different
sections of the `atmos.yaml` file. For the description of all the sections, refer to [CLI Configuration](/cli/configuration).

For the purpose of this Quick Start, below is the minimum configuration required for Atmos to work with Terraform and to
configure [Atmos components](/core-concepts/components) and [Atmos stacks](/core-concepts/stacks). Copy the YAML config below into your `atmos.yaml`
file.

<br/>

```yaml
# CLI config is loaded from the following locations (from lowest to highest priority):
# system dir ('/usr/local/etc/atmos' on Linux, '%LOCALAPPDATA%/atmos' on Windows)
# home dir (~/.atmos)
# current directory
# ENV vars
# Command-line arguments
#
# It supports POSIX-style Globs for file names/paths (double-star '**' is supported)
# https://en.wikipedia.org/wiki/Glob_(programming)

# Base path for components, stacks and workflows configurations.
# Can also be set using 'ATMOS_BASE_PATH' ENV var, or '--base-path' command-line argument.
# Supports both absolute and relative paths.
# If not provided or is an empty string, 'components.terraform.base_path', 'components.helmfile.base_path', 'stacks.base_path' 
# and 'workflows.base_path' are independent settings (supporting both absolute and relative paths).
# If 'base_path' is provided, 'components.terraform.base_path', 'components.helmfile.base_path', 'stacks.base_path' 
# and 'workflows.base_path' are considered paths relative to 'base_path'.
base_path: ""

components:
  terraform:
    # Can also be set using 'ATMOS_COMPONENTS_TERRAFORM_BASE_PATH' ENV var, or '--terraform-dir' command-line argument
    # Supports both absolute and relative paths
    base_path: "components/terraform"
    # Can also be set using 'ATMOS_COMPONENTS_TERRAFORM_APPLY_AUTO_APPROVE' ENV var
    apply_auto_approve: false
    # Can also be set using 'ATMOS_COMPONENTS_TERRAFORM_DEPLOY_RUN_INIT' ENV var, or '--deploy-run-init' command-line argument
    deploy_run_init: true
    # Can also be set using 'ATMOS_COMPONENTS_TERRAFORM_INIT_RUN_RECONFIGURE' ENV var, or '--init-run-reconfigure' command-line argument
    init_run_reconfigure: true
    # Can also be set using 'ATMOS_COMPONENTS_TERRAFORM_AUTO_GENERATE_BACKEND_FILE' ENV var, or '--auto-generate-backend-file' command-line argument
    auto_generate_backend_file: false

stacks:
  # Can also be set using 'ATMOS_STACKS_BASE_PATH' ENV var, or '--config-dir' and '--stacks-dir' command-line arguments
  # Supports both absolute and relative paths
  base_path: "stacks"
  # Can also be set using 'ATMOS_STACKS_INCLUDED_PATHS' ENV var (comma-separated values string)
  included_paths:
    - "orgs/**/*"
  # Can also be set using 'ATMOS_STACKS_EXCLUDED_PATHS' ENV var (comma-separated values string)
  excluded_paths:
    - "**/_defaults.yaml"
  # Can also be set using 'ATMOS_STACKS_NAME_PATTERN' ENV var
  name_pattern: "{tenant}-{environment}-{stage}"

workflows:
  # Can also be set using 'ATMOS_WORKFLOWS_BASE_PATH' ENV var, or '--workflows-dir' command-line arguments
  # Supports both absolute and relative paths
  base_path: "stacks/workflows"

logs:
  file: "/dev/stdout"
  # Supported log levels: Trace, Debug, Info, Warning, Off
  level: Info

# Custom CLI commands
commands: []

# Integrations
integrations: {}

# Validation schemas (for validating atmos stacks and components)
schemas: {}
```

<br/>

The `atmos.yaml` configuration file defines the following sections.

__NOTE:__ below is the description of the sections relevant to this Quick Start guide. For the description of all the sections, refer
to [CLI Configuration](/cli/configuration).

- `base_path` - the base path for components, stacks and workflows configurations. We set it to an empty string because we've decided to use the ENV
  var `ATMOS_BASE_PATH` to point to the absolute path of the root of the repo

- `components.terraform.base_path` - the base path to the Terraform components (Terraform root modules). As described in
  [Configure Repository](/quick-start/configure-repository), we've decided to put the Terraform components into the `components/terraform` directory,
  and this setting tells Atmos where to find them. Atmos will join the base path (set in the `ATMOS_BASE_PATH` ENN var)
  with `components.terraform.base_path` to calculate the final path to the Terraform components

- `components.terraform.apply_auto_approve` - if set to `true`, Atmos automatically adds the `-auto-approve` option to instruct Terraform to apply the
  plan without asking for confirmation when executing `terraform apply` command

- `components.terraform.deploy_run_init` - if set to `true`, Atmos runs `terraform init` before
  executing [`atmos terraform deploy`](/cli/commands/terraform/deploy) command

- `components.terraform.init_run_reconfigure` - if set to `true`, Atmos automatically adds the `-reconfigure` option to update the backend
  configuration when executing `terraform init` command

- `components.terraform.auto_generate_backend_file` - if set to `true`, Atmos automatically generates the Terraform backend file from the component
  configuration when executing `terraform plan` and `terraform apply` commands

- `stacks.base_path` - the base path to the Atmos stacks. As described in
  [Configure Repository](/quick-start/configure-repository), we've decided to put the stack configurations into the `stacks` directory,
  and this setting tells Atmos where to find them. Atmos will join the base path (set in the `ATMOS_BASE_PATH` ENN var)
  with `stacks.base_path` to calculate the final path to the stacks

- `stacks.included_paths` - list of file paths to the top-level stacks in the `stacks` directory to include in search when Atmos searches for the
  stack where the component is defined when executing `atmos` commands

- `stacks.excluded_paths` - list of file paths to the top-level stacks in the `stacks` directory to exclude from search when Atmos searches for the
  stack where the component is defined when executing `atmos` commands

- `stacks.name_pattern` - Atmos stack name pattern. When executing `atmos` commands, Atmos does not use the configuration file names and their
  filesystem locations to search for the stack where the component is defined. Instead, Atmos uses the context
  variables (`namespace`, `tenant`, `environment`, `stage`) to search for the stack. The stack config file names can be anything, and they can be in
  any folder in any sub-folder in the `stacks` directory. For example, when executing the `atmos terraform apply infra/vpc -s tenant1-ue2-dev`
  command, the stack `tenant1-ue2-dev` is specified by the `-s` flag. By looking at `name_pattern: "{tenant}-{environment}-{stage}"` and processing
  the tokens, Atmos knows that the first part of the stack name is `tenant`, the second part is `environment`, and the third part is `stage`. Then
  Atmos searches for the stack configuration file (in the `stacks` directory) where `tenant: tenant1`, `environment: ue2` and `stage: dev` are
  defined (inline or via imports)

- `workflows.base_path` - the base path to Atmos [workflow](/core-concepts/workflows) files

- `logs.verbose` - set to `true` to increase log verbosity. When set to `true`, Atmos prints to the console all the steps it takes to find and
  process the `atmos.yaml` CLI config file, and all the steps it takes to find the stack and find and process the component in the stack
