package main

import (
	"fmt"
	"time"

	html "github.com/ercileycoimbra/curso_go_pkg_html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	select {
	case a := <-c1:
		return a
	case b := <-c2:
		return b
	case c := <-c3:
		return c
	case <-time.After(time.Millisecond * 300):
		return "Timeout - Todos perderam"
	}
}

func main() {
	c1 := oMaisRapido("https://www.amazon.com", "https://www.youtube.com.br", "https://www.google.com")
	fmt.Println(c1)

}
