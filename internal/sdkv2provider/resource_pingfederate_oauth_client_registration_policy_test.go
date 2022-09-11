package sdkv2provider

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClientRegistrationPolicies"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("oauth_client_registration_policy", &resource.Sweeper{
		Name:         "oauth_client_registration_policy",
		Dependencies: []string{},
		F: func(r string) error {
			svc := oauthClientRegistrationPolicies.New(cfg)
			results, _, err := svc.GetDynamicClientRegistrationPolicies()
			if err != nil {
				return fmt.Errorf("unable to list client registration policies %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := svc.DeleteDynamicClientRegistrationPolicy(&oauthClientRegistrationPolicies.DeleteDynamicClientRegistrationPolicyInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep client registration policy %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateOAuthClientRegistrationPolicyResource(t *testing.T) {
	resourceName := "pingfederate_oauth_client_registration_policy.demo"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOAuthClientRegistrationPolicyResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOAuthClientRegistrationPolicyResourceConfig("false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientRegistrationPolicyResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "plugin_descriptor_ref.0.id", "com.pingidentity.pf.client.registration.ResponseTypesConstraintsPlugin"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.0.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.1.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.1.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.2.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.2.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.3.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.4.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.4.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.4.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.5.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.5.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.6.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.6.value"),
					resource.TestCheckNoResourceAttr(resourceName, "configuration.0.fields.7.name"),
					resource.TestCheckNoResourceAttr(resourceName, "configuration.0.fields.7.value"),
				),
			},
			{
				Config: testAccPingFederateOAuthClientRegistrationPolicyResourceConfig("true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientRegistrationPolicyResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "plugin_descriptor_ref.0.id", "com.pingidentity.pf.client.registration.ResponseTypesConstraintsPlugin"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.0.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.1.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.1.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.2.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.2.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.3.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.4.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.4.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.4.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.5.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.5.value"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.6.name"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.fields.6.value"),
					resource.TestCheckNoResourceAttr(resourceName, "configuration.0.fields.7.name"),
					resource.TestCheckNoResourceAttr(resourceName, "configuration.0.fields.7.value"),
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

func testAccCheckPingFederateOAuthClientRegistrationPolicyResourceDestroy(_ *terraform.State) error {
	return nil
}

func testAccPingFederateOAuthClientRegistrationPolicyResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_client_registration_policy" "demo" {
  policy_id = "mypolicy"
  name      = "mypolicy"
  plugin_descriptor_ref {
    id = "com.pingidentity.pf.client.registration.ResponseTypesConstraintsPlugin"
  }
  configuration {
    fields {
      name  = "code"
      value = "%s"
    }
    fields {
      name  = "code id_token"
      value = "false"
    }
    fields {
      name  = "code id_token token"
      value = "false"
    }
    fields {
      name  = "code token"
      value = "false"
    }
    fields {
      name  = "id_token"
      value = "false"
    }
    fields {
      name  = "id_token token"
      value = "false"
    }
    fields {
      name  = "token"
      value = "true"
    }
  }
}`, configUpdate)
}

func testAccCheckPingFederateOAuthClientRegistrationPolicyResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthClientRegistrationPolicies
		result, _, err := conn.GetDynamicClientRegistrationPolicy(&oauthClientRegistrationPolicies.GetDynamicClientRegistrationPolicyInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("error: ClientRegistrationPolicy (%s) not found", n)
		}

		if *result.PluginDescriptorRef.Id != rs.Primary.Attributes["plugin_descriptor_ref.0.id"] {
			return fmt.Errorf("error: ClientRegistrationPolicy response (%s) didnt match state (%s)", *result.PluginDescriptorRef.Id, rs.Primary.Attributes["plugin_descriptor_ref.0.id"])
		}

		return nil
	}
}

type registrationPoliciesMock struct {
	oauthClientRegistrationPolicies.OauthClientRegistrationPoliciesAPI
}

func (s registrationPoliciesMock) GetDynamicClientRegistrationDescriptor(_ *oauthClientRegistrationPolicies.GetDynamicClientRegistrationDescriptorInput) (output *pf.ClientRegistrationPolicyDescriptor, resp *http.Response, err error) {
	return &pf.ClientRegistrationPolicyDescriptor{
		AttributeContract: nil,
		ClassName:         String("com.pingidentity.pf.client.registration.ResponseTypesConstraintsPlugin"),
		ConfigDescriptor: &pf.PluginConfigDescriptor{
			ActionDescriptors: nil,
			Description:       nil,
			Fields:            &[]*pf.FieldDescriptor{},
			Tables: &[]*pf.TableDescriptor{
				{
					Columns: &[]*pf.FieldDescriptor{
						{
							TextFieldDescriptor: pf.TextFieldDescriptor{
								Encrypted: Bool(true),
							},
							Name: String("Password"),
							Type: String("TEXT"),
						},
					},
					Name:              String("Networks"),
					RequireDefaultRow: nil,
				},
			},
		},
		Id:                       String("com.pingidentity.pf.client.registration.ResponseTypesConstraintsPlugin"),
		Name:                     String("REST API"),
		SupportsExtendedContract: nil,
	}, nil, nil
}

func Test_resourcePingFederateOAuthClientRegistrationPolicyResourceReadData(t *testing.T) {
	svc := &registrationPoliciesMock{}
	cases := []struct {
		Resource pf.ClientRegistrationPolicy
	}{
		{
			Resource: pf.ClientRegistrationPolicy{
				Id:   String("example"),
				Name: String("terraform"),
				PluginDescriptorRef: &pf.ResourceLink{
					Id: String("com.pingidentity.pf.client.registration.ResponseTypesConstraintsPlugin"),
				},
				ParentRef: &pf.ResourceLink{
					Id: String("com.pingidentity.pf.client.registration.ResponseTypesConstraintsPlugin"),
				},
				Configuration: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{
						{
							Name:      String("code"),
							Value:     String("true"),
							Inherited: Bool(false),
						},
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateOAuthClientRegistrationPolicyResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOAuthClientRegistrationPolicyResourceReadResult(resourceLocalData, &tc.Resource, svc)

			assert.Equal(t, tc.Resource, *resourcePingFederateOAuthClientRegistrationPolicyResourceReadData(resourceLocalData))
		})
	}
}
