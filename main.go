package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Directories Hunter")
	var site string
	fmt.Println("Enter the site: ")
	fmt.Scanf("%s", &site)

	file, _ := os.Open("wordlist.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		subdomain := scanner.Text()

		url := fmt.Sprintf("%s%s", site, subdomain)

		http_response, err := http.Get(url)

		if err != nil {
			continue
		}
		if http_response.StatusCode < 400 || http_response.StatusCode >= 500 {
			fmt.Println(url, "|", http_response.StatusCode)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
}
