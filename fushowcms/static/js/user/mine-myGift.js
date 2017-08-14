$(function () {
	var menu = $('#mine-myGift-top li');
	menu.click(function () {
		var index = menu.index(this);
		$('.tabs_content').eq(index).show().siblings().hide();
		$('#mine-myGift-top').show();

		$(this).css({
			color: "white",
			background: "deepskyblue"
		}).siblings().css({
			color: "gray",
			background: "white"
		})
	})
	//余额
	var uid = getStorage("Id");
	reqAjax("/page/getuser", { UID: uid }, function (msg) {
		if (msg.ErrorCode != 0) {
			Dialog(msg.ErrorMsg, true, "确定", null, function () {
				$('.dialog').remove();
			}, null);
		} else {
			var yue = msg.Data.Balance;
			$("#myBlance").text(yue);
		}
	}, true);
	//交易记录
	var pages = getStorage("zspage");
	if (pages == null) {
		pages = 1;
	}
	var page;

	reqAjax("/user/findByUserIdGiveSlz", { UID: uid, page: pages, rows: 10 }, function (msg) {

		if (msg.Data == null) {
			return;
		}
		total = msg.Data.total;
		page = Math.ceil(total / 10);

		console.log(msg);
		$.each(msg.Data.data, function (key, val) {
			if (uid == val.RecipientId) {
				val.Num = "+" + val.Num;
			}
			if (uid == val.BenefactorId) {
				val.Num = "-" + val.Num;
			}
			var str = "<tr class='removediv'><td>" + val.GiveDate + "</td><td>" + val.BenefactorId + "</td><td>" + val.BenefactorName + "</td><td>" + val.RecipientId + "</td><td>" + val.RecipientName + "</td><td>" + val.Num + "</td></tr>"
			$("#transaction_list").append(str);
		});
		if (total >= 10) {
			$(".page-component").createPage({
				id: uid,
				pageCount: page,
				current: pages,
				backFn: function (p) {
					console.log(p);
				}
			});
		}
	});

	$("#mine-myGift-Yes").click(function () {
		var uid = getStorage("Id");
		var yhid;
		SlzPassUp();

	});

	var menu = $('.my-gift-nav li');
	menu.each(function (index) {
		var menuindex = index;
		$(this).click(function () {
			$(this).addClass('my-gift-nav-click').siblings().removeClass('my-gift-nav-click');
			$('.mySetting-table-content').css('display', 'none');
			$('.mySetting-table-content').eq(menuindex).css('display', 'block');
		})
	})

});

function sLZ() {
	var uid = getStorage("Id");
	var num = $("#resaver").val();
	var num2 = $("#resaverNum").val();
	var tel = "^[0-9]*[1-9][0-9]*$"
	if (!num2.match(tel)) {
		alertShowYNznx("赠送", "请输入正整数!", null);
		return;
	} else if (!num.match(tel)) {
		alertShowYNznx("赠送", "请输入有效id!", null);
		return;
	}
	alertShowYNznx("赠送金币", "<p style='font-size:16px;line-height:32px;'>用户你好</p><p style='font-size:16px;line-height:32px;'>您将为ID：" + num + "的用户</p><p style='font-size:16px;line-height:32px;'>赠送<span style='color:#df0050;font-size:25px'>" + num2 + "</span>金币</p>", "赠送", "我再看看");
	$('.alert-yes').bind('click', function () {
		setTimeout(function () {
			reqAjax("/user/givenumber", { UID: uid, Number: num2, ToId: num }, function (msg) {
				if (msg.ErrorCode == 0) {
					alertShowYNznx("提示", "成功", null);
					$('.alert-OK,.alert-close').live("click", function () {
						window.location.reload();
					})
				} else {
					alertShowYNznx("提示", msg.ErrorMsg, null);
					$('.alert-OK').live("click", function () {
						window.location.reload();
					})
				}
			}, true);
		}, 1000)
	});
	$('.alert-no').bind("click", function () {
		return;
	})
}


//赠送金币籽密码验证
function SlzPassUp() {
	reqAjax("/user/slzPassUp", { UID: getStorage('Id'), password: $('#resaverPassword').val() }, function (msg) {
		console.log(msg);
		if (msg.ErrorCode == 0) {
			isbindingbank();
		} else {
			alertShowYNznx("提示", msg.ErrorMsg, null);
		}
	})
}
//赠送金币籽
function isbindingbank() {
	reqAjax("/user/isbindingbank", { UID: getStorage('Id') }, function (msg) {
		if (msg.ErrorCode != 0) {
			Dialog(msg.ErrorMsg, true, "确定", null, function () {
				$('.dialog').remove();
			}, null);
		} else {
			$.each(msg.Data.rows, function (key, val) {
				if (val.Phone == "") {
					alertShowYNznx("提示", "检测到您未绑定手机,请先绑定手机再进行操作", null);
					return;
				} else {
					sLZ();
				}
			});
		}
	}, true);
}