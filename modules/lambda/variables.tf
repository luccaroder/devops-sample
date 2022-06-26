variable account_environment {}
variable "tags" {
  type                = map
  default = {
    deployment_tool   = "Terraform"
    service = "lambda"
  }
}

variable "LAMBDA_VARIABLES" {
  type = "map"
}

variable "LAMBDA_SETTINGS" {
  type = "map"
}
