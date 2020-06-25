package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//<-chan - canal somente leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string) //cria um channel sem buffer
	for _, url := range urls {
		go func(cUrl string) { //chama uma função anônima com goroutine, pro processamento não ficar amarrado esperando retorno
			resp, _ := http.Get(cUrl)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("Primeiro retorno:", <-t1, "|", <-t2)
	fmt.Println("Segundo retorno:", <-t1, "|", <-t2)
}
