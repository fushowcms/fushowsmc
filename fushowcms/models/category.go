package models

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-xorm/xorm"
)

//查询一级类目 addby liuhan
func (co *CategoryOne) GetCategoryList(page, rows int) (int64, []CategoryOne) {
	var list []CategoryOne
	total, _ := Engine.Where("id>?", 0).Count(co)
	Engine.Where("id >?", 0).Limit(rows, (page-1)*rows).Find(&list)
	return total, list
}

//查询一级类目 addby liuhan
func (co *CategoryOne) GetCategoryLists() []CategoryOne {
	var list []CategoryOne
	Engine.Where("id >?", 0).Find(&list)
	return list
}

//增加一级类目 addby liuhan
func (co *CategoryOne) CategoryAdd() bool {
	row, err := Engine.Insert(co) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//删除一级类目 addby liuhan
func (co *CategoryOne) CategoryDelete() bool {
	row, err := Engine.Where("Id=?", co.Id).Delete(co) //删除一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//根据一级类目id，查询对应的所有二级类目 addby liuhan
func (ct *CategoryTwo) GetTwoCategoryByOneId(id int64) []CategoryTwo {
	var list []CategoryTwo
	Engine.Where("pid =?", id).Find(&list)
	return list
}

//查询所有二级类目 addby liuhan
func (ct *CategoryTwo) GetTwoCategoryList() []CategoryTwo {
	var list []CategoryTwo
	Engine.Where("id >?", 0).Find(&list)
	return list
}

//删除一级类目 addby liuhan
func (ct *CategoryTwo) CategoryTwoDelete() bool {
	row, err := Engine.Where("Id=?", ct.Id).Delete(ct) //删除一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//增加二级类目 addby liuhan
func (ct *CategoryTwo) CategoryTwoAdd() bool {
	row, err := Engine.Insert(ct) //插入一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//判断二级类目名称时否存在  addby liuhan
func (ct *CategoryTwo) GetTwoCategoryByName(name string) []CategoryTwo {
	var list []CategoryTwo
	Engine.Where("two_category_name =?", name).Find(&list)
	return list
}

//判断一级类目名称是否已存在
func (co *CategoryOne) GetOneCategoryByName(name string) []CategoryOne {
	var list []CategoryOne
	Engine.Where("one_category_name =?", name).Find(&list)
	return list
}

//删除一级类目 addby liuhan
func (cot *CategoryOneTwo) CategoryAllByOneId(oneId int64) bool {
	row, err := Engine.Where("one_id=?", oneId).Delete(cot) //删除一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//根据二级类目删除关联类目 addby liuhan 071424
func (cot *CategoryOneTwo) CategoryAllByTwoId(twoId int64) bool {
	row, err := Engine.Where("two_id=?", twoId).Delete(cot) //删除一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//添加类目关联表 addby liuhan
func (cot *CategoryOneTwo) CategoryOneTwoAdd(oneId, twoId int64) bool {
	cot.OneId = oneId
	cot.TwoId = twoId
	row, err := Engine.Insert(cot) //删除一行数据
	if row > 0 || err == nil {
		return true
	}
	return false
}

//根据一级类目id，查询对应的所有二级类目 addby liuhan
func (ct *CategoryOneTwo) GetOneTwoCategoryByOneId(id int64) []CategoryOneTwo {
	var list []CategoryOneTwo
	Engine.Where("one_id =?", id).Find(&list)
	return list
}

//通过id查询数据 addby liuhan
func (ct *CategoryTwo) GetTwoCategoryById(id int64) []CategoryTwo {
	var list []CategoryTwo
	Engine.Where("id =?", id).Find(&list)
	return list
}

//根据二级类目地址查询 addby liuhan
func (ct *CategoryTwo) GetTwoCategoryByAddress() []CategoryTwo {
	var list []CategoryTwo
	Engine.Where("two_category_address =?", ct.TwoCategoryAddress).Find(&list)
	return list
}
