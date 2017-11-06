package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
)

func downloadFromURL(url string, outputFolder string) {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(outputFolder, url)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
}

func downloadMultipleFiles(urls []string, outputFolder string) {
	log.Printf("-> Downloading %d files...\n", len(urls))
	_, err := grab.GetBatch(3, outputFolder, urls...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	log.Printf("** DONE: %d files successfully downloaded.\n", len(urls))
}
