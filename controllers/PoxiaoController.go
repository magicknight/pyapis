package controllers

import (
	"github.com/astaxie/beego"
	"pyapis/models"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"net/http"
	"pyapis/utils"
)

type PoxiaoController struct {
	beego.Controller
}

var poxiaoBaseUrl = "http://www.poxiao.com"

func (c *PoxiaoController) Index() {
	result := models.Result{Code: 200}
	detail := make(map[string]interface{})

	client := &http.Client{}
	req, err := http.NewRequest("GET", poxiaoBaseUrl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("Referer", poxiaoBaseUrl)
	res, err := client.Do(req)

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
			// lastest update movies
			lastest_update_movies := make([]interface{}, doc.Find("#indextopleft li").Size())
			doc.Find("#indextopleft li").Each(func(i int, s *goquery.Selection) {
				movieType := utils.ConvGB2312ToUTF8(strings.TrimSpace(s.Find("span").First().Text()))
				movieDate := strings.TrimSpace(s.Find("span").Last().Text())
				movieName := utils.ConvGB2312ToUTF8(strings.TrimSpace(s.Find("a").Last().Text()))
				movieUrl, _ := s.Find("a").Last().Attr("href")

				poxiaoIndex := make(map[string]interface{})
				poxiaoIndex["MovieName"] = movieName
				poxiaoIndex["MovieUrl"] = poxiaoBaseUrl + movieUrl
				poxiaoIndex["MovieType"] = movieType
				poxiaoIndex["MovieDate"] = movieDate

				lastest_update_movies[i] = poxiaoIndex
			})
			detail["lastest_update_movies"] = &lastest_update_movies

			// five star movies
			five_star_movies := make([]interface{}, doc.Find(".dlx-content li").Size())
			doc.Find(".dlx-content li").Each(func(i int, s *goquery.Selection) {
				s.Find("em").Remove()
				movieName := utils.ConvGB2312ToUTF8(s.Find("a").Text())
				movieUrl, _ := s.Find("a").Attr("href")
				movieImage, _ := s.Find("img").Attr("src")

				poxiaoFiveStarMovie := make(map[string]interface{})
				poxiaoFiveStarMovie["movieName"] = movieName
				poxiaoFiveStarMovie["movieImage"] = movieImage
				poxiaoFiveStarMovie["movieUrl"] = poxiaoBaseUrl + movieUrl

				five_star_movies[i] = poxiaoFiveStarMovie
			})
			detail["five_star_movies"] = &five_star_movies
		}
	}

	result.Detail = &detail
	c.Data["json"] = &result
	c.ServeJSON()
}

// movie detail
func (c *PoxiaoController) Detail() {
	result := models.Result{Code: 200}
	detail := make(map[string]interface{})

	movieDetailUrl := c.Input().Get("url")

	client := &http.Client{}
	req, err := http.NewRequest("GET", movieDetailUrl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("Referer", poxiaoBaseUrl)
	res, err := client.Do(req)

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
			movieNameSub := utils.ConvGB2312ToUTF8(doc.Find("#film h1 em").Text())

			doc.Find("#film h1 em").Remove()
			movieName := utils.ConvGB2312ToUTF8(doc.Find("#film h1").First().Text())

			detailPic, _ := doc.Find(".detail_pic1 img").Attr("src")

			doc.Find("#starlist").Parent().Parent().Remove()
			detailIntro := make(map[string]interface{})
			doc.Find(".detail_intro tr").Each(func(i int, s *goquery.Selection) {
				firstTd := utils.ConvGB2312ToUTF8(s.Find("td").First().Text())
				if strings.Contains(firstTd, "导演") {
					detailIntro["director"] = utils.ConvGB2312ToUTF8(s.Find("td").Last().Text())
				} else if strings.Contains(firstTd, "主演") {
					players := make([]string, s.Find("td").Last().Find("a").Size())
					s.Find("td").Last().Find("a").Each(func(ii int, ss *goquery.Selection) {
						players[ii] = utils.ConvGB2312ToUTF8(ss.Text())
					})
					detailIntro["players"] = players
				} else if strings.Contains(firstTd, "国家") {
					detailIntro["country"] = utils.ConvGB2312ToUTF8(s.Find("td").Last().Text())
				} else if strings.Contains(firstTd, "类型") {
					detailIntro["type"] = utils.ConvGB2312ToUTF8(s.Find("td").Last().Text())
				} else if strings.Contains(firstTd, "上映日期") {
					detailIntro["playDate"] = utils.ConvGB2312ToUTF8(s.Find("td").Last().Text())
				} else if strings.Contains(firstTd, "版本") {
					detailIntro["version"] = utils.ConvGB2312ToUTF8(s.Find("td").Last().Text())
				} else if strings.Contains(firstTd, "对白语言") {
					detailIntro["language"] = utils.ConvGB2312ToUTF8(s.Find("td").Last().Text())
				}
			})

			doc.Find(".filmcontents .filmc-title").Remove()
			filmContents := doc.Find(".filmcontents").Text()
			filmContents = utils.ConvGB2312ToUTF8(filmContents)

			downloadUrls := make([]interface{}, doc.Find("input[name='checkbox2']").Size())
			doc.Find("input[name='checkbox2']").Each(func(i int, s *goquery.Selection) {
				downloadUrl := make(map[string]interface{})
				downloadUrl["name"] = utils.ConvGB2312ToUTF8(s.Parent().Text())
				url, _ := s.Attr("value")
				url = utils.ConvGB2312ToUTF8(url)
				downloadUrl["url"] = strings.Replace(url, "xzurl=", "", -1)

				downloadUrls[i] = downloadUrl
			})

			detail["movieName"] = movieName
			detail["movieNameSub"] = movieNameSub
			detail["detailPic"] = detailPic
			detail["detailIntro"] = detailIntro
			detail["filmContents"] = filmContents
			detail["downloadUrls"] = downloadUrls
			result.Detail = detail
		}
	}

	result.Detail = &detail
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *PoxiaoController) Movie() {
	result := models.Result{Code: 200}
	detail := make(map[string]interface{})

	url := "http://www.poxiao.com/type/movie/";

	pageNum := c.Input().Get("pageNum")
	if len(pageNum) == 0 {
		pageNum = "0"
	} else if pageNum == "1" {
		pageNum = "1"
	} else {
		url += "index_" + pageNum + ".html"
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Add("Referer", poxiaoBaseUrl)
	res, err := client.Do(req)

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
			sDoc := doc.Find(".yp-list-box .content ul:first-child>li")
			movies := make([]map[string]interface{}, sDoc.Size())
			sDoc.Each(func(i int, s *goquery.Selection) {
				movie := make(map[string]interface{})

				imgUrl, _ := s.Find(".gkpic img").Attr("src")
				movie["image"] = imgUrl

				movieUrl, _ := s.Find("h3 a").Attr("href")
				movie["url"] = poxiaoBaseUrl + movieUrl

				movie["name"] = utils.ConvGB2312ToUTF8(s.Find("h3 a").Text())

				movie["score"] = s.Find("h3 a em").Text()

				others := make([]string, s.Find(".jjzl li").Size())
				s.Find(".jjzl li").Each(func(ii int, ss *goquery.Selection) {
					others[ii] = utils.ConvGB2312ToUTF8(ss.Text())
				})
				movie["others"] = others

				movies[i] = movie
			})
			detail["movies"] = &movies

			totalCount := strings.TrimSpace(doc.Find(".list-pager>a").First().Text())
			beego.Debug(totalCount)
			page:=models.Paginate(utils.ConvStringToInt(pageNum), 20, utils.ConvStringToInt(totalCount), &detail)

			result.Detail = &page
		}
	}

	c.Data["json"] = &result
	c.ServeJSON()
}
