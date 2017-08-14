$("#mine-Tab-ul li").click(function() {
	var index = $("#mine-Tab-ul li").index($(this));
	//alert(index);
	//	//$("#mine-Tab-ul li").css("background","white");
	//  //$("#mine-Tab-ul li a").css("color","black");
	//	$("#mine-Tab-ul li").eq(index).css("background","#e84c3d");
	//	$("#mine-Tab-ul li a").eq(index).css("color","white");

	//alert(i);
});

$("#mine-Tab-ul li").hover(function() {
	var i = $("#mine-Tab-ul li").index($(this));
	$("#mine-Tab-ul li a").eq(i).css("color", "#e84c3d");
	if(i == 0) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "-59px 2px");
	} else if(i == 1) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "-59px -365px");
	} else if(i == 2) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "-59px -188px");
	} else if(i == 3) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "-59px -318px");
	} else if(i == 4) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "-59px -140px");
	} else if(i == 5) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "-59px -94px");
	} else if(i == 6) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "-59px -450px");
	} else if(i == 7) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "-59px -414px");
	}
	//  switch(i){
	//  	case:0
	//  	//$("#mine-Tab-ul li i").eq(0).css("background-position","-59px 2px");	
	//     alert("111");
	//     break;
	//      case:1
	//       $("#mine-Tab-ul li i").eq(1).css("background-position","-59px 2px");	
	//      break;
	//  }
}, function() {
	var i = $("#mine-Tab-ul li").index($(this));
	$("#mine-Tab-ul li a").eq(i).css("color", "black");
	//$("#mine-Tab-ul li i").eq(i).css("left","");
	if(i == 0) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "2px 2px");
	} else if(i == 1) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "2px -365px");
	} else if(i == 2) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "0px -188px");
	} else if(i == 3) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "0px -318px");
	} else if(i == 4) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "0px -140px");
	} else if(i == 5) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "0px -94px");
	} else if(i == 6) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "0px -450px");
	} else if(i == 7) {
		$("#mine-Tab-ul li i").eq(i).css("background-position", "0px -414px");
	}
})

$(function(){
	//******************关于主播********************//
$("#beAnchor").toggle(function() {
	var Uid = getStorage("Id");
	
	if(Uid == null) {
		alert("您还没有登录")
	} else {
		reqAjax("/page/getuser",{UID:Uid},function(msg){
			if(msg.ErrorCode!=0) {
				Dialog(msg.ErrorMsg,true,"确定",null,function() {
					$('.dialog').remove();
				},null);
			}else {
				if(msg.Data.Type == 0) {
					Dialog("仅主播有该权限",true,"确定",null,function() {
						$('.dialog').remove();
					},null);
				} else {
					$(".anchor-Inform").show("slow");
				}
			}
		},true);
	}
	//alert("123");
	//	$(".anchor-Inform").css("display","block");
	//	$(".anchor-Inform").css("transition","5s");
}, function() {
	//	$(".anchor-Inform").css("display","none");
	//	$(".anchor-Inform").css("transition","2s");
	$(".anchor-Inform").hide("slow");
})
})

$("#tuichu").click(function() {
	//alert(123456);
	if(confirm('确定要退出登录吗?')) {
		//alert("成功退出!");
		var id = getStorage("Id");
		
		$.ajax({
			type: "post",
			url: "/user/unlogin",
			async: true,
			data: {
				UID: id
			},
			success: function(msg) {
				if(msg.flag == true) {
					
					removeStorage("Id");
					removeStorage("nicheng");
					removeStorage("username");
					//window.location.href = "./";
					alert("成功退出");
					window.location.href = "/";
				} else {
					alert("退出失败，错误代码205");
				}
			},
			error: function(msg) {
				alert("退出失败");
			}
		});

	} else {
		//alert('您未退出');
	}
})