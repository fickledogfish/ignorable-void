package uuid

// If err is nil, returns uuid, otherwise, it panics with err.
func Must(uuid UUID, err error) UUID {
	if err != nil {
		panic(err)
	}

	return uuid
}
