$(function(){
	//sidebar分类 addby liuhan
//	reqAjax("/page/findCategoryAll",{},function(msg){
		
//		if(msg.ErrorCode == 0){
//			for(var i=0;i<msg.Data.length;i++){
//				var	str = ""
//				str = '<div class="sidebar-category-list-container">'
//				str += '<div class="sidebar-list-head">'
//				str += '<span class="sidebar-list-head-tit oneCategory" data-ename="jingji">'+msg.Data[i].COneName+'</span>'
//				str += '</div>'	  
//				str += '<div class="sidebar-site-category">'
//				str += '<ul class="sidebar-list sidebar-list-1 clearfix twoCategory">'                  
//				for(var j=0;j<msg.Data[i].ArrCategory.length;j++){				
//					str += '<li class="sidebar-site-category-item" data-ename="lol">'
//				    str += ' <a href="'+msg.Data[i].ArrCategory[j].TwoCategoryAddress+'" target="_top" title="'+msg.Data[i].ArrCategory[j].TwoCategoryName+'">'+msg.Data[i].ArrCategory[j].TwoCategoryName+'</a>'
//					str += '</li>'
//				}
//				str += '</ul></div></div>'
//				$('.sidebar-category').append(str)
//			}
//		}
//	},true);


	reqAjax("/page/getSdbadvertising",{},function(msg){
		console.log(msg)
		if(msg.ErrorCode!=0) {
			Dialog(msg.ErrorMsg,true,"确定",null,function() {
				$('.dialog').remove();
			},null);
		}else {
			var str = '';
			if(msg.Data != null){
				$.each(msg.Data, function(index,eq) {
					if(eq.SdbadURL.indexOf("http://") != -1){
						str += "<a onclick=\"window.open(\'"+eq.SdbadURL+"\')\" href='javascript:void(0)'><li><img src='"+eq.PicURL+"'/></li></a>";
					}else{
						str += "<a onclick=\"window.open(\'http://"+eq.SdbadURL+"\')\" href='javascript:void(0)'><li><img src='"+eq.PicURL+"'/></li></a>";
					}
					$('.sidebar-box-classification-two').append(str);
					str = '';
				});
			}else{
				return false;
			}	
		}
	},true);


	var test = window.location.pathname;
	if(test=="/page/listsClass"){
		$('.sidebar-head-nav-item-cate').css('background-color','#364040');
	}else if(test=="/page/lists"){
		$('.sidebar-head-nav-item-all').css('background-color','#364040');
	}
	var sidebarshowBtn=true;
	var webwidths=$(window).width();
	if(webwidths<1260){
		$('.sidebar-box').css({'left':'-174px'});
		$('.sidebar-content-bigShow,.sidebar-search,.sidebar-logo,.sidebar-login-boxS,.sidebar-box-classification-two').css({'left':'-66px'});
		$('.sidebar-content-bigShow,.sidebar-search,.sidebar-box-classification-two').css('display','none');
		$('.sidebar-collapse-navList').css('top','0');
		$(".sidebar-collapse-navList").css({'left':'174px'});
		$(".sidebar-shwo-LOGO,.sidebar-collapse-footer").css({'right':'0'});
		$('.sidebar-show-btn').addClass('sidebar-show-btnShow');
		$('.lists-big-box,.right_area').css({'left':'85px'});
		$('.right_area').css('left','85px');
		$('#room-video').css('height',$('.video-Box').width()*9/16);
		$('.sidebar-box-login').css({'left':'-240px'});
		sidebarshowBtn=false;
	}else{
		$('.sidebar-box').css({'left':'0'});
		$('.sidebar-collapse-navList').css('top','-172px');
		$('.lists-big-box').css({'left':'260px'});
		$('.right_area').css('left','260px');
		$('#room-video').css('height',$('.video-Box').width()*9/16);
		$('.sidebar-content-bigShow,.sidebar-search,.sidebar-box-classification-two').css('display','block');
		$('.sidebar-content-bigShow,.sidebar-search,.sidebar-logo,.sidebar-login-boxS,.sidebar-box-classification-two').css({'left':'0'});
		$(".sidebar-shwo-LOGO,.sidebar-collapse-footer").css({'right':'-66px'});
		$(".sidebar-collapse-navList").css({'left':'240px'});
		$('.sidebar-show-btn').removeClass('sidebar-show-btnShow');
		$('.sidebar-box-login').css({'left':'0'});
		sidebarshowBtn=true;
	}
	$('.sidebar-show-btn').click(function(){
		if(sidebarshowBtn){
			$('.sidebar-box').stop().animate({'left':-174},500);
			$('.sidebar-content-bigShow,.sidebar-search,.sidebar-logo,.sidebar-login-boxS,.sidebar-box-classification-two').stop().animate({'left':-66},500,function(){
				
				$('.sidebar-content-bigShow,.sidebar-search,.sidebar-box-classification-two').css('display','none');
				$('.sidebar-collapse-navList').css('top','0');
			});
			$(".sidebar-collapse-navList").stop().animate({'left':174},500);
			$(".sidebar-shwo-LOGO,.sidebar-collapse-footer").stop().animate({'right':0},500);
			$('.sidebar-show-btn').addClass('sidebar-show-btnShow');
			$('.lists-big-box,.right_area').stop().animate({'left':85},500);
			$('.right_area').css('left',85);
			$('#room-video').css('height',$('.video-Box').width()*9/16);
			$('.sidebar-box-login').stop().animate({'left':-240},500);
			sidebarshowBtn=false;
		}else{
			$('.sidebar-box').stop().animate({'left':0},500);
			$('.sidebar-collapse-navList').css('top',-172);
			$('.lists-big-box').stop().animate({'left':260},500);
			$('.right_area').css('left',260);
			$('#room-video').css('height',$('.video-Box').width()*9/16);
			$('.sidebar-content-bigShow,.sidebar-search,.sidebar-box-classification-two').css('display','block');
			$('.sidebar-content-bigShow,.sidebar-search,.sidebar-logo,.sidebar-login-boxS,.sidebar-box-classification-two').stop().animate({'left':0},500);
			$(".sidebar-shwo-LOGO,.sidebar-collapse-footer").stop().animate({'right':-66},500);
			$(".sidebar-collapse-navList").stop().animate({'left':240},500);
			$('.sidebar-show-btn').removeClass('sidebar-show-btnShow');
			$('.sidebar-box-login').stop().animate({'left':0},500);
			sidebarshowBtn=true;
		}	
	})
	userOff('.out');
	$('.sidebar-login-box').mouseover(function(){
		$('.sidebar-box-footer').css('height','auto');
	})
			var uid = getStorage("Id");
			$('.sidebar-content').mCustomScrollbar({
				autoHideScrollbar:true,
				scrollbarPosition:"outside"
			});	
		
			$('.sidebar-box-btn').click(function(){
				if (uid == null) {
					$('.sidebar-login-btn').click();
				}else{
					reqAjax("/page/getuser",{UID:uid},function(msg){
						if(msg.ErrorCode!=0) {
							Dialog(msg.ErrorMsg,true,"确定",null,function() {
								$('.dialog').remove();
							},null);
						}else {
							var type = msg.Data.Type;
							if (type != 1) {
								alertShowYNznx("您还不是主播呦","确定要前往个人中心申请主播？",null,"取消");
								$('.alert-yes').live("click",function () {
									setStorage("from","sidebarunbtn");
									window.location.href= '/user/mine_myInform';
								});
								$('.alert-no').live("click",function () {
									return;
								})
							}else{
								myaddr();
							}
						}
					},true);
				}
			})
		});
		
			function myaddr(){
				
				var uid = getStorage("Id");
				reqAjax("/user/getroominfo",{UID:uid},function(data) {
					if(data.ErrorCode!=0){
						return;
					}else{
						window.location.href= '/user/my_setting';
					}
				},true);
			}
$(window).resize(function(){
	var webwidths=$(window).width();
	if(webwidths<1260){
		$('.sidebar-box').css({'left':'-174px'});
		$('.sidebar-content-bigShow,.sidebar-search,.sidebar-logo,.sidebar-login-boxS,.sidebar-box-classification-two').css({'left':'-66px'});
		$('.sidebar-content-bigShow,.sidebar-search,.sidebar-box-classification-two').css('display','none');
		$('.sidebar-collapse-navList').css('top','0');
		$(".sidebar-collapse-navList").css({'left':'174px'});
		$(".sidebar-shwo-LOGO,.sidebar-collapse-footer").css({'right':'0'});
		$('.sidebar-show-btn').addClass('sidebar-show-btnShow');
		$('.lists-big-box,.right_area').css({'left':'85px'});
		$('.right_area').css('left','85px');
		$('#room-video').css('height',$('.video-Box').width()*9/16);
		$('.sidebar-box-login').css({'left':'-240px'});
		sidebarshowBtn=false;
	}else{
		$('.sidebar-box').css({'left':'0'});
		$('.sidebar-collapse-navList').css('top','-172px');
		$('.lists-big-box').css({'left':'260px'});
		$('.right_area').css('left','260px');
		$('#room-video').css('height',$('.video-Box').width()*9/16);
		$('.sidebar-content-bigShow,.sidebar-search,.sidebar-box-classification-two').css('display','block');
		$('.sidebar-content-bigShow,.sidebar-search,.sidebar-logo,.sidebar-login-boxS,.sidebar-box-classification-two').css({'left':'0'});
		$(".sidebar-shwo-LOGO,.sidebar-collapse-footer").css({'right':'-66px'});
		$(".sidebar-collapse-navList").css({'left':'240px'});
		$('.sidebar-show-btn').removeClass('sidebar-show-btnShow');
		$('.sidebar-box-login').css({'left':'0'});
		sidebarshowBtn=true;
	}
})

//搜索
function searchF(){
	
		var headerSearch = $(".search-key").val();
		if(headerSearch && $.trim(headerSearch)!=null && $.trim(headerSearch)!=''){
			setStorage("cxzbpage",1);
			$('.search-form').submit();
	//					window.location.href = "/page/search?kw="+headerSearch;
		}else{
			Dialog('请输入搜索内容');
		}
	
}
$(".sidebar-search-btn,.search-submit").on("click",function(){
	searchF()
});
//按下回车搜索
$('.search-key').bind('keypress',function(e){
	if(e.keyCode == "13"){
		searchF()
	}
});
