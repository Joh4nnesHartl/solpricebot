package commands

var helpSelector selector = func(words []string) (bool, handler) {
	if len(words) == 2 && words[1] == "help" {
		return true, handleHelp
	}

	return false, nil
}

var defaultSelector selector = func(words []string) (bool, handler) {
	if len(words) == 1 {
		return true, handleDefault
	}

	return false, nil
}

var specificSelector selector = func(words []string) (bool, handler) {
	if len(words) == 2 && words[1] != "help" {
		return true, handleSpecific
	}

	return false, nil
}
