var client,player,timer;//聊天客户端、播放器、定时器
//相关参数配置
var ProId,PerId,RoomId,UserId;
var IsFullScreen = true;
var chatconf = {
	host: "192.168.1.200",
//	host: "139.129.227.66",
	port: "61623",
	clientId: "example" + (Math.floor(Math.random() * 100000)),
	user: "admin",
	password: "www.fushow.cn"
};

var vpage = {
	roomId: 0,
	userId: getStorage('Id'),
	anchorId: 0,
	chatstate: false,
	authority: false,
	allowsendmsg: true,
	danmubtn: true,
	dialogmsg: '发送时间间隔过短，请稍候重试'
};


var userinfo = {};
var guessparams = {};


var text; 
//var text = "http://www.shiliulianmeng.com";
var mousemove = '';
$(function(){
	
        
		var dd = document.documentElement;
		if(!(dd.requestFullScreen || dd.mozRequestFullScreen || dd.webkitRequestFullScreen || dd.msRequestFullscreen)) {
			$('.fullscreenbtn').css('display','none');
			$('#room-video').hover(function(){
				$('#room-video').append('<p class="room-video-tips">可双击屏幕打开全屏模式</p>');					
			},function(){
				$('#room-video').find('.room-video-tips').remove();
			});
			
		}
	
	pageinit();
	liaotianlianjie();
	//if(!vpage.roomId){
	//	window.location.href="/page/index";
	//}
	var url = window.location.href;
	RoomId = url.substr(url.lastIndexOf('/')+1,url.length);
		if (anchor != 0){
			UserId = anchor;
			vpage.anchorId =anchor;
			vpage.roomId =RoomId;
			text ="http://tv.fushow.cn/wechat/exp/bofang0.html?roomid="+RoomId+"&anchorid="+UserId;
		}else{
			window.location.href="/";
		}



	//聊天服务相关
	// if(!window.WebSocket){
	// 	$('#message11').html('<p style="color:red">因当前浏览器版本过低，聊天功能可能无法正常使用，请尝试升级或更换浏览器</p>');
	// }else{
	// 	$.getScript('/static/js/fushowim.js',function(){
	// 		client = new Messaging.Client(chatconf.host, Number(chatconf.port), chatconf.clientId);
	// 		//建立连接
	// 		client.connect({
	// 			userName: chatconf.user,
	// 			password: chatconf.password,
	// 			onSuccess: onConnect,
	// 			onFailure: onFailure
	// 		});
	// 		//监听事件的调用方法
	// 		client.onConnect = onConnect;
	// 		client.onMessageArrived = onMessageArrived;
	// 		client.onConnectionLost = onConnectionLost;
	// 	});	
	// }

//	//获取直播间信息
//	reqAjax('/page/getroom',{AnchorId: vpage.anchorId},function(data){
//		$("#roomAlias").text(data.Data.state[0].RoomAlias);
//		$("#nickName").text('主播：'+data.Data.state[0].NickName);
//		var avatarurl = data.Data.state[0].Favicon?data.Data.state[0].Favicon:'/static/images/default_avatar.jpg';
//		$('.room-head-info-cover').html('<img src="'+avatarurl+'" onerror="this.src=\'/static/images/default_avatar.jpg\'"/>');
//		$("#liveAnnouncement").text(data.Data.state[0].RoomNotice);
//		$('.ordernum').text(data.Data.attention);
//		$('.viewnum').text(data.Data.state[0].LiveNumber);
//		$('.hotnum').text(data.Data.state[0].GiftNum);
//	});
	
	if(vpage.userId){
		//是否关注该主播并获取用户信息
			if(IsConcern){
				$('.order').html('取消关注').addClass('orderstate');
				$('.ordernum').addClass('ordernumstate');
				$('.order-area').addClass('order-areastate');
			}

		
		//获取当前用户信息
			if(getUserInfo) {
				userinfo = Ui;
				if(Ui.Id == vpage.anchorId)
				vpage.authority = true;//给予权限
			}else {
				Dialog("数据获取失败！",true,"确定",null,function() {
					$('.dialog').remove();
				},null);
			}

		
		//当前用户是否为房管findbyroomiduseridall
			if(Fball != null){
				if(userinfo.Id == Fball[0].UserId){
					userinfo.roomclass = 2 ;
					vpage.authority = true;//给予权限
				}
			}
		
		//当前用户是否被禁言bannedtopostall
			if(Fnall != null){
				var len = Fnall.length - 1;
				var timeer = Fnall[len].ModifyTime;
				var timestamp = Date.parse(new Date(timeer));
				var nowtime = $.now();
 				var bannedtime = (timestamp + 900000) - nowtime ;
				vpage.allowsendmsg = false;
				vpage.dialogmsg = '你已被禁言';
				setTimeout(function(){
					vpage.allowsendmsg = true;
					vpage.dialogmsg = '发送时间间隔过短，请稍候重试';
				},bannedtime)
			}
		
		//聊天域下点击用户名相关操作
		$('#message11').on('click','.chat-uname',function(){
			
			$('.chat-usercard').addClass('chat-usercard-show');
			var parentsel = $(this).parent();
			var cardmod = '<div class="chat-usercard-avatar"><img src="'+parentsel.data('avatar')+'" onerror="this.src=\'/static/images/default_avatar.jpg\'"/></div>';
			cardmod += '<div class="chat-usercard-close">关闭</div>';
			cardmod += '<div class="chat-usercard-nickname">'+parentsel.data('name')+'</div>';
			cardmod += '<div class="chat-usercard-level" style="display:none">'+parentsel.data('level')+'</div>';
			cardmod += '<div class="chat-usercard-button" data-uname="' + parentsel.data('name') + '" data-gid="' + parentsel.data('cid') + '" data-pid="' + parentsel.data('pid') + '" data-class="' + parentsel.attr('data-class') + '">';
			
			if(userinfo.Id == vpage.anchorId && parentsel.data('pid')!=vpage.anchorId){
				if(parentsel.attr('data-class') > 0){
					cardmod += '<i class="rmpoint">撤销房管</i>';
				}else{
					cardmod += '<i class="appoint">任命房管</i>';
				}
			}

			if(parentsel.data('pid') != vpage.anchorId && vpage.authority){
				cardmod += '<i class="mute">禁止发言</i>';
			}
			
			$('.chat-usercard').html(cardmod).on('click','.chat-usercard-close',function(){
				$('.chat-usercard').empty().removeClass('chat-usercard-show');
			});
			
			//可进行房间管理
			if(vpage.authority){
				//如果是主播可设置房管
				$('.chat-usercard').find('.appoint').click(function(){
					authAlert();
				});
				
				//如果是主播可撤销房管
				$('.chat-usercard').find('.rmpoint').click(function(){
					authAlert();
				});
				
				//如果是主播或房管可设置禁言
				$('.chat-usercard').find('.mute').click(function(){
					authAlert();
				});
			}
		});
		//失去焦点隐藏
		$("body").on("click",function(e){
			if(e.target!=$('.chat-usercard').get(0) && !$(e.target).hasClass('chat-uname')){
				$('.chat-usercard').empty().removeClass('chat-usercard-show');
			}   
		})

		
	}

	//点击举报、分享、下载APP add by liuhan -----start--------
	$(".report").click(function(){
		var html = "<iframe style='display:none;' src='tencent://message/?uin=3323684573&Site=&menu=yes'><img border='0' src='http://wpa.qq.com/pa?p=2:******:41' alt='哟哟哟' title='哟哟哟'/></iframe>"; 
		$('#report').append(html);
	});
	
	$("#room-head-lists li").hover(function() {
		$(this).addClass('hover');
		
	},function(){
		$(this).removeClass('hover');
	});
	
	$('#asd').qrcode({
//		render:'table',
		text: text, 
		width: 150, 
		height: 150,
		background: "#fff",
		foreground: "black",
		position: "relative",
		src: "/static/images/fxlogo.png"   
	});
	//add by liuhan -------end--------

	//点击关注按钮
	$(".order").click(function(){
		if(!getStorage('Id')){
			$('#loginBox,.logins-alert-bg').css('display','block');
			return;
		}
		
		if($('.order').is('.orderstate')){
			RoomConcernDel(vpage.userId, vpage.anchorId);
		}else{
			reqAjax('/user/roomconcernadd',{User:vpage.anchorId,UID:vpage.userId},function(data){
				console.log(data)
				if(data.ErrorCode != 0){
					if(data.ErrorCode == "2035"){
						return;
					}
				}else if(data.ErrorCode == 0){
					$('.order').html('取消关注').addClass('orderstate');
					$('.ordernum').addClass('ordernumstate').text(parseInt($('.ordernum').text(), 10)+1);
					$('.order-area').addClass('order-areastate');
				}
			
			});
		}
//
	});
	
	//关闭弹幕按钮点击事件
	$('.danmubtn').click(function(){
		if(vpage.danmubtn){
			vpage.danmubtn = false;
			$(this).addClass('off')//.text('开启弹幕');
			$('#danmu').empty();		
		}else{
			vpage.danmubtn = true;
			$(this).removeClass('off')//.text('关闭弹幕');
		}
	});
	
	//静音按钮点击事件
	$('.nosound').click(function(){
		var state = $(this).attr('data-state');
		if(state == 'yes'){
			$(this).attr('data-state','no').addClass('off');
			if(player)
			player.setVolume(0);
		}else{
			$(this).attr('data-state','yes').removeClass('off');
			if(player)
			player.setVolume(1);
		}
	});
	
	//全屏按钮点击事件
	$('.fullscreenbtn').click(function(){
		if(document.documentElement.requestFullScreen) {
			document.documentElement.requestFullScreen();
		} else if(document.documentElement.mozRequestFullScreen) {
			document.documentElement.mozRequestFullScreen();
		} else if(document.documentElement.webkitRequestFullScreen) {
			document.documentElement.webkitRequestFullScreen();
		}else if (document.documentElement.msRequestFullscreen) {
			document.documentElement.msRequestFullscreen();
		}
	});
	//退出全屏按钮点击事件
	$('.unfullscreenbtn').click(function(){
		clearTimeout(timer);
		if(document.cancelFullScreen) {
			document.cancelFullScreen();
		} else if(document.mozCancelFullScreen) {
			document.mozCancelFullScreen();
		} else if(document.webkitCancelFullScreen) {
			document.webkitCancelFullScreen();
		}else if (document.msExitFullscreen) {
			document.msExitFullscreen();
		}	
	});
	//点击视频双击全屏

	$('#room-video').dblclick(function(){	
		if (IsFullScreen == true) {
			var element = document.documentElement;
			if(element.requestFullScreen) {
				element.requestFullScreen();
			} else if(element.mozRequestFullScreen) {
				element.mozRequestFullScreen();
			} else if(element.webkitRequestFullScreen) {
				element.webkitRequestFullScreen();
			}else if (element.msRequestFullscreen) {
				element.msRequestFullscreen();
			}
			IsFullScreen = false;
		}else{
			clearTimeout(timer);
			if(document.cancelFullScreen) {
				document.cancelFullScreen();
			} else if(document.mozCancelFullScreen) {
				document.mozCancelFullScreen();
			} else if(document.webkitCancelFullScreen) {
				document.webkitCancelFullScreen();
			}else if (document.msExitFullscreen) {
				document.msExitFullscreen();
			}	
			IsFullScreen = true;
		}
		return IsFullScreen;
	})
	//为聊天区域绑定滚动条样式
	$('#message11').mCustomScrollbar({
		autoHideScrollbar:true,
		scrollbarPosition:"outside"
	});
	
	//获取直播地址并执行播放器初始化getinflow
	
//	reqAjax('/page/getinflow',{AnchorId:vpage.anchorId},function(data){
		if(Giflow){
			player = new prismplayer({
			    id: "J_prismPlayer", // 容器id
			    source: Giflow,
			    autoplay: true,    
				width: "100%",       
				height: "100%",
				isLive: true,
				waterMark: "http://shiliulianmeng.com/static/images/chat/watermark.png|TL|0.2|0.5",
			    skinLayout:false
			});
			
			player.on('liveStreamStop',function(){
				player = null;
				$('#J_prismPlayer').empty();
				$('#livestop').removeClass('none');
				
				$('#playerretry').click(function(){
					$('#livestop').addClass('none');
					player = new prismplayer({
					    id: "J_prismPlayer", // 容器id
					    source: data.Data.errMsg,
					    autoplay: true,    
						width: "100%",       
						height: "100%",
						isLive: true,
						waterMark: "http://www.shiliulianmeng.com/static/images/chat/watermark.png|TL|0.2|0.5",
					    skinLayout:false
					});
					player.on('ready',function(){
						var jinyin=$('.nosound').attr('data-state');
						if(jinyin == 'yes'){
							$(this).attr('data-state','yes').removeClass('off');
							if(player)
							player.setVolume(1);
						}else{
							$(this).attr('data-state','no').addClass('off');
							if(player)
							player.setVolume(0);
							
						}
					})
					
				});
				$('#returnindex').click(function(){
					location = '/';
				});
			});
		}else{
			$('#J_prismPlayer').remove();
			$('#room-video').css('background','black');
			$('#livestop').removeClass('none');
				
				$('#playerretry').click(function(){
					player = new prismplayer({
					    id: "J_prismPlayer", // 容器id
					    source: data.Data.errMsg,
					    autoplay: true,    
						width: "100%",       
						height: "100%",
						isLive: true,
						waterMark: "http://www.shiliulianmeng.com/static/images/chat/watermark.png|TL|0.2|0.5",
					    skinLayout:false
					});
				});
				$('#returnindex').click(function(){
					location = '/';
				});
		}
//	});

	//表情存放的路径
	$('.emoj').qqFace({
		id: 'facebox',
		assign: 'room-chat-input',
		path: '/static/arclist/' 
	});
	weekTop();
	var t = setInterval("weekTop()", 30000); 
	
	AllTop();
	var t2 = setInterval("AllTop()", 30000); 
	
	hotNum();
	var t3 = setInterval("hotNum()", 30000); 
	//鼠标经过公告展示
	$('#room-chat-notic').hover(function(){
		$(this).addClass('hover');
	},function(){
		$(this).removeClass('hover');
	});
	
	//鼠标点选排行tab标签
	$('.room-ranklist-hd .tab').click(function(){
		var tabnum = $(this).data('tab');
		$(this).addClass('active').siblings().removeClass('active');
		$('.room-rank-'+tabnum).addClass('active').siblings().removeClass('active');
		
	});
	
	//鼠标经过排行榜
	$('.room-ranklist-content').hover(function(){
		$(this).animate({height:"316px"},'fast');
	},function(){
		$(this).animate({height:"124px"},'fast');
	});
	
	//聊天发送按钮
	$('#room-chat-send').click(function(){
		if (!getStorage('Id')){
			$('#loginBox,.logins-alert-bg').css('display','block');
			return;
		}
		
		if(!vpage.allowsendmsg){
			$('#room-chat-sebdMessage-btn').prepend('<div class="dialogmsg">'+vpage.dialogmsg+'</div>');
			return;
		}
		$('.dialogmsg').remove();
		var sendmsg = $('#room-chat-input').val();
		
		if (sendmsg.length >= 80) {
			Dialog("您说的话超出了规定长度!");
			return;
		}
		var uName = userinfo.NickName
		if(uName == null || uName == ""){
			Dialog("请重新登录后再发言!");
			return;
		}
		if (sendmsg != "") {
			sendmsg = checkData(sendmsg);
			sendmsg = Filter(sendmsg);//敏感词过滤
			var msgobj = {
				from: uName,
				type: 'text',
				ext:{
					userinfo:{
						device:'pc',
						userid: getStorage('Id'),
						uid: userinfo.Id,
						avatar: userinfo.Favicon,
						roomclass:userinfo.roomclass,
						level:userlevel(userinfo.Integral).number
					}
				},
				msgbody: sendmsg,
			};
			if(userinfo.Id == vpage.anchorId)
				msgobj.ext.userinfo.roomclass = 1;
				
			var str = JSON.stringify(msgobj);
			if (str) {
				message = new Messaging.Message(str);
				message.destinationName = vpage.roomId;
				client.send(message);
				
				vpage.allowsendmsg = false;
				setTimeout(function(){
					vpage.allowsendmsg = true;
				},2000);
				
				$('#room-chat-input').val("");
			}
			return false;
		}
	});
	//全屏弹幕发送
	$('.Fullscreen-Barrage-btn').click(function(){
		if (!getStorage('Id')){
			$('#loginBox,.logins-alert-bg').css('display','block');
			return;
		}
		
		if(!vpage.allowsendmsg){
			$('#room-chat-sebdMessage-btn').prepend('<div class="dialogmsg">'+vpage.dialogmsg+'</div>');
			return;
		}
		$('.dialogmsg').remove();
		var sendmsg = $('.Fullscreen-Barrage input').val();
		
		if (sendmsg.length >= 80) {
			Dialog("您说的话超出了规定长度!",true);
			return;
		}
		var uName = userinfo.NickName
		if (sendmsg != "") {
			sendmsg = checkData(sendmsg);
			sendmsg = Filter(sendmsg);//敏感词过滤
			var msgobj = {
				from: uName,
				type: 'text',
				ext:{
					userinfo:{
						device:'pc',
						userid: getStorage('Id'),
						uid: userinfo.Id,
						roomclass:userinfo.roomclass,
						level:userlevel(userinfo.Integral).number
					}
				},
				msgbody: sendmsg,
			};
			if(userinfo.Id == vpage.anchorId)
				msgobj.ext.userinfo.roomclass = 1;
				
			var str = JSON.stringify(msgobj);
			if (str) {
				message = new Messaging.Message(str);
				message.destinationName = vpage.roomId;
				client.send(message);
				
				vpage.allowsendmsg = false;
				setTimeout(function(){
					vpage.allowsendmsg = true;
				},2000);
				
				$('.Fullscreen-Barrage input').val("");
			}
			return false;
		}
	});
	
	//按下回车发送聊天内容
	$('#room-chat-input,.Fullscreen-Barrage').bind('keypress',function(e){
		if(e.keyCode == "13"){
			$('#room-chat-send,.Fullscreen-Barrage-btn').click();
		}
	});
	
	//充值按钮点击事件
	$('#recharge').click(function(){
		if(!getStorage('Id')){
			$('#loginBox,.logins-alert-bg').css('display','block');
			
		}else{
//			location = '/user/recharge';
			window.open("/user/recharge");  
		}
	});
	
	//获取礼物列表
//	reqAjax('/page/getgiftlist',{},function(data){
		if(!Grows){
			return;
		}else{
			var giftNamesList = [];
			var giftPriceList = [];
			var giftAccountList = [];
			var giftID = [];
			var iconmod = '';	
			for(var i=0 ; i< Grows.length ;i++){
				var imageviewGif = Grows[i].GiftPicture;
				var imageviewPng = Grows[i].GiftPicStatic;
				var giftName = Grows[i].GiftName;
				giftNamesList.push(giftName);
				var giftPrice = Grows[i].BuyNumber;
				giftPriceList.push(giftPrice);
				var giftAccount = Grows[i].GiftAccount;
				giftAccountList.push(giftAccount);
				var giftId = Grows[i].Id;
				giftID.push(giftId);
				iconmod += '<li class="giftBtn-Pic"><img src="'+imageviewPng+'"/><div class="giftcard">';
				iconmod += '<div class="giftcard_con"><div class="giftcard_img"><img class="giftPic_gif" src="'+imageviewGif+'"/>';
				iconmod += '</div><div class="giftcard_text"><i class="giftname">'+giftName+'</i><i class="giftprice">'+giftPrice+'金币</i><p class="describe">'+giftAccount+'</p>';
				iconmod += '<div class="giftnumbtn" data-giftId="'+giftId;
				iconmod += '" data-giftname="'+giftName;
				iconmod += '" data-imagegif ="'+imageviewGif;
				iconmod +='"><i class="i_color" data-num="10">10</i><i class="i_color" data-num="100" >100</i><i class="i_color" data-num="520">520</i><i class="i_color" data-num="1314">1314</i><input type="text" placeholder="其他" class="numinput"/>';
				iconmod += '<input type="button" value="赠送" class="batchbtn"/></div>';
				iconmod +='</div></div></div></li>';
			}
			$("#room-video-bottom-ul").append(iconmod);
			
			//赠送
			$('#room-video-bottom').find('.giftnumbtn input[type="button"]').click(function(){
				
				var num =	$(this).siblings('.numinput').val();
				var reg = /^[1-9]\d*$/;
				if(!reg.test(num)){
					Dialog("请输入正整数");
					return 
				}
				if(num == ""||num=="0"){
					return
				}
				var giftId = $(this).parent().data('giftid');
				var giftName = $(this).parent().data('giftname');
				var giftimgurl = $(this).parent().data('imagegif');
				
				var crm = parseInt($(".comm-remain-money").html());
				var crn = parseInt($(".comm-remain-nums").html());
				var giftUnitPrice = $(this).parents('.giftcard_text').find('.giftprice').html();
				giftUnitPrice = parseInt(giftUnitPrice.substring(0, giftUnitPrice.length - 2)); 
				
				//判断是否登陆
				if (!getStorage('Id')){
					$('#loginBox,.logins-alert-bg').css('display','block');
					return;
				}
	
				reqAjax('/user/givegiftnumadd',{GiftId: giftId,UID: vpage.userId,AnchorId: vpage.anchorId,Number: num},function(data){
					if(data.ErrorCode != 0) {
						Dialog("当前余额不足",true,'获取','取消',function(){
							location = '/user/recharge';
						});
					}else{
						$(".comm-remain-money").html(crm - giftUnitPrice * num);
//						$(".comm-remain-nums").html(crn + (giftUnitPrice * num * 0.1));
						var uName = userinfo.NickName
						var sendmsg = '感谢' + uName + '送给主播' + giftName +num +'个';
						var msgobj = {
							from: uName,
							type: 'gift',
							msgbody: sendmsg,
							img_src:giftimgurl
						};
						var str = JSON.stringify(msgobj);
						if (str) {
							message = new Messaging.Message(str);
							message.destinationName = vpage.roomId;
							client.send(message);
						}
					}
				},true);
				$(this).siblings('.numinput').val("");
				$(".giftBtn-Pic").removeClass('giftBtn-Pic-Cur');
			});
	
			$('#room-video-bottom').find('.giftnumbtn i').click(function(){
				var num =  $(this).data('num');
				var giftId = $(this).parent().data('giftid');
				var giftName = $(this).parent().data('giftname');
				var giftimgurl = $(this).parent().data('imagegif');
				
				var crm = parseInt($(".comm-remain-money").html());
				var crn = parseInt($(".comm-remain-nums").html());
				var giftUnitPrice = $(this).parents('.giftcard_text').find('.giftprice').html();
				giftUnitPrice = parseInt(giftUnitPrice.substring(0, giftUnitPrice.length - 2)); 
	
				//判断是否登陆
				if (!getStorage('Id')){
					$('#loginBox,.logins-alert-bg').css('display','block');
					return;
				}
	
				reqAjax('/user/givegiftnumadd',{GiftId: giftId,UID: vpage.userId,AnchorId: vpage.anchorId,Number: num},function(data){
					if(data.ErrorCode != 0) {
						
							Dialog("您的余额不足！",true,'获取','取消',function(){
							location = '/user/recharge';
							});
						
					}else{
						$(".comm-remain-money").html(crm - giftUnitPrice * num);
//						$(".comm-remain-nums").html(crn + (giftUnitPrice * num * 0.1));
						var uName = userinfo.NickName
						var sendmsg = '感谢' + uName + '送给主播' + giftName +num +'个';
						var msgobj = {
							from: uName,
							type: 'gift',
							msgbody: sendmsg,
							img_src:giftimgurl
						};
						var str = JSON.stringify(msgobj);
						if (str) {
							message = new Messaging.Message(str);
							message.destinationName = vpage.roomId;
							client.send(message);
						}
					}
				},true);
				$(this).siblings('.numinput').val("");
				$(".giftBtn-Pic").removeClass('giftBtn-Pic-Cur');
			});
			
			//鼠标经过礼物图标效果
			$(".giftBtn-Pic").hover(function(){
				$(this).addClass('giftBtn-Pic-Cur');
			},function() {
				$(this).removeClass('giftBtn-Pic-Cur');
			});
			
			//点击礼物图标
			$('#room-video-bottom-ul').find('.giftBtn-Pic > img').click(function(){
				//判断是否登陆
				if (!getStorage('Id')){
					$('#loginBox').css('display','block');
					return;
				}
				var index = $(".giftBtn-Pic > img").index($(this));
				var giftName = giftNamesList[index];
				var giftimgurl = $(this).siblings('.giftcard').find('.giftPic_gif').attr('src');			
				var giftUnitPrice = giftPriceList[index];
				var crm = parseInt($(".comm-remain-money").html());
				var crn = parseInt($(".comm-remain-nums").html());
				
				reqAjax('/user/givegiftnumadd',{GiftId: giftID[index],UID: vpage.userId,AnchorId: vpage.anchorId,Number: 1},function(data){
					if(data.ErrorCode != 0) {
						Dialog('您的余额不足！',true,'获取','取消',function(){
							location = '/user/recharge';
						});
					}else{
						$(".comm-remain-money").html(crm - giftUnitPrice);
//						$(".comm-remain-nums").html(crn + (giftUnitPrice * 0.1));
						var uName = userinfo.NickName
						var sendmsg = '感谢' + uName + '送给主播' + giftName+1+'个';
						var msgobj = {
							from: uName,
							type: 'gift',
							msgbody: sendmsg,
							img_src:giftimgurl
						};
						var str = JSON.stringify(msgobj);
						if (str) {
							message = new Messaging.Message(str);
							message.destinationName = vpage.roomId;
							client.send(message);
						}
					}
				},true);
				
			});
		}
//	});	
	
	//竞猜相关 by bi periodinfo
		if(aaa == null){
			$('#guess').css("display","none");
		}else {
			$('#guess').css("display","block");
			$('#guess').click(function(){
				$('#guess').css("display","none");
				if(!vpage.userId){
					$('#loginBox,.logins-alert-bg').css('display','block');
					return;
				}else{
					$(this).css("display", "none");
					$("#guess-kinds-1,#guessing-box").css("display", "block");
					$(".Gusse-choices11 li:first").mouseover();
				}
				if(!ATeam && !BTeam){
					var vsmod1 = '<div>'+ATeam+'</div>'
					var vsmod2 = '<div>'+BTeam+'</div>'
					$('#guess-kinds-1 .bothsides .bothsides-1').html(vsmod1);
					$('#guess-kinds-1 .bothsides .bothsides-2').html(vsmod2);
					//遍历模板
					$('.abcd').remove();
					$.each(aaa, function(i,obj){
						var plistmod = '<li class="abcd" data-proid="'+obj.ProductId+'" data-perid="'+obj.PeriodsId+'">'+obj.ProductName+'</li>';
						$('.Gusse-choices11').append(plistmod);
					});
					$('.Gusse-choices11').on('click','li',function() { 
				        $("#guess-kinds-li").empty();
				        $("#guess-kinds-2").css("display", "block");
						$('#guess-kinds-1').css('display','none');
						ProId = $(this).data('proid');
						PerId = $(this).data('perid');
						
						reqAjax('/page/getnowperioddetails',{PeriodId: PerId,ProductId: ProId},function(data){
							for(var i = 1; i <= 8; i++){
								var state = "State" + i;
								state = data.Data[state];
								var hot = "State" + i + "Hot";
								hot = data.Data[hot];
								ProId = data.Data.ProductId;
								var odds = "State" + i + "Odds";
								odds = data.Data[odds];
								guessparams.PeriodsId  = data.Data.PeriodsId;
								if(odds!=0){
									$('#guess-kinds-txt').html(data.Data.ProductName);
									var str = '';
									str += '<div class="row" data-num="'+ i +'">';
									str += '<i><em class="dot"></em></i>';
									str += '<i>' + state + '</i>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;';
									str += '<i>赔率:</i><i class="odds">' + odds + '</i>&nbsp;&nbsp;<i>热度:' + hot + '</i></div>';
									$("#guess-kinds-li").append(str);
	
								}
							}
							$('#guess-kinds-2').find('.row').click(function(){
								$(this).addClass('hover').siblings().removeClass('hover');
							});
						});
					});
				}
			});


		}
	
	$("#guess-kinds-1-Btn").click(function() {
		$("#guess").css("display", "block");
           $("#guess-kinds-1,#guessing-box").css("display", "none");
      	$('.slz').remove();
	});
	
	$("#guess-kinds-2-back").click(function() {
//		$("#guess").css("display", "block");
        $("#guess-kinds-2").css("display", "none");
		$("#guess-kinds-1").css("display", "block");
		$('.slz').remove();
	});
	
	$("#guess-kinds-2-yes").click(function() {	
		var val=$('#guess-kinds-li').find('.hover').data('num');
		if (!val){
			Dialog("请选择投注项目",true);
			return;
		}
		var num = $("#guess-kinds-2-Num").val();
		if(num == "") {
			Dialog("请输入投注数量",true);
			return;
	    }
		if(num<=0||num!=parseInt(num)||num>50000){
			Dialog("请输入小于50000的正确数字",true);
			return;
		}
		
		guessparams.SupporNumber = $("#guess-kinds-2-Num").val();
		guessparams.Odds = $('#guess-kinds-li').find('.hover .odds').text();	
		var procode = ProId;
		if(ProId < 10){
			procode = '0'+ProId;
		}
		if(val < 10){
			val = '0'+ val;
		}
		guessparams.SupEncoding =  '#'+procode+'>'+val;
		guessparams.UID = vpage.userId;
		var str = '<p class="slz">是否确认支持<span style="color:#FF0861">' + guessparams.SupporNumber + '</span>石榴籽？</p>';
		$("#guess-kinds-Finalyes").css("display", "block");
		$("#guess-kinds-Finalyes").append(str);
		$("#guess-kinds-2").css("display", "none");
		$("#guess-kinds-1").css("display", "none");
	});
	
	$("#guess-kinds-Finalyes-no").click(function(){
		$("#guess-kinds-Finalyes").css("display", "none");
		$("#guess-kinds-2").css("display", "block");
		$('.slz').remove();
	});
	
	$("#guess-kinds-Finalyes-yes").click(function() {
		if(!vpage.userId){
			$('#loginBox,.logins-alert-bg').css('display','block');
			return;
		}
		$('.slz').remove();
		reqAjax('/user/supportadd',guessparams,function(data){
			if(data.ErrorCode == 0){
				// Dialog("恭喜你投注成功",true);
				// $('#guess').css("display", "block");
				// $("#guess-kinds-Finalyes,#guessing-box").css("display", "none");
				$('#guess-kinds-Success').css("display","block");
				$("#guess-kinds-Finalyes").css("display", "none");
			}else{
				Dialog(data.ErrorMsg,true);
				$('#guess').css("display", "block");
				$("#guess-kinds-Finalyes,#guessing-box").css("display", "none");
			}
		});
		$('.slz').remove();
	});
	$('#guess-kinds-Success-yes').click(function() {
		$('#guess-kinds-Success').css("display","none");
		$('#guess-kinds-1').css("display","block");
		$('.slz').remove();
});
	$('#guess-kinds-Success-no').click(function() {
		location.href = '/user/my_guess';
		$('.slz').remove();

	});
	$('#guess-kinds-Success-back').click(function() {
		$("#guess-kinds-Success,#guessing-box").css("display","none");
		$('#guess').css("display", "block");
	})
});

//当客户端连接到服务器时通知客户端
var onConnect = function(frame){
	client.subscribe(vpage.roomId);
	vpage.chatstate = true;
	//当聊天服务器连接成功后发送欢迎语
	if(vpage.chatstate&&vpage.userId){
		var nickname = getStorage('nicheng') ? getStorage('nicheng') : getStorage('username');
		var text = '欢迎';
		var msgobj = {
			from:nickname,
			type:'welcome',
			msgbody:text,
			ext:{
				userinfo:{
					userid:userinfo.Id
				}
			}
		};
		var str = JSON.stringify(msgobj);
		if (str && nickname) {
			message = new Messaging.Message(str);
			message.destinationName = vpage.roomId;
			client.send(message);
		}
	}	
};

//接收信息
var onMessageArrived = function(message) {
	var msg = $.parseJSON( message.payloadString );
	console.log(msg);
	var msgb = msg.msgbody;
	
	var str = '';
	if(msg.type == 'text'){   //纯文字聊天
		str += '<li class="chat-msg" data-cid="' + msg.ext.userinfo.userid 
		+ '" data-name="' + msg.from 
		+ '" data-pid="' + msg.ext.userinfo.uid
		+ '" data-class="' + msg.ext.userinfo.roomclass
		+ '"  data-level="' + msg.ext.userinfo.level
		+ '"  data-avatar="' + msg.ext.userinfo.avatar
		+ '" >';

		if(msg.ext.userinfo.device == 'mobile'){
			str += '<i class="chat-tags-mobile"></i>';
		}
		if(msg.ext.userinfo.roomclass == 1 || msg.ext.userinfo.roomclass == 2){
			str += '<i class="chat-tags-anchor anchor-'+msg.ext.userinfo.roomclass+'"></i>';
		}
		if(msg.ext.userinfo.level || msg.ext.userinfo.level=='0'){
			if(msg.ext.userinfo.roomclass != 1)
			str += '<i class="chat-tags-level level-'+msg.ext.userinfo.level+'"></i>';
		}
		str += '<span class="chat-uname">'+msg.from+'：</span>';
		str += '<span class="chat-content">'+msg.msgbody+'</span>';
		str=replace_em(str);
		$('#message11').find('.mCSB_container').append(str);
		$('#message11').mCustomScrollbar('scrollTo','bottom',{scrollInertia:200});
		var	rText = msg.msgbody;
		rText =replace_em(rText);
		if(msg.ext.userinfo.userid == vpage.userId){
			showdanmu(rText,true);
		}else{
			showdanmu(rText);
		}
		
	}else if(msg.type == 'gift'){ //赠礼相关
		str += '<li class="chat-msg" style="color:#e8983e" data-name="' + msg.from + '">' + msg.msgbody + '</li>';
		$('#message11').find('.mCSB_container').append(str);
		$('#message11').mCustomScrollbar('scrollTo','bottom',{scrollInertia:200});
		showdanmu(msg.msgbody,null,msg.img_src);
		
	}else if(msg.type == 'welcome'){ //欢迎提示语
		if(msg.ext.userinfo.userid == userinfo.Id){
			return;
		}
		str += '<li class="chat-msg" data-name="' + msg.from + '">欢迎' + msg.from + '来到直播间</li>';
		$('#message11').find('.mCSB_container').append(str);
		$('#message11').mCustomScrollbar('scrollTo','bottom',{scrollInertia:200});
	
	
	}else if(msg.type == 'banned'){ //禁言相关
		var uuname = msg.ext.info.username;
		if(msg.ext.info.userid == userinfo.Id){
			vpage.allowsendmsg = false;
			vpage.dialogmsg = '你已被禁言';
			uuname = '您';
			setTimeout(function(){
					vpage.allowsendmsg = true;
			},900000);
		}
		str += '<li class="chat-msg" data-name="' + msg.from + '">' + uuname + msg.msgbody + '15分钟</li>';
		$('#message11').find('.mCSB_container').append(str);
		$('#message11').mCustomScrollbar('scrollTo','bottom',{scrollInertia:200});

	}else if(msg.type == 'welcome'){ //进入房间相关
		str += '<li class="chat-msg" data-name="' + msg.from + '">欢迎<i style="color:#ff0861">' + msg.from + '</i>进入房间</li>';
		$('#message11').find('.mCSB_container').append(str);
		$('#message11').mCustomScrollbar('scrollTo','bottom',{scrollInertia:200});

	}else if(msg.type == 'promote'){ //提升管理员相关
		if(msg.ext.info.userid == userinfo.Id){
			vpage.authority = true;
			userinfo.roomclass = 2;
		}
		
	}else if(msg.type == 'revoke'){//撤销管理员相关
		if(msg.ext.info.userid == userinfo.Id){
			vpage.authority = false;
			userinfo.roomclass = 0;
		}
	}
}

//失去连接
var onConnectionLost = function(responseObject) {
	//if(responseObject.errorCode !== 0) {
	//debug(client.clientId + ": " + responseObject.errorCode + "\n");
	//}
	console.log('聊天服务器已断开');
}

//连接失败
var onFailure = function(failure) {
	failact();
}


function showdanmu(text,self,imgsSrc){
	if(!vpage.danmubtn){
		return;
	}

	if(self){
		var resDanmu = $('<i class="danmuM danmuM_self" style="top:-200px">');
	}else{
		var resDanmu = $('<i class="danmuM" style="top:-200px">');
	}
	if(imgsSrc){
		resDanmu.html('<img src="'+imgsSrc+'" class="liwuImg"/>'+text);
	}else{
		resDanmu.html(text);
	}
	
	$('#danmu').append(resDanmu);
	setTimeout(function(){
		resDanmu.find('.emoji_img').remove();
		var rTop= Math.random()*95;
		resDanmu.css({
			'width': resDanmu.outerWidth(),
			'top': rTop+'%',
			'right': '-'+resDanmu.width()+'px'
		});
		resDanmu.animate({right:'100%'},10000,"linear",function(){
			$(this).remove();
		});
	},100);
}

function Filter(word){
	if (typeof wordmap === 'undefined') {
		return word;
	} else {
		for(var i in wordmap){
			while (word.indexOf(wordmap[i]) > -1) {
				var filterwordlen = wordmap[i].length;
				var after = '';
				for(var n=0;n<filterwordlen;n++){
					after+='*';
				}
				word = word.replace(wordmap[i],after);
			}
		}
		return word;
	}
}
var tfvideo=true;
function fullscreenstate(){
	$('.right_area').on('mousemove',function(e){
		var thispage = e.pageX+','+e.pageY;
		
		if(thispage == mousemove){
			return;
		}
		
		if(tfvideo){
			clearTimeout(timer);
			$('#room-video-controller').css('height','40px');
			$('html').css({  
				cursor: 'default'  
			});
			
			
			timer = setTimeout(function(){
				$('html').css({cursor: 'none'});
				$('#room-video-controller').css('height',0);
			}, 5000);
			mousemove = e.pageX+','+e.pageY;
		}
		
		$('#room-video-controller').hover(function(){
			tfvideo=false;
			clearTimeout(timer);
			$('#room-video-controller').css('height','40px');
			$('html').css({  
				cursor: 'default'  
			});
		},function(){
			tfvideo=true;
		})
		
	});
}

function unfullscreenstate(){
	$('#room-video-controller').css('height','40px');
	$('html').css('cursor','default');
	clearTimeout(timer);
	setTimeout(function() {
		$('#message11').mCustomScrollbar('scrollTo','bottom',{scrollInertia:200});
	},100);	
	$('.right_area').off();
}

//全屏状态监听
document.addEventListener("webkitfullscreenchange", function () {	
	if(document.webkitIsFullScreen){
		timer = setTimeout(function(){
			$('html').css({cursor: 'none'});
			$('#room-video-controller').css('height',0);
		}, 5000);
		fullscreenstate();
	}else{
		unfullscreenstate();
	}
});

document.addEventListener("mozfullscreenchange", function () {	
	if(document.mozFullScreen){
		timer = setTimeout(function(){
			$('html').css({cursor: 'none'});
			$('#room-video-controller').css('height',0);
		}, 5000);
		fullscreenstate();
	}else{
		unfullscreenstate();
	}
});

document.addEventListener("MSFullscreenChange", function () {
	if(document.msFullscreenElement){
		timer = setTimeout(function(){
			$('html').css({cursor: 'none'});
			$('#room-video-controller').css('height',0);
		}, 5000);
		fullscreenstate();
	}else{
		unfullscreenstate();
	}
});

function weekTop(){
	//获取周排行榜getgiftgiveweeks
		if(!Gggw){
			return;
		}else{
			var mod = '<div class="room-rank-li-top3">';
			if(Gggw.length > 1){
				mod += '<div class="user-top3 rank-2"><div class="user-avater"></div><img class="user-avater-img" src="'+(Gggw[1].Favicon?Gggw[1].Favicon:'/static/images/default_avatar.jpg')+'"/><div class="user-nickname">' + Gggw[1].NickName + '</div><div class="user-level level-' + userlevel(Gggw[1].Integral).number + '"></div></div>';
			}else{
				mod += '<div class="user-top3 rank-2"><div class="user-avater"></div><div class="user-nickname">--虛位以待--</div></div>';
			}	
			if(Gggw.length > 0){
				mod += '<div class="user-top3 rank-1"><div class="user-avater"></div><img class="user-avater-img" src="'+(Gggw[0].Favicon?Gggw[0].Favicon:'/static/images/default_avatar.jpg')+'"/><div class="user-nickname">' + Gggw[0].NickName + '</div><div class="user-level level-' + userlevel(Gggw[0].Integral).number + '"></div></div>';
			}else{
				mod += '<div class="user-top3 rank-1"><div class="user-avater"></div><div class="user-nickname">--虛位以待--</div></div>';
			}
			if(Gggw.length > 2){
				mod += '<div class="user-top3 rank-3"><div class="user-avater"></div><img class="user-avater-img" src="'+(Gggw[2].Favicon?Gggw[2].Favicon:'/static/images/default_avatar.jpg')+'"/><div class="user-nickname">' + Gggw[2].NickName + '</div><div class="user-level level-' + userlevel(Gggw[2].Integral).number + '"></div></div>';
			}else{
				mod += '<div class="user-top3 rank-3"><div class="user-avater"></div><div class="user-nickname">--虛位以待--</div></div>';
			}	
			mod += '</div>';
			for(var i=3;i<10;i++){
				if(i<Gggw.length){
					mod += '<div class="room-rank-li"><i class="room-rank-order">'+(i+1)+'</i><i class="room-rank-user-level level-' + userlevel(Gggw[i].Integral).number + '"></i><i class="room-rank-user-nickname">' + Gggw[i].NickName + '</i></div>';
				}else{
					mod += '<div class="room-rank-li"><i class="room-rank-order">'+(i+1)+'</i><i class="room-rank-user-nickname">--虛位以待--</i></div>';
				}
			}
			$('.room-rank-tab1').html(mod);
		}
}
//获取总排行榜getgiftgiveall
function AllTop(){
		if(!Gggm){
			return;
		}else{
			var mod = '<div class="room-rank-li-top3">';
			if(Gggm.length > 1){
				mod += '<div class="user-top3 rank-2"><div class="user-avater"></div><img class="user-avater-img" src="'+(Gggm[1].Favicon?Gggm[1].Favicon:'/static/images/default_avatar.jpg')+'"/><div class="user-nickname">' + Gggm[1].NickName + '</div><div class="user-level level-' + userlevel(Gggm[1].Integral).number + '"></div></div>';
			}else{
				mod += '<div class="user-top3 rank-2"><div class="user-avater"></div><div class="user-nickname">--虛位以待--</div></div>';
			}	
			if(Gggm.length > 0){
				mod += '<div class="user-top3 rank-1"><div class="user-avater"></div><img class="user-avater-img" src="'+(Gggm[0].Favicon?Gggm[0].Favicon:'/static/images/default_avatar.jpg')+'"/><div class="user-nickname">' + Gggm[0].NickName + '</div><div class="user-level level-' + userlevel(Gggm[0].Integral).number + '"></div></div>';
			}else{
				mod += '<div class="user-top3 rank-1"><div class="user-avater"></div><div class="user-nickname">--虛位以待--</div></div>';
			}
	
			if(Gggm.length > 2){
				mod += '<div class="user-top3 rank-3"><div class="user-avater"></div><img class="user-avater-img" src="'+(Gggm[2].Favicon?Gggm[2].Favicon:'/static/images/default_avatar.jpg')+'"/><div class="user-nickname">' + Gggm[2].NickName + '</div><div class="user-level level-' + userlevel(Gggm[2].Integral).number + '"></div></div>';
			}else{
				mod += '<div class="user-top3 rank-3"><div class="user-avater"></div><div class="user-nickname">--虛位以待--</div></div>';
			}	
			mod += '</div>';
			for(var i=3;i<10;i++){
			
				if(i<Gggm.length){
					mod += '<div class="room-rank-li"><i class="room-rank-order">'+(i+1)+'</i><i class="room-rank-user-level level-' + userlevel(Gggm[i].Integral).number + '"></i><i class="room-rank-user-nickname">' + Gggm[i].NickName + '</i></div>';
				}else{
					mod += '<div class="room-rank-li"><i class="room-rank-order">'+(i+1)+'</i><i class="room-rank-user-nickname">--虛位以待--</i></div>';
				}
			}
			
			$('.room-rank-tab2').html(mod);
		}
}
function hotNum(){
//	getroom
		$('.hotnum').text(state[0].GiftNum);
		$('.viewnum').text(state[0].LiveNumber);
}
function timerliu(intDiff){

	window.setInterval(function(){
		var day=0,
			hour=0,
			minute=0,
			second=0;//时间默认值		
		if(intDiff > 0){
			day = Math.floor(intDiff / (60 * 60 * 24));
			hour = Math.floor(intDiff / (60 * 60)) - (day * 24);
			minute = Math.floor(intDiff / 60) - (day * 24 * 60) - (hour * 60);
			second = Math.floor(intDiff) - (day * 24 * 60 * 60) - (hour * 60 * 60) - (minute * 60);
		}
		if (minute <= 9) minute = '0' + minute;
		if (second <= 9) second = '0' + second;
		$('#day_show').html(day+"天");
		$('#hour_show').html('<s id="h"></s>'+hour+'时');
		$('#minute_show').html('<s></s>'+minute+'分');
		$('#second_show').html('<s></s>'+second+'秒');
		intDiff--;
		
		if(intDiff ==0){
			$(".time-item").css("display", "none");
			$("#guess").css("display", "none");
		}
		}, 1000);
}

function RoomConcernDel(uid, aid){
	reqAjax('/user/cancelroomcon',{User:aid,UID:uid},function(data){
		if(data.ErrorCode==0){
			Dialog("您成功取消关注");
			$(".order").html("关注").removeClass('orderstate');
			$(".ordernum").removeClass('ordernumstate').text(parseInt($('.ordernum').text(), 10)-1);
			$(".order-area").removeClass('order-areastate');
		}else{
			Dialog(data.errMsg);
		}
	});
}
function liaotianlianjie(){
	//聊天服务相关
	if(!window.WebSocket){
		$('#message11').html('<p style="color:red">因当前浏览器版本过低，聊天功能可能无法正常使用，请尝试升级或更换浏览器</p>');
	}else{
		if(getStorage('Id')){
			reqAjax('/page/getroomim',{UID:vpage.userId},function(ret){
				if(ret.Data !=""){
					chatconf.user = ret.Data.username;
					$.getScript('/static/js/md5.js',function(){
						chatconf.password = hex_md5(ret.Data.password);
					});	
				}
			});
		}
		
		$.getScript('/static/js/fushowim.js',function(){
			client = new Messaging.Client(chatconf.host, Number(chatconf.port), chatconf.clientId);
			client.connect({
				userName: chatconf.user,
				password: chatconf.password,
				onSuccess: onConnect,
				onFailure: onFailure
			});
			client.onConnect = onConnect;
			client.onMessageArrived = onMessageArrived;
			client.onConnectionLost = onConnectionLost;
		});	
	}
}



function failact(){
	setTimeout(function(){
		if(getStorage('Id')){
			reqAjax('/page/getroomim',{UID:getStorage('Id')},function(ret){
				client.connect({
					userName: ret.Data.username,
					password: hex_md5(ret.Data.password),
					onSuccess: onConnect,
					onFailure: onFailure
				});
			});
			
		}else{
			client.connect({
				userName: chatconf.user,
				password: chatconf.password,
				onSuccess: onConnect,
				onFailure: onFailure
			});
		}
	},5000)
}

