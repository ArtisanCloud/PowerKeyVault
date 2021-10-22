package app

import (
	"github.com/ArtisanCloud/PowerKeyVault/app/http"
	"github.com/ArtisanCloud/PowerKeyVault/config"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/gin-gonic/gin"
)

type Delete struct {
	ID int `json:"id" form:"id" xml:"id" binding:"required"`
}

// 删除app验证参数 id
func ValidateRequestAppDelete(ctx *gin.Context) {
	var form Delete
	if err := ctx.ShouldBind(&form); err != nil {
		if err := ctx.ShouldBindJSON(&form); err != nil {
			apiResponse := &http.APIResponse{}
			apiResponse.Context = ctx
			apiResponse.SetCode(
				config.API_ERR_CODE_REQUEST_PARAM_ERROR,
				config.API_RETURN_CODE_ERROR,
				"参数缺失 id", "参数缺失 id").SetData(object.HashMap{
				"message": err.Error(),
			}).ThrowJSONResponse(ctx)
		}
	}
	ctx.Set("appDeleteParams", form)
	ctx.Next()
}
