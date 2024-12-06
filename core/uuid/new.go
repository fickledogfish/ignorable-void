package uuid

// Creates a new random UUID or panics. It is equivalent to
//
//	uuid.Must(uuid.NewRandom())
func New() UUID {
	return Must(NewRandom())
}
