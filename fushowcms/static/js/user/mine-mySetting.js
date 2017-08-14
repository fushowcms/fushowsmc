$(function () {
	var roomId;
	var uid = getStorage("Id");
	var bankTF;
	var roomIds;
	var chatconf = {
		//host: "192.168.1.200",
		host: "114.55.134.4",
		port: "61623",
		clientId: "example" + (Math.floor(Math.random() * 100000)),
		user: "admin",
		password: "www.fushow.cn"
	};

	reqAjax("/page/getTwoCategoryList", {}, function (result) {
		if (result.Data == null) {
			return;
		} else {
			$.each(result.Data, function (key, val) {
				$("#roomType").append(" <option value =" + val.Id + ">" + val.TwoCategoryName + "</option>");
			});
		}
	});

	reqAjax("/user/getroominfo", { UID: uid }, function (msg) {
		if (msg.ErrorCode != 0) {
			return;
		} else {
			var urlstr = window.location.host;
			roomId = msg.Data.Id;
			$("#mysetting_roomid").text(msg.Data.Id);
			$("#biaoti-xiugai").attr("placeholder", msg.Data.RoomAlias);
			$("#inputM").text(msg.Data.RoomNotice);
			$("#setting-zhiboBtn").wrap("<a href='/roomlive/" + msg.Data.Id + "' ></a>");
			$("#wangzhiaddr").val("http://" + urlstr + "/roomlive/" + msg.Data.Id);
			$("#setting-zhiboBtn").attr('ddd', msg.Data.Id);
			$("#inputM").text(msg.Data.RoomNotice);
			$("#live_cover_img").css("background-image", "url(." + msg.Data.LiveCover + ")");
			$("#roomType").val(msg.Data.RoomType);
			//直播流
			plugflow();
		}
	}, true);

	// 复制功能
	var clipboard = new Clipboard('.my_clip_button');
	clipboard.on('success', function(e) {
		alertShowYNznx("复制成功", "复制成功", null);
		e.clearSelection();
	});

	//直播记录
	$('#mine-setting-ul li').eq(2).click(function () {
		var pages = getStorage("zbjlpage");
		var page;
		applycashingnumtext = '', cashingtext = '';
		if (pages == null) {
			pages = 1;
		}
		var userId = getStorage("Id");


		reqAjax("/user/getanchortime", { UID: uid, page: pages, rows: 20 }, function (msg) {

			total = msg.Data.total;
			page = Math.ceil(total / 20);
			$('.lll').remove();
			if (msg.ErrorCode != 0) {
				return;
			}
			if (msg.Data.total > 0) {
				$.each(msg.Data.data, function (key, val) {
					var str = "<tr class = 'lll removediv'><td>" + val.StartTime + " -- " + val.EndTime + "</td><td>" + val.RoomId + "</td><td>" + val.AnchormTime + "</td></tr>";
					$("#liverecord").append(str);
				});
			}
			if (total >= 20) {
				$(".page-component").createPage({
					id: userId,
					pageCount: page,
					current: pages,
					backFn: function (p) {
						//console.log(p);
					}
				});
			}
		});


	})


	//显示人民币金额
	$('#applycashingnumtext').bind('input propertychange', function () {
		$('#convertedinto').html(($(this).val()) * 80 / 10000);
	});

	//提交结算
	$('#applycashingnumbtn').click(function () {
		var date = new Date();
		var strDate;
		reqAjax("/user/getNowDate", {}, function (msg) {
			strDate = parseInt(msg.Data.split("-")[2]);
		}, false);
		var num = "^[0-9]*[1-9][0-9]*$"
		var num1 = $('#applycashingnumtext').val() / 100 + ""

		if (!num1.match(num)) {
			alertShowYNznx("赠送", "积分兑换只能为100的倍数!", null);
			return;
		}
		if (strDate >= 29) {
			alertShowYNznx("提示", "请在每月28号(包括28)之前申请结算！", null);
			return;
		} else {
			reqAjax("/user/isbindingbank", { UID: uid }, function (msg) {

				if (msg.ErrorCode != 0) {
					Dialog(msg.ErrorMsg, true, "确定", null, function () {
						$('.dialog').remove();
					}, null);
				} else {
					$.each(msg.Data.rows, function (key, val) {
						if (val.Phone == "") {
							alert("检测到您未绑定手机,请先绑定手机再进行操作");
							return;
						} else {
							iscashmonth();
						}
					});
				}
			}, true);
		}
	});
	//直播秘钥重置
	$(".plugag").click(function () {
		plugflows();
	});

	//绑定银行卡
	$('#bangbankbtn').click(function () {
		var bankcardNum = $('#bankcardNum').val();
		var bankName = $('#bankName').val();
		var bankdePosit = $('#bankdePosit').val();
		if (bankName != "" && bankcardNum != "" && bankdePosit != "") {
			$.ajax({
				type: "post",
				url: "/user/anchorbindingbank",
				dataType: "json",
				data: {
					UID: uid,
					bankcard: bankcardNum,
					bankname: bankName,
					bankdeposit: bankdePosit
				},
				async: true,
				success: function (msg) {
					if (msg.ErrorCode == 0) {
						alertShowYNznx("提示", "成功绑定", null);
						$('#bangbank').css('display', 'none');
						$('#bankcardmsg').css('display', 'block');
						$('.bankName').html(bankName);
						$('.bankcardNum').html(bankcardNum.substring(0, 3) + "****" + bankcardNum.substring(14, 18));
						$('.bankdePosit').html(bankdePosit);
					} else {
						return;
					}
				},
				error: function () {
					alert("失败!");
				}
			});
		}
		if (bankcardNum == "") {
			$('#bankcardNum').css('background-color', 'coral');
			$('#bankcardNumshow').show();
		}
		if (bankName == "") {
			$('#bankName').css('background-color', 'coral');
			$('#bankNameshow').show();
		}
		if (bankdePosit == "") {
			$('#bankdePosit').css('background-color', 'coral');
			$('#bankdePositshow').show();
		}
	});

//	//是否绑定银行卡
//	reqAjax("/user/isbindingbank", { UID: uid }, function (msg) {
//		if (msg.ErrorCode != 0) {
//			Dialog(msg.ErrorMsg, true, "确定", null, function () {
//				$('.dialog').remove();
//			}, null);
//		} else {
//			if (msg == null) {
//				return;
//			}
//			$.each(msg.Data.rows, function (key, val) {
//				$('#pomegranateNum').html(val.PomegranateNum);
//				if (val.IsBandingBank) {
//					$('#bangbank').css('display', 'none');
//					$('#bankcardmsg').css('display', 'block');
//					$('#zsname').attr('value', val.RealName);
//					$('.zsname').html(val.RealName);
//					$('.bankName').html(val.BankName);
//					var BankCardtext = val.BankCard.toString().substring(0, 4) + "**** ****" + val.BankCard.toString().substring(14, 18);
//					$('.bankcardNum').html(BankCardtext);
//					$('.bankdePosit').html(val.BankDeposit);
//					bankTF = true;
//				} else {
//					$('#zsname').attr('value', val.RealName);
//					$('.zsname').html(val.RealName);
//					bankTF = false;
//				}
//			});
//		}
//	}, true);

	var pages = getStorage("sydhpage");
	var page;
	if (pages == null) {
		pages = 1;
	}

	reqAjax("/user/settlementdetail", { UID: uid, page: pages, rows: 20 }, function (msg) {
		if (msg.ErrorCode != 0) {
			// Dialog(msg.ErrorMsg,true,"确定",null,function() {
			// 	$('.dialog').remove();
			// },null);
		} else {
			$('.xxx').remove();
			if (msg == null) {
				return;
			} else {
				total = msg.Data.total;
				page = Math.ceil(total / 20);
			}
			$.each(msg.Data.rows, function (key, val) {
				if (val.IsCashing) {
					val.IsCashing = "已结算";
				} else {
					val.IsCashing = "未结算";
				}
				var str = "<tr class = 'xxx removediv'><td>" + val.CashingDate + "</td><td>" + val.ApplyCashingNum + "</td><td>" + val.Cashing + "</td><td>" + val.IsCashing + "</td></tr>";
				$("#returnexchange").append(str);
			});
			//分页

			if (total >= 20) {
				$(".page-component1").createPage1({
					id: uid,
					pageCount: page,
					current: pages,
					backFn: function (p) {
						//console.log(p);
					}
				});
			}
		}
	}, true);
	var name;
	//js截取图片函数
	$("#select_img").live("change", function () {
		var uploadFile = $(this).val();
		var file = uploadFile.lastIndexOf("\\");
		name = uploadFile.substring(file + 1);
		var photoExt = name.substring((name.lastIndexOf(".")) + 1).toLowerCase();
		if (!(photoExt == "jpg" || photoExt == "jpeg" || photoExt == "png" || photoExt == "gif")) {
			alert("请选择正确的图片格式,jpg/jpeg/png/gif");
			return;
		}
		$.ajaxFileUpload({
			type: "post",
			url: "/uploadj",
			secureuri: true,
			fileElementId: "select_img",
			dataType: "json",
			success: function (msg) {
				if (msg.state == "success") {
					$("#edit_img").fadeIn(1000);
					$("#overlay").css("display", "block");
					$("#preview").attr("src", "../../static/upload/" + name + "");
					$("#target").attr("src", "../../static/upload/" + name + "");
					$("#file_name").val(name);
					$("#file_ext").val(photoExt);
					new cutImage().init();
				} else {
					alert(msg.state);
				}
			},
			error: function (msg) {
				alert("网络环境异常，请稍好再试");
			}
		});
	});
	$("#cancel").click(function () {
		$.ajax({
			type: "post",
			url: "/user/roomcoverdel",
			dataType: "json",
			async: false,
			data: {
				FileName: $("#file_name").val(),
			},
			success: function (msg) {
				if (msg.state == "success") {
					$("#edit_img").css("display", "none");
					$("#overlay").css("display", "none");
					window.location.reload(); //刷新当前页面
				} else {
					alert(msg.state)
				}
			},
			error: function (msg) {
				alert("系统错误");
				$("#edit_img").css("display", "none");
				$("#overlay").css("display", "none");

			}
		});

	});
	$("#submit_img").click(function () {
		var uid = getStorage("Id");
		//确认剪切
		var img = new Image();
		img.src = $('#preview').attr("src");
		var w = img.width;
		var ratio = img.width / 390;
		var X1 = parseInt($("#x1").val() * ratio);
		var Y1 = parseInt($("#y1").val() * ratio);
		var Cw = parseInt($("#cw").val() * ratio);
		var Ch = parseInt($("#ch").val() * ratio);
		$.ajax({
			type: "post",
			url: "/user/roomcoverup",
			dataType: "json",
			async: false,
			data: {
				UID: uid,
				FileName: $("#file_name").val(),
				FileExt: $("#file_ext").val(),
				X1: X1,
				Y1: Y1,
				Cw: Cw,
				Ch: Ch,
			},
			success: function (msg) {
				if (msg.state == "success") {
					$("#edit_img").css("display", "none");
					$("#overlay").css("display", "none");
					$("#live_cover_img").css("background-image", "url(../../static/upload/J" + name + ")");
					alert("修改成功");
				} else {
					alert(msg.state);
					$("#edit_img").css("display", "none");
					$("#overlay").css("display", "none");
				}
				window.location.reload(); //刷新当前页面
			},
			error: function (msg) {
				alert("系统错误");
				$("#edit_img").css("display", "none");
				$("#overlay").css("display", "none");

			}
		});
	});
	$(".room-xiugai").click(function () {
		var id = $("#mysetting_roomid").text();
		var name = $("#biaoti-xiugai").val();
		var gonggao = $("#inputM").val();
		var roomType = $("#roomType").val();
		if (name == "" && gonggao == "") {
			alertShowYNznx("提示", "公告不能为空", null);
			return;
		}
		reqAjax("/user/roomupis", { RoomId: id, RoomAlias: name, RoomNotice: gonggao, RoomType: roomType }, function (msg) {
			if (msg.ErrorCode != 0) {
				alertShowYNznx("提示", msg.ErrorMsg, null);

				$('.alert-OK').live("click", function () {
					return;
				});
				return;
			} else {
				alertShowYNznx("提示", "修改成功", null);
				$('.alert-OK').live("click", function () {
					return;
				});
			}
		});
	})
	//直播管理 table
	var menu = $('#mine-setting-ul li');
	menu.each(function (index) {
		var menuindex = index;
		$(this).click(function () {
			$('.mySetting-table-content').eq(menuindex).css('display', 'block').siblings().css('display', 'none');
			menu.eq(menuindex).addClass('mySetting-navClick').siblings().removeClass('mySetting-navClick');
		})
	})

	$('.exchange-nav li').each(function (index) {
		var exchangeIndex = index;
		$(this).click(function () {

			$('.exchange-nav li').eq(exchangeIndex).addClass('exchange-nav-click').siblings().removeClass('exchange-nav-click');
			$('.exchange-table').eq(exchangeIndex).css('display', 'block').siblings('.exchange-table').css('display', 'none');;
		})
	})
	//房管管理列表显示
	$('#mine-setting-ul li').eq(5).click(function () {
		var pages = getStorage("glypage");
		var page;
		if (pages == null) {
			pages = 1;
		}
		var total;
		reqAjax("/user/findByRoomIdAll", { RoomId: roomId, page: pages, rows: 20 }, function (msg) {
			if (msg.ErrorCode == 0) {
				total = msg.Data.total;
				if (total == 0) {
					return;
				}
				page = Math.ceil(total / 20);
				$('.jjj').remove();
				$.each(msg.Data.data, function (key, val) {
					var str = "<tr class='jjj removediv' ><td>" + val.NickName + "</td><td>" + val.ModifyTime + "</td><td><button class='cancelmanage'  onclick='' data-userid='" + val.UserId + "' data-usernames='" + val.NickName + "'  data-roomId='" + val.RoomId + "' data-rowid='" + val.Id + "'>撤销</button></td></tr>";
					$("#roommanage").append(str);
				});
			} else {
				return
			}
			if (total >= 20) {
				$(".page-component2").createPage2({
					roomId: roomId,
					pageCount: page,
					current: pages,
					backFn: function (p) {
						//console.log(p);
					}
				});
			}
		});



	});


	//删除房管
	var cancelmanage = $('.cancelmanage');
	cancelmanage.live("click", function confirmAct() {
		var rowid = $(this).data('rowid');
		var userId = $(this).attr('data-userid');
		var usernames = $(this).attr('data-usernames');
		roomIds = $(this).attr('data-roomId');
		alertShowYNznx("删除房管", "确定要执行此操作吗？", null, "取消");
		$('.alert-yes').live("click", function () {



			reqAjax("/user/delRoomusermanage", { Id: rowid }, function (msg) {
				//当客户端连接到服务器时通知客户端
				var onConnect = function (frame) {
					client.subscribe(roomIds);

					//当聊天服务器连接成功后发送欢迎语

					var uName = getStorage('nicheng') ? getStorage('nicheng') : getStorage('username');
					var msgobj = {
						from: uName,
						type: 'revoke',
						ext: {
							info: {
								userid: userId,
								username: usernames
							}
						},
						msgbody: '撤销房管',
					};
					var str = JSON.stringify(msgobj);
					if (str) {
						message = new Messaging.Message(str);
						message.destinationName = roomIds;
						client.send(message);
					}

				};
				//连接失败
				var onFailure = function (failure) {

				}
				$.getScript('/static/js/fushowim.js', function () {
					client = new Messaging.Client(chatconf.host, Number(chatconf.port), chatconf.clientId);
					//建立连接
					client.connect({
						userName: chatconf.user,
						password: chatconf.password,
						onSuccess: onConnect,
						onFailure: onFailure
					});

				});
				reqAjax("/user/findByRoomIdAll", { RoomId: roomId }, function (msg) {
					$('.jjj').remove();
					if (msg.ErrorCode == 0) {


						if (msg.Data.data == null) {
							return;
						}
						$.each(msg.Data.data, function (key, val) {
							var str = "<tr class='jjj'><td>" + val.ModifyBy + "</td><td>" + val.ModifyTime + "</td><td><button class='cancelmanage'  onclick=''  data-rowid='" + val.Id + "'>撤销</button></td></tr>";
							$("#roommanage").append(str);
						});
					} else {
						return;
					}
				}, true);
			});




		})
		$('.alert-no').live("click", function () {
			return
		})
	});

	//直播收益
	var pages = getStorage("settingpage");
	var page;
	if (pages == null) { pages = 1; }
	var uid = getStorage("Id");
	reqAjax("/user/findByRecipient", { UID: uid, page: pages, rows: 10 }, function (msg) {
		if (msg.ErrorCode != 0) {
			return;
		} else {
			total = msg.Data.total;
			if (total == 0) {
				return;
			}
			page = Math.ceil(total / 10);
		}
		$.each(msg.Data.data, function (key, val) {
			if (msg.Data.id == val.RecipientId) {
				var str = "<tr class='removedivlive'><td>" + val.GiveDate + "</td><td>" + val.BenefactorName + "</td><td>" + val.GiftName + "*" + val.GiftNum + "</td></tr>";
				$("#liveearnings").append(str);
			}
		});
		if (total >= 10) {
			$(".livepage").createPagelive({
				id: uid,
				pageCount: page,
				current: pages,
				backFn: function (p) { }
			});
		}
	}, true);

	$(".plugag").click(function () {
		plugflow();
	});
	//绑定银行卡
	$('#bangbankbtn').click(function () {
		var bankcardNum = $('#bankcardNum').val();
		var bankName = $('#bankName').val();
		var bankdePosit = $('#bankdePosit').val();
		if (bankName != "" && bankcardNum != "" && bankdePosit != "") {
			reqAjax("/user/anchorbindingbank", { UID: uid, bankcard: bankcardNum, bankname: bankName, bankdeposit: bankdePosit }, function (msg) {
				if (msg.ErrorCode != 0) {
					Dialog(msg.ErrorMsg, true, "确定", null, function () {
						$('.dialog').remove();
					}, null);
				} else {
					alertShow($('#bangbankbtn'), "提示", "成功绑定");
					$('#bangbank').css('display', 'none');
					$('#bankcardmsg').css('display', 'block');
					$('.bankName').html(bankName);
					$('.bankcardNum').html(bankcardNum.substring(0, 3) + "****" + bankcardNum.substring(14, 18));
					$('.bankdePosit').html(bankdePosit);
				}
			}, true);
		}
		if (bankcardNum == "") {
			$('#bankcardNum').css('background-color', 'coral');
			$('#bankcardNumshow').show();
		}
		if (bankName == "") {
			$('#bankName').css('background-color', 'coral');
			$('#bankNameshow').show();
		}
		if (bankdePosit == "") {
			$('#bankdePosit').css('background-color', 'coral');
			$('#bankdePositshow').show();
		}
	});

	//是否绑定银行卡
	reqAjax("/user/isbindingbank", { UID: uid }, function (msg) {
		if (msg.ErrorCode != 0) {
			Dialog(msg.ErrorMsg, true, "确定", null, function () {
				$('.dialog').remove();
			}, null);
		} else {
			if (msg == null) {
				return;
			}
			$.each(msg.Data.rows, function (key, val) {
				$('#pomegranateNum').html(val.PomegranateNum);
				if (val.IsBandingBank) {
					$('#bangbank').css('display', 'none');
					$('#bankcardmsg').css('display', 'block');
					$('#zsname').attr('value', val.RealName);
					$('.zsname').html(val.RealName);
					$('.bankName').html(val.BankName);
					var BankCardtext = val.BankCard.toString().substring(0, 4) + "**** ****" + val.BankCard.toString().substring(14, 18);
					$('.bankcardNum').html(BankCardtext);

					$('.bankdePosit').html(val.BankDeposit);
				} else {
					$('#zsname').attr('value', val.RealName);
					$('.zsname').html(val.RealName);
				}
			});
		}
	}, true);
});

function iscashmonth() {
	var uid = getStorage("Id");
	$.ajax({
		type: "post",
		url: "/user/ismonthcashing",
		dataType: "json",
		data: {
			UID: uid
		},
		async: true,
		success: function (msg) {
			if (bankTF) {
				if (msg.Data != null) {
					var getymd = msg.Data[0].CashingDate.substring(0, 10);
					var yy = getymd.substring(0, 4);
					var mm = getymd.substring(5, 7);
					var ymd = getNowFormatDate(1);
					var yys = ymd.substring(0, 4);
					var mms = ymd.substring(5, 7);


					if (yy == yys && mm == mms) {
						alertShowYNznx("提示", "每月只能申请一次结算", null);
						return;
					} else {
						alertShowYNznx("提示", "您将兑换<span style='color:#df0050; font-size:22px;'>" + $('#applycashingnumtext').val() + "</span>积分", null, "取消");
						$('body').find('.alert-yes').click(function () {
							maX();
						})
						$('body').find('.alert-no,.alert-close').click(function () {
							return;
						})
					}

				} else {
					alertShowYNznx("提示", "您将兑换<span style='color:#df0050; font-size:22px;'>" + $('#applycashingnumtext').val() + "</span>积分", null, "取消");
					$('body').find('.alert-yes').click(function () {
						maX();
					})
					$('body').find('.alert-no,.alert-close').click(function () {
						return;
					})
				}
			} else {
				alertShowYNznx("提示", "您还未绑定银行卡", null);
				return;
			}
		},
		error: function () {
			alert("失败!");
		}
	});
}
function maX() {
	applycashingnumtext = $('#applycashingnumtext').val();
	cashingtext = (applycashingnumtext * 80) / 10000;
	if (applycashingnumtext == "") {
		alertShowYNznx("提示", "不能为空", null);
		return;
	}
	if (applycashingnumtext == 0) {
		alertShowYNznx("提示", "不能结算0积分", null);
		return;
	}
	if (applycashingnumtext < 0) {
		alertShowYNznx("提示", "结算积分不能为负数", null);
		$('#applycashingnumtext').val('');
		$('#convertedinto').text('');
		return;
	}
	//判断申请结算数量是否足够
	reqAjax("/user/isenough", { UID: uid, applycashingnum: applycashingnumtext }, function (msg) {
		if (msg.ErrorCode != 0) {
			Dialog(msg.ErrorMsg, true, "确定", null, function () {
				$('.dialog').remove();
			}, null);
		} else {
			tJJs();
			hQYe();
		}
	}, true);
}

function getNowFormatDate(val) {
	var date = new Date();
	var seperator1 = "-";
	var seperator2 = ":";
	var month = date.getMonth() + 1;
	var strDate = date.getDate();
	if (month >= 1 && month <= 9) {
		month = "0" + month;
	}
	if (strDate >= 0 && strDate <= 9) {
		strDate = "0" + strDate;
	}
	var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate +
		" " + date.getHours() + seperator2 + date.getMinutes() +
		seperator2 + date.getSeconds();
	var ymd = date.getFullYear() + seperator1 + month + seperator1 + strDate;
	if (val == 1) {
		return ymd;
	} else {
		return currentdate;
	}
}
//提交结算ajax
function tJJs() {
	reqAjax("/user/anchorapplycashing", { UID: uid, applycashingnum: applycashingnumtext, cashing: cashingtext, cashingdate: getNowFormatDate() }, function (msg) {
		if (msg.ErrorCode != 0) {
			Dialog(msg.ErrorMsg, true, "确定", null, function () {
				$('.dialog').remove();
			}, null);
			return;
		} else {
			sttle();
		}
	}, true);
}

function sttle() {
	var pages = getStorage("sydhpage");
	var page;
	if (pages == null) {
		pages = 1;
	}
	reqAjax("/user/settlementdetail", { UID: uid, page: pages, rows: 20 }, function (msg) {
		if (msg.ErrorCode != 0) {
			// Dialog(msg.ErrorMsg,true,"确定",null,function() {
			// 	$('.dialog').remove();
			// },null);
			return;
		} else {
			total = msg.Data.total;
			page = Math.ceil(total / 20);
			$('.xxx').remove();
			if (msg.Data.rows == null) {
				return;
			}
			$.each(msg.Data.rows, function (key, val) {
				if (val.IsCashing) {
					val.IsCashing = "已结算";
				} else {
					val.IsCashing = "未结算";
				}
				var str = "<tr class = 'xxx'><td>" + val.CashingDate + "</td><td>" + val.ApplyCashingNum + "</td><td>" + val.Cashing + "</td><td>" + val.IsCashing + "</td></tr>";
				$("#returnexchange").append(str);
			});
			//分页

			if (total >= 20) {
				$(".page-component1").createPage1({
					id: uid,
					pageCount: page,
					current: pages,
					backFn: function (p) {
						//console.log(p);
					}
				});
			}
		}
	}, true);
}
//获取余额
function hQYe() {
	reqAjax("/user/isbindingbank", { UID: uid }, function (msg) {
		if (msg.ErrorCode != 0) {
			Dialog(msg.ErrorMsg, true, "确定", null, function () {
				$('.dialog').remove();
			}, null);
		} else {
			if (msg == null) {
				return;
			}
			$.each(msg.Data.rows, function (key, val) {
				$('#pomegranateNum').html(val.PomegranateNum);
			});
		}
	}, true);
}
//hQYe();
//串码流
function plugflow() {
	var uid = getStorage("Id");
	reqAjax("/page/getplugflow", { UID: uid }, function (g) {
		if (g.ErrorCode == 0) {
			$('.plugflow').val("");
			var plugmsg = g.Data.errMsg;
			$('.plugflow').val(plugmsg);
		} else {
			alert(g.ErrorMsg);
		}
	}, true);
};

function plugflows() {
	var uid = getStorage("Id");
	reqAjax("/page/getplugflow", { UID: uid, Type: 1 }, function (g) {
		if (g.ErrorCode == 0) {
			$('.plugflow').val("");
			var plugmsg = g.Data.errMsg;
			$('.plugflow').val(plugmsg);
		} else {
			alert(g.ErrorMsg);
		}
	}, true);
};

function cutImage() {
	var oop = this;
	this.option = {
		x: 0,
		y: 0,
		w: 100,
		h: 100,
		t: 'target',
		p: 'preview',
		o: 'preview_div'
	}
	this.init = function () {
		oop.target();
	}
	this.target = function () {
		$('#' + oop.option['t']).Jcrop({
			onChange: oop.updatePreview,
			onSelect: oop.updatePreview,
			minSize: [100, 100],
			aspectRatio: 1,
			setSelect: [oop.option['x'], oop.option['y'], oop.option['w'], oop.option['h']],
			bgFade: true,
			bgOpacity: .5
		});
	}
	this.updatePreview = function (obj) {
		if (parseInt(obj.w) > 0) {
			var rx = $('#' + oop.option['o']).width() / obj.w;
			var ry = $('#' + oop.option['o']).height() / obj.h;

			$('#' + oop.option['p']).css({
				width: Math.round(rx * $('#' + oop.option['t']).width()) + 'px',
				height: Math.round(ry * $('#' + oop.option['t']).height()) + 'px',
				marginLeft: '-' + Math.round(rx * obj.x) + 'px',
				marginTop: '-' + Math.round(ry * obj.y) + 'px'
			});
			$("#x1").val(obj.x);
			$("#y1").val(obj.y);
			$("#cw").val(obj.w);
			$("#ch").val(obj.h);
		}
	}
}



