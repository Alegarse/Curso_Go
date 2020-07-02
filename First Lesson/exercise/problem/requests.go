package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	sites := []string{
		"https://www.google.com",
		"https://drive.googe.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(len(sites))

	for _, site := range sites {
		go func(site string) {
			defer wg.Done()

			res, err := http.Get(site)
			if err != nil {
				cancel()
				fmt.Println("El link " + site + " es erróneo")
			}

			select {
			case <-ctx.Done():
				return
			default:
				io.WriteString(os.Stdout, res.Status+"\n")
			}
		}(site)
	}
	wg.Wait()
}
