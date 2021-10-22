package app

import "github.com/gin-gonic/gin"

type List struct {
	Page     int `json:"page" form:"page" xml:"page"`
	PageSize int `json:"page_size" form:"page_size" xml:"page_size"`
}

// 列表接口参数初始化默认值
// 未做查询参数
func ValidateRequestAppList(ctx *gin.Context) {
	var params List
	if ctx.Param("page") == "" {
		params.Page = 1
	}

	if ctx.Param("page_size") == "" {
		params.PageSize = 10
	}

	ctx.Set("appListParams", params)
	ctx.Next()
}
