resource "pingfederate_jdbc_data_store" "demo" {
  name           = "terraform"
  driver_class   = "org.hsqldb.jdbcDriver"
  user_name      = "sa"
  password       = ""
  max_pool_size  = 10
  connection_url = "jdbc:hsqldb:mem:mymemdb"
  connection_url_tags {
    connection_url = "jdbc:hsqldb:mem:mymemdb"
    default_source = true
  }
}
