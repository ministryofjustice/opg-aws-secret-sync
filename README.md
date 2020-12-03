# opg-aws-secret-sync
Lambda to sync secrets across regions: Managed by opg-org-infra &amp; Terraform

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.13.5 |
| aws | >= 3.19.0 |

## Providers

| Name | Version |
|------|---------|
| aws | >= 3.19.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| description | Description of your Lambda Function (or Layer) | `string` | `""` | no |
| environment | Environment name | `string` | n/a | yes |
| environment\_variables | A map that defines environment variables for the Lambda Function. | `map(string)` | `{}` | no |
| handler | Lambda Function entrypoint in your code | `string` | n/a | yes |
| image\_arn | The ARN for the coontainer image to use | `string` | `null` | no |
| lambda\_name | A unique name for your Lambda Function | `string` | n/a | yes |
| runtime | Lambda Function runtime | `string` | `"go1.x"` | no |
| source\_secret\_region | Region to copy secrets from. | `string` | n/a | yes |
| tags | A map of tags to assign to resources. | `map(string)` | `{}` | no |
| target\_secret\_region | Region to copy secrets to. | `string` | n/a | yes |
| timeout | The amount of time your Lambda Function has to run in seconds. | `number` | `3` | no |

## Outputs

No output.

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
