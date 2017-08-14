(function ($) {
	var ms = {
		init: function (obj, args) {
			return (function () {
				ms.fillHtml(obj, args);
				ms.bindEvent(obj, args);
			})();
		},
		fillHtml: function (obj, args) {
			return (function () {
				obj.empty();
				uid = args.id;
				if (args.current > 1) {
					obj.append('<a href="javascript:;" class="prevPage">上一页</a>');
				} else {
					obj.remove('.prevPage');
					obj.append('<a href="javascript:;" style="border:none !important" class="disabled">上一页</a>');
				}
				if (args.current != 1 && args.current >= 4 && args.pageCount != 4) {
					obj.append('<a href="javascript:;" class="tcdNumber">' + 1 + '</a>');
				}
				if (args.current - 2 > 2 && args.current <= args.pageCount && args.pageCount > 5) {
					obj.append('<b>...</b>');
				}
				var start = args.current - 2, end = parseInt(args.current) + 2;
				if ((start > 1 && args.current < 4) || args.current == 1) {
					end++;
				}
				if (args.current > args.pageCount - 4 && args.current >= args.pageCount) {
					start--;
				}
				for (; start <= end; start++) {
					if (start <= args.pageCount && start >= 1) {
						if (start != args.current) {
							obj.append('<a href="javascript:;" class="tcdNumber">' + start + '</a>');
						} else {
							obj.append('<a href="javascript:;" class="current">' + start + '</a>');
						}
					}
				}
				if (parseInt(args.current) + 2 < parseInt(args.pageCount) - 1 && parseInt(args.current) >= 1 && parseInt(args.pageCount) > 5) {
					obj.append('<b>...</b>');
				}
				if (args.current != args.pageCount && args.current < args.pageCount - 2 && args.pageCount != 4) {
					obj.append('<a href="javascript:;" class="tcdNumber">' + args.pageCount + '</a>');
				}
				if (args.current < args.pageCount) {
					obj.append('<a href="javascript:;" class="nextPage">下一页</a>');
				} else {
					obj.remove('.nextPage');
					obj.append('<a href="javascript:;" style="border:none !important" class="disabled">下一页</a>');
				}
			})();
		},
		bindEvent: function (obj, args) {
			return (function () {
				obj.on("click", "a.tcdNumber", function () {
					var current = parseInt($(this).text());
					ms.fillHtml(obj, { "current": current, "pageCount": args.pageCount });
					if (typeof (args.backFn) == "function") {
						args.backFn(current);
						var url = "/user/getgoodslist";

						reqAjax(url, { page: current, rows: 8 }, function (data) {

							setStorage("page", current);
							//							window.location.reload();
							$("#ul_list li img").parent().remove();
							$.each(data.rows, function (key, val) {
								var str = '<li calss="removediv" data-id="' + val.Id + '" id="' + val.Id + '"><img src="' + val.GoodsPic + '" class="img-responsive"/><div class="myOrderList-font"><div class="myOrderList-font-top"><p id="goodsName">' + val.GoodsName + '</p><span id="goodsAccount">库存：' + val.GoodsStock + '</span><div style="clear:both"></div></div><div class="myOrderList-font-bottom"><p id="goodsPirce">' + val.GoodsPirce + '石榴籽</p><span class="myOrderList-font-btn">兑换</span><div style="clear:both"></div></div></div></li>'
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
										var stringAlert = '<div class="product-content-AlertBg"><div class="product-close">X</div><div class="product-content"><div class="product-content-img"><img src="' + data.Data.GoodsPic + '"/></div><div class="product-content-font"><div class="product-content-fontName">' + data.Data.GoodsName + '</div><div class="product-content-fontMoney"><b>单价：</b>' + data.Data.GoodsPirce + '石榴籽</div><div class="product-content-fontMain"><b>产品描述：</b>' + data.Data.GoodsAccount + '</div><div class="product-content-fontBtn">立即购买</div></div><div style="clear:both;"></div></div></div>'
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

						});

					}
				});
				obj.on("click", "a.prevPage", function () {
					var current = parseInt(obj.children("a.current").text());
					ms.fillHtml(obj, { "current": current - 1, "pageCount": args.pageCount });
					if (typeof (args.backFn) == "function") {
						args.backFn(current - 1);
						var url = "/user/getgoodslist";
						reqAjax(url, { page: current - 1, rows: 8 }, function (data) {

							setStorage("page", current);
							//							window.location.reload();
							$("#ul_list li img").parent().remove();
							$.each(data.rows, function (key, val) {
								var str = '<li calss="removediv" data-id="' + val.Id + '" id="' + val.Id + '"><img src="' + val.GoodsPic + '" class="img-responsive"/><div class="myOrderList-font"><div class="myOrderList-font-top"><p id="goodsName">' + val.GoodsName + '</p><span id="goodsAccount">库存：' + val.GoodsStock + '</span><div style="clear:both"></div></div><div class="myOrderList-font-bottom"><p id="goodsPirce">' + val.GoodsPirce + '石榴籽</p><span class="myOrderList-font-btn">兑换</span><div style="clear:both"></div></div></div></li>'
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
										var stringAlert = '<div class="product-content-AlertBg"><div class="product-close">X</div><div class="product-content"><div class="product-content-img"><img src="' + data.Data.GoodsPic + '"/></div><div class="product-content-font"><div class="product-content-fontName">' + data.Data.GoodsName + '</div><div class="product-content-fontMoney"><b>单价：</b>' + data.Data.GoodsPirce + '石榴籽</div><div class="product-content-fontMain"><b>产品描述：</b>' + data.Data.GoodsAccount + '</div><div class="product-content-fontBtn">立即购买</div></div><div style="clear:both;"></div></div></div>'
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

						});
					}
				});
				obj.on("click", "a.nextPage", function () {
					var current = parseInt(obj.children("a.current").text());
					ms.fillHtml(obj, { "current": current + 1, "pageCount": args.pageCount });
					if (typeof (args.backFn) == "function") {
						args.backFn(current + 1);
						var url = "/user/getgoodslist";
						reqAjax(url, { page: current + 1, rows: 8 }, function (data) {

							setStorage("page", current);
							//							window.location.reload();
							$("#ul_list li img").parent().remove();
							$.each(data.rows, function (key, val) {
								var str = '<li calss="removediv" data-id="' + val.Id + '" id="' + val.Id + '"><img src="' + val.GoodsPic + '" class="img-responsive"/><div class="myOrderList-font"><div class="myOrderList-font-top"><p id="goodsName">' + val.GoodsName + '</p><span id="goodsAccount">库存：' + val.GoodsStock + '</span><div style="clear:both"></div></div><div class="myOrderList-font-bottom"><p id="goodsPirce">' + val.GoodsPirce + '石榴籽</p><span class="myOrderList-font-btn">兑换</span><div style="clear:both"></div></div></div></li>'
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
										var stringAlert = '<div class="product-content-AlertBg"><div class="product-close">X</div><div class="product-content"><div class="product-content-img"><img src="' + data.Data.GoodsPic + '"/></div><div class="product-content-font"><div class="product-content-fontName">' + data.Data.GoodsName + '</div><div class="product-content-fontMoney"><b>单价：</b>' + data.Data.GoodsPirce + '石榴籽</div><div class="product-content-fontMain"><b>产品描述：</b>' + data.Data.GoodsAccount + '</div><div class="product-content-fontBtn">立即购买</div></div><div style="clear:both;"></div></div></div>'
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

						});
					}
				});
			})();
		}
	}
	$.fn.createPage1 = function (options) {
		var args = $.extend({
			pageCount: options.pageCount,
			current: options.current,
			backFn: function () { }
		}, options);
		ms.init(this, args);
	}
})(jQuery);