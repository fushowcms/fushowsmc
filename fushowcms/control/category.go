package control

import (
	"fushowcms/comm"
	m "fushowcms/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//查询一级类目 addby liuhan
func GetCategoryList(c *gin.Context) {
	var (
		co m.CategoryOne
	)
	page, _ := strconv.Atoi(c.PostForm("page"))
	rows, _ := strconv.Atoi(c.PostForm("rows"))
	total, list := co.GetCategoryList(page, rows)
	if len(list) <= 0 {
		c.JSON(200, gin.H{"total": total, "rows": []int{}})
		return
	}
	c.JSON(200, gin.H{"total": total, "rows": list})
}

//增加一级类目 addby liuhan
func CategoryAdd(c *gin.Context) {
	var (
		co m.CategoryOne
	)
	if c.PostForm("CategoryName") == "" {
		comm.ResponseError(c, 3200) //一级类目不能为空
		return
	}
	co.OneCategoryName = c.PostForm("CategoryName")
	//判断一级类目名称是否已存在
	list := co.GetOneCategoryByName(co.OneCategoryName)
	if len(list) > 0 {
		comm.ResponseError(c, 3211) //一级类目增加失败，名称已存在
		return
	}
	if !co.CategoryAdd() {
		comm.ResponseError(c, 3201) //一级类目增加失败
		return
	}
	m := make(map[string]interface{})
	m["state"] = true
	m["message"] = "添加成功"
	comm.Response(c, m)
}

//删除一级类目 addby liuhan
func CategoryDelete(c *gin.Context) {
	var (
		co  m.CategoryOne
		cot m.CategoryOneTwo
	)
	co.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64)
	if co.Id == 0 {
		comm.ResponseError(c, 3202) //一级类目Id不存在
		return
	}
	var
	//判断一级类目下是否存在二级类目，如果有，不允许删除
	list = cot.GetOneTwoCategoryByOneId(co.Id)
	if len(list) > 0 {
		comm.ResponseError(c, 3205) //一级类目下有二级类目存在，不允许删除
		return
	}
	if !co.CategoryDelete() {
		comm.ResponseError(c, 3203) //一级类目删除失败
		return
	}
	m := make(map[string]interface{})
	m["state"] = true
	m["message"] = "删除成功"
	comm.Response(c, m)
}

//根据一级类目id，查询对应的所有二级类目 addby liuhan
func GetTwoCategoryByOneId(c *gin.Context) {
	var (
		ct m.CategoryTwo
	)
	id, _ := strconv.ParseInt(c.PostForm("Id"), 10, 64)
	list := ct.GetTwoCategoryByOneId(id)
	if len(list) <= 0 {
		comm.ResponseError(c, 3204) //二级类目不存在
		return
	}
	comm.Response(c, list)
}

//查询所有二级类目 addby liuhan
func GetTwoCategoryList(c *gin.Context) {
	var (
		ct m.CategoryTwo
	)
	list := ct.GetTwoCategoryList()
	comm.Response(c, list)
}

//删除二级类目 addby liuhan
func CategoryTwoDelete(c *gin.Context) {
	var (
		ct m.CategoryTwo
	)
	ct.Id, _ = strconv.ParseInt(c.PostForm("Id"), 10, 64)
	if ct.Id == 0 {
		comm.ResponseError(c, 3206) //二级类目Id不存在
		return
	}
	if !ct.CategoryTwoDelete() {
		comm.ResponseError(c, 3207) //二级类目删除失败
		return
	}
	m := make(map[string]interface{})
	m["state"] = true
	m["message"] = "删除成功"
	comm.Response(c, m)
}

//增加二级类目 addby liuhan
func CategoryTwoAdd(c *gin.Context) {
	var (
		ct m.CategoryTwo
	)
	if c.PostForm("TwoCategoryName") == "" {
		comm.ResponseError(c, 3214) //二级类目不能为空
		return
	}
	ct.TwoCategoryName = c.PostForm("TwoCategoryName")
	if c.PostForm("TwoCategoryImage") == "" {
		comm.ResponseError(c, 3215) //二级类目图片不能为空
		return
	}
	ct.TwoCategoryImage = c.PostForm("TwoCategoryImage")
	if c.PostForm("TwoCategoryAddress") == "" {
		comm.ResponseError(c, 3213) //二级类目图片不能为空
		return
	}
	ct.TwoCategoryAddress = c.PostForm("TwoCategoryAddress")
	//判断二级类目名称时否存在
	list := ct.GetTwoCategoryByName(ct.TwoCategoryName)
	if len(list) > 0 {
		comm.ResponseError(c, 3209) //名称已存在
		return
	}
	if !ct.CategoryTwoAdd() {
		comm.ResponseError(c, 3210) //二级类目增加失败
		return
	}
	m := make(map[string]interface{})
	m["state"] = true
	m["message"] = "添加成功"
	comm.Response(c, m)
}

//增加二级类目 addby liuhan
func CategoryTwoAdds(c *gin.Context) {
	var (
		cot m.CategoryOneTwo
	)
	oneId, _ := strconv.ParseInt(c.PostForm("OneId"), 10, 64)
	twoIdList := c.PostForm("TwoIdList")
	//分割所有二级类目id
	arr := strings.Split(twoIdList, ",")
	//删除所有子类关联
	if !cot.CategoryAllByOneId(oneId) {
		comm.ResponseError(c, 3212) //增加失败
		return
	}
	//添加关联
	for j := 0; j < len(arr)-1; j++ {
		var (
			cott m.CategoryOneTwo
		)
		if arr[j] != "" {
			twoId, _ := strconv.ParseInt(arr[j], 10, 64)
			cott.CategoryOneTwoAdd(oneId, twoId)
		}

	}
	m := make(map[string]interface{})
	m["state"] = true
	m["message"] = "添加成功"
	comm.Response(c, m)
}

//根据一级类目id，查询对应的所有二级类目 addby liuhan
func GetOneTwoCategoryByOneId(c *gin.Context) {
	var (
		cot m.CategoryOneTwo
	)
	id, _ := strconv.ParseInt(c.PostForm("Id"), 10, 64)
	list := cot.GetOneTwoCategoryByOneId(id)
	if len(list) <= 0 {
		comm.ResponseError(c, 3204) //二级类目不存在
		return
	}
	comm.Response(c, list)
}

//二级类目管理 addby liuhan
type Category struct {
	COneId      int64           //一级类目ID
	COneName    string          //一级类目name
	ArrCategory []m.CategoryTwo `xorm:"extends"`
}

// 查询所有类目 addby liuhan
func FindCategoryAll(c *gin.Context) {
	categoryList := GetList()
	comm.Response(c, categoryList)
}

func GetList() []Category {
	var (
		categoryList []Category
		co           m.CategoryOne
		cot          m.CategoryOneTwo
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
			categoryList = append(categoryList, now)
		}
	}
	return categoryList
}

//根据一级类目id，查询对应的所有二级类目 addby liuhan
func GetTwoCategoryByOneIds(c *gin.Context) {
	var (
		cot  m.CategoryOneTwo
		ctt1 []m.CategoryTwo
	)
	id, _ := strconv.ParseInt(c.PostForm("Id"), 10, 64)
	//通过一级类目取二级类目
	twoList := cot.GetOneTwoCategoryByOneId(id)
	//判断没有二级类目时不存在时
	for k := 0; k < len(twoList); k++ {
		//通过twoid查询
		var ctt m.CategoryTwo
		cttList := ctt.GetTwoCategoryById(twoList[k].TwoId)
		ctt1 = append(ctt1, cttList[0])
	}

	comm.Response(c, ctt1)
}

//根据二级类目地址查询 addby liuhan
func GetTwoCategoryByAddress(c *gin.Context) {
	var (
		ct m.CategoryTwo
	)
	ct.TwoCategoryAddress = c.PostForm("Address")
	twoList := ct.GetTwoCategoryByAddress()
	if len(twoList) == 0 {
		comm.ResponseError(c, 3216) //地址不存在
		return
	}
	comm.Response(c, twoList) //一级类目Id不存在
}

//二级类目下的房间 addby liuhan 161226
type TwoArCategory struct {
	Total              int64
	Id                 int64          //二级类目ID
	TwoCategoryName    string         //二级类目名称
	TwoCategoryImage   string         //二级类目图片
	TwoCategoryAddress string         //二级类目地址
	ArCategoryTwo      []m.AnchorInfo `xorm:"extends"`
}

//查询所有二级分类下的房间 addby liuhan 161226
func GetTwoCategoryRoom(c *gin.Context) {
	var (
		ct           m.CategoryTwo
		ar           m.AnchorInfo
		categoryList []TwoArCategory
	)
	//查询所有二级类目
	twoList := ct.GetTwoCategoryList()
	//遍历二级类目，找到对应的房间
	for j := 0; j < len(twoList); j++ {
		var tac TwoArCategory
		//通过二级类目找到对应的房间
		data, _, total := ar.GetRoomAliasByRoomTypes(strconv.FormatInt(twoList[j].Id, 10))
		if len(data) > 0 {
			tac.Id = twoList[j].Id
			tac.Total = total
			tac.TwoCategoryName = twoList[j].TwoCategoryName
			tac.TwoCategoryImage = twoList[j].TwoCategoryImage
			tac.TwoCategoryAddress = twoList[j].TwoCategoryAddress
			tac.ArCategoryTwo = data
			categoryList = append(categoryList, tac)
		}
	}
	comm.Response(c, categoryList) //一级类目Id不存在
}

//删除二级类目 addby liuhan 170124
func DeleteCategoryTwo(c *gin.Context) {
	twoIdList := c.PostForm("TwoIdList")
	arr := strings.Split(twoIdList, ",")
	//删除二级类目
	for j := 0; j < len(arr)-1; j++ {
		var (
			cott m.CategoryOneTwo
		)
		if arr[j] != "" {
			twoId, _ := strconv.ParseInt(arr[j], 10, 64)
			if cott.CategoryAllByTwoId(twoId) {
				var ct1 m.CategoryTwo
				ct1.Id = twoId
				if !ct1.CategoryTwoDelete() {
					comm.ResponseError(c, 3207) //二级类目删除失败
					return
				}
			}
		}
	}
	m := make(map[string]interface{})
	m["state"] = true
	m["message"] = "删除成功"
	comm.Response(c, m)
}
