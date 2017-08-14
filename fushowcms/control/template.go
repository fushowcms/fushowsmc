package control

import (
	"fmt"
	"fushowcms/comm"
	"github.com/gin-gonic/gin"
	m "fushowcms/models"
	"strconv"
	"strings"
)

//模板输出
func GetIndexPage(c *gin.Context) {
	var (
		ar            m.AnchorRoom
		s_list_normal []m.AnchorInfo
	)
	//精彩推荐8个
	stick, _, _ := ar.GetHotLiveList()
	var stickId []string
	//去重
	for i := 0; i < len(stick); i++ {
		stickId = append(stickId, strconv.FormatInt(stick[i].Id, 10))
	}
	if len(stick) < 8 {
		//不足4个房间时
		roomnum := 8 - len(stick)
		_, s_list_normal = ar.GetRoomStickIng(roomnum, stickId)
	}

	//二级目录
	var (
		ct  m.CategoryTwo
		ars m.AnchorInfo
		ccc []TwoArCategory
	)
	//查询所有二级类目
	twoList := ct.GetTwoCategoryList()
	//遍历二级类目，找到对应的房间
	for j := 0; j < len(twoList); j++ {
		var tac TwoArCategory
		//通过二级类目找到对应的房间
		data, _, total := ars.GetRoomAliasByRoomTypes(strconv.FormatInt(twoList[j].Id, 10))
		if len(data) > 0 {
			tac.Id = twoList[j].Id
			tac.Total = total
			tac.TwoCategoryName = twoList[j].TwoCategoryName
			tac.TwoCategoryImage = twoList[j].TwoCategoryImage
			tac.TwoCategoryAddress = twoList[j].TwoCategoryAddress
			tac.ArCategoryTwo = data
			ccc = append(ccc, tac)
		}
	}
	var cp m.CarouselPic

	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))

	list, _ := cp.GetCarPicList(page, rows)

	c.HTML(200, "page/index.html", gin.H{
		"data_stick":       stick,
		"data_stick_state": s_list_normal,
		"Categorydata":     ccc,
		"CarouselPic":      list,
	})
}
func GetListsPage(c *gin.Context) {
	fmt.Println("start")
	//查询类目
	ccc := GetList()
	//查询房间列表
	ing_list, ing_list_normal, stick, s_list_normal, total := GetAlllist(1, 15)

	c.HTML(200, "page/lists.html", gin.H{
		"CategoryList": ccc,

		"data_normal":       ing_list,
		"data_normal_state": ing_list_normal,
		"data_stick":        stick,
		"data_stick_state":  s_list_normal,
		"total":             total,
	})
}

func GetRoomLivePage(c *gin.Context) {
	fmt.Println("start")
	var (
		ccc []Category
		co  m.CategoryOne
		cot m.CategoryOneTwo
	)
	//查询所有一级类目
	oneList := co.GetCategoryLists()
	for j := 0; j < len(oneList); j++ {
		//通过一级类目取二级类目
		twoList := cot.GetOneTwoCategoryByOneId(oneList[j].Id)
		//判断没有二级类目时不存在时
		if len(twoList) != 0 {
			var now Category
			for k := 0; k < len(twoList); k++ {
				//通过twoid查询
				var ctt m.CategoryTwo
				cttList := ctt.GetTwoCategoryById(twoList[k].TwoId)
				now.ArrCategory = append(now.ArrCategory, cttList[0])
			}
			now.COneId = oneList[j].Id
			now.COneName = oneList[j].OneCategoryName
			ccc = append(ccc, now)
		}
	}
	var (
		arc m.AnchorRoomConcern
	)
	url := c.Request.RequestURI
	str := strings.LastIndex(url, "/")
	res := Substr(url, str+1, len(url)-str)
	roomId, _ := strconv.ParseInt(res, 10, 64)

	var rm m.AnchorRoom
	rm.Id = roomId
	if !rm.GetRoom() {
		fmt.Println("房间不存在")
	}
	cookie, err := c.Request.Cookie(comm.InitConf.GetValue("SESSION", "cookie_name"))
	fmt.Println(cookie)
	if err == nil {
		cookievalue := cookie.Value
		arc.Uid, _ = strconv.ParseInt(comm.GetDecValue(cookievalue, "fushow.cms"), 10, 64)
	}
	arc.User = rm.Uid
	isConcern := arc.IsConcern()
	if arc.IsConcern() {
		fmt.Println("没有关注")
	}

	var ui m.UidInfoJson
	//获取客户端参数
	id := arc.Uid
	ui.Id = id

	// 根据id 获取用户信息
	getUserInfo := ui.GetUserInfo()
	if ui.GetUserInfo() {
		ui.PassWord = ""
		ui.UID = comm.SetAesValue(strconv.FormatInt(ui.Id, 10), "fushow.cms")
	}
	//根据房间id和用户id判断该用户是否是房管
	Fball := FindByRoomIdUserIdAll_com(roomId, arc.Uid)
	//查询房间内禁言人员
	Fnall := FindNotUserSpeakAll_com(arc.Uid, roomId)
	//直播流
	Giflow := GetInFlow_com(rm.Uid, "0")
	//礼物列表
	Grows, Gtotal := GetGiftList_com(0, 0, "", "", "")
	//总榜
	_, Gggm := GetGiftGiveMonths_com(rm.Uid)
	//周榜
	_, Gggw := GetGiftGiveWeeks_com(rm.Uid)
	//获取房间单个
	list, count := GetRoom_com(rm.Uid)

	c.HTML(200, "page/roomlive.html", gin.H{
		"CategoryList": ccc,
		"state":        list,
		"attention":    count,
		"anchor":       rm.Uid,
		"IsConcern":    isConcern,
		"GetUserInfo":  getUserInfo,
		"Ui":           ui,
		"Fball":        Fball,
		"Fnall":        Fnall,
		"Giflow":       Giflow,
		"Grows":        Grows,
		"Gtotal":       Gtotal,
		"Gggm":         Gggm,
		"Gggw":         Gggw,
	})
}
func GetListsClassPage(c *gin.Context) {
	fmt.Println("start")
	var (
		ccc []Category
		co  m.CategoryOne
		cot m.CategoryOneTwo
	)
	//查询所有一级类目
	oneList := co.GetCategoryLists()
	fmt.Println("oneList", oneList)
	if len(oneList) <= 0 {
		fmt.Println("一级类目不存在")
	}
	for j := 0; j < len(oneList); j++ {
		//通过一级类目取二级类目
		twoList := cot.GetOneTwoCategoryByOneId(oneList[j].Id)
		//判断没有二级类目时不存在时
		if len(twoList) != 0 {
			var now Category
			for k := 0; k < len(twoList); k++ {
				//通过twoid查询
				var ctt m.CategoryTwo
				cttList := ctt.GetTwoCategoryById(twoList[k].TwoId)
				now.ArrCategory = append(now.ArrCategory, cttList[0])
			}
			now.COneId = oneList[j].Id
			now.COneName = oneList[j].OneCategoryName
			ccc = append(ccc, now)
		}
	}
	c.HTML(200, "page/listsClass.html", gin.H{
		"CategoryList": ccc,
	})
}
func GetCatePage(c *gin.Context) {
	var (
		ccc []Category
		co  m.CategoryOne
		cot m.CategoryOneTwo
	)
	//查询所有一级类目
	oneList := co.GetCategoryLists()
	fmt.Println("oneList", oneList)
	if len(oneList) <= 0 {
		fmt.Println("一级类目不存在")
	}
	for j := 0; j < len(oneList); j++ {
		//通过一级类目取二级类目
		twoList := cot.GetOneTwoCategoryByOneId(oneList[j].Id)
		//判断没有二级类目时不存在时
		if len(twoList) != 0 {
			var now Category
			for k := 0; k < len(twoList); k++ {
				//通过twoid查询
				var ctt m.CategoryTwo
				cttList := ctt.GetTwoCategoryById(twoList[k].TwoId)
				now.ArrCategory = append(now.ArrCategory, cttList[0])
			}
			now.COneId = oneList[j].Id
			now.COneName = oneList[j].OneCategoryName
			ccc = append(ccc, now)
		}
	}
	c.HTML(200, "page/cate.html", gin.H{
		"CategoryList": ccc,
	})
}

func GetAlltvLivePage(c *gin.Context) {
	var (
		ccc []Category
		co  m.CategoryOne
		cot m.CategoryOneTwo
	)
	//查询所有一级类目
	oneList := co.GetCategoryLists()
	fmt.Println("oneList", oneList)
	if len(oneList) <= 0 {
		fmt.Println("一级类目不存在")
	}
	for j := 0; j < len(oneList); j++ {
		//通过一级类目取二级类目
		twoList := cot.GetOneTwoCategoryByOneId(oneList[j].Id)
		//判断没有二级类目时不存在时
		if len(twoList) != 0 {
			var now Category
			for k := 0; k < len(twoList); k++ {
				//通过twoid查询
				var ctt m.CategoryTwo
				cttList := ctt.GetTwoCategoryById(twoList[k].TwoId)
				now.ArrCategory = append(now.ArrCategory, cttList[0])
			}
			now.COneId = oneList[j].Id
			now.COneName = oneList[j].OneCategoryName
			ccc = append(ccc, now)
		}
	}
	c.HTML(200, "page/alltvlive.html", gin.H{
		"CategoryList": ccc,
	})
}

func GetOutLivePage(c *gin.Context) {
	var (
		ccc []Category
		co  m.CategoryOne
		//	ct  m.CategoryTwo
		cot m.CategoryOneTwo
	)
	//查询所有一级类目
	oneList := co.GetCategoryLists()
	fmt.Println("oneList", oneList)
	if len(oneList) <= 0 {
		fmt.Println("一级目录不存在")
	}
	for j := 0; j < len(oneList); j++ {
		//通过一级类目取二级类目
		twoList := cot.GetOneTwoCategoryByOneId(oneList[j].Id)
		//判断没有二级类目时不存在时
		if len(twoList) != 0 {
			var now Category
			for k := 0; k < len(twoList); k++ {
				//通过twoid查询
				var ctt m.CategoryTwo
				cttList := ctt.GetTwoCategoryById(twoList[k].TwoId)
				now.ArrCategory = append(now.ArrCategory, cttList[0])
			}
			now.COneId = oneList[j].Id
			now.COneName = oneList[j].OneCategoryName
			ccc = append(ccc, now)
		}
	}
	fmt.Println("ccc", ccc)
	c.HTML(200, "page/outlive.html", gin.H{
		"CategoryList": ccc,
	})
}
