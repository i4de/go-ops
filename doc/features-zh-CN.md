# 详细功能描述


## 1. 节点

### 1.0 节点管理
 * [x] 连接节点
 * [x] 断开节点
 * [ ] 自动重连
 * [x] 节点列表
 * [ ] 给节点打标签

### 1.1 节点状态
 * [x] cpu
 * [x] 内存
 * [x] 硬盘使用情况
 * [x] 网卡信息

### 1.2 节点文件操作
 * [ ] 查看节点文件夹
 * [ ] 在节点上创建文件夹
 * [ ] 删除文件
 * [ ] 移动文件
 * [ ] 上传文件
 * [ ] 下载文件



## 2. 资源管理

### 2.1 脚本资源
 * [ ] 创建脚本
 * [ ] 删除脚本
 * [ ] 编辑脚本
 * [ ] 克隆（复制别人的）脚本
 * [ ] 公共脚本库  
 * [ ] 用户脚本库
   
### 2.2 文件资源
  * [ ] 创建文件空间
  * [ ] 修改文件空间
  * [ ] 删除文件空间
  * [ ] 文件上传
  * [ ] 文件下载
  * [ ] 文件锁
  * [ ] 文件删除 

### 2.3 巡检项
  * [ ] 创建巡检项
  * [ ] 删除巡检项
  * [ ] 编辑巡检项

### 2.4 巡检模板
  * [ ] 创建巡检模板
  * [ ] 删除巡检模板
  * [ ] 编辑巡检模板



### 2.5 插件资源
  * [ ] 添加插件资源
  * [ ] 编辑插件
  * [ ] 删除插件

### 2.6 agent资源管理
  * [ ] 增加默认agent
  * [ ] 更新默认agent
  * [ ] 默认agent-包管理
  * [ ] 删除agent资源





## 3. 脚本任务

### 3.1 脚本任务执行支持
 * [x] 指定特定用户运行(只支持linux)
 * [x] 脚本运行的工作目录
 * [x] 脚本运行的超时时间
 * [x] 脚本参数
 * [x] 自定义解释器
 * [x] 脚本输入(脚本运行过程中需要用户的输入)
 * [x] 中断脚本执行（脚本在运行过程中，可以中断取消）

### 3.2 脚本任务执行方式
 * [x] 命令下发 (执行简单的命令)
 * [x] 脚本内容下发 (执行复杂脚本)
 * [x] 执行本地脚本 (脚本存放在本地) 
 * [x] 执行URL脚本 (从URL下载脚本到本地执行)
 * [ ] 插件任务 (可以通过插件方式执行任务)


## 4. 文件分发任务
 * [x] 通过给定文件URL下载地址分发到指定节点（文件下载）
 * [ ] 通过文件流方式推送文件到指定节点（文件下载）


## 5.Agent管理

### 5.1 节点agent管理
 * [x] agent安装
 * [x] agent卸载
 * [x] agent启动 
 * [x] agent停止
 * [ ] 查看agent状态





## 6. 调用方（APP）管理
 * [x] 新增调用方
 * [x] 更新调用方
 * [x] 删除调用方
 * [x] 查询调用方

## 7. 用户管理
 * [ ] 新增用户
 * [ ] 更新用户
 * [ ] 删除用户
 * [ ] 用户角色

## 8. 权限控制
 * [ ] 调用方权限控制
 * [ ] 用户权限控制


## 9. 文件服务器
 * [ ] 支持本地存储,s3协议的文件服务器存储
 * [ ] 文件上传 (上传文件到文件服务器)
 * [ ] 文件锁 （被引用文件不可被删除,但是可以被更新）
 * [ ] 文件删除 (从文件服务器删除文件)
 * [ ] 获取文件列表信息 (获取文件服务器的列表信息)


## 10. 任务管理

### 10.1 创建、编辑任务
 * [ ] 创建、编辑脚本任务
 * [ ] 创建、编辑巡检任务
 * [ ] 创建、编辑文件分发任务
 * [ ] 创建、编辑定时任务
 * [ ] 创建、编辑预设任务

### 10.2 任务记录
 * [ ] 任务运行记录列表
 * [ ] 任务运行记录详情查看
 * [ ] 取消任务

### 10.3 预设任务
 * [ ] 预设任务列表
 * [ ] 预设任务详情
 * [ ] 预设任务删除
 * [ ] 运行预设任务

### 10.4 定时任务
 * [ ] 定时任务列表
 * [ ] 定时任务详情
 * [ ] 定时任务删除
 * [ ] 启动定时任务
 * [ ] 停止定时任务


## 11. 作业管理

### 11.1 批次作业
 * [ ] 任务分批次运行

### 11.2 作业编排
 * [ ] 任务通过作业编排方式运行

## 12. 远程连接

### 12.1 ssh连接
### 12.2 vnc连接