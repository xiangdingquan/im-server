更新的文档在https://gitee.com/caocun/interface/blob/im/json.md，此文档不再维护

# [error_code](#error_code)
# [system](#system) 
- [synctime](#systemsynctime)
- [relayData](#systemrelaydata)
- [sendSmsCode](#systemsendsmscode)

# [wallet](#wallet)
- [setPassword](#walletsetpassword)
- [info](#walletinfo)
- [records](#walletrecords)

# [call](#call)
- [token.create](#calltokencreate)
- [create](#callcreate)
- [getInfo](#callgetinfo)
- [cancel](#callcancel)
- [start](#callstart)
- [stop](#callstop)
- [queryOffline](#callqueryoffline)
- [queryRecord](#callqueryrecord)

# [redpacket](#redpacket)
- [create](#redpacketcreate)
- [get](#redpacketget)
- [detail](#redpacketdetail)
- [record](#redpacketrecord)
- [statistics](#redpacketstatistics)

# [remittance](#remittance)
- [remit](#remittanceremit)
- [receive](#remittancereceive)
- [refund](#remittancerefund)
- [getRecord](#remittancegetrecord)
- [getRecords](#remittancegetrecords)
- [remind](#remittanceremind)

# [users](#users)
- [info](#usersinfo)
- [searchByPhone](#userssearchbyphone)
- [queryPrivacySettings](#usersqueryprivacysettings)

# [accounts](#accounts)
- [toggleMultiOnline](#accountstogglemultionline)
- [getMultiOnline](#accountsgetmultionline)

# [channel](#channel)
- [cleanMessages](#channelcleanMessages)

# [chats](#chats)
- [getBannedRightex](#chatsgetbannedrightex)
- [modifyBannedRightex](#chatsmodifybannedrightex)
- [getFilterKeywords](#chatsgetfilterkeywords)
- [setFilterKeywords](#chatssetfilterkeywords)
- [disableInviteLink](#chatsdisableinvitelink)

# [web](#web)
- [discover](#webdiscover)
- [recharge](#webrecharge)
- [withdraw](#webwithdraw)
- [app.getVersion](#webappgetversion)
- [system.getCSNumbers](#websystemgetcsnumbers)
- [system.getCustomMenus](#websystemgetcustommenus)
- [system.getAppConfig](#websystemgetappconfig)
- [account.myInfo](#webaccountmyinfo)
- [account.setPassword](#webaccountsetpassword)
- [account.setWalletPassword](#webaccountsetwalletpassword)
- [account.checkWalletPassword](#webaccountcheckwalletpassword)

# [moment](#moment)
- [post](#momentpost)

# error_code
错误码列表
```
501:自己为对方的黑名单
```

# system
## system.synctime
获取服务器UTC时间戳 同步服务器时间 计算时差(毫秒) 建议首次调用后一直保存.
```
<-{
    "client_time":Number //客户端UTC13位时间戳
}
->{
    "client_time":Number //客户端UTC13位时间戳
    "server_time":Number //服务器UTC13位时间戳
}
```

## system.relayData
转发数据
```
<-{
    "action":String //转发的事件名称
    "to":Array Number //接收者
    "data":Json //自定义携带数据
}
->{}
->on_event{
    "action":String //转发的处理名称
    "from":Number //提交的所有者
    "to":Array Number //接收者
    "data":Json //自定义携带数据
}
```

## system.sendSmsCode
发送验证码
```
<-{
    "type":Number //短信类型 1修改钱包密码
}
->{}
```

# wallet
## wallet.setPassword
创建钱包/修改钱包密码
```
<-{
    "smsCode":String //接收的验证码
    "password":String //新密码(MD5)
}
->{
    "address":String //钱包地址
}
400:验证码认证错误
```

## wallet.records
获取钱包记录
```
<-{
    "count":Number //每页数量
    "page":Number //获取第N页
}
->{
    "type":Number //类型 1.充值 2提现 3.收款 4.转账 5.创建红包 6.领取红包 7.红包退回
    "amount":Number //变动金额
    "remarks":String //备注
    "createAt":Number //创建时间
}
```

## wallet.info
查询钱包信息
```
<-{}
->{
    "address":String //钱包地址
    "balance":Number //钱包余额
    "hasPaymentPassword":Bool //是否设置支付密码
}
400:未开通钱包
```

# call
## call.token.create
生成声网RTCToken
```
<-{
    "channelName":String //通道名称
    "uid":Number //用户标识
}
->{
    "token":String //声网RTCToken
}
```

## call.create
创建通话
```
<-{
    "channelName":String //声网通道名称
    "to":Array{Number} //接收者
    "chatId":Number //IM会话标识
    "isMeetingAV":Bool //是否会议群聊
    "isVideo":Bool //是否开启视频
}
->{
    "callId":Number //通话记录标识
    "channelName":String //通道名称
    "from":Number //通话发起者/所有者
    "to":Array{Number} //接收者
    "chatId":Number //IM会话标识
    "isMeetingAV":Bool //是否会议群聊
    "isVideo":Bool //是否开启视频
    "createAt":Number //服务器10位时间戳
}
->call.onInvite{
    "callId":Number //通话记录标识
    "channelName":String //通道名称
    "from":Number //通话发起者/所有者
    "to":Array{Number} //接收者
    "chatId":Number //IM会话标识
    "isMeetingAV":Bool //是否会议群聊
    "isVideo":Bool //是否开启视频
    "createAt":Number //服务器10位时间戳
}
```

## call.getInfo
获取通话信息
```
<-{
    "callId":Number //通话记录标识
}
->{
    "callId":Number //通话记录标识
    "channelName":String //声网通道名称
    "from":Number //通话发起者/所有者
    "to":Array{Number} //接收者
    "chatId":Number //IM会话标识
    "isMeetingAV":Bool //是否会议群聊
    "isVideo":Bool //是否开启视频
    "isClose":Bool //是否已取消通话
    "createAt":Number //服务器10位时间戳
}
```

## call.cancel
取消通话只有 创建者/所有者可调用 其他无效
```
<-{
    "callId":Number //通话记录标识
}
->{}
->call.onCancel{
    "callId":Number //通话记录标识
    "channelName":String //声网通道名称
    "chatId":Number //IM会话标识
    "createAt":Number //创建时间10位时间戳
}
```

## call.start
加入/开始通话
```
<-{
    "callId":Number //通话记录标识
}
->{}
```

## call.stop
离开/停止通话
```
<-{
    "callId":Number //通话记录标识
}
->{}
->call.onLeave{
    "from":Number //通话发起者/所有者
    "to":Array{Number} //接收者
    "callId":Number //通话记录标识
    "channelName":String //声网通道名称
    "chatId":Number //IM会话标识
    "createAt":Number //创建时间10位时间戳
}
```

## call.queryOffline
* 获取离线记录
```
<-{}
->{
    "count":Number //返回记录数量
    "records":Array{
        "callId":Number //通话记录标识
        "channelName":String //声网通道名称
        "chatId":Number //IM会话标识
        "from":Number //创建者
        "to":Array{Number} //接收者 被邀请的成员
        "createAt":Number //创建时间10位时间戳
        "closeAt":Number //关闭时间10位时间戳
        "isMeetingAV":Bool //是否会议群聊
        "isVideo":Bool //是否开启视频
        "enterAt":Number //进入时间10位时间戳
        "leaveAt":Number //离开时间10位时间戳
    }
}
```

## call.queryRecord
查询自己所有通话记录
```
<-{
    "type":Number //0:单聊呼出和接听 1:单聊呼出 2:单聊接听 3:会议
    "count":Number //每页数量
    "page":Number //获取第N页
}
->{
    Array{
        "callId":Number //通话记录标识
        "channelName":String //声网通道名称
        "chatId":Number //IM会话标识
        "from":Number //创建者
        "to":Array{Number} //接收者 被邀请的成员
        "createAt":Number //创建时间10位时间戳
        "closeAt":Number //关闭时间10位时间戳
        "isMeetingAV":Bool //是否会议群聊
        "isVideo":Bool //是否开启视频
        "enterAt":Number //进入时间10位时间戳
        "leaveAt":Number //离开时间10位时间戳
    }
}
```

# redpacket
## redpacket.create
创建红包
```
<-{
    "chatId":Number //IM会话标识
    "type":Number //1.单聊红包 2.拼手气红包 3.普通红包
    "title":String //标题
    "price":Number //单个红包金额
    "total_price":Number //红包总金额
    "count": Number //红包数量
    "password":String //钱包密码(MD5)
}
->{}
->onMessage(2,100){
    "redPacketId":Number //红包id
    "from":Number //创建者
    "title":String //红包标题
}
400:余额少于红包总金额
401:红包数量大于1000
402:单个红包金额少于0.01
403:单个红包金额大于200
404:支付密码错误
```

## redpacket.get
领取红包
```
<-{
    "redPacketId":Number //红包id
}
->{
    users:Array{
        "userId":Number //用户id
        "price":Number //获取到的金额
        "gotAt":Number //领取的时间
    }
}
->onMessage(2,200){
    "redPacketId":Number //红包id
    "from":Number //创建者
    "type":Number //1.单聊红包 2.拼手气红包 3.普通红包
    "price":Number //获得金额
    "isLast":Bool //是否最后一个红包
}
400:不能领取自己的红包
401:没权限领取
402:红包已经领完
```

## redpacket.detail
红包详情
```
<-{
    "redPacketId":Number //红包id
}
->{
    "redPacketId":Number //红包id
    "from":Number //创建者
    "createAt":Number //创建时间
    "chatId":Number //IM会话标识
    "type":Number //1.单聊红包 2.拼手气红包 3.普通红包
    "title":String //红包标题
    "price":Number //单个红包金额
    "total_price":Number //红包总金额
    "count": Number //红包数量
    "isExpire":Bool //是否过期
    "users":Array{
        "userId":Number //用户id
        "price":Number //获取到的金额
        "gotAt":Number //领取时间
    }
}
```

## redpacket.record
红包记录
```
<-{
    "type":Number //1创建红包 2领取红包
    "count":Number //每页数量
    "page":Number //获取第N页
}
->{
    Array{
        "redPacketId":Number //红包id
        "from":Number //创建者
        "createAt":Number //创建时间
        "chatId":Number //IM会话标识
        "type":Number //1.单聊红包 2.拼手气红包 3.普通红包
        "title":String //红包标题
        "price":Number //单个红包金额
        "total_price":Number //红包总金额
        "count": Number //红包数量
        "isExpire":Bool //是否过期
        "users":Array{
            "userId":Number //用户id
            "price":Number //获取到的金额
            "gotAt":Number //领取时间
        }
    }
}
```

## redpacket.statistics
红包统计数据
```
<-{
    "type":Number //1创建红包 2领取红包
    "year":Number //年份，传0为查找所有
}
->{
    "total_price":Number //总金额
    "count":Number //个数
    "top_price_count":Number //手气最佳个数
}
```

# remittance
## remittance.remit
转账-付款
```
<-{
    "chatId":Number //IM会话标识
    "description":String //转账说明
    "amount":Number //转账金额
    "password":String //钱包密码(MD5)
    "type":Number //1.单聊转账 2.群聊转账
    "payee":Number //收款人
}
->{}
->onMessage(5,100){
    "remittanceId":Number //转账id
    "payer":Number //付款人
    "payee":Number //收款人
    "amount":Number //转账金额
}
400:余额不足
401:转账金额少于0.01
402:转账金额大于xxx
403:支付密码错误
```

## remittance.receive
转账-收款
```
<-{
    "remittanceId":Number //转账id
}
->{
}
->onMessage(5,200){
    "remittanceId":Number //转账id
    "payer":Number //付款人
    "payee":Number //收款人
    "amount":Number //转账金额
}
400:你不是收款人
```

## remittance.refund
转账-退款
```
<-{
    "remittanceId":Number //转账id
}
->{
}
->onMessage(5,300){
    "remittanceId":Number //转账id
    "payer":Number //付款人
    "payee":Number //收款人
    "amount":Number //转账金额
}
400:你不是收款人
```

## remittance.getRecord
查询转账
```
<-{
    "remittanceId":Number
}
->{
    {
        "type":Number //1.单聊转账 2.群聊转账
        "remittanceId":Number //转账id
        "payerUID":Number //付款人
        "payeeUID":Number //收款人
        "remittedAt":Number //创建时间
        "receivedAt":Number //收款时间
        "refundedAt":Number //退款时间
        "chatId":Number //IM会话标识
        "description":String //转账说明
        "amount":Number //转账金额
        "status":Number //状态
    }
}
```

## remittance.getRecords
查询转账
```
<-{
    "type":Number //1.支出 2.收入
    "from_id":Number //填0从头开始查，之后可以填拿到的最小的id
    "limit":Number
}
->{
    Array{
        "type":Number //1.单聊转账 2.群聊转账
        "remittanceId":Number //转账id
        "payerUID":Number //付款人
        "payeeUID":Number //收款人
        "remittedAt":Number //创建时间
        "receivedAt":Number //收款时间
        "refundedAt":Number //退款时间
        "chatId":Number //IM会话标识
        "description":String //转账说明
        "amount":Number //转账金额
        "status":Number //状态
    }
}
```

## remittance.remind
发送提醒消息
```
<-{
    "remittanceId":Number
}
->{
}
->onMessage(5,400){
    "remittanceId":Number //转账id
}
```



# users
## users.info
获取多个用户信息
```
<-{
    "uIds":Array{Number} //用户id列表
}
->{
    Array{
        "uId":Number //用户id
        "firstName":String //姓
        "lastName":String //名
        "userName":String //用户名
        "isInternal":Bool //是否为内部号
    }
}
```

## users.searchByPhone
根据手机号查找手机号但需要对方开放隐私权限
```
<-{
    "phone":String //电话号码
}
->{
    "uId":Number //用户id
}
```

## users.queryPrivacySettings
查询多个用户针对自己的设置
```
<-{
    "uIds":Array{Number} //用户id列表
}
->{
    Array{
        "uId":Number //用户id
        "isBlacklist":Bool //是否为黑名单
    }
}
```

# accounts
## accounts.toggleMultiOnline
设置多端登录开关
```
<-{
    "isOn":Bool //是否可以多端登录
}
->{}
```

## accounts.getMultiOnline
查询多端登录开关
```
<-{}
->{
    "isOn":Bool
}
```

# channel
## channel.cleanMessages
屏蔽频道信息
```
<-{
    "channelId":Number //频道id
    "messageIds":Array{Number} //消息id列表
}
->{}
```

# chats
## chats.getBannedRightex
获取群聊扩展权限
```
<-{
    "chatId":Number //群组id 普通群-id 超级群+id
}
->{
    "banWhisper":Bool //禁止私聊
    "banSendWebLink":Bool //禁止发送网页链接
    "banSendQRcode":Bool //禁止发送二维码
    "banSendKeyword":Bool //禁止发送关键字
    "banSendDmMention":Bool //禁止发送dm@
}
```

## chats.modifyBannedRightex
修改群聊扩展权限
```
<-{
    "chatId":Number //群组id 普通群-id 超级群+id
    "banWhisper":Bool //禁止私聊
    "banSendWebLink":Bool //禁止发送网页链接
    "banSendQRcode":Bool //禁止发送二维码
    "banSendKeyword":Bool //禁止发送关键字
    "banSendDmMention":Bool //禁止发送dm@
}
->{}
->chats.rights.onUpdate{
    "isChannel":Bool //是否超级群
    "chatId":Number //群组id
    "banWhisper":Bool //禁止私聊
    "banSendWebLink":Bool //禁止发送网页链接
    "banSendQRcode":Bool //禁止发送二维码
    "banSendKeyword":Bool //禁止发送关键字
    "banSendDmMention":Bool //禁止发送dm@
}
```

## chats.getFilterKeywords
获取群聊屏蔽关键字
```
<-{
    "chatId":Number //群组id 普通群-id 超级群+id
}
->{
    Array{String} //关键字
}
```

## chats.setFilterKeywords
设置群聊屏蔽关键字
```
<-{
    "chatId":Number //群组id 普通群-id 超级群+id
    "keywords":Array{String} //关键字
}
->{}
->chats.keywords.onUpdate{
    "isChannel":Bool //是否超级群
    "chatId":Number //群组id
    "keywords":Array{String} //关键字
}
```

## chats.disableInviteLink
停用群邀请链接
```
<-{
    "chatId":Number //群组id 普通群-id 超级群+id
}
->{}
```

# web
## web.discover
发现页信息
```
<-{}
->{
    Array{
        "title":String //分组标题
        "menus":Array{
            "title":String //菜单标题
            "icon":String //图标uri
            "url":String //http Url
        }
    }
}
```

## web.recharge
获取充值页面
```
<-{}
->{
    "csUserId":Number //客服uid
    "payUrl":String?Null //充值页面
}
```

## web.withdraw
获取提现页面
```
<-{
    "amount":Number //提现金额
    "password":String //钱包密码(MD5)
}
->{
    "csUserId":Number //客服uid
    "checkoutUrl":String?Null //提现页面
}
400:余额不足
401:钱包密码错误
402:钱包密码未设置
```

## web.app.getVersion
获取平台应用版本号
```
<-{
    "platform":String //所属平台固定"android"
    "versionCode":Number //当前应用版本序列号
}
->{
    "versionCode":Number //最新版本序列号
    "version":String //最新版本号
    "downloadUrl":String //下载地址
}
201:已是最新版本
```

## web.system.getCSNumbers
获取系统客服列表
```
<-{}
->{
    Array{Number} //客服uid
}
```

## web.system.getCustomMenus
自定义菜单
```
<-{}
->{
    "title":String //菜单标题
    "url":String //http Url
}
```

## web.system.getAppConfig
客户端相关配置
```
<-{}
->{
    "permitModifyUserName":Bool //是否可以修改用户名
    "onlyFriendChat":Bool //true 仅好友可以聊天
    "onlyWhiteAddFriend":Bool //true 仅能添加白名单里的好友
}
```

## web.account.myInfo
获取自己账号相关信息
```
<-{}
->{
    "hasPassword":Bool //是否有密码
}
```

## web.account.setPassword
设置登录密码
```
<-{
    'old_password':String, //MD5 旧密码
    'new_password':String  //MD5 新密码
}
->{}

200 成功
400 无效新密码
401 旧密码错误
```

## web.account.setWalletPassword
修改支付密码
```
<-{
    'old_password':String, //MD5 旧密码
    'new_password':String  //MD5 新密码
}
->{}

200 成功
400 无效新密码
401 旧密码错误
```

## web.account.checkWalletPassword
支付密码校验
```
<-{
    'password':String //MD5 支付密码
}
->{}

200 正确
400 无效密码
401 密码错误
```

# moment
## moment.post
发现页信息
```
<-{}
->{
    Array{
        "title":String //分组标题
        "menus":Array{
            "title":String //菜单标题
            "icon":String //图标uri
            "url":String //http Url
        }
    }
}
```