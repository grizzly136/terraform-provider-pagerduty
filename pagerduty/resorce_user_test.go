package pagerduty

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccItem_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("pagerduty_user_resource.tharun2", "email", "tharunchunchu@gmail.com"),
					resource.TestCheckResourceAttr("pagerduty_user_resource.tharun2", "name", "tharunkumar"),
					resource.TestCheckResourceAttr("pagerduty_user_resource.tharun2", "type", "user"),
					resource.TestCheckResourceAttr("pagerduty_user_resource.tharun2", "role", "admin"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprint(`
resource "pagerduty_user_resource" "tharun2" {
  email        = "tharunchunchu@gmail.com"
  name   = "tharunkumar"
  type         =  "user"
  role="admin"
}
`)
}
func TestAccItem_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"pagerduty_user_resource.kunal", "email", "kunal@gmail.com"),
					resource.TestCheckResourceAttr(
						"pagerduty_user_resource.kunal", "name", "gohire"),
					resource.TestCheckResourceAttr(
						"pagerduty_user_resource.kunal", "type", "user"),
					resource.TestCheckResourceAttr(
						"pagerduty_user_resource.kunal", "role", "admin"),
				),
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"pagerduty_user_resource.kunal", "email", "kunaln@gmail.com"),
					resource.TestCheckResourceAttr(
						"pagerduty_user_resource.kunal", "name", "kunal"),
					resource.TestCheckResourceAttr(
						"pagerduty_user_resource.kunal", "type", "user"),
					resource.TestCheckResourceAttr(
						"pagerduty_user_resource.kunal", "role", "admin"),
				),
			},
		},
	})
}

func testAccCheckItemUpdatePre() string {
	return fmt.Sprintf(`
	resource "pagerduty_user_resource" "kunal" {
		email        = "kunal@gmail.com"
		name   = "gohire"
		type         =  "user"
		role="admin"
	  }
`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
	resource "pagerduty_user_resource" "kunal" {
		email        = "kunaln@gmail.com"
		name   = "kunal"
		type         =  "user"
		role="admin"
	  }
`)
}
