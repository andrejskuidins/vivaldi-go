package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	filePath := "/opt/vivaldi/resources/vivaldi/window.html"

	multiline_window, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	content := string(multiline_window)

	// Check if the script is already present
	if !strings.Contains(content, "custom.js") {
		newContent := strings.Replace(content, "<body>", "<body>\n  <script src=\"custom.js\"></script>", 1)

		err = os.WriteFile(filePath, []byte(newContent), 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("File updated successfully")
	} else {
		log.Println("Script already present, no changes made")
	}
}
