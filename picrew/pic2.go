package picrew

import (
	"math/rand"
	"time"
)

var Pic2 = SetCf(`47161`)

func ToPic2() string {
	rand.Seed(time.Now().UnixNano())
	Main := make(ImsMap)
	var PNm = []string{
		"顔",
		"眉毛",
		"右目",
		"左目",
		"口",
		"洋服",
		"背景",
	}
	var Hair0 = []string{
		"後ろ髪",
		"前髪",
	}
	var Hair1 = []string{
		"後ろ髪2",
		"横髪",
	}
	var Hair2 = []string{
		"あほ毛",
		// "前髪あほ毛",
	}

	var PNm2 = []string{
		"右目ハイライト",
		"左目ハイライト",
		"顔装飾",
		"ほくろ",
		"首装飾",
		"羽",
	}
	var PNm3 = []string{
		"顔装飾",
		"顔装飾２",
		"頭装飾",
	}
	//必须组件
	Main.SetIm(Pic2, PNm, -1, 1)
	//头发
	hcol := rand.Intn(Pic0[Hair0[0]].Cols)
	Main.SetIm(Pic2, Hair0, hcol, 1)
	Main.SetIm(Pic2, Hair1, hcol, 3)
	Main.SetIm(Pic2, Hair2, -1, 4)

	Main.SetIm(Pic2, PNm2, -1, 4)
	Main.SetIm(Pic2, PNm3, -1, 3)
	// 制图
	return Main.Save(22)
}
