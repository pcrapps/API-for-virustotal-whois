terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

provider "aws" {
   region = "us-east-2"
}

resource "aws_lambda_function" "ipinfo" {
  filename      = "./lambda/lambdaapi.zip"
   function_name = "ipinfo"
   handler = "lambdaapi"
   runtime = "go1.x"
   role = aws_iam_role.lambda_exec.arn
}
resource "aws_iam_role" "lambda_exec" {
   name = "ipinfo_lambda"

   assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

}
resource "aws_lambda_permission" "apigw" {
   statement_id  = "AllowAPIGatewayInvoke"
   action        = "lambda:InvokeFunction"
   function_name = aws_lambda_function.ipinfo.function_name
   principal     = "apigateway.amazonaws.com"

   # The "/*/*" portion grants access from any method on any resource
   # within the API Gateway REST API.
   source_arn = "${aws_api_gateway_rest_api.ipinfo.execution_arn}/*/*"
}


