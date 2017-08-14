var goodsStock;
var goodsName;
var goodsPic;

//商品列表
$(function () {

	var pages = getStorage("page");
	var page;
	if (pages == null) {
		pages = 1;
	}
	var total;

	reqAjax("/user/getgoodslist", { page: pages, rows: 8 }, function (data) {
		if (data.rows == null) {
			return false;
		}
		total = data.total;
		page = Math.ceil(total / 8);

		$.each(data.rows, function (key, val) {
			var str = '<li calss="removediv" data-id="' + val.Id + '" id="' + val.Id + '"><img src="' + val.GoodsPic + '" class="img-responsive"/><div class="myOrderList-font"><div class="myOrderList-font-top"><p id="goodsName">' + val.GoodsName + '</p><span id="goodsAccount">库存：' + val.GoodsStock + '</span><div style="clear:both"></div></div><div class="myOrderList-font-bottom"><p id="goodsPirce">' + val.GoodsPirce + '金币</p><span class="myOrderList-font-btn">兑换</span><div style="clear:both"></div></div></div></li>'

			$("#ul_list").append(str);


		});
		//然后点击
		$("#ul_list li").on("click", function () {

			var dataIDN = $(this).data("id");
			reqAjax("/user/getgoodsById", { Id: dataIDN }, function (data) {
				if (data.ErrorMsg == "ok") {
					goodsId = data.Data.Id
					PgoodsPirce = data.Data.GoodsPirce;
					$("#goodsPic").attr("src", data.Data.GoodsPic);
					$("#goodsDetail").attr("src", data.Data.GoodsDetail);
					$("#emStock").html(data.Data.GoodsStock);
					goodsStock = data.Data.GoodsStock;
					goodsName = data.Data.GoodsName;
					goodsPic = data.Data.GoodsPic;
					var stringAlert = '<div class="product-content-AlertBg"><div class="product-close"></div><div class="product-content"><div class="product-content-img"><img src="' + data.Data.GoodsPic + '"/></div><div class="product-content-font"><div class="product-content-fontName">' + data.Data.GoodsName + '</div><div class="product-content-fontMoney"><b style="font-size:18px;color:#5a5959;font-weight:100;font-size: 16px;">单价：</b>' + data.Data.GoodsPirce + '金币</div><div class="product-content-fontMain"><b style="color:#5a5959;float:left;font-size:18px;font-weight: 100;font-size: 16px; margin-top: 11px;">产品描述：</b><p class="product-p">' + data.Data.GoodsAccount + '</p></div><div class="product-content-fontBtn">立即购买</div></div><div style="clear:both;"></div></div></div>'
					$('#product-show').append(stringAlert);

					$('.product-close').on("click", function () {
						$('.product-content-AlertBg').remove();
					})
					$('.product-content-fontBtn').on("click", function () {
						$('.product-form-bg').css('display', 'block');
					})
					$('.product-btn-no').on("click", function () {
						$('.product-content-AlertBg').remove();
						$('.product-form-bg').css('display', 'none');
					})
				}
			});
		})


		//分页

		if (total >= 8) {
			$(".page-component1").createPage1({
				pageCount: page,
				current: pages,
				backFn: function (p) {
					//console.log(p);
				}
			});
		}

	});
	function onclick_click(e) { // 在页面任意位置点击而触发此事件
		$("a").removeClass("selected");
		$('#a' + e).addClass("selected");
		setStorage("page", e);
		window.location.reload();
	}
});
//兑换记录
$(function () {
	var pages = getStorage("orderpage");
	var page;
	if (pages == null) {
		pages = 1;
	}
	var userId = getStorage("Id");
	reqAjax("/user/getorderbyuserid", { page: pages, rows: 8, UID: userId }, function (data) {

		if (data.ErrorCode == "0") {
			total = data.Data.total;
			page = Math.ceil(total / 8);
			$.each(data.Data.order, function (key, val) {
				var state;
				var remove = " ";
				var remove1;
				var fontColor;
				var delfontColor;
				var delstate;
				var delfont;
				if (val.Setstate == 0) {
					state = "未发货"
					fontColor = "style='color: red;'"
				}
				if (val.Setstate == 0 && val.Delstate == 1) {
					state = "未发货"
				}
				if (val.Setstate == 1) {
					state = "已发货"
					fontColor = "style='color:#1eab24;'"
				}
				if (val.Delstate == 1) {
					delfont = "已取消"
					delfontColor = "style='color: red;'"
				}
				if (val.Delstate == 0) {
					delfont = "未取消"
					delfontColor = "style='color:#1eab24;'"
				}
				var str = "<tr class='removediv' data-id='"
					+ val.Id
					+ "' id ='"
					+ val.Id
					+ "'><td><div style=' height: 50px;width: 50px;overflow: hidden; margin:0 auto; padding:5px;'><img src='"
					+ val.GoodsPic
					+ "' style=' height: 50px;width:50px;'/><div></td><td><span>"
					+ val.GoodsName
					+ "</span></td><td><span class=''>"
					+ val.GoodsNum
					+ "</span></td><td><span class=''>"
					+ val.GoodsTotal
					+ "</span></td>"
					+ "<td><span>"
					+ val.CreateTime
					+ "</span></td>"
					+ "<td id='setstate'><span class=''"
					+ fontColor
					+ ">"
					+ state
					+ "</span></td>"
					+ "<td data-id='"
					+ val.Id
					+ "' >"
					+ "</td></tr>";

				$("#order_list").append(str);
			});
			if (total >= 8) {
				$(".page-component").createPage({
					id: userId,
					pageCount: page,
					current: pages,
					backFn: function (p) {
						//console.log(p);
					}
				});
			}
		}
	});

	function onclick_remove(r) {
		if (confirm("确认删除么！此操作不可恢复")) {

			var id = r.dataset.id;
			$.ajax({
				type: "post",
				url: "/user/orderdel",
				dataType: "json",
				data: {
					Id: id
				},
				success: function (data) {

					window.location.reload();//刷新当前页面.
				}
			});
		}
	}
	function onclick_quxiao(q) {
		if (confirm("确认取消订单么！此操作不可恢复")) {

			var id = q.dataset.id;
			$.ajax({
				type: "post",
				url: "/user/chargeback",
				dataType: "json",
				data: {
					Id: id
				},
				success: function (data) {

					window.location.reload();//刷新当前页面.
				}
			});
		}
	}
});
$(function () {
	var menu = $('.my-gift-nav li');
	menu.each(function (index) {
		var menuindex = index;
		$(this).click(function () {
			$(this).addClass('my-gift-nav-click').siblings().removeClass('my-gift-nav-click');
			$('.mySetting-table-content').css('display', 'none');
			$('.mySetting-table-content').eq(menuindex).css('display', 'block');
		})
	})
})