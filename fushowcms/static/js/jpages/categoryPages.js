(function($){
	var uid;
	var ms = {
		init:function(obj,args){
			return (function(){
				ms.fillHtml(obj,args);
				ms.bindEvent(obj,args);
			})();
		},
		fillHtml:function(obj,args){
			return (function(){
				obj.empty();
				uid = args.id;
				if(args.current > 1){
					obj.append('<a href="javascript:;" class="prevPage">上一页</a>');
				}else{
					obj.remove('.prevPage');
					obj.append('<a href="javascript:;" style="border:none !important" class="disabled">上一页</a>');
				}
				if(args.current != 1 && args.current >= 4 && args.pageCount != 4){
					obj.append('<a href="javascript:;" class="tcdNumber">'+1+'</a>');
				}
				if(args.current-2 > 2 && args.current <= args.pageCount && args.pageCount > 5){
					obj.append('<b>...</b>');
				}
				var start = args.current -2,end = parseInt(args.current)+2;
				if((start > 1 && args.current < 4)||args.current == 1){
					end++;
				}
				if(args.current > args.pageCount-4 && args.current >= args.pageCount){
					start--;
				}
				for (;start <= end; start++) {
					if(start <= args.pageCount && start >= 1){
						if(start != args.current){
							obj.append('<a href="javascript:;" class="tcdNumber">'+ start +'</a>');
						}else{
							obj.append('<a href="javascript:;" class="current">'+ start +'</a>');
						}
					}
				}
				if( parseInt(args.current) + 2 <  parseInt(args.pageCount) - 1 &&  parseInt(args.current) >= 1 && parseInt(args.pageCount) > 5){
					obj.append('<b>...</b>');
				}
				if(args.current != args.pageCount && args.current < args.pageCount -2  && args.pageCount != 4){
					obj.append('<a href="javascript:;" class="tcdNumber">'+args.pageCount+'</a>');
				}
				if(args.current < args.pageCount){
					obj.append('<a href="javascript:;" class="nextPage">下一页</a>');
				}else{
					obj.remove('.nextPage');
					obj.append('<a href="javascript:;" style="border:none !important" class="disabled">下一页</a>');
				}
			})();
		},
		bindEvent:function(obj,args){
			return (function(){
				obj.on("click","a.tcdNumber",function(){
					var current = parseInt($(this).text());
					ms.fillHtml(obj,{"current":current,"pageCount":args.pageCount});
					if(typeof(args.backFn)=="function"){
						args.backFn(current);
						
						reqAjax("/page/getRoomAliasByRoomType",{RoomType:args.RoomType,page:current,rows:15},function(msg){
							if(msg.ErrorCode==0){
								setStorage("categorypage",current);
								$(".removeabc").remove();
								if (!msg.Data.data){
									var list_type = document.getElementById('list_type');
									list_type.innerHTML = "<div style='width: 100%;text-align: center;color: gainsboro;font-size: 50px;'>没有搜到相关信息</div>";
									return;
								}
								var html = '';
								for(var i = 0; i < msg.Data.data.length; i++) {
									str = '<a href="' + msg.Data.data[i].LiveAddress + '"class="removeabc"><li><div class="lists-box-list-img"><img class="images-show" src="' + msg.Data.data[i].LiveCover + '" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>' + msg.Data.data[i].RoomAlias + '</h1><p>' + msg.Data.data[i].NickName + '<span><img src="/static/images/ui/look-ioc.png">' + msg.Data.data[i].LiveNumber + '</span></p></div></li></a>'
									html += str;
								}
								var list_type = document.getElementById('list_type');
								list_type.innerHTML += html;
								$('.images-show').each(function(index) {
									$('.lists-box-list-img').css('max-height', $(this).width() / 2);
								});
							}
						},true);										
					}
				});
				obj.on("click","a.prevPage",function(){
					var current = parseInt(obj.children("a.current").text());
					ms.fillHtml(obj,{"current":current-1,"pageCount":args.pageCount});
					if(typeof(args.backFn)=="function"){
						args.backFn(current-1);
						
						reqAjax("/page/getRoomAliasByRoomType",{RoomType:args.RoomType,page:current-1,rows:15},function(msg){
							if(msg.ErrorCode==0){
								$(".removeabc").remove();
								setStorage("categorypage",current-1);
								if (!msg.Data.data){
									var list_type = document.getElementById('list_type');
									list_type.innerHTML = "<div style='width: 100%;text-align: center;color: gainsboro;font-size: 50px;'>没有搜到相关信息</div>";
									return;
								}
								var html = '';
								for(var i = 0; i < msg.Data.data.length; i++) {
									str = '<a href="' + msg.Data.data[i].LiveAddress + '"class="removeabc"><li><div class="lists-box-list-img"><img class="images-show" src="' + msg.Data.data[i].LiveCover + '" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>' + msg.Data.data[i].RoomAlias + '</h1><p>' + msg.Data.data[i].NickName + '<span><img src="/static/images/ui/look-ioc.png">' + msg.Data.data[i].LiveNumber + '</span></p></div></li></a>'
									html += str;
								}
								var list_type = document.getElementById('list_type');
								list_type.innerHTML += html;
								$('.images-show').each(function(index) {
									$('.lists-box-list-img').css('max-height', $(this).width() / 2);
								});
							}
						},true);			
					}
				});
				obj.on("click","a.nextPage",function(){
					var current = parseInt(obj.children("a.current").text());
					ms.fillHtml(obj,{"current":current+1,"pageCount":args.pageCount});
					if(typeof(args.backFn)=="function"){
						args.backFn(current+1);
						
						reqAjax("/page/getRoomAliasByRoomType",{RoomType:args.RoomType,page:current+1,rows:15},function(msg){
							if(msg.ErrorCode==0){
							$(".removeabc").remove();
							setStorage("categorypage",current+1);
								if (!msg.Data.data){
									var list_type = document.getElementById('list_type');
									list_type.innerHTML = "<div style='width: 100%;text-align: center;color: gainsboro;font-size: 50px;'>没有搜到相关信息</div>";
									return;
								}
								var html = '';
								for(var i = 0; i < msg.Data.data.length; i++) {
									str = '<a href="' + msg.Data.data[i].LiveAddress + '"class="removeabc"><li><div class="lists-box-list-img"><img class="images-show" src="' + msg.Data.data[i].LiveCover + '" onerror="this.src=\'/static/images/ui/noPlayimg.png\'"/><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div></div><div class="lists-box-list-font"><h1>' + msg.Data.data[i].RoomAlias + '</h1><p>' + msg.Data.data[i].NickName + '<span><img src="/static/images/ui/look-ioc.png">' + msg.Data.data[i].LiveNumber + '</span></p></div></li></a>'
									html += str;
								}
								var list_type = document.getElementById('list_type');
								list_type.innerHTML += html;
								$('.images-show').each(function(index) {
									$('.lists-box-list-img').css('max-height', $(this).width() / 2);
								});
							}
						},true);		
					}
				});
			})();
		}
	}
	$.fn.createPage = function(options){
		
		var args = $.extend({
			id :options.RoomType,
			pageCount : options.pageCount,
			current : options.current,
			backFn : function(){}
		},options);
		ms.init(this,args);
	}
})(jQuery);