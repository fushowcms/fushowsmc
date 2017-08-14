package routers

import (
	"fushowcms/comm"
	"fushowcms/control"

	"github.com/gin-gonic/gin"
)

func init() {

}

func Run() {
	go control.GetDoMainName()
	router := gin.Default()
	router.LoadHTMLGlobNoSVN("views/**/*")
	router.Static("/static", "./static")
	router.POST("/upload", control.Upload)
	router.POST("/getSLTVMessage", control.GetSLTVMessage) // 石榴报错返回信息
	v1 := router.Group("")
	{
		v1.GET("/", control.GetActionPage)
	}
	v1.GET("/overweb", control.OverWeb) //礼物赠送
	v1.GET("/init", control.InitDatabase)
	v1.GET("/publishcallback", control.PublishCallBack) //推流信息
	v1.GET("/stopfushow", comm.StopServer)
	v1.GET("/updatefushow", comm.UpdateServer)
	temp := router.Group("")
	temp.GET("/index", control.GetIndexPage)
	v1.GET("/pandainfo", control.PandaInfo) //熊猫
	v1.GET("/getimagecode", control.Getimagecode)
	user := router.Group("/user", control.IsAccess())
	{
		user.GET("/*action", control.GetActionUser)
		user.POST("/userup", control.UserUp)                               //后台返回格式已修改
		user.POST("/passup", control.PassUp)                               //无用，先保留
		user.POST("/pcpassup", control.PcPassUp)                           //后台返回格式已修改
		user.POST("/supportadd", control.SupportAdd)                       //支持增删改查
		user.POST("/getAnchorRoomTimeList", control.GetAnchorRoomTimeList) //获取直播记录
		user.POST("/watchrecord", control.WatchRecord)                     //查询我的观看记录
		user.POST("/applyadd", control.ApplyAdd)
		user.POST("/isapplyExit", control.IsApplyExit)       //判断审核状态
		user.POST("/useridnumber", control.UserIdNumber)     //web端成为主播
		user.POST("/givegiftnumadd", control.GiveGiftNumAdd) //礼物赠送
		user.POST("/userupload", control.UserUpload)         //  WEB上传头像
		user.POST("/unlogin", control.UnLogin)               //退出登录//已修改后台返回参数20161130
		user.POST("/roomconcernadd", control.RoomConcernAdd) //添加一条关注记录
		user.POST("/roomconcerndel", control.RoomConcernDel)
		user.POST("/getroomconcern", control.GetRoomConcern)
		user.POST("/roomupis", control.RoomUp)
		user.POST("/getroominfo", control.GetRoomInfo)               //获取直播信息
		user.POST("/givenumber", control.GiveNumber)                 //赠送石榴籽
		user.POST("/isconcern", control.IsConcern)                   //是否已经关注
		user.POST("/getMyOrderRoomList", control.GetMyOrderRoomList) //关注列表
		user.POST("/cancelroomcon", control.CancelRoomCon)           // 取消关注
		user.POST("/cancelorderoom", control.CancelOrderRoom)        //手机端取消关注
		user.POST("/getmyattentionlist", control.GetMyAttention)     // 手机端关注 我的所有关注列表
		user.POST("/getSupportUidList", control.GetSupportUidList)   //我的支持列表
		user.POST("/isOpenMyAttention", control.IsOpenMyAttention)   //查看我关注的正在直播的主播
		user.POST("/test", control.GoodsUpload)
		user.POST("/checknicknamechange", control.CheckNickNameChange) //查询是否修改过昵称
		user.POST("/changenickname", control.ChangeNickname)           //修改昵称
		user.POST("/userpayrecord", control.UserPayRecord)             //人充值记录
		user.POST("/findByRoomIdAll", control.FindByRoomIdAll)
		user.POST("/addRoomusermanage", control.AddRoomUserManage)
		user.POST("/delRoomusermanage", control.DelRoomUserManage)
		user.POST("/delfrontRoomusermanage", control.DelFrontUserManage)
		user.POST("/signinadd", control.SigninAdd)                     //签到
		user.POST("/getsigninlist", control.GetSignInList)             //签到记录
		user.POST("/findByUserIdAll", control.FindByUserIdAll)         //礼物记录
		user.POST("/findByUserIdGiveSlz", control.FindByUserIdGiveSlz) //根据userId查询赠送石榴籽记录
		user.POST("/findOrgan", control.FindOrgan)
		user.POST("/findOrganById", control.FindOrganById)
		user.POST("/findOrganByOrganCode", control.FindOrganByOrganCode) //根据机构编码查询
		user.POST("/findOrganByOrganName", control.FindOrganByOrganName)
		user.POST("/addOrgan", control.AddOrgan)
		user.POST("/updateOrgan", control.UpdateOrgan)
		user.POST("/delOrgan", control.DelOrgan)
		user.POST("/findOrganManageByOrganId", control.FindOrganManageByOrganId)
		user.POST("/addOrganManage", control.AddOrganManage) //绑定送石榴籽
		user.POST("/sumOrganManageByOrganId", control.SumOrganManageByOrganId)
		user.POST("/findBroadcast", control.FindBroadcast)
		user.POST("/addBroadcast", control.AddBroadcast)
		user.POST("/findOrganByUserId", control.FindOrganByUserId)
		user.POST("/addbannedtopost", control.AddNotUserSpeak)
		user.POST("/getEventInfo", control.GetEventInfo)             //判断绑定联盟活动是否存在，得到多少石榴籽
		user.POST("/anchorbindingbank", control.AnchorBindingBank)   //主播绑定银行卡
		user.POST("/anchorapplycashing", control.AnchorApplyCashing) //主播申请结算
		user.POST("/settlementdetail", control.SettlementDetails)    //主播申请结算明细
		user.POST("/isbindingbank", control.IsBindingBank)           //判断是否绑定银行卡
		user.POST("/isenough", control.IsEnough)                     //判断申请结算数量是否足够
		user.POST("/ismonthcashing", control.IsMonthCashing)         //判断主播是否已经申请结算
		user.POST("/getanchortime", control.GetAnchorRoomTimeList)   //直播记录
		user.POST("/IsSigned", control.IsSigned)                     // 查询是否今天签到过
		user.POST("/findByBenefactor", control.FindByBenefactor)
		user.POST("/findByRecipient", control.FindByRecipient) //根据Recipient查询刷礼物记录
		user.POST("/userUpFavicon", control.UserUpFavicon)     //修改用户头像
		user.POST("/categoryUpload", control.CategoryUpload)   //上传类目图片
		user.POST("/slzPassUp", control.SlzPassUp)             //赠送石榴籽密码验证
	}
	root := router.Group("/root", control.IsRootAccess())
	{
		root.GET("/*action", control.GetActionRoot)
		root.POST("/getsidebar", control.GetSidebar)                 //获取网站后台侧边栏 cnxulin//后台返回格式已修改
		root.POST("/getroomconcernlist", control.GetRoomConcernList) //获取所有关注记录
		root.POST("/getOfficialList", control.GetOfficialList)
		root.POST("/officialadd", control.OfficialAdd)
		root.POST("/officialdel", control.OfficialDel)
		root.POST("/officialup", control.OfficialUp)
		root.POST("/unlogin", control.UnLogin) //退出登录,后台返回格式已修改
		root.POST("/userdel", control.UserDel) //后台删除用户
		root.POST("/roomadd", control.RoomAdd)
		root.POST("/roomdel", control.RoomDel)
		root.POST("/roomup", control.RoomUp)
		root.POST("/chooseroom", control.ChooseRoom)   //筛选房间
		root.POST("/getroomlist", control.GetRoomList) //后台房间列表
		root.POST("/productadd", control.ProductAdd)   //添加一期（添加期数产品过程表）
		root.POST("/productdel", control.ProductDel)
		root.POST("/productup", control.ProductUp)
		root.POST("/getproductlist", control.GetProductList)     //后台获取产品列表
		root.POST("/getmyproductlist", control.GetMyProductList) //后台checkbox调用
		root.POST("/getnowperprolist", control.GetPerProList)    //Phone端获取当前期数所有的产品   房管选择产品
		root.POST("/supportdel", control.SupportDel)
		root.POST("/supportup", control.SupportUp)
		root.POST("/getsupport", control.GetSupport)
		root.POST("/getsupportlist", control.GetSupportList)       //后台支持列表获取
		root.POST("/getapplylist", control.GetApplyList)           //后台申请列表
		root.POST("/getApplyLists", control.GetApplyLists)         //后台申请列表---->搜索
		root.POST("/getapplyinfo", control.GetApplyInfo)           //获取某一条申请详情
		root.POST("/applyrefused", control.ApplyRefused)           //拒绝申请
		root.POST("/applyarg", control.ApplyArg)                   //同意申请
		root.POST("/getcarpiclist", control.GetCarPicLists)        //获取轮播图列表（总表后台
		root.POST("/periodadd", control.PeriodAdd)                 //添加期数
		root.POST("/perioddel", control.PeriodDel)                 //删除期数
		root.POST("/periodup", control.PeriodUP)                   //更新期数
		root.POST("/periodendcodingup", control.PeriodEndCodingUP) //投注结果更新 State=1
		root.POST("/getperiodlist", control.GetPeriodList)         //后台期数列表
		root.POST("/getrootuserlist", control.GetRootUserList)     //感觉无用，没改
		root.POST("/getproduct", control.GetProduct)               //获取单个产品
		root.POST("/getAnchorInfos", control.GetAnchorInfos)       //获取后台主播信息
		root.POST("/anchorBalance", control.AnchorBalance)         //后台主播结算,后台返回已修改
		root.POST("/periodidlist", control.GetPeriodIdProduct)     // 获取某一期所有的竞猜产品(期过程属性表)
		root.POST("/rootuseradd", control.RootAddUser)             //后台添加用户
		root.POST("/giftadd", control.GiftAdd)
		root.POST("/giftdel", control.GiftDel)
		root.POST("/giftup", control.GiftUp)
		root.POST("/findGiveAll", control.FindGiveAll)               //查询所有赠送礼物
		root.POST("/affiliationadd", control.AffiliationAdd)         //添加机构
		root.POST("/affiliationdel", control.AffiliationDel)         //删除机构
		root.POST("/affiliationup", control.AffiliationUp)           //修改机构
		root.POST("/getaffiliationlist", control.GetAffiliationList) //列表信息
		root.POST("/carpicadd", control.CarPicAdd)                   //添加轮播图
		root.POST("/carpicdel", control.CarPicDel)                   //删除轮播图
		root.POST("/carpicup", control.CarPicUp)                     //修改轮播图
		root.POST("/setsite", control.SetSite)                       //信息设置，后台返回已修改
		root.POST("/getsite", control.GetSite)                       //获取网站信息，后台返回已修改
		root.POST("/getpaylist", control.GetPayList)                 //充值记录//后台查询暂不修改
		root.POST("/peraccounting", control.PerAccounting)           //期核算
		root.POST("/getallfundDesc", control.GetAllFundDesc)
		root.POST("/userorderbylevelADsc", control.UserOrderByLevelADSC)
		root.POST("/userorderbysearch", control.UserOrderBySearch)
		root.POST("/usersearchtype", control.SearchOfType)     //筛选  等级与权限,未找到使用的地方，暂留
		root.POST("/geteventlist", control.GetEventList)       //活动列表
		root.POST("/eventadd", control.EventAdd)               //添加活动
		root.POST("/eventup", control.EventUp)                 //修改活动
		root.POST("/eventover", control.EventOver)             //测试活动截止
		root.POST("/geteventrecinfo", control.GetEventRecInfo) //活动详情
		root.POST("/earningsdetails", control.EarningsDetails) //获取一期中每个产品的支持热度值 @author xulin
		root.POST("/getRootPlayList", control.GetRootPlayList) //后台直播列表
		root.POST("/getblocklist", control.GetBlockList)       //禁止推流列表
		root.POST("/updateOrgan", control.UpdateOrgan)
		root.POST("/findOrgan", control.FindOrgan)
		root.POST("/delOrgan", control.DelOrgan)
		root.POST("/addOrgan", control.AddOrgan)
		root.POST("/addBroadcast", control.AddBroadcast)
		root.POST("/findOrganById", control.FindOrganById)
		root.POST("/sumOrganManageByOrganId", control.SumOrganManageByOrganId)
		root.POST("/findOrganManageByOrganId", control.FindOrganManageByOrganId)
		root.POST("/getgift", control.GetGift) //获取礼物
		root.POST("/getgiftlist", control.GetGiftList)
		root.POST("/getnumberlist", control.GetNumberList)                       //后台  赠送石榴籽列表
		root.POST("/addBroadcastRedis", control.AddBroadcastRedis)               //添加广播信息
		root.POST("/getuidtypelist", control.GetUidTypeList)                     //权限列表
		root.POST("/uidtypeadd", control.UidTypeAdd)                             //权限添加
		root.POST("/uidtypeup", control.UidTypeUp)                               //权限修改
		root.POST("/getauthoritylist", control.GetAuthorityList)                 //sitebar表数据
		root.POST("/typeprocessadd", control.TypeProcessAdd)                     //修改或添加 人物 权限过程
		root.POST("/getmyaulist", control.GetMyAuList)                           //权限的过程  是否存在
		root.POST("/getwebearning", control.GetWebEarnings)                      //平台收益情况，后台返回已修改
		root.POST("/getwebinfo", control.GetWebInfoAll)                          //运营总数据，后台返回已修改
		root.POST("/regnumber", control.RegNumber)                               //平台运营信息，后台返回已修改
		root.POST("/getCategoryList", control.GetCategoryList)                   //一级类目列表 addby liuhan
		root.POST("/categoryAdd", control.CategoryAdd)                           //增加一级类目 addby liuhan
		root.POST("/categoryDelete", control.CategoryDelete)                     //删除一级类目 addby liuhan
		root.POST("/getTwoCategoryByOneId", control.GetTwoCategoryByOneId)       //根据一级类目id，查询对应的所有二级类目 addby liuhan
		root.POST("/getTwoCategoryList", control.GetTwoCategoryList)             //查询所有二级类目 addby liuhan
		root.POST("/categoryTwoDelete", control.CategoryTwoDelete)               //删除二级类目 addby liuhan
		root.POST("/categoryTwoAdd", control.CategoryTwoAdd)                     //增加二级类目 addby liuhan
		root.POST("/categoryTwoAdds", control.CategoryTwoAdds)                   //增加二级类目 addby liuhan
		root.POST("/getOneTwoCategoryByOneId", control.GetOneTwoCategoryByOneId) //根据一级类目id，查询对应的所有二级类目 addby liuhan
		root.POST("/deleteCategoryTwo", control.DeleteCategoryTwo)               //删除二级类目 addby liuhan 170124
		root.POST("/getDbadvertisinglist", control.GetDbadvertisinglist)
		root.POST("/dbadvertisingAdd", control.DbadvertisingAdd)         //添加直播间广告
		root.POST("/dbadvertisingDel", control.DbadvertisingDel)         //删除直播间广告
		root.POST("/dbadvertisingUp", control.DbadvertisingUp)           //修改直播间广告
		root.POST("/getSdadvertisinglist", control.GetSdadvertisinglist) //左侧导航
		root.POST("/sdbadvertisingAdd", control.SdbadvertisingAdd)
		root.POST("/sdbadvertisingDel", control.SdbadvertisingDel)
		root.POST("/sdbadvertisingUp", control.SdbadvertisingUp)
	}
	page := router.Group("/page")
	{
		page.GET("/*action", control.GetActionPage)
		page.POST("/getanchorinfo", control.GetAnchorInfo)               //直播页面
		page.POST("/getuser", control.GetUser)                           //获取用户信息,后台返回格式已修改
		page.POST("/getusername", control.GetUserName)                   //获取用户信息 根据username,无用，没改
		page.POST("/getuserlist", control.GetUserList)                   //用户搜索
		page.POST("/getsearchlist", control.GetSearchList)               //  手机端
		page.POST("/getafflist", control.GetAllAffiliation)              //  注册时机构列表
		page.POST("/getIndexCarPicList", control.GetIndexCarPicList)     //获取轮播图列表
		page.POST("/getcarpiclist", control.GetCarPicLists)              //*手机端首页轮播*/
		page.POST("/getroom", control.GetRoom)                           //获取房间单个
		page.POST("/getroomtypelist", control.GetRoomTypeListIng)        //正在直播 分页
		page.POST("/byphonereg", control.ByPhoneRegPhone)                //手机注册  验证手机是否注册过
		page.POST("/isverification", control.IsVerification)             //手机注册  验证码是否正确
		page.POST("/disanfangdenglu", control.DiSanFangDengLu)           //第三方登陆
		page.POST("/getOfficial", control.GetOfficial)                   //获取单个官方活动信息
		page.POST("/getOfficialList", control.GetOfficialList)           //获取官方信息列表
		page.POST("/getStartOfficialList", control.GetStartOfficialList) //获取显示活动信息
		page.POST("/getproduct", control.GetProduct)                     //获取单个产品
		page.POST("/loginrootmyadmin", control.LoginRootMyAdmin)         //后台登陆,后端返回格式已修改
		page.POST("/getIndexroomlist", control.GetIndexRoomList)         // 手机端精彩推荐
		page.POST("/getroomlisting", control.GetALiRoomInfo)             //正在直播的房间  2016-11-07
		page.POST("/getListroomlist", control.GetListRoomList)           //列表页
		page.POST("/passmate", control.PassMate)                         //验证密码是否匹配
		page.POST("/selroomalias", control.SelRoomAlias)                 //根据别名查询(模糊)
		page.POST("/getperiod", control.GetPeriod)                       //获取某一期
		page.POST("/getgift", control.GetGift)                           //获取礼物
		page.POST("/getgiftlist", control.GetGiftList)
		page.POST("/perproname", control.PerProName)                         //一期所有的产品及名称
		page.POST("/periodmores", control.PeriodPhoneMore)                   //Phone更多期.
		page.POST("/periodidlist", control.GetPeriodIdProduct)               // 获取某一期所有的竞猜产品(期过程属性表)
		page.POST("/getperiodidproductname", control.GetPeriodIdProductName) //根据期ID-->显示期ID、产品ID、产品名称（Phone）
		page.POST("/periodinfo", control.CurrentPeriodBase)                  //最新当前期----房间
		page.POST("/gettimeperprolist", control.CurrentPeriodBase)           //web当前期
		page.POST("/getnowperioddetails", control.CurrentPeriodDetails)      //获取当前期产品明细
		page.POST("/periodmore", control.PeriodMore)                         //获取更多期
		page.POST("/sms", control.MobileSms)                                 //手机端连接短信服务器
		page.POST("/verificationCode", control.VerificationCode)             //fushou短信验证 addby liuhan 170122
		page.POST("/mobilesms", control.MobileSmS)                           //手机端连接短信服务器 (通用) 忘记密码 2016_10_12_徐
		page.POST("/checkcode", control.CheckCode)                           //判断手机端输入的验证码与数据库是否相同 2016_10_12_徐
		page.POST("/userup", control.UserUp)
		page.POST("/getgiftgiveall", control.GetGiftGiveMonths)  //总榜
		page.POST("/getgiftgiveweeks", control.GetGiftGiveWeeks) //周榜
		page.POST("/updateOrgan", control.UpdateOrgan)
		page.POST("/findOrgan", control.FindOrgan)
		page.POST("/delOrgan", control.DelOrgan)
		page.POST("/addOrgan", control.AddOrgan)
		page.POST("/addBroadcast", control.AddBroadcast)
		page.POST("/findOrganById", control.FindOrganById)
		page.POST("/sumOrganManageByOrganId", control.SumOrganManageByOrganId)
		page.POST("/findOrganManageByOrganId", control.FindOrganManageByOrganId)
		page.POST("/qQLogin", control.QQLogin)
		page.POST("/wBLogin", control.WBLogin)
		page.POST("/weiXinUserInfo", control.WeiXinUserInfo) //微信code登陆
		page.POST("/qQUserInfo", control.QQUserInfo)         //QQ code登陆
		page.POST("/wBGetUserInfo", control.WBGetUserInfo)   //微博code登陆
		page.POST("/phoneLogin", control.PhoneLogin)
		page.POST("/regval", control.RegVal)                                   //手机注册--验证手机是否注册过
		page.POST("/isvercode", control.IsVerCode)                             //手机注册--查看输入验证码是否正确
		page.POST("/reg", control.Reg)                                         //手机注册--注册
		page.POST("/checknick", control.CheckNick)                             //手机注册--查看昵称是否被使用
		page.POST("/nick", control.Nick)                                       //手机注册--填写昵称
		page.POST("/log", control.Log)                                         //手机登录
		page.POST("/loseregval", control.LoseRegval)                           //忘记密码--发验证码
		page.POST("/resetpass", control.ResetPass)                             //忘记密码--重置密码
		page.POST("/getplugflow", control.GetPlugFlow)                         //主播串流码
		page.POST("/getinflow", control.GetInFlow)                             //主播地址返回
		page.POST("/imageresizerUpload", control.ImageresizerUpload)           //头像上传
		page.POST("/byPhoneBindEditPhone", control.ByPhoneBindEditPhone)       //修改绑定手机--发送验证码  -->通过手机号
		page.POST("/findCategoryAll", control.FindCategoryAll)                 //查询分类
		page.POST("/getTwoCategoryByOneIds", control.GetTwoCategoryByOneIds)   //根据一级类目id，查询对应的所有二级类目 addby liuhan
		page.POST("/getTwoCategoryByAddress", control.GetTwoCategoryByAddress) //根据二级类目地址查询 addby liuhan.
		page.POST("/getRoomAliasByRoomType", control.GetRoomAliasByRoomType)   //根据分类查询 addby liuhan
		page.POST("/getTwoCategoryList", control.GetTwoCategoryList)           //查询所有二级类目 addby liuhan
		page.POST("/getCategoryList", control.GetCategoryList)                 //一级类目列表 addby liuhan
		page.POST("/getTwoCategoryRoom", control.GetTwoCategoryRoom)           //查询所有二级分类下的房间 addby liuhan 161226
		page.POST("/getHotLiveList", control.GetHotLiveList)                   // 热门直播8个 addby liuhan 161227
		page.POST("/getIndexCarPicLists", control.GetIndexCarPicLists)
		page.POST("/getroombyrid", control.GetRoomByRId) //获取房间信息
		page.POST("/test1", control.TestGiftNumber)
		page.POST("/givegiftnumadd", control.GiveGiftNumAdd)             //礼物赠送
		page.POST("/getDbadvertisinglist", control.GetDbadvertisinglist) //获取直播广告列表
		page.POST("/getDbadvertising", control.GetDbadvertising)
		page.POST("/getSdadvertisinglist", control.GetSdadvertisinglist) //获取直播广告列表
		page.POST("/getSdbadvertising", control.GetSdbadvertising)
		page.POST("/getroomim", control.GetRoomIm) //连接房间IM
		page.POST("/setUp", control.SetUp)
		page.POST("/addsetUp", control.AddSetUp)
	}
	help := router.Group("/help")
	{
		help.GET("/*action", control.GetActionHelp)
	}
	cate := router.Group("/cate")
	{
		cate.GET("/*action", control.GetActionCate)
	}

	roomlive := router.Group("/roomlive")
	{
		roomlive.GET("/*action", control.GetActionRoomLive)
	}
	outlive := router.Group("/outlive")
	{
		outlive.GET("/*action", control.GetActionOutLive)
	}
	router.Run(":80")
}
