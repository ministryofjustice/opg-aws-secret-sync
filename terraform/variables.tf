variable "aws_subnet_ids" {
  type = list(string)
}

variable "lambda_name" {
  type = string
}

variable "handler" {
  type = string
}

variable "lambda_function_subdir" {
  type = string
}

variable "tags" {}

variable "memory" {}

variable "environment" {}
