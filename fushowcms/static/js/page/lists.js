var pages = getStorage("zbpage");

if(pages==null){
	pages=1;
}
var page;
$("body").css("overflow-x","hidden");
$(function(){
	pageinit();
	
	Item_adaptation();
	
			page = Math.ceil(total / 15);				
			var html='';
			var html1='';					
			if(data_normal!=null&&data_normal.length >= 8){
				if(data_normal!=null) {
					for(var i=0;i<data_normal.length;i++){
						str ='<a href="/roomlive/'+data_normal[i].Id+'"><li><div class="lists-box-list-img"><img class="images-show" onerror="this.src=\'/static/images/ui/noPlayimg.png\'" src="'+data_normal[i].LiveCover+'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>'+data_normal[i].RoomAlias+'</h1><span class="lists-classFont">'+data_normal[i].TwoCategoryName+'</span><p>'+data_normal[i].NickName+'<span><img src="/static/images/ui/look-ioc.png">'+ data_normal[i].LiveNumber +'</span></p></div></li></a>'
						html+=str;
					}
				}
			}else if(data_normal!=null&&data_normal.length< 8){
				for(var i=0;i<data_normal.length;i++){
					str ='<a href="/roomlive/'+data_normal[i].Id+'"><li><div class="lists-box-list-img"><img class="images-show" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"  src="'+data_normal[i].LiveCover+'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>'+data_normal[i].RoomAlias+'</h1><span class="lists-classFont">'+data_normal[i].TwoCategoryName+'</span><p>'+data_normal[i].NickName+'<span><img src="/static/images/ui/look-ioc.png">'+ data_normal[i].LiveNumber +'</span></p></div></li></a>'
					html+=str;
				}
				if(data_normal_state!=null) {
					for(var i=0;i<data_normal_state.length;i++){
						str1 ='<a href="/roomlive/'+data_normal_state[i].Id+'"><li><div class="lists-box-list-img"><img class="images-show" onerror="this.src=\'/static/images/ui/noPlayimg.png\'" src="'+data_normal_state[i].LiveCover+'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>'+data_normal_state[i].RoomAlias+'</h1><span class="lists-classFont">'+data_normal_state[i].TwoCategoryName+'</span><p>'+data_normal_state[i].NickName+'<span><img src="/static/images/ui/look-ioc.png">'+data_normal_state[i].LiveNumber +'</span></p></div></li></a>'
						html+=str1;
					}
				}
			}else if (data_normal==null){
				if(data_normal_state!=null) {
					for(var i=0;i<data_normal_state.length;i++){
						str1 = '<a href="/roomlive/'+data_normal_state[i].Id+'"><li><div class="lists-box-list-img"><img class="images-show" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"  src="'+data_normal_state[i].LiveCover+'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>'+data_normal_state[i].RoomAlias+'</h1><span class="lists-classFont">'+data_normal_state[i].TwoCategoryName+'</span><p>'+data_normal_state[i].NickName+'<span><img src="/static/images/ui/look-ioc.png">'+ data_normal_state[i].LiveNumber +'</span></p></div></li></a>';
						html+=str1;
					}	
				}
			}
			if(data_stick!=null &&data_stick.length >=5){
				if(data_stick!=null) {
					for(var i=0;i<data_stick.length;i++){	
						str2 ='<a href="/roomlive/'+data_stick[i].Id+'"><li><div class="lists-box-list-img"><img class="images-show" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"  src="'+data_stick[i].LiveCover+'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>'+data_stick[i].RoomAlias+'</h1><span class="lists-classFont">'+data_stick[i].TwoCategoryName+'</span><p>'+data_stick[i].NickName+'<span><img src="/static/images/ui/look-ioc.png">'+ data_stick[i].LiveNumber +'</span></p></div></li></a>'
						html1+=str2;
					}
				}
						
			}else if(data_stick!=null &&data_stick.length <5){
				if(data_stick!=null) {
					for(var i=0;i<data_stick.length;i++){	
						str3 ='<a href="/roomlive/'+data_stick[i].Id+'"><li><div class="lists-box-list-img"><img  class="images-show" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"  src="'+data_stick[i].LiveCover+'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>'+data_stick[i].RoomAlias+'</h1><span class="lists-classFont">'+data_stick[i].TwoCategoryName+'</span><p>'+data_stick[i].NickName+'<span><img src="/static/images/ui/look-ioc.png">'+ data_stick[i].LiveNumber +'</span></p></div></li></a>'
						html1+=str3;
					}
				}
				if(data_stick_state!=null) {
					for(var i=0;i<data_stick_state.length;i++){	
						str2 ='<a href="/roomlive/'+data_stick_state[i].Id+'"><li><div class="lists-box-list-img"><img class="images-show" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"  src="'+data_stick_state[i].LiveCover+'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>'+data_stick_state[i].RoomAlias+'</h1><span class="lists-classFont">'+data_stick_state[i].TwoCategoryName+'</span><p>'+data_stick_state[i].NickName+'<span><img src="/static/images/ui/look-ioc.png">'+ data_stick_state[i].LiveNumber +'</span></p></div></li></a>'
						html1+=str2;
					}
				}
				
			}else if(data_stick==null){
				if(data_stick_state!= null) {
					for(var i=0;i<data_stick_state.length;i++){	
						str2 ='<a href="/roomlive/'+data_stick_state[i].Id+'"><li><div class="lists-box-list-img"><img class="images-show" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"  src="'+data_stick_state[i].LiveCover+'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>'+data_stick_state[i].RoomAlias+'</h1><span class="lists-classFont">'+data_stick_state[i].TwoCategoryName+'</span><p>'+data_stick_state[i].NickName+'<span><img src="/static/images/ui/look-ioc.png">'+data_stick_state[i].LiveNumber +'</span></p></div></li></a>'
						html1+=str2;
					}
				}
				
			}
					
			// add by liuhan
			//--------------------------------------------------------
			if(total>=1){
				$(".page-component").createPage({
					pageCount:page,
					current:pages,
					backFn:function(p){
				    console.log(p);
					}		
				});
			}
					//--------------------------------------------------------
			var hot_title = document.getElementById('hot_title');
			var list_type = document.getElementById('list_type');
			hot_title.innerHTML += html1;
			list_type.innerHTML += html;
			$('.images-show').each(function(index){
				$('.images-show').css('height',$(this).width()*9/16);
			})

})
$("#page-list").on('click','li',function(){
	
	var pages = setStorage("zbpage",1);
});
function Item_adaptation(){
	var winWidth = $(window).width();
	var boxWidth = $('#wrap').width();
			
	if(winWidth <= 1280){
		var itemWidth = $('#wrap').width()*0.315;
		var itemHeight = itemWidth/1.7;
		var styleElement = '<style>.video-list-item{width:' + itemWidth + 'px !important;';
		styleElement += '}.video-list-item .video-img{height:' + itemHeight + 'px !important;}';
		styleElement += '#list_type{height:' + (itemHeight+55) + 'px !important;overflow:hidden;}';
		styleElement += '</style>';
	}
			
	if(winWidth>1280&&winWidth<=1600){
		var itemWidth = $('#wrap').width()*0.235;
		var itemHeight = itemWidth/1.7;
		var styleElement = '<style>.video-list-item{width:' + itemWidth + 'px !important;';
		styleElement += '}.video-list-item .video-img{height:' + itemHeight + 'px !important;}';
		styleElement += '#list_type{height:' + (itemHeight+55) + 'px !important;overflow:hidden;}';
		styleElement += '</style>';
	}
	if(winWidth>1600){
		var itemWidth = $('#wrap').width()*0.19;
		var itemHeight = itemWidth/1.7;
		var styleElement = '<style>.video-list-item{width:' + itemWidth + 'px !important;';
		styleElement += '}.video-list-item .video-img{height:' + itemHeight + 'px !important;}';
		styleElement += '#list_type{height:' + (itemHeight+55) + 'px !important;overflow:hidden;}';
		styleElement += '</style>';
	}
	$('#pstyle').html(styleElement);
}
$('.lists-big-box').resize(function() { 
	$('.lists-big-box').css('width',$(window).width()-260);
	$('.images-show').each(function(index){
		$('.images-show').css('height',$(this).width()*9/16);
	})
	Item_adaptation();
});