package proxy

// Database interface to access
type Database interface {
	Access(url string, port string, user string, pass string) (string, error)
}
