package sdkv2provider

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/notificationPublishers"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("notification_publisher", &resource.Sweeper{
		Name:         "notification_publisher",
		Dependencies: []string{},
		F: func(r string) error {
			svc := notificationPublishers.New(cfg)
			results, _, err := svc.GetNotificationPublishers()
			if err != nil {
				return fmt.Errorf("unable to list notification publishers %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := svc.DeleteNotificationPublisher(&notificationPublishers.DeleteNotificationPublisherInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep notification publisher %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateNotificationPublisher(t *testing.T) {
	resourceName := "pingfederate_notification_publisher.demo"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateNotificationPublisherDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateNotificationPublisherConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateNotificationPublisherExists(resourceName),
					//testAccCheckPingFederateNotificationPublisherAttributes(),
				),
			},
			{
				Config: testAccPingFederateNotificationPublisherConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateNotificationPublisherExists(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config:      testAccPingFederateNotificationPublisherConfigWrongPlugin(),
				ExpectError: regexp.MustCompile(`unable to find plugin_descriptor for com\.pingidentity\.adapters\.opentoken\.wrong available plugins:`),
			},
		},
	})
}

func testAccCheckPingFederateNotificationPublisherDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateNotificationPublisherConfig(configUpdate string) string {
	return fmt.Sprintf(`
data "pingfederate_version" "instance" {}

locals {
  isSupported = length(regexall("(10|11).[0-9]", data.pingfederate_version.instance.version)) > 0
  isPF11_2    = length(regexall("(11).[2-9]", data.pingfederate_version.instance.version)) > 0
}

resource "pingfederate_notification_publisher" "demo" {
  name         = "acc_test_bar"
  publisher_id = "foo1"
  plugin_descriptor_ref {
    id = "com.pingidentity.email.SmtpNotificationPlugin"
  }

  configuration {
    dynamic "fields" {
      for_each = local.isSupported ? [1] : []
      content {
        name  = "UTF-8 Message Header Support"
        value = "false"
      }
    }
    dynamic "fields" {
      for_each = local.isPF11_2 ? [1] : []
      content {
        name  = "Sender Name"
        value = "false"
      }
    }
    fields {
      name  = "From Address"
      value = "help@foo.org"
    }
    fields {
      name  = "Email Server"
      value = "%s"
    }
    fields {
      name  = "SMTP Port"
      value = "25"
    }
    fields {
      name  = "Encryption Method"
      value = "NONE"
    }
    fields {
      name  = "SMTPS Port"
      value = "465"
    }
    fields {
      name  = "Verify Hostname"
      value = "true"
    }
    fields {
      name  = "Username"
      value = ""
    }
    fields {
      name  = "Password"
      value = ""
    }
    fields {
      name  = "Test Address"
      value = ""
    }
    fields {
      name  = "Connection Timeout"
      value = "30"
    }
    fields {
      name  = "Retry Attempt"
      value = "2"
    }
    fields {
      name  = "Retry Delay"
      value = "2"
    }
    fields {
      name  = "Enable SMTP Debugging Messages"
      value = "false"
    }
  }
}
`, configUpdate)
}

func testAccPingFederateNotificationPublisherConfigWrongPlugin() string {
	return `
resource "pingfederate_notification_publisher" "demo" {
  name         = "acc_test_bar2"
  publisher_id = "acc_test_bar2"
  plugin_descriptor_ref {
    id = "com.pingidentity.adapters.opentoken.wrong"
  }

  configuration {
    fields {
      name  = "Use Verbose Error Messages"
      value = "false"
    }
  }
}
`
}

func testAccCheckPingFederateNotificationPublisherExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).NotificationPublishers
		result, _, err := conn.GetNotificationPublisher(&notificationPublishers.GetNotificationPublisherInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: NotificationPublisher (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: NotificationPublisher response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

type notificationPublishersMock struct {
	notificationPublishers.NotificationPublishersAPI
}

func (s notificationPublishersMock) GetNotificationPublisherPluginDescriptor(input *notificationPublishers.GetNotificationPublisherPluginDescriptorInput) (output *pf.NotificationPublisherDescriptor, resp *http.Response, err error) {
	return &pf.NotificationPublisherDescriptor{
		AttributeContract: nil,
		ClassName:         String("com.pingidentity.pf.selectors.cidr.CIDRAdapterSelector"),
		ConfigDescriptor: &pf.PluginConfigDescriptor{
			ActionDescriptors: nil,
			Description:       nil,
			Fields:            &[]*pf.FieldDescriptor{},
			Tables: &[]*pf.TableDescriptor{
				{
					Columns: &[]*pf.FieldDescriptor{
						{
							Type: String("TEXT"),
							Name: String("Username"),
						},
					},
					Description:       nil,
					Label:             nil,
					Name:              String("Networks"),
					RequireDefaultRow: nil,
				},
			},
		},
		Id:                       String("com.pingidentity.pf.selectors.cidr.CIDRAdapterSelector"),
		Name:                     String("CIDR Authentication Selector"),
		SupportsExtendedContract: nil,
	}, nil, nil
}

func Test_resourcePingFederateNotificationPublisherResourceReadData(t *testing.T) {
	m := &notificationPublishersMock{}
	cases := []struct {
		Resource pf.NotificationPublisher
	}{
		{
			Resource: pf.NotificationPublisher{
				Name: String("foo"),
				Id:   String("foo"),
				PluginDescriptorRef: &pf.ResourceLink{
					Id: String("com.pingidentity.email.SmtpNotificationPlugin"),
				},
				Configuration: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{
						{
							Name:      String("Result Attribute Name"),
							Value:     String(""),
							Inherited: Bool(false),
						},
					},
				},
				ParentRef: &pf.ResourceLink{
					Id: String("foo"),
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			resourceSchema := resourcePingFederateNotificationPublisherResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateNotificationPublisherResourceReadResult(resourceLocalData, &tc.Resource, m)

			assert.Equal(t, tc.Resource, *resourcePingFederateNotificationPublisherResourceReadData(resourceLocalData))
		})
	}
}
