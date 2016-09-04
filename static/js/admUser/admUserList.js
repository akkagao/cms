$(function () {
    searAdmUserObj = {
        search: function () {
            $('#admUser_list').datagrid('load', {
                admaccout: $('input[name="searchAdmUserAccount"]').val(),
                admusername: $('input[name="searchAdmUserName"]').val(),
                admusermail: $('input[name="searchAdmUserEmail"]').val(),
                admuserphone: $('input[name="searchAdmUserPhone"]').val(),
                admuserid: $('input[name="searchAdmUserId"]').val()
            });
        }
    }
    //datagrid初始化
    $('#admUser_list').datagrid({
        url: 'admuser/gridlist',

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
            {field: 'ck', checkbox: true}
        ]],//设置表单复选框
        toolbar: admUser_toolbar
    });
})

//添加修改按钮
function admUserOpt(val, row, index) {
    return '<a href="#" onclick="openModifyAdmUserWin(' + row.id + ')">修改</a>';
}

//打开添加管理员组窗口
function openAddAdmUserWin() {
    $('#addAdmUser').window({
        width: 800,
        height: 600,
        modal: true,
        // maximizable: false,
        minimizable: false,
        collapsible: false,//是否可折叠的
        href: "/admuser/toaddadmuser"
    });
}

//打开修改管理员组窗口
function openModifyAdmUserWin(admUserId) {
    $('#modifyAdmUser').window({
        width: 800,
        height: 600,
        modal: true,
        // maximizable: false,
        minimizable: false,
        collapsible: false,//是否可折叠的
        href: "/admuser/tomodifyadmuser?admUserId=" + admUserId
    });
}

deleteAdmUser


//删除方法
function deleteAdmUser() {
    var selections = $('#admUser_list').datagrid('getSelections')
    if (selections.length == 0) {
        alert("请先选择要删除的记录")
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

    url = "/admuser/delete"
    var data = { userids: ids };

    $.post(url, data, function (result) {
        if (result == "success") {
            loadModifyAdmUserGrid()
            $.messager.alert('操作提示', "删除成功", 'info');
            selections=0;
        } else {
            $.messager.alert('操作提示', result, 'warning');
            selections=0;
        }
    });
}

function loadModifyAdmUserGrid() {
    $('#admUser_list').datagrid('load', {
    });
}