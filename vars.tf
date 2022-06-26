variable "account_environment" {
  type        = string
  description = "Setting Account environment"
}

variable "tags" {
  type                = map
  default = {
    deployment_tool   = "Terraform"
  }
}

variable "aws_region" {
  type = string
  default = "us-east-1"
}

variable "HELLO_LAMBDA" {
  type = "map"

  default = {
    settings = {
      function_name = "hello_lambda"
      handler       = "hello_lambda.lambda_handler"
      runtime       = "python3.6"
      source_file   = "hello_lambda.py"
    }

    variables = {
      hello = "Hello"
    }
  }
}
