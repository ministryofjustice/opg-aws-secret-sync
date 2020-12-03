variable "description" {
  description = "Description of your Lambda Function (or Layer)"
  type        = string
  default     = ""
}

variable "environment" {
  description = "Environment name"
  type        = string
}

variable "environment_variables" {
  description = "A map that defines environment variables for the Lambda Function."
  type        = map(string)
  default     = {}
}

variable "handler" {
  description = "Lambda Function entrypoint in your code"
  type        = string
}

variable "image_arn" {
  description = "The ARN for the coontainer image to use"
  type        = string
  default     = null
}

variable "lambda_name" {
  description = "A unique name for your Lambda Function"
  type        = string
}

variable "runtime" {
  description = "Lambda Function runtime"
  type        = string
  default     = "go1.x"
}

variable "timeout" {
  description = "The amount of time your Lambda Function has to run in seconds."
  type        = number
  default     = 3
}


variable "tags" {
  description = "A map of tags to assign to resources."
  type        = map(string)
  default     = {}
}

variable "source_secret_region" {
  description = "Region to copy secrets from."
  type        = string
}

variable "target_secret_region" {
  description = "Region to copy secrets to."
  type        = string
}
