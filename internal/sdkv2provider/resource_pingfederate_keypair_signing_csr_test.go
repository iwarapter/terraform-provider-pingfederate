package sdkv2provider

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/certificatesCa"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSigning"
)

func TestAccPingFederateKeyPairSigningCsr(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateKeyPairSigningCsrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateKeyPairSigningCsrConfigGenerateWithCSR(generateAndSignCSR(t)),
			},
		},
	})
}

func generateAndSignCSR(t *testing.T) string {
	if os.Getenv("TF_ACC") != "" {
		key, _ := rsa.GenerateKey(rand.Reader, 2048)
		ca := &x509.Certificate{
			SerialNumber: big.NewInt(2019),
			Subject: pkix.Name{
				Organization: []string{"Testing"},
				Country:      []string{"GB"},
			},
			NotBefore:             time.Now(),
			NotAfter:              time.Now().AddDate(10, 0, 0),
			IsCA:                  true,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
			BasicConstraintsValid: true,
		}

		caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &key.PublicKey, key)
		if err != nil {
			log.Fatalf("Failed to create certificateCa: %s", err)
		}
		caBuf := new(bytes.Buffer)
		_ = pem.Encode(caBuf, &pem.Block{Type: "CERTIFICATE", Bytes: caBytes})
		trustSvc := certificatesCa.New(cfg)
		_, _, _ = trustSvc.ImportTrustedCA(&certificatesCa.ImportTrustedCAInput{Body: pf.X509File{FileData: String(base64.StdEncoding.EncodeToString(caBuf.Bytes()))}})
		svc := keyPairsSigning.New(cfg)
		_, _, err = svc.CreateKeyPair(&keyPairsSigning.CreateKeyPairInput{Body: pf.NewKeyPairSettings{
			City:             String("Test"),
			CommonName:       String("CSR Test"),
			Country:          String("GB"),
			Id:               String("csr-test-1"),
			KeyAlgorithm:     String("RSA"),
			KeySize:          Int(2048),
			Organization:     String("Test"),
			OrganizationUnit: String("Test"),
			State:            String("Test"),
			ValidDays:        Int(365),
		}})
		if err != nil {
			t.Fatalf("unable to create keypair")
		}
		csrPem, _, err := svc.ExportCsr(&keyPairsSigning.ExportCsrInput{Id: "csr-test-1"})
		if err != nil {
			t.Fatalf("unable to get CSR")
		}
		*csrPem = strings.ReplaceAll(*csrPem, " NEW ", " ")
		b, _ := pem.Decode([]byte(*csrPem))
		csr, err := x509.ParseCertificateRequest(b.Bytes)
		if err != nil {
			t.Fatalf("unable to parse csr: %s", err)
		}
		template := x509.Certificate{
			Signature:          csr.Signature,
			SignatureAlgorithm: csr.SignatureAlgorithm,

			PublicKeyAlgorithm: csr.PublicKeyAlgorithm,
			PublicKey:          csr.PublicKey,

			SerialNumber: big.NewInt(2),
			Issuer:       ca.Subject,
			Subject:      csr.Subject,
			NotBefore:    time.Now(),
			NotAfter:     time.Now().Add(24 * time.Hour),
			KeyUsage:     x509.KeyUsageDigitalSignature,
			ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		}
		certBytes, err := x509.CreateCertificate(rand.Reader, &template, ca, csr.PublicKey, key)
		if err != nil {
			t.Fatalf("unable to sign certificate request: %s", err)
		}
		buf := new(bytes.Buffer)
		err = pem.Encode(buf, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
		if err != nil {
			t.Fatalf("unable to encode certificate: %s", err)
		}
		signedCert := buf.String()
		return signedCert
	}
	return ""
}

func testAccCheckPingFederateKeyPairSigningCsrDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateKeyPairSigningCsrConfigGenerateWithCSR(signedCert string) string {
	return fmt.Sprintf(`
resource "pingfederate_keypair_signing_csr" "test" {
  keypair_id = "csr-test-1"
  file_data  = "%s"
}
`, base64.StdEncoding.EncodeToString([]byte(signedCert)))
}
