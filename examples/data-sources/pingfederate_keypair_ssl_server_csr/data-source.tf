data "pingfederate_keypair_ssl_server_csr" "csr" {
  id = pingfederate_keypair_ssl_server.demo_generate.id
}
