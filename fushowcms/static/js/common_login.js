$(function(){
	var code =getPar("code")
	if(code!=""){
		$.ajax({
			type: "post",
			url: "/page/disanfangdenglu",
			dataType: "json",
			data: {
				code:code
			},success:function(msg){
				if(msg.state){
					setStorage("Id", msg.data.UID);
					setStorage("username", msg.data.UserName);
					setStorage("nicheng", msg.data.NickName);
					
					$("#loginBox").css("display", "none");
					$("#head-login").css({width:"120",color:"orange"});
					$("#head-setup").css("display","none");
					//var helloUser = $("<a>");
					var str = "";
					str += '<li><a href="/user/mine_myInform" style="color:orange;">'+msg.data.UserName+'</a></li>';
					str+='<li style="width:100px;color:skyblue" class="out">退出登录</li>';
					$("#header-Setup-Login").html(str);
					window.location.href="/user/mine_myInform";
				}
			}
		});
	}

	
	var ID = getStorage("Id");
	var niCheng = getStorage("nicheng");
	var userName = getStorage("username");
	
	if(ID!=null&&userName!=null&&niCheng!=null){
		var str = "";
		str += '<li><a href="/user/mine_myInform"style =" color:orange;">'+userName+'你好</a></li>';
	    str+='<li style="color:skyblue" class="out">退出登录</li>';
	    $("#header-Setup-Login").html(str);
	}else{
		//console.log('尚未登陆');
	}
	var aff ='';
		//推荐机构列表
		$.ajax({  
			type: "post",  
			url: "/page/getafflist",  
			dataType: "json",  
			data: null, 
			error:function(){
				//Dialog("机构获取错误");
			},success:function(msg){
				if(msg.state){
					$.each(msg.data,function(k,v){
						aff+='<option value="'+v.AffId+'" >'+v.InstitutionName+'</option>'
					})
					$("#aff_id").append(aff);
				}
			}
		});
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
			$('#get').css('display','block');
            var resend = setInterval(function(){
                count--;
                if (count > 0){
                    btn.val(count+"秒后可重新获取");
                   
                }else {
                    clearInterval(resend);
                    btn.val("获取验证码").removeAttr('disabled');
                    $("#Phone").attr("disabled",false);
                }
            }, 1000);
            btn.attr('disabled',true).css({'cursor':'not-allowed'});
            
			reqAjax("/page/byphonereg",{mobile: pho},function(msg){
				if(msg.ErrorCode!=0) {
					Dialog(msg.ErrorMsg,true,"确定",null,function() {
						$('.dialog').remove();
					},null);
				}else {
					Dialog("发送成功",true,"确定",null,function() {
						$('.dialog').remove();
					},null);
            		$("#Phone").attr("disabled",true);
				}
			},true);          
        });
});