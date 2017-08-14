$(function() {
		alert("好用好用！");
		$.ajax({
			type: "post",
			url: "/page/periodmores",
			dataType: "json",
			//	data: ,
			async: true,
			success: function(msg, xhr) {
				alert("请求成功！");
				var i = 0;
				var str = "";

				$.each(msg.date, function() {
					var kaishiTime = msg.date[i].StartTime;
					$(".itemName").text(kaishiTime + "开始竞猜");
					//alert(kaishiTime);

					$(".deadline-Time").text(msg.date[i].EndTime);
					var qishu = msg.date[i].PeriodsId;
					var oli = $("<li>");
					oli.html(i + 1);
					$(".newGusse-Ul").append(oli);

					i++;
					console.log(i);
				});
				/***************TabClick******************/
				$(".newGusse-Ul li").click(function() {
					var index = $(".newGusse-Ul li").index($(this));
					var kaishiTime = msg.date[index].StartTime;
					$(".itemName").text(kaishiTime + "开始竞猜");
					console.log(kaishiTime);
					var qishu = msg.date[index].PeriodsId;
					var PeriodsIdNow = msg.date[index].PeriodsId;
					console.log(PeriodsIdNow);
					//alert(index);
					var str2 = "";
					$(".newGusse-Tabcontent").empty();
					/*************/
					str2 += "	<p class='newGusse-deadline'>本场比赛预测时间截止至";
					str2 += "<span class='deadline-Time'>" + msg.date[index].EndTime + "</span>";
					str2 += "<span style='color:darkred;font-size:14px ;  display: inline-block; margin-left: 20px;'>期数:" + qishu + "</span></p>";
					str2 += "<div class='titleBox'><div class='newGusse-group' style='float: left;'>";
					str2 += "<div class='group-Pic'></div><div class='group-Iform'>";
					str2 += "</div></div><div class='vsB'></div><div class='newGusse-group' style='float: left;'>";
					str2 += "<div class='group-Pic'></div><div class='group-Iform'></div></div></div>";

					$(".newGusse-Tabcontent").append(str2);
					/***********TabContent*************/

					$.ajax({
						type: "post",
						url: "/page/perproname",
						async: true,
						dataType: "json",
						data: {
							"PeriodsId": PeriodsIdNow
						},

						success: function(msg2, xhr) {
							var i = 0;
							var Name = msg2.data[i].ProductName;

							var Product11 = msg2.data[1].State2Hot;

							console.log(Product11);
							$.each(msg2.data, function() {
								//alert("json遍历！");
								/******项目父*******/
								var Product = msg2.data[i].ProductName;

								var odiv = "";
								odiv += "<div class='newGusse-ChoiceB'><p style='font-weight: bolder;color: #c3313c;' class='newGusse-kindName'>本场比赛" + Product + "？</p>";
								odiv += "<ul class='newGusse-ChoiceB-UL'></ul>";
								odiv += "<div class='newGusse-ChoiceYesBox'><input placeholder='请输入支持数量！' class='newGusse-yesBtn' value =''/>";
								odiv += "<div class='newGusse-Yes'>确认</div></div></div>";
								$(".newGusse-Tabcontent").append(odiv);

								var choiceInput = "";

								// $(".newGusse-ChoiceB").append(choiceInput);	
								for(var j = 1; j <= 8; j++) {
									var state = "State" + j;
									state = msg2.data[i][state];

									var hot = "State" + j + "Hot";
									hot = msg2.data[i][hot];

									var odds = "State" + j + "Odds";
									odds = msg2.data[i][odds];
									//alert(hot);
									if(state != "") {
										var choiceStr = "";
										choiceStr += "<li class='newGusse-Choice'><input type='radio' name='chioces" + i + "'class='newGusse-Check'/>";
										choiceStr += "<div class='group-Pic2'></div>";
										choiceStr += "<span class='newGusse-ChoiceName'>" + state + "</span>";
										choiceStr += "<div class='group-Iform2'><ul class='group-IformUL2'>";
										choiceStr += "	<li>热度:<span >" + hot + "</span></li><li>赔率:<span class='peilv'>" + odds + "</span></li>";
										choiceStr += "</ul></div></li>";
										$(".newGusse-ChoiceB-UL").eq(i).append(choiceStr);

										choiceStr = "";
									}
								}
								i++;
							});

							$(".newGusse-Yes").click(function() {
								var indexInput = $(".newGusse-Yes").index($(this));

								var uid = getStorage("Id");
								var productid = msg2.data[indexInput].ProductId;
								var Num = $(".newGusse-yesBtn").eq(indexInput).val();
								//var eee = "input[name='chioces"+indexInput+"']:checked";
								var nameStr = $("input[name='chioces" + indexInput + "']");
								var c = 0;

								$.each(nameStr, function() {
									if($(this).attr("checked")) {
										return(c);
									};
									c++;
								});
								var supE = "";
								supE += "#" + msg2.data[indexInput].ProductId + ">" + c;
								//合成码
								var odds = $(".peilv").eq(c).text(); //赔率

								alert(supE);

								if(Num == NaN) {
									alert("请输入支持数量!");
								} else {
									//ajax
									$.ajax({
										type: "post",
										url: "/user/supportadd",
										data: {
											UID: uid,
											ProductId: productid,
											supEncoding: supE,
											supporNumber: Num,
											Odds: odds
										},
										async: true,
										success: function() {
											alert("投注成功");
										},
										error: function() {
											alert("投注失败");
										}

									});

								}

							});
						}, //success
						error: function() {
								alert("456");
							} //error
					});
				}); //ajax
			}, //success
			error: function() {
					alert("请求失败！");
				} //error
		}); //ajax
	}) //function