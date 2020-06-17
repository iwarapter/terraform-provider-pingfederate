package metadata

// ClientInfo wraps immutable data from the client.Client structure.
type ClientInfo struct {
	ServiceName string
	APIVersion  string
	Endpoint    string
	BasePath    string
}
