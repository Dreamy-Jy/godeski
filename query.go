package odeskidb

type query interface {
	Execute() (query, string, error)
}
