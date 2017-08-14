
		var PgoodsPirce;
		//提交按钮
		
		function CheckForm(){
			alertShowYNznx("兑换商品","是否确定购买此商品",null,"取消");
			$('.alert-yes').on('click',function(){
			var a = getStorage("Id");
			
			reqAjax("/user/isbindingbank",{UID: a},function(msg){
				if(msg.ErrorCode!=0) {
					Dialog(msg.ErrorMsg,true,"确定",null,function() {
						$('.dialog').remove();
					},null);
				}else {
					if (msg != null) {
						$.each(msg.Data.rows, function(key, val) {
							if(val.Phone == "") {
								alertShow(".product-btn","检测到您未绑定手机","请先前往个人中心绑定手机再进行操作");
								return;
							}else{
								shopping();
							}
						});
					}				
				}
			},true);
			})
			$('.alert-no').on("click",function(){
				return false;
			})
		}; 
		function shopping(){
			if(getStorage("Id")==null){
					alertShowYNznx("兑换商品", "请登录后购买!",null);
					return
				}
				if($("#receiver").val().length == 0){
					alertShowYNznx("兑换商品", "收件人不能为空!",null);
					return
				}
				
				if($("#tel").val().length == 0){
					alertShowYNznx("兑换商品", "联系方式不能为空!",null);
					return	
				}else{
					var tel = "^1(3[0-9]|4[57]|5[0-35-9]|7[01678]|8[0-9])\\d{8}$";
					if(!($("#tel").val()).match(tel)){
						alertShowYNznx("兑换商品", "手机号格式不正确!",null);
						return
					}
				}
				
				if($("#address").val().length == 0){
					alertShowYNznx("兑换商品", "收货地址不能为空!",null);
					return
				}
				if($("#num").val()<=0) {
					alertShowYNznx("兑换商品", "商品数量必须大于0!",null);
					return
				}
				if($("#num").val()>goodsStock) {
					alertShowYNznx("兑换商品", "库存不足!",null);
					return
				}
				var goodsTotal = PgoodsPirce * parseInt($("#num").val())
				
				reqAjax("/user/orderadd",{GoodsId:goodsId,OrderID:goodsId,GoodsName:goodsName,GoodsNum:$("#num").val(),GoodsPirce:PgoodsPirce,GoodsTotal:goodsTotal,
					Receiver:$("#receiver").val(),Address:$("#address").val(),Tel:$("#tel").val(),GoodsAccount:$("#goodsAccount").html(),GoodsPic:goodsPic,UID:getStorage("Id")
					},function(msg){
						if (msg.ErrorCode != 0){
							alertShowYNznx("兑换商品",msg.ErrorMsg,null)
						}else if(msg.ErrorCode == 0){
								alertShowYNznx("兑换商品",msg.Data,null);
								$('.alert-OK,.alert-close').bind("click",function(){
									window.location.href='my_order_list';
								});
								return;
							}	
					});
					
					
				
				
				
		}	
		/*var gid = localStorage.getItem("goodsDetailId");
		var goodsId;
		$.ajax({  
				type: "post",  
				url: "/user/getgoodsById",  
				dataType: "json",  
				data: {
					Id:gid
				},  
				success: function(data){
					goodsId = data.goods.Id
					PgoodsPirce = data.goods.GoodsPirce;
					$("#goodsPic").attr("src", data.goods.GoodsPic);
					$("#goodsDetail").attr("src", data.goods.GoodsDetail);
					$("#goodsNameAccount").html(data.goods.GoodsName+"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+data.goods.GoodsAccount);
					$("#goodsName").html(data.goods.GoodsName);
					$("#goodsAccount").html(data.goods.GoodsAccount);
					$("#goodsPirce").html(data.goods.GoodsPirce);
					$("#emStock").html(data.goods.GoodsStock);
					$("#goodsTotal").html(data.goods.GoodsPirce);
				}  
				});*/



			$("#img_add").click(function() {
				var num = $("#num").val();
				var j_EmStock = $("#emStock").html();
				J_EmStock = parseInt(j_EmStock);
				num = parseInt(num);
				num += 1;
				if(num < j_EmStock) {
					$("#num").val(num);
					$("#goodsTotal").html(parseInt($("#goodsPirce").html())*num);
					$("#img_jian").attr("src", "../static/images/goods/jian.png");
				}
				if(num == j_EmStock) {
					$("#num").val(num);
					$("#goodsTotal").html(parseInt($("#goodsPirce").html())*num);
					$("#img_add").attr("src", "../static/images/goods/add1.png");
				}
				if(num > j_EmStock) {
					$("#img_add").attr("src", "../static/images/goods/add1.png");
				}
			});
			$("#img_jian").click(function() {

				var num = $("#num").val();
				var j_EmStock = $("#emStock").html();
				J_EmStock = parseInt(j_EmStock);
				num = parseInt(num);
				num -= 1;
				if(num < j_EmStock) {
					if(num == 1) {
						$("#num").val(1);
						$("#goodsTotal").html(parseInt($("#goodsPirce").html())*num);
						$("#img_jian").attr("src", "../static/images/goods/jian1.png");
					} else {
						$("#num").val(num);
						$("#goodsTotal").html(parseInt($("#goodsPirce").html())*num);
						$("#img_add").attr("src", "../static/images/goods/add.png");
					}
				}
				if(num >= j_EmStock) {
					$("#num").val(num);
					$("#goodsTotal").html(parseInt($("#goodsPirce").html())*num);
					$("#img_add").attr("src", "../static/images/goods/add1.png");
				}
				if(num < 1) {
					$("#num").val(1);
					$("#goodsTotal").html(parseInt($("#goodsPirce").html())*1);
				}
			});

			$("#num").keyup(function() {
				//		    alert("keyup")
				var num = $("#num").val();
				var j_EmStock = $("#emStock").html();
				J_EmStock = parseInt(j_EmStock);
				num = parseInt(num);
				$("#goodsTotal").html(parseInt($("#goodsPirce").html())*num);
				if(num >= j_EmStock) {
					$("#img_add").attr("src", "../static/images/goods/add1.png");
					$("#img_jian").attr("src", "../static/images/goods/jian.png");
				}
				if(j_EmStock < 1) {
					$("#img_jian").attr("src", "../static/images/goods/jian1.png");
					$("#img_add").attr("src", "../static/images/goods/add.png");
				}
				if(num < j_EmStock) {
					$("#img_jian").attr("src", "../static/images/goods/jian.png");
					$("#img_add").attr("src", "../static/images/goods/add.png");
				}
				if(num < 1) {
					$("#img_jian").attr("src", "../static/images/goods/jian1.png");
					$("#img_add").attr("src", "../static/images/goods/add.png");
				}
			});