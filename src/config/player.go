package config

// SamplePlayer is a sample player listed in the status and query info of a Java Edition server.
type SamplePlayer struct {
	Username string `json:"name" yaml:"username"`
	UUID     string `json:"id" yaml:"uuid"`
}
