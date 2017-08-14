$(function () {
	pageinit()
});
//页面初始化
function pageinit() {

	if (getStorage('Id')) {

		reqAjax("/page/getuser", { UID: getStorage('Id') }, function (data) {
			console.log(data);

			$('.comm-nickname').html(data.Data.NickName);
			$('.comm-level').addClass("level-" + userlevel(data.Data.Integral).number);
			if (data.Data.Favicon) {
				$('.comm-avatar').attr('src', data.Data.Favicon);
			} else {
				$('.comm-avatar').attr('src', '/static/images/default_avatar.jpg');
			}
			$('.comm-remain-money').html(data.Data.Balance);
			$('.comm-remain-nums').html(data.Data.PomegranateNum);
			$('.comm-level-number').html(userintegral(data.Data.Integral));
		});
		var uname = getStorage('nicheng') ? getStorage('nicheng') : getStorage('username'), Integral = getStorage('Integral')
		//侧边栏
		$('.sidebar-collapse-user,.sidebar-login-boxS').css('display', 'block');
		$('.sidebar-box-login,.sidebar-collapse-logions').css('display', 'none');
		$('.sidebar-box-login').css('display', 'none');
		$('.sidebar-box-ladingState').css('display', 'block');
		userOff('.out_head');
		//头部样式
		$('.header-loginOff').css('display', 'none');
		$('.header-loginUserImg').css('display', 'block');
		$('.header-tool-user-info').css('padding', '10px 0px');

	} else {
		$('.sidebar-collapse-user,.sidebar-login-boxS').css('display', 'none');
		$('.sidebar-box-login,.sidebar-collapse-logions').css('display', 'block');
		$('.sidebar-box-login').css('display', 'block');
		$('.sidebar-box-ladingState').css('display', 'none');
		//头部样式
		$('.header-loginUserImg').css('display', 'none');
		$('.header-loginOff').css('display', 'block');
	}
	$(".sidebar-search-btn,.search-submit").on("click", function () {
		var headerSearch = $(".search-key").val();
		if (headerSearch && $.trim(headerSearch) != null && $.trim(headerSearch) != '') {
			setStorage("cxzbpage", 1);
			$('.search-form').submit();
		} else {
			Dialog('请输入搜索内容');
		}
	});
	//按下回车搜索
	$('.search-key').bind('keypress', function (e) {
		if (e.keyCode == "13") {
			$('.search-submit').click();
		}
	});
}

var username = document.cookie.split(";")[0].split("=")[1];

//获取url参数
function getUrlpramas(name) {
	var urlstr = window.location.search.substr(1);
	var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
	var r = urlstr.match(reg);
	if (r != null) {
		return unescape(r[2]);
	} else {
		return null;
	}
}

//获取?后的参数
function getstrpramas(str, name) {
	var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
	var r = str.match(reg);
	if (r != null)
		return unescape(r[2]);
	return null;
}

//获取本地存储
function getStorage(key) {
	if (key == 'Id' || key == 'Rid' || key == 'uid') {
		return getcookie(key);
	} else {
		return window.localStorage.getItem(key);
	}
}

//设定本地存储
function setStorage(key, val) {
	if (key == 'Id' || key == 'Rid' || key == 'uid') {
		var date = new Date();
		date.setTime(date.getTime() + 60 * 60 * 1000);
		document.cookie = key + '=' + val + ';expires=' + date.toGMTString() + ';path=/';
		timercookie(key, val);
	} else {
		window.localStorage.setItem(key, val);
	}
}

//清除本地存储
function removeStorage(key) {
	window.localStorage.removeItem(key);
	if (key == 'Id' || key == 'Rid' || key == 'uid') {
		var date = new Date();
		date.setTime(date.getTime() - 1000);
		document.cookie = key + '=' + getcookie('' + key + '') + ';expires=' + date.toGMTString() + ';path=/';
	}
}

//获取cookie
function getcookie(key) {
	var arr, reg = new RegExp("(^| )" + key + "=([^;]*)(;|$)");
	if (arr = document.cookie.match(reg))
		return unescape(arr[2]);
	else
		return null;
}

//设定cookie
function setcookie(key, value, time) {
	var exp = new Date();
	exp.setTime(exp.getTime() + time);
	document.cookie = key + "=" + escape(value) + ";expires=" + exp.toGMTString();
}

//删除cookie
function delcookie(key) {
	var exp = new Date();
	exp.setTime(exp.getTime() - 1);
	var cval = getCookie(key);
	if (cval != null)
		document.cookie = key + "=" + cval + ";expires=" + exp.toGMTString();
}

//html标记对转换为转义符
function checkData(v) {
	var entry = { "'": "&apos;", '"': '&quot;', '<': '&lt;', '>': '&gt;' };
	v = v.replace(/(['")-><&\\\/\.])/g, function ($0) { return entry[$0] || $0; });
	return v;
}

//表情转义符转换为图片
function replace_em(str) {
	str = str.replace(/\[em_([0-9]*)\]/g, '<i class="emoji_img" style="background:url(/static/arclist/$1.gif)"></i>');
	return str;
}

/*
 * Api接口请求
 * @param_1 url(string) 请求地址
 * @param_2 params(obj) 请求参数
 * @param_3 callbak(function) 回调函数
 * @param_4 asyncbool(bool) 是否为异步处理
*/
function reqAjax(url, params, callbak, asyncbool) {
	if (!asyncbool) {
		asyncbool = false;
	}
	$.ajax({
		type: "post",
		async: asyncbool,
		url: url,
		dataType: "json",
		data: params,
		success: function (data) {
			callbak(data);
		},
		error: function (xhr) {
			//			console.log(xhr,url);
			//Dialog('服务连接异常，请稍后重试');
		}
	});
}

/*
 * 自定义会话窗口（alert）
 * @param_1 prompt(string) 提示语
 * @param_2 mask(bool) 是否需要遮罩层
 * @param_3 oktbn(string) 确认按钮（左侧按钮）文字
 * @param_4 cancelbtn(string) 取消按钮（右侧按钮）文字  （若只显示一个按钮时该参数值传null或空字符串）
 * @param_5 okcallfun(function) 点击确认按钮（左侧按钮）后的执行方法
 * @param_6 cancelcallfun(function) 点击取消按钮（右侧按钮）后的执行方法（若只显示一个按钮时该参数值传null或空字符串）

*/
function Dialog(prompt, mask, oktbn, cancelbtn, okcallfun, cancelcallfun) {
	var randomid = $.now();
	var okbtnid = 'ok' + randomid;
	var cancelbtnid = 'cancel' + randomid;
	var mod = '<div class="dialog">';
	if (mask == true) {
		mod += '<div class="dialog_mask"></div>';
	}
	mod += '<div class="dialog_wrap">';
	mod += '<div class="dialog_wrap_box">';
	mod += '<div class="dialog_wrap_body">';
	mod += '<div class="dialog_wrap_content">' + prompt + '</div>';
	mod += '<div class="dialog_wrap_button">';
	mod += '<button class="dialog_wrap_button_ok" id="' + okbtnid + '">' + (oktbn ? oktbn : '确认') + '</button>';
	if (cancelbtn) {
		mod += '<button class="dialog_wrap_button_cancel" id="' + cancelbtnid + '">' + cancelbtn + '</button>';
	}
	mod += '</div>';
	mod += '</div>';
	mod += '</div>';
	mod += '</div>';
	mod += '</div>';
	$('body').append(mod);
	if (okcallfun) {
		$('#' + okbtnid).on("click", function () {
			if (jQuery.isFunction(okcallfun)) {
				okcallfun($(this));
			} else {

			}
		});
	} else {
		$('#' + okbtnid).on("click", function () {
			$(this).parents('.dialog').remove();
		});
	}

	if (cancelcallfun) {
		$('#' + cancelbtnid).on("click", function () {
			if (jQuery.isFunction(cancelcallfun)) {
				cancelcallfun($(this));
			} else {

			}
		});
	} else {
		$('#' + cancelbtnid).on("click", function () {
			$(this).parents('.dialog').remove();
		});
	}
}




/*登出*/
function userOff(clickName, Refresh) {
	$(clickName).click(function () {
		Dialog('确定要退出登录吗', true, '确定', '取消', function () {
			var id = getStorage("Id");
			$.ajax({
				type: "post",
				url: "/user/unlogin",
				async: true,
				data: {
					UID: id
				},
				success: function (msg) {
					if (msg.flag == true) {
						removeStorage("Id");
						removeStorage("nicheng");
						removeStorage("username");
						removeStorage("Integral");
						removeStorage("Favicon");
						removeStorage("from");
						if (location.pathname.indexOf('/user') != -1) {
							window.location.href = '/';
						} else {
							location.reload();
						}
					} else {
						removeStorage("Id");
						removeStorage("nicheng");
						removeStorage("username");
						removeStorage("Integral");
						removeStorage("Favicon");
						removeStorage("from");
						if (location.pathname.indexOf('/user') != -1) {
							window.location.href = '/';
						} else {
							location.reload();
						}
					}
				},
				error: function (msg) {
					removeStorage("Id");
					removeStorage("nicheng");
					removeStorage("username");
					removeStorage("Integral");
					removeStorage("Favicon");
					removeStorage("from");
					if (location.pathname.indexOf('/user') != -1) {
						window.location.href = '/';
					} else {
						location.reload();
					}
					// if(Refresh){
					// 	window.location.href = Refresh;
					// }else{
					// 	location.reload();
					// }
				}
			});
		}, function (e) {
			e.parents('.dialog').remove();
		});
	});
}
/*弹框*/
//yes OR no弹出框
//alertid:需要显示弹出框的 子节点名称
//titleFont：弹出框的标题文字内容
//alertFont：弹出框显示的文字
//oktbn：确定按钮文字 默认“确定”
//cancelbtn：取消按钮文字 null为默认单选

//function alertShowYN(alertid,titleFont,alertFont){
//	var altstr='<div class="alert-bg-box"><div class="alert-Box"><div class="alert-title"><p>'+titleFont+'</p><span class="alert-close">X</span></div><div class="alert-content"><div class="alert-content-font">'+alertFont+'</div><ul class="alert-yAndN"><li class="alert-yes">确定</li><li class="alert-no">取消</li><div style="clear:both"></div></ul></div></div></div>'
//	$(alertid).parents('.mySetting-table-content').append(altstr);
//	$('.alert-close,.alert-yes,.alert-no').click(function(){
//		$('.alert-bg-box').css('display','none');
//	})
//}
////弹出框
//function alertShow(alertid,titleFont,alertFont){
//	var altstr='<div class="alert-bg-box"><div class="alert-Box"><div class="alert-title"><p>'+titleFont+'</p><span class="alert-close">X</span></div><div class="alert-content"><div class="alert-content-font">'+alertFont+'</div><div class="alert-OK">确定</div></div></div></div>'
//	$(alertid).parents('.mySetting-table-content').append(altstr);
//	$('.alert-close,.alert-yes,.alert-no,.alert-OK').click(function(){
//		$('.alert-bg-box').css('display','none');
//	})
//}
//

function alertShowYNznx(titleFont, alertFont, oktbn, cancelbtn) {
	var altstr = '<div class="alert-bg-box"><div style="display: table-cell;vertical-align: middle;"><div class="alert-Box">'
	altstr += '<div class="alert-title"><p>' + titleFont + '</p><span class="alert-close">×</span></div><div class="alert-content">'
	altstr += '<div class="alert-content-font">' + alertFont + '</div>';
	if (cancelbtn) {
		altstr += '<ul class="alert-yAndN"><li class="alert-yes">' + (oktbn ? oktbn : '确定') + '</li>';
		altstr += '<li class="alert-no">' + cancelbtn + '</li>';
	} else {
		altstr += '<ul class="alert-yAndN"><li class="alert-OK" style="margin-left: 57px;">' + (oktbn ? oktbn : '确定') + '</li>';
	}
	altstr += '<div style="clear:both"></div></ul></div></div></div></div>';
	$("body").append(altstr);
	$('.alert-close,.alert-yes,.alert-no,.alert-OK').click(function () {
		$('.alert-bg-box').remove();
	})
}
//uzi
function authAlert() {
	if(window.location.pathname.indexOf("/root/") >= 0){
		return jQuery.messager.show({title:"提示",msg:"此版本仅供学习，禁止商用"});
	}else{
		var randomid = $.now();
		var okbtnid = 'ok' + randomid;
		var mod = '<div class="dialog">';
		mod += '<div class="dialog_mask"></div>';
		mod += '<div class="dialog_wrap">';
		mod += '<div class="dialog_wrap_box">';
		mod += '<div class="dialog_wrap_body">';
		mod += '<div class="dialog_wrap_content">此版本仅供学习，禁止商用</div>';
		mod += '<div class="dialog_wrap_button">';
		mod += '<button class="dialog_wrap_button_ok" id="' + okbtnid + '">确认</button>';
		mod += '</div>';
		mod += '</div>';
		mod += '</div>';
		mod += '</div>';
		mod += '</div>';
		$('body').append(mod);
		$('#' + okbtnid).on("click", function () {
			$(this).parents('.dialog').remove();
		});
	}
}

function userlevel(val) {
	var level = {};
	if (val <= 99) { level.name = "白板"; level.number = 0; level.tiao = val; level.dvalue = 1; level.next = "青铜5"; }
	if (val >= 100 && val < 200) { level.name = "青铜5"; level.number = 1; level.tiao = val - 100; level.dvalue = 100; level.next = "青铜4"; }
	if (val >= 200 && val < 300) { level.name = "青铜4"; level.number = 2; level.tiao = val - 200; level.dvalue = 100; level.next = "青铜3"; }
	if (val >= 300 && val < 500) { level.name = "青铜3"; level.number = 3; level.tiao = val - 300; level.dvalue = 200; level.next = "青铜2"; }
	if (val >= 500 && val < 800) { level.name = "青铜2"; level.number = 4; level.tiao = val - 500; level.dvalue = 300; level.next = "青铜1"; }
	if (val >= 800 && val < 1000) { level.name = "青铜1"; level.number = 5; level.tiao = val - 800; level.dvalue = 200; level.next = "白银5"; }
	if (val >= 1000 && val < 2000) { level.name = "白银5"; level.number = 6; level.tiao = val - 1000; level.dvalue = 1000; level.next = "白银4"; }
	if (val >= 2000 && val < 3000) { level.name = "白银4"; level.number = 7; level.tiao = val - 2000; level.dvalue = 1000; level.next = "白银3"; }
	if (val >= 3000 && val < 5000) { level.name = "白银3"; level.number = 8; level.tiao = val - 3000; level.dvalue = 2000; level.next = "白银2"; }
	if (val >= 5000 && val < 8000) { level.name = "白银2"; level.number = 9; level.tiao = val - 5000; level.dvalue = 3000; level.next = "白银1"; }
	if (val >= 8000 && val < 10000) { level.name = "白银1"; level.number = 10; level.tiao = val - 8000; level.dvalue = 2000; level.next = "黄金5"; }
	if (val >= 10000 && val < 20000) { level.name = "黄金5"; level.number = 11; level.tiao = val - 10000; level.dvalue = 10000; level.next = "黄金4"; }
	if (val >= 20000 && val < 30000) { level.name = "黄金4"; level.number = 12; level.tiao = val - 20000; level.dvalue = 10000; level.next = "黄金3"; }
	if (val >= 30000 && val < 50000) { level.name = "黄金3"; level.number = 13; level.tiao = val - 30000; level.dvalue = 20000; level.next = "黄金2"; }
	if (val >= 50000 && val < 80000) { level.name = "黄金2"; level.number = 14; level.tiao = val - 50000; level.dvalue = 30000; level.next = "黄金1"; }
	if (val >= 80000 && val < 100000) { level.name = "黄金1"; level.number = 15; level.tiao = val - 80000; level.dvalue = 20000; level.next = "铂金5"; }
	if (val >= 100000 && val < 200000) { level.name = "铂金5"; level.number = 16; level.tiao = val - 100000; level.dvalue = 100000; level.next = "铂金4"; }
	if (val >= 200000 && val < 300000) { level.name = "铂金4"; level.number = 17; level.tiao = val - 200000; level.dvalue = 100000; level.next = "铂金3"; }
	if (val >= 300000 && val < 500000) { level.name = "铂金3"; level.number = 18; level.tiao = val - 300000; level.dvalue = 200000; level.next = "铂金2"; }
	if (val >= 500000 && val < 800000) { level.name = "铂金2"; level.number = 19; level.tiao = val - 500000; level.dvalue = 300000; level.next = "铂金1"; }
	if (val >= 800000 && val < 1000000) { level.name = "铂金1"; level.number = 20; level.tiao = val - 800000; level.dvalue = 200000; level.next = "钻石5"; }
	if (val >= 1000000 && val < 2000000) { level.name = "钻石5"; level.number = 21; level.tiao = val - 1000000; level.dvalue = 1000000; level.next = "钻石4"; }
	if (val >= 2000000 && val < 3000000) { level.name = "钻石4"; level.number = 22; level.tiao = val - 2000000; level.dvalue = 1000000; level.next = "钻石3"; }
	if (val >= 3000000 && val < 5000000) { level.name = "钻石3"; level.number = 23; level.tiao = val - 3000000; level.dvalue = 2000000; level.next = "钻石2"; }
	if (val >= 5000000 && val < 8000000) { level.name = "钻石2"; level.number = 24; level.tiao = val - 5000000; level.dvalue = 3000000; level.next = "钻石1"; }
	if (val >= 8000000 && val < 10000000) { level.name = "钻石1"; level.number = 25; level.tiao = val - 8000000; level.dvalue = 2000000; level.next = "大师"; }
	if (val >= 10000000 && val < 20000000) { level.name = "大师"; level.number = 26; level.tiao = val - 10000000; level.dvalue = 10000000; level.next = "王者"; }
	if (val >= 20000000) { level.name = "王者"; level.number = 27; level.tiao = val - 20000000; level.dvalue = ""; }
	return level;
}
function userintegral(val) {
	var level = {};
	if (val <= 99) { return 100 - val }
	if (val >= 100 && val < 200) { return 200 - val }
	if (val >= 200 && val < 300) { return 300 - val }
	if (val >= 300 && val < 500) { return 500 - val }
	if (val >= 500 && val < 800) { return 800 - val }
	if (val >= 800 && val < 1000) { return 1000 - val }
	if (val >= 1000 && val < 2000) { return 2000 - val }
	if (val >= 2000 && val < 3000) { return 3000 - val }
	if (val >= 3000 && val < 5000) { return 5000 - val }
	if (val >= 5000 && val < 8000) { return 8000 - val }
	if (val >= 8000 && val < 10000) { return 10000 - val }
	if (val >= 10000 && val < 20000) { return 20000 - val }
	if (val >= 20000 && val < 30000) { return 30000 - val }
	if (val >= 30000 && val < 50000) { return 50000 - val }
	if (val >= 50000 && val < 80000) { return 80000 - val }
	if (val >= 80000 && val < 100000) { return 10000 - val }
	if (val >= 100000 && val < 200000) { return 200000 - val }
	if (val >= 200000 && val < 300000) { return 300000 - val }
	if (val >= 300000 && val < 500000) { return 500000 - val }
	if (val >= 500000 && val < 800000) { return 800000 - val }
	if (val >= 800000 && val < 1000000) { return 1000000 - val }
	if (val >= 1000000 && val < 2000000) { return 2000000 - val }
	if (val >= 2000000 && val < 3000000) { return 3000000 - val }
	if (val >= 3000000 && val < 5000000) { return 5000000 - val }
	if (val >= 5000000 && val < 8000000) { return 8000000 - val }
	if (val >= 8000000 && val < 10000000) { return 10000000 - val }
	if (val >= 10000000 && val < 20000000) { return 20000000 - val }
	if (val >= 20000000) { return 200 - val }
}
function bindchangeMobile(pho) {

	var btn = $(this);
	var count = 60;
	$('#bindcheck_num,#bindphCheck').css('display', 'block');

	var resend = setInterval(function () {
		count--;
		if (count > 0) {

			$('#bindphCheck').css('display', 'block');
			$("#bindgetting").css({ "float": "left", "width": "170px" });
			$("#bindgetting").val(count + "秒后可重新获取").attr('disabled', 'disabled');;
			$("#bindgetting").css({ 'cursor': 'not-allowed' });

		} else {
			clearInterval(resend);
			$('#bindnewphone,#bindordpass').removeAttr('disabled');
			$("#bindgetting").css({ "float": "right", "width": "102px" });
			$("#bindgetting").val("再次发送").removeAttr('disabled').removeAttr('disabled');
			$("#bindgetting").css({ 'cursor': 'pointer' });
		}
	}, 1000);
	btn.attr('disabled', true).css({
		'cursor': 'not-allowed',
		'border': '1px solid gainsboro'
	});
	$('#bindchNum').css('display', 'block');
	//	reqAjax("/page/byPhoneBindEditPhone",{mobile: pho},function(msg){
	//		if(msg.ErrorCode!=0) {
	//			alertShowYNznx("提示", msg.ErrorMsg, null);
	//		}else {
	//			alertShowYNznx("提示", "发送成功", null);
	//			$("#bindnewphone").attr("disabled", true);
	//		}
	//	},true); 
}

function bindPhone(phone, newpass) {
	var uid = getStorage("Id");
	reqAjax("/user/userup", { UID: uid, Phone: phone, PassWord: newpass }, function (msg) {
		if (msg.ErrorCode != 0) {
			alertShowYNznx("提示", msg.ErrorMsg, null);
			$(".alert-OK,.alert-close").bind("click", function () {
				location.reload();
			});
		} else {

			alertShowYNznx("提示", "恭喜您成功绑定手机号。你绑定的手机号为：" + phone, null);
			$(".alert-OK,.alert-close").bind("click", function () {
				location.reload();
			});

		}
	}, true);
}
function timercookie(key, val) {
	setInterval(function () {
		var date = new Date();
		date.setTime(date.getTime() + 600 * 60 * 1000);
		document.cookie = key + '=' + val + ';expires=' + date.toGMTString() + ';path=/';
		console.log("timercookie>>>>>>>>>>>>coo");
	}, 10 * 60 * 1000);
}

function isAndroid() {
	return navigator.userAgent.toLowerCase().indexOf('android') != -1;
}
function isIos() {
	var ua = navigator.userAgent.toLowerCase();
	return ua.indexOf('iphone') != -1 || ua.indexOf('ipad') != -1;
}
function isWechat() {
	return navigator.userAgent.toLowerCase().indexOf('micromessenger') != -1;
}