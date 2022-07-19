package l2sellerprofile

import (
	"fmt"
	"testing"

	"github.com/equinix/terraform-provider-equinix/internal/tfacc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const tstL2SellerProfileEnvVar = "TF_ACC_ECX_SELLER_PROFILE_NAME"

func TestAccDataSourceECXL2SellerProfile_basic(t *testing.T) {
	profileName, _ := schema.EnvDefaultFunc(tstL2SellerProfileEnvVar, "AWS Direct Connect")()
	context := map[string]interface{}{
		"resourceName": "tf-aws",
		"name":         profileName,
	}
	resourceName := fmt.Sprintf("data.equinix_ecx_l2_sellerprofile.%s", context["resourceName"].(string))
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { tfacc.PreCheck(t) },
		Providers: tfacc.AccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceECXL2SellerProfileConfig_basic(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "uuid"),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
					resource.TestCheckResourceAttrSet(resourceName, "speed_from_api"),
					resource.TestCheckResourceAttrSet(resourceName, "speed_customization_allowed"),
					resource.TestCheckResourceAttrSet(resourceName, "redundancy_required"),
					resource.TestCheckResourceAttrSet(resourceName, "encapsulation"),
					resource.TestCheckResourceAttrSet(resourceName, "organization_name"),
				),
			},
		},
	})
}

func testAccDataSourceECXL2SellerProfileConfig_basic(ctx map[string]interface{}) string {
	return tfacc.NPrintf(`
data "equinix_ecx_l2_sellerprofile" "%{resourceName}" {
  name = "%{name}"
}
`, ctx)
}