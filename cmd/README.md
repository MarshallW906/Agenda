# Agenda子命令设计文档 by ZM-J
- help
  - usage: 列出命令说明
  - args: command string (= "all")
  - notes: 参数默认为all，打印所有命令的说明。
- register
  - usage: 用户注册
  - args: username string, password string, email string, phone string
  - notes: username和password为必要项，email和phone为选填。
- login
  - usage: 用户登录
  - args: username string, password string
  - notes: 均为必要项。
- logout
  - usage: 用户登出
  - args: None
  - notes: None
- query
  - usage: 用户查询
  - args: username string (= ""), email string (= ""), password string(= "")
  - notes: 空串等价为通配符。
- delete
  - usage: 用户删除
  - args: None
  - notes: None
- createMeeting
  - usage: 创建会议
  - args: title string, start_time string, end_time string
  - notes: 均为必要项。end_time > start_time。
- addParticipant
  - usage: 增加会议参与者
  - args: title string, member string
  - notes: 均为必要项。
- removeParticipant
  - usage: 删除会议参与者
  - args: title string, member string
  - notes: 均为必要项。
- queryMeeting
  - usage: 查询会议
  - args: start_time string, end_time string
  - notes: end_time >= start_time，相等时查询某个时间点的会议（在该时间点开始或结束的会议也被计入）。
- cancelMeeting
  - usage: 取消会议
  - args: title string
  - notes: 均为必要项。
- quitMeeting
  - usage: 退出会议
  - args: title string
  - notes: 均为必要项。
- clearMeeting
  - usage: 清空会议
  - args: None
  - notes: None

**TODO：时间格式及精度（年月日or时分秒）待进一步明确**
时间格式可以用go官方的time包，可以直接将表示时间的字符串格式化为时间。