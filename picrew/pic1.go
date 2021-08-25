package picrew

import (
	"math/rand"
	"time"
)

var Pic1 = SetCf(`407340`)

func ToPic1() string {
	rand.Seed(int64(time.Now().Minute()))
	Main := make(ImsMap)
	var PNm = []string{
		"肌",
		"右目",
		"左目",
		"右ハイライト",
		"左ハイライト",
		"眉",
		"口",
		"下着",
	}
	var Hair0 = []string{
		"前髪",
		"横髪",
		"後ろ髪",
	}
	var Hair1 = []string{
		"インナーカラー",
		"まとめ髪",
		"下ろし髪",
	}
	var Hair2 = []string{
		"前髪メッシュ",
		"アホ毛",
	}

	var PNm2 = []string{
		"ほくろ",
		"頬",
		// "涙",
		"服",
		"上着",
	}
	var PNm3 = []string{
		"眼鏡",
		"マスク",
		"耳飾り",
		"首飾り",
		"髪飾り",
		"帽子",
		"羽",
		"背景１",
		"背景２",
	}
	//必须组件
	Main.SetIm(Pic1, PNm, "", 1)
	//头发
	hcolall := Pic1[Hair0[0]].Cols
	li := rand.Intn(len(hcolall))
	hcol := hcolall[li]
	Main.SetIm(Pic1, Hair0, hcol, 1)
	Main.SetIm(Pic1, Hair1, hcol, 3)
	Main.SetIm(Pic1, Hair2, hcol, 4)

	Main.SetIm(Pic1, PNm2, "", 3)
	Main.SetIm(Pic1, PNm3, "", 3)
	// 制图
	return Main.Save(45)
}
