package main

func reformat(message string, formatter func(string) string) string {
	// Apply the formatter three times
	result := formatter(formatter(formatter(message)))
	
	// Add the prefix "TEXTIO: " and return the result
	return "TEXTIO: " + result
}

