package main

import (
	"fmt"
	"log"
	"mathops/calculator"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>tes</title>
			</head>
			<body>
				<form action="/calculate" method="get">
					<input type="text" name="number1" placeholder="Number 1">
					<input type="text" name="number2" placeholder="Number 2">
					<select name="operator">
						<option value="+">Tambah</option>
						<option value="-">Kurang</option>
						<option value="*">Kali</option>
						<option value="/">Bagi</option>
						<option value="^">Pangkat</option>
					</select>
					<button type="submit">Hitung</button>
				</form>
			</body>
			</html>
			`
		fmt.Fprint(w, html)
	})

	http.HandleFunc("/calculate", calculator.CalculateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
