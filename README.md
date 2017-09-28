# UPLOAD -- 图片上传服务器

# 配置文件
- Release: 是否是正式模式，1:正式模式, 0:测试模式
- BindPort: 绑定端口
- UploadDir: 中转图片的临时目录
- AllowFileType: 允许上传的图片类型
- MaxFileSize: 允许上传的最大文件大小
- FTPServer: 开启传到FTP服务器模式
- FTPServerAddr: FTP服务器地址
- FTPUploadPath: FTP上传路径
- FTPUsername: FTP登陆账号
- FTPPassword: FTP登陆密码
- HTTPServerURL: FTP服务器对应的HTTP访问地址

# 启动方式
- 默认启动
> ./upload
- 指定配置文件启动
> ./upload -conf=config.json

# 接口测试
- Release为0时开启测试模式，打开首页出现接口测试界面

# 接口接入
- 接口：/upload<br/>
- 类型：POST enctype=“multipart/form-data”<br/>

PARAM   |TYPE   |INTRO
--------|-------|-------
file    |file   |图片文件
module  |string |模块名，avatar:头像, share:分享
openid  |string |用户唯一ID，如OpenID或UID
unique	|int0/1 |是否是用户唯一资源：如：头像类unique传1，分享类非唯一资源传0

- 成功返回
> {"status":0,"data":"{HOST}/upload/bg/20170107_bbbbb_1483785506667.jpg"}
- 失败返回
> {"status":-1,"message":"get file error"}
