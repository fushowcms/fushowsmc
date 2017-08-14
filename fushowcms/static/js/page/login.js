var LogCaptcha;
$(function(){
    //点击登录按钮
    $('#head-login,.sidebar-login-btn,.header-login-btn').click(function(){
        $('#loginBox').show();
        $('.logregbox').removeClass('none');
        getrcode();
        $('.logins-alert-inner-tabs').eq(0).click();
    });
    //点击注册按钮
    $('#head-setup,.sidebar-register-btn,.header-register-btn').click(function(){
        $('#loginBox').show();
        $('.logregbox').removeClass('none');
        getrcode();
        $('.logins-alert-inner-tabs').eq(1).click();
    });
    //点击关闭按钮
    $('.logins-alert-inner-wrap-close').click(function(){
        $('#loginBox').hide();
        $('.logregbox').siblings().addClass('none');
        //$('.losebox').addClass('none');
    });
    //选项卡切换按钮
    $('.logins-alert-inner-tabs').click(function(){
        $(this).addClass('tabs-active').siblings().removeClass('tabs-active');
        $('.tabs-pannel').eq($('.logins-alert-inner-tabs').index(this)).removeClass('none').siblings('.tabs-pannel').addClass('none');
        getrcode();
    });
    //点击忘记密码
    $('.lose-password').click(function(){
        $('.logregbox').addClass('none');
        $('.losebox').removeClass('none');
        getrcode();
    });
    //点击图形验证码
    $('.code-img').click(function(){
        getrcode();
    });

    //注册下一步按钮
    $('.regists-form-submit').click(function() {
        regnewuser();
        
        
    });
    $('#reg-form-submit').click(function(){
        nicknames();
    });

    //忘记密码获取验证码
    $('.lose-form-verify-btn').click(function() {
        identifycode($(this));
    });
    //重置密码按钮
    $('.lose-form-submit').click(function() {
        resetingpass();
    });



    //点击登陆
    $('.logins-form-submit').click(function() {
        var data = {};
        data.mobile = $('#userName').val();
        data.password = $('#userPassword').val();
        data.captcha = LogCaptcha;
        data.keycode = $('#log-img-verify').val();
        if(!data.mobile){
            $('.login-font').html('×用户名不能为空');
            return;
        }
        if(!data.password){
            $('.login-font').html('×密码不能为空');
            return;
        }
        if(!data.keycode){
            $('.login-font').html('×请填写验证码');
            return;
        }
        if(!getStorage("Id")){
            reqAjax("/page/log",data,function(ret) {
                if(ret.ErrorCode==0) {
                    setStorage("Id", ret.Data.list.UID);
                    setStorage("nicheng",ret.Data.list.NickName);
                    setStorage("from","login");
                    setStorage("Favicon",ret.Data.list.Favicon);
                    setStorage("Integral",ret.Data.list.Integral);
                    window.location.reload();
                }else{
                    getrcode();
                    $('#log-img-verify').val('');
                    $('.logins-form-feedback').html(ret.ErrorMsg);
                    if(ret.ErrorCode==4010){
                        $('#userName').val('');
                        $('#userPassword').val('');
                    }else if(ret.ErrorCode==4015){
                        $('.logins-alert-login').css('display','none');
                        $('.logins-alert-register').css('display','none');
                        $('.logins-alert-lose').css('display','none');
                    }else{
                        $('#userPassword').val('');
                    }
                }
            });
        }
    });
    //手机注册获取验证码
    $('.logins-form-verify-btn').click(function(){
        var data = {};
        data.mobile = $("#phone-number").val();
        data.keycode = $('#register-img-verify').val();
        data.captcha = LogCaptcha;
        var reg_phone=/^1[34578]\d{9}$/;
        
        if(!data.mobile){
            $(".register-font").html('×手机号不能为空');
            return;
        }
        if(!reg_phone.test(data.mobile)){
            $(".register-font").html('×手机格式不正确');
            return;
        }
        var verifybtn = $(this);
        var count = 60;
        reqAjax("/page/verificationCode",data,function(ret) {
            if(ret.ErrorCode==0) {
                //发送成功
                var resend = setInterval(function(){
                    count--;
                    if (count > 0){
                        verifybtn.val(count+"秒后可重新获取").attr('disabled','disabled');
                        verifybtn.css({'cursor':'not-allowed'});		                   
                    }else {
                        clearInterval(resend);
                        verifybtn.val("获取验证码").removeAttr('disabled');
                        verifybtn.css({'cursor':'pointer'});
                    }
                }, 1000);
            }else{
                alert(ret.ErrorMsg);
                return;					
            }
        });
    });
    



    //三方登录按钮
    $("#wechat-login").click(function(){		
        location = "https://open.weixin.qq.com/connect/qrconnect?appid=wx39232dd6a07ff4b8&redirect_uri=http://tv.fushow.cn&response_type=code&scope=snsapi_login&state=weixin#wechat_redirect";
    });
    $("#qq-login").click(function(){
        location = "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=101382938&redirect_uri=http://tv.fushow.cn&state=qq";
    });
    $("#weibo-login").click(function(){
        location = "https://api.weibo.com/oauth2/authorize?client_id=1099460121&response_type=code&redirect_uri=http://tv.fushow.cn&state=weibo";	
    });
});

function getrcode(){
    $.ajax({
        cache: true,
        type: "get",
        url:"/getimagecode",
        async: false,
        error: function(request) {
            alert("图形验证码获取失败");
        },
        success: function(data) {
            LogCaptcha=data.Data.CaptchaId;
            $(".code-img").attr("src",data.Data.ImageURL)
        }
    });
}

/*忘记密码验证码*/
function identifycode(e) {
    var data = {};
    data.mobile = $("#lose-phone-number").val();
    data.keycode = $('#lose-img-verify').val();
    data.captcha = LogCaptcha;
    var reg_phone=/^1[34578]\d{9}$/;
    if(!data.mobile){
        $(".lose-font").html('×手机号不能为空')
        return;
    }else if(!reg_phone.test(data.mobile)){
        $(".lose-font").html('×手机格式不正确')
        return;
    }else{
        $(".lose-font").html('')
    }
    var verifybtns = e;
    var count = 60;

    reqAjax('/page/loseregval',data,function(ret){
        if(ret.ErrorCode!=0) {
            alert(ret.ErrorMsg);
            return;
        }else {
            var resend = setInterval(function(){
            count--;
                if (count > 0){
                    verifybtns.val(count+"秒后可重新获取").attr('disabled','disabled');
                    verifybtns.css({'cursor':'not-allowed'});		                   
                }else {
                    clearInterval(resend);
                    verifybtns.val("获取验证码").removeAttr('disabled');
                    verifybtns.css({'cursor':'pointer'});
                }
            }, 1000);
        }
    });
}


/*注册*/
function regnewuser() {
    var reg_pass=/^(?![0-9]+$)[0-9A-Za-z]{6,16}$/;
    var data = {};
    data.mobile = $('#phone-number').val();
    data.password = $('#passWord').val();
    data.code = $('#phone-verify').val();
    data.way = 1;
    if(data.password == ""){
        //密码不能为空
        alert("密码不能为空")
        return;
    }else if(!reg_pass.test(data.password)) {
        //密码格式不正确,应为6-9位的字母数字组合
        alert("密码格式不正确,应为6-9位的字母数字组合")
        return;
    }
    reqAjax('/page/reg',data,function(ret){
        console.log(ret);
        if(ret.ErrorCode==0){


            setStorage("Id", ret.Data.list.UID);
            setStorage("nicheng",ret.Data.list.NickName);
            setStorage("from","login");
            setStorage("Favicon",ret.Data.list.Favicon);
            setStorage("Integral",ret.Data.list.Integral);
            //window.location.reload();

            //完成注册后设置昵称
            $('.logregbox').addClass('none');
            $('.regsecbox').removeClass('none');
            

        }else{
            $('.logins-alert-inner-wrap-close').click();
            Dialog(ret.ErrorMsg);
        }
    });

}

/*重置密码*/
function resetingpass() {
    reg_pass=/^(?![0-9]+$)[0-9A-Za-z]{6,16}$/;
    var validate = $('#lose-phone-verify').val();
    var passnew = $('#passnew').val();
    var confirmpass = $('#confirmpass').val();
    if(!validate){
        $(".lose-font").html('×验证码不能为空')
        return false;
    }else if(!passnew){
        $(".lose-font").html('×密码不能为空')
        return false;
    }else if(passnew!=confirmpass) {
        $(".lose-font").html('×两次密码输入不一致')
        return false;
    }else if(!reg_pass.test(passnew)){
        $(".lose-font").html('×密码应为6-16位')
        return false;
    }else{
        $(".lose-font").html('')
    }
    reqAjax("/page/resetpass",{mobile:$("#lose-phone-number").val(),code:validate,password:confirmpass},function(ret) {
        if(ret.ErrorCode!=0) {
            alert(ret.ErrorMsg);
            return;
        }else{
            //修改成功
            window.location.reload();
        }
    });
}

function nicknames() {
    var nickname = $("#nickname").val();
    //nick_check = /^[\u4e00-\u9fff\w]{4,10}$/;	
    nick_check=/^.{3,10}$/;
        
    if(!nickname) {
        $('.nick-font').html('×昵称不能为空');
        return;
    }else if(!nick_check.test(nickname)) {
        $('.nick-font').html('×昵称格式不正确,昵称长度应为3~10位');
        return;
    }else {
        $('.nick-font').html('');
    }
    checknick(nickname);
}
/*判断昵称是否存在*/
function checknick(nk) {
    
    reqAjax("/page/checknick",{nickname:nk},function(ret) {
        if(ret.Data.state) {
            nkupdate(nk);
        }else {
            Dialog("昵称已存在");
        }
    });
}
/*填写昵称*/
function nkupdate(nk) {
    //var Id = getStorage("Id");
    var data = {};
    data.UID = getStorage("Id");
    data.nickname = nk;
    reqAjax("/page/nick",data,function(ret) {
        if(ret.ErrorCode!=0) {
            alert(ret.ErrorMsg);
            return;
        }else {
            window.location.reload();
        }
    });
}

function getPar(par){
    var local_url = document.location.href; 
    var get = local_url.indexOf(par +"=");
    if(get == -1){
        return "";   
    }   
    var get_par = local_url.slice(par.length + get + 1);
    var nextPar = get_par.indexOf("&");
    if(nextPar != -1){
        get_par = get_par.slice(0, nextPar);
    }
    return get_par;
}