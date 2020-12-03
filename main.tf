resource "aws_lambda_function" "lambda_function" {
  function_name = var.lambda_name
  handler       = var.handler
  image_uri     = var.image_arn
  role          = aws_iam_role.lambda_role.arn
  runtime       = var.runtime
  timeout       = var.timeout


  dynamic "environment" {
    for_each = length(keys(var.environment_variables)) == 0 ? [] : [true]
    content {
      variables = var.environment_variables
    }
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
