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
| account\_id | AWS Account ID that your lambda will be in | `string` | n/a | yes |
| command | The CMD for the docker image. | `list(string)` | <pre>[<br>  ""<br>]</pre> | no |
| description | Description of your Lambda Function (or Layer) | `string` | `""` | no |
| ecr\_arn | The ARN for the Elastic Contaier Repository. | `string` | `"*"` | no |
| entry\_point | The ENTRYPOINT for the docker image. | `list(string)` | <pre>[<br>  ""<br>]</pre> | no |
| environment | Environment name | `string` | n/a | yes |
| environment\_variables | A map that defines environment variables for the Lambda Function. | `map(string)` | `{}` | no |
| image\_uri | The URI for the coontainer image to use | `string` | `null` | no |
| lambda\_name | A unique name for your Lambda Function | `string` | n/a | yes |
| package\_type | The Lambda deployment package type. | `string` | `"Image"` | no |
| source\_secret\_region | Region to copy secrets from. | `string` | n/a | yes |
| tags | A map of tags to assign to resources. | `map(string)` | `{}` | no |
| target\_secret\_region | Region to copy secrets to. | `string` | n/a | yes |
| timeout | The amount of time your Lambda Function has to run in seconds. | `number` | `3` | no |
| working\_directory | The working directory for the docker image. | `string` | `null` | no |

## Outputs

No output.

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
