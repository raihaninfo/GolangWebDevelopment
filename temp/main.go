package main

import "fmt"

func main() {
	name := "Md Abu Raihan"
	heading := "This website heading one"
	paragraph := "this is my website first paragraph"
	webtem := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>` + name + `</title>
	</head>
		<body>
			<h1>` + heading + `</h1>
			<p>` + paragraph + `</p>
		</body>
	</html> 
	`

	fmt.Println(webtem)
}
