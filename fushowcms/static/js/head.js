$(function(){
	broadcast();
});
$(function() {
	$("#listZhibo").hover(
		function() {
			$("#listZhibo").css("background", " #e84c3d");
			$("#listZhibo a").css("color", "white");
		},
		function() {
			$("#listZhibo").css("background", " white");
			$("#listZhibo a").css("color", "#e84c3d");
		}
	);

	$("#listZhibo").on('click', function() {
		var pages = setStorage("zbpage", 1);
	})

	//*************登陆点击****************//
	$(".header-login-btn").click(function() {
		$("#loginBox").css("display", "block");
		$('#switch_login').removeClass("switch_btn_focus").addClass('switch_btn');
		$('#switch_qlogin').removeClass("switch_btn").addClass('switch_btn_focus');
		$('#switch_bottom').css({
			left: '0px',
			width: '70px'
		});
		$('#qlogin').css('display', 'none');
		$('#web_qr_login').css('display', 'block');
		document.documentElement.style.overflow = 'hidden';

	});
	//*************注册点击****************//
	$(".header-register-btn").click(function() {
		$("#loginBox").css("display", "block");
		$('#switch_login').removeClass("switch_btn").addClass('switch_btn_focus');
		$('#switch_qlogin').removeClass("switch_btn_focus").addClass('switch_btn');
		$('#switch_bottom').css({
			left: '154px',
			width: '70px'
		});
		$('#qlogin').css('display', 'block');
		$('#web_qr_login').css('display', 'none');
		document.documentElement.style.overflow = 'hidden';
	});

	$("body").on('click', '#login-quit', function() {
		$("#loginBox").css("display", "none");
		document.documentElement.style.overflow = 'scroll';
	})

	$('#switch_qlogin').click(function() {
		$('#switch_login').removeClass("switch_btn_focus").addClass('switch_btn');
		$('#switch_qlogin').removeClass("switch_btn").addClass('switch_btn_focus');
		$('#switch_bottom').animate({
			left: '0px',
			width: '70px'
		});
		$('#qlogin').css('display', 'none');
		$('#web_qr_login').css('display', 'block');

	});
	$('#switch_login').click(function() {
		$('#switch_login').removeClass("switch_btn").addClass('switch_btn_focus');
		$('#switch_qlogin').removeClass("switch_btn_focus").addClass('switch_btn');
		$('#switch_bottom').animate({
			left: '154px',
			width: '70px'
		});
		$('#qlogin').css('display', 'block');
		$('#web_qr_login').css('display', 'none');
	});

	//	if(getParam("a") == '0') {
	//		$('#switch_login').trigger('click');
	//	}

	//搜索
	$("#header-serch-Btn").on("click", function() {
		var headerSearch = $("#header-Search").val();
		setStorage("cxzbpage", 1);
		window.location.href = "/page/search?kw=" + headerSearch;
	});

	function getPar(par) {
		//获取当前URL
		var local_url = document.location.href;
		//获取要取得的get参数位置
		var get = local_url.indexOf(par + "=");
		if(get == -1) {
			return "";
		}
		//截取字符串
		var get_par = local_url.slice(par.length + get + 1);
		//判断截取后的字符串是否还有其他get参数
		var nextPar = get_par.indexOf("&");
		if(nextPar != -1) {
			get_par = get_par.slice(0, nextPar);
		}
		return get_par;
	}
	var code = getPar("code")
	if(code != "") {
		$.ajax({
			type: "post",
			url: "/page/disanfangdenglu",
			dataType: "json",
			data: {
				code: code
			},
			success: function(msg) {
				if(msg.state) {
					setStorage("Id", msg.data.UID)
					setStorage("username", msg.data.UserName)
					setStorage("nicheng", msg.data.NickName)
					$("#loginBox").css("display", "none");
					$("#head-login").css({
						width: "120",
						color: "orange"
					});
					$("#head-setup").css("display", "none");
					var str = "";
					str += '<li><a href="/user/mine_myInform" style="color:orange;">' + msg.data.UserName + '</a></li>';
					str += '<li style="width:100px;color:skyblue" class="out">退出登录</li>';
					$("#header-Setup-Login").html(str);
					window.location.href = "/user/mine_myInform";
				}
			}
		});
	}
});