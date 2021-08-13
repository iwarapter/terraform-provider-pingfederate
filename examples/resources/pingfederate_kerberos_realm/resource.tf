resource "pingfederate_kerberos_realm" "demo" {
  kerberos_realm_name      = "test"
  key_distribution_centers = ["foo.com", "bar.com"]
  kerberos_username        = "user"
  kerberos_password        = "secret"
}
