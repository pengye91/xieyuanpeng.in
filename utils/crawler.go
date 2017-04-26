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
	fmt.Println("test")
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
			}
		}()
		subWG1.Add(1)
		func(url string, subwg *sync.WaitGroup) {
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
		}(url, subWG1)
		fmt.Printf("%v\n", classChan)
		subWG2.Add(1)
		go getImgs(title, &classChan, subWG2)
	}
	fmt.Println("new new new new")
	subWG1.Wait()
	subWG2.Wait()
	// close(classChan)
	fmt.Println("getImgUrls")
}

func getImgs(title string, imgChan *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	// controlChan := make(chan struct{}, 10)
	fmt.Printf("%v %s\n", *imgChan, "xixixi")
	subWG := &sync.WaitGroup{}
	for _, imgUrl := range *imgChan {
		fmt.Println(imgUrl)
		imgName := strings.Split(imgUrl, "/")[len(strings.Split(imgUrl, "/"))-1]
		// controlChan <- struct{}{}
		fmt.Println(title + "/" + imgName)
		subWG.Add(1)
		go func(imgUrl string) {
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
			fmt.Println("test test test test")
			// subWG.Done()
		}(imgUrl)
		fmt.Println("test getImgs 1")
	}
	fmt.Println("test 2")
	subWG.Wait()
	fmt.Println("test getImgs")
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
