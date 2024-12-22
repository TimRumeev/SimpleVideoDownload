package main

import (
	"SimpleDownload/src/internal/lib/urls"
	"fmt"
)

func main() {
	url := "https://ya.ru/"
	domain := urls.GetDomain(url)
	fmt.Println(domain)
}
