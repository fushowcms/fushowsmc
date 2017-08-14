$(function() {
	setStorage("categorypage","1");
	var local_url = window.location.pathname;
	$("body").css("overflow-x", "hidden");
	var pages = getStorage("categorypage");
	if(pages == null) {
		pages = 1;
	}
	var page;
	$(function() {
		pageinit();
		$('.lists-big-box').css('width', $(window).width() - 260);
		Item_adaptation();
		$(function() {
			reqAjax("/page/getTwoCategoryByAddress",{Address:local_url},function(msg) {	
			var RoomType= msg.Data[0].Id
				if (msg.ErrorCode == 0){
					$('.pClass').html(msg.Data[0].TwoCategoryName)
					reqAjax("/page/getRoomAliasByRoomType",{RoomType:msg.Data[0].Id,page:pages,rows:15},function(msg) {		
						if(msg.ErrorCode==0){
							total = msg.Data.total;
							page = Math.ceil(total / 15);
							if (!msg.Data.data){
								var list_type = document.getElementById('list_type');
								list_type.innerHTML = "<div style='width: 100%;text-align: center;color: gainsboro;font-size: 50px;' class='removeabc'>没有搜到相关信息</div>";
								return;
							}
							var html = '';
							for(var i = 0; i < msg.Data.data.length; i++) {
								str = '<a href="/roomlive/' + msg.Data.data[i].Id + '" class="removeabc"><li><div class="lists-box-list-img"><img class="images-show" src="' + msg.Data.data[i].LiveCover + '" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>'+msg.Data.data[i].RoomAlias+'</h1><span class="lists-classFont">'+msg.Data.data[i].TwoCategoryName+'</span><p>' + msg.Data.data[i].NickName + '<span><img src="/static/images/ui/look-ioc.png">' + msg.Data.data[i].LiveNumber + '</span></p></div></li></a>'
								html += str;
							}
							if(total>=15){
								$(".page-component").createPage({
									RoomType:RoomType,
							        pageCount:page,
							        current:pages,
							        backFn:function(p){
					           			 console.log(p);
						      		  }		
						   	 });
							}
							var list_type = document.getElementById('list_type');
							list_type.innerHTML += html;
							$('.images-show').each(function(index) {
								$('.lists-box-list-img').css('max-height', $(this).width() / 2);
							});
						}
					});
				}
			});
		});
	})

	function Item_adaptation() {
		var winWidth = $(window).width();
		var boxWidth = $('#wrap').width();

		if(winWidth <= 1280) {
			var itemWidth = $('#wrap').width() * 0.315;
			var itemHeight = itemWidth / 1.7;
			var styleElement = '<style>.video-list-item{width:' + itemWidth + 'px !important;';
			styleElement += '}.video-list-item .video-img{height:' + itemHeight + 'px !important;}';
			styleElement += '#list_type{height:' + (itemHeight + 55) + 'px !important;overflow:hidden;}';
			styleElement += '</style>';
		}

		if(winWidth > 1280 && winWidth <= 1600) {
			var itemWidth = $('#wrap').width() * 0.235;
			var itemHeight = itemWidth / 1.7;
			var styleElement = '<style>.video-list-item{width:' + itemWidth + 'px !important;';
			styleElement += '}.video-list-item .video-img{height:' + itemHeight + 'px !important;}';
			styleElement += '#list_type{height:' + (itemHeight + 55) + 'px !important;overflow:hidden;}';
			styleElement += '</style>';
		}
		if(winWidth > 1600) {
			var itemWidth = $('#wrap').width() * 0.19;
			var itemHeight = itemWidth / 1.7;
			var styleElement = '<style>.video-list-item{width:' + itemWidth + 'px !important;';
			styleElement += '}.video-list-item .video-img{height:' + itemHeight + 'px !important;}';
			styleElement += '#list_type{height:' + (itemHeight + 55) + 'px !important;overflow:hidden;}';
			styleElement += '</style>';
		}
		$('#pstyle').html(styleElement);
	}
	$(window).resize(function() {
		$('.lists-big-box').css('width', $(window).width() - 260);
		$('.images-show').each(function(index) {
			$('.lists-box-list-img').css('max-height', $(this).width() / 2);
		})
		Item_adaptation();
	});
})
