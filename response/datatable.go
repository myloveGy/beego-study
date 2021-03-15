package response

// 表格返回数据
type DataTable struct {
	Draw            int64       `json:"draw"`
	RecordsFiltered int64       `json:"recordsFiltered"`
	RecordsTotal    int64       `json:"recordsTotal"`
	Data            interface{} `json:"data"`
}

type PageData struct {
	Total        int64       `json:"iTotal"`
	TotalRecords int         `json:"iTotalRecords"`
	Data         interface{} `json:"aData"`
}
