$(function () {
  modifyAdmUserObj = {
        search: function () {
            $('#updateAdmUserGroup').datagrid('load', {
                groupName: $('input[name="modifyAdmUser_UserGroupName"]').val()         
            });
        }
    }
    //datagrid初始化
    $('#updateAdmUserGroup').datagrid({
        url: 'admuser/gridgrouplist',
        queryParams: { admUserId: $("input[name='admUserId']").val() },
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
        toolbar: modifyAdmUser_toolbar,
        onLoadSuccess:function(row){//当表格成功加载时执行               
                var rowData = row.rows;
                $.each(rowData,function(idx,val){//遍历JSON
                      if(val.check==true){
                        $("#updateAdmUserGroup").datagrid("selectRow", idx);//如果数据行为已选中则选中改行
                      }
                });              
            }
    });
})



function submitModifyAmdUserForm() {
    var selections = $('#updateAdmUserGroup').datagrid('getSelections')
    if (selections.length == 0) {
        $.messager.alert('操作提示', "请至少选择一个组", 'info');
        return false
    }

    var idArray = new Array(selections.length)
    for (var i = 0; i < selections.length; i++) {
        idArray[i] = selections[i].id
    }
    ids = idArray.join(",")

    url = "/admuser/modifyyadmuser"
    var data = {
        groupids: ids,
        userId:$("input[name='admUserId']").val(),
        account: $("input[name='modifyAdmUserAcout']").val(),
        name: $("input[name='modifyAdmUserName']").val(),
        phone: $("input[name='modifyAdmUserPhone']").val(),
        department: $("input[name='modifyAdmUserDepartment']").val(),
        password: $("input[name='modifyAdmUserPassword']").val(),
        mail: $("input[name='modifyAdmUserEmail']").val()
    };

    if (data.account.length < 1 || data.name.length < 1 || data.phone.length < 1 || data.department.length < 1 || data.mail.length < 1) {
        $.messager.alert('操作提示', "信息填写不完整,请补充后重新提交", 'info');
        return
    }


    $.post(url, data, function (result) {
        if (result == "success") {
            $('#modifyAdmUser').window("close")
            $.messager.alert('操作提示', "修改成功", 'info');
            loadModifyAdmUserGrid()
        } else {
            $.messager.alert('操作提示', result, 'info');
        }
    });
}


function clearModifyAmdUserForm() {
    $('#modifyAdmUser').form('clear');
}


function loadModifyAdmUserGrid() {
    $('#admUser_list').datagrid('load', {
    });
}