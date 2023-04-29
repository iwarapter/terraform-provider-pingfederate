data "pingfederate_keypair_ssl_server_certificate" "csr" {
  id = pingfederate_keypair_ssl_server.demo_generate.id
}
