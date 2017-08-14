//礼物展示//
$(function() {
	reqAjax("/page/getgiftlist",{},function(msg) {
			var i = 0;
			var giftList = [];
			var giftNamesList = [];
			var giftPriceList = [];
			var giftID = [];

			$.each(msg.rows, function() {
				var imageG = msg.rows[i].GiftPicture;
				giftList.push(imageG);
				var giftName = msg.rows[i].GiftName;
				giftNamesList.push(giftName);
				var giftPrice = msg.rows[i].BuyNumber;
				giftPriceList.push(giftPrice);
				var giftId = msg.rows[i].Id;
				giftID.push(giftId);
				str = "";
				str += "<li class='giftBtn-Pic' style='background-image:url(" + imageG + ");background-size:100%'></li>"
				$("#room-video-bottom-ul").append(str);
				i++;
				//鼠标经过礼物图标效果
				$(".giftBtn-Pic").hover(function() {
					var index = $("#room-video-bottom-ul li").index($(this));
					var left = -108 + 36 * index;
					$(".gift-Btn").css("display", "block");
					$(".gift-Btn").css("left", left);
					var giftPic = giftList[index];
					$(".gift-Btn img").attr("src", giftPic);
					$(".giftName1").text(giftNamesList[index]);
					$(".giftPrice1").text(giftPriceList[index] + "石榴籽");
				}, function() {
					var index = $("#room-video-bottom-ul li").index($(this));
					$(".gift-Btn").css("display", "none");
				});
			});

			//点击礼物图标
			$("body").on('click', '.giftBtn-Pic', function() {
				var ID = getStorage("Id");
				var niCheng = getStorage("nicheng");
				var userName = getStorage("username");
				if(!ID && !userName && !niCheng) {
					$('#loginBox').css('display', 'block');
					return;
				}
				var index = $(".giftBtn-Pic").index($(this));
				reqAjax("/page/getuser",{UID:1},function(msg){
					if(msg.ErrorCode!=0) {
						Dialog(msg.ErrorMsg,true,"确定",null,function() {
							$('.dialog').remove();
						},null);
					}else {
						var uid = getStorage("Id");
						var money = msg.Data.Balance;
						var giftName = giftNamesList[index];
						if(money <= 100) {
							$(".moneyBox").css("display", "block");
							$(".moneyBox").load("/user/roomGift_noMenoy", function() {
								$(".cancel-btn").click(function() {
									$(".moneyBox").css("display", "none");
								});
							});
						} else {
							$(".moneyBox").css("display", "block");
							$(".gift-Balance").text(money);
							var aaa = giftList[index];
							$(".moneyBox").load("/page/roomGift_gift", function() {
								var sss = $(".roomGift-icon2");
								sss.attr("src", aaa);
								$(".roomGiftName1").text(giftName);

								$(".cancel-btn").click(function() {
									$(".moneyBox").css("display", "none");
								});

								$(".roomGift-pay").click(function() {
									var giftimgurl = $(this).parents('.roomGift-message2').find('.roomGift-icon2').attr('src');
									var giftname = $(this).parents('.roomGift-message2').find('.roomGiftName1').text();

									uid = getStorage("Id");
									anchorN = getPar("anchorId");
									
									reqAjax("/user/givegiftnumadd",{GiftId: giftID[index],UID: uid,AnchorId: anchorN,Number: 1},function(msg){
										if(msg.ErrorCode!=0) {
											Dialog(msg.ErrorMsg,true,"确定",null,function() {
												$('.dialog').remove();
											},null);
										}else {
											client.send(giftname);
											var oDanmu = $("<span class='danmuM'>");
											var rTop = Math.random() * 650;
											oDanmu.css("top", rTop);
											oDanmu.animate({
												right: "1360px"
											},10000, function() {
												if(oDanmu.css("right") == "1340px") {
													oDanmu.detach();
												}
											});
											var val = '<img src="'+imgurl+'"/>'
											oDanmu.html(val);
											$("#danmu").append(oDanmu);
											$(".moneyBox").css("display", "none");
											}			
										}
									},true);
									
								})
							});
						}
					}
				},true);
				});
			});
	});
});

function getPar(par) {
	//获取当前URL
	var local_url = document.location.href;
	//获取要取得的get参数位置
	var get = local_url.indexOf(par + "=");
	if(get == -1) {
		return "";
	}
	//截取字符串
	var get_par = local_url.slice(par.length + get + 1);
	//判断截取后的字符串是否还有其他get参数
	var nextPar = get_par.indexOf("&");
	if(nextPar != -1) {
		get_par = get_par.slice(0, nextPar);
	}
	return get_par;
}