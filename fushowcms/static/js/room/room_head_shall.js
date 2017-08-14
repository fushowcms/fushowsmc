$(function() {
	var uid = getStorage("Id");
	var nicheng1 = getStorage("nicheng");
	if(uid == null) {
		$(".sidebar-login-btns a").eq(0).click(function() {
			$("#loginBox").css("display", "block");
			$('#switch_login').removeClass("switch_btn_focus").addClass('switch_btn');
			$('#switch_qlogin').removeClass("switch_btn").addClass('switch_btn_focus');
			$('#switch_bottom').css({
				left: '0px',
				width: '70px'
			});
			$('#qlogin').css('display', 'none');
			$('#web_qr_login').css('display', 'block');
		});

		$(".sidebar-login-btns a").eq(1).click(function() {
			$("#loginBox").css("display", "block");
			$('#switch_login').removeClass("switch_btn").addClass('switch_btn_focus');
			$('#switch_qlogin').removeClass("switch_btn_focus").addClass('switch_btn');
			$('#switch_bottom').css({
				left: '154px',
				width: '70px'
			});
			$('#qlogin').css('display', 'block');
			$('#web_qr_login').css('display', 'none');
		});
	} else {
		$(".sidebar-login-btns").detach();
		var str = "<li style='color:white;'><a  href='/user/mine_myInform' style='font-size:20px;color:skyblue;font-family:Microsoft Yahei'> " + nicheng1 + "</a></li>";
		$(".sidebar-login-btns1").append(str);
		$(".sidebar-login-banner").detach();
	}
	var anchorId = getPar("anchorId");
	if(uid==null){
	}else{
		$.ajax({
		type: "post",
		url: "/user/isconcern",
		dataType: "json",
		data: {
			UID: uid,
			User: anchorId,
		},
		success: function(msg) {
			if(msg.state == 'not exist') {}
			if(msg.state == 'exist') {
				$(".order").html("已关注").css({
					background: "#00A0FF",
					color: "white"
				});
			}
		},
		error: function() {
			alert("加载失败");
		}
	});
	}
	
})

//#################Head中举报分享等##################//
$("#room-head-lists li").hover(function() {
		var index = $("#room-head-lists li").index($(this));
		$("#room-head-lists li").eq(index).css("color", "#00a0ff");

	}, function() {
		var index = $("#room-head-lists li").index($(this));
		$("#room-head-lists li").eq(index).css("color", "#8c8c8c");

	})
	//###############分享子Div##################//
$("#room-head-lists li").eq(1).mouseover(function() {
	$(" #room-head-shallSon").css("display", "block")
})

$("#room-head-lists li").eq(1).mouseout(function() {
//	var timer = setTimeout(function() {
		$(" #room-head-shallSon").css("display", "none")
//	}, 600);
	$("#room-head-shallSon").mouseover(function() {
		clearTimeout(timer);
		$(" #room-head-shallSon").css("display", "block")
	});
	$("#room-head-shallSon").mouseout(function() {
		$(" #room-head-shallSon").css("display", "none")
	});
})

$("#room-head-lists li").eq(1).mouseover(function() {
		$(" #room-head-shallSon").css("display", "block")
	})
	//#################下载APP##################//
$("#room-head-lists li").eq(2).mouseover(function() {
	$("#room-head-appSon").css("display", "block")
})

$("#room-head-lists li").eq(2).mouseout(function() {

		$("#room-head-appSon").css("display", "none")
	})
	//###############下载APP子Div##################//

//点击举报
$(".report").click(function(){
	// 客户提供的要QQ需去http://wp.qq.com/，进行设置
	var html = "<iframe style='display:none;' src='tencent://message/?uin=3323684573&Site=&menu=yes'><img border='0' src='http://wpa.qq.com/pa?p=2:******:41' alt='哟哟哟' title='哟哟哟'/></iframe>"; 
	var target = document.getElementById('report'); 
	target.innerHTML = html; 

});
//###############点击关注##################//
$(".order").click(function() {
	var anchorId = getPar("anchorId");
	var id = getStorage("Id");
	if(id==null){
		alert("请登录后在关注喜欢的主播");
		return;
	}
	//添加关注
	reqAjax("/user/roomconcernadd",{UID: id,User: anchorId},function(msg) {
		if(msg.ErrorCode!=0){
			alert(msg.ErrorMsg);
			return;
		}else{
			alert(msg.Data);
		}
	});

})

function RoomConcernDel(id, anchorId) {
	$.ajax({
		type: "post",
		url: "/user/cancelroomcon",
		dataType: "json",
		async: false,
		data: {
			UID: id,
			User: anchorId,
		},
		success: function(msg) {
			if(msg.state == "success") {
				alert("您成功取消关注");
				$(".order").html("关注").css({
					background: "white",
					color: "gray"
				});
			} else {
				alert(msg.ErrorMsg);
			}
		},
		error: function() {
			alert("加载失败");
		}
	});
}





$(function() {

	var roomId = getPar("roomId");
	$.ajax({
		type: "post",
		dataType: "json",
		data: {
			RoomId: roomId
		},
		url: "/page/getanchorinfo",
		success: function(msg) {
			$("#roomAlias").text(msg.state[0].RoomAlias);
			$("#nickName").text(msg.state[0].NickName);
			if(msg.state[0].Favicon==""){
			$(".room-head-info-cover").css("background-image","url(/static/images/default_avatar_512_512.jpg)");
			}else{
			$("#favicon").attr("src", msg.state[0].Favicon);
			}
			$("#liveAnnouncement").text(msg.state[0].RoomNotice);
		},
		error: function(sss) {}
	});
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




$("#long").click(function() {
	$("#guess-kinds-2").css("display", "block");
	$("#guess-kinds-1").css("display", "none");
})
$("#qita").click(function() {
	$("#guess-kinds-2").css("display", "block");
	$("#guess-kinds-1").css("display", "none");
})

								
	

$("#guess-kinds-Finalyes-no").click(function() {
		$("#guess-kinds-Finalyes").css("display", "none");
		$("#guess").css("display", "block");
	})
	//#################第一个界面##################//

//##########选中圆点###########//
$(".guess-kinds-2-select").eq(0).click(function() {
	$(".guess-kinds-2-select").eq(0).css("background", "#00A0FF");
	$(".guess-kinds-2-select").eq(1).css("background", "white");
})

$(".guess-kinds-2-select").eq(1).click(function() {
	$(".guess-kinds-2-select").eq(1).css("background", "#00A0FF");
	$(".guess-kinds-2-select").eq(0).css("background", "white");
})

//--####################即时通讯#########################-//

//--#########################tab####################-//
$("#room-chat-tab-head li").eq(0).click(function() {
	/*自己加清*/
	$("#room-chat-tab-head li").eq(0).addClass("tabOn");
	$("#room-chat-tab-head li").eq(0).removeClass("tabOff");
	/*自己列表加清*/
	$("#room-chat-tab-weekRanking").addClass("tab-list-on");
	$("#room-chat-tab-weekRanking").removeClass("tab-list-off");
	/*其他加清*/
	$("#room-chat-tab-head li").eq(1).addClass("tabOff");
	$("#room-chat-tab-head li").eq(1).removeClass("tabOn");
	/*其他列表加清*/
	$("#room-chat-tab-finalRanking").addClass("tab-list-off");
	$("#room-chat-tab-finalRanking").removeClass("tab-list-on");
})

$("#room-chat-tab-head li").eq(1).click(function() {
	/*自己加清*/
	$("#room-chat-tab-head li").eq(1).addClass("tabOn");
	$("#room-chat-tab-head li").eq(1).removeClass("tabOff");
	/*自己列表加清*/
	$("#room-chat-tab-finalRanking").addClass("tab-list-on");
	$("#room-chat-tab-finalRanking").removeClass("tab-list-off");
	/*其他加清*/
	$("#room-chat-tab-head li").eq(0).addClass("tabOff");
	$("#room-chat-tab-head li").eq(0).removeClass("tabOn");
	/*其他列表加清*/
	$("#room-chat-tab-weekRanking").addClass("tab-list-off");
	$("#room-chat-tab-weekRanking").removeClass("tab-list-on");
})