package models

/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
import (
	"fmt"
	"fushowcms/comm"
	"os"
	"strconv"
	"time"

	"github.com/go-xorm/core"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var (
	Engine  *xorm.Engine
	SEngine *xorm.Engine
	nowfile *os.File
)

type UidInfo struct {
	Id             int64
	UserName       string    `xorm:"notnull"` //用户名
	PassWord       string    `xorm:"notnull"` //密码
	NickName       string    `xorm:"index"`   //昵称
	Phone          string    `xorm:"index"`   //手机号
	Type           uint64    `xorm:"notnull"` //普通用户:0,主播:1,房管:2,总管:3,管理员:255,联盟管理员:4
	Balance        float64   //余额
	Form           string    //第三方登录来自于
	GiftNum        int64     //礼物数
	Level          uint64    `xorm:"index"` //等级
	RealName       string    `xorm:"index"` //真实姓名
	IdNumber       string    //身份证号
	IdentityPic    string    //身份证照片
	Remark         string    //备注
	Favicon        string    //头像
	RegTime        string    `xorm:"datetime created"` //注册时间
	LoginTime      time.Time //登录时间
	Integral       int64     //积分
	PomegranateNum int64     //石榴币
	CheckToken     string    //认证检测
	AffId          int64     //推荐机构、人等
	GuavaCoin      int64     //石榴币不清除  --->热度
	ModifyTime     string    `xorm:"datetime updated"` //修改时间	修改时间
	DelTime        string    `xorm:"datetime deleted"` //记录删除时间
	Version        int64     `xorm:"version"`          //乐观锁
	VerifyCode     string    //验证码
	Attention      int64     //关注数
	NickFlag       bool      //判断是否修改过昵称
	BankName       string    //绑定银行
	BankDeposit    string    //开户行
	BankCard       int64     //绑定银行卡号
	IsBandingBank  bool      //判断是否绑定银行卡
	RegWay         int64     //注册设备 1、PC  3、IOS   4、安卓H5   5、Android原生 addby mao
	Ip             string    //注册Ip地址
}

//用户权限  type ---> 名称
type UidType struct {
	Id       uint64 //类型   1--->n   id= 1 --->主播
	TypeName string //类型名称
}

//用户操作过程表
type TypeProcess struct {
	Id              int64  //记录Id
	TypeId          uint64 `xorm:"index"` //权限Id   如 1--->主播  255--->超管
	AuthorityListId string //目录Id  如1---->用户管理
	Uid             int64  //操作人ID
	DateTime        string `xorm:"datetime created"` //操作日期
}

//登陆记录
type UidLoginInfo struct {
	Id        int64  //记录Id
	Uid       int64  //用户Id
	LoginWay  int64  `xorm:"index"`            //登陆方式 1、PC  2、后台管理员  3、IOS  4、Android H5   5、Android原生
	LoginTime string `xorm:"datetime created"` //登录时间
	Ip        string //登陆的ip地址
}

type UidInfoJson struct {
	UidInfo        //用户表
	UID     string //用户ID
}

//活动表
//功能：注册、签到送石榴籽
//time : 2016-11-1 txl
type Event struct {
	Id        int64
	EventName string //活动名称
	StartTime string `xorm:"datetime"` //活动开始时间
	EndTime   string `xorm:"datetime"` //活动结束时间
	NowState  int64  //活动状态   //-1因活动资金不足停止 0.手动结束  1.正在进行中
	Number    int64  //活动赠送数量 （石榴籽）
	AllNumber int64  //活动预计石榴籽
	NickName  string //开启人
	EventType int64  `xorm:"index"`            //活动类型   0:注册  1：签到  2：充值
	DelTime   string `xorm:"datetime deleted"` //记录删除时间
}

//活动记录
type EventRec struct {
	Id            int64  //记录ID
	EventId       int64  `xorm:"index"` //活动Id
	EventType     int64  `xorm:"index"` //活动类型   0:注册  1：签到  2：充值
	UserId        int64  //参加用户Id
	SponsorId     int64  //发起人Id
	BalanceNumber int64  //活动剩余的金额
	DateTime      string `xorm:"datetime created"` //操作时间
	Version       int64  `xorm:"version"`          //乐观锁
}

//签到表
//功能：签到
//time : 2016-11-1 txl
type SignIn struct {
	Id         int64  //签到ID
	UserId     int64  //用户ID
	SignInTime string `xorm:"date"` //签到时间
}

//创建时间 2016-10-13 Txl
//用户注册推荐人--->注册时填写推荐机构Id
type UserAff struct {
	Uid   int64  //用户ID
	AffId string //机构ID
}

//创建时间 2016-10-13 Txl
//机构表-->网吧、推荐人
type Affiliation struct {
	Id              int64
	AffId           string    `xorm:"unique"` //机构编号
	InstitutionName string    //机构名称
	DelTime         time.Time `xorm:"deleted"` //记录删除时间
	Version         int64     `xorm:"version"` //乐观锁
}

type UserKey struct {
	Id       int64  `xorm:"index"`
	UserName string `xorm:"notnull"` //用户名
}

//用户观看记录
type UserWatch struct {
	Id        int64
	Uid       int64     `xorm:"index"`   //观看者ID
	RoomId    int64     `xorm:"index"`   //房间ID
	WatchTime time.Time `xorm:"created"` //观看时间
	Version   int64     `xorm:"version"` //乐观锁
}

//直播房间表
type AnchorRoom struct {
	Id           int64  //房间Id	房间Id：主播Id = 1:1
	Uid          int64  `xorm:"index"` //主播Id
	RoomType     uint64 `xorm:"index"`
	RoomAlias    string //房间别名
	RoomNotice   string //房间公告
	LiveAddress  string //直播地址
	LiveCover    string //直播封面
	LiveState    uint64 //直播状态	0:未直播，1:直播中 2 禁止推流
	RoomStick    int64  //是否置顶   0:不置顶 1:置顶
	PeriodsId    int64  //期Id	房管设定，包括房间分类和期
	AddTime      int64  //增加时间	添加时间
	ModifyTime   string `xorm:"datetime updated"` //修改时间	修改时间
	DelTime      string `xorm:"datetime deleted"` //记录删除时间
	ModifyUserId int64  `xorm:"index"`            //用户Id	记录修改者账号
	Version      int64  `xorm:"version"`          //乐观锁
}

//房间在现虚拟人数
type AnchorLiveNumber struct {
	Id         int64
	RoomId     int64 //房间id
	LiveNumber int64 //虚拟人数
}
type AnchorRoomInfo struct {
	AnchorRoom       `xorm:"extends"` //直播房间表
	UidInfo          `xorm:"extends"` //用户表
	AnchorLiveNumber `xorm:"extends"` //房间在线虚拟人数表
	CategoryTwo      `xorm:"extends"`
}

type PeriodsAllContent struct {
	Periods           `xorm:"extends"` //期数基本属性表
	Product           `xorm:"extends"` //产品表
	SupportManagement `xorm:"extends"` //支持管理表
}

//直播记录表
type AnchorRoomTime struct {
	Id          int64  //直播记录ID
	RoomId      int64  //直播房间ID
	Uid         int64  `xorm:"index"`    //主播ID
	StartTime   string `xorm:"datetime"` //开始时间
	EndTime     string `xorm:"datetime"` //结束时间
	AnchormTime int64  //直播时间
}

//商品列表
type Goods struct {
	Id           int64     //	商品Id
	GoodsName    string    `xorm:"notnull"` //商品名称
	GoodsType    string    //商品种类
	GoodsAccount string    //商品描述
	GoodsPirce   int64     //商品价格
	GoodsPic     string    //商品图片
	RegisterDate int64     `xorm:"created"` //上架时间
	SalesVolume  int64     //销售数量
	GoodsStock   int64     //库存数量
	GoodsDetail  string    //商品详情
	ModifyTime   int64     `xorm:"updated"` //修改时间
	DelTime      time.Time `xorm:"deleted"` //删除时间
	DelState     string    // 0存在1删除
	Version      int64     `xorm:"version"` //乐观锁
}

//订单列表
type Order struct {
	Id           int64  //	Id
	OrderID      string `xorm:"notnull"` //订单编号
	GoodsPic     string //商品图片
	GoodsId      int64  //商品id
	GoodsName    string //商品名称
	GoodsNum     int64  //商品数量
	GoodsPirce   int64  //商品价格
	GoodsAccount string //商品描述
	GoodsTotal   int64  //商品总价
	UserId       int64  //用户id
	Setstate     string //商家发货状态 1已发货 0未发货
	Delstate     string //取消订单 1已取消 0未取消
	Receiver     string //收货人
	Address      string //收货地址
	Tel          string //收货人电话
	CreateTime   string `xorm:"datetime created"` //上架时间
	ModifyTime   string `xorm:"datetime updated"` //修改时间
	DelTime      string `xorm:"datetime deleted"` //删除时间
	Version      int64  `xorm:"version"`          //乐观锁
}

type SupportManagementTest struct {
	Uid              int64  //用户Id
	PeriodsId        int64  //期数Id
	SupEncoding      string //投注结果编码  字符串开始$:#02>08    //$:#01>03#02>08#03->05#
	SupporTime       string //投注时间
	SupporNumber     int64  //石榴籽	支持数（投注数）
	Odds             string //赔率
	ComputationState int64  //核算状态，0：未核算，1：核算结束
	IsWin            string //支持状态		胜负状态，0：负，1：胜
	PrizenNmber      int64  //胜负状态，0：负，1：胜
	ProductName      string //产品名称
	State1           string //状态1
	State2           string //状态1
	State3           string //状态1
	State4           string //状态1
	State5           string //状态1
}

type MySupport struct {
	Uid          int64  //用户ID
	PeriodsId    int64  //期数Id
	SupporTime   string //投注时间
	SupporNumber int64  //石榴籽	支持数（投注数）
	Odds         string //赔率
	IsWin        int64  //支持状态		胜负状态，0：负，1：胜
	ProductName  string //产品名称
	SupportState string //状态1
}

/**
 * 结构体名称：CarousePic,轮播图数据表
 * @author 徐林->修正
 * @Time 2016-10-31
 */
type CarouselPic struct {
	Id            int64
	CarouselType  int64  //轮播类型(视频：1，图片：3)
	VideoLivePage string //视频直播页
	PicPath       string //图片路径
	PicName       string `xorm:"index"`    //图片名称
	StartTime     string `xorm："datetime"` //开始时间
	EndTime       string `xorm："datetime"` //结束时间
	State         int64  //显示状态 1、显示 2 未显示
	Sort          uint64 //排序
	Litming       string //限定
	Version       int64  `xorm:"version"` //乐观锁
}

//直播房间表关注表
type AnchorRoomConcern struct {
	Id   int64     //房间Id	房间Id：主播Id = 1:1
	Uid  int64     `xorm:"index"`   //关注者
	User int64     `xorm:"index"`   //被关注着ID
	Date time.Time `xorm:"created"` //关注日期
}

type PeriodsEarnings struct {
	Id             int64  //记录Id
	PeriodsId      int64  //期数Id
	EarningsNumber int64  //收益石榴籽
	AllNumber      int64  //所有投注
	WinNumber      int64  //所有赢得钱
	CreatTime      string `xorm:"date created"` //创建时间
}

// 关注返回  web表现形式
type AnchorRoomConcernList struct {
	Id          int64     //房间Id	房间Id：主播Id = 1:1
	Uid         int64     `xorm:"index"` //关注者
	User        int64     `xorm:"index"` //被关注着ID
	LiveId      int64     `xorm:"index"` //直播ID
	NickName    string    `xorm:"index"` //昵称
	RoomAlias   string    //房间别名
	LiveAddress string    `xorm:"index"`   //直播地址
	LiveCover   string    `xorm:"index"`   //直播封面
	Attention   int64     `xorm:"index"`   //关注数
	Date        time.Time `xorm:"created"` //关注日期
}

// 观看直播记录表
type WatchRecord struct {
	Id     int64
	Uid    int64     `xorm:"index"`   //主播Id
	LiveId int64     `xorm:"index"`   //直播ID
	Date   time.Time `xorm:"created"` //关注日期
}

// 观看返回
type WatchRecordList struct {
	Id          int64
	Uid         int64     `xorm:"index"` //主播Id
	LiveId      int64     `xorm:"index"` //直播ID
	NickName    string    `xorm:"index"` //昵称
	RoomAlias   string    //房间别名
	LiveAddress string    `xorm:"index"`   //直播地址
	LiveCover   string    `xorm:"index"`   //直播封面
	Attention   int64     `xorm:"index"`   //关注数
	Date        time.Time `xorm:"created"` //关注日期
}

type MyOrderRoom struct {
	User        int64  //主播Id
	RoomType    uint64 //房间分类	0:普通房间，1:竞猜房间
	RoomAlias   string //房间别名
	LiveAddress string //直播地址
	LiveCover   string //直播封面
	LiveNumber  int64  //直播数
	LiveState   int64  //直播状态	0:未直播，1:直播中
	Id          int64
	NickName    string //昵称
}

type CallBackDownBill struct {
	Return_code string `xml:"return_code"` //返回状态码
	Return_msg  string `xml:"return_msg"`  //返回信息
}

//充值记录
type RechargingRecords struct {
	Uid         int64  `xorm:"notnull"`          //用户ID
	TradeNo     string `xorm:"notnull unique"`   //订单号
	Money       int64  `xorm:"notnull"`          //购买金额
	Balance     int64  `xorm:"notnull"`          //石榴籽数目
	State       int    `xorm:"default 0"`        //支付状态 0：失败 1：成功
	RefundState int    `xorm:"default 0"`        //退款状态 0：未退款 1：退款成功
	Origin      string `xorm:"notnull"`          //来源 例如：alipay weixin
	Time        string `xorm:"datetime created"` //购买时间
}

type WxPayNotice struct {
	Return_code          string `xml:"return_code"`          //返回状态码
	Return_msg           string `xml:"return_msg"`           //返回信息
	Appid                string `xml:"appid"`                //公众账号ID
	Mch_id               string `xml:"mch_id"`               //商户号
	Device_info          string `xml:"device_info"`          //设备号
	Nonce_str            string `xml:"nonce_str"`            //随机字符串
	Sign                 string `xml:"sign"`                 //签名
	Result_code          string `xml:"result_code"`          //业务结果
	Err_code             string `xml:"err_code"`             //错误代码
	Err_code_des         string `xml:"err_code_des"`         //错误代码描述
	Openid               string `xml:"openid"`               //用户标识
	Is_subscribe         string `xml:"is_subscribe"`         //是否关注公众账号
	Trade_type           string `xml:"trade_type"`           //交易类型
	Bank_type            string `xml:"bank_type"`            //付款银行
	Total_fee            string `xml:"total_fee"`            //订单金额
	Settlement_total_fee string `xml:"settlement_total_fee"` //应结订单金额
	Fee_type             string `xml:"fee_type"`             //货币种类
	Cash_fee             string `xml:"cash_fee"`             //现金支付金额
	Cash_fee_type        string `xml:"cash_fee_type"`        //现金支付货币类型
	Coupon_fee           string `xml:"coupon_fee"`           //代金券金额
	Coupon_count         string `xml:"coupon_count"`         //代金券使用数量
	Transaction_id       string `xml:"transaction_id"`       //微信支付订单号
	Out_trade_no         string `xml:"out_trade_no"`         //商户订单号
	Attach               string `xml:"attach"`               //商家数据包
	Time_end             string `xml:"time_end"`             //支付完成时间
}

//支持管理表
type SupportManagement struct {
	Id               int64     //支持Id	自增Id
	Uid              int64     `xorm:"index"` //用户Id
	PeriodsId        int64     `xorm:"index"` //期数Id
	SupEncoding      string    //投注结果编码  字符串开始$:#02>08    //$:#01>03#02>08#03->05#
	SupporNumber     int64     //石榴籽	支持数（投注数）
	Odds             string    //赔率
	ComputationState int64     //核算状态，0：未核算，1：核算结束
	IsWin            int64     //支持状态		胜负状态，0：竞猜中，1：胜 2：负
	SupporTime       string    `xorm："datetime"` //投注时间	投注时间
	PrizenNmber      int64     //胜负状态，0：负，1：胜
	AccounTingTime   time.Time `xorm:"updated"` //核算时间	核算时间，由总管录入
	DelTime          time.Time `xorm:"deleted"` //记录删除时间
	//Version          int64     `xorm:"version"`

}

//期数基本属性表
//开始时间、结束时间、A队、B队、产品ID、状态、房管提交时间、总管审核时间
//StartTime、EndTime、Ateam、Bteam、ProductId、SubmitTime、VerifyTime
type Periods struct {
	Id          int64
	PeriodsId   int64     `xorm:"index unique"`     //期数ID
	StartTime   string    `xorm:"datetime notnull"` //开始时间
	EndTime     string    `xorm:"datetime notnull"` //结束时间
	RoomId      string    //同一期竞猜房间集合
	ATeam       string    `xorm:"notnull"` //A队
	BTeam       string    `xorm:"notnull"` //B队
	State       int64     `xorm:"index"`   //状态  0：未提交，1：房管提交，2：总管审核
	SubmitTime  string    `xorm:"created"` //房管提交时间
	VerifyTime  string    `xorm:"updated"` //总管审核时间
	ProEncoding string    //产品选择编码  字符串开$:#01>00#02>00#03->00
	SupEncoding string    //投注结果编码  字符串开$:#01>01#02>05#03->05
	DelTime     time.Time `xorm:"deleted"` //记录删除时间
	Version     int64     `xorm:"version"` //乐观锁
}

//每期核算资金
type PeriodsAccont struct {
	Id          int64
	PeriodsId   int64     `xorm:"index"`
	BetNumber   float64   //投注总额，所有投注的石榴籽，投注额小于支出 时，从资金池划拨，投注额大于支出）是向资金池划拨，
	Expend      float64   //支出：所有用户盈利的总额
	AllotNumber float64   //向资金池划拨数量，正值时向资金池加计算，负值时向资金池减计算：划拨数量 =投注总额 -支出
	CheckTime   time.Time `xorm:"created"` //核算时间
}

//增加核算资金流向查询
type FundAccounting struct {
	Id        int64
	Uid       int64   `xorm:"index"` //用户ID
	Money     float64 //资金
	PeriodsId int64   `xorm:"index"`            //期ID
	AddTime   string  `xorm:"datetime created"` //增加时间
	Version   int64   `xorm:"version"`          //乐观锁
}

//产品基本属性表
type Product struct {
	Id          int64
	ProductName string    `xorm:"notnull"` //产品名称
	State1      string    //状态1
	State2      string    //状态2
	State3      string    //状态3
	State4      string    //状态4
	State5      string    //状态5
	State6      string    //状态6
	State7      string    //状态7
	State8      string    //状态8
	State9      string    //状态9
	State10     string    //状态10
	AddTime     time.Time `xorm:"created"` //添加时间
	ModifyTime  time.Time `xorm:"updated"` //修改时间
	DelTime     time.Time `xorm:"deleted"` //删除时间
	Version     int64     `xorm:"version"` //乐观锁
}

//期数产品过程属性表
type PeriodsProduct struct {
	Id           int64     //Id
	PeriodsId    int64     `xorm:"index"` //期数ID
	ProductId    int64     `xorm:"index"` //产品ID
	State1Odds   string    //赔率	  每个选项的赔率
	State2Odds   string    //赔率
	State3Odds   string    //赔率
	State4Odds   string    //赔率
	State5Odds   string    //赔率
	State6Odds   string    //赔率
	State7Odds   string    //赔率
	State8Odds   string    //赔率
	State9Odds   string    //赔率
	State10Odds  string    //赔率
	State1Hot    int64     //热度1
	State2Hot    int64     //热度2
	State3Hot    int64     //热度3
	State4Hot    int64     //热度4
	State5Hot    int64     //热度5
	State6Hot    int64     //热度6
	State7Hot    int64     //热度7
	State8Hot    int64     //热度8
	State9Hot    int64     //热度9
	State10Hot   int64     //热度10
	AddTime      int64     `xorm:"created"` //增加时间	资金池追加总额是添加，别时为空
	ModifyTime   int64     `xorm:"updated"` //修改时间	变化是修改
	ModifyUserId int64     `xorm:"index"`   //用户id	   	记录核算者账号
	DelTime      time.Time `xorm:"deleted"` //删除时间
	Version      int64     `xorm:"version"` //乐观锁
}

//资金池过程表
type Fund struct {
	Id            int64     //资金池ID，每次资金池变化都增加一条记录，最新资金池取最后一条
	OfferAmount   float64   //发行总额：发行总额=储备金+ 流通金
	StorageFund   float64   //储备金：资金池划拨时，增加储备金同时减少流通金
	CurrencyMoney float64   //流通金
	AddTime       time.Time `xorm:"created"` //增加时间
	ModifyTime    time.Time `xorm:"updated"` //修改时间
	ModifyUserid  int64     `xorm:"index"`   //用户id
	Version       int64     `xorm:"version"` //乐观锁
}

//资金池过程表
type FundGift struct {
	Id         int64 //资金池ID，每次资金池变化都增加一条记录，最新资金池取最后一条
	AddTime    int64 `xorm:"created"` //增加时间
	ModifyTime int64 `xorm:"updated"` //修改时间
	GiftNumber int64 //礼物数量
	Version    int64 `xorm:"version"` //乐观锁
}

//用户充值资金池记录表
type Fund1 struct {
	Id           int64   //资金池ID，每次资金池变化都增加一条记录，最新资金池取最后一条
	Uid          int64   //用户ID
	Money        float64 //充值金额
	Number       float64 //购买石榴籽数
	AddTime      int64   `xorm:"created"` //增加时间
	ModifyTime   int64   `xorm:"updated"` //修改时间
	ModifyUserid int64   `xorm:"index"`   //用户id
	Version      int64   `xorm:"version"` //乐观锁
}

//定期汇总表
type Collect struct {
	Id      int64 //资金池ID，每次资金池变化都增加一条记录，最新资金池取最后一条
	firstId int64 //开始ID
	endId   int64 //结束ID
}

//礼物属性表
type Gift struct {
	Id            int64
	GiftName      string    `xorm:"notnull"` //礼物名称
	GiftType      int64     //礼物类型
	GiftAccount   string    //礼物描述明细
	BuyNumber     int64     //购买时所需石榴籽
	ToNumber      int64     //主播可获得石榴籽
	GiftPicture   string    //礼物图片
	GiftPicStatic string    // 静态图
	RegisterDate  time.Time `xorm:"created"` //注册时间
	State         int64     //礼物状态
	Version       int64     `xorm:"version"` //乐观锁
}

//礼物赠送表  -->向主播刷礼物
type GiftGive struct {
	Id           int64  //记录Id
	BenefactorId int64  `xorm:"index"`            //赠送人Id
	RecipientId  int64  `xorm:"index"`            //接收人Id
	GiveDate     string `xorm:"datetime created"` //赠送时间
	GiftId       int64  `xorm:"index"`            //礼物Id
	GiftNum      int64  //礼物数量
	AllNumber    int64  //礼物总价值
	DelTime      string `xorm:"datetime deleted"` //删除时间
	Version      int64  `xorm:"version"`          //乐观锁
}

//主播申请表
type Applicant struct {
	Id            int64     //申请Id	自增，用户申请记录标志Id
	UserId        int64     `xorm:"index"`            //用户Id
	ApplicantTime string    `xorm:"datetime created"` //申请时间
	State         int64     `xorm:"notnull"`          //审核状态	0：未审核，1：审核通过，2：审核未通过
	VerifyUserId  int64     //审核用户Id	记录核算者账号
	VerifyTime    int64     `xorm:"updated"` //审核时间
	DelTime       time.Time `xorm:"deleted"` //删除时间
	Version       int64     `xorm:"version"` //乐观锁
}

//后台单条申请信息
type ApplicantInfo struct {
	Id            int64
	Uid           int64  //用户ID
	NickName      string `xorm:"index"`   //昵称
	Phone         string `xorm:"index"`   //手机号
	Type          uint64 `xorm:"notnull"` //普通用户:0,主播:1,房管:2,总管:3,管理员:255
	Level         uint64 `xorm:"index"`   //等级
	RealName      string `xorm:"index"`   //真实姓名
	IdNumber      string //身份证号
	IdentityPic   string //身份证照片
	ApplicantTime string `xorm:"datetime created"` //申请时间
	ApplyId       int64  //申请ID
	State         int64  `xorm:"notnull"` //审核状态	0：未审核，1：审核通过，2：审核未通过
}

//后台申请样式
type ApplicantDem struct {
	Id            int64
	UserId        int64     //用户ID
	Level         uint64    //等级
	NickName      string    //昵称
	ApplicantTime time.Time `xorm:"created"` //申请时间
	State         int64     //状态
}

//充值记录表
type Seed struct {
	Id            int64   //充值Id
	Uid           int64   `xorm:"index"` //用户Id
	RechargeMoney float32 //充值金额
	SeednNmber    int64   //获取石榴籽数
	RechargeDate  string  `xorm:"datetime created"` //充值时间
	DelTime       string  `xorm:"datetime deleted"` //删除时间
	Version       int64   `xorm:"version"`          //乐观锁
}

//商品列表
type Commodity struct {
	Id               int64     //	商品Id
	CommodityName    string    `xorm:"notnull"` //商品名称
	CommodityType    string    //商品种类
	CommodityAccount string    //商品描述
	CommodityPirce   float64   //商品价格
	CommodityPic     string    //商品图片
	RegisterDate     int64     `xorm:"created"` //上架时间
	SalesVolume      float64   //销售数量
	State            int64     //状态(是否首页)
	ModifyTime       int64     `xorm:"updated"` //修改时间
	DelTime          time.Time `xorm:"deleted"` //删除时间
	Version          int64     `xorm:"version"` //乐观锁
}

//兑换商品列表
type Convert struct {
	Id           int64     //兑换Id	自增
	UserId       int64     `xorm:"index"`   //用户Id	送出石榴籽
	ToUserId     int64     `xorm:"index"`   //接受Id	获取石榴籽
	ExchangeDate int64     `xorm:"created"` //兑换时间
	SeedNum      int64     //兑换积分
	DelTime      time.Time `xorm:"deleted"` //删除时间
	Version      int64     `xorm:"version"` //乐观锁
}

//网站信息
type Website struct {
	Id           int64
	Title        string //网站标题
	Keywords     string //网站关键字
	Descrip      string //网站描述
	Name         string //网站名称
	Logo         string //网站LOGO
	Icon         string //网站地址栏图标
	URL          string //网址
	ModifyUserId int64
	ModifyTime   string `xorm:"datetime created"`
}

type AnchorInfo struct {
	Id              int64  //	房间Id	房间Id：主播Id = 1:1
	Uid             int64  `xorm:"index"` //主播Id
	RoomType        uint64 `xorm:"index"`
	TwoCategoryName string //类目名称
	RoomAlias       string `xorm:"index"` //房间别名
	LiveAddress     string `xorm:"index"` //直播地址
	LiveNumber      int64  `xorm:"index"` //直播地址
	LiveCover       string `xorm:"index"` //直播封面
	NickName        string //昵称
}

//官方活动
type OfficialFunctions struct {
	Id               int64  //活动编码
	PicURL           string //图片地址
	OfficialURL      string //活动路径
	OfficialName     string `xorm:"index"` //活动名称
	OfficialBriefing string `xorm:"TEXT"`  //活动简介
	LiveState        int64  //直播状态
	StartTime        int64  //开始时间
	EndTime          int64  //结束时间
}

//直播间广告
type Dbadvertising struct {
	Id           int64  //广告编码
	PicURL       string //广告图片
	DbadURL      string //广告地址
	DbadName     string `xorm:"index"` //广告名称
	DbadBriefing string `xorm:"TEXT"`  //广告简介
	LiveState    int64  //直播状态
	StartTime    int64  //开始时间
	EndTime      int64  //结束时间
}

//左侧导航 广告
type Sdbadvertising struct {
	Id            int64
	PicURL        string //广告图片
	SdbadURL      string //广告地址
	SdbadName     string //广告名称
	SdbadBriefing string //广告简介
	LiveState     int64  //直播状态
	StartTime     int64  //开始时间
	EndTime       int64  //结束时间
}
type Unifiedorder struct {
	Return_code string `xml:"return_code"` //返回状态码
	Return_msg  string `xml:"return_msg"`  //返回信息
	Appid       string `xml:"appid"`       //公众账号ID
	Mch_id      string `xml:"mch_id"`      //商户号
	Nonce_str   string `xml:"nonce_str"`   //随机字符串
	Sign        string `xml:"sign"`        //签名
	Result_code string `xml:"result_code"` //业务结果
	Trade_type  string `xml:"trade_type"`  //交易类型
	Prepay_id   string `xml:"prepay_id"`   //预支付交易会话标识
	Code_url    string `xml:"code_url"`    //二维码链接
	Partnerid   string //合作ID
	Package     string //包
	Timestamp   int64  //时间戳
}

type OrderQuery struct {
	Return_code          string `xml:"return_code"`          //返回状态码
	Return_msg           string `xml:"return_msg"`           //返回信息
	Appid                string `xml:"appid"`                //公众账号ID
	Mch_id               string `xml:"mch_id"`               //商户号
	Device_info          string `xml:"device_info"`          //设备号
	Nonce_str            string `xml:"nonce_str"`            //随机字符串
	Sign                 string `xml:"sign"`                 //签名
	Result_code          string `xml:"result_code"`          //业务结果
	Openid               string `xml:"openid"`               //用户标识
	Is_subscribe         string `xml:"is_subscribe"`         //是否关注公众账号
	Trade_type           string `xml:"trade_type"`           //交易类型
	Trade_state          string `xml:"trade_state"`          //交易状态
	Bank_type            string `xml:"bank_type"`            //付款银行
	Total_fee            string `xml:"total_fee"`            //订单金额
	Settlement_total_fee string `xml:"settlement_total_fee"` //应结订单金额
	Fee_type             string `xml:"fee_type"`             //货币种类
	Cash_fee             string `xml:"cash_fee"`             //现金支付金额
	Transaction_id       string `xml:"transaction_id"`       //微信支付订单号
	Out_trade_no         string `xml:"out_trade_no"`         //商户订单号
	Time_end             string `xml:"time_end"`             //支付完成时间
	Trade_state_desc     string `xml:"trade_state_desc"`     //交易状态描述
}

type Refund struct {
	Return_code    string `xml:"return_code"`    //返回状态码
	Return_msg     string `xml:"return_msg"`     //返回信息
	Appid          string `xml:"appid"`          //公众账号ID
	Mch_id         string `xml:"mch_id"`         //商户号
	Nonce_str      string `xml:"nonce_str"`      //随机字符串
	Sign           string `xml:"sign"`           //签名
	Result_code    string `xml:"result_code"`    //业务结果
	Transaction_id string `xml:"transaction_id"` //微信订单号
	Refund_id      string `xml:"refund_id"`      //微信退款单号
	Refund_channel string `xml:"refund_channel"` //退款渠道
	Refund_fee     string `xml:"refund_fee"`     //申请退款金额
	Total_fee      string `xml:"total_fee"`      //订单金额
	Out_refund_no  string `xml:"out_refund_no"`  //商户退款单号
	Out_trade_no   string `xml:"out_trade_no"`   //商户订单号
}

type FundQuery struct {
	Return_code          string `xml:"return_code"`          //返回状态码
	Return_msg           string `xml:"return_msg"`           //返回信息
	Result_code          string `xml:"result_code"`          //公众账号ID
	Err_code             string `xml:"err_code"`             //商户号
	Err_code_des         string `xml:"err_code_des"`         //随机字符串
	Appid                string `xml:"appid"`                //公众账号ID
	Mch_id               string `xml:"mch_id"`               //商户号
	Nonce_str            string `xml:"nonce_str"`            //随机字符串
	Sign                 string `xml:"sign"`                 //签名
	Transaction_id       string `xml:"transaction_id"`       //微信订单号
	Refund_channel       string `xml:"refund_channel"`       //退款渠道
	Refund_fee           string `xml:"refund_fee"`           //申请退款金额
	Cash_fee             string `xml:"cash_fee"`             //现金支付金额
	Fee_type             string `xml:"fee_type"`             //费用类型
	Total_fee            string `xml:"total_fee"`            //订单金额
	Refund_count         string `xml:"refund_count"`         //商户退款单号
	Out_trade_no         string `xml:"out_trade_no"`         //商户订单号
	Settlement_total_fee string `xml:"settlement_total_fee"` //商户订单号
}

type AnchorRoomAllInfo struct {
	NickName   string //主播名
	Favicon    string //头像
	RoomAlias  string //房间别名
	RoomNotice string //房间公告
	Attention  int64  //关注数
}

type MobileSMS struct {
	Account  string //描述
	Password string //密码
	Moible   string //手机号
	Content  string //内容
}

//验证码表
type VerificationCode struct {
	Id    int64
	Phone string `xorm:"unique index"` //手机号
	Code  string //验证码
}

//支持管理结果表
type SupportManagementResult struct {
	Uid              int64 `xorm:"index"` //用户Id
	PeriodsId        int64 `xorm:"index"` //期数Id
	ComputationState int64 //支持状态		胜负状态，0：未核算，1：核算结束
	IsWin            bool  //支持状态		胜负状态，0：负，1：胜
}

//查询支持石榴总数和支持人数
type Number struct {
	SupportNumber string //热度
	RowNumber     int64  //支持数
	Odds          string //赔率
	SupEncoding   string //支持编码
}

type Test struct {
	Id           int64
	Uid          int64   `xorm:"index"` //用户Id
	PeriodsId    int64   `xorm:"index"` //期数Id
	Balance      float64 //余额
	Version      int64   `xorm:"version"` //乐观锁
	SupporNumber float64 //石榴籽	支持数（投注数）
	Odds         string  //赔率
}

//房间房管管理  addby liuhan
type RoomUserManage struct {
	Id         int64  //	Id
	RoomId     int64  `xorm:"notnull"` //房间id
	UserId     int64  `xorm:"notnull"` //用户id
	NickName   string //昵称
	ModifyBy   int64  //修改人id
	ModifyTime string `xorm:"datetime created"` //修改时间
	DelTime    string `xorm:"datetime deleted"` //删除时间
}

//主播结算明细表
type SettlementDetail struct {
	Id              int64   //明细id
	Uid             int64   `xorm:"index"` //用户ID
	Cashing         float64 //结算金额(主播石榴币结算80%)
	IsApply         bool    //是否申请结算(0是未申请，1是已经申请)
	CashingDate     string  `xorm:"datetime"` //申请结算时间
	ApplyCashingNum int64   //申请结算石榴币
	IsCashing       bool    //结算状态(是在申请结算后，看后台管理员是否结算)(0是申请后未结算，1是申请后已结算)
}

type UidInfoSettlementDetail struct {
	SettlementDetail `xorm:"extends"` //主播结算明细表
	UserName         string           `xorm:"notnull"` //用户名
	PassWord         string           `xorm:"notnull"` //密码
	NickName         string           `xorm:"index"`   //昵称
	Phone            string           `xorm:"index"`   //手机号
	Type             uint64           `xorm:"notnull"` //普通用户:0,主播:1,房管:2,总管:3,管理员:255,联盟管理员:4
	Balance          float64          //余额
	Form             string           //第三方登录来自于
	GiftNum          int64            //礼物数
	Level            uint64           `xorm:"index"` //等级
	RealName         string           `xorm:"index"` //真实姓名
	IdNumber         string           //身份证号
	IdentityPic      string           //身份证照片
	Remark           string           //备注
	Favicon          string           //头像
	RegTime          string           `xorm:"date created"` //注册时间
	LoginTime        time.Time        //登录时间
	Integral         int64            //积分
	PomegranateNum   int64            //石榴币
	CheckToken       string           //检查标识
	AffId            int64            //推荐机构、人等
	GuavaCoin        int64            //石榴币不清除  --->热度
	ModifyTime       string           `xorm:"datetime updated"` //修改时间	修改时间
	DelTime          string           `xorm:"datetime deleted"` //记录删除时间
	Version          int64            `xorm:"version"`          //乐观锁
	VerifyCode       string           //验证码
	Attention        int64            //关注数
	NickFlag         bool             //判断是否修改过昵称
	BankName         string           //绑定银行
	BankDeposit      string           //开户行
	BankCard         int64            //绑定银行卡号
	IsBandingBank    bool             //判断是否绑定银行卡
}

//赠送石榴籽表	  addby liuhan
type GiveSlz struct {
	Id           int64  //记录Id
	BenefactorId int64  `xorm:"index"`            //赠送人Id
	RecipientId  int64  `xorm:"index"`            //接收人Id
	GiveDate     string `xorm:"datetime created"` //赠送时间
	Num          int64  //数量
	DelTime      string `xorm:"datetime  deleted"` //删除时间
	Version      int64  `xorm:"version"`           //乐观锁
}

type PeriodsRoom struct {
	Periods
	RoomAlias string //房间名
	Earnings  string //竞猜收益
}

//机构表  addby liuhan
type Organ struct {
	Id         int64
	OrganCode  string `xorm:"notnull"`          //机构编号
	OrganName  string `xorm:"notnull"`          //机构名称
	PassWord   string `xorm:"notnull"`          //密码
	ModifyTime string `xorm:"datetime created"` //修改时间
	DelTime    string `xorm:"datetime deleted"` //记录删除时间
	Version    int64  `xorm:"version"`          //乐观锁
}

//机构表管理  addby liuhan
type OrganManage struct {
	Id             int64
	UserId         int64  //用户名称
	OrganId        int64  //机构id
	RechargeNum    int64  //充值数量
	RechargeMethod string //充值方式
	ModifyTime     string `xorm:"datetime created"` //修改时间
	DelTime        string `xorm:"datetime deleted"` //记录删除时间
	Version        int64  `xorm:"version"`          //乐观锁
}

//禁止用户发言表  addby liuhan
type NotUserSpeak struct {
	Id         int64
	UserId     int64  //用户id
	RoomId     int64  //房间id
	ModifyTime string `xorm:"datetime created"` //修改时间
	DelTime    string `xorm:"datetime deleted"` //记录删除时间
	Version    int64  `xorm:"version"`          //乐观锁
}

//广播管理  addby liuhan
type Broadcast struct {
	Id         int64
	UserId     int64  //用户ID
	Content    string //内容
	StartTime  string //开始时间
	EndTime    string //结束时间
	ModifyTime string `xorm:"datetime created"` //修改时间
	DelTime    string `xorm:"datetime deleted"` //记录删除时间
	Version    int64  `xorm:"version"`          //乐观锁
}

//赠送表DTO addby liuhan
type GiftGives struct {
	Id           int64  //记录Id
	BenefactorId int64  `xorm:"index"`            //赠送人Id`
	RecipientId  int64  `xorm:"index"`            //接收人Id
	GiveDate     string `xorm:"datetime created"` //赠送时间
	GiftId       int64  //礼物Id
	GiftName     string //礼物名称
	GiftNum      int64  //数量
	AllNumber    int64  //总价值
	DelTime      string `xorm:"datetime  deleted"` //删除时间
	Version      int64  `xorm:"version"`           //乐观锁
}

//用户权限表 cnxulin
type AuthorityList struct {
	Id        int64
	Type      uint64 `xorm:"notnull"` //普通用户:0,主播:1,房管:2,总管:3,管理员:255,联盟管理员:4
	Icon      string //图标
	Menus     int64  `xorm:"notnull"` //一级目录
	SecondDir int64  `xorm:"notnull"` //二级目录
	ListName  string //显示名称
	Address   string //地址
	Sort      uint64 //排序
}

//期Id
type PerId struct {
	Id     int64
	Date   string //日期
	Pid    int64  //期数ID
	PersId string //期ID
}

// 后台提交支持记录表 @liuhan 170314
type SupportRecord struct {
	Id          int64
	PerId       int64  //期id
	SubmitterId int64  //提交人
	ModifyTime  string `xorm:"datetime created"` //创建时间
	Result      string //提交的结果是什么
	Type        int    //类型  1：提交结果2：核算
}

//一级类目管理 addby liuhan
type CategoryOne struct {
	Id              int64  //一级类目ID
	OneCategoryName string //一级类目名称
}

//二级类目管理 addby liuhan
type CategoryTwo struct {
	Id                 int64  //二级类目ID
	TwoCategoryName    string //二级类目名称
	TwoCategoryImage   string //二级类目图片
	TwoCategoryAddress string //二级类目地址
}

//二级类目管理 addby liuhan
type CategoryOneTwo struct {
	Id    int64 //二级类目ID
	OneId int64 //一级类目ID
	TwoId int64 //一级类目ID
}

type SLTVMessage struct {
	Id         int64
	UrlHost    string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPass     string
	DbDbname   string
	DbSelectdb string
	RedisDial  string
	RedisDeal  string
	RedisPass  string
	RedisKey   string
	ModifyTime string `xorm:"datetime created"`
}

type SetUp struct {
	Id         int64
	Ip         string
	Phone      string
	ModifyTime string `xorm:"datetime created"`
}

func InitDataBase(userName, psw string) {
	var (
		user     UidInfo
		user_key UserKey
		fund     Fund
	)
	user.UserName = "admin"
	user.PassWord = "Hoi3di/cjFRY38JsKYG8kg==" // 123qwe
	user.NickName = "系统"
	user.Phone = "15900000000"
	user.Type = 255
	user.Balance, _ = strconv.ParseFloat(comm.GetConfig("Fund", "user_money"), 10)
	Engine.Insert(&user)
	// -----------------------------------------------------------------------------------------------
	fund.StorageFund, _ = strconv.ParseFloat(comm.GetConfig("Fund", "smoney"), 10) //储备金
	nowc, _ := strconv.ParseFloat(comm.GetConfig("Fund", "cmoney"), 10)            //初始化流通金
	userc, _ := strconv.ParseFloat(comm.GetConfig("Fund", "user_money"), 10)       //管理员初始化流通金
	fund.CurrencyMoney = nowc + userc                                              //流通金
	fund.OfferAmount = fund.StorageFund + fund.CurrencyMoney
	Engine.Insert(&fund)
	fmt.Println("uid_info", user.Id)
	//user_key.UserName = userName
	Engine.Insert(&user_key)
	fmt.Println("userKey", user_key.Id)
	Engine.Exec("ALTER TABLE anchor_room AUTO_INCREMENT = 10001")
	Engine.Exec("ALTER TABLE uid_info AUTO_INCREMENT = 10001;")
	Engine.Exec("DROP TABLE `authority_list`;")
	Engine.Exec("CREATE TABLE `authority_list` (`id` bigint(20) NOT NULL, `type` bigint(20) NOT NULL, `list_name` varchar(255) DEFAULT NULL, `address` varchar(255) DEFAULT NULL, `sort` bigint(20) DEFAULT NULL, `menus` bigint(20) NOT NULL, `second_dir` bigint(20) NOT NULL, `icon` varchar(255) DEFAULT NULL ) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	Engine.Exec("ALTER TABLE `authority_list` ADD PRIMARY KEY (`id`);")
	Engine.Exec("ALTER TABLE `authority_list` MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;")
	Engine.Exec("INSERT INTO `uid_type` VALUES (255,'超级管理员');")
	Engine.Exec("INSERT INTO `uid_type` VALUES (99,'普通用户');")
	Engine.Exec("INSERT INTO `uid_type` VALUES (4,'联盟单位');")
	Engine.Exec("INSERT INTO `uid_type` VALUES (1,'主播');")
	Engine.Exec("INSERT INTO `type_process` VALUES (1,255,'1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,20,21,22,23,24,25,26,27,28,29,30,31',1,'2016-11-23 10:21:45');")
	Engine.Exec("INSERT INTO `category_two` VALUES (99999999,'其他');")
	//后台菜单栏
	Engine.Exec("INSERT INTO `authority_list` VALUES (1, 2, '用户管理', '', 1001, 1, 0, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (2, 2, '用户管理', '/root/user_manage', 1002, 0, 1, '/static/images/user_manage.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (3, 2, '审核管理', '/root/apply_manage', 1003, 0, 1, '/static/images/shgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (4, 2, '房间管理', '', 2001, 1, 0, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (5, 2, '房间管理', '/root/room_manage', 2002, 0, 4, '/static/images/fang.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (6, 2, '直播记录管理', '/root/play_manage', 2003, 0, 4, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (7, 2, '资金管理', '', 3001, 1, 0, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (8, 2, '资金池管理', '/root/pool_manage', 3002, 0, 7, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (9, 2, '订单管理', '/root/order_manage', 3003, 0, 7, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (10, 2, '结算管理', '/root/anchor', 3004, 0, 7, '/static/images/user_manage.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (11, 2, '充值记录', '/root/pay_manage', 3005, 0, 7, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (12, 2, '送礼物记录', '/root/gift_record', 3006, 0, 7, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (17, 2, '联盟管理', '', 5001, 1, 0, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (18, 2, '管理员联盟管理', '/root/root_organ_manage', 5002, 0, 17, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (19, 4, '联盟单位', '/root/organ_unit', 5003, 0, 17, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (20, 2, '联盟管理', '/root/affiliation_manage', 5003, 0, 17, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (21, 2, '活动管理', '', 6001, 1, 0, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (22, 2, '活动管理', '/root/event_manage', 6002, 0, 21, '/static/images/shgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (23, 2, '轮播图管理', '/root/carousel_manage', 6003, 0, 21, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (24, 2, '官方活动', '/root/official_manage', 6004, 0, 21, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (25, 2, '全频道广播', '/root/full_channel_broadcasting', 6005, 0, 21, '/static/images/user_manage.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (26, 2, '赠送管理', '/root/givenumber_manage', 6006, 0, 21, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (27, 2, '礼物管理', '/root/gift_manage', 7001, 0, 0, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (28, 2, '商品管理', '/root/root_goods_manage', 8001, 0, 0, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (29, 2, '栏目管理', '/root/category_manage', 9001, 0, 0, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (30, 2, '权限管理', '/root/uidtype_manage', 9002, 0, 0, '/static/images/lwgl.png');")
	Engine.Exec("INSERT INTO `authority_list` VALUES (31, 2, '设置', '/root/website', 10001, 0, 0, '/static/images/lwgl.png');")

}
func init() {
	var err error
	host := comm.GetConfig("DB", "host")
	port := comm.GetConfig("DB", "port")
	user := comm.GetConfig("DB", "user")
	pass := comm.GetConfig("DB", "pass")
	dbname := comm.GetConfig("DB", "dbname")
	selectdb := comm.GetConfig("DB", "selectdb")
	if host == "no value" {
		return
	}
	myconn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	Engine, err = xorm.NewEngine(selectdb, myconn)
	err = Engine.Sync2(
		new(CategoryOne),
		new(CategoryTwo),
		new(CategoryOneTwo),
		new(AnchorLiveNumber),
		new(Sdbadvertising),
		new(Dbadvertising),
		new(GiveSlz),
		new(Organ),
		new(OrganManage),
		new(PerId),
		new(UidLoginInfo),
		new(Broadcast),
		new(GiftGives),
		new(TypeProcess),
		new(RoomUserManage),
		new(Commodity),
		new(SupportManagementResult),
		new(NotUserSpeak),
		new(Convert),
		new(Event),
		new(SignIn),
		new(Product),
		new(UidInfo),
		new(CarouselPic),
		new(Periods),
		new(PeriodsProduct),
		new(Fund),
		new(Gift),
		new(GiftGive),
		new(EventRec),
		new(Applicant),
		new(Seed),
		new(WatchRecord),
		new(VerificationCode),
		new(FundGift),
		new(AnchorRoom),
		new(AnchorRoomTime),
		new(AnchorRoomConcern),
		new(UserKey),
		new(UserAff),
		new(PeriodsEarnings),
		new(Affiliation),
		new(OfficialFunctions),
		new(SupportManagement),
		new(Goods),
		new(UidType),
		new(Order),
		new(RechargingRecords),
		new(Website),
		new(AuthorityList),
		new(SettlementDetail),
		new(SLTVMessage),
		new(SetUp),
	)

	if err != nil {
		fmt.Println(err)
	}
	//权限配合
	usertype()
	//日志执行
	logfile()
}

func usertype() {
	var list []UidType
	err := Engine.Table("uid_type").Where("id>?", 0).Find(&list)
	if err != nil {
		fmt.Println(err)
	}
	//存在
	comm.UserTyped = make(map[string]uint64)
	if len(list) > 0 {
		for i := 0; i < len(list); i++ {
			if list[i].TypeName == "" {
				continue
			}
			comm.UserTyped[list[i].TypeName] = list[i].Id
		}
	}
}

//Log日志
func logfile() {
	//首次创建文件
	_, err := os.Stat("Logger")
	if err != nil {
		os.Mkdir("Logger", 0777)
	}

	date := time.Now().Format("2006-01-02")

	nowfile, err = os.Create("Logger/" + date + ".log")
	if err != nil {
		println(err.Error())
		return
	}
	Engine.SetLogger(xorm.NewSimpleLogger(nowfile))
	Engine.Logger().SetLevel(core.LOG_INFO)
	Engine.ShowSQL(true)

	ticker := time.NewTicker(time.Hour * 24)
	go func() {
		for _ = range ticker.C {
			nowfile.Close()
			date := time.Now().Format("2006-01-02")

			datefile, err := os.Create("Logger/" + date + ".log")
			if err != nil {
				println(err.Error())
				return
			}
			Engine.SetLogger(xorm.NewSimpleLogger(datefile))
			Engine.Logger().SetLevel(core.LOG_INFO)

			fmt.Println("ticked at %v", date)
		}
	}()
}
