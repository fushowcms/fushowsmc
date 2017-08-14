//$("#mine-myGusse-ul li").eq(0).click(function(){
//	//alert(123);
//	$("#mine-myGusse-ul li").eq(0).removeClass("myGift-Off");
//	$("#mine-myGusse-ul li").eq(0).addClass("myGift-On");
//	//alert(123);
//	$("#mine-myGusse-ul li").eq(1).removeClass("myGift-On");
//	$("#mine-myGusse-ul li").eq(1).addClass("myGift-Off");
//	
//	$("#winBox").css("display","block");
//	$("#failBox").css("display","none");
//})
//$("#mine-myGusse-ul li").eq(1).click(function(){
//	//alert(123);
//	$("#mine-myGusse-ul li").eq(1).removeClass("myGift-Off");
//	$("#mine-myGusse-ul li").eq(1).addClass("myGift-On");
//	//alert(123);
//	$("#mine-myGusse-ul li").eq(0).removeClass("myGift-On");
//	$("#mine-myGusse-ul li").eq(0).addClass("myGift-Off");
//	
//	$("#winBox").css("display","none");
//	$("#failBox").css("display","block");
//})
$(function() {
	var uid = getStorage("Id");
	var pages = getStorage("guesspage");
	var page;
	
	reqAjax("/page/getuser",{UID:uid},function(msg){
		if(msg.ErrorCode!=0) {
			Dialog(msg.ErrorMsg,true,"确定",null,function() {
				$('.dialog').remove();
			},null);
		}else {
			var yue = msg.Data.Balance;
			var nicheng = msg.Data.NickName;
			var phone = msg.Data.Phone;
			$("#myBlance").text(yue + "石榴籽");
		}
	},true);

	if(pages == null) {
		pages = 1;
	}
	var uid = getStorage('Id');
	$(function() {
		$.ajax({
			type: "post",
			url: "/user/getSupportUidList",
			data: {
				UID: uid,
				page: pages,
				rows: 20
			},
			dataType: "json",
			success: function(msg) {
				total = msg.total;
				page = Math.ceil(total / 20);
				var html = '';
				var winBox = document.getElementById('winBox');
				var failBox = document.getElementById('failBox');
				var prizenNmber;
				var isWin;

				if(msg.state) {
					//IE8适配foreach
					if(!Array.prototype.forEach) {
						Array.prototype.forEach = function forEach(callback, thisArg) {
							var T, k;
							if(this == null) {
								throw new TypeError("this is null or not defined");
							}
							var O = Object(this);
							var len = O.length >>> 0;
							if(typeof callback !== "function") {
								throw new TypeError(callback + " is not a function");
							}
							if(arguments.length > 1) {
								T = thisArg;
							}
							k = 0;
							while(k < len) {
								var kValue;
								if(k in O) {
									kValue = O[k];
									callback.call(T, kValue, k, O);
								}
								k++;
							}
						};
					}
					msg.data.forEach(function(e) {
						if(e.PrizenNmber == 0) { //胜负状态，0：负，1：胜
							prizenNmber = "负";
						}
						if(e.PrizenNmber == 1) { //胜负状态，0：负，1：胜
							prizenNmber = "胜";
						}
						if(e.IsWin == 1) {
							isWin = "胜";
						}
						if(e.IsWin == 2) {
							isWin = "负";
						}
						if(e.IsWin == 0) {
							isWin = "竞猜中";
						}
						winBox.innerHTML += '<ul class="mine-liveGusseBox-ul2"><li style="width: 20%;">' + e.PeriodsId + '</li><li style="width: 20%;">' + e.ProductName + '</li><li style="width: 20%;">' + prizenNmber + '</li><li style="width: 20%;">' + e.SupporNumber + '</li><li style="width: 20%;">' + isWin + '</li></ul>';
					});
				}
				//分页
				if(total >= 20) {
					$("#tcdPageCode").createPage({
						pageCount: page,
						current: pages,
						backFn: function(p) {
							//console.log(p);
						}
					});
				}
			}
		});
		//礼物记录
		reqAjax("/user/findByUserIdAll",{UID:uid},function(data) {
			if(data.ErrorCode!=0){
				return;
			}else{
				$.each(data.Data, function(key, val) {
					if(data.Data.id == val.BenefactorId) {
						var str = "<tr class='removediv'><td>"+val.GiveDate +"</td><td>"+val.RecipientId +"</td><td>"+val.RecipientName +"</td><td>"+ val.GiftName + "×" +val.GiftNum +"</td></tr>";
						$("#transaction_list").append(str);
					}
				});
			}
		},true);
	});	
});
$(function() {
	var menu = $('#mine-liveGusseBox-top li');
	menu.click(function() {
		var index = menu.index(this);
		$('.tabs_content').eq(index).show().siblings().hide();
		$('#mine-liveGusseBox-top').show();

		$(this).css({
			color: "white",
			background: "deepskyblue"
		}).siblings().css({
			color: "gray",
			background: "white"
		})
	})
})