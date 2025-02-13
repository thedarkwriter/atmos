---
title: atmos terraform generate varfile
sidebar_label: generate varfile
sidebar_class_name: command
id: generate-varfile
---

:::note purpose
Use this command to generate a varfile (`.tfvar` ) for an Atmos terraform [component](/core-concepts/components) in a [stack](/core-concepts/stacks).
:::

## Usage

Execute the `terraform generate varfile` command like this:

```shell
atmos terraform generate varfile <command> <component> -s <stack>
```

This command generates a varfile for an Atmos terraform component in a stack.

:::tip
Run `atmos terraform generate varfile --help` to see all the available options
:::

## Examples

```shell
atmos terraform generate varfile top-level-component1 -s tenant1-ue2-dev
atmos terraform generate varfile infra/vpc -s tenant1-ue2-staging
atmos terraform generate varfile test/test-component -s tenant1-ue2-dev
atmos terraform generate varfile test/test-component-override-2 -s tenant2-ue2-prod
atmos terraform generate varfile test/test-component-override-3 -s tenant1-ue2-dev -f vars.json
```

## Arguments

| Argument    | Description               | Required |
|:------------|:--------------------------|:---------|
| `component` | Atmos terraform component | yes      |

## Flags

| Flag        | Description | Alias | Required |
|:------------|:------------|:------|:---------|
| `--stack`   | Atmos stack | `-s`  | yes      |
| `--dry-run` | Dry run     |       | no       |
