
//设置tree的初始化参数
var admusergroupsetting = {
    check: {
        enable: true
        // chkboxType: { "Y": "", "N": "" }
    },
    data: {
        simpleData: {
            enable: true
        }
    }
};




//初始化左边tree
$(document).ready(loadTree());
function loadTree() {
    url = "/admusergroup/loadtreewithoutroot"
    var data;
    $.post(url, data, function (result) {
        // zNodes = result
        $.fn.zTree.init($("#addadmgrouproletree"), admusergroupsetting, result);
    });
}



function submitAddAmdUserGroupForm() {

    var zTree = $.fn.zTree.getZTreeObj("addadmgrouproletree");
    nodes = zTree.getCheckedNodes(true);
    checkCount = nodes.length;
    //判断选中的节点数，如果没有选中节点则提示操作错误
    if (checkCount == 0) {
        $.messager.alert('操作提示', "请至少选择一个权限", 'info');
        return false;
    }
    //获取所有选中的节点ID
    var idArray = new Array(checkCount)
    for (var i = 0; i < nodes.length; i++) {
        idArray[i] = nodes[i].id
    }
    ids = idArray.join(",")

    url = "/admusergroup/addadmusergroup"

    var data = {
        ids: ids,
        groupname: $("input[name='admgroupusername']").val(),
        describe: $("input[name='admgroupuserdescribe']").val()
    };

    $.post(url, data, function (result) {
        if (result == "success") {
            $('#addadmusergroup').window("close")
            $.messager.alert('操作提示', "添加成功", 'info');
            loadAdmUserGroupDatagrid()
        }else{
             $.messager.alert('操作提示', result, 'info');
        }
    });
}

function clearAddAmdUserGroupForm() {
    $('#addamdusergroup').form('clear');
}