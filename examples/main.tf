terraform {
  required_providers {
    pagerduty = {
      version = "1.0"
      source  = "tharun/edu/pagerduty"
    }
  }
}

provider "pagerduty" {
  token = "Token token=[token]"
}
//to fetch the user data
data "pagerduty_user_data" "name" {
  id = "PQXBTF7"
}
//to create user resource
resource "pagerduty_user_resource" "test1" {
  name  = "tharun"
  email = "teovjsofv@gmail.com"
  type  = "user"
  role  = "admin"
}

output "user" {
  value = data.pagerduty_user_data.name
}
output "resource_test" {
  value = pagerduty_user_resource.test1
}
