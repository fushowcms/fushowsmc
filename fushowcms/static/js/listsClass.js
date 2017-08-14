var widtnNums,
	listHeight;
$(function(){
	//刚进页面加载的全部
	
	reqAjax("/page/getTwoCategoryList",{},function(msg){
		$('.removeLi').remove()
		for(var i=0;i<msg.Data.length;i++){
			var	str ="";
			str += '<a href="'+msg.Data[i].TwoCategoryAddress+'">'
			str += '<li class="removeLi"><div class="listsClass-list-img">'
			str += '<img src="/static/upload/category/'+msg.Data[i].TwoCategoryImage+'" /></div>'	
			str += '<div class="listsClass-list-title">'+msg.Data[i].TwoCategoryName+'</div></li>'	
			str += '</a>'
			$('.listsClass-list').append(str)
		}
		mediaWidth(webWidth,boxWidth);
	},true);
	//全部二级类目 addby liuhan
	$(".category_all").on("click",function(){
		reqAjax("/page/getTwoCategoryList",{},function(msg){
			$('.removeLi').remove()
			for(var i=0;i<msg.Data.length;i++){
				var	str ="";
				str += '<a href="'+msg.Data[i].TwoCategoryAddress+'">'
				str += '<li class="removeLi"><div class="listsClass-list-img">'
				str += '<img src="/static/upload/category/'+msg.Data[i].TwoCategoryImage+'" /></div>'	
				str += '<div class="listsClass-list-title">'+msg.Data[i].TwoCategoryName+'</div></li>'	
				str += '</a>'
				$('.listsClass-list').append(str)
			}
		mediaWidth(webWidth,boxWidth);
		},true);
	});

	//分类 addby liuhan
	reqAjax("/page/getCategoryList",{},function(msg){
		
		for(var i=0;i<msg.rows.length;i++){
			var	str ="";
			str += '<li data-id="'+msg.rows[i].Id+'">'+msg.rows[i].OneCategoryName+'</li>'
			$('.listsClass-Nav-lists').append(str)
		}
		//点击
		$(".listsClass-Nav-lists li").on("click",function(){
			
			$('.removeLi').remove()
			$('.listsClass-Nav-lists li').removeClass('listsClass-Nav-lists-after');
			$(this).addClass('listsClass-Nav-lists-after');
			var dataId=$(this).data("id");
			reqAjax("/page/getTwoCategoryByOneIds",{Id:dataId},function(msg){
				for(var i=0;i<msg.Data.length;i++){
					var	str ="";
					str += '<a href="'+msg.Data[i].TwoCategoryAddress+'">'
					str += '<li class="removeLi"><div class="listsClass-list-img">'
					str += '<img src="/static/upload/category/'+msg.Data[i].TwoCategoryImage+'" /></div>'	
					str += '<div class="listsClass-list-title">'+msg.Data[i].TwoCategoryName+'</div></li>'	
					str += '</a>'
					$('.listsClass-list').append(str)
				}
				mediaWidth(webWidth,boxWidth);
			},true);
		})
	},true);

	var boxWidth=$('.listsClass-list').width(),
		webWidth=$(window).width(),
		listWidth;
	mediaWidth(webWidth,boxWidth);
	
	$('.listsClass-Nav-lists li').each(function(index){
		var cIndex=index;
		$(this).click(function(){
			$('.listsClass-Nav-lists li').removeClass('listsClass-Nav-lists-after');
			$('.listsClass-Nav-lists li').eq(cIndex).addClass('listsClass-Nav-lists-after');
		})
	})
	
})
$(window).resize(function(){
	var boxWidth=$('.listsClass-list').width(),
		webWidth=$(window).width(),
		listWidth;
	mediaWidth(webWidth,boxWidth);
})
function mediaWidth(windowWidth,boxWidth){
	var windowWidth=windowWidth,
		boxWidth=boxWidth;
	if(windowWidth > 1600){	
		$('.listsClass-list li').css({'width':"182px"});
		listHeight=$('.listsClass-list li').width()*14/9;
		$('.listsClass-list-img').css({'height':"252px"});
	}else if(windowWidth > 1200 && windowWidth < 1600 ){
		widtnNums= Math.floor((boxWidth-120)/6);
		$('.listsClass-list li').css({'width':""+widtnNums+"px"});
		listHeight=$('.listsClass-list li').width()*14/9;
		$('.listsClass-list-img').css({'height':""+listHeight+"px"});
	}else if(windowWidth > 1000 && windowWidth < 1200){
		widtnNums= Math.floor((boxWidth-80)/4);
		$('.listsClass-list li').css({'width':""+widtnNums+"px"});
		listHeight=$('.listsClass-list li').width()*14/9;
		$('.listsClass-list-img').css({'height':""+listHeight+"px"});
	}
}