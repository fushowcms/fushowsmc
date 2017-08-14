//热门直播ajax
	var code =window.localStorage.getItem("sqcode");
	var state =window.localStorage.getItem("sqstate");
   	$(function() {
		
		var url;


		if(code!=""&&state=="weixin"){
			url =  "/page/weiXinUserInfo";
			shouquan(url);
		}
		if(code!=""&&state=="qq"){
			url =  "/page/qQUserInfo";
			shouquan(url);
		}
		if(code!=""&&state=="weibo"){
			url =  "/page/wBGetUserInfo";
			shouquan(url);
		}
		
		
		var code1 =getPar("code");
		var state1 =getPar("state");
		if(code1!=""&&state1!=""){
			window.localStorage.setItem("sqcode", code1);
			window.localStorage.setItem("sqstate", state1);
			window.location.href = '/';	
			return ;
		}
	}); 

    $(".gotop-btn").click(function(){
        var speed=400;//滑动的速度
        $('body,html').animate({ scrollTop: 0 }, speed);
        return false;
    });
	
	$('#indexShouye').css('background','#df0050').find('a').eq(0).css('color','#fff');
$(document).ready(function() {
	var height = $('.index-suggestion').position().top;
	$(window).scroll(function() {
		var scroll = $(window).scrollTop();
		if(scroll >= height) {
			$(".index-lift .gotop-btn").fadeIn(1000);
		} else {
			$(".index-lift .gotop-btn").fadeOut(1000);
		}
	})
	$(".index-lift .gotop-btn").click(function() {
		$('body,html').animate({
			scrollTop: 0
		}, 1000);
		return false;
	})
	$('#lunbo-Btn').on('click', '#lunbo-Btn li', function() {
		var that = this;
		$('#lunbo-Btn li div').removeClass('pics-Border2');
		$(this).children('div').addClass('pics-Border2');
	});
	$('.index-channel>div:even').addClass('index-channel-even');
	
})