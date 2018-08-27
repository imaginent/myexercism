package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {

	output := "One for " + name + ", one for me."

	if name == "" {
		return "One for you, one for me."
	}
	return output
}
