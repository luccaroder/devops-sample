# Terraform

## Requirements

- Terraform = 0.14.4

## Directories

| Directory | Description |
| --------- | ----------- |
| modules | Contains the internal modules, created per our time. |
| stg | Contains the all globals envs use in STG environment |
| prd | Contains the all globals envs use in PRD environment |

## Important files

| File | Description |
| --------- | ----------- |
| backend.tf | Backends are configured directly in Terraform files in the terraform section. |
| vars.tf | Define the global variable values |

## How run?

### Terraform init

```sh
terraform init -backend-config=stg/variables.tfvars
```

### Terraform plan
Run plan in all targets:
```sh
terraform plan -var-file=stg/variables.tfvars
```

Run plan in one target:
```sh
terraform plan -var-file=tsg/variables.tfvars -target=module.modulename
```

### Terraform apply

```sh
terraform apply -var-file=stg/variables.tfvars
```

Run apply in one target:
```sh
terraform apply -var-file=tsg/variables.tfvars -target=module.modulename
```

## How can I test it?

Uses terratest and the tests files uses golang language.

1 - Install and configure go
2 - follow instructions: [Tutorial](https://terratest.gruntwork.io/docs/getting-started/quick-start/)
3 - Access folder `/test` and run `go test -v -timeout 30m`

### Terraform validate

```sh
terraform validate
```
