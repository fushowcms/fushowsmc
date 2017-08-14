var opindex;
$(function(){	
	pageinit();
	var url = window.location.href;
	var RoomId = url.substr(url.lastIndexOf('/')+1,url.length);
	reqAjax('/page/getthisroom',{ID:RoomId},function(msg) {
		console.log(msg); 
		if(msg.Data == null) {
			return false;
		} else {
			var id = msg.Data.RId;
			var form = msg.Data.Form;
			$('#roomAlias').html(msg.Data.Title);
			$('#nickName').html('主播:'+msg.Data.NickName);
			$('.hotnum').html(msg.Data.Number);
			$('.source').html('来源：'+msg.Data.Form);
			if(form == "斗鱼"){
				$('#room-video iframe').attr("src",'http://staticlive.douyutv.com/common/share/play.swf?room_id='+id);
				setTimeout(function(){
					$('#room-video iframe').attr("src",'http://staticlive.douyutv.com/common/share/play.swf?room_id='+id);
				},500)
			} else if(form == "虎牙") {
				var arr = id.split("&");
				var chTopId = arr[0].split(":")[1];
				var subChId = arr[1].split(":")[1]
				$('#room-video iframe').attr("src",'http://weblbs.yystatic.com/s/'+chTopId+'/'+subChId+'/huyacoop.swf');
				setTimeout(function(){
					$('#room-video iframe').attr("src",'http://weblbs.yystatic.com/s/'+chTopId+'/'+subChId+'/huyacoop.swf');
				},500)		
			} else if(form == "战旗") {
				$('#room-video iframe').attr("src",'http://www.zhanqi.tv/live/embed?roomId='+id);
			} else if(form == "熊猫"){
				$.ajax({  
					type: "get",
					async: false,
					url: "/pandainfo?roomid="+id, 
					dataType: "json",  
					data:'',
					success: function(data){
	
						var infodata = $.parseJSON(data.Data);
						console.log(infodata);
						$('#room-video iframe').attr("src",'http://s5.pdim.gs/static/46b6dba3c7147b28.swf?roomId='+infodata.data.roominfo.id+'&videoId='+infodata.data.videoinfo.room_key+'&plflag='+infodata.data.videoinfo.plflag+'&display_type=1&isWebPlayer=false');
						setTimeout(function(){
							$('#room-video iframe').attr("src",'http://s5.pdim.gs/static/46b6dba3c7147b28.swf?roomId='+infodata.data.roominfo.id+'&videoId='+infodata.data.videoinfo.room_key+'&plflag='+infodata.data.videoinfo.plflag+'&display_type=1&isWebPlayer=true');
						},500)
					},
					error: function(xhr){
			//			console.log(xhr,url);
						//Dialog('服务连接异常，请稍后重试');
					}
				});			
			}
		}
		
	},true);
	//直播间广告js
	advertisement();
	$(".adcol").bind("click",function(){
		$("#advertisement").empty();
	})
	$(".adcol1").hover(function(){  
        $(this).css("opacity","0");  
    },function(){  
        $(this).css("opacity","1");  
    }); 
	function advertisement(){
	reqAjax('/page/getDbadvertising',{},function(data){
		console.log(data)
		var str1 = "";
		if(data.ErrorCode == 0){
			if(data.Data != null) {
				$.each(data.Data, function(index,eq) {
					if(eq.DbadURL.indexOf("http://") != -1){
						str1 += "<div id='adv"+ (index+1) +"'><a onclick=\"window.open(\'"+eq.DbadURL+"\')\" href='javascript:void(0)'><div class='advs'><img src='"+ eq.PicURL +"'alt=''></div></a></div>";
					}else{
						str1 += "<div id='adv"+ (index+1) +"'><a onclick=\"window.open(\'http://"+eq.DbadURL+"\')\" href='javascript:void(0)'><div class='advs'><img src='"+ eq.PicURL +"'alt=''></div></a></div>";
					}
					
					if (index == 3) {
						return false;
					}
					opindex = index;
				});
				if (opindex == 0) {
					for(var i = 1;i<3;i++){
						str1 += "<div id='adv"+ (i+1) +"'><a href='javascript:void(0)'><div class='advs'></div></a></div>";
					}
				}
				if (opindex == 1) {
					str1 += "<div id='adv"+ 3 +"'><a href='javascript:void(0)'><div class='advs'></div></a></div>";
				}
				$("#advertisement").append(str1);
			}
		}else{
			return false;
		}
			
	});	
}


})