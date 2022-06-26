module "lambda" {
  source                        = "./modules/lambda"
  LAMBDA_SETTINGS               = "${var.HELLO_LAMBDA["settings"]}"
  LAMBDA_VARIABLES              = "${var.HELLO_LAMBDA["variables"]}"
  account_environment           = var.account_environment
}
