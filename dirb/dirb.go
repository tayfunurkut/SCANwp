package dirb

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/cheggaaa/pb"
	"github.com/tayfun8/scanwp/models"
	"github.com/tayfun8/scanwp/utils"
)

var fuzz []models.Dirb

func Dirb(wordlist, target string) {

	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Dosya açma hatası:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	totalRequests := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalRequests++
	}

	bar := pb.StartNew(totalRequests)

	file.Seek(0, 0)

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		url := target + word
		wg.Add(1)
		go func(url, word string) {
			defer wg.Done()
			statusCode, err := utils.SendHTTPGetRequest(url)
			if err != nil {
				fmt.Printf("HTTP isteği hatası (%s): %v\n", url, err)
				return
			}
			result := models.Dirb{
				Path:       url,
				StatusCode: strconv.Itoa(statusCode),
			}
			fuzz = append(fuzz, result)
			bar.Increment()
		}(url, word)
	}

	wg.Wait()
	utils.WriteDirb(fuzz)

	bar.Finish()
}
