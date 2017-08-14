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
					var url = "/user/findByRecipient"; 
					reqAjax(url,{UID:args.id,page:current,rows:10},function(msg) {
						localStorage.setItem("settingpage",current);
						$(".removedivlive").remove();
						$.each(msg.Data.data, function(key,val){
							 if(msg.Data.id == val.RecipientId) {
								var str = "<tr class='removedivlive'><td>" + val.GiveDate + "</td><td>" + val.BenefactorName + "</td><td>"+val.GiftName+"*" + val.GiftNum + "</td></tr>";
								$("#liveearnings").append(str);
							}
						});
					});
					
					
					
//					$.ajax({  
//						type: "post",  
//						url: url,  
//						dataType: "json",  
//						data: {
//								UID:args.id,
//								page:current,
//								rows:20
//						},  
//						success: function(msg){
//							localStorage.setItem("settingpage",current);
//							$(".removedivlive").remove();
//							$.each(msg.data, function(key,val){
//								 if(msg.id == val.RecipientId) {
//									var str = "<tr class='removedivlive'><td>" + val.GiveDate + "</td><td>" + val.BenefactorName + "</td><td>"+val.GiftName+"*" + val.GiftNum + "</td></tr>";
//									$("#liveearnings").append(str);
//								}
//							});
//						}
//						});											
					}
				});
				obj.on("click","a.prevPage",function(){
					var current = parseInt(obj.children("a.current").text());
					ms.fillHtml(obj,{"current":current-1,"pageCount":args.pageCount});
					if(typeof(args.backFn)=="function"){
						args.backFn(current-1);
					var url = "/user/findByRecipient";  
					
					reqAjax(url,{UID:args.id,page:current-1,rows:10},function(msg) {
						localStorage.setItem("settingpage",current-1);
						$(".removedivlive").remove();
						$.each(msg.Data.data, function(key,val){
							 if(msg.Data.id == val.RecipientId) {
								var str = "<tr class='removedivlive'><td>" + val.GiveDate + "</td><td>" + val.BenefactorName + "</td><td>"+val.GiftName+"*" + val.GiftNum + "</td></tr>";
								$("#liveearnings").append(str);
							}
						});
					});
					
//					$.ajax({  
//						type: "post",  
//						url: url,  
//						dataType: "json",  
//						data: {
//								UID:args.id,
//								page:current-1,
//								rows:20
//						},  
//						success: function(msg){
//							localStorage.setItem("settingpage",current-1);
//							$(".removedivlive").remove();
//							$.each(msg.data, function(key,val){
//								 if(msg.id == val.RecipientId) {
//									var str = "<tr class='removedivlive'><td>" + val.GiveDate + "</td><td>" + val.BenefactorName + "</td><td>"+val.GiftName+"*" + val.GiftNum + "</td></tr>";
//									$("#liveearnings").append(str);
//								}
//							});
//						}
//					});	
					}
				});
				obj.on("click","a.nextPage",function(){
					var current = parseInt(obj.children("a.current").text());
					ms.fillHtml(obj,{"current":current+1,"pageCount":args.pageCount});
					if(typeof(args.backFn)=="function"){
						args.backFn(current+1);
					var url = "/user/findByRecipient";  
					// 
					reqAjax(url,{UID:args.id,page:current+1,rows:10},function(msg) {
						localStorage.setItem("settingpage",current+1);
						$(".removedivlive").remove();
						$.each(msg.Data.data, function(key,val){
							 if(msg.Data.id == val.RecipientId) {
								var str = "<tr class='removedivlive'><td>" + val.GiveDate + "</td><td>" + val.BenefactorName + "</td><td>"+val.GiftName+"*" + val.GiftNum + "</td></tr>";
								$("#liveearnings").append(str);
							}
						});
					});
//					$.ajax({  
//						type: "post",  
//						url: url,  
//						dataType: "json",  
//						data: {
//								UID:args.id,
//								page:current+1,
//								rows:20
//						},  
//						success: function(msg){
//							localStorage.setItem("settingpage",current+1);
//							$(".removedivlive").remove();
//						$.each(msg.data, function(key,val){
							
//								 if(msg.id == val.RecipientId) {
//									var str = "<tr class='removedivlive'><td>" + val.GiveDate + "</td><td>" + val.BenefactorName + "</td><td>"+val.GiftName+"*" + val.GiftNum + "</td></tr>";
//									$("#liveearnings").append(str);
//								}
//							});
//							}
//						});	
					}
				});
			})();
		}
	}
	$.fn.createPagelive = function(options){
		
		var args = $.extend({
			id :options.id,
			pageCount : options.pageCount,
			current : options.current,
			backFn : function(){}
		},options);
		
		ms.init(this,args);
	}
})(jQuery);