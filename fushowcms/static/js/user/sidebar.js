$(function () {
    var url = window.location.pathname;
    var loc = url.lastIndexOf("\/");
    var file = url.substring(loc + 1);
    if ($("#" + file + "").size() > 0) {
        $("#" + file + "").addClass("li_selected");

    }
    if ($("#" + file + "").hasClass("anchor-Inform")) {
        $(".anchor-Inform").show("slow");
    }
    //******************关于主播********************//
    $("#beAnchor").toggle(function () {
        var Uid = getStorage("Id");
        if (Uid == null) {
            Dialog("您还没有登录", true, "确定", null, function () {
                $('.dialog').remove();
            }, null);
        } else {
            reqAjax("/page/getuser", { UID: Uid }, function (msg) {
                if (msg.ErrorCode != 0) {
                    Dialog(msg.ErrorMsg, true, "确定", null, function () {
                        $('.dialog').remove();
                    }, null);
                } else {
                    if (msg.Data.Type != 1) {
                        Dialog("仅主播有该权限", true, "确定", null, function () {
                            $('.dialog').remove();
                        }, null);
                    } else {
                        $(".anchor-Inform").stop().show("slow");
                    }
                }
            }, true);
        }
    }, function () {
        $(".anchor-Inform").stop().hide("slow");
    })
});