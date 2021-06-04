terraform {
  required_providers {
    pagerduty = {
      version = "1.0"
      source  = "tharun/edu/pagerduty"
    }
  }
}

provider "pagerduty" {
  token = "Token token=[Token]"
}
data "pagerduty_user_data" "name" {
  id = "PQXBTF7"
}
output "user" {
  value = data.pagerduty_user_data.name
}
resource "pagerduty_user_resource" "test1" {
  name  = "tharun"
  email = "teovjsofv@gmail.com"
  type  = "user"
  role  = "admin"
}
output "resource_test" {
  value = pagerduty_user_resource.test1
}
# resource "pagerduty_user_resource" "test2" {
#   name  = "test121233"
#   email = "test12@gmail.com"
#   type  = "user"
#   role  = "observer"
# }
# output "resource_test2" {
#   value = pagerduty_user_resource.test2
# }
# resource "pagerduty_user_resource" "tharun"  {
#   name  = "test"
#   email = "test123@gmail.com"
#   type  = "user"
#   role  = "admin"
# }

