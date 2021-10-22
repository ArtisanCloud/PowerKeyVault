package controllers

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/http/request/app"
	"github.com/ArtisanCloud/PowerKeyVault/app/models"
	. "github.com/ArtisanCloud/PowerKeyVault/app/services"
	"github.com/ArtisanCloud/PowerKeyVault/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppAPIController struct {
	*APIController
	ServiceApp *AppService
}

func NewAppAPIController(c *gin.Context) (ctl *AppAPIController) {
	return &AppAPIController{
		ServiceApp: NewAppService(c),
	}
}

// app 列表
func APIAPPGetList(ctx *gin.Context) {
	ctl := NewAppAPIController(ctx)
	contextParams, _ := ctx.Get("appListParams")
	params := contextParams.(app.List)

	r, err := ctl.ServiceApp.Index(params.Page, params.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    500,
			"message": "创建数据失败" + err.Error(),
			"data":    nil,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    201,
		"message": "获取app列表成功",
		"data":    r,
	})

	return

}

// 新增app
func APIAPPUpsert(context *gin.Context) {
	ctl := NewAppAPIController(context)
	params, _ := context.Get("app")
	app := params.(*models.App)
	err := ctl.ServiceApp.Upsert(models.UNIQUE_ID,[]*models.App{app})

	if err != nil {
		ctl.RS.SetCode(config.API_ERR_CODE_FAIL_TO_UPSERT_APP, config.API_RETURN_CODE_ERROR, "", err.Error())
		panic(ctl.RS)
		return
	}

	ctl.RS.Success(context,nil )

	return
}


// 删除app
func APIAPPDelete(context *gin.Context) {
	ctl := NewAppAPIController(context)
	contextParams, _ := context.Get("appDeleteParams")
	params := contextParams.(app.Delete)

	r, err := ctl.ServiceApp.Delete(params.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    500,
			"message": "删除app失败" + err.Error(),
			"data":    r,
		})
	}

	context.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "删除app成功",
		"data":    r,
	})

	return
}
