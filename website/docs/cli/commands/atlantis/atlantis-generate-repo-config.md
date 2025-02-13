---
title: atmos atlantis generate repo-config
sidebar_label: generate repo-config
sidebar_class_name: command
id: generate-repo-config
description: Use this command to generate a repository configuration for Atlantis.
---

:::info Purpose
Use this command to generate a repository configuration for Atlantis.
:::

<br/>

```shell
atmos atmos atlantis generate repo-config [options]
```

<br/>

:::tip
Run `atmos atlantis generate repo-config --help` to see all the available options
:::

## Examples

```shell
atmos atlantis generate repo-config

atmos atlantis generate repo-config --output-path /dev/stdout

atmos atlantis generate repo-config --config-template config-1 --project-template project-1

atmos atlantis generate repo-config --config-template config-1 --project-template project-1 --stacks <stack1, stack2>

atmos atlantis generate repo-config --config-template config-1 --project-template project-1 --components <component1, component2>

atmos atlantis generate repo-config --config-template config-1 --project-template project-1 --stacks <stack1> --components <component1, component2>

atmos atlantis generate repo-config --affected-only=true

atmos atlantis generate repo-config --affected-only=true --output-path /dev/stdout

atmos atlantis generate repo-config --affected-only=true --verbose=true

atmos atlantis generate repo-config --affected-only=true --output-path /dev/stdout --verbose=true

atmos atlantis generate repo-config --affected-only=true --repo-path <path_to_cloned_target_repo>

atmos atlantis generate repo-config --affected-only=true --ref refs/heads/main

atmos atlantis generate repo-config --affected-only=true --ref refs/tags/v1.1.0

atmos atlantis generate repo-config --affected-only=true --sha 3a5eafeab90426bd82bf5899896b28cc0bab3073

atmos atlantis generate repo-config --affected-only=true --ref refs/tags/v1.2.0 --sha 3a5eafeab90426bd82bf5899896b28cc0bab3073

atmos atlantis generate repo-config --affected-only=true --ssh-key <path_to_ssh_key>

atmos atlantis generate repo-config --affected-only=true --ssh-key <path_to_ssh_key> --ssh-key-password <password>
```

## Flags

| Flag                 | Description                                                                                                                                                      | Required |
|:---------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------|:---------|
| `--config-template`  | Atlantis config template name                                                                                                                                    | no       |
| `--project-template` | Atlantis project template name                                                                                                                                   | no       |
| `--output-path`      | Output path to write `atlantis.yaml` file                                                                                                                        | no       |
| `--stacks`           | Generate Atlantis projects for the specified stacks only (comma-separated values)                                                                                | no       |
| `--components`       | Generate Atlantis projects for the specified components only (comma-separated values)                                                                            | no       |
| `--affected-only`    | Generate Atlantis projects only for the Atmos components changed<br/>between two Git commits                                                                     | no       |
| `--ref`              | [Git Reference](https://git-scm.com/book/en/v2/Git-Internals-Git-References) with which to compare the current working branch                                    | no       |
| `--sha`              | Git commit SHA with which to compare the current working branch                                                                                                  | no       |
| `--ssh-key`          | Path to PEM-encoded private key to clone private repos using SSH                                                                                                 | no       |
| `--ssh-key-password` | Encryption password for the PEM-encoded private key if the key contains<br/>a password-encrypted PEM block                                                       | no       |
| `--repo-path`        | Path to the already cloned target repository with which to compare the current branch.<br/>Conflicts with `--ref`, `--sha`, `--ssh-key` and `--ssh-key-password` | no       |
| `--verbose`          | Print more detailed output when cloning and checking out the target<br/>Git repository and processing the result                                                 | no       |

<br/>

:::info

Refer to [Atlantis Integration](/integrations/atlantis.md) for more details on the Atlantis integration in Atmos

:::
