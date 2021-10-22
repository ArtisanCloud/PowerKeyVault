package app

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/http"
	"github.com/ArtisanCloud/PowerKeyVault/config"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/gin-gonic/gin"
)

type Update struct {
	Delete
	Create
}

// 修改app 验证参数
func ValidateRequestAppUpdate(ctx *gin.Context) {
	var form Update
	if err := ctx.ShouldBind(&form); err != nil {
		if err := ctx.ShouldBindJSON(&form); err != nil {
			apiResponse := &http.APIResponse{}
			apiResponse.Context = ctx
			apiResponse.SetCode(
				config.API_ERR_CODE_REQUEST_PARAM_ERROR,
				config.API_RETURN_CODE_ERROR,
				"参数缺失", "参数缺失").SetData(object.HashMap{
				"message": err.Error(),
			}).ThrowJSONResponse(ctx)
		}
	}
	ctx.Set("appUpdateParams", form)
	ctx.Next()
}
