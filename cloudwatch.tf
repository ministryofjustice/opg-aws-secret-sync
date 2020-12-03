resource "aws_cloudwatch_event_rule" "event" {
  name          = "capture-secret-update-${var.environment}"
  description   = "Capture when a secret is updated in Secrets Manager"
  event_pattern = local.event_pattern
}

resource "aws_cloudwatch_event_target" "event" {
  rule      = aws_cloudwatch_event_rule.event.name
  target_id = var.lambda_name
  arn       = aws_lambda_function.lambda_function.arn
}

locals {

  event_pattern = jsonencode({
    detail-type = [
      "AWS Service Event via CloudTrail"
    ],
    source = [
      "aws.secretsmanager"
    ],
    detail = {
      eventSource = [
        "secretsmanager.amazonaws.com"
      ],
      eventName = [
        "RotationSucceeded"
      ]
    }
  })
}
