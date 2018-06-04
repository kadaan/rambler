package env

// Environment is the execution environment of a command. It contains every information
// about the database and migrations to use.
type Environment struct {
	Database  string `json:"database"`
	Directory string `json:"directory"`
	Driver    string `json:"driver"`
	Host      string `json:"host"`
	Password  string `json:"password"`
	Port      uint64 `json:"port"`
	Protocol  string `json:"protocol"`
	Table     string `json:"table"`
	User      string `json:"user"`
}
