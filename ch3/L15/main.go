package main

func reformat(message string, formatter func(string) string) string {
	// Apply the formatter three times
	result := formatter(formatter(formatter(message)))
	
	// Add the prefix "TEXTIO: " and return the result
	return "TEXTIO: " + result
}

// Example formatter function (adds a period at the end)
func addPeriod(message string) string {
	return message + "..."
}

func main() {
	// Example message
	message := "General Kenobi"
	
	// Call reformat with the message and formatter
	formattedMessage := reformat(message, addPeriod)
	
	// Print the result
	fmt.Println(formattedMessage)
}

