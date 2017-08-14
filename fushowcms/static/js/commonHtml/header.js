$("li").mouseover(function() {
	$(this).css("cursor", "pointer");
})

$(function() {
	//$.getScript('/static/js/jquery-1.9.0.min.js');
	//$.getScript('http://open.51094.com/user/myscript/157ca89018b36a.html');
	
	var ID = getStorage("Id");
	var niCheng = getStorage("nicheng");
	var userName = getStorage("username");

	if(ID != null && userName != null && niCheng != null || ID != undefined || userName != undefined) {
		var str = "";
		str += '<li><a href="/user/mine_myInform"style ="color:orange;">' + niCheng + '</a></li>';
		str += '<li style="color:skyblue" class="out">退出登录</li>';
		$("#header-Setup-Login").html(str);
		//alert("hah");
		$("li").mouseover(function() {
			$(this).css("cursor", "pointer");
		});
	} else {
		//alert("这是else");
	}


	$('header').on('click', '.out', function() {
		if(confirm('确定要退出登陆吗?')) {
			//alert("成功退出!");
			var id = getStorage("Id")
			$.ajax({
				type: "post",
				url: "/user/unlogin",
				async: true,
				data: {
					UID: id
				},
				success: function(msg) {
					//console.log(msg);
					if(msg.flag == true) {
						removeStorage("Id");
						removeStorage("nicheng");
						removeStorage("username");
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
	});
});