$(function(){
	var pageurl;
	var type;
	reqAjax('/page/getIndexCarPicLists',{page:1,rows:5},function(data){
		if(!data.total){
			return;
		}
		var listmod = ''
		for(var i=0; i<data.rows.length; i++){
			if(data.rows[i].CarouselType==1) {
				listmod += '<li data-url="' + data.rows[i].PicPath + '" data-type="' + data.rows[i].CarouselType + '" data-video="' + data.rows[i].VideoLivePage + '"><img src=" '+ data.rows[i].PicPath +' " height="95" width="165"/><div class="pics-Border2"></div></li>';
			}else {
				listmod += '<li data-url="' + data.rows[i].PicPath + '" data-type="' + data.rows[i].CarouselType + '" data-video="' + data.rows[i].VideoLivePage + '"><img src=" '+ data.rows[i].Litming +' " height="95" width="165"/><div class="pics-Border2"></div></li>';
			}
		}
		$('#lunbo-Btn').html(listmod).find('li').click(function(){
			$("#lunbo-Pic").html("");
			$(this).addClass('active').siblings().removeClass('active');
			type = $(this).attr('data-type');
			pageurl = $(this).attr('data-video');
			if(type==1){//直播视频
				var str = pageurl.split('?')[1];
				var roomId = getstrpramas(str,'roomId');
				var anchorId = getstrpramas(str,'anchorId');	

				reqAjax('/page/getinflow',{AnchorId:anchorId},function(data){
					var player = new prismplayer({
						id: "lunbo-Pic", // 容器id
						source: data.Data.errMsg,
						autoplay: true,    
//						width: "1168px",       
//						height: "657px",  
						isLive: true,
						skinLayout:false
					});
				});
				var btnmod = '<a href="/roomlive/'+roomId+'" target="_blank" class="golivebtn none">进入该直播间</a>';
				$('#lunbo-Pic').append(btnmod).hover(function(){
					$(this).find('.golivebtn').removeClass('none');
				},function(){
					$(this).find('.golivebtn').addClass('none');
				});

			}else{
				$('#lunbo-Pic').html('<img src="'+$(this).data('url')+'" width="100%" height="100%"/>');				
			}			
		});
		$('#lunbo-Pic').click(function() {
			if(type!=0) {
				return
			} else {
				if(pageurl.indexOf('http://')!=-1) {
					window.open(pageurl);
				}else {
					window.open("http://" + pageurl);
				}
			}		
		});
		$('#lunbo-Btn').find('li').eq(0).click();
	});
});

