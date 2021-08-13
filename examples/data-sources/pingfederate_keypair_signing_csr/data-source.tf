data "pingfederate_keypair_signing_csr" "csr" {
  id = pingfederate_keypair_signing.demo_generate.id
}
