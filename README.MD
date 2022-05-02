# 基于gin和gorm的个人博客

### 如何启动
1. go build -o app
2. 在配置文件conf里修改你的数据库
3. `./app -migrate [数据库名字]`迁移数据
4. `./app -run` 启动服务


### 目录结构
1. api：restful Api
2. cmd：实现的命令行
    1. run命令：启动服务
    2. migrate命令：自动建表
3. common：各种工具类，常量
4. conf: 配置文件的位置
5. core: 功能模块包括：
    1. jwt刷新验证
    2. 验证码
    3. 上传文件
6. models: 实体类以及crud相关
7. routers：路由相关，各种中间件
8. static: 静态资源目录
9. ui：博客展示页面html
10. view：渲染博客展示页面

### 后端地址
- GitHub：https://github.com/SSjunLee/GolangBlog.git

### 管理系统ui地址
- 码云：https://gitee.com/ljn1997/my-golang-blog-ui.git


### TODU
1. 图床管理，把图片放到云盘
2. 配置管理
3. 优化CSS样式
4. 完善用户角色权限，动态加载管理系统菜单