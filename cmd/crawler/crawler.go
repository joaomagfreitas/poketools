package main

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"path"
	"slices"
	"sync"
	"time"

	"github.com/joaomagfreitas/stdx/httpx"
	"github.com/joaomagfreitas/stdx/slicesx"
)

// Mimic Safari browser to avoid bot detection
var ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 15_7_2) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15"

// Generate urls of pages to crawl
var pages = slicesx.Gen(151, func(index int) string {
	return fmt.Sprintf("https://pokemondb.net/pokedex/%d", (index + 1))
})

// Max number of requests simultaneously
var parallel = 12

// Time variance before firing requests
var wait = 3200 * time.Millisecond

func main() {
	h := http.Header{}
	h.Set("user-agent", ua)

	for pages := range slices.Chunk(pages, parallel) {
		var wg sync.WaitGroup
		wg.Add(len(pages))

		for _, p := range pages {
			go func() {
				defer func() {
					wg.Done()
				}()

				w := time.Duration(rand.UintN(uint(wait)))
				log.Printf("> sleeping for %v before crawling %s", w, p)
				time.Sleep(w)

				log.Printf("> crawling for %s", p)
				resp, err := httpx.Get(p, nil, h)

				if err != nil {
					log.Printf("> failed to request %s, %v", p, err)
					return
				}

				if resp.StatusCode != 200 {
					log.Printf("> expected 200 status but got %d in request to %s, %v", resp.StatusCode, p, err)
					return
				}

				bs, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Printf("> failed to read response body, %v", err)
				}

				log.Printf("> saving %s in disk", p)
				fp := path.Base(resp.Request.URL.Path)
				err = os.WriteFile(fp, bs, 0660)
				if err != nil {
					log.Printf("> failed to write %s to disk, %v", p, err)
				}
			}()
		}

		wg.Wait()
	}
}
