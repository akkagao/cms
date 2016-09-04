function submitAddRoleForm() {
    url = "/role/addrole"
    var pid = $("input[name='searchRolepid']").val()
    var data = {
        pid: pid,
        name: $("input[name='rolename']").val(),
        roleurl: $("input[name='roleurl']").val(),
        module: $("input[name='module']").val(),
        action: $("input[name='action']").val(),
        ismenu: $("#roleismenu").combobox('getValue'),
        describe: $("input[name='roledescribe']").val()
    };

    $.post(url, data, function (result) {
        if (result == "success") {
            clearAddRoleForm()
            loadTree(pid)
            loaddatagrid(pid)
            $.messager.alert('操作提示', "添加成功", 'info');
        }
    });
}

function clearAddRoleForm() {
    $('#addrole').form('clear');
    $("#roleismenu").combobox({value:"1"})
}