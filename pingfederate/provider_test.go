package pingfederate

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/services/version"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ory/dockertest"
)

func TestMain(m *testing.M) {
	_, acceptanceTesting := os.LookupEnv("TF_ACC")
	if acceptanceTesting {
		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}

		devOpsUser, devOpsUserExists := os.LookupEnv("PING_IDENTITY_DEVOPS_USER")
		devOpsKey, devOpsKeyExists := os.LookupEnv("PING_IDENTITY_DEVOPS_KEY")

		var options *dockertest.RunOptions
		dir, _ := os.Getwd()

		if devOpsUserExists && devOpsKeyExists {
			options = &dockertest.RunOptions{
				Hostname:   "pingfederate",
				Repository: "pingidentity/pingfederate",
				Mounts:     []string{dir + "/pingfederate-data.zip:/opt/in/instance/server/default/data/drop-in-deployer/data.zip"},
				Env:        []string{fmt.Sprintf("PING_IDENTITY_DEVOPS_USER=%s", devOpsUser), fmt.Sprintf("PING_IDENTITY_DEVOPS_KEY=%s", devOpsKey), "PING_IDENTITY_ACCEPT_EULA=YES"},
				Tag:        "10.0.2-edge",
			}
		} else {
			options = &dockertest.RunOptions{
				Hostname:   "pingfederate",
				Repository: "pingidentity/pingfederate",
				Mounts: []string{
					dir + "/pingfederate.lic:/opt/in/instance/server/default/conf/pingfederate.lic",
					dir + "/pingfederate-data.zip:/opt/in/instance/server/default/data/drop-in-deployer/data.zip",
				},
				Tag: "10.0.2-edge",
			}
		}

		// pulls an image, creates a container based on it and runs it
		resource, err := pool.RunWithOptions(options)
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}
		pool.MaxWait = time.Minute * 2

		// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
		if err := pool.Retry(func() error {
			var err error
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
			pfUrl, _ := url.Parse(fmt.Sprintf("https://localhost:%s/pf-admin-api/v1", resource.GetPort("9999/tcp")))
			client := version.New(config.NewConfig().WithUsername("Administrator").WithPassword("2Federate").WithEndpoint(pfUrl.String()))

			log.Println("Attempting to connect to PingFederate admin API....")
			_, _, err = client.GetVersion()
			return err
		}); err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}
		resource.Expire(180)
		os.Setenv("PINGFEDERATE_BASEURL", fmt.Sprintf("https://localhost:%s", resource.GetPort("9999/tcp")))
		log.Println("Connected to PingFederate admin API....")
		code := m.Run()
		log.Println("Tests complete shutting down container")

		// You can't defer this because os.Exit doesn't care for defer
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}

		os.Exit(code)
	} else {
		m.Run()
	}
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

//var testAccProviderFactories func(providers *[]*schema.Provider) map[string]*schema.Provider
//var testAccTemplateProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	//testAccTemplateProvider = template.Provider().(*schema.Provider)
	testAccProviders = map[string]*schema.Provider{
		"pingfederate": testAccProvider,
	}
	//testAccProviderFactories = func(providers *[]*schema.Provider) map[string]*schema.Provider {
	//	return map[string]*schema.Provider{
	//		"pingfederate": func() (schema.Provider, error) {
	//			p := Provider()
	//			*providers = append(*providers, p.(*schema.Provider))
	//			return p, nil
	//		},
	//	}
	//}
}

func testAccPreCheck(t *testing.T) {
	err := testAccProvider.Configure(context.TODO(), terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
