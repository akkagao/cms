
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

//初始化tree
$(document).ready(loadTree());
function loadTree() {
    var admgroupuserid = $("input[name='admgroupuserid']").val()
    url = "/admusergroup/loadtreechecked"
    var data = {
        admgroupuserid: admgroupuserid,
    };
    $.post(url, data, function (result) {
        // zNodes = result
        $.fn.zTree.init($("#modifyadmgrouproletree"), admusergroupsetting, result);
    });
}

/**
 * 修改管理员组
 */
function submitModifyAmdUserGroupForm() {
    var zTree = $.fn.zTree.getZTreeObj("modifyadmgrouproletree");
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

    url = "/admusergroup/modifyadmusergroup"

    var data = {
        ids: ids,
        id: $("input[name='admgroupuserid']").val(),
        groupname: $("input[name='ag_m_name']").val(),
        describe: $("input[name='ag_m_describe']").val()
    };

    $.post(url, data, function (result) {
        if (result == "success") {
            $('#modifyadmusergroup').window("close")
            $.messager.alert('操作提示', "修改成功", 'info');
            loadAdmUserGroupDatagrid()
        } else {
            $.messager.alert('操作提示', result, 'info');
        }
    });
}

function clearModifyAmdUserGroupForm() {
    $('#modifyadmusergroup').form('clear');
}


