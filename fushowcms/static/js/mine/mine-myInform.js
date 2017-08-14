
var cwzbphone;
var form;
var math = [];
var BindCaptcha = "";
var newNick;

$(function () {
	var url = location.search;
	if (url.indexOf("?") != -1) {
		$('#realName').show().siblings().hide();
		$('#cAnchor').css('border-bottom', '4px solid #009cfa').siblings().css('border-bottom', '4px solid #fff');
	}
	var uid = getStorage("Id");
	var from = getStorage("from");
	var pomegranateNum,
		thisSLZ,
		yue,
		yueText;

	reqAjax("/page/getuser", { UID: uid }, function (msg) {
		if (msg.ErrorCode != 0) {
			Dialog(msg.ErrorMsg, true, "确定", null, function () {
				$('.dialog').remove();
			}, null);
		} else {
			var userId = msg.Data.Id;
			var nicheng = msg.Data.UserName;
			var phone = msg.Data.Phone;
			cwzbphone = phone;
			yue = msg.Data.Balance;
			var nickname = msg.Data.NickName;
			var pic = msg.Data.Favicon;
			pomegranateNum = msg.Data.PomegranateNum;
			var integral = msg.Data.Integral;
			var type = msg.Data.Type;
			var lvw = userlevel(integral).tiao;
			var lvmaxw = userlevel(integral).dvalue;

			form = msg.Data.Form; //第三方登录成为主播用

			$("#userId").text(userId);
			$("#userNames").text(nickname);
			$("#nowNichen").text(nickname);
			$("#pomegranateNum").text(pomegranateNum);
			thisSLZ = $("#pomegranateNum").text();

			$("#nowPhone").val(phone);

			if (msg.Data.Form) {
				$("#mine-myInform-userPic").attr("style", "background-image:url(" + pic + ")");
			} else {
				if (pic != "") {
					$("#mine-myInform-userPic").attr("style", "background-image:url(" + pic.split('.')[0] + ".jpg)");
					$("#chengePic-big").attr("style", "background-image:url(" + pic.split('.')[0] + ".jpg)");
				} else {
					$("#mine-myInform-userPic").attr("style", "background-image:url(/static/images/default_avatar.jpg)");
				}
			}

			if (phone != "") {
				$("#bindphoneli").hide();
				var phonepictext = "<img src='/static/images/mine/phonestate.png' alt='' />";
				$("#phonestatePic").html(phonepictext);
				$("#phonestate").html("手机已验证");
				var phonestr = "<span>您已绑定手机：</span><span>" + phone.substring(0, 3) + "****" + phone.substring(7, 11) + "</span><span id='change'>修改</span>";
				$('#phonestr').html(phonestr);
				$('#phonestr').show().siblings().hide();
				$('#unPhone').hide();
				if (type == 1) {
					$('#realshow').show().siblings().hide();
				} else {
					$('#realshow').hide();
					$('#realphone').hide();
				}
			} else {
				$('#unPhone').show().siblings().hide();
				$('#realphone').show().siblings().hide();
				$("#editPassword").hide();
				$("#changephoneli").hide();
				$("#phonestatePic").html("");
				$("#phonestate").html("手机未验证");
			}

			$('#change').click(function () {
				$('#phonestr').hide().siblings().show();
				$('#phCheck').hide();
				// $('#check_num').hide();
			});

			$('.lvmin').addClass('level-' + userlevel(integral).number);
			if (integral == 0) {
				$('.lvmin').addClass('level-' + userlevel(integral).number);
				$('.lv_text').text('还差' + userintegral(integral) + '个经验值到');
				$('.lvmax').addClass('level-' + (userlevel(integral).number + 1));
			}
			if (userlevel(integral).name == "王者") {
				$('.lvmin').removeClass('sidebar-user-level-user').addClass('sidebar-user-level-user-max');
				$('.lvmax').remove();
				$(".lv_text").remove();
				$('.lv').css('width', '100%');
			} else if (integral != 0) {
				$('.lv_text').text('还差' + userintegral(integral) + '个经验值到');
				$('.lvmax').addClass('level-' + (userlevel(integral).number + 1));
				if (lvw < 100) {
					$('.lv').css('width', lvw + '%');
				} else {
					$('.lv').css('width', (lvw / lvmaxw) * 100 + '%');
				}
			}
			$("#phoneNum").text(phone);
			$(".yue").text(yue);
			yueText = parseInt($(".yue").text());

			if (form != "" && cwzbphone == "") {
				$('#changePhone').html("绑定手机")
			}
		}
	}, true);

	reqAjax('/user/isapplyExit', { UID: uid }, function (ret) {
		if (ret.ErrorCode != null) {
			if (ret.ErrorCode == "0") {
				$("#errorMsg").html("正在审核中");
				return;
			}
			if (ret.ErrorCode == "2003") {
				$("#errorMsg").html(ret.ErrorMsg);
				return;
			}
			$("#errorMsg").css("opcity", "0");
		} else {
			return;
		}
	});

	//签到活动检测
	reqAjax("/root/geteventlist", {}, function (msg) {
		if (msg == null) {
			$('.signin').hide();
			return;
		} else {
			if (msg.rows != null) {
				for (var i = 0; i < msg.rows.length; i++) {
					math.push(msg.rows[i].EventType);
					if (msg.rows[i].Number == null || msg.rows[i].Number == 0) {
						$('.signin').hide();
					}
				}
				if (math.indexOf(1) != -1) {
					IsSigned();
				} else {
					$('.signin').hide();
				}
			}
		}
	}, true);

	//点击签到
	$('.signin').click(function () {
		reqAjax("/user/signinadd", { UID: uid }, function (msg) {
			if (msg.ErrorCode == 0) {
				if (msg.Data.errMsg == 0) {
					return;
				} else {
					alertShowYNznx("提示", "签到成功,您成功获得" + msg.Data.errMsg + "个金币", null);
					$('.alert-OK,.alert-close').bind("click", function () {
						$('#signinpic').attr("src", "/static/images/mine/qiandao2.png");
						var zhival = yueText + parseInt(msg.Data.errMsg);
						$(".yue").text(zhival);
						yueText = parseInt($(".yue").html());
					});
				}
			} else {
				alertShowYNznx("提示", "今天您已经完成了签到，请明天再来哦~", null);
				return;
			}
		}, true);
	});


	//获取验证码
	$("#getting").click(function () {
		var pho = $("#newphone").val();
		if (!pho) {
			alertShowYNznx("提示", "手机号不能为空", null);
			return false;
		}
		reg_phone = /^1[34578]\d{9}$/;
		if (!reg_phone.test(pho)) {
			alertShowYNznx("提示", "手机格式不正确", null);
			return false;
		}
		var oldpas = $('#ordpass').val();
		var success = false;
		if (!oldpas) {
			alertShowYNznx("提示", "密码不能为空", null);
			return false;
		} else {
			reqAjax("/page/passmate", { UID: getStorage("Id"), pass: oldpas }, function (msg) {
				if (msg.ErrorCode != 0) {
					alertShowYNznx("提示", msg.ErrorMsg, null);
					return false;
				} else {
					var phonenew = $('#newphone').val();
					changeMobile(phonenew);
				}
			}, true);
		}
	});

	$("#chengePhone-Btn").click(function () {
		var code = $("#phone_check").val();
		var phone1 = $("#newphone").val();
		if (code == "") {
			alertShowYNznx("提示", "请输入验证码", null);
			return
		}
		reqAjax("/page/isverification", { mobile: phone1, code: code }, function (msg) {
			if (msg.ErrorCode != 0) {
				alertShowYNznx("提示", msg.ErrorMsg, null);
				return
			} else {
				// $("#check_num").css({
				// 	'display': 'none'
				// });
				$("#chengePhone-Btn").css({
					'display': 'block'
				});
				var phone = $("#newphone").eq(0).val();
				if ($("#newphone").val().length != 11) {
					alertShowYNznx("提示", "你输入的手机号位数不对,请重新输入", null);
					return;
				}
				var uid = getStorage("Id");
				reqAjax("/user/userup", { UID: uid, Phone: phone }, function (msg) {
					if (msg.ErrorCode != 0) {
						alertShowYNznx("提示", msg.ErrorMsg, null);
					} else {
						alertShowYNznx("提示", "恭喜您成功修改手机号。你绑定的手机号为：" + phone, null);
						$('.alert-OK,.alert-close').bind("click", function () {
							location.reload();
						});
					}
				}, true);
			}
		}, true);
	})

	//上传图片
	$("#file_upload").bind("change", function () {
		var uid = getStorage("Id");
		$("#user_id").val(uid);
		upload(this, $("#insertPicForm"), $("#mine-myInform-userPic"));
		location.reload();
	});

	$('#chengeName-Btn').click(function () {
		var params = {};
		params.UID = getStorage("Id");
		reqAjax("/user/checknicknamechange", params, function (msg) {
			if (msg.ErrorCode != 0) {
				alertShowYNznx("提示", msg.ErrorMsg, null);
			} else {
				nickflag = msg.Data.NickFlag;
			}
		});
		if (nickflag) {
			alertShowYNznx("提示", "您已经修改过昵称，昵称只能修改一次", null);
			return;
		} else {
			var nickname = $("#nicheng").val()
			var nick_check = /^.{3,10}$/;
			if (!nickname) {
				alertShowYNznx("提示", "昵称不能为空", null);
				return;
			} else if (!nick_check.test(nickname)) {
				alertShowYNznx("提示", "昵称格式不正确,昵称长度应为3~10位", null);
				return;
			}

			reqAjax("/page/checknick", { nickname: nickname }, function (ret) {
				if (ret.Data.state) {
					var pars = {};
					pars.UID = getStorage("Id");
					pars.nickname = $("#nicheng").val();
					pars.nickflag = "1";
					reqAjax("/user/changenickname", pars, function (msg) {
						if (msg.ErrorCode != 0) {
							alertShowYNznx("提示", msg.ErrorMsg, null);
						} else {
							setStorage("nicheng", pars.nickname);
							location.reload();
						}
					}, true)
				} else {
					alertShowYNznx("提示", "昵称已存在", null);
					return;
				}
			});
		}
	});

	$("#chengePassw-Btn").click(function () {
		var zz6 = /^(?![0-9]+$)[0-9A-Za-z]{6,16}$/;
		var num1 = $("#newpwd1").val();
		var num2 = $("#newpwd2").val();
		var num3 = $("#newpwd3").val();
		if (num2 != num3) {
			alertShowYNznx("提示", "两次密码输入不一致！", null);
			$("#newpwd1").val("");
			$("#newpwd2").val("");
			$("#newpwd3").val("");
			return;
		} else if (!zz6.test(num2) || !zz6.test(num3)) {
			alertShowYNznx("提示", "密码应为6-16位", null);
			$("#newpwd1").val("");
			$("#newpwd2").val("");
			$("#newpwd3").val("");
			return;
		}
		var username = getStorage("username");
		var userid = getStorage("Id");

		reqAjax("/user/pcpassup", { username: username, UID: userid, password: num1, newpassword: num3 }, function (msg) {
			if (msg.ErrorCode != 0) {
				alertShowYNznx("提示", msg.ErrorMsg, null);
				$(".alert-OK,.alert-close").bind("click", function () {
					$('#newpwd1').val("");
					$('#newpwd2').val("");
					$('#newpwd3').val("");
				});
			} else {
				alertShowYNznx("提示", "修改成功", null);
				$(".alert-OK,.alert-close").bind("click", function () {
					location.reload();
				});
			}
		}, true);
	});
	//切换选项卡
	var $menu = $('#mine-myInform-ul li');
	$menu.click(function () {
		var index = $menu.index(this);
		$('.tabs_content').eq(index).show().siblings().hide();
		$('#mine-myInform-ul').show();

		$(this).css({
			color: "#cc00000",
			borderBottom: "4px solid #009CFA"
		}).siblings().css({
			color: "#cc00000",
			background: "white",
			borderBottom: ""
		});
	});

	$('#realphone span').click(function () {
		$("#bindphoneli").click();
	});
	$('#unPhone span').click(function () {
		$("#bindphoneli").click();
	});

	$('.unPhoneSpan').click(function () {
		if (form == "") {
			$("#changephoneli").click();
		} else {
			$("#bindphoneli").click();
		}
	});

	$('#changePhone').click(function () {
		if (cwzbphone == "") {
			$("#bindphoneli").click();
		} else {
			$("#changephoneli").click();
		}
		$.ajax({
			cache: false,
			type: "get",
			url: "/getimagecode",
			async: false,
			error: function (request) {
				alert("图形验证码获取失败");
			},
			success: function (data) {
				BindCaptcha = data.Data.CaptchaId;
				$(".bind-img").attr("src", data.Data.ImageURL)
			}
		});
	})

	$('#picno,.picheader-col').click(function () {
		$('.chengePic-divshow2').hide();
		$('.chengePic-divshow1').show();
	});

	if (from == "sidebarunbtn") {
		$('#cAnchor').click();
		setStorage("from", "");
	}

	$.ajax({
		cache: false,
		type: "get",
		url: "/getimagecode",
		async: false,
		error: function (request) {
			alert("Connection error");
		},
		success: function (data) {
			BindCaptcha = data.Data.CaptchaId;
			$('.bind-img').attr("src", data.Data.ImageURL)
		}
	});

	$('.bind-img').click(function () {
		$.ajax({
			cache: false,
			type: "get",
			url: "/getimagecode",
			async: false,
			error: function (request) {
				alert("图形验证码获取失败");
			},
			success: function (data) {
				BindCaptcha = data.Data.CaptchaId;
				$(".bind-img").attr("src", data.Data.ImageURL);
			}
		});
	});

	//修改昵称
	nickName();

	//jquery.form.js 的ajax提交表单
	$('#picyes').click(function () {
		$('form').on('submit', function () {
			var picmsg = $("#inputImage").val();
			var xmsg = Math.round(parseInt($('#xx').val()));
			var ymsg = Math.round(parseInt($('#yy').val()));
			var imgdivwidth = parseInt($(".cropper-canvas img").width());
			var imgdivheight = parseInt($(".cropper-canvas img").height());
			var x = xmsg;
			var y = ymsg;
			var x1 = Math.round(maxWidth * 122 / imgdivwidth);
			var y1 = Math.round(maxHeight * 122 / imgdivheight);
			var datatimename = new Date().getTime();
			$(this).ajaxSubmit({
				url: '/page/imageresizerUpload',
				type: 'post',
				dataType: "json",
				timeout: "3000",
				data: {
					uploadFile: picmsg,
					FileName: datatimename + "." + imgfilename,
					X: x,
					Y: y,
					X1: x1,
					Y1: y1
				},
				success: function (msg) {
					if (msg.ErrorMsg == "ok") {
						$('#mine-myInform-userPic').css("background-image", "url(../../static/upload/" + datatimename + "." + "jpg" + ")");
						reqAjax("/user/userUpFavicon", { UID: getStorage("Id"), Favicon: "/static/upload/" + datatimename + ".jpg" }, function (msg) {
							if (msg.ErrorCode != 0) {
								alertShowYNznx("提示", msg.ErrorMsg, null);
							} else {
								alertShowYNznx("提示", "修改成功", null);
								$(".alert-OK,.alert-close").bind("click", function () {
									setStorage("Favicon", '/static/upload/' + datatimename + '.jpg');
									window.location.reload();
								})
							}
						}, true);
					} else {
						alertShowYNznx("提示", msg.ErrorMsg, null);
						$(".alert-OK,.alert-close").bind("click", function () {
							window.location.reload();
						})
					};
				},
				error: function (msg) { },
			});
			return false; //阻止表单默认提交
		});
	});

	//选择图片按钮跳转
	$('.picbody-btn').click(function () {
		$('.changepicbtn').click();
	})

	$('.zhuce-agree-box0').click(function () {
		$('.zhuce-agree-box0').css('display', 'none');
		$('.zhuce-agree-box').attr('data-check', 'no');
	})
	$('.zhuce-agree-box img').click(function () {
		$('.zhuce-agree-box0').css('display', 'block');
		$('.zhuce-agree-box').attr('data-check', 'yes');
	})


	$("#realName-Btn").click(function () {

		var phone = $("#myphone").val();
		var myname = $("#myname").val();
		var myidnumber = $("#myidnumber").val();
		var checked = $(".zhuce-agree-box").attr('data-check');
		if (checked == "yes") {
			alertShowYNznx("提示", "请同意《富秀直播个人直播协议》", null);
			return;
		}

		if ($("#myphone").val() == "") {
			alertShowYNznx("提示", "请输入手机号", null);
			return;
		}
		if ($("#myname").val().length == 0) {
			alertShowYNznx("提示", "请输入真实姓名", null);
			return;
		}
		if ($("#myidnumber").val().length == 0) {
			alertShowYNznx("提示", "请输入身份证号", null);
			return;
		}
		if ($("#myphone").val() != cwzbphone) {
			alertShowYNznx("提示", "请输入该账号绑定的手机号", null);
			return;
		}
		if ($("#myphone").val().length != 11) {
			alertShowYNznx("提示", "手机号位数不正确，请重新输入", null);
			return
		}
		if ($("#myidnumber").val().length != 18) {
			alertShowYNznx("提示", "身份证位数不正确，请重新输入", null);
			return
		}
		if ($("#idp_upload_name").val().length == 0) {
			alertShowYNznx("提示", "请上传图片", null);
			return;
		}
		var myidnumber = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/;
		if (!myidnumber.test($("#myidnumber").val())) {
			alertShowYNznx("提示", "身份证号格式不正确", null);
			return;
		}
		var uid = getStorage("Id");
		var uploadFile = $("#idp_upload_name").val();


		reqAjax("/user/useridnumber", { uploadFile: uploadFile, UID: uid, Phone: phone, IdNumber: $("#myidnumber").val(), Name: myname }, function (msg) {

			if (msg.ErrorCode != 0) {
				alertShowYNznx("提示", msg.ErrorMsg, null);
			} else {
				alertShowYNznx("提示", "申请成功", null);
				$(".alert-OK,.alert-close").click(function () {
					window.location.reload();
				})
			}
		}, true);
	})

	/*************绑定手机****************/
	//获取验证码
	$("#bindgetting").click(function () {
		var pho = $("#bindnewphone").val();
		if (!pho) {
			alertShowYNznx("提示", "手机号不能为空", null);
			return false;
		}
		reg_phone = /^1[34578]\d{9}$/;
		if (!reg_phone.test(pho)) {
			alertShowYNznx("提示", "手机格式不正确", null);
			return false;
		}

		var oldpas = $('#bindordpass').val();
		var success = false;
		var oldpas_pass = /^.{6,16}$/;
		if (!oldpas_pass.test(oldpas)) {
			alertShowYNznx("提示", "密码格式不正确,密码长度应为6~16位", null);
			return false;
		}
		if (!oldpas) {
			alertShowYNznx("提示", "密码不能为空", null);
			return false;
		} else {
			var phone = $("#bindnewphone").val();
			reqAjax("/page/verificationCode", { mobile: phone, keycode: $('#bind-img-verify').val(), captcha: BindCaptcha }, function (msg) {
				if (!msg) {
					alertShowYNznx("提示", "发送失败", null);
					return
				}
				if (msg.ErrorCode != 0) {
					alertShowYNznx("提示", msg.ErrorMsg, null);

				} else {
					bindchangeMobile(pho);
					$('#bindnewphone,#bindordpass').attr("disabled", "true");
				}
			}, true);

		}
	});

	//添写验证码
	$('#bindcheck_num').click(function () {
		var code = $("#bindphone_check").val();
		var phone = $("#bindnewphone").val();
		var newpass = $("#bindordpass").val();
		if (code == "") {
			alertShowYNznx("提示", "请输入验证码", null);
			return
		}
		reqAjax("/page/isverification", { mobile: phone, code: code }, function (msg) {
			if (msg.ErrorCode != 0) {
				alertShowYNznx("提示", msg.ErrorMsg, null);

			} else {
				bindPhone(phone, newpass);
			}
		}, true);
	});
})

//上传图片通用函数定义
function upload(o, f, i, n) { //o,上传文件的input;f,上传的form；i上传成功后预览图片的为位置；n存取图片名称的地方
	var uploadFile = $(o).val();
	if (uploadFile == "") {
		alertShowYNznx("提示", "上传文件为空", null);
		return;
	}
	var file = uploadFile.lastIndexOf("\\");
	var name = uploadFile.substring(file + 1);
	var patrn = /[@#\)(\$%\^&\*]+/g;
	if (patrn.test(name)) {
		alertShowYNznx("提示", "您输入的数据含有非法字符！", null);
		return;
	}
	var fileName = name.substring(name.lastIndexOf(".") + 1).toLowerCase();
	if (fileName != "jpg" && fileName != "jpeg" && fileName != "pdf" && fileName != "png" && fileName != "dwg" && fileName != "gif") {
		alertShowYNznx("提示", "请上传正确格式图片！", null);
		return;
	}
	var hideForm = f,
		$file = $(o).remove();
	hideForm.append($file);
	var options = {
		dataType: "json",
		beforeSubmit: function () {
			alertShowYNznx("提示", "正在上传！", null);
		},
		success: function (msg) {
			if (msg.state == "success") {
				alertShowYNznx("提示", "上传成功", null);
				i.css("background-image", "url(../../static/upload/" + msg.imgNmae + ")");
				i.css("background-repeat", "no-repeat");
				n.val(msg.imgNmae);
			} else {
				alertShowYNznx("提示", msg.state, null);
			}
			return 1;
		}
	};
	hideForm.ajaxSubmit(options);
	return !1;
}

function changeMobile(phone) {
	reqAjax("/page/verificationCode", { mobile: phone, keycode: $('#change-img-verify').val(), captcha: BindCaptcha }, function (msg) {
		if (!msg) {
			alertShowYNznx("提示", "发送失败", null);
			return;
		}
		if (msg.ErrorCode != 0) {
			alertShowYNznx("提示", msg.ErrorMsg, null);
			return;
		} else {
			alertShowYNznx("提示", "发送成功", null);
			$('#chengePhone-Btn').show();
			var btn = $(this);
			var count = 60;
			$('#phCheck').css('display', 'block');
			$('#getting').css({ 'width': '165px', 'left': '188px' });
			var resend = setInterval(function () {
				count--;
				if (count > 0) {
					$("#getting").val(count + "秒后可重新获取");
					$('#getting').attr("disabled", true).css({
						'cursor': 'not-allowed'
					});
				} else {
					clearInterval(resend);
					$('#getting').css({ 'width': '165px', 'left': '188px' });
					$('#newphone,#ordpass').removeAttr('disabled');
					$("#getting").val("再次发送").removeAttr('disabled');

				}
			}, 1000);
			btn.attr('disabled', true).css({
				'cursor': 'not-allowed',
				'border': '1px solid gainsboro'
			});
			$('#newphone,#ordpass').attr("disabled", true);
		}
	}, true);
}

//身份证图片上传
function idpclearfile() {
	$("#idp_upload").unbind().change(function () {
		upload(this, $("#insertIdPicForm"), $("#realName-Pre"), $("#idp_upload_name"));
	});
}

//签到记录
function IsSigned() {
	var uid = getStorage("Id");
	reqAjax("/user/IsSigned", { UID: uid }, function (dat) {
		if (dat.ErrorCode == 0) {
			if (dat.Data.flag == true) {
				$('#signinpic').attr("src", "/static/images/mine/qiandao2.png");
			} else {
				$('#signinpic').attr("src", "/static/images/mine/qiandao1.png");
			}
		} else {
			return false;
		}
	}, true);
}

//修改昵称
function nickName() {
	var params1 = {};
	params1.UID = getStorage("Id");
	reqAjax("/user/checknicknamechange", params1, function (msg) {
		if (msg.ErrorCode != 0) {
			alertShowYNznx("提示", msg.ErrorMsg, null);
		} else {
			newNick = msg.Data.NickFlag;
		}
	});
	if (newNick) {
		$('#cNickname').hide();
		$('#chengeName').hide();
	}
}