
	var pages = getStorage("guesspage");
	var page;
	if(pages==null){pages=1;}
	var uid = getStorage('Id')
	$(function(){
		reqAjax('/user/getSupportUidList',{UID:uid,page:pages,rows:8},function(msg){
			if(msg.ErrorCode == 0){
				total = msg.Data.total;
				page = Math.ceil(total / 8);
				var html ='';
				var winBox = document.getElementById('winBox');
				var failBox = document.getElementById('failBox');
				var prizenNmber;
				var isWin;
				if(msg.ErrorMsg == "ok"){
				$.each(msg.Data.list, function(e){
					console.log(e)
					if(msg.Data.list[e].PrizenNmber==0){//胜负状态，0：负，1：胜
						prizenNmber = "负";
					}
					if(msg.Data.list[e].PrizenNmber==1){//胜负状态，0：负，1：胜
						prizenNmber = "胜";
					}
					if(msg.Data.list[e].IsWin==1){
						isWin = "胜";
					}
					if(msg.Data.list[e].IsWin==2){
						isWin = "负";
					}
					if(msg.Data.list[e].IsWin==0){
						isWin = "竞猜中";
					}
					var r6r = "<tr class='removedivzc'><td>"+ msg.Data.list[e].PeriodsId +"</td><td>"+ msg.Data.list[e].SupporTime +"</td><td>"+ msg.Data.list[e].ProductName +"</td><td>"+ msg.Data.list[e].SupportState +"</td><td>"+ msg.Data.list[e].SupporNumber +"</td><td>"+ msg.Data.list[e].Odds +"</td><td style='color:#fba535'>"+ isWin +"</td></tr>";
					$("#myBetting").append(r6r);
					
				});
			}
						//分页
				if(total>=8){
					$(".support").createPagezc({
						id:uid,
						pageCount:page,
						current:pages,
						backFn:function(p){}
					});
				}
			}

			});

		
	        	//礼物记录
	    reqAjax('/user/findByBenefactor',{UID: uid,page:pages,rows:8},function(msg){
			if (msg.data == null) {
				//console.log("礼物记录没有数据");
				return;
			}
			total = msg.total;
			page = Math.ceil(total / 8);
			$.each(msg.data, function(key,val){
					var str = "<tr class='removediv'><td>"+val.GiveDate +"</td><td>"+val.RecipientName +"</td><td>"+ val.GiftName + "×" +val.GiftNum +"</td><td>"+val.AllNumber +"</td></tr>";
					$("#myGuess").append(str);
			});
			if(total>=8){
				$(".gift").createPage1({
					id:uid,
					pageCount:page,
					current:pages,
					backFn:function(p){}
				});
			}
		});	
	});	