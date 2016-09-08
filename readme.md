# 项目简介：
    此项目的设计目标主要用于需要多人协作、需要对操作权限做精确控制的管理系统使用。目前只提供了权限管理、用户组管理、
    用户管理基础功能，后续如果有可以公用的功能会陆续增加进来。
    本项目使用go语言开发
- web框架使用 [beego](http://beego.me/)
- 前台页面使用 [easyUI](http://www.jeasyui.net/demo/380.html) （easyUI 中文网）
- 页面中使用的树组件为  [zTree](http://www.treejs.cn/v3/main.php#_zTreeInfo)

## 软件环境
- 开发工具： VS Code
- 数据库： mysql

## 安装指南
- doc 目录下的database.sql  是数据库初始化脚本
- 默认账号 `admin` 密码 `111111`

## 功能介绍： 
    1：管理员管理 
        管理可以使用此系统的用户
        功能：查询、新增、修改（可以直接重置密码）删除、管联用户组（管理用户组后可以拥有改组的所有权限）
    2：管理员组管理
        主要是用于用户分组和权限分组
        功能：查询、新增、修改、删除、关联权限（设置这个组拥有那些权限）
    3：权限管理
        用于管理系统中所有的操作权限、导航菜单中的菜单目录
        功能：查询、新增、修改、删除
        ps：这里的权限有三种用途
            A：作为导航菜单中的目录使用（新增的时候只需要填写 权限名称 和 选择是否作为菜单）
            B：仅作为菜单（新增的时候必须填写 权限名称、请求地址、模块名称、操作名称、并选择作为菜单为 是）
            C：仅作为一个普通权限（新增的时候必须填写 权限名称、模块名称、操作名称、并选择作为菜单为 否）

## 功能截图
- 管理员管理
![image](https://github.com/crazy-wolf/cms/blob/master/doc/img/user.png)
![image](https://github.com/crazy-wolf/cms/blob/master/doc/img/adduser.png)
- 管理员组管理
![image](https://github.com/crazy-wolf/cms/blob/master/doc/img/usergroup.png)
![image](https://github.com/crazy-wolf/cms/blob/master/doc/img/addusergroup.png)
- 权限管理
![image](https://github.com/crazy-wolf/cms/blob/master/doc/img/role.png)

## 待完成功能
- [X] 管理员管理
- [X] 修改管理员组功能
- [X] 删除管理员组功能
- [X] 登陆校验
- [X] 权限校验
- [X] 添加权限后刷新左侧菜单
- [X] 权限页面打开后 在打开组管理  权限页面点击tree刷新表
- [X] 建库sql
- [X] 权限删除的时候判断当前权限是否有子节点，如果有子节点不能删除（2016-9-6）
- [X] 增加退出功能 （2016-9-6）
- [X] 添加修改参数校验
- [ ] 功能测试

`ps:系统中所有html标签的name不能相同，调用的js方法名也不能相同，
否则会引起功能混乱，easyui的限制由于它是把所有的js和html加载到一个页面中了`