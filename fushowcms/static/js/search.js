$(function() {
	function getQueryString(key) {
		var reg = new RegExp("(^|&)" + key + "=([^&]*)(&|$)");
		var result = window.location.search.substr(1).match(reg);
		return result ? decodeURIComponent(result[2]) : null;
	}
	var local_url = getQueryString("kw");
	
	var pages = getStorage("cxzbpage");
	var pages = getStorage("zbpage");
	if(pages == null) {
		pages = 1;
	}
	var page;
	$("body").css("overflow-x", "hidden");
	$(function() {
		pageinit();
		$('.lists-big-box').css('width', $(window).width() - 260);
		Item_adaptation();
		$(function() {
			reqAjax("/page/selroomalias",{roomAlias:local_url,page: pages,rows: 24},function(msg) {
				console.log(msg);
				if(msg.ErrorCode!=0){
//					var list_type = document.getElementById('list_type');
//					list_type.innerHTML = "<div style='width: 100%;text-align: center;color: gainsboro;font-size: 30px;'>没有搜到相关信息</div>";
					return;
				}
				total = msg.Data.total;
				$('.lists-box-title p span').html('房间（' + total + '）');
				if(total == 0) {
					var list_type = document.getElementById('list_type');
					list_type.innerHTML = "<div style='width: 100%;text-align: center;color: gainsboro;font-size: 30px;'>没有搜到相关信息</div>";
					return ;
				}
				page = Math.ceil(total / 1);
				var html = '';
				for(var i = 0; i < msg.Data.data.length; i++) {
					str = '<a href="/roomlive/' + msg.Data.data[i].Id + '"><li><div class="lists-box-list-img"><img class="images-show" src="' + msg.Data.data[i].LiveCover + '" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>' + msg.Data.data[i].RoomAlias + '</h1><p>' + msg.Data.data[i].NickName + '<span><img src="/static/images/ui/look-ioc.png">' + msg.Data.data[i].LiveNumber + '</span></p></div></li></a>'
					html += str;
				}
				if(total >= 24) {
					$("#tcdPageCode").createPage({
						pageCount: page,
						current: pages,
						backFn: function(p) {
							//console.log(p);
						}
					});
				}
				var list_type = document.getElementById('list_type');
				list_type.innerHTML += html;
				$('.images-show').each(function(index) {
					$('.lists-box-list-img').css('max-height', $(this).width() / 2);
				});
			});

		});

	})
	$("#page-list").on('click', 'li', function() {
		var pages = setStorage("zbpage", 1);
	});

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

$(function() {
	$("#page-list").on('click', 'li', function() {
		var pages = setStorage("cxzbpage", 1)
	})
})