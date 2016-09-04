//此方法 用于测试 不用每次进来都要点一次左侧菜单
$(document).ready(function(){
    // addTab('系统用户管理auto','admuser/list')
    // addTab('权限管理auto','role/list')
    // addTab('系统用户组管理auto','admusergroup/list')
    // addTab('添加系统用户组管理auto','admusergroup/toadd')
    // addTab('添加系统用户auto','/admuser/toaddadmuser')


});
//添加标签页的方法
function addTab(title, url) {
    if ($('#mainTab').tabs('exists', title)) {
        $('#mainTab').tabs('select', title);
    } else {
        //var content = '<iframe scrolling="auto" frameborder="0"  src="'+url+'" style="width:100%;height:100%;"></iframe>';
        $('#mainTab').tabs('add', {
            title: title,
            //content:content,
            href: url,
            closable: true
        });
    }
}

/**
 * 格式化时间
 */
function dataformatter(value, row, index) {
    if (value) return phpjs.date("Y-m-d H:i:s", phpjs.strtotime(value));
    return value;
}