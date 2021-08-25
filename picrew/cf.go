package picrew

import (
	"github.com/tdf1939/img"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

var YUrl = `https://ghproxy.com/https://raw.githubusercontent.com/tdf1939/ZeroBot-Plugin-picrew/main/`

//图片结构体
type Im struct {
	Lyr int    //图层
	Url string //地址
}

//素材全部图层切片
type Urls [][]Im

//部件结构体
type ItemsMap struct {
	Items []Urls //插件
	Cols  int    //颜色数
}

//项目字典
type Main map[string]ItemsMap

//选用图片
type ImsMap map[int]string

//部件素材列表 素材-->图层-->颜色-->url
func SetCf(id string) Main {
	it := make(Main)
	cf, _ := req.Get(YUrl + id + `/cf.json`)
	img, _ := req.Get(YUrl + id + `/img.json`)
	imgjson := img.String()
	cfjson := cf.String()
	pListAll := gjson.Get(cfjson, `pList`).Array() //组件
	for _, v := range pListAll {
		var li ItemsMap
		pList := v.String()
		var lyrs []string //组件图层
		lyrsAll := gjson.Get(pList, `lyrs`).Array()
		for _, v1 := range lyrsAll {
			lyrs = append(lyrs, v1.String())
		}
		cpId := gjson.Get(pList, "cpId").String()
		ColAll := gjson.Get(cfjson, "cpList."+cpId).Array()
		li.Cols = len(ColAll)
		var cIds []string //部件颜色id
		for _, v1 := range ColAll {
			cIds = append(cIds, gjson.Get(v1.String(), "cId").String())
		}
		itemsAll := gjson.Get(pList, `items`).Array()
		for _, v1 := range itemsAll {
			item := gjson.Get(v1.String(), `itmId`).String()
			urls := gjson.Get(imgjson, `lst.`+item).String()
			var k Urls
			for _, v2 := range lyrs {
				tc := int(gjson.Get(cfjson, `lyrList.`+v2).Int()) //图层数
				var t []Im                                        //素材单个图层地址切片
				if lyrurl := gjson.Get(urls, v2).String(); gjson.Get(urls, v2).Exists() {
					for _, v3 := range cIds {
						if colurl := gjson.Get(lyrurl, v3+".url").String(); gjson.Get(lyrurl, v3).Exists() {
							t = append(t, Im{Lyr: tc, Url: colurl})
						}
					}
					k = append(k, t)
				}
			}
			li.Items = append(li.Items, k)
		}
		it[gjson.Get(pList, `pNm`).String()] = li
	}
	return it
}

//选择部件对应素材, nm 部件列表, col指定颜色小于0则随机，r  1/r概率选择此部件
func (ims ImsMap) SetIm(m Main, nm []string, col int, r int) {
	rand.Seed(time.Now().UnixNano())
	for _, v := range nm {
		if rr := rand.Intn(r); rr == 0 {
			it := m[v].Items[rand.Intn(len(m[v].Items))]
			for _, v1 := range it {
				if col <= len(v1) && col >= 0 {
					ims[v1[col].Lyr] = `https://cdn.picrew.me` + v1[col].Url
				} else {
					col := rand.Intn(len(v1))
					ims[v1[col].Lyr] = `https://cdn.picrew.me` + v1[col].Url
				}
			}
		}
	}
}

//图层叠加 r 最大图层
func (m ImsMap) Save(r int) string {
	rand.Seed(time.Now().UnixNano())
	f := strconv.Itoa(rand.Int())
	path := "data/image/picrew/" + f + "/"
	os.MkdirAll(path, 0777)

	dc := img.NewDc(600, 600, color.NRGBA{0, 0, 0, 0})
	for i := 1; i <= r; i++ {
		if v, ok := m[i]; ok {
// 			req.SetProxyUrl("http://127.0.0.1:10809")   //代理，按需设置
			re, _ := req.Get(v)
			re.ToFile(path + strconv.Itoa(i) + ".png")
			dc.Over(img.ImDc(path+strconv.Itoa(i)+".png", 0, 0).Im, 0, 0, 0, 0)
			fmt.Print(f, ": 已成功", i, "张；还剩", r-i, "张！\n")
		}
	}
	img.SavePng(dc.Im, path+"pic.png")
	imgurl := img.SGpic(path + "pic.png")
	os.RemoveAll(path)
	return imgurl
}
