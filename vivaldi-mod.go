package main

import (
	"log"
	"os"
)

func main() {
	multiline_window := `<!-- Vivaldi window document -->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <title>Vivaldi</title>
  <link rel="stylesheet" href="style/common.css" />
  <link rel="stylesheet" href="chrome://vivaldi-data/css-mods/css" />
</head>

<body>
  <script src="custom.js"></script>
</body>

</html>
`
	err := os.WriteFile("/opt/vivaldi/resources/vivaldi/window.html", []byte(multiline_window), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
