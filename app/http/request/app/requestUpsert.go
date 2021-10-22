package app

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/http"
	"github.com/ArtisanCloud/PowerKeyVault/app/models"
	"github.com/ArtisanCloud/PowerKeyVault/config"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/gin-gonic/gin"
)

type RequestUpsert struct {
	UUID string `form:"uuid" json:"uuid" xml:"uuid"`

	Name   string `json:"name" form:"name" xml:"name" binding:"required"`
	AppID  string `json:"appID" form:"appID" xml:"appID" binding:"required"`
	Secret string `json:"secret" form:"secret" xml:"secret" binding:"required"`
}

// 新增app验证参数
func ValidateRequestAppUpsert(ctx *gin.Context) {
	var form RequestUpsert

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
	app := convertParaToAppForUpsert(&form)
	ctx.Set("app", app)
	ctx.Next()

}

// 对模型初始化
func convertParaToAppForUpsert(form *RequestUpsert) (App *models.App) {

	var uuid string = ""
	if form.UUID != "" {
		uuid = form.UUID
	}

	App = &models.App{
		MyModel: &models.MyModel{
			UUID: uuid,
		},
		Name: form.Name,
		AppID: form.AppID,
		Secret: form.Secret,
	}

	return App
}
