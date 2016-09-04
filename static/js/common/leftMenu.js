
var leftMenuSetting = {
    data: {
        simpleData: {
            enable: true
        }
    }
    // ,
    // callback: {
    // 	onClick: zTreeOnClick
    // }
};



// 初始化左边菜单tree
$(document).ready(loadTree());
function loadTree() {
    url = "/loadMenu"
    var data;
    $.post(url, data, function (result) {
        // zNodes = result
        $.fn.zTree.init($("#leftMenuTree"), leftMenuSetting, result);
    });
}