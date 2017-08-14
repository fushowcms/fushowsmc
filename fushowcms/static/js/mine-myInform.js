$(function() {
	// alert("jibenziliao!");

	var uid = getStorage("Id")
	
	$.ajax({
		type: "post",
		url: "/page/getuser",
		dataType: "json",
		data: {
			UID: uid
		},
		async: true,
		success: function(msg, xhr) {
			var nicheng = msg.NickName;
			var phone = msg.Phone;
			var yue = msg.Balance;
			//alert(nicheng);
			$("#userName").text(nicheng);
			$("#phoneNum").text(phone);
			$(".yue").text(yue);
			// alert("个人信息请求成功!");
		},
		error: function() {
			alert("个人信息请求失败!");
		}
	});
})

$("#basic").click(function() {
	$("#mine-myInform-basicInf").addClass("mine-myInform-On");
	$("#mine-myInform-basicInf").removeClass("mine-myInform-Off");
	//$("#mine-myInform-ul li").eq(0)

	$("#chengePic").addClass("mine-myInform-Off");
	$("#chengePassw").addClass("mine-myInform-Off");
	$("#chengeName").addClass("mine-myInform-Off");
	$("#chengePhone").addClass("mine-myInform-Off");
	$("#realName").addClass("mine-myInform-Off");
	$("#chengeEmail").addClass("mine-myInform-Off");

	$("#chengePic").removeClass("mine-myInform-On");
	$("#chengePassw").removeClass("mine-myInform-On");
	$("#chengeName").removeClass("mine-myInform-On");
	$("#chengePhone").removeClass("mine-myInform-On");
	$("#chengeEmail").removeClass("mine-myInform-On");

	$("#mine-myInform-TabBox").css("display", "block");

});

$("#mine-myInform-ul li").eq(1).click(function() {
	$("#chengePic").removeClass("mine-myInform-Off");
	$("#chengePic").addClass("mine-myInform-On");

	//$("#mine-myInform-ul li").eq(0)

	$("#mine-myInform-basicInf").addClass("mine-myInform-Off");
	$("#chengePassw").addClass("mine-myInform-Off");
	$("#chengeName").addClass("mine-myInform-Off");
	$("#chengePhone").addClass("mine-myInform-Off");
	$("#realName").addClass("mine-myInform-Off");
	$("#chengeEmail").addClass("mine-myInform-Off");

	$("#mine-myInform-basicInf").removeClass("mine-myInform-On");
	$("#chengePassw").removeClass("mine-myInform-On");
	$("#chengeName").removeClass("mine-myInform-On");
	$("#chengePhone").removeClass("mine-myInform-On");
	$("#chengeEmail").removeClass("mine-myInform-On");

	$("#mine-myInform-TabBox").css("display", "none");
})

$("#mine-myInform-ul li").eq(2).click(function() {
	$("#chengePassw").removeClass("mine-myInform-Off");
	$("#chengePassw").addClass("mine-myInform-On");

	//$("#mine-myInform-ul li").eq(0)

	$("#mine-myInform-basicInf").addClass("mine-myInform-Off");
	$("#chengePhone").addClass("mine-myInform-Off");
	$("#chengePic").addClass("mine-myInform-Off");
	$("#chengeName").addClass("mine-myInform-Off");
	$("#realName").addClass("mine-myInform-Off");
	$("#chengeEmail").addClass("mine-myInform-Off");

	$("#mine-myInform-basicInf").removeClass("mine-myInform-On");
	$("#chengePhone").removeClass("mine-myInform-On");
	$("#chengePic").removeClass("mine-myInform-On");
	$("#chengeName").removeClass("mine-myInform-On");
	$("#chengeEmail").removeClass("mine-myInform-On");

	$("#mine-myInform-TabBox").css("display", "none");
})

$("#mine-myInform-ul li").eq(3).click(function() {
	$("#chengeName").removeClass("mine-myInform-Off");
	$("#chengeName").addClass("mine-myInform-On");

	//$("#mine-myInform-ul li").eq(0)

	$("#mine-myInform-basicInf").addClass("mine-myInform-Off");
	$("#chengePassw").addClass("mine-myInform-Off");
	$("#chengePic").addClass("mine-myInform-Off");
	$("#chengePhone").addClass("mine-myInform-Off");
	$("#realName").addClass("mine-myInform-Off");
	$("#chengeEmail").addClass("mine-myInform-Off");

	$("#mine-myInform-basicInf").removeClass("mine-myInform-On");
	$("#chengePassw").removeClass("mine-myInform-On");
	$("#chengePic").removeClass("mine-myInform-On");
	$("#chengePhone").removeClass("mine-myInform-On");
	$("#chengeEmail").removeClass("mine-myInform-On");

	$("#mine-myInform-TabBox").css("display", "none");
})

$("#mine-myInform-ul li").eq(4).click(function() {

	$("#realName").removeClass("mine-myInform-Off");

	//$("#mine-myInform-ul li").eq(0)

	$("#mine-myInform-basicInf").addClass("mine-myInform-Off");
	$("#chengePassw").addClass("mine-myInform-Off");
	$("#chengePic").addClass("mine-myInform-Off");
	$("#chengeName").addClass("mine-myInform-Off");
	$("#chengePhone").addClass("mine-myInform-Off");
	$("#chengeEmail").addClass("mine-myInform-Off");

	$("#mine-myInform-basicInf").removeClass("mine-myInform-On");
	$("#chengePassw").removeClass("mine-myInform-On");
	$("#chengePic").removeClass("mine-myInform-On");
	$("#chengeName").removeClass("mine-myInform-On");
	$("#chengePhone").removeClass("mine-myInform-On");
	$("#chengeEmail").removeClass("mine-myInform-On");

	$("#mine-myInform-TabBox").css("display", "none");
})

$("#mine-myInform-ul li").eq(5).click(function() {
	$("#chengePhone").removeClass("mine-myInform-Off");
	$("#chengePhone").addClass("mine-myInform-On");

	//$("#mine-myInform-ul li").eq(0)

	$("#mine-myInform-basicInf").addClass("mine-myInform-Off");
	$("#chengePassw").addClass("mine-myInform-Off");
	$("#chengePic").addClass("mine-myInform-Off");
	$("#chengeName").addClass("mine-myInform-Off");
	$("#realName").addClass("mine-myInform-Off");
	$("#chengeEmail").addClass("mine-myInform-Off");

	$("#mine-myInform-basicInf").removeClass("mine-myInform-On");
	$("#chengePassw").removeClass("mine-myInform-On");
	$("#chengePic").removeClass("mine-myInform-On");
	$("#chengeName").removeClass("mine-myInform-On");
	$("#chengeEmail").removeClass("mine-myInform-On");

	$("#mine-myInform-TabBox").css("display", "none");
})

$("#mine-myInform-ul li").eq(6).click(function() {
	$("#chengeEmail").removeClass("mine-myInform-Off");
	$("#chengeEmail").addClass("mine-myInform-On");

	//$("#mine-myInform-ul li").eq(0)

	$("#mine-myInform-basicInf").addClass("mine-myInform-Off");
	$("#chengePassw").addClass("mine-myInform-Off");
	$("#chengePic").addClass("mine-myInform-Off");
	$("#chengeName").addClass("mine-myInform-Off");
	$("#realName").addClass("mine-myInform-Off");
	$("#chengePhone").addClass("mine-myInform-Off");

	$("#mine-myInform-basicInf").removeClass("mine-myInform-On");
	$("#chengePassw").removeClass("mine-myInform-On");
	$("#chengePic").removeClass("mine-myInform-On");
	$("#chengeName").removeClass("mine-myInform-On");
	$("#chengePhone").removeClass("mine-myInform-On");
	$("#chengeEmail").removeClass("mine-myInform-On");

	$("#mine-myInform-TabBox").css("display", "none");
})

/******************背景色********************/
$("#mine-myInform-ul li").click(function() {
	var index = $("#mine-myInform-ul li").index($(this));
	$("#mine-myInform-ul li").css({
		color: "gray",
		background: "white"
	});
	$("#mine-myInform-ul li").eq(index).css({
		color: "white",
		background: "deepskyblue"
	});
	//alert(index);	 
})

/****************基本资料--->修改昵称********************/
$("#amend").click(function() {
	$("#mine-myInform-TabBox").css("display", "none");
	$("#chengeName").css("display", "block");

	$("#basic").click(function() {
		$("#mine-myInform-TabBox").css("display", "block");
		$("#chengeName").css("display", "none");
	})
})

$("span").mouseover(function() {
	$(this).css("cursor", "pointer");
})
$("p").mouseover(function() {
		$(this).css("cursor", "pointer");
	})
	/****************基本资料--->修手机********************/
$("#span-changePhone").click(function() {
	$("#mine-myInform-TabBox").addClass("mine-myInform-Off");
	$("#chengePhone").addClass("mine-myInform-On");

})

$("#basic").click(function() {
	$("#mine-myInform-TabBox").addClass("mine-myInform-On");
	$("#chengePhone").addClass("mine-myInform-Off");
})

/****************基本资料--->修改手机********************/
$("#chengePhone-Btn").click(function() {
	var num = $("#chengePassw-Btn input").eq(0).val();
	var num2 = $("#phoneNum").text();
	var telNum = $("#telNum").val();
	var uid = getStorage("Id")
	$.ajax({
		type: "post",
		url: "/user/userup",
		dataType: "json",
		data: {
			"UID": uid,
			"Phone": telNum
		},
		success: function(msg) {
			alert(msg);
		}
	});

})
$("#chengeName-Btn").click(function() {
	var num = $("#nicheng").val();
	var uid = getStorage("Id")
	$.ajax({
		type: "post",
		url: "/user/userup",
		dataType: "json",
		data: {
			"UID": uid,
			"NickName": num
		},
		success: function(msg) {
			alert("修改成功")
		}
	});
})
$("#chengePassw-Btn").click(function() {
	var num = $("#newpwd").val();
	var num2 = $("#newpwd2").val();
	var uid = getStorage("Id");
	$.ajax({
		type: "post",
		url: "/user/passup",
		dataType: "json",
		data: {
			"username": uid,
			"password": num,
			"newpassword": num2
		},
		success: function(msg) {
			alert(msg.state)
		}
	});
})

//成为主播	
$("#submit").click(function() {
	var real = $("#realeName").val();
	var uid = getStorage("Id")
	var idcard = $("#idcard").val();
	var phone = $("#phone").val();
	//txl -del
	reqAjax("/user/applyadd",{UID:uid,RealName:real,IdNumber:idcard,Phone:phone,upload:"E:\WuWork\src\fushowcms\static\images\default_avatar_64_64.jpg"},function(data) {
		alert(msg.data.data);
	});
})

//修改邮箱	
$("#chengePhone-Btn").click(function() {

})