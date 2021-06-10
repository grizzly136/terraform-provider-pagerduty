* This terraform provider allows to perform Create, Read, Update, Delete and Import of pagerduty Users. 
* To fetch and import a user server generated User ID is needed.
## Requirements 
* [Go](https://golang.org/doc/install) >= 1.16 (To build the provider plugin)<br>
* [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x <br/>
* Application:[Pagerduty Developer account](https://www.pagerduty.com/)
## Application setup 
### setup
* create a [pagerduty](https://www.pagerduty.com/) developer account
### Authentication
* API uses the basic token for authentication
* Login to your developer account.
* In the dashboard go to `Integrations -> Api acess keys -> click on  **create new api key**` <br>
## Building The Provider
1. Clone the repository "https://github.com/tharunforu/terraform-provider-pagerduty"<br>
2. Run the following commands :
 ```
go mod init terraform-provider-pagerduty
go mod tidy
go mod vendor
```
## Managing plugins for terraform
1. Run the following command to create a vendor subdirectory which will comprise of  all provider dependencies. <br>
```
~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}
``` 
Command: 
```bash
mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/pagerduty/1.0/[OS_ARCH]
```
For eg. `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/pagerduty/1.0/windows_amd64`<br>

2. Run `go build -o terraform-provider-pagerduty.exe`. This will save the Generate a binary file. <br>
3. Run this command to move this binary file to the desired location From were we can access the provider.
 ```
 move terraform-provider-pagerduty.exe %APPDATA%\terraform.d\plugins\hashicorp.com\edu\pagerduty\1.0\[OS_ARCH]
 ``` 
[OR]
1. Download required binaries <br>
2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Working with terraform
### application credential integration in terraform
* In the provider block add api token in the respective field
### Basic terraform commands
* `terraform init` initializes the plugins
* `terraform plan` suggests if any changes required
* `terraform apply` executes the code
### Create User
1. To create a user Add the user email, name, type and role in the respective field 
2. You will see that a user has been successfully created after executing basic terraform commands.

### Update the User
Update the data of the user file and apply using `terraform apply`

### Read the User Data
Add data and output blocks and run basic terraform commands

### Delete the user
Run `terraform destroy` to delete all the users or just comment out the unwated user resource block and run basic terraform commands
### Import User Data
1. Write manually a resource configuration block for the User, to which the imported object will be mapped or define the empty resource block.
2. Run the command `terraform pagerduty_resource_user.import “[USER_ID]”`
3. To validate Check for the attributes in the `.tfstate` file .
### naunces
* To get the user Id go to pagerduty `dashboard -> people -> users -> particular_user`  last 7 digits will be the user id
eg :  URL:……users/PDFJIJD  here PDFJIJD is the user id of the user
## Example Usage
```
terraform {
  required_providers {
    pagerduty = {
      version = "1.0.8"
      source  = "tharunforu/pagerduty"
    }
  }
}

provider "pagerduty" {
  token = "Token token=examepletoken"
}
data "pagerduty_user_data" "name" {
  id = "PDH5H9D"
}
output "user" {
  value = data.pagerduty_user_data.name
}
 resource "pagerduty_user_resource" "test1" {
   name  = "testname"
   email = "test@gmail.com"
   type  = "user"
   role  = "admin"
 }
 output "resource_test" {
   value = pagerduty_user_resource.test1
 }
```
## Argument Reference

* `email` -(Required)  The email id. we should give the email id of the user we want to create. Type `string`.
* `name` -(Required)  Name of the User.Type `string`.
* `type` -(Required)  Type of the User.Type `string`.
* `role` -(Required)  Job role of user in organization.Type `string`.It can take values `read_only_user`, `read_only_limited_user`, `owner`, `admin`, `user`, `observer`, `limited_user`, `team_responder`, or `restricted_access`. 
-