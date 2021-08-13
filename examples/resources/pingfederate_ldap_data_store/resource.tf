resource "pingfederate_ldap_data_store" "demo_ldap" {
  name             = "terraform_ldap"
  ldap_type        = "PING_DIRECTORY"
  hostnames        = ["host.docker.internal:1389"]
  bind_anonymously = true
  min_connections  = 1
  max_connections  = 1
}
