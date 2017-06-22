package models

type Page struct {
	PageNum    int `json:"pageNum"`
	PageSize   int `json:"pageSize"`
	TotalPage  int `json:"totalPage"`
	TotalCount int `json:"totalCount"`
	FirstPage  bool `json:"firstPage"`
	LastPage   bool `json:"lastPage"`
	List       interface{} `json:"list"`
}

func Paginate(pageNum int, pageSize int, totalCount int, list interface{}) Page {
	page := Page{PageNum: pageNum, PageSize: pageSize, TotalCount: totalCount, List: list}

	if totalCount%pageSize > 0 {
		page.TotalPage = (totalCount / pageSize) + 1
	} else {
		page.TotalPage = totalCount / pageSize
	}

	if pageNum == 1 {
		page.FirstPage = true
	}

	if page.TotalPage == pageNum {
		page.LastPage = true
	}

	return page;
}
