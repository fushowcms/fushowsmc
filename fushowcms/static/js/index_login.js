
$(function(){
//关闭登陆注册弹窗
$("#login-quit").click(function(){
	
    $("#loginBox").css("display","none");   
})

//登陆弹窗
$("body #head-login").click(function() {
	$("#loginBox").css("display", "block");

});
//注册弹窗
$("body #head-setup").click(function() {
	$("#loginBox").css("display", "block");
	$('#switch_login').click();
});


$('#switch_qlogin').click(function(){
		$('#switch_login').removeClass("switch_btn_focus").addClass('switch_btn');
		$('#switch_qlogin').removeClass("switch_btn").addClass('switch_btn_focus');
		$('#switch_bottom').animate({left:'0px',width:'70px'});
		$('#qlogin').css('display','none');
		$('#web_qr_login').css('display','block');
		
		});
	
$('#switch_login').click(function(){
		
		$('#switch_login').removeClass("switch_btn").addClass('switch_btn_focus');
		$('#switch_qlogin').removeClass("switch_btn_focus").addClass('switch_btn');
		$('#switch_bottom').animate({left:'154px',width:'70px'});
		$('#qlogin').css('display','block');
		$('#web_qr_login').css('display','none');
});

	//加载推荐机构列表
	$.ajax({
				type: "post",
				url: "/page/getallaffiliation",
				dataType: "json",
				success: function(msg) {
					if(msg.state){
						$.each(msg.data, function(){
						$("#aff_id").append('<option value="'+this.AffId+'">'+this.InstitutionName+'</option>');
						});
					}
				}
			});

	var ID = getStorage("Id");
	var niCheng = getStorage("nicheng");
	var userName = getStorage("username");
	
	if(ID!=null&&userName!=null&&niCheng!=null){
		var str = "";
		str += '<li><a href="/user/mine_myInform"style =" color:orange;">'+userName+'你好</a></li>';
	    str+='<li style="color:skyblue" class="out">退出登录</li>';
	    $("#header-Setup-Login").html(str);
	  
	}else{
		console.log('尚未登陆');
		
	}
	//点击获取验证码
	$("#getting").click(function(){
		    var pho =$("#Phone").val();
			if(!pho){
				$("#userCue").html("<font color='red'><b>×手机号不能为空</b></font>");
				return false;
			}
			reg_phone=/^1[34578]\d{9}$/;
			if(!reg_phone.test(pho)){
				$("#userCue").html("<font color='red'><b>×手机格式不正确</b></font>");
				return false;
			}
            var btn = $(this);
            var count = 60;
			$('#phone_check').css('display','block');
            var resend = setInterval(function(){
                count--;
                if (count > 0){
                    btn.val(count+"秒后可重新获取");
                   
                }else {
                    clearInterval(resend);
                    btn.val("获取验证码").removeAttr('disabled');
                }
            }, 1000);
            btn.attr('disabled',true).css({'cursor':'not-allowed'});
			$.ajax('/page/sms', {
					data: {
						account: 'cf_hxwy2014',
						password: '123456',
						mobile: pho,
						content: '您的验证码是：' + token + '。请不要把验证码泄露给其他人。如非本人操作，可不用理会！'
					},
					dataType: 'json', //服务器返回json格式数据
					type: 'post', //HTTP请求类型
					timeout: 10000, //超时时间设置为10秒；
					success: function(data, textStatus, xhr) {
						alert("验证码已发送！");
					},error:function(){
						alert("短信服务器错误");
					}
        });
		
	   


});
    

function cli(){
			//登陆
			var id = document.getElementById('u').value;
			var name = document.getElementById('u').value;
			var pwd = document.getElementById('p').value;
			
			if(!name){
				$('#login_prompt').html('用户名不能为空');
				return;
			}
			if(!pwd){
				$('#login_prompt').html('密码不能为空');
				return;
			}

			$.ajax({
				type: "post",
				url: "/page/login",
				dataType: "json",
				data: {
					//uid:id,
					username: name,
					password: pwd
				},
				success: function(msg, data, xhr) {
					if (name == msg.UserName && pwd == msg.PassWord) {
						
						setStorage("Id", msg.UID);
						setStorage("username", msg.UserName);
						setStorage("nicheng", msg.NickName);
						
						alert("登陆成功!");
						$("#loginBox").css("display", "none");
						$("#head-login").css({width:"120",color:"orange"});
						$("#head-setup").css("display","none");
						//var helloUser = $("<a>");
						var str = "";
						str += '<li><a href="/user/mine_myInform" style="color:orange;">'+msg.UserName+'你好</a></li>';
						str+='<li style="width:100px;color:skyblue" class="out">退出登录</li>';
						$("#header-Setup-Login").html(str);
						window.location.href="/user/mine_myInform";
					}else{
						
					}
				},
				error: function(msg) {
					alert("网络环境异常，请稍好再试");
				}
			});
		}
		
function regs() {//注册
			if ($('#user').val() == "") {
				$('#user').focus().css({
					border: "1px solid red",
					boxShadow: "0 0 2px red"
				});
				$('#userCue').html("<font color='red'><b>×用户名不能为空</b></font>");
				return;
			}
	
			if ($('#user').val().length < 4 || $('#user').val().length > 16) {
	
				$('#user').focus().css({
					border: "1px solid red",
					boxShadow: "0 0 2px red"
				});
				$('#userCue').html("<font color='red'><b>×用户名位4-16字符</b></font>");
				return;
			}
	
			if ($('#passwd').val().length < pwdmin) {
				$('#passwd').focus();
				$('#userCue').html("<font color='red'><b>×密码不能小于" + pwdmin + "位</b></font>");
				return;
			}
			if ($('#passwd2').val() != $('#passwd').val()) {
				$('#passwd2').focus();
				$('#userCue').html("<font color='red'><b>×两次密码不一致！</b></font>");
				return;
			}
			
			if ($('#NickName').val().length == 0) {
				$('#NickName').focus();
				$('#userCue').html("<font color='red'><b>×昵称不能为空！</b></font>");
				return;
			}
			
			if ($('#Phone').val().length < 1) {
				$('#Phone').focus();
				$('#userCue').html("<font color='red'><b>×手机号不能为空！</b></font>");
				return;
			}
			if ($('#Phone').val().length != 11) {
				$('#Phone').focus();
				$('#userCue').html("<font color='red'><b>×手机号格式不正确！</b></font>");
				return;
			}

			var username = document.getElementById('user').value;
			var pwd = document.getElementById('passwd').value;
			var nickname = document.getElementById('NickName').value;
			var phone = document.getElementById('Phone').value;
			$.ajax({  
				type: "post",  
				url: "/page/regist",  
				dataType: "json",  
				data: {
					UserName:username,
					PassWord:pwd,
					NickName:nickname,
					Phone:phone
				},  
				success: function(msg){
					 if(msg.Id > 0 && msg.UserName != ""){
						
						setStorage("Id", msg.UID);
						setStorage("username", msg.UserName);
						setStorage("nicheng", msg.NickName);
						
						alert("注册成功");
						$('.o-login').css('display','block');
						$('#loginBox').css('display','none');
						var str = "";
						str += '<li><a href="/user/mine_myInform" style="color:orange;">'+msg.UserName+'你好</a></li>';
						str+='<li style="width:100px;color:skyblue" class="out">退出登录</li>';
						$("#header-Setup-Login").html(str);
					} if(msg.state=='exist'){
				            alert("用户已存在");
				    }if(msg.state=='failuk'){
				            alert("增加用户失败");
				    } 
					
				},
				error: function(msg){
						alert("网络环境异常，请稍好再试");
				}	  
			});
		}



$(document).ready(function() { 
		
		// var code = getPar("code");
		// if(code!=""){
		// 	$.ajax({  
		// 		type:"post",
		// 		url: "http://open.51094.com/user/auth.html",
		// 		dataType: "json",
		// 		data: {
		// 			type:"get_user_info",
		// 			code:code,
		// 			appid:"157ca89018b36a",
		// 			token:"96e26c0180780fec8d7ab49b18edd1d1"
		// 		},
		// 		asynch:true,
		// 		success: function(msg){
		// 			alert(msg.name);
		// 		}  
		// 	});
		// }
	});
	
	
	function getPar(par){
	    //获取当前URL
	    var local_url = document.location.href; 
	    //获取要取得的get参数位置
	    var get = local_url.indexOf(par +"=");
	    if(get == -1){
	        return "";   
	    }   
	    //截取字符串
	    var get_par = local_url.slice(par.length + get + 1);    
	    //判断截取后的字符串是否还有其他get参数
	    var nextPar = get_par.indexOf("&");
	    if(nextPar != -1){
	        get_par = get_par.slice(0, nextPar);
	    }
	    return get_par;
	}

});