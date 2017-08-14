$(function() {  
				var pages = getStorage("gzpage");
				if(pages==null){
					pages=1;
				}
				var page;
			var id =getStorage("Id");
			var url = "/user/getmyattentionlist";
			$.ajax({  
				type: "post",  
				url: url,  
				dataType: "json",  
				data: {
					UID:id,
					page:pages,
					rows:8
				},  
				success: function(msg){
					
				total = msg.total;
				page = Math.ceil(total / 8);
					var html='';
					if(!msg.message){
						return;
					}else{
						//IE8适配foreach
						if ( !Array.prototype.forEach ) {
						    Array.prototype.forEach = function forEach( callback, thisArg ) {						
						    var T, k;						
						    if ( this == null ) {
						      throw new TypeError( "this is null or not defined" );
						    }
						    var O = Object(this);
						    var len = O.length >>> 0; 
						    if ( typeof callback !== "function" ) {
						      throw new TypeError( callback + " is not a function" );
						    }
						    if ( arguments.length > 1 ) {
						      T = thisArg;
						    }
						    k = 0;						
						    while( k < len ) {						
						      var kValue;
						      if ( k in O ) {						
						        kValue = O[ k ];
						        callback.call( T, kValue, k, O );
						      }
						      k++;
						    }
						  };
						}
						msg.state.forEach(function(e){
							console.log(e)
							var roomname = '';
							var roomstate = '';
							
							if(e.RoomType == 1){
								 roomname = "竞猜直播间";	
							}else{
								roomname = "普通直播间";
							}
							if(e.LiveState == 0){
								roomstate = "休息中";
							}else{
								roomstate = "直播中";
							}
							var str = '<a href="'+e.LiveAddress+'"><li class="anchorPart"><img class="anchorPic" src="'+e.LiveCover+'/" onerror="this.src=\'/static/images/thumb.jpg\'"><div class="anchorWord"><div class="anchor-Time">'+e.RoomAlias+'</div><div class="anchor-Name">'+e.NickName+'</div><div class="anchor-yes">'+roomname+'</div><div class="anchor-audience"><img class="anchor-sPic2" src="../static/images/mine/时间.png"></img>'+roomstate+'</div><div class="anchor-game"><img class="anchor-sPic1" src="../static/images/mine/好友.png"></img>'+e.RoomType+'</div></div></li></a>';
							html+=str;
						});
					}
					if(total>=8){
						// 
						$(".page-component").createPage({
							id:id,
					        pageCount:page,
					        current:pages,
					        backFn:function(p){
			           			 console.log(p);
				      		  }		
				    });
					}
					var hot_title = document.getElementById('anchorPart-ul');
					hot_title.innerHTML += html;
				}  
			});  
		}); 