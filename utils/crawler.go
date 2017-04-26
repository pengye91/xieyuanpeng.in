package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"
)

const (
	BASEURL     = "https://tuchong.com/267406/"
	BASEIMGPATH = "../xiyuanpeng_front/public/images2/"
)

var (
	req           = fasthttp.AcquireRequest()
	resp          = fasthttp.AcquireResponse()
	secondUrlChan = make(chan [2]string, 10)
	imgUrlChan    = make(chan string, 10)
)

func main() {
	go getUrls()
	mainWg := &sync.WaitGroup{}
	mainWg.Add(1)
	go getImgUrls(mainWg)
	mainWg.Wait()
	fmt.Println("Done")
}

func getImgUrls(wg *sync.WaitGroup) {
	defer wg.Done()
	subWG1 := &sync.WaitGroup{}
	subWG2 := &sync.WaitGroup{}
	for titlePair := range secondUrlChan {
		// classChan := make(chan string)
		classChan := []string{}
		title, url := titlePair[0], titlePair[1]
		go func() {
			if err := os.Mkdir(BASEIMGPATH+title, os.ModePerm); err != nil {
				fmt.Println("nothing hurt, just return")
				return
			}
		}()
		subWG1.Add(1)
		go func(url string, subwg *sync.WaitGroup) {
			defer subwg.Done()
			specPage, queryErr := goquery.NewDocument(url)
			if queryErr != nil {
				fmt.Println(queryErr)
				return
			}
			specPage.Find(".multi-photo-image").Each(func(i int, s *goquery.Selection) {
				imgUrl, _ := s.Attr("src")
				// classChan <- imgUrl
				func(s *[]string) {
					*s = append(*s, imgUrl)
				}(&classChan)
				// fmt.Printf("%v\n", classChan)
			})
			subWG2.Add(1)
			go getImgs(title, &classChan, subWG2)
		}(url, subWG1)
	}
	fmt.Println("1")
	subWG1.Wait()
	fmt.Println("2")
	subWG2.Wait()
	fmt.Println("3")
}

func getImgs(title string, imgChan *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	// controlChan := make(chan struct{}, 10)
	subWG := &sync.WaitGroup{}
	for _, imgUrl := range *imgChan {
		imgName := strings.Split(imgUrl, "/")[len(strings.Split(imgUrl, "/"))-1]
		// controlChan <- struct{}{}
		subWG.Add(1)
		go func(imgUrl string) {
			defer subWG.Done()
			client := &fasthttp.Client{}
			req.SetRequestURI(imgUrl)
			if clientDoErr := client.DoTimeout(req, resp, 30*time.Second); clientDoErr != nil {
				fmt.Println(clientDoErr)
				return
			}
			imgBody := resp.Body()
			imgFile, createErr := os.Create(BASEIMGPATH + title + "/" + imgName)
			if createErr != nil {
				fmt.Println(createErr)
				return
			}
			_, writeErr := imgFile.Write(imgBody)
			if writeErr != nil {
				fmt.Println(writeErr)
				return
			}
			fmt.Println("saving", BASEIMGPATH+title+"/"+imgName)
			// <-controlChan
			imgFile.Close()
		}(imgUrl)
		fmt.Println("4")
	}
	fmt.Println("5")
	subWG.Wait()
	fmt.Println("6")
}

func getUrls() {
	file, err := os.Open("./index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	basePage, queryErr := goquery.NewDocumentFromReader(file)
	if queryErr != nil {
		fmt.Println(queryErr)
		return
	}
	basePage.Find(".post-photo").Each(func(i int, s *goquery.Selection) {
		var titlePair [2]string
		secondUrl, _ := s.Attr("data-url")
		title := strings.TrimSpace(s.Find(".post-title").Text())
		titlePair[0], titlePair[1] = title, secondUrl
		secondUrlChan <- titlePair
	})
	close(secondUrlChan)
}
