package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsLambda(t *testing.T) {
	t.Parallel()
	helloLambdaFolder := test_structure.CopyTerraformFolderToTemp(t, "../", "/modules/lambda")
	functionName := fmt.Sprintf("aws-lambda-%s", random.UniqueId())

	awsRegion := "us-east-1"

	lambdaSettings := map[string]string{
		"function_name": functionName,
		"handler": "hello_lambda.lambda_handler",
		"runtime": "python3.6",
		"source_file": "hello_lambda.py",
	}

	lambdaVariables := map[string]string{
		"hello": "Hello",
	}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: helloLambdaFolder,

		Vars: map[string]interface{}{
			"account_environment": "test",
			"LAMBDA_SETTINGS": lambdaSettings,
			"LAMBDA_VARIABLES": lambdaVariables,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	response := aws.InvokeFunction(t, awsRegion, functionName + "_test", "")
	assert.Equal(t, `"Hello from Lambda!"`, string(response))
}

type FunctionPayload struct {
	Echo       string
	ShouldFail bool
}