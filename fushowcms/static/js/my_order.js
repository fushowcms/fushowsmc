var anchorId = getUrlpramas('anchorId');
$(function() {
	var pages = getStorage("gzpage");
	if(pages==null){
		pages=1;
	}
	var page;
	var id = getStorage("Id");
	
	reqAjax('/user/getmyattentionlist',{UID:id,page:pages,rows:9},function(msg){
			var html='';
			if(msg.ErrorCode!=0){
				return;
			}else if (msg.ErrorCode=="0"){
			total = msg.Data.total;
			page = Math.ceil(total / 9);
			$.each(msg.Data.state, function(key,e){
				var roomname = '';
				var roomstate = '';
				var roomstatic='';
				if(e.RoomType == 1){
				roomname = "竞猜直播间";	
				}else{
					roomname = "普通直播间";
				}
				if(e.LiveState == 0){
					roomstate = "none";
					roomstatic="未开播";
					
				}else{
					roomstate = "block";
					roomstatic="正在直播";
				}
				var str = '<li><a href="/roomlive/'+ e.Id +'"><div class="lists-box-list-img"><img class="images-show" src="'+e.LiveCover+'" onerror="this.src=\'/static/images/thumb.jpg\'"><div class="list-mask"></div><div class="list-mask-img"><img src="/static/images/ui/play-ioc.png"/></div><div class="myOrder-title">'+e.RoomAlias+'</div><div class="myOrder-static" style="display:'+roomstate+'"></div></div></a><div class="lists-box-list-font"><h1><img src="/static/images/ui/clock-ioc.png" class="clock-ioc">'+roomstatic+'</h1><p>'+e.NickName+'<span class="offfollow" data-rowuser="'+ e.User + '">取消关注</span></p></div></li>';
							
				html+=str;
			});
		}
		if(total>=9){
			$(".page-component").createPage({
				id:id,
				pageCount:page,
				current:pages,
				backFn:function(p){
			        console.log(p);
				}		
			});
		}
			var hot_title =document.getElementById('myOrder-lists-show');
			hot_title.innerHTML += html;
	});
	
//	关注
	var offfollow = $('.offfollow');
	offfollow.bind("click", function confirmAct() {
		alertShowYNznx("提示", "确定要取消关注吗？", null, "取消");
		var rowUser = $(this).data('rowuser');
		$('.alert-yes').bind("click",function () {
			RoomConcernDel(id,rowUser);
		});
		$('.alert-no').bind("click",function () {
			return;
		});
	})

	function RoomConcernDel(userId, anchorId){
		reqAjax('/user/cancelroomcon',{User:anchorId,UID:userId},function(data){
			if(data.ErrorCode == "0") {
					alertShowYNznx("提示", "成功取消关注", null);
					$('.alert-OK').bind("click",function () {
						window.location.reload();
					});
			} else {
				alertShowYNznx("提示", data.state, null);
					$('.alert-OK').bind("click",function () {
						window.location.reload();
					});
			}
		});
	}
}); 
