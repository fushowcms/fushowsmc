var num;
var ip = returnCitySN["cip"];
$(function () {
    var money = getStorage("rechargeMoney");
    if (money == "" || money == 0 || money == null) {
        window.location = "recharge";
        return;
    }
    var txt = money / 100;
    $(".uNeedPay").text(txt);

    reqAjax("/user/erweima", { UID: getStorage("Id"), money: money, ip: ip }, function (msg) {
        if (msg.ErrorCode == 0) {
            $("#outTradeNo").val(msg.Data.outTradeNo);
            $('#qrcode').qrcode({
                text: msg.Data.state.Code_url,
                width: 264,
                height: 264,
                background: "#fff",
                foreground: "black"
            });
        }
    }, true);
    removeStorage("rechargeMoney");
    num = window.setInterval(function () { queryBill() }, 5000);
});
function queryBill() {
    var outTradeNo = $("#outTradeNo").val();

    reqAjax("/user/billquery", { bill: outTradeNo }, function (msg) {
        if (msg.ErrorCode == 4021) {
            Dialog(msg.ErrorMsg, true, "确定", null, function () {
                window.location.href = "recharge-Success";
            }, null);
        } else if (msg.ErrorCode == 4026) {
            Dialog(msg.ErrorMsg, true, "确定", null, function () {
                window.location.href = "recharge-Fail";
            }, null);
        }
    }, true);
}