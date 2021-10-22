package app

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/http"
	"github.com/ArtisanCloud/PowerKeyVault/app/models"
	"github.com/ArtisanCloud/PowerKeyVault/config"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
)

type RequestCreate struct {
	Name   string `json:"name" form:"name" xml:"name" binding:"required"`
	AppID  string `json:"appID" form:"name" xml:"appID" binding:"required"`
	Secret string `json:"secret" form:"secret" xml:"secret" binding:"required"`
}

// 新增app验证参数
func ValidateRequestAppCreate(ctx *gin.Context) {
	var form RequestCreate

	if err := ctx.ShouldBind(&form); err != nil {
		if err := ctx.ShouldBindJSON(&form); err != nil {
			apiResponse := http.NewAPIResponse(ctx)
			apiResponse.SetCode(
				config.API_ERR_CODE_REQUEST_PARAM_ERROR,
				config.API_RETURN_CODE_ERROR,
				"", "").SetData(object.HashMap{
				"message": err.Error(),
			}).ThrowJSONResponse(ctx)
		}
	}
	app := convertParaToAppForUpsert(form)
	ctx.Set("app", app)
	ctx.Next()

}

// request struct

// 对模型初始化
func convertParaToAppForUpsert(form *RequestCreate) (App *models.App) {

	var uuid string = ""
	if form.UUID != "" {
		uuid = form.UUID
	}
	//fmt.Dump(form)

	startDate := carbon.ParseByFormat(form.StartDate, carbon2.DATETIME_FORMAT).Carbon2Time()
	endDate := carbon.ParseByFormat(form.EndDate, carbon2.DATETIME_FORMAT).Carbon2Time()

	App = &models.App{
		MyModel: &models.MyModel{
			UUID: uuid,
		},
		AccountUUID:    form.AccountUUID,
		StartDate:      startDate,
		EndDate:        endDate,
		ExtendPeriod:   form.ExtendPeriod,
		MainAppUUID:    form.MainAppUUID,
		MainTotalClass: form.TotalClasses,
		MainUsedClass:  0,
		TotalClasses:   form.TotalClasses,
		Name:           form.Name,
		OrderUUID:      form.OrderUUID,
		OrderItemUUID:  form.OrderItemUUID,
		Plan:           form.Plan,
		ProductUUID:    form.ProductUUID,
		Status:         form.Status,
	}

	return App
}
