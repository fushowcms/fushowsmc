$(function() {  

			var uid = getStorage("Id");
			var pages = getStorage("czpage");
			var page;
			if(pages==null){
				pages=1;
			} 
			
			reqAjax("/user/userpayrecord",{UID: uid,page:pages,rows:20},function(msg){
				if(msg.ErrorCode!=0) {
					// Dialog(msg.ErrorMsg,true,"确定",null,function() {
					// 	$('.dialog').remove();
					// },null);
				}else {
					
					if (msg != null) {
							total = msg.Data.total;
							page = Math.ceil(total / 20);
							if (msg != null) {
								$.each(msg.Data.rows, function(key,val){
								var state
								//State        //支付状态 0：失败 1：成功
								//RefundState   //退款状态 0：未退款 1：退款成功
										if (!val.State){
											state = "支付失败";
										}else{
											state = "支付成功";
										}
										var str = "<tr class='removediv'><td>"+val.TradeNo +"</td><td>"+val.Time +"</td><td>"+ val.Money +"</td><td>"+ state +"<td></tr>";
										$("#rechargerecord_list").append(str);
								});
	
							//分页
							if(total>=20){
								$(".page-component").createPage({
									id:uid,
							       pageCount:page,
							        current:pages,
							        backFn:function(p){
					           			 //console.log(p);
								}
						    });
							}
							}else{
								return;
							}
						}
				}
			},true);
		}); 