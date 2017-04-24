package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	BASEURL     = "https://tuchong.com/267406/"
	BASEIMGPATH = "../xiyuanpeng_front/public/images/"
)

var (
	req           = fasthttp.AcquireRequest()
	resp          = fasthttp.AcquireResponse()
	secondUrlChan = make(chan [2]string, 10)
	imgUrlChan    = make(chan string, 10)
)

func main() {
	generalWg := &sync.WaitGroup{}
	go getUrls()
	generalWg.Add(1)
	go getImgUrls(generalWg)
	fmt.Println("test")
	generalWg.Wait()
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
}

func getImgUrls(generalWg *sync.WaitGroup) {
	defer generalWg.Done()
	wg := &sync.WaitGroup{}
	wg1 := &sync.WaitGroup{}
	for titlePair := range secondUrlChan {
		classChan := make(chan string, 10)
		title, url := titlePair[0], titlePair[1]
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			specPage, queryErr := goquery.NewDocument(url)
			if queryErr != nil {
				fmt.Println(queryErr)
				return
			}
			specPage.Find(".multi-photo-image").Each(func(i int, s *goquery.Selection) {
				imgUrl, _ := s.Attr("src")
				classChan <- imgUrl
			})
		}(url)
		wg1.Add(1)
		go getImgs(title, classChan, wg1)
	}
	wg.Wait()
	wg1.Wait()
	fmt.Println("getImgUrls")
}

func getImgs(title string, imgChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	controlChan := make(chan struct{}, 10)
	wg1 := &sync.WaitGroup{}
	for imgUrl := range imgChan {
		imgName := strings.Split(imgUrl, "/")[len(strings.Split(imgUrl, "/"))-1]
		controlChan <- struct{}{}
		wg1.Add(1)
		go func(imgUrl string, wg *sync.WaitGroup) {
			defer wg.Done()
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
			defer imgFile.Close()
			fmt.Println("saving", BASEIMGPATH+title+"/"+imgName)
			<-controlChan
		}(imgUrl, wg1)
	}
	fmt.Println("test getImgs")
	wg1.Wait()
	fmt.Println("test getImgs")
}
