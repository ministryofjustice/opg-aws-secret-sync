resource "aws_lambda_function" "lambda_function" {
  filename         = var.filename
  function_name    = var.lambda_name
  handler          = var.handler
  layers           = [aws_lambda_layer_version.lambda_layer.arn]
  memory_size      = var.memory
  role             = aws_iam_role.lambda_role.arn
  runtime          = ""
  source_code_hash = data.archive_file.lambda_archive.output_base64sha256
  timeout          = 30

  environment {
    variables = {
      SOURCE_REGION = var.source_secret_region
      TARGET_REGION = var.target_secret_region
    }
  }

  vpc_config {
    subnet_ids         = var.aws_subnet_ids
    security_group_ids = [data.aws_security_group.lambda_api_ingress.id]
  }

  depends_on = [aws_cloudwatch_log_group.lambda]
  tags       = var.tags
}

resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "AllowCloudWatchEventInvoke_${var.environment}"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_function.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.event.arn
}

resource "aws_cloudwatch_log_group" "lambda" {
  name = "/aws/lambda/${var.lambda_name}"
  tags = var.tags
}
