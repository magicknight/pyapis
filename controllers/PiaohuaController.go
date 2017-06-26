package controllers

import (
	"github.com/astaxie/beego"
	"pyapis/models"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"pyapis/utils"
)

type PiaohuaController struct {
	beego.Controller
}

var piaohuaBaseUrl string = "http://www.piaohua.com"
var piaohuaClient = &http.Client{}

func piaohuaHttpClient(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", piaohuaBaseUrl, nil)
	beego.Debug(req.Header.Get("Cookie"))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("Referer", piaohuaBaseUrl)
	req.Header.Add("Cookie", "cf_clearance=7df8b9cebe9d133d3b8d3a71b8575d897855c7de-1498455381-3600; __cfduid=d0864bf384208cfca5ec08dbd4b3e270a1498183457")
	res, err := piaohuaClient.Do(req)
	return res, err
}

func (c *PiaohuaController) Index() {
	result := models.Result{Code:200}
	detail := make(map[string]interface{})

	res, err := piaohuaHttpClient(piaohuaBaseUrl)

	if err != nil {
		result.Code = 201
		result.Desc = err.Error()
	} else {
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromResponse(res)
		if err != nil {
			result.Code = 202
			result.Desc = err.Error()
		} else {
			//今日热门电影推荐
			li1 := doc.Find("#iml1>ul").First().Find("li")
			movies := make([]map[string]interface{}, li1.Size())
			li1.Each(func(i int, s *goquery.Selection) {
				movie := make(map[string]interface{})

				image, _ := s.Find("img").First().Attr("src")
				movie["image"] = image

				name := s.Find("strong>font>font").Text()
				movie["name"] = name

				s.Find("strong>font>font").Remove()
				movieType := s.Find("strong>font").Text()
				movie["type"] = movieType

				date := s.Find("span").Text()
				movie["date"] = date

				url, _ := s.Find("a").First().Attr("href")
				movie["url"] = piaohuaBaseUrl + url

				movies[i] = movie
			})
			detail["movies"] = &movies

			//今日热门电视剧推荐
			li2 := doc.Find("#iml1>ul").Last().Find("li")
			tvs := make([]map[string]interface{}, li2.Size())
			li2.Each(func(i int, s *goquery.Selection) {
				tv := make(map[string]interface{})

				image, _ := s.Find("img").First().Attr("src")
				tv["image"] = image

				name := s.Find("strong>font").Text()
				tv["name"] = name

				date := s.Find("span").Text()
				tv["date"] = date

				url, _ := s.Find("a").First().Attr("href")
				tv["url"] = piaohuaBaseUrl + url

				tvs[i] = tv
			})
			detail["tvs"] = &tvs
		}
	}

	result.Detail = &detail
	c.Data["json"] = &result
	c.ServeJSON()
}

//movie detail
func (c *PiaohuaController) Detail() {
	result := models.Result{Code:200}
	detail := make(map[string]interface{})

	url := c.Input().Get("url")

	res, err := piaohuaHttpClient(url)

	if err != nil {
		result.Code = 201
		result.Desc = err.Error()
	} else {
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromResponse(res)
		if err != nil {
			result.Code = 202
			result.Desc = err.Error()
		} else {
			//download url
			tables := doc.Find("#showinfo>table")
			downloads := make([]string, tables.Size())
			tables.Each(func(i int, s *goquery.Selection) {
				download, _ := s.Find("a").First().Attr("href")
				downloads[i] = strings.TrimSpace(download)
			})
			detail["downloads"] = &downloads

			//intro
			doc.Find("#showinfo .play-list-box").Remove()
			doc.Find("#showinfo strong").Remove()
			doc.Find("#showinfo table").Remove()
			if doc.Find("#showinfo img").Size() > 1 {
				doc.Find("#showinfo img").Last().Remove()
			}
			intro, _ := doc.Find("#showinfo").Html()
			detail["intro"] = strings.TrimSpace(intro)
		}
	}

	result.Detail = &detail
	c.Data["json"] = &result
	c.ServeJSON()
}

// movies by type
func (c *PiaohuaController) Movie() {
	result := models.Result{Code:200}
	detail := make(map[string]interface{})

	movieType := c.Input().Get("type")
	pageNum := c.Input().Get("pageNum")

	url := piaohuaBaseUrl + "/html/" + movieType + "/list_" + pageNum + ".html"

	if len(pageNum) == 0 {
		pageNum = "1"
		url = piaohuaBaseUrl + "/html/" + movieType + "/index.html"
	} else if pageNum == "1" {
		url = piaohuaBaseUrl + "/html/" + movieType + "/index.html"
	} else {
		url = piaohuaBaseUrl + "/html/" + movieType + "/list_" + pageNum + ".html"
	}

	res, err := piaohuaHttpClient(url)

	if err != nil {
		result.Code = 201
		result.Desc = err.Error()
	} else {
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromResponse(res)
		if err != nil {
			result.Code = 202
			result.Desc = err.Error()
		} else {
			//ranking
			lis := doc.Find("#nmr").Find("li")
			ranking := make([]map[string]interface{}, lis.Size())
			lis.Each(func(i int, s *goquery.Selection) {
				movie := make(map[string]interface{})

				name := s.Find("a>font>font").Text()
				movie["name"] = name

				url, _ := s.Find("a").Attr("href")
				movie["url"] = piaohuaBaseUrl + url

				s.Find("a>font>font").Remove()
				_type := s.Find("a>font").Text()
				movie["type"] = _type

				ranking[i] = movie
			})
			detail["ranking"] = &ranking

			//page
			dls := doc.Find("#list").Find("dl")
			list := make([]map[string]interface{}, dls.Size())
			dls.Each(func(i int, s *goquery.Selection) {
				movie := make(map[string]interface{})

				image, _ := s.Find("dt>a>img").Attr("src")
				movie["image"] = image

				name := s.Find("dd>strong>a").Children().Text()
				movie["name"] = name

				s.Find("dd>strong>a").Children().Remove()
				_type := s.Find("dd>strong>a").Text()
				movie["type"] = _type

				other := s.Find("dd>p").Text()
				movie["other"] = other

				s.Find("dd>span>a").Remove()
				date := s.Find("dd>span").Text()
				movie["date"] = strings.TrimSpace(strings.Replace(date, "更新时间：", "", -1))

				list[i] = movie
			})
			totalCount := doc.Find(".page").First().Find("strong").Last().Text()
			page := models.Paginate(utils.ConvStringToInt(pageNum), 14, utils.ConvStringToInt(totalCount), list)
			detail["page"] = &page
		}
	}

	result.Detail = &detail
	c.Data["json"] = &result
	c.ServeJSON()
}

//lastest movies
func (c *PiaohuaController) Lastest() {
	result := models.Result{Code:200}
	detail := make(map[string]interface{})

	url := "http://www.piaohua.com/html/zuixindianying.html"

	res, err := piaohuaHttpClient(url)

	if err != nil {
		result.Code = 201
		result.Desc = err.Error()
	} else {
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromResponse(res)
		if err != nil {
			result.Code = 202
			result.Desc = err.Error()
		} else {
			doc.Find("#im").Find(".bunews").Each(func(i int, s *goquery.Selection) {
				lis := s.Find("li")
				movies := make([]map[string]interface{}, lis.Size())
				lis.Each(func(ii int, ss *goquery.Selection) {
					movie := make(map[string]interface{})

					name := ss.Find("a").Text()
					movie["name"] = name

					url, _ := ss.Find("a").Attr("href")
					movie["url"] = piaohuaBaseUrl + url

					date := ss.Find("span").Text()
					movie["date"] = date

					movies[ii] = movie
				})

				href, _ := s.Find(".title>a").Attr("href")
				href = strings.Replace(href, "/html/", "", -1)
				href = strings.Replace(href, "/index.html", "", -1)
				detail[href] = &movies
			})
		}
	}

	result.Detail = &detail
	c.Data["json"] = &result
	c.ServeJSON()
}