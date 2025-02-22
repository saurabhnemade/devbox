package devbox

func exclude(s []string, elems []string) []string {
	excluded := make(map[string]bool, len(elems))
	for _, ex := range elems {
		excluded[ex] = true
	}

	filtered := make([]string, 0, len(s))
	for _, str := range s {
		if !excluded[str] {
			filtered = append(filtered, str)
		}
	}
	return filtered
}
