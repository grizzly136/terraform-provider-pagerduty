## Pagerduty Terraform Provider

* This terraform provider allows to perform Create, Read, Update, Delete and Import of pagerduty Users. 
* To fetch and import a user server generated User ID is needed.
## Requirements 
* [Go](https://golang.org/doc/install) <br>
* [Terraform](https://www.terraform.io/downloads.html)  <br/>
* [Pagerduty Developer account](https://www.pagerduty.com/)

## Initialize pagerduty Provider
1. Clone the repository  to from "https://github.com/tharunforu/terraform-provider-pagerduty"<br>
2. Add the Token in `main.tf` in provider for api acess<br>
3. Run the following commands :
 ```
go mod init terraform-provider-pagerduty
go mod tidy
```
4. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Installation
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


## Run the Terraform provider

### Authentication
1. Add the API TOKEN in  provider in  `main.tf` <br>

#### Create User
1. Add the user email,  name,type,and role in the respective field in `main.tf`
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created.

#### Update the user
Update the data of the user in the `main.tf` file and apply using `terraform apply`

#### Read the User Data
Add data and output blocks in the `main.tf` file after that add email field  and user email  and run `terraform plan` to read user data

#### Delete the user
Run `terraform destroy` to delete all the users in `main.tf` or just comment out the unwated user and run `terrafrom apply`
#### Import a User Data
1. Write manually a resource configuration block for the User in `main.tf`, to which the imported object will be mapped or define the empty resource block.
2. Run the command `terraform pagerduty_resource_user.import “[USER_ID]”`
* To get the user Id go to pagerduty dashboard in people select the users then select a particular user in the url last 7 digits will be the user id
eg :  URL:……users/PDFJIJD  here PDFJIJD is the user id of the user
3. Check for the attributes in the `.tfstate` file and fill them accordingly in the resource block.


### Testing the Provider
1. Navigate to the test file directory.
2. Run command `go test` for unit testing and for acceptance testing run command `TF_ACC=1 go test` . These commands will give combined test results for the execution or errors if any failure occurs.
3. If you want to see test result of each test function individually while running test in a single go, run command `go test -v`
4. To check test cover run `go test -cover`


## Example Usage
```
terraform {
  required_providers {
    pagerduty = {
      version = "1.0"
      source  = "tharun/edu/pagerduty"
    }
  }
}

provider "pagerduty" {
  token = [Api Token]
}
data "pagerduty_user_data" "name" {
  id = [userid]
}
output "user" {
  value = data.pagerduty_user_data.name
}
 resource "pagerduty_user_resource" "test1" {
 name  =[Name of he user ]
   email = [Email of user]
   type  = [Type]
   role  = [ROLE of user]
 }
 output "resource_test" {
   value = pagerduty_user_resource.test1
 }


```

## Argument Reference

* `email` (string) - The email id associated with the user account.
* `name` (string) - name of the User.
* `type` (string) - type of the User.
* `role` (string) -  Job role of user in organization. 
