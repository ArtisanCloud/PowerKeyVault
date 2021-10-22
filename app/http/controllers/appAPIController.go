package controllers

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/http/request/app"
	"github.com/ArtisanCloud/PowerKeyVault/app/models"
	. "github.com/ArtisanCloud/PowerKeyVault/app/services"
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
func APIAPPCreate(context *gin.Context) {
	ctl := NewAppAPIController(context)
	params, _ := context.Get("appCreateParams")
	para := params.(app.Create)
	ctl.ServiceApp.App = &models.App{
		Name:   para.Name,
		AppID:  para.AppID,
		Secret: para.Secret,
	}

	r, err := ctl.ServiceApp.Create()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    500,
			"message": "创建数据失败" + err.Error(),
			"data":    nil,
		})
	}

	context.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "新增app成功",
		"data":    r,
	})

	return
}

// 更新app
func APIAPPUpdate(context *gin.Context) {
	ctl := NewAppAPIController(context)
	contextParams, _ := context.Get("appUpdateParams")
	params := contextParams.(app.Update)
	ctl.ServiceApp.App = &models.App{
		Name:   params.Name,
		AppID:  params.AppID,
		Secret: params.Secret,
	}

	r, err := ctl.ServiceApp.Update(params.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    500,
			"message": "更新数app信息失败" + err.Error(),
			"data":    r,
		})
	}

	context.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "更新数app据成功",
		"data":    r,
	})

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
