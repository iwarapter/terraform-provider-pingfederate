package pingfederate

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicies"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("authentication_policy_fragment", &resource.Sweeper{
		Name:         "authentication_policy_fragment",
		Dependencies: []string{"authentication_policies"},
		F: func(r string) error {
			svc := authenticationPolicies.New(cfg)
			results, _, err := svc.GetFragments(&authenticationPolicies.GetFragmentsInput{})
			if err != nil {
				return fmt.Errorf("unable to list authentication policy fragments %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := svc.DeleteFragment(&authenticationPolicies.DeleteFragmentInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep authentication policy fragment %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateAuthenticationPolicyFragmentResource(t *testing.T) {
	re := regexp.MustCompile(`^((10)\.[2-9])`)
	if !re.MatchString(pfVersion) {
		t.Skipf("This test only runs against PingFederate 10.2 and above, not: %s", pfVersion)
	}
	resourceName := "pingfederate_authentication_policy_fragment.demo"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateAuthenticationPolicyFragmentResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateAuthenticationPolicyFragmentResourceConfig("example"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPolicyFragmentResourceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "inputs.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "outputs.0.id"),
					resource.TestCheckResourceAttr(resourceName, "description", "functional test example"),
				),
			},
			{
				Config: testAccPingFederateAuthenticationPolicyFragmentResourceConfig("words"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPolicyFragmentResourceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "inputs.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "outputs.0.id"),
					resource.TestCheckResourceAttr(resourceName, "description", "functional test words"),
					resource.TestCheckResourceAttr(resourceName, "root_node.0.action.0.type", "AUTHN_SELECTOR"),
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

func testAccCheckPingFederateAuthenticationPolicyFragmentResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateAuthenticationPolicyFragmentResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`resource "pingfederate_authentication_policy_fragment" "demo" {
  name = "fragtest"
  description = "functional test %s"
  inputs {
	id = pingfederate_authentication_policy_contract.input.id
  }
  outputs {
	id = pingfederate_authentication_policy_contract.output.id
  }

  root_node {
	action {
      type = "AUTHN_SELECTOR"
	  authentication_selector_ref {
		id = pingfederate_authentication_selector.demo.id
	  }
	}
	children {
	  action {
		type    = "DONE"
		context = "No"
	  }
	}
	children {
	  action {
		type    = "DONE"
		context = "Yes"
	  }
	}
  }
}

resource "pingfederate_authentication_policy_contract" "input" {
  name                = "fragmenttest1"
  extended_attributes = ["one", "two"]
}

resource "pingfederate_authentication_policy_contract" "output" {
  name                = "fragmenttest2"
  extended_attributes = ["three", "four"]
}

resource "pingfederate_authentication_selector" "demo" {
  name = "fragmenttest"
  plugin_descriptor_ref {
    id = "com.pingidentity.pf.selectors.cidr.CIDRAdapterSelector"
  }

  configuration {
    fields {
      name  = "Result Attribute Name"
      value = ""
    }
    tables {
      name = "Networks"
      rows {
        fields {
          name  = "Network Range (CIDR notation)"
          value = "127.0.0.1/32"
        }
      }
    }
  }
}
`, configUpdate)
}

func testAccCheckPingFederateAuthenticationPolicyFragmentResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).AuthenticationPolicies
		result, _, err := conn.GetFragment(&authenticationPolicies.GetFragmentInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("error: AuthenticationPolicyFragment (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("error: AuthenticationPolicyFragment response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateAuthenticationPolicyFragmentResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.AuthenticationPolicyFragment
	}{
		{
			Resource: pf.AuthenticationPolicyFragment{
				Name: String("foo"),
				Id:   String("foo"),
				Inputs: &pf.ResourceLink{
					Id: String("1"),
				},
				Outputs: &pf.ResourceLink{
					Id: String("1"),
				},
				RootNode: &pf.AuthenticationPolicyTreeNode{
					Action: &pf.PolicyAction{
						ApcMappingPolicyAction: pf.ApcMappingPolicyAction{
							AttributeMapping: &pf.AttributeMapping{
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
									"foo": {
										Value:  String("bar"),
										Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
									},
								},
								AttributeSources: &[]*pf.AttributeSource{},
							},
							AuthenticationPolicyContractRef: &pf.ResourceLink{Id: String("foo")},
						},
						LocalIdentityMappingPolicyAction: pf.LocalIdentityMappingPolicyAction{
							LocalIdentityRef: &pf.ResourceLink{Id: String("foo")},
							InboundMapping: &pf.AttributeMapping{
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
									"foo": {
										Value:  String("bar"),
										Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
									},
								},
								AttributeSources: &[]*pf.AttributeSource{},
							},
							OutboundAttributeMapping: &pf.AttributeMapping{
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
									"foo": {
										Value:  String("bar"),
										Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
									},
								},
								AttributeSources: &[]*pf.AttributeSource{},
							},
						},
						AuthnSelectorPolicyAction: pf.AuthnSelectorPolicyAction{
							AuthenticationSelectorRef: &pf.ResourceLink{Id: String("foo")},
						},
						AuthnSourcePolicyAction: pf.AuthnSourcePolicyAction{
							AuthenticationSource: &pf.AuthenticationSource{
								SourceRef: &pf.ResourceLink{Id: String("foo")},
								Type:      String("foo"),
							},
							InputUserIdMapping: &pf.AttributeFulfillmentValue{
								Value:  String("bar"),
								Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
							},
							AttributeRules: &pf.AttributeRules{
								FallbackToSuccess: Bool(true),
								Items: &[]*pf.AttributeRule{
									{
										AttributeName: String("foo"),
										Condition:     String("foo"),
										ExpectedValue: String("foo"),
										Result:        String("foo"),
									},
								},
							},
						},
						FragmentPolicyAction: pf.FragmentPolicyAction{
							Fragment: &pf.ResourceLink{
								Id: String("2"),
							},
							FragmentMapping: &pf.AttributeMapping{
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
									"foo": {
										Value:  String("bar"),
										Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
									},
								},
								AttributeSources: &[]*pf.AttributeSource{},
							},
						},
						Context: String("Success"),
						Type:    String("DONE"),
					},
					Children: &[]*pf.AuthenticationPolicyTreeNode{
						{
							Action: &pf.PolicyAction{
								Type:    String("RESTART"),
								Context: String("Fail"),
							},
							Children: &[]*pf.AuthenticationPolicyTreeNode{
								{
									Action: &pf.PolicyAction{
										Type:    String("DONE"),
										Context: String("Fail"),
									},
								},
								{
									Action: &pf.PolicyAction{
										Type:    String("CONTINUE"),
										Context: String("Fail"),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			resourceSchema := resourcePingFederateAuthenticationPolicyFragmentResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateAuthenticationPolicyFragmentResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateAuthenticationPolicyFragmentResourceReadData(resourceLocalData))
		})
	}
}
