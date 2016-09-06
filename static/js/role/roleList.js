$(function () {
    searRoleObj = {
        search: function () {
            $('#role_list').datagrid('load', {
                roleName: $('input[name="searchRoleName"]').val(),
                roleUrl: $('input[name="searchRoleUrl"]').val(),
                roleid: $("input[name='searchRolepid']").val()
            });
        }
    }
    //datagrid初始化
    $('#role_list').datagrid({
        url: 'role/gridlist',
        // queryParams: { roleid: 0 },
        iconCls: 'icon-edit',//图标
        width: 700,
        height: 'auto',
        nowrap: false,
        striped: true,
        border: true,
        collapsible: false,//是否可折叠的
        fit: true,//自动大小
        //sortName: 'code',
        //sortOrder: 'desc',
        remoteSort: false,
        idField: 'id',
        singleSelect: false,//是否单选
        pagination: true,//分页控件
        rownumbers: true,//行号
        fitColumns: true,//列宽自适应（列设置width=100）
        frozenColumns: [[
            { field: 'ck', checkbox: true }
        ]],//设置表单复选框
        toolbar: role_toolbar
    });
})

//添加修改按钮
function roleOpt(val, row, index) {
    return '<a href="#" onclick="openModifyRoleWin(' + row.id + ')">修改</a>';
}
//判断是否是菜单
function roleIsMenu(val, row, index) {
    if (row.ismenu == 0) {
        return "是"
    } else if (row.ismenu == 1) {
        return "否"
    }

}


//设置tree的初始化参数
var rolesetting = {
    data: {
        simpleData: {
            enable: true
        }
    }
    ,
    callback: {
        onClick: changeRoleList
    }
};

//初始化左边tree
$(document).ready(loadTree());
function loadTree(id) {
    url = "/role/listtree"
    // var zNodes = [{ id: 0, name: "Root", open: true }]
    var data = { id: id };
    $.post(url, data, function (result) {
        // zNodes = result
        $.fn.zTree.init($("#roletree"), rolesetting, result);
    });
}
// 点击tree节点的时候 重新加载右边的权限列表
function changeRoleList(event, treeId, treeNode) {
    loaddatagrid(treeNode.id)
    $('#searchRolepid').val(treeNode.id)
}

//加载表格
function loaddatagrid(id) {
    $('#role_list').datagrid('load', {
        roleid: id
    });
}

//打开添加权限目录窗口
function openAddRoleWin() {
    $('#addRole').window({
        width: 400,
        height: 300,
        modal: true,
        maximizable: false,
        minimizable: false,
        collapsible: false,//是否可折叠的
        href: "/role/toadd"
    });
}

//打开添加权限目录窗口
function openAddRoleDirWin() {
    $('#addRoleDir').window({
        width: 400,
        height: 300,
        modal: true,
        maximizable: false,
        minimizable: false,
        collapsible: false,//是否可折叠的
        href: "/role/toadddir"
    });
}

//打开修改权限窗口
function openModifyRoleWin(roleid) {
    $("#roleid").attr("value", roleid);
    $('#modifyrole').window({
        width: 400,
        height: 300,
        modal: true,
        maximizable: false,
        minimizable: false,
        collapsible: false,//是否可折叠的
        href: "/role/tomodify?roleid=" + roleid
    });
}

//删除方法
function deleteRole() {
    var selections = $('#role_list').datagrid('getSelections')
    if (selections.length == 0) {
        alert("请先选择要删除的列")
        return false
    }

    if (!confirm("确定要删除选中的数据吗？")) {
        return false
    }
    var idArray = new Array(selections.length)
    for (var i = 0; i < selections.length; i++) {
        idArray[i] = selections[i].id
    }
    ids = idArray.join(",")

    url = "/role/deleterole"
    var data = { ids: ids };

    var pid = $("input[name='searchRolepid']").val()
    $.post(url, data, function (result) {
        loadTree(pid)
        loaddatagrid(pid)
        selections.length = 0;
        if (result == "success") {
            $.messager.alert('操作提示', "删除成功", 'info');
        } else {
            $.messager.alert('操作提示', result, 'warning');
        }
    });
}