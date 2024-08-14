package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	multiline_window, err := os.ReadFile("/opt/vivaldi/resources/vivaldi/window.html")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("window.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content := string(multiline_window)
	newContent := strings.Replace(content, "<body>", "<body>\n  <script src=\"custom.js\"></script>", 1)

	err = os.WriteFile("/opt/vivaldi/resources/vivaldi/window.html", []byte(newContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
