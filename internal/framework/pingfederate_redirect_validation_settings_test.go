package framework

import (
	"context"
	"fmt"
	"testing"

	fresource "github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateRedirectValidationSettingsResource(t *testing.T) {
	resourceName := "pingfederate_redirect_validation_settings.settings"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateRedirectValidationSettingsResourceConfig("demo"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_in_error_resource_validation", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_idp_discovery", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_slo", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_sso", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.allow_query_and_fragment", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.idp_discovery", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.in_error_resource", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.require_https", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.target_resource_slo", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.target_resource_sso", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.valid_domain", "*.foo"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.valid_path", "/demo"),
				),
			},
			{
				Config: testAccPingFederateRedirectValidationSettingsResourceConfig("update"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_in_error_resource_validation", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_idp_discovery", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_slo", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_sso", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.allow_query_and_fragment", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.idp_discovery", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.in_error_resource", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.require_https", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.target_resource_slo", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.target_resource_sso", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.valid_domain", "*.foo"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.valid_path", "/update"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccPingFederateRedirectValidationSettingsResourceDefaultValues(t *testing.T) {
	resourceName := "pingfederate_redirect_validation_settings.settings"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateRedirectValidationSettingsResourceConfigWithDefaults("bar.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_in_error_resource_validation", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_idp_discovery", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_slo", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_sso", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.allow_query_and_fragment", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.idp_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.in_error_resource", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.require_https", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.target_resource_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.target_resource_sso", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.valid_domain", "bar.com"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.valid_path", ""),
				),
			},
			{
				Config: testAccPingFederateRedirectValidationSettingsResourceConfigWithDefaults("foo.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_in_error_resource_validation", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_idp_discovery", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_slo", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.enable_target_resource_validation_for_sso", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.allow_query_and_fragment", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.idp_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.in_error_resource", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.require_https", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.target_resource_slo", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.target_resource_sso", "false"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.valid_domain", "foo.com"),
					resource.TestCheckResourceAttr(resourceName, "redirect_validation_local_settings.white_list.0.valid_path", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccPingFederateRedirectValidationSettingsResourceConfigWithDefaults(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_redirect_validation_settings" "settings" {
  redirect_validation_local_settings = {
    enable_in_error_resource_validation                 = true
    enable_target_resource_validation_for_idp_discovery = true
    enable_target_resource_validation_for_slo           = true
    enable_target_resource_validation_for_sso           = true
    white_list = [
      {
        valid_domain = "%s"
      }
    ]
  }
}`, configUpdate)
}

func testAccPingFederateRedirectValidationSettingsResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_redirect_validation_settings" "settings" {
  redirect_validation_local_settings = {
    enable_in_error_resource_validation                 = true
    enable_target_resource_validation_for_idp_discovery = true
    enable_target_resource_validation_for_slo           = true
    enable_target_resource_validation_for_sso           = true
    white_list = [
      {
        allow_query_and_fragment = true
        idp_discovery            = true
        in_error_resource        = true
        require_https            = true
        target_resource_slo      = true
        target_resource_sso      = true
        valid_domain             = "*.foo"
        valid_path               = "/%s"
      }
    ]
  }
  redirect_validation_partner_settings = {
    enable_wreply_validation_slo = true
  }
}`, configUpdate)
}

func Test_resourcePingFederateRedirectValidationSettingsResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.RedirectValidationSettings
	}{
		{
			Resource: pf.RedirectValidationSettings{
				RedirectValidationLocalSettings: &pf.RedirectValidationLocalSettings{
					EnableInErrorResourceValidation:               Bool(true),
					EnableTargetResourceValidationForIdpDiscovery: Bool(true),
					EnableTargetResourceValidationForSLO:          Bool(true),
					EnableTargetResourceValidationForSSO:          Bool(true),
					WhiteList: &[]*pf.RedirectValidationSettingsWhitelistEntry{
						{
							AllowQueryAndFragment: Bool(true),
							IdpDiscovery:          Bool(true),
							InErrorResource:       Bool(true),
							RequireHttps:          Bool(true),
							TargetResourceSLO:     Bool(true),
							TargetResourceSSO:     Bool(true),
							ValidDomain:           String("*.foo.com"),
							ValidPath:             String("/bar"),
						},
					},
				},
				RedirectValidationPartnerSettings: &pf.RedirectValidationPartnerSettings{
					EnableWreplyValidationSLO: Bool(true),
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			res := &pingfederateRedirectValidationSettingsResource{}
			ctx := context.Background()
			schResp := &fresource.SchemaResponse{}
			res.Schema(ctx, fresource.SchemaRequest{}, schResp)
			require.False(t, schResp.Diagnostics.HasError())

			state := tfsdk.State{Schema: schResp.Schema}
			require.False(t, state.Set(ctx, flattenRedirectValidationSettings(&tc.Resource)).HasError())

			check := RedirectValidationSettingsData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandRedirectValidationSettings(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
}
