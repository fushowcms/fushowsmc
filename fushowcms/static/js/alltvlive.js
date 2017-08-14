var pages = getStorage("zbpage");

if(pages==null){
	pages=1;
}
var page,total;
$("body").css("overflow-x","hidden");
$(function(){
	pageinit();
	Item_adaptation();		
			reqAjax("/page/getalldata",{page:pages,rows:15},function(msg) {
				console.log(msg)
				if(msg.ErrorCode == 0){
					total = msg.Data.total;
					page = Math.ceil(total / 15);
					if(msg.Data.data){
						$("#polybox").show();
						for(var i = 0;i<msg.Data.data.length;i++){
							var mod = '<a href="/outlive/'+ msg.Data.data[i].SId +'">'
								+'<li>'
									+'<div class="lists-box-list-img">'
										+'<img class="images-show" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"  src="'+msg.Data.data[i].Thumbimg+'"/>'
										+'<div class="list-mask"></div>'
										+'<div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div>'
									+'</div>'
									+'<div class="lists-box-list-font">'
										+'<h1>'+msg.Data.data[i].Title+'</h1>'
										// +'<span class="lists-classFont">'+data_stick_state[i].TwoCategoryName+'</span>'
										+'<p>'+msg.Data.data[i].NickName+'<span><img src="/static/images/ui/look-ioc.png">'+msg.Data.data[i].Number +'</span></p>'
									+'</div>'
								+'</li>'
							+'</a>';
							$('#alllive').append(mod);
						}
					}
				}
			});

			if(total>=1){
				$(".page-component").createPage({
					pageCount:page,
					current:pages,
					backFn:function(p){
				    	console.log(p);
					}		
				});
			}
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