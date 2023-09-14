package jpush

import "encoding/json"

// Message represents jpush request message
type (
	Payload struct {
		//必填 推送平台设置 推送到所有平台 "all"
		//指定特定推送平台 ["android", "ios", "quickapp","winphone"]
		Platform interface{} `json:"platform"`

		//必填 推送目标 如果要发广播（全部设备），则直接填写 “all” Audience
		Audience interface{} `json:"audience"`

		//可选 通知内容体。是被推送到客户端的内容。与 message 一起二者必须有其一，可以二者并存
		//“通知”对象，是一条推送的实体内容对象之一（另一个是“消息”），是会作为“通知”推送到客户端的。
		//其下属属性包含 4 种，3 个平台属性，以及一个 "alert" 属性。
		Notification Notification `json:"notification,omitempty"`

		//自定义消息 应用内消息。或者称作：透传消息。
		//可选 消息内容体。是被推送到客户端的内容。与 notification 一起二者必须有其一，可以二者并存
		//此部分内容不会展示到通知栏上，JPush SDK 收到消息内容后透传给 App。需要 App 自行处理。
		//iOS 在推送应用内消息通道（非 APNS）获取此部分内容，即需 App 处于前台。Windows Phone 暂时不支持应用内消息。
		//Message Message `json:"message,omitempty"`

		//此功能生效需Android push SDK≥V3.9.0、iOS push SDK≥V3.4.0，若低于此版本按照原流程执行
		//InappMessage InappMessage `json:"inapp_message,omitempty"`

		//可选 短信渠道补充送达内容体
		//用于设置短信推送内容以及短信发送的延迟时间。
		//SmsMessage SmsMessage `json:"sms_message,omitempty"`

		//自定义消息转厂商通知
		//Push API 发起自定义消息类型的推送请求时，针对 Android 设备，如果 APP 长连接不在线，则消息没法及时的下发；
		//针对这种情况，极光推出了“自定义消息转厂商通知”的功能。
		//也就是说，针对用户一些重要的自定义消息，可以申请开通极光 VIP 厂商通道功能，
		//开通后，通过 APP 长连接不在线时没法及时下发的消息，可以通过厂商通道下发以厂商通知形式展示，及时提醒到用户。
		//极光内部会有去重处理，您不用担心消息重复下发问题。
		//notification_3rd 只针对开通了厂商通道的用户生效；
		//notification 和 notification_3rd 不能同时有内容，如果这两块同时有内容，则会返回错误提示；
		//notification_3rd 的内容对 iOS 和 WinPhone 平台无效，只针对 Android 平台生效；
		//notification_3rd 是用作补发厂商通知的内容，只有当 message 部分有内容，才允许传递此字段，且要两者都不为空时，才会对离线的厂商设备转发厂商通道的通知。
		//Notification3rd Notification3rd `json:"notification_3rd,omitempty"`

		// 可选 推送参数
		//Options Options `json:"options,omitempty"`

		// 可选 回调参数
		//CallBack CallBack `json:"callback,omitempty"`

		//推送唯一标识符
		//可选 用于防止 api 调用端重试造成服务端的重复推送而定义的一个标识符。
		//CId string `json:"cid,omitempty"`
	}

	Notification3rd struct {
		//可选 补发通知标题，如果为空则默认为应用名称
		Title string `json:"title,omitempty"`

		//必填 补发通知的内容，如果存在 notification_3rd 这个key，content 字段不能为空，且值不能为空字符串。
		Content string `json:"content,omitempty"`

		//可选 不超过1000字节，Android 8.0开始可以进行 NotificationChannel配置，这里根据channel ID 来指定通知栏展示效果
		ChannelId string `json:"channel_id,omitempty"`

		//可选 该字段用于指定开发者想要打开的 activity，值为 activity 节点的 “android:name”属性值;适配华为、小米、vivo厂商通道跳转；针对 VIP 厂商通道用户使用生效。
		UriActivity string `json:"uri_activity,omitempty"`

		//可选 指定跳转页面；该字段用于指定开发者想要打开的 activity，值为 "activity"-"intent-filter"-"action" 节点的 "android:name" 属性值;适配 oppo、fcm跳转；针对 VIP 厂商通道用户使用生效。
		UriAction string `json:"uri_action,omitempty"`

		//可选 角标数字，取值范围1-99；此属性目前仅针对华为 EMUI 8.0 及以上、小米 MIUI 6 及以上设备生效；
		//此字段如果不填，表示不改变角标数字（小米设备由于系统控制，不论推送走极光通道下发还是厂商通道下发，即使不传递依旧是默认+1的效果。）；
		//否则下一条通知栏消息配置的badge_add_num数据会和之前角标数量进行增加； 建议badge_add_num配置为1；
		//举例：badge_add_num配置1，应用之前角标数为2，发送此角标消息后，应用角标数显示为3。
		BadgeAddNum string `json:"badge_add_num,omitempty"`

		//可选 桌面图标对应的应用入口Activity类， 比如“com.test.badge.MainActivity；
		//配合badge_add_num使用，二者需要共存，缺少其一不可；
		//针对华为设备推送时生效（此值如果填写非主Activity类，走厂商推送以华为厂商限制逻辑为准；走极光通道下发，默认以APP的主Activity为准）
		BadgeClass string `json:"badge_class,omitempty"`

		//可选 填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀；注意：针对Android 8.0以上，当传递了channel_id 时，此属性不生效。
		Sound string `json:"sound,omitempty"`

		//JSON Object 可选 扩展字段；这里自定义 JSON 格式的 Key / Value 信息，以供业务使用。
		Extras interface{} `json:"extras,omitempty"`
	}

	Audience struct {
		// [标签OR] 数组。多个标签之间是 OR 的关系，即取并集。
		Tag []string `json:"tag,omitempty"`

		// [标签AND] 数组。多个标签之间是 AND 关系，即取交集。
		TagAnd []string `json:"tag_and,omitempty"`

		// [标签NOT] 数组。多个标签之间，先取多标签的并集，再对该结果取补集。
		TagNot []string `json:"tag_not,omitempty"`

		// [别名] 数组。多个别名之间是 OR 关系，即取并集。
		Alias []string `json:"alias,omitempty"`

		// [注册ID] 数组。多个注册 ID 之间是 OR 关系，即取并集。
		RegistrationId []string `json:"registration_id,omitempty"`

		// [用户分群 ID] 在页面创建的用户分群的 ID。定义为数组，但目前限制一次只能推送一个。
		Segment []string `json:"segment,omitempty"`

		// [A/B Test ID] 在页面创建的 A/B 测试的 ID。定义为数组，但目前限制是一次只能推送一个。
		Abtest []string `json:"abtest,omitempty"`
	}

	Notification struct {
		//类型为bool类型
		//取值为true表示推送采用智能时机策略推送，取值为false表示采用默认规则推送，字段默认值是false
		AiOpportunity bool `json:"ai_opportunity,omitempty"`

		//这个位置的 "alert" 属性（直接在 notification 对象下），是一个快捷定义，各平台的 alert 信息如果都一样，则可不定义。如果各平台有定义，则覆盖这里的定义
		//Alert string `json:"alert,omitempty"`

		//Android 平台上的通知，JPush SDK 按照一定的通知栏样式展示。
		Android Android `json:"android,omitempty"`

		//iOS 平台上 APNs 通知结构。
		//该通知内容会由 JPush 代理发往 Apple APNs 服务器，并在 iOS 设备上在系统通知的方式呈现。
		//Ios Ios `json:"ios,omitempty"`

		//快应用平台上通知结构。
		//QuickApp QuickApp `json:"quickapp,omitempty"`

		//Windows Phone 平台上的通知。
		//WinPhone WinPhone `json:"winphone,omitempty"`

		//OS VOIP功能。
		//该类型推送支持和 iOS 的 Notification 通知并存
		//任意自定义key/value对，会透传给APP
		//Voip map[string]string `json:"voip,omitempty"`
	}

	Android struct {
		//必填 通知内容 这里指定了，则会覆盖上级统一指定的 alert 信息；内容可以为空字符串，则表示不展示到通知栏。各推送通道对此字段的限制详见推送限制
		Alert string `json:"alert"`

		//可选 通知标题 如果指定了，则通知里原来展示 App 名称的地方，将展示成这个字段。各推送通道对此字段的限制详见推送限制
		Title string `json:"title,omitempty"`

		//可选 通知栏样式 ID Android SDK 可设置通知栏样式，这里根据样式 ID 来指定该使用哪套样式，android 8.0 开始建议采用NotificationChannel配置。
		BuilderId int `json:"builder_id,omitempty"`

		// 可选 android通知channel_id 不超过1000字节，Android 8.0开始可以进行NotificationChannel配置，这里根据channel ID 来指定通知栏展示效果。
		ChannelId string `json:"channel_id,omitempty"`

		//可选 通知栏展示优先级 默认为 0，范围为 -2～2。
		Priority int `json:"priority,omitempty"`

		//可选 通知栏条目过滤或排序 完全依赖 rom 厂商对 category 的处理策略
		Category string `json:"category,omitempty"`

		//可选 通知栏样式类型 默认为 0，还有 1，2，3 可选，用来指定选择哪种通知栏样式，其他值无效。有三种可选分别为 bigText=1，Inbox=2，bigPicture=3。
		Style int `json:"style,omitempty"`

		//可选 通知提醒方式 可选范围为 -1～7 ，对应 Notification.DEFAULT_ALL = -1 或者 Notification.DEFAULT_SOUND = 1, Notification.DEFAULT_VIBRATE = 2, Notification.DEFAULT_LIGHTS = 4 的任意 “or” 组合。默认按照 -1 处理。
		AlertType int `json:"alert_type,omitempty"`

		//可选 大文本通知栏样式 当 style = 1 时可用，内容会被通知栏以大文本的形式展示出来,厂商没有填充big_text,则也默认使用该big_text字段展示。厂商big_text, 支持 api 16 以上的 rom。
		BigText string `json:"big_text,omitempty"`

		//可选 文本条目通知栏样式 当 style = 2 时可用， json 的每个 key 对应的 value 会被当作文本条目逐条展示,厂商没有填充inbox，则也默认使用该inbox字段展示。厂商inbox，支持 api 16 以上的 rom。
		Inbox interface{} `json:"inbox,omitempty"`

		//可选 大图片通知栏样式 当 style = 3 时可用，可以是网络图片 url，或本地图片的 path，目前支持 .jpg 和 .png 后缀的图片,也可以是极光media_id。图片内容会被通知栏以大图片的形式展示出来。如果是 http／https 的 url，会自动下载；如果要指定开发者准备的本地图片就填 sdcard 的相对路径。厂商big_pic_path，支持 api 16 以上的 rom。
		BigPicPath string `json:"big_pic_path,omitempty"`

		//可选 扩展字段 这里自定义 JSON 格式的 Key / Value 信息，以供业务使用。针对部分厂商跳转地址异常，可通过 third_url_encode 兼容处理 "extras": { "third_url_encode": true //notification - android - extras ，true表示需要极光encode处理，值需要是布尔类型 } "extras": { "third_url_encode": false //notification - android - extras ，false，或者无此字段，表示不需要极光encode处理，值需要是布尔类型 }
		Extras interface{} `json:"extras,omitempty"`

		//可选 通知栏大图标 图标路径可以是以http或https开头的网络图片，如：http:jiguang.cn/logo.png ,图标大小不超过 30 k（注：从JPush Android SDK v4.0.0版本开始，图片大小限制提升至 300 k）;
		//也可以是位于drawable资源文件夹的图标路径，如：R.drawable.lg_icon；
		//也可以是通过极光图片上传接口得到的media_id值，如：jgmedia-2-14b23451-0001-41ce-89d9-987b465122da。
		//若是极光media_id,则对其它厂商通道也会使用这个media_id下发，若非media_id，则对走华硕通道下发和极光自有通道下发生效，不影响请求走其它厂商通道。厂商large_icon
		LargeIcon string `json:"large_icon,omitempty"`

		//可选 通知栏小图标 图标路径可以是以http或https开头的网络图片,如：http:jiguang.cn/logo.png,图标大小不超过 30 k （注：从JPush Android SDK v4.0.0版本开始，图片大小限制提升至 300 k）;
		//也可以是通过极光图片上传接口得到的media_id值，如：jgmedia-2-14b23451-0001-41ce-89d9-987b465122da。
		//此字段值，若是极光media_id,则对其它厂商通道也会使用这个media_id下发，若非media_id，则对走华硕通道下发和极光自有通道下发生效，不影响请求走其它厂商通道。厂商small_icon_uri
		SmallIconUri string `json:"small_icon_uri,omitempty"`

		//可选 指定跳转页面 使用 intent 里的 url 指定点击通知栏后跳转的目标页面;
		//此字段值，则仅对走华硕通道和极光自有通道下发生效，不影响请求走其它厂商通道。
		//SDK≥420的版本，API推送时建议填写intent字段（intent:#Intent;component=您的包名/Activity全名;end），否则点击通知可能无跳转动作。若使用onNotifyMessageOpened方法实现通知跳转，也可intent字段为空
		Intent interface{} `json:"intent,omitempty"`

		//可选 指定跳转页面 该字段用于指定开发者想要打开的 activity，值为 activity 节点的 “android:name”属性值;
		//适配华为、小米、vivo厂商通道跳转；
		//Jpush SDK≥V4.2.0，可不再填写本字段，仅设置intent字段即可
		UriActivity string `json:"uri_activity,omitempty"`

		//可选 指定跳转页面 该字段用于指定开发者想要打开的 activity，值为 "activity"-"intent-filter"-"action" 节点的 "android:name" 属性值;
		//适配 oppo、fcm跳转；
		//Jpush SDK≥V4.2.0，可不再填写本字段，仅设置intent字段即可，但若需兼容旧版SDK必须填写该字段
		UriAction string `json:"uri_action,omitempty"`

		//可选 角标数字，取值范围1-99 此属性目前仅针对华为 EMUI 8.0 及以上、小米 MIUI 6 及以上设备生效； 此字段如果不填，表示不改变角标数字（小米设备由于系统控制，不论推送走极光通道下发还是厂商通道下发，即使不传递依旧是默认+1的效果。）； 否则下一条通知栏消息配置的badge_add_num数据会和之前角标数量进行增加； 建议badge_add_num配置为1； 举例：badge_add_num配置1，应用之前角标数为2，发送此角标消息后，应用角标数显示为3。
		BadgeAddNum int `json:"badge_add_num,omitempty"`

		//可选 桌面图标对应的应用入口Activity类， 比如“com.test.badge.MainActivity” 配合badge_add_num使用，二者需要共存，缺少其一不可； 针对华为设备推送时生效（此值如果填写非主Activity类，走厂商推送以华为厂商限制逻辑为准；走极光通道下发，默认以APP的主Activity为准）
		BadgeClass string `json:"badge_class,omitempty"`

		//可选 填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀 注意：针对Android 8.0以上，当传递了channel_id 时，此属性不生效。
		Sound string `json:"sound,omitempty"`

		//可选 定时展示开始时间（yyyy-MM-dd HH:mm:ss） 此属性不填写，SDK默认立即展示；此属性填写，则以填写时间点为准才开始展示。
		//JPush Android SDK v3.5.0 版本开始支持。
		ShowBeginTime string `json:"show_begin_time,omitempty"`

		//可选 定时展示结束时间（yyyy-MM-dd HH:mm:ss） 此属性不填写，SDK会一直展示；此属性填写，则以填写时间点为准，到达时间点后取消展示。
		//JPush Android SDK v3.5.0 版本开始支持。
		ShowEndTime string `json:"show_end_time,omitempty"`

		//可选 APP在前台，通知是否展示 值为 "1" 时，APP 在前台会弹出通知栏消息；
		//值为 "0" 时，APP 在前台不会弹出通知栏消息。
		//注：默认情况下 APP 在前台会弹出通知栏消息。
		//JPush Android SDK v3.5.8 版本开始支持。
		DisplayForeground string `json:"display_foreground,omitempty"`
	}

	Ios struct {
		//string或JSON Object 必填 通知内容 这里指定内容将会覆盖上级统一指定的 alert 信息；内容为空则不展示到通知栏。支持字符串形式也支持官方定义的 alert payload 结构，在该结构中包含 title 和 subtitle 等官方支持的 key
		Alert interface{} `json:"alert"`

		//string 或 JSON Object 可选 通知提示声音或警告通知 普通通知： string类型，如果无此字段，则此消息无声音提示；有此字段，如果找到了指定的声音就播放该声音，否则播放默认声音，如果此字段为空字符串，iOS 7 为默认声音，iOS 8 及以上系统为无声音。说明：JPush 官方 SDK 会默认填充声音字段，提供另外的方法关闭声音，详情查看各 SDK 的源码。
		//告警通知： JSON Object ,支持官方定义的 payload 结构，在该结构中包含 critical 、name 和 volume 等官方支持的 key .
		Sound interface{} `json:"sound,omitempty"`

		//可选 应用角标 如果不填，表示不改变角标数字，否则把角标数字改为指定的数字；为 0 表示清除。JPush 官方 SDK 会默认填充 badge 值为 "+1",详情参考：badge +1
		Badge int `json:"badge,omitempty"`

		//可选 推送唤醒 推送的时候携带 "content-available":true 说明是 Background Remote Notification，如果不携带此字段则是普通的 Remote Notification。详情参考：Background Remote Notification
		ContentAvailable bool `json:"content-available,omitempty"`

		//可选 通知扩展 推送的时候携带 ”mutable-content":true 说明是支持iOS10的UNNotificationServiceExtension，如果不携带此字段则是普通的 Remote Notification。详情参考：UNNotificationServiceExtension
		MutableContent bool `json:"mutable-content,omitempty"`

		//可选  IOS 8 才支持。设置 APNs payload 中的 "category" 字段值
		Category string `json:"category,omitempty"`

		//JSON Object 可选 附加字段 这里自定义 Key / value 信息，以供业务使用。
		Extras interface{} `json:"extras,omitempty"`

		//可选 通知分组 ios 的远程通知通过该属性来对通知进行分组，同一个 thread-id 的通知归为一组。
		ThreadId string `json:"thread-id,omitempty"`
	}

	QuickApp struct {
		//必填 通知标题 必填字段，快应用推送通知的标题
		Title string `json:"title"`

		//必填 通知内容 这里指定了，则会覆盖上级统一指定的 alert 信息。
		Alert string `json:"alert"`

		//必填 跳转页面 快应用通 知跳转地址。
		Page string `json:"page"`

		//可选 JSON Object 附加字段 这里自定义 Key / value 信息，以供业务使用。
		Extras interface{} `json:"extras,omitempty"`
	}

	WinPhone struct {
		//必填 通知内容 会填充到 toast 类型 text2 字段上。这里指定了，将会覆盖上级统一指定的 alert 信息；内容为空则不展示到通知栏。
		Alert string `json:"alert"`
		//可选 通知标题 会填充到 toast 类型 text1 字段上。
		Title string `json:"title"`
		//可选 点击打开的页面名称 点击打开的页面。会填充到推送信息的 param 字段上，表示由哪个 App 页面打开该通知。可不填，则由默认的首页打开。
		OpenPage string `json:"_open_page,omitempty"`
		//JSON Object 可选 扩展字段 作为参数附加到上述打开页面的后边。
		Extras interface{} `json:"extras,omitempty"`
	}

	Message struct {
		// 必填 消息内容本身
		MsgContent string `json:"msg_content"`

		//可选 消息标题
		Title string `json:"title,omitempty"`

		//可选 消息内容类型
		ContentType string `json:"content_type,omitempty"`

		//JSON Object 可选 JSON 格式的可选参数
		Extras interface{} `json:"extras,omitempty"`
	}

	InappMessage struct {
		//inapp_message： 面向通知栏消息，Boolean类型；
		//值为 true 表示启用应用内提醒功能；
		//值为 false 表示禁用应用内提醒功能
		InappMessage bool `json:"inapp_message,omitempty"`
	}

	SmsMessage struct {
		//必填 单位为秒，不能超过 24 小时。设置为 0，表示立即发送短信。该参数仅对 android 和 iOS 平台有效，Winphone 平台则会立即发送短信。
		DelayTime int `json:"delay_time"`

		//选填 签名ID，该字段为空则使用应用默认签名。
		Signid int `json:"signid,omitempty"`

		//必填 短信补充的内容模板 ID。没有填写该字段即表示不使用短信补充功能。
		TempId int `json:"temp_id"`

		//json 可选 短信模板中的参数。
		TempPara string `json:"temp_para,omitempty"`

		//可选 active_filter 字段用来控制是否对补发短信的用户进行活跃过滤，默认为 true ，做活跃过滤；为 false，则不做活跃过滤；
		ActiveFilter bool `json:"active_filter,omitempty"`
	}

	Options struct {
		//可选 推送序号 纯粹用来作为 API 调用标识，API 返回时被原样返回，以方便 API 调用方匹配请求与返回。值为 0 表示该 messageid 无 sendno，所以字段取值范围为非 0 的 int.
		SendNo int `json:"sendno,omitempty"`

		//可选 离线消息保留时长(秒) 推送当前用户不在线时，为该用户保留多长时间的离线消息，以便其上线时再次推送。默认 86400 （1 天），最长 10 天。设置为 0 表示不保留离线消息，只有推送当前在线的用户可以收到。该字段对 iOS 的 Notification 消息无效。
		TimeToLive int `json:"time_to_live,omitempty"`

		//可选 要覆盖的消息 ID 如果当前的推送要覆盖之前的一条推送，这里填写前一条推送的 msg_id 就会产生覆盖效果，即：1）该 msg_id 离线收到的消息是覆盖后的内容；2）即使该 msg_id Android 端用户已经收到，如果通知栏还未清除，则新的消息内容会覆盖之前这条通知；覆盖功能起作用的时限是：1 天。如果在覆盖指定时限内该 msg_id 不存在，则返回 1003 错误，提示不是一次有效的消息覆盖操作，当前的消息不会被推送；该字段仅对 Android 有效。
		OverrideMsgId int `json:"override_msg_id,omitempty"`

		//可选 APNs 是否生产环境 True 表示推送生产环境，False 表示要推送开发环境；如果不指定则为推送生产环境。但注意，JPush 服务端 SDK 默认设置为推送 “开发环境”。该字段仅对 iOS 的 Notification 有效。
		ApnsProduction bool `json:"apns_production,omitempty"`

		//可选 更新 iOS 通知的标识符 APNs 新通知如果匹配到当前通知中心有相同 apns-collapse-id 字段的通知，则会用新通知内容来更新它，并使其置于通知中心首位。collapse id 长度不可超过 64 bytes。
		ApnsCollapseId string `json:"apns_collapse_id,omitempty"`

		//可选 定速推送时长(分钟) 又名缓慢推送，把原本尽可能快的推送速度，降低下来，给定的 n 分钟内，均匀地向这次推送的目标用户推送；最大值为 1400；最多能同时存在 20 条定速推送；未设置则不是定速推送。
		BigPushDuration int `json:"big_push_duration,omitempty"`

		//JSON Object 可选 推送请求下发通道 仅针对配置了厂商用户使用有效
		ThirdPartyChannel interface{} `json:"third_party_channel,omitempty"`
	}

	CallBack struct {
		//可选 数据临时回调地址，指定后以此处指定为准，仅针对这一次推送请求生效；不指定，则以极光后台配置为准
		Url string `json:"url,omitempty"`

		//JSON Object 可选 需要回调给用户的自定义参数
		Params interface{} `json:"params,omitempty"`

		//可选 回调数据类型，1:送达回执, 2:点击回执, 3:送达和点击回执, 8:推送成功回执, 9:成功和送达回执, 10:成功和点击回执, 11:成功和送达以及点击回执
		Type string `json:"type,omitempty"`
	}

	// Response represents fcm response message - (tokens and topics)
	Response struct {
		StatusCode int
		MsgId      string `json:"msg_id"`
		Sendno     string `json:"sendno,omitempty"`
		Error      Error  `json:"error,omitempty"`
	}

	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (p *Payload) Marshal() []byte {
	payload, _ := json.Marshal(p)
	return payload
}
