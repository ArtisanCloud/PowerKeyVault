package lang

import (
	"github.com/ArtisanCloud/PowerKeyVault/resources/lang/en_US"
	"github.com/ArtisanCloud/PowerKeyVault/resources/lang/zh_TW"
)

func LoadLanguages()  {
	en_US.LoadLang()
	en_TW.LoadLang()
}