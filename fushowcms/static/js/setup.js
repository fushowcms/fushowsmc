$(function(){
    $('.setup').click(function(){
        var dbAddress = $('#dbAddress').val();
        var dbPort = $('#dbPort').val();
        var dbAccount = $('#dbAccount').val();
        var dbPassword = $('#dbPassword').val();
        var dbName = $('#dbName').val();
        var dbSelect = $('#dbSelect').val();
        var redisDial = $('#redisDial').val();
        var redisDeal = $('#redisDeal').val();
        var redisPass = $('#redisPass').val();
        var redisKey = $('#redisKey').val();
        var phone = $('#phone').val();
        if (dbAddress=="") {
            dialog("数据库地址不能为空!");
            return;
        }
        if (dbPort=="") {
            dialog("数据库端口不能为空!");
            return;
        }
        if (dbAccount=="") {
           dialog("数据库账号不能为空!");
            return;
        }
        if (dbPassword=="") {
            dialog("数据库密码不能为空!");
            return;
        }
        if (dbName=="") {
           dialog("数据库名称不能为空!");
            return;
        }
        if (dbSelect=="") {
            dialog("数据库类型不能为空!");
            return;
        }
        if (redisDial=="") {
            dialog("redisDial不能为空!");
            return;
        }
        if (redisDeal=="") {
           dialog("redisDeal不能为空!");
            return;
        }
        if (redisPass=="") {
            dialog("redisPass不能为空!");
            return;
        }
        if (redisKey=="") {
            dialog("redisKey不能为空!");
            return;
        }
        if (phone=="") {
            dialog("电话不能为空!");
            return;
        }
        reqAjax('/page/setUp',{DbHost:dbAddress,DbPort:dbPort,DbUser:dbAccount,DbPass:dbPassword,DbDbname:dbName,DbSelectdb:dbSelect,RedisDial:redisDial,RedisDeal:redisDeal,RedisPass:redisPass,RedisKey:redisKey,Phone:phone},function(data){
			if(data.ErrorCode==0){
                dialog("设置成功,请重启服务");
            }
		});
    })

    function dialog(msg) {
        Dialog(msg,true,"确定",null,function() {
				$('.dialog').remove();
            },null);
    }
})