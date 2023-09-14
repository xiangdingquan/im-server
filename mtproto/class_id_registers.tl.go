// ConstructorList
// RequestList

package mtproto

var clazzIdRegisters2 = map[int32]func() TLObject{
	// parsed_manually_types
	1538843921: func() TLObject { // CRC32_message2
		return &TLMessage2{}
	},
	1945237724: func() TLObject { // CRC32_msg_container
		return &TLMsgContainer{}
	},
	530561358: func() TLObject { // CRC32_msg_copy
		return &TLMsgCopy{}
	},
	812830625: func() TLObject { // CRC32_gzip_packed
		return &TLGzipPacked{}
	},
	-212046591: func() TLObject { // CRC32_rpc_result
		return &TLRpcResult{}
	},
	// Constructor
	1715713620: func() TLObject { // 0x6643b654
		o := MakeTLClient_DHInnerData(nil)
		o.Data2.Constructor = 1715713620
		return o
	},
	1003222836: func() TLObject { // 0x3bcbf734
		o := MakeTLDhGenOk(nil)
		o.Data2.Constructor = 1003222836
		return o
	},
	1188831161: func() TLObject { // 0x46dc1fb9
		o := MakeTLDhGenRetry(nil)
		o.Data2.Constructor = 1188831161
		return o
	},
	-1499615742: func() TLObject { // 0xa69dae02
		o := MakeTLDhGenFail(nil)
		o.Data2.Constructor = -1499615742
		return o
	},
	-161422892: func() TLObject { // 0xf660e1d4
		o := MakeTLDestroyAuthKeyOk(nil)
		o.Data2.Constructor = -161422892
		return o
	},
	178201177: func() TLObject { // 0xa9f2259
		o := MakeTLDestroyAuthKeyNone(nil)
		o.Data2.Constructor = 178201177
		return o
	},
	-368010477: func() TLObject { // 0xea109b13
		o := MakeTLDestroyAuthKeyFail(nil)
		o.Data2.Constructor = -368010477
		return o
	},
	85337187: func() TLObject { // 0x5162463
		o := MakeTLResPQ(nil)
		o.Data2.Constructor = 85337187
		return o
	},
	-2083955988: func() TLObject { // 0x83c95aec
		o := MakeTLPQInnerData(nil)
		o.Data2.Constructor = -2083955988
		return o
	},
	-1443537003: func() TLObject { // 0xa9f55f95
		o := MakeTLPQInnerDataDc(nil)
		o.Data2.Constructor = -1443537003
		return o
	},
	1013613780: func() TLObject { // 0x3c6a84d4
		o := MakeTLPQInnerDataTemp(nil)
		o.Data2.Constructor = 1013613780
		return o
	},
	1459478408: func() TLObject { // 0x56fddf88
		o := MakeTLPQInnerDataTempDc(nil)
		o.Data2.Constructor = 1459478408
		return o
	},
	1973679973: func() TLObject { // 0x75a3f765
		o := MakeTLBindAuthKeyInner(nil)
		o.Data2.Constructor = 1973679973
		return o
	},
	2043348061: func() TLObject { // 0x79cb045d
		o := MakeTLServer_DHParamsFail(nil)
		o.Data2.Constructor = 2043348061
		return o
	},
	-790100132: func() TLObject { // 0xd0e8075c
		o := MakeTLServer_DHParamsOk(nil)
		o.Data2.Constructor = -790100132
		return o
	},
	-1249309254: func() TLObject { // 0xb5890dba
		o := MakeTLServer_DHInnerData(nil)
		o.Data2.Constructor = -1249309254
		return o
	},
	-1835453025: func() TLObject { // 0x9299359f
		o := MakeTLHttpWait(nil)
		o.Data2.Constructor = -1835453025
		return o
	},
	1182381663: func() TLObject { // 0x4679b65f
		o := MakeTLAccessPointRule(nil)
		o.Data2.Constructor = 1182381663
		return o
	},
	1108910436: func() TLObject { // 0x4218a164
		o := MakeTLTlsBlockString(nil)
		o.Data2.Constructor = 1108910436
		return o
	},
	1296942110: func() TLObject { // 0x4d4dc41e
		o := MakeTLTlsBlockRandom(nil)
		o.Data2.Constructor = 1296942110
		return o
	},
	154352379: func() TLObject { // 0x9333afb
		o := MakeTLTlsBlockZero(nil)
		o.Data2.Constructor = 154352379
		return o
	},
	283665263: func() TLObject { // 0x10e8636f
		o := MakeTLTlsBlockDomain(nil)
		o.Data2.Constructor = 283665263
		return o
	},
	-428498495: func() TLObject { // 0xe675a1c1
		o := MakeTLTlsBlockGrease(nil)
		o.Data2.Constructor = -428498495
		return o
	},
	-1632019620: func() TLObject { // 0x9eb95b5c
		o := MakeTLTlsBlockPublicKey(nil)
		o.Data2.Constructor = -1632019620
		return o
	},
	-416951217: func() TLObject { // 0xe725d44f
		o := MakeTLTlsBlockScope(nil)
		o.Data2.Constructor = -416951217
		return o
	},
	-1477445615: func() TLObject { // 0xa7eff811
		o := MakeTLBadMsgNotification(nil)
		o.Data2.Constructor = -1477445615
		return o
	},
	-307542917: func() TLObject { // 0xedab447b
		o := MakeTLBadServerSalt(nil)
		o.Data2.Constructor = -307542917
		return o
	},
	2105940488: func() TLObject { // 0x7d861a08
		o := MakeTLMsgResendReq(nil)
		o.Data2.Constructor = 2105940488
		return o
	},
	558156313: func() TLObject { // 0x2144ca19
		o := MakeTLRpcError(nil)
		o.Data2.Constructor = 558156313
		return o
	},
	-1370486635: func() TLObject { // 0xae500895
		o := MakeTLFutureSalts(nil)
		o.Data2.Constructor = -1370486635
		return o
	},
	-501201412: func() TLObject { // 0xe22045fc
		o := MakeTLDestroySessionOk(nil)
		o.Data2.Constructor = -501201412
		return o
	},
	1658015945: func() TLObject { // 0x62d350c9
		o := MakeTLDestroySessionNone(nil)
		o.Data2.Constructor = 1658015945
		return o
	},
	-1631450872: func() TLObject { // 0x9ec20908
		o := MakeTLNewSessionCreated(nil)
		o.Data2.Constructor = -1631450872
		return o
	},
	-734810765: func() TLObject { // 0xd433ad73
		o := MakeTLIpPort(nil)
		o.Data2.Constructor = -734810765
		return o
	},
	932718150: func() TLObject { // 0x37982646
		o := MakeTLIpPortSecret(nil)
		o.Data2.Constructor = 932718150
		return o
	},
	1658238041: func() TLObject { // 0x62d6b459
		o := MakeTLMsgsAck(nil)
		o.Data2.Constructor = 1658238041
		return o
	},
	-630588590: func() TLObject { // 0xda69fb52
		o := MakeTLMsgsStateReq(nil)
		o.Data2.Constructor = -630588590
		return o
	},
	661470918: func() TLObject { // 0x276d3ec6
		o := MakeTLMsgDetailedInfo(nil)
		o.Data2.Constructor = 661470918
		return o
	},
	-2137147681: func() TLObject { // 0x809db6df
		o := MakeTLMsgNewDetailedInfo(nil)
		o.Data2.Constructor = -2137147681
		return o
	},
	155834844: func() TLObject { // 0x949d9dc
		o := MakeTLFutureSalt(nil)
		o.Data2.Constructor = 155834844
		return o
	},
	1579864942: func() TLObject { // 0x5e2ad36e
		o := MakeTLRpcAnswerUnknown(nil)
		o.Data2.Constructor = 1579864942
		return o
	},
	-847714938: func() TLObject { // 0xcd78e586
		o := MakeTLRpcAnswerDroppedRunning(nil)
		o.Data2.Constructor = -847714938
		return o
	},
	-1539647305: func() TLObject { // 0xa43ad8b7
		o := MakeTLRpcAnswerDropped(nil)
		o.Data2.Constructor = -1539647305
		return o
	},
	880243653: func() TLObject { // 0x347773c5
		o := MakeTLPong(nil)
		o.Data2.Constructor = 880243653
		return o
	},
	1817363588: func() TLObject { // 0x6c52c484
		o := MakeTLTlsClientHello(nil)
		o.Data2.Constructor = 1817363588
		return o
	},
	81704317: func() TLObject { // 0x4deb57d
		o := MakeTLMsgsStateInfo(nil)
		o.Data2.Constructor = 81704317
		return o
	},
	-1933520591: func() TLObject { // 0x8cc0d131
		o := MakeTLMsgsAllInfo(nil)
		o.Data2.Constructor = -1933520591
		return o
	},
	1515793004: func() TLObject { // 0x5a592a6c
		o := MakeTLHelpConfigSimple(nil)
		o.Data2.Constructor = 1515793004
		return o
	},
	-644365371: func() TLObject { // 0xd997c3c5
		o := MakeTLHelpConfigSimple(nil)
		o.Data2.Constructor = -644365371
		return o
	},
	-1148011883: func() TLObject { // 0xbb92ba95
		o := MakeTLMessageEntityUnknown(nil)
		o.Data2.Constructor = -1148011883
		return o
	},
	-100378723: func() TLObject { // 0xfa04579d
		o := MakeTLMessageEntityMention(nil)
		o.Data2.Constructor = -100378723
		return o
	},
	1868782349: func() TLObject { // 0x6f635b0d
		o := MakeTLMessageEntityHashtag(nil)
		o.Data2.Constructor = 1868782349
		return o
	},
	1827637959: func() TLObject { // 0x6cef8ac7
		o := MakeTLMessageEntityBotCommand(nil)
		o.Data2.Constructor = 1827637959
		return o
	},
	1859134776: func() TLObject { // 0x6ed02538
		o := MakeTLMessageEntityUrl(nil)
		o.Data2.Constructor = 1859134776
		return o
	},
	1692693954: func() TLObject { // 0x64e475c2
		o := MakeTLMessageEntityEmail(nil)
		o.Data2.Constructor = 1692693954
		return o
	},
	-1117713463: func() TLObject { // 0xbd610bc9
		o := MakeTLMessageEntityBold(nil)
		o.Data2.Constructor = -1117713463
		return o
	},
	-2106619040: func() TLObject { // 0x826f8b60
		o := MakeTLMessageEntityItalic(nil)
		o.Data2.Constructor = -2106619040
		return o
	},
	681706865: func() TLObject { // 0x28a20571
		o := MakeTLMessageEntityCode(nil)
		o.Data2.Constructor = 681706865
		return o
	},
	1938967520: func() TLObject { // 0x73924be0
		o := MakeTLMessageEntityPre(nil)
		o.Data2.Constructor = 1938967520
		return o
	},
	1990644519: func() TLObject { // 0x76a6d327
		o := MakeTLMessageEntityTextUrl(nil)
		o.Data2.Constructor = 1990644519
		return o
	},
	892193368: func() TLObject { // 0x352dca58
		o := MakeTLMessageEntityMentionName(nil)
		o.Data2.Constructor = 892193368
		return o
	},
	546203849: func() TLObject { // 0x208e68c9
		o := MakeTLInputMessageEntityMentionName(nil)
		o.Data2.Constructor = 546203849
		return o
	},
	-1687559349: func() TLObject { // 0x9b69e34b
		o := MakeTLMessageEntityPhone(nil)
		o.Data2.Constructor = -1687559349
		return o
	},
	1280209983: func() TLObject { // 0x4c4e743f
		o := MakeTLMessageEntityCashtag(nil)
		o.Data2.Constructor = 1280209983
		return o
	},
	-1672577397: func() TLObject { // 0x9c4e7e8b
		o := MakeTLMessageEntityUnderline(nil)
		o.Data2.Constructor = -1672577397
		return o
	},
	-1090087980: func() TLObject { // 0xbf0693d4
		o := MakeTLMessageEntityStrike(nil)
		o.Data2.Constructor = -1090087980
		return o
	},
	34469328: func() TLObject { // 0x20df5d0
		o := MakeTLMessageEntityBlockquote(nil)
		o.Data2.Constructor = 34469328
		return o
	},
	1981704948: func() TLObject { // 0x761e6af4
		o := MakeTLMessageEntityBankCard(nil)
		o.Data2.Constructor = 1981704948
		return o
	},
	1984755728: func() TLObject { // 0x764cf810
		o := MakeTLBotInlineMessageMediaAuto(nil)
		o.Data2.Constructor = 1984755728
		return o
	},
	175419739: func() TLObject { // 0xa74b15b
		o := MakeTLBotInlineMessageMediaAuto(nil)
		o.Data2.Constructor = 175419739
		return o
	},
	-1937807902: func() TLObject { // 0x8c7f65e2
		o := MakeTLBotInlineMessageText(nil)
		o.Data2.Constructor = -1937807902
		return o
	},
	85477117: func() TLObject { // 0x51846fd
		o := MakeTLBotInlineMessageMediaGeo(nil)
		o.Data2.Constructor = 85477117
		return o
	},
	-1222451611: func() TLObject { // 0xb722de65
		o := MakeTLBotInlineMessageMediaGeo(nil)
		o.Data2.Constructor = -1222451611
		return o
	},
	-1970903652: func() TLObject { // 0x8a86659c
		o := MakeTLBotInlineMessageMediaVenue(nil)
		o.Data2.Constructor = -1970903652
		return o
	},
	1130767150: func() TLObject { // 0x4366232e
		o := MakeTLBotInlineMessageMediaVenue(nil)
		o.Data2.Constructor = 1130767150
		return o
	},
	416402882: func() TLObject { // 0x18d1cdc2
		o := MakeTLBotInlineMessageMediaContact(nil)
		o.Data2.Constructor = 416402882
		return o
	},
	904770772: func() TLObject { // 0x35edb4d4
		o := MakeTLBotInlineMessageMediaContact(nil)
		o.Data2.Constructor = 904770772
		return o
	},
	1314881805: func() TLObject { // 0x4e5f810d
		o := MakeTLPaymentsPaymentResult(nil)
		o.Data2.Constructor = 1314881805
		return o
	},
	-666824391: func() TLObject { // 0xd8411139
		o := MakeTLPaymentsPaymentVerificationNeeded(nil)
		o.Data2.Constructor = -666824391
		return o
	},
	1800845601: func() TLObject { // 0x6b56b921
		o := MakeTLPaymentsPaymentVerficationNeeded(nil)
		o.Data2.Constructor = 1800845601
		return o
	},
	42930452: func() TLObject { // 0x28f1114
		o := MakeTLTheme(nil)
		o.Data2.Constructor = 42930452
		return o
	},
	-136770336: func() TLObject { // 0xf7d90ce0
		o := MakeTLTheme(nil)
		o.Data2.Constructor = -136770336
		return o
	},
	1211967244: func() TLObject { // 0x483d270c
		o := MakeTLThemeDocumentNotModified(nil)
		o.Data2.Constructor = 1211967244
		return o
	},
	-1038136962: func() TLObject { // 0xc21f497e
		o := MakeTLEncryptedFileEmpty(nil)
		o.Data2.Constructor = -1038136962
		return o
	},
	1248893260: func() TLObject { // 0x4a70994c
		o := MakeTLEncryptedFile(nil)
		o.Data2.Constructor = 1248893260
		return o
	},
	522914557: func() TLObject { // 0x1f2b0afd
		o := MakeTLUpdateNewMessage(nil)
		o.Data2.Constructor = 522914557
		return o
	},
	1318109142: func() TLObject { // 0x4e90bfd6
		o := MakeTLUpdateMessageID(nil)
		o.Data2.Constructor = 1318109142
		return o
	},
	-1576161051: func() TLObject { // 0xa20db0e5
		o := MakeTLUpdateDeleteMessages(nil)
		o.Data2.Constructor = -1576161051
		return o
	},
	1548249383: func() TLObject { // 0x5c486927
		o := MakeTLUpdateUserTyping(nil)
		o.Data2.Constructor = 1548249383
		return o
	},
	-1704596961: func() TLObject { // 0x9a65ea1f
		o := MakeTLUpdateChatUserTyping(nil)
		o.Data2.Constructor = -1704596961
		return o
	},
	125178264: func() TLObject { // 0x7761198
		o := MakeTLUpdateChatParticipants(nil)
		o.Data2.Constructor = 125178264
		return o
	},
	469489699: func() TLObject { // 0x1bfbd823
		o := MakeTLUpdateUserStatus(nil)
		o.Data2.Constructor = 469489699
		return o
	},
	-1489818765: func() TLObject { // 0xa7332b73
		o := MakeTLUpdateUserName(nil)
		o.Data2.Constructor = -1489818765
		return o
	},
	-1791935732: func() TLObject { // 0x95313b0c
		o := MakeTLUpdateUserPhoto(nil)
		o.Data2.Constructor = -1791935732
		return o
	},
	314359194: func() TLObject { // 0x12bcbd9a
		o := MakeTLUpdateNewEncryptedMessage(nil)
		o.Data2.Constructor = 314359194
		return o
	},
	386986326: func() TLObject { // 0x1710f156
		o := MakeTLUpdateEncryptedChatTyping(nil)
		o.Data2.Constructor = 386986326
		return o
	},
	-1264392051: func() TLObject { // 0xb4a2e88d
		o := MakeTLUpdateEncryption(nil)
		o.Data2.Constructor = -1264392051
		return o
	},
	956179895: func() TLObject { // 0x38fe25b7
		o := MakeTLUpdateEncryptedMessagesRead(nil)
		o.Data2.Constructor = 956179895
		return o
	},
	-364179876: func() TLObject { // 0xea4b0e5c
		o := MakeTLUpdateChatParticipantAdd(nil)
		o.Data2.Constructor = -364179876
		return o
	},
	1851755554: func() TLObject { // 0x6e5f8c22
		o := MakeTLUpdateChatParticipantDelete(nil)
		o.Data2.Constructor = 1851755554
		return o
	},
	-1906403213: func() TLObject { // 0x8e5e9873
		o := MakeTLUpdateDcOptions(nil)
		o.Data2.Constructor = -1906403213
		return o
	},
	-1094555409: func() TLObject { // 0xbec268ef
		o := MakeTLUpdateNotifySettings(nil)
		o.Data2.Constructor = -1094555409
		return o
	},
	-337352679: func() TLObject { // 0xebe46819
		o := MakeTLUpdateServiceNotification(nil)
		o.Data2.Constructor = -337352679
		return o
	},
	-298113238: func() TLObject { // 0xee3b272a
		o := MakeTLUpdatePrivacy(nil)
		o.Data2.Constructor = -298113238
		return o
	},
	314130811: func() TLObject { // 0x12b9417b
		o := MakeTLUpdateUserPhone(nil)
		o.Data2.Constructor = 314130811
		return o
	},
	-1667805217: func() TLObject { // 0x9c974fdf
		o := MakeTLUpdateReadHistoryInbox(nil)
		o.Data2.Constructor = -1667805217
		return o
	},
	-1721631396: func() TLObject { // 0x9961fd5c
		o := MakeTLUpdateReadHistoryInbox(nil)
		o.Data2.Constructor = -1721631396
		return o
	},
	791617983: func() TLObject { // 0x2f2f21bf
		o := MakeTLUpdateReadHistoryOutbox(nil)
		o.Data2.Constructor = 791617983
		return o
	},
	2139689491: func() TLObject { // 0x7f891213
		o := MakeTLUpdateWebPage(nil)
		o.Data2.Constructor = 2139689491
		return o
	},
	1757493555: func() TLObject { // 0x68c13933
		o := MakeTLUpdateReadMessagesContents(nil)
		o.Data2.Constructor = 1757493555
		return o
	},
	-352032773: func() TLObject { // 0xeb0467fb
		o := MakeTLUpdateChannelTooLong(nil)
		o.Data2.Constructor = -352032773
		return o
	},
	-1227598250: func() TLObject { // 0xb6d45656
		o := MakeTLUpdateChannel(nil)
		o.Data2.Constructor = -1227598250
		return o
	},
	1656358105: func() TLObject { // 0x62ba04d9
		o := MakeTLUpdateNewChannelMessage(nil)
		o.Data2.Constructor = 1656358105
		return o
	},
	856380452: func() TLObject { // 0x330b5424
		o := MakeTLUpdateReadChannelInbox(nil)
		o.Data2.Constructor = 856380452
		return o
	},
	1108669311: func() TLObject { // 0x4214f37f
		o := MakeTLUpdateReadChannelInbox(nil)
		o.Data2.Constructor = 1108669311
		return o
	},
	-1015733815: func() TLObject { // 0xc37521c9
		o := MakeTLUpdateDeleteChannelMessages(nil)
		o.Data2.Constructor = -1015733815
		return o
	},
	-1734268085: func() TLObject { // 0x98a12b4b
		o := MakeTLUpdateChannelMessageViews(nil)
		o.Data2.Constructor = -1734268085
		return o
	},
	-1232070311: func() TLObject { // 0xb6901959
		o := MakeTLUpdateChatParticipantAdmin(nil)
		o.Data2.Constructor = -1232070311
		return o
	},
	1753886890: func() TLObject { // 0x688a30aa
		o := MakeTLUpdateNewStickerSet(nil)
		o.Data2.Constructor = 1753886890
		return o
	},
	196268545: func() TLObject { // 0xbb2d201
		o := MakeTLUpdateStickerSetsOrder(nil)
		o.Data2.Constructor = 196268545
		return o
	},
	1135492588: func() TLObject { // 0x43ae3dec
		o := MakeTLUpdateStickerSets(nil)
		o.Data2.Constructor = 1135492588
		return o
	},
	-1821035490: func() TLObject { // 0x9375341e
		o := MakeTLUpdateSavedGifs(nil)
		o.Data2.Constructor = -1821035490
		return o
	},
	1417832080: func() TLObject { // 0x54826690
		o := MakeTLUpdateBotInlineQuery(nil)
		o.Data2.Constructor = 1417832080
		return o
	},
	239663460: func() TLObject { // 0xe48f964
		o := MakeTLUpdateBotInlineSend(nil)
		o.Data2.Constructor = 239663460
		return o
	},
	457133559: func() TLObject { // 0x1b3f4df7
		o := MakeTLUpdateEditChannelMessage(nil)
		o.Data2.Constructor = 457133559
		return o
	},
	-415938591: func() TLObject { // 0xe73547e1
		o := MakeTLUpdateBotCallbackQuery(nil)
		o.Data2.Constructor = -415938591
		return o
	},
	-469536605: func() TLObject { // 0xe40370a3
		o := MakeTLUpdateEditMessage(nil)
		o.Data2.Constructor = -469536605
		return o
	},
	-103646630: func() TLObject { // 0xf9d27a5a
		o := MakeTLUpdateInlineBotCallbackQuery(nil)
		o.Data2.Constructor = -103646630
		return o
	},
	634833351: func() TLObject { // 0x25d6c9c7
		o := MakeTLUpdateReadChannelOutbox(nil)
		o.Data2.Constructor = 634833351
		return o
	},
	-299124375: func() TLObject { // 0xee2bb969
		o := MakeTLUpdateDraftMessage(nil)
		o.Data2.Constructor = -299124375
		return o
	},
	1461528386: func() TLObject { // 0x571d2742
		o := MakeTLUpdateReadFeaturedStickers(nil)
		o.Data2.Constructor = 1461528386
		return o
	},
	-1706939360: func() TLObject { // 0x9a422c20
		o := MakeTLUpdateRecentStickers(nil)
		o.Data2.Constructor = -1706939360
		return o
	},
	-1574314746: func() TLObject { // 0xa229dd06
		o := MakeTLUpdateConfig(nil)
		o.Data2.Constructor = -1574314746
		return o
	},
	861169551: func() TLObject { // 0x3354678f
		o := MakeTLUpdatePtsChanged(nil)
		o.Data2.Constructor = 861169551
		return o
	},
	1081547008: func() TLObject { // 0x40771900
		o := MakeTLUpdateChannelWebPage(nil)
		o.Data2.Constructor = 1081547008
		return o
	},
	1852826908: func() TLObject { // 0x6e6fe51c
		o := MakeTLUpdateDialogPinned(nil)
		o.Data2.Constructor = 1852826908
		return o
	},
	433225532: func() TLObject { // 0x19d27f3c
		o := MakeTLUpdateDialogPinned(nil)
		o.Data2.Constructor = 433225532
		return o
	},
	-686710068: func() TLObject { // 0xd711a2cc
		o := MakeTLUpdateDialogPinned(nil)
		o.Data2.Constructor = -686710068
		return o
	},
	-99664734: func() TLObject { // 0xfa0f3ca2
		o := MakeTLUpdatePinnedDialogs(nil)
		o.Data2.Constructor = -99664734
		return o
	},
	-364071333: func() TLObject { // 0xea4cb65b
		o := MakeTLUpdatePinnedDialogs(nil)
		o.Data2.Constructor = -364071333
		return o
	},
	-657787251: func() TLObject { // 0xd8caf68d
		o := MakeTLUpdatePinnedDialogs(nil)
		o.Data2.Constructor = -657787251
		return o
	},
	-2095595325: func() TLObject { // 0x8317c0c3
		o := MakeTLUpdateBotWebhookJSON(nil)
		o.Data2.Constructor = -2095595325
		return o
	},
	-1684914010: func() TLObject { // 0x9b9240a6
		o := MakeTLUpdateBotWebhookJSONQuery(nil)
		o.Data2.Constructor = -1684914010
		return o
	},
	-523384512: func() TLObject { // 0xe0cdc940
		o := MakeTLUpdateBotShippingQuery(nil)
		o.Data2.Constructor = -523384512
		return o
	},
	1563376297: func() TLObject { // 0x5d2f3aa9
		o := MakeTLUpdateBotPrecheckoutQuery(nil)
		o.Data2.Constructor = 1563376297
		return o
	},
	-1425052898: func() TLObject { // 0xab0f6b1e
		o := MakeTLUpdatePhoneCall(nil)
		o.Data2.Constructor = -1425052898
		return o
	},
	1180041828: func() TLObject { // 0x46560264
		o := MakeTLUpdateLangPackTooLong(nil)
		o.Data2.Constructor = 1180041828
		return o
	},
	281165899: func() TLObject { // 0x10c2404b
		o := MakeTLUpdateLangPackTooLong(nil)
		o.Data2.Constructor = 281165899
		return o
	},
	1442983757: func() TLObject { // 0x56022f4d
		o := MakeTLUpdateLangPack(nil)
		o.Data2.Constructor = 1442983757
		return o
	},
	-451831443: func() TLObject { // 0xe511996d
		o := MakeTLUpdateFavedStickers(nil)
		o.Data2.Constructor = -451831443
		return o
	},
	-1987495099: func() TLObject { // 0x89893b45
		o := MakeTLUpdateChannelReadMessagesContents(nil)
		o.Data2.Constructor = -1987495099
		return o
	},
	1887741886: func() TLObject { // 0x7084a7be
		o := MakeTLUpdateContactsReset(nil)
		o.Data2.Constructor = 1887741886
		return o
	},
	1893427255: func() TLObject { // 0x70db6837
		o := MakeTLUpdateChannelAvailableMessages(nil)
		o.Data2.Constructor = 1893427255
		return o
	},
	-513517117: func() TLObject { // 0xe16459c3
		o := MakeTLUpdateDialogUnreadMark(nil)
		o.Data2.Constructor = -513517117
		return o
	},
	-1398708869: func() TLObject { // 0xaca1657b
		o := MakeTLUpdateMessagePoll(nil)
		o.Data2.Constructor = -1398708869
		return o
	},
	1421875280: func() TLObject { // 0x54c01850
		o := MakeTLUpdateChatDefaultBannedRights(nil)
		o.Data2.Constructor = 1421875280
		return o
	},
	422972864: func() TLObject { // 0x19360dc0
		o := MakeTLUpdateFolderPeers(nil)
		o.Data2.Constructor = 422972864
		return o
	},
	1786671974: func() TLObject { // 0x6a7e7366
		o := MakeTLUpdatePeerSettings(nil)
		o.Data2.Constructor = 1786671974
		return o
	},
	-1263546448: func() TLObject { // 0xb4afcfb0
		o := MakeTLUpdatePeerLocated(nil)
		o.Data2.Constructor = -1263546448
		return o
	},
	967122427: func() TLObject { // 0x39a51dfb
		o := MakeTLUpdateNewScheduledMessage(nil)
		o.Data2.Constructor = 967122427
		return o
	},
	-1870238482: func() TLObject { // 0x90866cee
		o := MakeTLUpdateDeleteScheduledMessages(nil)
		o.Data2.Constructor = -1870238482
		return o
	},
	-2112423005: func() TLObject { // 0x8216fba3
		o := MakeTLUpdateTheme(nil)
		o.Data2.Constructor = -2112423005
		return o
	},
	-2027964103: func() TLObject { // 0x871fb939
		o := MakeTLUpdateGeoLiveViewed(nil)
		o.Data2.Constructor = -2027964103
		return o
	},
	1448076945: func() TLObject { // 0x564fe691
		o := MakeTLUpdateLoginToken(nil)
		o.Data2.Constructor = 1448076945
		return o
	},
	1123585836: func() TLObject { // 0x42f88f2c
		o := MakeTLUpdateMessagePollVote(nil)
		o.Data2.Constructor = 1123585836
		return o
	},
	654302845: func() TLObject { // 0x26ffde7d
		o := MakeTLUpdateDialogFilter(nil)
		o.Data2.Constructor = 654302845
		return o
	},
	-1512627963: func() TLObject { // 0xa5d72105
		o := MakeTLUpdateDialogFilterOrder(nil)
		o.Data2.Constructor = -1512627963
		return o
	},
	889491791: func() TLObject { // 0x3504914f
		o := MakeTLUpdateDialogFilters(nil)
		o.Data2.Constructor = 889491791
		return o
	},
	643940105: func() TLObject { // 0x2661bf09
		o := MakeTLUpdatePhoneCallSignalingData(nil)
		o.Data2.Constructor = 643940105
		return o
	},
	1708307556: func() TLObject { // 0x65d2b464
		o := MakeTLUpdateChannelParticipant(nil)
		o.Data2.Constructor = 1708307556
		return o
	},
	1854571743: func() TLObject { // 0x6e8a84df
		o := MakeTLUpdateChannelMessageForwards(nil)
		o.Data2.Constructor = 1854571743
		return o
	},
	482860628: func() TLObject { // 0x1cc7de54
		o := MakeTLUpdateReadChannelDiscussionInbox(nil)
		o.Data2.Constructor = 482860628
		return o
	},
	1178116716: func() TLObject { // 0x4638a26c
		o := MakeTLUpdateReadChannelDiscussionOutbox(nil)
		o.Data2.Constructor = 1178116716
		return o
	},
	610945826: func() TLObject { // 0x246a4b22
		o := MakeTLUpdatePeerBlocked(nil)
		o.Data2.Constructor = 610945826
		return o
	},
	-13975905: func() TLObject { // 0xff2abe9f
		o := MakeTLUpdateChannelUserTyping(nil)
		o.Data2.Constructor = -13975905
		return o
	},
	-309990731: func() TLObject { // 0xed85eab5
		o := MakeTLUpdatePinnedMessages(nil)
		o.Data2.Constructor = -309990731
		return o
	},
	-2054649973: func() TLObject { // 0x8588878b
		o := MakeTLUpdatePinnedChannelMessages(nil)
		o.Data2.Constructor = -2054649973
		return o
	},
	-1038026530: func() TLObject { // 0xc220f8de
		o := MakeTLUpdateNewBlog(nil)
		o.Data2.Constructor = -1038026530
		return o
	},
	-1147315158: func() TLObject { // 0xbb9d5c2a
		o := MakeTLUpdateBlogID(nil)
		o.Data2.Constructor = -1147315158
		return o
	},
	1430149653: func() TLObject { // 0x553e5a15
		o := MakeTLUpdateDeleteBlog(nil)
		o.Data2.Constructor = 1430149653
		return o
	},
	-1139005339: func() TLObject { // 0xbc1c2865
		o := MakeTLUpdateBlogFollow(nil)
		o.Data2.Constructor = -1139005339
		return o
	},
	-1285237824: func() TLObject { // 0xb364d3c0
		o := MakeTLUpdateBlogComment(nil)
		o.Data2.Constructor = -1285237824
		return o
	},
	323077162: func() TLObject { // 0x1341c42a
		o := MakeTLUpdateBlogLike(nil)
		o.Data2.Constructor = 323077162
		return o
	},
	1798460205: func() TLObject { // 0x6b32532d
		o := MakeTLUpdateBlogUserBlocked(nil)
		o.Data2.Constructor = 1798460205
		return o
	},
	-1995270778: func() TLObject { // 0x89129586
		o := MakeTLUpdateNewBlogGroupTag(nil)
		o.Data2.Constructor = -1995270778
		return o
	},
	565242618: func() TLObject { // 0x21b0eafa
		o := MakeTLUpdateEditBlogGroupTag(nil)
		o.Data2.Constructor = 565242618
		return o
	},
	-582768106: func() TLObject { // 0xdd43aa16
		o := MakeTLUpdateDeleteBlogGroupTag(nil)
		o.Data2.Constructor = -582768106
		return o
	},
	-1216284815: func() TLObject { // 0xb780f771
		o := MakeTLUpdateAddBlogGroupTagMember(nil)
		o.Data2.Constructor = -1216284815
		return o
	},
	1829842702: func() TLObject { // 0x6d112f0e
		o := MakeTLUpdateDeleteBlogGroupTagMember(nil)
		o.Data2.Constructor = 1829842702
		return o
	},
	108681335: func() TLObject { // 0x67a5877
		o := MakeTLUpdateBlogReadHistory(nil)
		o.Data2.Constructor = 108681335
		return o
	},
	-1671982488: func() TLObject { // 0x9c579268
		o := MakeTLUpdateBlogReward(nil)
		o.Data2.Constructor = -1671982488
		return o
	},
	1299120167: func() TLObject { // 0x4d6f0027
		o := MakeTLUpdateTopic(nil)
		o.Data2.Constructor = 1299120167
		return o
	},
	-1738988427: func() TLObject { // 0x98592475
		o := MakeTLUpdateChannelPinnedMessage(nil)
		o.Data2.Constructor = -1738988427
		return o
	},
	1279515160: func() TLObject { // 0x4c43da18
		o := MakeTLUpdateUserPinnedMessage(nil)
		o.Data2.Constructor = 1279515160
		return o
	},
	-519195831: func() TLObject { // 0xe10db349
		o := MakeTLUpdateChatPinnedMessage(nil)
		o.Data2.Constructor = -519195831
		return o
	},
	579418918: func() TLObject { // 0x22893b26
		o := MakeTLUpdateChatPinnedMessage(nil)
		o.Data2.Constructor = 579418918
		return o
	},
	-2131957734: func() TLObject { // 0x80ece81a
		o := MakeTLUpdateUserBlocked(nil)
		o.Data2.Constructor = -2131957734
		return o
	},
	-1657903163: func() TLObject { // 0x9d2e67c5
		o := MakeTLUpdateContactLink(nil)
		o.Data2.Constructor = -1657903163
		return o
	},
	1855224129: func() TLObject { // 0x6e947941
		o := MakeTLUpdateChatAdmins(nil)
		o.Data2.Constructor = 1855224129
		return o
	},
	628472761: func() TLObject { // 0x2575bbb9
		o := MakeTLUpdateContactRegistered(nil)
		o.Data2.Constructor = 628472761
		return o
	},
	1873186369: func() TLObject { // 0x6fa68e41
		o := MakeTLUpdateReadFeed(nil)
		o.Data2.Constructor = 1873186369
		return o
	},
	-1723313495: func() TLObject { // 0x994852a9
		o := MakeTLUpdateReadFeed(nil)
		o.Data2.Constructor = -1723313495
		return o
	},
	-2083620338: func() TLObject { // 0x83ce7a0e
		o := MakeTLUpdateBizDataRaw(nil)
		o.Data2.Constructor = -2083620338
		return o
	},
	1892275565: func() TLObject { // 0x70c9d56d
		o := MakeTLChannelParticipant(nil)
		o.Data2.Constructor = 1892275565
		return o
	},
	1348540138: func() TLObject { // 0x506116ea
		o := MakeTLChannelParticipant(nil)
		o.Data2.Constructor = 1348540138
		return o
	},
	367766557: func() TLObject { // 0x15ebac1d
		o := MakeTLChannelParticipant(nil)
		o.Data2.Constructor = 367766557
		return o
	},
	-1845189537: func() TLObject { // 0x9204a45f
		o := MakeTLChannelParticipantSelf(nil)
		o.Data2.Constructor = -1845189537
		return o
	},
	-1557620115: func() TLObject { // 0xa3289a6d
		o := MakeTLChannelParticipantSelf(nil)
		o.Data2.Constructor = -1557620115
		return o
	},
	678497813: func() TLObject { // 0x28710e15
		o := MakeTLChannelParticipantCreator(nil)
		o.Data2.Constructor = 678497813
		return o
	},
	1149094475: func() TLObject { // 0x447dca4b
		o := MakeTLChannelParticipantCreator(nil)
		o.Data2.Constructor = 1149094475
		return o
	},
	-2138237532: func() TLObject { // 0x808d15a4
		o := MakeTLChannelParticipantCreator(nil)
		o.Data2.Constructor = -2138237532
		return o
	},
	-471670279: func() TLObject { // 0xe3e2e1f9
		o := MakeTLChannelParticipantCreator(nil)
		o.Data2.Constructor = -471670279
		return o
	},
	1654253591: func() TLObject { // 0x6299e817
		o := MakeTLChannelParticipantAdmin(nil)
		o.Data2.Constructor = 1654253591
		return o
	},
	-859915345: func() TLObject { // 0xccbebbaf
		o := MakeTLChannelParticipantAdmin(nil)
		o.Data2.Constructor = -859915345
		return o
	},
	1571450403: func() TLObject { // 0x5daa6e23
		o := MakeTLChannelParticipantAdmin(nil)
		o.Data2.Constructor = 1571450403
		return o
	},
	-1473271656: func() TLObject { // 0xa82fa898
		o := MakeTLChannelParticipantAdmin(nil)
		o.Data2.Constructor = -1473271656
		return o
	},
	-1582292988: func() TLObject { // 0xa1b02004
		o := MakeTLChannelParticipantBanned(nil)
		o.Data2.Constructor = -1582292988
		return o
	},
	470789295: func() TLObject { // 0x1c0facaf
		o := MakeTLChannelParticipantBanned(nil)
		o.Data2.Constructor = 470789295
		return o
	},
	573315206: func() TLObject { // 0x222c1886
		o := MakeTLChannelParticipantBanned(nil)
		o.Data2.Constructor = 573315206
		return o
	},
	-1010402965: func() TLObject { // 0xc3c6796b
		o := MakeTLChannelParticipantLeft(nil)
		o.Data2.Constructor = -1010402965
		return o
	},
	-11252123: func() TLObject { // 0xff544e65
		o := MakeTLFolder(nil)
		o.Data2.Constructor = -11252123
		return o
	},
	1654593920: func() TLObject { // 0x629f1980
		o := MakeTLAuthLoginToken(nil)
		o.Data2.Constructor = 1654593920
		return o
	},
	110008598: func() TLObject { // 0x68e9916
		o := MakeTLAuthLoginTokenMigrateTo(nil)
		o.Data2.Constructor = 110008598
		return o
	},
	957176926: func() TLObject { // 0x390d5c5e
		o := MakeTLAuthLoginTokenSuccess(nil)
		o.Data2.Constructor = 957176926
		return o
	},
	-827700856: func() TLObject { // 0xceaa4988
		o := MakeTLBlogsGroupTag(nil)
		o.Data2.Constructor = -827700856
		return o
	},
	406307684: func() TLObject { // 0x1837c364
		o := MakeTLInputEncryptedFileEmpty(nil)
		o.Data2.Constructor = 406307684
		return o
	},
	1690108678: func() TLObject { // 0x64bd0306
		o := MakeTLInputEncryptedFileUploaded(nil)
		o.Data2.Constructor = 1690108678
		return o
	},
	1511503333: func() TLObject { // 0x5a17b5e5
		o := MakeTLInputEncryptedFile(nil)
		o.Data2.Constructor = 1511503333
		return o
	},
	767652808: func() TLObject { // 0x2dc173c8
		o := MakeTLInputEncryptedFileBigUploaded(nil)
		o.Data2.Constructor = 767652808
		return o
	},
	-1056001329: func() TLObject { // 0xc10eb2cf
		o := MakeTLInputPaymentCredentialsSaved(nil)
		o.Data2.Constructor = -1056001329
		return o
	},
	873977640: func() TLObject { // 0x3417d728
		o := MakeTLInputPaymentCredentials(nil)
		o.Data2.Constructor = 873977640
		return o
	},
	178373535: func() TLObject { // 0xaa1c39f
		o := MakeTLInputPaymentCredentialsApplePay(nil)
		o.Data2.Constructor = 178373535
		return o
	},
	-905587442: func() TLObject { // 0xca05d50e
		o := MakeTLInputPaymentCredentialsAndroidPay(nil)
		o.Data2.Constructor = -905587442
		return o
	},
	-1078332329: func() TLObject { // 0xbfb9f457
		o := MakeTLHelpPassportConfigNotModified(nil)
		o.Data2.Constructor = -1078332329
		return o
	},
	-1600596305: func() TLObject { // 0xa098d6af
		o := MakeTLHelpPassportConfig(nil)
		o.Data2.Constructor = -1600596305
		return o
	},
	1490799288: func() TLObject { // 0x58dbcab8
		o := MakeTLInputReportReasonSpam(nil)
		o.Data2.Constructor = 1490799288
		return o
	},
	505595789: func() TLObject { // 0x1e22c78d
		o := MakeTLInputReportReasonViolence(nil)
		o.Data2.Constructor = 505595789
		return o
	},
	777640226: func() TLObject { // 0x2e59d922
		o := MakeTLInputReportReasonPornography(nil)
		o.Data2.Constructor = 777640226
		return o
	},
	-1376497949: func() TLObject { // 0xadf44ee3
		o := MakeTLInputReportReasonChildAbuse(nil)
		o.Data2.Constructor = -1376497949
		return o
	},
	-512463606: func() TLObject { // 0xe1746d0a
		o := MakeTLInputReportReasonOther(nil)
		o.Data2.Constructor = -512463606
		return o
	},
	-1685456582: func() TLObject { // 0x9b89f93a
		o := MakeTLInputReportReasonCopyright(nil)
		o.Data2.Constructor = -1685456582
		return o
	},
	-606798099: func() TLObject { // 0xdbd4feed
		o := MakeTLInputReportReasonGeoIrrelevant(nil)
		o.Data2.Constructor = -606798099
		return o
	},
	1556570557: func() TLObject { // 0x5cc761bd
		o := MakeTLEmojiKeywordsDifference(nil)
		o.Data2.Constructor = 1556570557
		return o
	},
	-1118798639: func() TLObject { // 0xbd507cd1
		o := MakeTLInputThemeSettings(nil)
		o.Data2.Constructor = -1118798639
		return o
	},
	1163625789: func() TLObject { // 0x455b853d
		o := MakeTLMessageViews(nil)
		o.Data2.Constructor = 1163625789
		return o
	},
	-290164953: func() TLObject { // 0xeeb46f27
		o := MakeTLStickerSet(nil)
		o.Data2.Constructor = -290164953
		return o
	},
	1787870391: func() TLObject { // 0x6a90bcb7
		o := MakeTLStickerSet(nil)
		o.Data2.Constructor = 1787870391
		return o
	},
	1434820921: func() TLObject { // 0x5585a139
		o := MakeTLStickerSet(nil)
		o.Data2.Constructor = 1434820921
		return o
	},
	-852477119: func() TLObject { // 0xcd303b41
		o := MakeTLStickerSet(nil)
		o.Data2.Constructor = -852477119
		return o
	},
	1776236393: func() TLObject { // 0x69df3769
		o := MakeTLChatInviteEmpty(nil)
		o.Data2.Constructor = 1776236393
		return o
	},
	-64092740: func() TLObject { // 0xfc2e05bc
		o := MakeTLChatInviteExported(nil)
		o.Data2.Constructor = -64092740
		return o
	},
	-1519029347: func() TLObject { // 0xa575739d
		o := MakeTLEmojiURL(nil)
		o.Data2.Constructor = -1519029347
		return o
	},
	-1096616924: func() TLObject { // 0xbea2f424
		o := MakeTLGlobalPrivacySettings(nil)
		o.Data2.Constructor = -1096616924
		return o
	},
	1794557279: func() TLObject { // 0x6af6c55f
		o := MakeTLMicroBlogEmpty(nil)
		o.Data2.Constructor = 1794557279
		return o
	},
	-1079755243: func() TLObject { // 0xbfa43e15
		o := MakeTLMicroBlog(nil)
		o.Data2.Constructor = -1079755243
		return o
	},
	383369912: func() TLObject { // 0x16d9c2b8
		o := MakeTLMicroBlog(nil)
		o.Data2.Constructor = 383369912
		return o
	},
	475467473: func() TLObject { // 0x1c570ed1
		o := MakeTLWebDocument(nil)
		o.Data2.Constructor = 475467473
		return o
	},
	-971322408: func() TLObject { // 0xc61acbd8
		o := MakeTLWebDocument(nil)
		o.Data2.Constructor = -971322408
		return o
	},
	-104284986: func() TLObject { // 0xf9c8bcc6
		o := MakeTLWebDocumentNoProxy(nil)
		o.Data2.Constructor = -104284986
		return o
	},
	-177732982: func() TLObject { // 0xf568028a
		o := MakeTLBankCardOpenUrl(nil)
		o.Data2.Constructor = -177732982
		return o
	},
	-1932527041: func() TLObject { // 0x8ccffa3f
		o := MakeTLInt32(nil)
		o.Data2.Constructor = -1932527041
		return o
	},
	307276766: func() TLObject { // 0x1250abde
		o := MakeTLAccountAuthorizations(nil)
		o.Data2.Constructor = 307276766
		return o
	},
	157948117: func() TLObject { // 0x96a18d5
		o := MakeTLUploadFile(nil)
		o.Data2.Constructor = 157948117
		return o
	},
	-242427324: func() TLObject { // 0xf18cda44
		o := MakeTLUploadFileCdnRedirect(nil)
		o.Data2.Constructor = -242427324
		return o
	},
	-363659686: func() TLObject { // 0xea52fe5a
		o := MakeTLUploadFileCdnRedirect(nil)
		o.Data2.Constructor = -363659686
		return o
	},
	856375399: func() TLObject { // 0x330b4067
		o := MakeTLConfig(nil)
		o.Data2.Constructor = 856375399
		return o
	},
	-422959626: func() TLObject { // 0xe6ca25f6
		o := MakeTLConfig(nil)
		o.Data2.Constructor = -422959626
		return o
	},
	840162234: func() TLObject { // 0x3213dbba
		o := MakeTLConfig(nil)
		o.Data2.Constructor = 840162234
		return o
	},
	-344215200: func() TLObject { // 0xeb7bb160
		o := MakeTLConfig(nil)
		o.Data2.Constructor = -344215200
		return o
	},
	-2034927730: func() TLObject { // 0x86b5778e
		o := MakeTLConfig(nil)
		o.Data2.Constructor = -2034927730
		return o
	},
	-1669068444: func() TLObject { // 0x9c840964
		o := MakeTLConfig(nil)
		o.Data2.Constructor = -1669068444
		return o
	},
	1493171408: func() TLObject { // 0x58fffcd0
		o := MakeTLHighScore(nil)
		o.Data2.Constructor = 1493171408
		return o
	},
	1823064809: func() TLObject { // 0x6ca9c2e9
		o := MakeTLPollAnswer(nil)
		o.Data2.Constructor = 1823064809
		return o
	},
	-1432995067: func() TLObject { // 0xaa963b05
		o := MakeTLStorageFileUnknown(nil)
		o.Data2.Constructor = -1432995067
		return o
	},
	1086091090: func() TLObject { // 0x40bc6f52
		o := MakeTLStorageFilePartial(nil)
		o.Data2.Constructor = 1086091090
		return o
	},
	8322574: func() TLObject { // 0x7efe0e
		o := MakeTLStorageFileJpeg(nil)
		o.Data2.Constructor = 8322574
		return o
	},
	-891180321: func() TLObject { // 0xcae1aadf
		o := MakeTLStorageFileGif(nil)
		o.Data2.Constructor = -891180321
		return o
	},
	172975040: func() TLObject { // 0xa4f63c0
		o := MakeTLStorageFilePng(nil)
		o.Data2.Constructor = 172975040
		return o
	},
	-1373745011: func() TLObject { // 0xae1e508d
		o := MakeTLStorageFilePdf(nil)
		o.Data2.Constructor = -1373745011
		return o
	},
	1384777335: func() TLObject { // 0x528a0677
		o := MakeTLStorageFileMp3(nil)
		o.Data2.Constructor = 1384777335
		return o
	},
	1258941372: func() TLObject { // 0x4b09ebbc
		o := MakeTLStorageFileMov(nil)
		o.Data2.Constructor = 1258941372
		return o
	},
	-1278304028: func() TLObject { // 0xb3cea0e4
		o := MakeTLStorageFileMp4(nil)
		o.Data2.Constructor = -1278304028
		return o
	},
	276907596: func() TLObject { // 0x1081464c
		o := MakeTLStorageFileWebp(nil)
		o.Data2.Constructor = 276907596
		return o
	},
	1516793212: func() TLObject { // 0x5a686d7c
		o := MakeTLChatInviteAlready(nil)
		o.Data2.Constructor = 1516793212
		return o
	},
	-540871282: func() TLObject { // 0xdfc2f58e
		o := MakeTLChatInvite(nil)
		o.Data2.Constructor = -540871282
		return o
	},
	-613092008: func() TLObject { // 0xdb74f558
		o := MakeTLChatInvite(nil)
		o.Data2.Constructor = -613092008
		return o
	},
	1634294960: func() TLObject { // 0x61695cb0
		o := MakeTLChatInvitePeek(nil)
		o.Data2.Constructor = 1634294960
		return o
	},
	-599948721: func() TLObject { // 0xdc3d824f
		o := MakeTLTextEmpty(nil)
		o.Data2.Constructor = -599948721
		return o
	},
	1950782688: func() TLObject { // 0x744694e0
		o := MakeTLTextPlain(nil)
		o.Data2.Constructor = 1950782688
		return o
	},
	1730456516: func() TLObject { // 0x6724abc4
		o := MakeTLTextBold(nil)
		o.Data2.Constructor = 1730456516
		return o
	},
	-653089380: func() TLObject { // 0xd912a59c
		o := MakeTLTextItalic(nil)
		o.Data2.Constructor = -653089380
		return o
	},
	-1054465340: func() TLObject { // 0xc12622c4
		o := MakeTLTextUnderline(nil)
		o.Data2.Constructor = -1054465340
		return o
	},
	-1678197867: func() TLObject { // 0x9bf8bb95
		o := MakeTLTextStrike(nil)
		o.Data2.Constructor = -1678197867
		return o
	},
	1816074681: func() TLObject { // 0x6c3f19b9
		o := MakeTLTextFixed(nil)
		o.Data2.Constructor = 1816074681
		return o
	},
	1009288385: func() TLObject { // 0x3c2884c1
		o := MakeTLTextUrl(nil)
		o.Data2.Constructor = 1009288385
		return o
	},
	-564523562: func() TLObject { // 0xde5a0dd6
		o := MakeTLTextEmail(nil)
		o.Data2.Constructor = -564523562
		return o
	},
	2120376535: func() TLObject { // 0x7e6260d7
		o := MakeTLTextConcat(nil)
		o.Data2.Constructor = 2120376535
		return o
	},
	-311786236: func() TLObject { // 0xed6a8504
		o := MakeTLTextSubscript(nil)
		o.Data2.Constructor = -311786236
		return o
	},
	-939827711: func() TLObject { // 0xc7fb5e01
		o := MakeTLTextSuperscript(nil)
		o.Data2.Constructor = -939827711
		return o
	},
	55281185: func() TLObject { // 0x34b8621
		o := MakeTLTextMarked(nil)
		o.Data2.Constructor = 55281185
		return o
	},
	483104362: func() TLObject { // 0x1ccb966a
		o := MakeTLTextPhone(nil)
		o.Data2.Constructor = 483104362
		return o
	},
	136105807: func() TLObject { // 0x81ccf4f
		o := MakeTLTextImage(nil)
		o.Data2.Constructor = 136105807
		return o
	},
	894777186: func() TLObject { // 0x35553762
		o := MakeTLTextAnchor(nil)
		o.Data2.Constructor = 894777186
		return o
	},
	-58224696: func() TLObject { // 0xfc878fc8
		o := MakeTLPhoneCallProtocol(nil)
		o.Data2.Constructor = -58224696
		return o
	},
	-1564789301: func() TLObject { // 0xa2bb35cb
		o := MakeTLPhoneCallProtocol(nil)
		o.Data2.Constructor = -1564789301
		return o
	},
	354925740: func() TLObject { // 0x1527bcac
		o := MakeTLSecureSecretSettings(nil)
		o.Data2.Constructor = 354925740
		return o
	},
	-433014407: func() TLObject { // 0xe630b979
		o := MakeTLInputWallPaper(nil)
		o.Data2.Constructor = -433014407
		return o
	},
	1913199744: func() TLObject { // 0x72091c80
		o := MakeTLInputWallPaperSlug(nil)
		o.Data2.Constructor = 1913199744
		return o
	},
	-2077770836: func() TLObject { // 0x8427bbac
		o := MakeTLInputWallPaperNoFile(nil)
		o.Data2.Constructor = -2077770836
		return o
	},
	-1673717362: func() TLObject { // 0x9c3d198e
		o := MakeTLInputPeerNotifySettings(nil)
		o.Data2.Constructor = -1673717362
		return o
	},
	949182130: func() TLObject { // 0x38935eb2
		o := MakeTLInputPeerNotifySettings(nil)
		o.Data2.Constructor = 949182130
		return o
	},
	364538944: func() TLObject { // 0x15ba6c40
		o := MakeTLMessagesDialogs(nil)
		o.Data2.Constructor = 364538944
		return o
	},
	1910543603: func() TLObject { // 0x71e094f3
		o := MakeTLMessagesDialogsSlice(nil)
		o.Data2.Constructor = 1910543603
		return o
	},
	-253500010: func() TLObject { // 0xf0e3e596
		o := MakeTLMessagesDialogsNotModified(nil)
		o.Data2.Constructor = -253500010
		return o
	},
	-368018716: func() TLObject { // 0xea107ae4
		o := MakeTLChannelAdminLogEventsFilter(nil)
		o.Data2.Constructor = -368018716
		return o
	},
	1189204285: func() TLObject { // 0x46e1d13d
		o := MakeTLRecentMeUrlUnknown(nil)
		o.Data2.Constructor = 1189204285
		return o
	},
	-1917045962: func() TLObject { // 0x8dbc3336
		o := MakeTLRecentMeUrlUser(nil)
		o.Data2.Constructor = -1917045962
		return o
	},
	-1608834311: func() TLObject { // 0xa01b22f9
		o := MakeTLRecentMeUrlChat(nil)
		o.Data2.Constructor = -1608834311
		return o
	},
	-347535331: func() TLObject { // 0xeb49081d
		o := MakeTLRecentMeUrlChatInvite(nil)
		o.Data2.Constructor = -347535331
		return o
	},
	-1140172836: func() TLObject { // 0xbc0a57dc
		o := MakeTLRecentMeUrlStickerSet(nil)
		o.Data2.Constructor = -1140172836
		return o
	},
	-316748368: func() TLObject { // 0xed1ecdb0
		o := MakeTLSecureValueHash(nil)
		o.Data2.Constructor = -316748368
		return o
	},
	-1188055347: func() TLObject { // 0xb92fb6cd
		o := MakeTLPageListItemText(nil)
		o.Data2.Constructor = -1188055347
		return o
	},
	635466748: func() TLObject { // 0x25e073fc
		o := MakeTLPageListItemBlocks(nil)
		o.Data2.Constructor = 635466748
		return o
	},
	1605510357: func() TLObject { // 0x5fb224d5
		o := MakeTLChatAdminRights(nil)
		o.Data2.Constructor = 1605510357
		return o
	},
	-748155807: func() TLObject { // 0xd3680c61
		o := MakeTLContactStatus(nil)
		o.Data2.Constructor = -748155807
		return o
	},
	-2000710887: func() TLObject { // 0x88bf9319
		o := MakeTLInputBotInlineResult(nil)
		o.Data2.Constructor = -2000710887
		return o
	},
	750510426: func() TLObject { // 0x2cbbe15a
		o := MakeTLInputBotInlineResult(nil)
		o.Data2.Constructor = 750510426
		return o
	},
	-1462213465: func() TLObject { // 0xa8d864a7
		o := MakeTLInputBotInlineResultPhoto(nil)
		o.Data2.Constructor = -1462213465
		return o
	},
	-459324: func() TLObject { // 0xfff8fdc4
		o := MakeTLInputBotInlineResultDocument(nil)
		o.Data2.Constructor = -459324
		return o
	},
	1336154098: func() TLObject { // 0x4fa417f2
		o := MakeTLInputBotInlineResultGame(nil)
		o.Data2.Constructor = 1336154098
		return o
	},
	863093588: func() TLObject { // 0x3371c354
		o := MakeTLMessagesPeerDialogs(nil)
		o.Data2.Constructor = 863093588
		return o
	},
	-1868808300: func() TLObject { // 0x909c3f94
		o := MakeTLPaymentRequestedInfo(nil)
		o.Data2.Constructor = -1868808300
		return o
	},
	1648543603: func() TLObject { // 0x6242c773
		o := MakeTLFileHash(nil)
		o.Data2.Constructor = 1648543603
		return o
	},
	-2128640689: func() TLObject { // 0x811f854f
		o := MakeTLAccountSentEmailCode(nil)
		o.Data2.Constructor = -2128640689
		return o
	},
	986597452: func() TLObject { // 0x3ace484c
		o := MakeTLContactsLink(nil)
		o.Data2.Constructor = 986597452
		return o
	},
	-244016606: func() TLObject { // 0xf1749a22
		o := MakeTLMessagesStickersNotModified(nil)
		o.Data2.Constructor = -244016606
		return o
	},
	-463889475: func() TLObject { // 0xe4599bbd
		o := MakeTLMessagesStickers(nil)
		o.Data2.Constructor = -463889475
		return o
	},
	-1970352846: func() TLObject { // 0x8a8ecd32
		o := MakeTLMessagesStickers(nil)
		o.Data2.Constructor = -1970352846
		return o
	},
	2104790276: func() TLObject { // 0x7d748d04
		o := MakeTLDataJSON(nil)
		o.Data2.Constructor = 2104790276
		return o
	},
	-884757282: func() TLObject { // 0xcb43acde
		o := MakeTLStatsAbsValueAndPrev(nil)
		o.Data2.Constructor = -884757282
		return o
	},
	194458693: func() TLObject { // 0xb973445
		o := MakeTLString(nil)
		o.Data2.Constructor = 194458693
		return o
	},
	935395612: func() TLObject { // 0x37c1011c
		o := MakeTLChatPhotoEmpty(nil)
		o.Data2.Constructor = 935395612
		return o
	},
	-770990276: func() TLObject { // 0xd20b9f3c
		o := MakeTLChatPhoto(nil)
		o.Data2.Constructor = -770990276
		return o
	},
	1197267925: func() TLObject { // 0x475cdbd5
		o := MakeTLChatPhoto(nil)
		o.Data2.Constructor = 1197267925
		return o
	},
	1632839530: func() TLObject { // 0x6153276a
		o := MakeTLChatPhoto(nil)
		o.Data2.Constructor = 1632839530
		return o
	},
	-123988: func() TLObject { // 0xfffe1bac
		o := MakeTLPrivacyValueAllowContacts(nil)
		o.Data2.Constructor = -123988
		return o
	},
	1698855810: func() TLObject { // 0x65427b82
		o := MakeTLPrivacyValueAllowAll(nil)
		o.Data2.Constructor = 1698855810
		return o
	},
	1297858060: func() TLObject { // 0x4d5bbe0c
		o := MakeTLPrivacyValueAllowUsers(nil)
		o.Data2.Constructor = 1297858060
		return o
	},
	-125240806: func() TLObject { // 0xf888fa1a
		o := MakeTLPrivacyValueDisallowContacts(nil)
		o.Data2.Constructor = -125240806
		return o
	},
	-1955338397: func() TLObject { // 0x8b73e763
		o := MakeTLPrivacyValueDisallowAll(nil)
		o.Data2.Constructor = -1955338397
		return o
	},
	209668535: func() TLObject { // 0xc7f49b7
		o := MakeTLPrivacyValueDisallowUsers(nil)
		o.Data2.Constructor = 209668535
		return o
	},
	415136107: func() TLObject { // 0x18be796b
		o := MakeTLPrivacyValueAllowChatParticipants(nil)
		o.Data2.Constructor = 415136107
		return o
	},
	-1397881200: func() TLObject { // 0xacae0690
		o := MakeTLPrivacyValueDisallowChatParticipants(nil)
		o.Data2.Constructor = -1397881200
		return o
	},
	-305282981: func() TLObject { // 0xedcdc05b
		o := MakeTLTopPeer(nil)
		o.Data2.Constructor = -305282981
		return o
	},
	1678812626: func() TLObject { // 0x6410a5d2
		o := MakeTLStickerSetCovered(nil)
		o.Data2.Constructor = 1678812626
		return o
	},
	872932635: func() TLObject { // 0x3407e51b
		o := MakeTLStickerSetMultiCovered(nil)
		o.Data2.Constructor = 872932635
		return o
	},
	-326966976: func() TLObject { // 0xec82e140
		o := MakeTLPhonePhoneCall(nil)
		o.Data2.Constructor = -326966976
		return o
	},
	1694474197: func() TLObject { // 0x64ff9fd5
		o := MakeTLMessagesChats(nil)
		o.Data2.Constructor = 1694474197
		return o
	},
	-1663561404: func() TLObject { // 0x9cd81144
		o := MakeTLMessagesChatsSlice(nil)
		o.Data2.Constructor = -1663561404
		return o
	},
	-587149284: func() TLObject { // 0xdd00d01c
		o := MakeTLBlogsUserDatesNotModified(nil)
		o.Data2.Constructor = -587149284
		return o
	},
	376056903: func() TLObject { // 0x166a2c47
		o := MakeTLBlogsUserDates(nil)
		o.Data2.Constructor = 376056903
		return o
	},
	-322927481: func() TLObject { // 0xecc08487
		o := MakeTLBlogsGroupTags(nil)
		o.Data2.Constructor = -322927481
		return o
	},
	-1032140601: func() TLObject { // 0xc27ac8c7
		o := MakeTLBotCommand(nil)
		o.Data2.Constructor = -1032140601
		return o
	},
	2002815875: func() TLObject { // 0x77608b83
		o := MakeTLKeyboardButtonRow(nil)
		o.Data2.Constructor = 2002815875
		return o
	},
	1062645411: func() TLObject { // 0x3f56aea3
		o := MakeTLPaymentsPaymentForm(nil)
		o.Data2.Constructor = 1062645411
		return o
	},
	-1195615476: func() TLObject { // 0xb8bc5b0c
		o := MakeTLInputNotifyPeer(nil)
		o.Data2.Constructor = -1195615476
		return o
	},
	423314455: func() TLObject { // 0x193b4417
		o := MakeTLInputNotifyUsers(nil)
		o.Data2.Constructor = 423314455
		return o
	},
	1251338318: func() TLObject { // 0x4a95e84e
		o := MakeTLInputNotifyChats(nil)
		o.Data2.Constructor = 1251338318
		return o
	},
	-1311015810: func() TLObject { // 0xb1db7c7e
		o := MakeTLInputNotifyBroadcasts(nil)
		o.Data2.Constructor = -1311015810
		return o
	},
	-1540769658: func() TLObject { // 0xa429b886
		o := MakeTLInputNotifyAll(nil)
		o.Data2.Constructor = -1540769658
		return o
	},
	1815593308: func() TLObject { // 0x6c37c15c
		o := MakeTLDocumentAttributeImageSize(nil)
		o.Data2.Constructor = 1815593308
		return o
	},
	297109817: func() TLObject { // 0x11b58939
		o := MakeTLDocumentAttributeAnimated(nil)
		o.Data2.Constructor = 297109817
		return o
	},
	1662637586: func() TLObject { // 0x6319d612
		o := MakeTLDocumentAttributeSticker(nil)
		o.Data2.Constructor = 1662637586
		return o
	},
	250621158: func() TLObject { // 0xef02ce6
		o := MakeTLDocumentAttributeVideo(nil)
		o.Data2.Constructor = 250621158
		return o
	},
	-1739392570: func() TLObject { // 0x9852f9c6
		o := MakeTLDocumentAttributeAudio(nil)
		o.Data2.Constructor = -1739392570
		return o
	},
	358154344: func() TLObject { // 0x15590068
		o := MakeTLDocumentAttributeFilename(nil)
		o.Data2.Constructor = 358154344
		return o
	},
	-1744710921: func() TLObject { // 0x9801d2f7
		o := MakeTLDocumentAttributeHasStickers(nil)
		o.Data2.Constructor = -1744710921
		return o
	},
	1923290508: func() TLObject { // 0x72a3158c
		o := MakeTLAuthCodeTypeSms(nil)
		o.Data2.Constructor = 1923290508
		return o
	},
	1948046307: func() TLObject { // 0x741cd3e3
		o := MakeTLAuthCodeTypeCall(nil)
		o.Data2.Constructor = 1948046307
		return o
	},
	577556219: func() TLObject { // 0x226ccefb
		o := MakeTLAuthCodeTypeFlashCall(nil)
		o.Data2.Constructor = 577556219
		return o
	},
	1558266229: func() TLObject { // 0x5ce14175
		o := MakeTLPopularContact(nil)
		o.Data2.Constructor = 1558266229
		return o
	},
	-484987010: func() TLObject { // 0xe317af7e
		o := MakeTLUpdatesTooLong(nil)
		o.Data2.Constructor = -484987010
		return o
	},
	580309704: func() TLObject { // 0x2296d2c8
		o := MakeTLUpdateShortMessage(nil)
		o.Data2.Constructor = 580309704
		return o
	},
	-1857044719: func() TLObject { // 0x914fbf11
		o := MakeTLUpdateShortMessage(nil)
		o.Data2.Constructor = -1857044719
		return o
	},
	1076714939: func() TLObject { // 0x402d5dbb
		o := MakeTLUpdateShortChatMessage(nil)
		o.Data2.Constructor = 1076714939
		return o
	},
	377562760: func() TLObject { // 0x16812688
		o := MakeTLUpdateShortChatMessage(nil)
		o.Data2.Constructor = 377562760
		return o
	},
	2027216577: func() TLObject { // 0x78d4dec1
		o := MakeTLUpdateShort(nil)
		o.Data2.Constructor = 2027216577
		return o
	},
	1918567619: func() TLObject { // 0x725b04c3
		o := MakeTLUpdatesCombined(nil)
		o.Data2.Constructor = 1918567619
		return o
	},
	1957577280: func() TLObject { // 0x74ae4240
		o := MakeTLUpdates(nil)
		o.Data2.Constructor = 1957577280
		return o
	},
	301019932: func() TLObject { // 0x11f1331c
		o := MakeTLUpdateShortSentMessage(nil)
		o.Data2.Constructor = 301019932
		return o
	},
	618817830: func() TLObject { // 0x24e26926
		o := MakeTLUpdateAccountResetAuthorization(nil)
		o.Data2.Constructor = 618817830
		return o
	},
	-402498398: func() TLObject { // 0xe8025ca2
		o := MakeTLMessagesSavedGifsNotModified(nil)
		o.Data2.Constructor = -402498398
		return o
	},
	772213157: func() TLObject { // 0x2e0709a5
		o := MakeTLMessagesSavedGifs(nil)
		o.Data2.Constructor = 772213157
		return o
	},
	235081943: func() TLObject { // 0xe0310d7
		o := MakeTLHelpRecentMeUrls(nil)
		o.Data2.Constructor = 235081943
		return o
	},
	-55902537: func() TLObject { // 0xfcaafeb7
		o := MakeTLInputDialogPeer(nil)
		o.Data2.Constructor = -55902537
		return o
	},
	1684014375: func() TLObject { // 0x64600527
		o := MakeTLInputDialogPeerFolder(nil)
		o.Data2.Constructor = 1684014375
		return o
	},
	741914831: func() TLObject { // 0x2c38b8cf
		o := MakeTLInputDialogPeerFeed(nil)
		o.Data2.Constructor = 741914831
		return o
	},
	289586518: func() TLObject { // 0x1142bd56
		o := MakeTLSavedPhoneContact(nil)
		o.Data2.Constructor = 289586518
		return o
	},
	-386039788: func() TLObject { // 0xe8fd8014
		o := MakeTLPeerBlocked(nil)
		o.Data2.Constructor = -386039788
		return o
	},
	1330637553: func() TLObject { // 0x4f4feaf1
		o := MakeTLFeedBroadcasts(nil)
		o.Data2.Constructor = 1330637553
		return o
	},
	-1704428358: func() TLObject { // 0x9a687cba
		o := MakeTLFeedBroadcastsUngrouped(nil)
		o.Data2.Constructor = -1704428358
		return o
	},
	236446268: func() TLObject { // 0xe17e23c
		o := MakeTLPhotoSizeEmpty(nil)
		o.Data2.Constructor = 236446268
		return o
	},
	2009052699: func() TLObject { // 0x77bfb61b
		o := MakeTLPhotoSize(nil)
		o.Data2.Constructor = 2009052699
		return o
	},
	-374917894: func() TLObject { // 0xe9a734fa
		o := MakeTLPhotoCachedSize(nil)
		o.Data2.Constructor = -374917894
		return o
	},
	-525288402: func() TLObject { // 0xe0b0bc2e
		o := MakeTLPhotoStrippedSize(nil)
		o.Data2.Constructor = -525288402
		return o
	},
	1520986705: func() TLObject { // 0x5aa86a51
		o := MakeTLPhotoSizeProgressive(nil)
		o.Data2.Constructor = 1520986705
		return o
	},
	182326673: func() TLObject { // 0xade1591
		o := MakeTLContactsBlocked(nil)
		o.Data2.Constructor = 182326673
		return o
	},
	471043349: func() TLObject { // 0x1c138d15
		o := MakeTLContactsBlocked(nil)
		o.Data2.Constructor = 471043349
		return o
	},
	-513392236: func() TLObject { // 0xe1664194
		o := MakeTLContactsBlockedSlice(nil)
		o.Data2.Constructor = -513392236
		return o
	},
	-1878523231: func() TLObject { // 0x900802a1
		o := MakeTLContactsBlockedSlice(nil)
		o.Data2.Constructor = -1878523231
		return o
	},
	-4838507: func() TLObject { // 0xffb62b95
		o := MakeTLInputStickerSetEmpty(nil)
		o.Data2.Constructor = -4838507
		return o
	},
	-1645763991: func() TLObject { // 0x9de7a269
		o := MakeTLInputStickerSetID(nil)
		o.Data2.Constructor = -1645763991
		return o
	},
	-2044933984: func() TLObject { // 0x861cc8a0
		o := MakeTLInputStickerSetShortName(nil)
		o.Data2.Constructor = -2044933984
		return o
	},
	42402760: func() TLObject { // 0x28703c8
		o := MakeTLInputStickerSetAnimatedEmoji(nil)
		o.Data2.Constructor = 42402760
		return o
	},
	-427863538: func() TLObject { // 0xe67f520e
		o := MakeTLInputStickerSetDice(nil)
		o.Data2.Constructor = -427863538
		return o
	},
	2044861011: func() TLObject { // 0x79e21a53
		o := MakeTLInputStickerSetDice(nil)
		o.Data2.Constructor = 2044861011
		return o
	},
	1399245077: func() TLObject { // 0x5366c915
		o := MakeTLPhoneCallEmpty(nil)
		o.Data2.Constructor = 1399245077
		return o
	},
	462375633: func() TLObject { // 0x1b8f4ad1
		o := MakeTLPhoneCallWaiting(nil)
		o.Data2.Constructor = 462375633
		return o
	},
	-2014659757: func() TLObject { // 0x87eabb53
		o := MakeTLPhoneCallRequested(nil)
		o.Data2.Constructor = -2014659757
		return o
	},
	-2089411356: func() TLObject { // 0x83761ce4
		o := MakeTLPhoneCallRequested(nil)
		o.Data2.Constructor = -2089411356
		return o
	},
	-1719909046: func() TLObject { // 0x997c454a
		o := MakeTLPhoneCallAccepted(nil)
		o.Data2.Constructor = -1719909046
		return o
	},
	1828732223: func() TLObject { // 0x6d003d3f
		o := MakeTLPhoneCallAccepted(nil)
		o.Data2.Constructor = 1828732223
		return o
	},
	-2025673089: func() TLObject { // 0x8742ae7f
		o := MakeTLPhoneCall(nil)
		o.Data2.Constructor = -2025673089
		return o
	},
	-419832333: func() TLObject { // 0xe6f9ddf3
		o := MakeTLPhoneCall(nil)
		o.Data2.Constructor = -419832333
		return o
	},
	-1660057: func() TLObject { // 0xffe6ab67
		o := MakeTLPhoneCall(nil)
		o.Data2.Constructor = -1660057
		return o
	},
	1355435489: func() TLObject { // 0x50ca4de1
		o := MakeTLPhoneCallDiscarded(nil)
		o.Data2.Constructor = 1355435489
		return o
	},
	-2032041631: func() TLObject { // 0x86e18161
		o := MakeTLPoll(nil)
		o.Data2.Constructor = -2032041631
		return o
	},
	-716006138: func() TLObject { // 0xd5529d06
		o := MakeTLPoll(nil)
		o.Data2.Constructor = -716006138
		return o
	},
	-901375139: func() TLObject { // 0xca461b5d
		o := MakeTLPeerLocated(nil)
		o.Data2.Constructor = -901375139
		return o
	},
	-118740917: func() TLObject { // 0xf8ec284b
		o := MakeTLPeerSelfLocated(nil)
		o.Data2.Constructor = -118740917
		return o
	},
	-1676371894: func() TLObject { // 0x9c14984a
		o := MakeTLThemeSettings(nil)
		o.Data2.Constructor = -1676371894
		return o
	},
	739712882: func() TLObject { // 0x2c171f72
		o := MakeTLDialog(nil)
		o.Data2.Constructor = 739712882
		return o
	},
	-455150117: func() TLObject { // 0xe4def5db
		o := MakeTLDialog(nil)
		o.Data2.Constructor = -455150117
		return o
	},
	1908216652: func() TLObject { // 0x71bd134c
		o := MakeTLDialogFolder(nil)
		o.Data2.Constructor = 1908216652
		return o
	},
	906521922: func() TLObject { // 0x36086d42
		o := MakeTLDialogFeed(nil)
		o.Data2.Constructor = 906521922
		return o
	},
	-1058912715: func() TLObject { // 0xc0e24635
		o := MakeTLMessagesDhConfigNotModified(nil)
		o.Data2.Constructor = -1058912715
		return o
	},
	740433629: func() TLObject { // 0x2c221edd
		o := MakeTLMessagesDhConfig(nil)
		o.Data2.Constructor = 740433629
		return o
	},
	-1059185703: func() TLObject { // 0xc0de1bd9
		o := MakeTLJsonObjectValue(nil)
		o.Data2.Constructor = -1059185703
		return o
	},
	565550063: func() TLObject { // 0x21b59bef
		o := MakeTLSchemeParam(nil)
		o.Data2.Constructor = 565550063
		return o
	},
	-170975004: func() TLObject { // 0xf5cf20e4
		o := MakeTLWalletRecordTypeUnknown(nil)
		o.Data2.Constructor = -170975004
		return o
	},
	1692388072: func() TLObject { // 0x64dfcae8
		o := MakeTLWalletRecordTypeManual(nil)
		o.Data2.Constructor = 1692388072
		return o
	},
	-1602911425: func() TLObject { // 0xa075833f
		o := MakeTLWalletRecordTypeRecharge(nil)
		o.Data2.Constructor = -1602911425
		return o
	},
	981005824: func() TLObject { // 0x3a78f600
		o := MakeTLWalletRecordTypeWithdrawal(nil)
		o.Data2.Constructor = 981005824
		return o
	},
	-341632063: func() TLObject { // 0xeba31bc1
		o := MakeTLWalletRecordTypeWithdrawRefunded(nil)
		o.Data2.Constructor = -341632063
		return o
	},
	315650517: func() TLObject { // 0x12d071d5
		o := MakeTLWalletRecordTypeTransferIn(nil)
		o.Data2.Constructor = 315650517
		return o
	},
	-1517213532: func() TLObject { // 0xa59128a4
		o := MakeTLWalletRecordTypeTransferOut(nil)
		o.Data2.Constructor = -1517213532
		return o
	},
	1742991292: func() TLObject { // 0x67e3efbc
		o := MakeTLWalletRecordTypeRedpacketDeduct(nil)
		o.Data2.Constructor = 1742991292
		return o
	},
	239697919: func() TLObject { // 0xe497fff
		o := MakeTLWalletRecordTypeRedpacketGot(nil)
		o.Data2.Constructor = 239697919
		return o
	},
	-83819777: func() TLObject { // 0xfb0102ff
		o := MakeTLWalletRecordTypeRedpacketRefund(nil)
		o.Data2.Constructor = -83819777
		return o
	},
	-826266136: func() TLObject { // 0xcec02de8
		o := MakeTLWalletRecordTypeBlogReward(nil)
		o.Data2.Constructor = -826266136
		return o
	},
	230138437: func() TLObject { // 0xdb7a245
		o := MakeTLWalletRecordTypeBlogGotReward(nil)
		o.Data2.Constructor = 230138437
		return o
	},
	-2014519914: func() TLObject { // 0x87ecdd96
		o := MakeTLWalletRecordTypeRemittanceRemit(nil)
		o.Data2.Constructor = -2014519914
		return o
	},
	-1627947689: func() TLObject { // 0x9ef77d57
		o := MakeTLWalletRecordTypeRemittanceReceive(nil)
		o.Data2.Constructor = -1627947689
		return o
	},
	-1837805826: func() TLObject { // 0x92754efe
		o := MakeTLWalletRecordTypeRemittanceRefund(nil)
		o.Data2.Constructor = -1837805826
		return o
	},
	1474492012: func() TLObject { // 0x57e2f66c
		o := MakeTLInputMessagesFilterEmpty(nil)
		o.Data2.Constructor = 1474492012
		return o
	},
	-1777752804: func() TLObject { // 0x9609a51c
		o := MakeTLInputMessagesFilterPhotos(nil)
		o.Data2.Constructor = -1777752804
		return o
	},
	-1614803355: func() TLObject { // 0x9fc00e65
		o := MakeTLInputMessagesFilterVideo(nil)
		o.Data2.Constructor = -1614803355
		return o
	},
	1458172132: func() TLObject { // 0x56e9f0e4
		o := MakeTLInputMessagesFilterPhotoVideo(nil)
		o.Data2.Constructor = 1458172132
		return o
	},
	-1629621880: func() TLObject { // 0x9eddf188
		o := MakeTLInputMessagesFilterDocument(nil)
		o.Data2.Constructor = -1629621880
		return o
	},
	2129714567: func() TLObject { // 0x7ef0dd87
		o := MakeTLInputMessagesFilterUrl(nil)
		o.Data2.Constructor = 2129714567
		return o
	},
	-3644025: func() TLObject { // 0xffc86587
		o := MakeTLInputMessagesFilterGif(nil)
		o.Data2.Constructor = -3644025
		return o
	},
	1358283666: func() TLObject { // 0x50f5c392
		o := MakeTLInputMessagesFilterVoice(nil)
		o.Data2.Constructor = 1358283666
		return o
	},
	928101534: func() TLObject { // 0x3751b49e
		o := MakeTLInputMessagesFilterMusic(nil)
		o.Data2.Constructor = 928101534
		return o
	},
	975236280: func() TLObject { // 0x3a20ecb8
		o := MakeTLInputMessagesFilterChatPhotos(nil)
		o.Data2.Constructor = 975236280
		return o
	},
	-2134272152: func() TLObject { // 0x80c99768
		o := MakeTLInputMessagesFilterPhoneCalls(nil)
		o.Data2.Constructor = -2134272152
		return o
	},
	2054952868: func() TLObject { // 0x7a7c17a4
		o := MakeTLInputMessagesFilterRoundVoice(nil)
		o.Data2.Constructor = 2054952868
		return o
	},
	-1253451181: func() TLObject { // 0xb549da53
		o := MakeTLInputMessagesFilterRoundVideo(nil)
		o.Data2.Constructor = -1253451181
		return o
	},
	-1040652646: func() TLObject { // 0xc1f8e69a
		o := MakeTLInputMessagesFilterMyMentions(nil)
		o.Data2.Constructor = -1040652646
		return o
	},
	-419271411: func() TLObject { // 0xe7026d0d
		o := MakeTLInputMessagesFilterGeo(nil)
		o.Data2.Constructor = -419271411
		return o
	},
	-530392189: func() TLObject { // 0xe062db83
		o := MakeTLInputMessagesFilterContacts(nil)
		o.Data2.Constructor = -530392189
		return o
	},
	464520273: func() TLObject { // 0x1bb00451
		o := MakeTLInputMessagesFilterPinned(nil)
		o.Data2.Constructor = 464520273
		return o
	},
	-1334848398: func() TLObject { // 0xb06fd472
		o := MakeTLWalletInfo(nil)
		o.Data2.Constructor = -1334848398
		return o
	},
	-1290580579: func() TLObject { // 0xb3134d9d
		o := MakeTLContactsFound(nil)
		o.Data2.Constructor = -1290580579
		return o
	},
	453805082: func() TLObject { // 0x1b0c841a
		o := MakeTLDraftMessageEmpty(nil)
		o.Data2.Constructor = 453805082
		return o
	},
	-1169445179: func() TLObject { // 0xba4baec5
		o := MakeTLDraftMessageEmpty(nil)
		o.Data2.Constructor = -1169445179
		return o
	},
	-40996577: func() TLObject { // 0xfd8e711f
		o := MakeTLDraftMessage(nil)
		o.Data2.Constructor = -40996577
		return o
	},
	-209337866: func() TLObject { // 0xf385c1f6
		o := MakeTLLangPackDifference(nil)
		o.Data2.Constructor = -209337866
		return o
	},
	1444661369: func() TLObject { // 0x561bc879
		o := MakeTLContactBlocked(nil)
		o.Data2.Constructor = 1444661369
		return o
	},
	-543777747: func() TLObject { // 0xdf969c2d
		o := MakeTLAuthExportedAuthorization(nil)
		o.Data2.Constructor = -543777747
		return o
	},
	181652051: func() TLObject { // 0xad3ca53
		o := MakeTLWalletRecordsNotModified(nil)
		o.Data2.Constructor = 181652051
		return o
	},
	2035462031: func() TLObject { // 0x7952af8f
		o := MakeTLWalletRecords(nil)
		o.Data2.Constructor = 2035462031
		return o
	},
	-1417756512: func() TLObject { // 0xab7ec0a0
		o := MakeTLEncryptedChatEmpty(nil)
		o.Data2.Constructor = -1417756512
		return o
	},
	1006044124: func() TLObject { // 0x3bf703dc
		o := MakeTLEncryptedChatWaiting(nil)
		o.Data2.Constructor = 1006044124
		return o
	},
	1651608194: func() TLObject { // 0x62718a82
		o := MakeTLEncryptedChatRequested(nil)
		o.Data2.Constructor = 1651608194
		return o
	},
	-931638658: func() TLObject { // 0xc878527e
		o := MakeTLEncryptedChatRequested(nil)
		o.Data2.Constructor = -931638658
		return o
	},
	-94974410: func() TLObject { // 0xfa56ce36
		o := MakeTLEncryptedChat(nil)
		o.Data2.Constructor = -94974410
		return o
	},
	332848423: func() TLObject { // 0x13d6dd27
		o := MakeTLEncryptedChatDiscarded(nil)
		o.Data2.Constructor = 332848423
		return o
	},
	326715557: func() TLObject { // 0x137948a5
		o := MakeTLAuthPasswordRecovery(nil)
		o.Data2.Constructor = 326715557
		return o
	},
	-1419371685: func() TLObject { // 0xab661b5b
		o := MakeTLTopPeerCategoryBotsPM(nil)
		o.Data2.Constructor = -1419371685
		return o
	},
	344356834: func() TLObject { // 0x148677e2
		o := MakeTLTopPeerCategoryBotsInline(nil)
		o.Data2.Constructor = 344356834
		return o
	},
	104314861: func() TLObject { // 0x637b7ed
		o := MakeTLTopPeerCategoryCorrespondents(nil)
		o.Data2.Constructor = 104314861
		return o
	},
	-1122524854: func() TLObject { // 0xbd17a14a
		o := MakeTLTopPeerCategoryGroups(nil)
		o.Data2.Constructor = -1122524854
		return o
	},
	371037736: func() TLObject { // 0x161d9628
		o := MakeTLTopPeerCategoryChannels(nil)
		o.Data2.Constructor = 371037736
		return o
	},
	511092620: func() TLObject { // 0x1e76a78c
		o := MakeTLTopPeerCategoryPhoneCalls(nil)
		o.Data2.Constructor = 511092620
		return o
	},
	-1472172887: func() TLObject { // 0xa8406ca9
		o := MakeTLTopPeerCategoryForwardUsers(nil)
		o.Data2.Constructor = -1472172887
		return o
	},
	-68239120: func() TLObject { // 0xfbeec0f0
		o := MakeTLTopPeerCategoryForwardChats(nil)
		o.Data2.Constructor = -68239120
		return o
	},
	-958657434: func() TLObject { // 0xc6dc0c66
		o := MakeTLMessagesFeaturedStickersNotModified(nil)
		o.Data2.Constructor = -958657434
		return o
	},
	82699215: func() TLObject { // 0x4ede3cf
		o := MakeTLMessagesFeaturedStickersNotModified(nil)
		o.Data2.Constructor = 82699215
		return o
	},
	-1230257343: func() TLObject { // 0xb6abc341
		o := MakeTLMessagesFeaturedStickers(nil)
		o.Data2.Constructor = -1230257343
		return o
	},
	-123893531: func() TLObject { // 0xf89d88e5
		o := MakeTLMessagesFeaturedStickers(nil)
		o.Data2.Constructor = -123893531
		return o
	},
	-892239370: func() TLObject { // 0xcad181f6
		o := MakeTLLangPackString(nil)
		o.Data2.Constructor = -892239370
		return o
	},
	1816636575: func() TLObject { // 0x6c47ac9f
		o := MakeTLLangPackStringPluralized(nil)
		o.Data2.Constructor = 1816636575
		return o
	},
	695856818: func() TLObject { // 0x2979eeb2
		o := MakeTLLangPackStringDeleted(nil)
		o.Data2.Constructor = 695856818
		return o
	},
	1421174295: func() TLObject { // 0x54b56617
		o := MakeTLWebPageAttributeTheme(nil)
		o.Data2.Constructor = 1421174295
		return o
	},
	-855308010: func() TLObject { // 0xcd050916
		o := MakeTLAuthAuthorization(nil)
		o.Data2.Constructor = -855308010
		return o
	},
	1148485274: func() TLObject { // 0x44747e9a
		o := MakeTLAuthAuthorizationSignUpRequired(nil)
		o.Data2.Constructor = 1148485274
		return o
	},
	1489977929: func() TLObject { // 0x58cf4249
		o := MakeTLChannelBannedRights(nil)
		o.Data2.Constructor = 1489977929
		return o
	},
	-170029155: func() TLObject { // 0xf5dd8f9d
		o := MakeTLMessagesDiscussionMessage(nil)
		o.Data2.Constructor = -170029155
		return o
	},
	-266886085: func() TLObject { // 0xf017a43b
		o := MakeTLBlogsCommentNotModified(nil)
		o.Data2.Constructor = -266886085
		return o
	},
	-1463061087: func() TLObject { // 0xa8cb75a1
		o := MakeTLBlogsComments(nil)
		o.Data2.Constructor = -1463061087
		return o
	},
	-1729618630: func() TLObject { // 0x98e81d3a
		o := MakeTLBotInfo(nil)
		o.Data2.Constructor = -1729618630
		return o
	},
	-309659827: func() TLObject { // 0xed8af74d
		o := MakeTLChannelsAdminLogResults(nil)
		o.Data2.Constructor = -309659827
		return o
	},
	-2103600678: func() TLObject { // 0x829d99da
		o := MakeTLSecureRequiredType(nil)
		o.Data2.Constructor = -2103600678
		return o
	},
	41187252: func() TLObject { // 0x27477b4
		o := MakeTLSecureRequiredTypeOneOf(nil)
		o.Data2.Constructor = 41187252
		return o
	},
	84438264: func() TLObject { // 0x5086cf8
		o := MakeTLWallPaperSettings(nil)
		o.Data2.Constructor = 84438264
		return o
	},
	-1590738760: func() TLObject { // 0xa12f40b8
		o := MakeTLWallPaperSettings(nil)
		o.Data2.Constructor = -1590738760
		return o
	},
	818327999: func() TLObject { // 0x30c6b1bf
		o := MakeTLBlogsUser(nil)
		o.Data2.Constructor = 818327999
		return o
	},
	-350980120: func() TLObject { // 0xeb1477e8
		o := MakeTLWebPageEmpty(nil)
		o.Data2.Constructor = -350980120
		return o
	},
	-981018084: func() TLObject { // 0xc586da1c
		o := MakeTLWebPagePending(nil)
		o.Data2.Constructor = -981018084
		return o
	},
	-392411726: func() TLObject { // 0xe89c45b2
		o := MakeTLWebPage(nil)
		o.Data2.Constructor = -392411726
		return o
	},
	-94051982: func() TLObject { // 0xfa64e172
		o := MakeTLWebPage(nil)
		o.Data2.Constructor = -94051982
		return o
	},
	1594340540: func() TLObject { // 0x5f07b4bc
		o := MakeTLWebPage(nil)
		o.Data2.Constructor = 1594340540
		return o
	},
	1930545681: func() TLObject { // 0x7311ca11
		o := MakeTLWebPageNotModified(nil)
		o.Data2.Constructor = 1930545681
		return o
	},
	-2054908813: func() TLObject { // 0x85849473
		o := MakeTLWebPageNotModified(nil)
		o.Data2.Constructor = -2054908813
		return o
	},
	2013922064: func() TLObject { // 0x780a0310
		o := MakeTLHelpTermsOfService(nil)
		o.Data2.Constructor = 2013922064
		return o
	},
	-236044656: func() TLObject { // 0xf1ee3e90
		o := MakeTLHelpTermsOfService(nil)
		o.Data2.Constructor = -236044656
		return o
	},
	-373643672: func() TLObject { // 0xe9baa668
		o := MakeTLFolderPeer(nil)
		o.Data2.Constructor = -373643672
		return o
	},
	1042605427: func() TLObject { // 0x3e24e573
		o := MakeTLPaymentsBankCardData(nil)
		o.Data2.Constructor = 1042605427
		return o
	},
	-2001655273: func() TLObject { // 0x88b12a17
		o := MakeTLChannelsFeedSourcesNotModified(nil)
		o.Data2.Constructor = -2001655273
		return o
	},
	-1903441347: func() TLObject { // 0x8e8bca3d
		o := MakeTLChannelsFeedSources(nil)
		o.Data2.Constructor = -1903441347
		return o
	},
	-1369215196: func() TLObject { // 0xae636f24
		o := MakeTLDisabledFeature(nil)
		o.Data2.Constructor = -1369215196
		return o
	},
	398898678: func() TLObject { // 0x17c6b5f6
		o := MakeTLHelpSupport(nil)
		o.Data2.Constructor = 398898678
		return o
	},
	-1655957568: func() TLObject { // 0x9d4c17c0
		o := MakeTLPhoneConnection(nil)
		o.Data2.Constructor = -1655957568
		return o
	},
	1667228533: func() TLObject { // 0x635fe375
		o := MakeTLPhoneConnectionWebrtc(nil)
		o.Data2.Constructor = 1667228533
		return o
	},
	-1910892683: func() TLObject { // 0x8e1a1775
		o := MakeTLNearestDc(nil)
		o.Data2.Constructor = -1910892683
		return o
	},
	2134579434: func() TLObject { // 0x7f3b18ea
		o := MakeTLInputPeerEmpty(nil)
		o.Data2.Constructor = 2134579434
		return o
	},
	2107670217: func() TLObject { // 0x7da07ec9
		o := MakeTLInputPeerSelf(nil)
		o.Data2.Constructor = 2107670217
		return o
	},
	396093539: func() TLObject { // 0x179be863
		o := MakeTLInputPeerChat(nil)
		o.Data2.Constructor = 396093539
		return o
	},
	2072935910: func() TLObject { // 0x7b8e7de6
		o := MakeTLInputPeerUser(nil)
		o.Data2.Constructor = 2072935910
		return o
	},
	548253432: func() TLObject { // 0x20adaef8
		o := MakeTLInputPeerChannel(nil)
		o.Data2.Constructor = 548253432
		return o
	},
	398123750: func() TLObject { // 0x17bae2e6
		o := MakeTLInputPeerUserFromMessage(nil)
		o.Data2.Constructor = 398123750
		return o
	},
	-1667893317: func() TLObject { // 0x9c95f7bb
		o := MakeTLInputPeerChannelFromMessage(nil)
		o.Data2.Constructor = -1667893317
		return o
	},
	-88014124: func() TLObject { // 0xfac102d4
		o := MakeTLInputPeerUsername(nil)
		o.Data2.Constructor = -88014124
		return o
	},
	-1036572727: func() TLObject { // 0xc23727c9
		o := MakeTLAccountPasswordInputSettings(nil)
		o.Data2.Constructor = -1036572727
		return o
	},
	570402317: func() TLObject { // 0x21ffa60d
		o := MakeTLAccountPasswordInputSettings(nil)
		o.Data2.Constructor = 570402317
		return o
	},
	-2037289493: func() TLObject { // 0x86916deb
		o := MakeTLAccountPasswordInputSettings(nil)
		o.Data2.Constructor = -2037289493
		return o
	},
	946083368: func() TLObject { // 0x38641628
		o := MakeTLMessagesStickerSetInstallResultSuccess(nil)
		o.Data2.Constructor = 946083368
		return o
	},
	904138920: func() TLObject { // 0x35e410a8
		o := MakeTLMessagesStickerSetInstallResultArchive(nil)
		o.Data2.Constructor = 904138920
		return o
	},
	-1228606141: func() TLObject { // 0xb6c4f543
		o := MakeTLMessagesMessageViews(nil)
		o.Data2.Constructor = -1228606141
		return o
	},
	1450380236: func() TLObject { // 0x56730bcc
		o := MakeTLNull(nil)
		o.Data2.Constructor = 1450380236
		return o
	},
	-1560655744: func() TLObject { // 0xa2fa4880
		o := MakeTLKeyboardButton(nil)
		o.Data2.Constructor = -1560655744
		return o
	},
	629866245: func() TLObject { // 0x258aff05
		o := MakeTLKeyboardButtonUrl(nil)
		o.Data2.Constructor = 629866245
		return o
	},
	901503851: func() TLObject { // 0x35bbdb6b
		o := MakeTLKeyboardButtonCallback(nil)
		o.Data2.Constructor = 901503851
		return o
	},
	1748655686: func() TLObject { // 0x683a5e46
		o := MakeTLKeyboardButtonCallback(nil)
		o.Data2.Constructor = 1748655686
		return o
	},
	-1318425559: func() TLObject { // 0xb16a6c29
		o := MakeTLKeyboardButtonRequestPhone(nil)
		o.Data2.Constructor = -1318425559
		return o
	},
	-59151553: func() TLObject { // 0xfc796b3f
		o := MakeTLKeyboardButtonRequestGeoLocation(nil)
		o.Data2.Constructor = -59151553
		return o
	},
	90744648: func() TLObject { // 0x568a748
		o := MakeTLKeyboardButtonSwitchInline(nil)
		o.Data2.Constructor = 90744648
		return o
	},
	1358175439: func() TLObject { // 0x50f41ccf
		o := MakeTLKeyboardButtonGame(nil)
		o.Data2.Constructor = 1358175439
		return o
	},
	-1344716869: func() TLObject { // 0xafd93fbb
		o := MakeTLKeyboardButtonBuy(nil)
		o.Data2.Constructor = -1344716869
		return o
	},
	280464681: func() TLObject { // 0x10b78d29
		o := MakeTLKeyboardButtonUrlAuth(nil)
		o.Data2.Constructor = 280464681
		return o
	},
	-802258988: func() TLObject { // 0xd02e7fd4
		o := MakeTLInputKeyboardButtonUrlAuth(nil)
		o.Data2.Constructor = -802258988
		return o
	},
	-1144565411: func() TLObject { // 0xbbc7515d
		o := MakeTLKeyboardButtonRequestPoll(nil)
		o.Data2.Constructor = -1144565411
		return o
	},
	-265263912: func() TLObject { // 0xf03064d8
		o := MakeTLInputPeerNotifyEventsEmpty(nil)
		o.Data2.Constructor = -265263912
		return o
	},
	-395694988: func() TLObject { // 0xe86a2c74
		o := MakeTLInputPeerNotifyEventsAll(nil)
		o.Data2.Constructor = -395694988
		return o
	},
	-1219778094: func() TLObject { // 0xb74ba9d2
		o := MakeTLContactsContactsNotModified(nil)
		o.Data2.Constructor = -1219778094
		return o
	},
	-353862078: func() TLObject { // 0xeae87e42
		o := MakeTLContactsContacts(nil)
		o.Data2.Constructor = -353862078
		return o
	},
	-1613493288: func() TLObject { // 0x9fd40bd8
		o := MakeTLNotifyPeer(nil)
		o.Data2.Constructor = -1613493288
		return o
	},
	-1261946036: func() TLObject { // 0xb4c83b4c
		o := MakeTLNotifyUsers(nil)
		o.Data2.Constructor = -1261946036
		return o
	},
	-1073230141: func() TLObject { // 0xc007cec3
		o := MakeTLNotifyChats(nil)
		o.Data2.Constructor = -1073230141
		return o
	},
	-703403793: func() TLObject { // 0xd612e8ef
		o := MakeTLNotifyBroadcasts(nil)
		o.Data2.Constructor = -703403793
		return o
	},
	1959820384: func() TLObject { // 0x74d07c60
		o := MakeTLNotifyAll(nil)
		o.Data2.Constructor = 1959820384
		return o
	},
	488313413: func() TLObject { // 0x1d1b1245
		o := MakeTLInputAppEvent(nil)
		o.Data2.Constructor = 488313413
		return o
	},
	1996904104: func() TLObject { // 0x770656a8
		o := MakeTLInputAppEvent(nil)
		o.Data2.Constructor = 1996904104
		return o
	},
	-1282352120: func() TLObject { // 0xb390dc08
		o := MakeTLPageRelatedArticle(nil)
		o.Data2.Constructor = -1282352120
		return o
	},
	-242812612: func() TLObject { // 0xf186f93c
		o := MakeTLPageRelatedArticle(nil)
		o.Data2.Constructor = -242812612
		return o
	},
	-1237848657: func() TLObject { // 0xb637edaf
		o := MakeTLStatsDateRangeDays(nil)
		o.Data2.Constructor = -1237848657
		return o
	},
	-1938715001: func() TLObject { // 0x8c718e87
		o := MakeTLMessagesMessages(nil)
		o.Data2.Constructor = -1938715001
		return o
	},
	978610270: func() TLObject { // 0x3a54685e
		o := MakeTLMessagesMessagesSlice(nil)
		o.Data2.Constructor = 978610270
		return o
	},
	-923939298: func() TLObject { // 0xc8edce1e
		o := MakeTLMessagesMessagesSlice(nil)
		o.Data2.Constructor = -923939298
		return o
	},
	-1497072982: func() TLObject { // 0xa6c47aaa
		o := MakeTLMessagesMessagesSlice(nil)
		o.Data2.Constructor = -1497072982
		return o
	},
	189033187: func() TLObject { // 0xb446ae3
		o := MakeTLMessagesMessagesSlice(nil)
		o.Data2.Constructor = 189033187
		return o
	},
	1682413576: func() TLObject { // 0x64479808
		o := MakeTLMessagesChannelMessages(nil)
		o.Data2.Constructor = 1682413576
		return o
	},
	-1725551049: func() TLObject { // 0x99262e37
		o := MakeTLMessagesChannelMessages(nil)
		o.Data2.Constructor = -1725551049
		return o
	},
	1951620897: func() TLObject { // 0x74535f21
		o := MakeTLMessagesMessagesNotModified(nil)
		o.Data2.Constructor = 1951620897
		return o
	},
	-784000893: func() TLObject { // 0xd1451883
		o := MakeTLPaymentsValidatedRequestedInfo(nil)
		o.Data2.Constructor = -784000893
		return o
	},
	-914167110: func() TLObject { // 0xc982eaba
		o := MakeTLCdnPublicKey(nil)
		o.Data2.Constructor = -914167110
		return o
	},
	482797855: func() TLObject { // 0x1cc6e91f
		o := MakeTLInputSingleMedia(nil)
		o.Data2.Constructor = 482797855
		return o
	},
	1588230153: func() TLObject { // 0x5eaa7809
		o := MakeTLInputSingleMedia(nil)
		o.Data2.Constructor = 1588230153
		return o
	},
	-1567730343: func() TLObject { // 0xa28e5559
		o := MakeTLMessageUserVote(nil)
		o.Data2.Constructor = -1567730343
		return o
	},
	909603888: func() TLObject { // 0x36377430
		o := MakeTLMessageUserVoteInputOption(nil)
		o.Data2.Constructor = 909603888
		return o
	},
	244310238: func() TLObject { // 0xe8fe0de
		o := MakeTLMessageUserVoteMultiple(nil)
		o.Data2.Constructor = 244310238
		return o
	},
	1253220205: func() TLObject { // 0x4ab29f6d
		o := MakeTLLong(nil)
		o.Data2.Constructor = 1253220205
		return o
	},
	-1568590240: func() TLObject { // 0xa2813660
		o := MakeTLInt64(nil)
		o.Data2.Constructor = -1568590240
		return o
	},
	-116274796: func() TLObject { // 0xf911c994
		o := MakeTLContact(nil)
		o.Data2.Constructor = -116274796
		return o
	},
	-2082087340: func() TLObject { // 0x83e5de54
		o := MakeTLMessageEmpty(nil)
		o.Data2.Constructor = -2082087340
		return o
	},
	-38382484: func() TLObject { // 0xfdb6546c
		o := MakeTLMessage(nil)
		o.Data2.Constructor = -38382484
		return o
	},
	1487813065: func() TLObject { // 0x58ae39c9
		o := MakeTLMessage(nil)
		o.Data2.Constructor = 1487813065
		return o
	},
	1160515173: func() TLObject { // 0x452c0e65
		o := MakeTLMessage(nil)
		o.Data2.Constructor = 1160515173
		return o
	},
	1157215293: func() TLObject { // 0x44f9b43d
		o := MakeTLMessage(nil)
		o.Data2.Constructor = 1157215293
		return o
	},
	678405636: func() TLObject { // 0x286fa604
		o := MakeTLMessageService(nil)
		o.Data2.Constructor = 678405636
		return o
	},
	-1642487306: func() TLObject { // 0x9e19a1f6
		o := MakeTLMessageService(nil)
		o.Data2.Constructor = -1642487306
		return o
	},
	-438840932: func() TLObject { // 0xe5d7d19c
		o := MakeTLMessagesChatFull(nil)
		o.Data2.Constructor = -438840932
		return o
	},
	-2066640507: func() TLObject { // 0x84d19185
		o := MakeTLMessagesAffectedMessages(nil)
		o.Data2.Constructor = -2066640507
		return o
	},
	2131196633: func() TLObject { // 0x7f077ad9
		o := MakeTLContactsResolvedPeer(nil)
		o.Data2.Constructor = 2131196633
		return o
	},
	-1798033689: func() TLObject { // 0x94d42ee7
		o := MakeTLChannelMessagesFilterEmpty(nil)
		o.Data2.Constructor = -1798033689
		return o
	},
	-847783593: func() TLObject { // 0xcd77d957
		o := MakeTLChannelMessagesFilter(nil)
		o.Data2.Constructor = -847783593
		return o
	},
	1611985938: func() TLObject { // 0x6014f412
		o := MakeTLStatsGroupTopAdmin(nil)
		o.Data2.Constructor = 1611985938
		return o
	},
	-1378534221: func() TLObject { // 0xadd53cb3
		o := MakeTLPeerNotifyEventsEmpty(nil)
		o.Data2.Constructor = -1378534221
		return o
	},
	1830677896: func() TLObject { // 0x6d1ded88
		o := MakeTLPeerNotifyEventsAll(nil)
		o.Data2.Constructor = 1830677896
		return o
	},
	-1649296275: func() TLObject { // 0x9db1bc6d
		o := MakeTLPeerUser(nil)
		o.Data2.Constructor = -1649296275
		return o
	},
	-1160714821: func() TLObject { // 0xbad0e5bb
		o := MakeTLPeerChat(nil)
		o.Data2.Constructor = -1160714821
		return o
	},
	-1109531342: func() TLObject { // 0xbddde532
		o := MakeTLPeerChannel(nil)
		o.Data2.Constructor = -1109531342
		return o
	},
	-276825834: func() TLObject { // 0xef7ff916
		o := MakeTLStatsMegagroupStats(nil)
		o.Data2.Constructor = -276825834
		return o
	},
	506920429: func() TLObject { // 0x1e36fded
		o := MakeTLInputPhoneCall(nil)
		o.Data2.Constructor = 506920429
		return o
	},
	1577067778: func() TLObject { // 0x5e002502
		o := MakeTLAuthSentCode(nil)
		o.Data2.Constructor = 1577067778
		return o
	},
	955951967: func() TLObject { // 0x38faab5f
		o := MakeTLAuthSentCode(nil)
		o.Data2.Constructor = 955951967
		return o
	},
	-302941166: func() TLObject { // 0xedf17c12
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = -302941166
		return o
	},
	1951750604: func() TLObject { // 0x745559cc
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = 1951750604
		return o
	},
	-1901811583: func() TLObject { // 0x8ea4a881
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = -1901811583
		return o
	},
	253890367: func() TLObject { // 0xf220f3f
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = 253890367
		return o
	},
	-288727837: func() TLObject { // 0xeeca5ce3
		o := MakeTLLangPackLanguage(nil)
		o.Data2.Constructor = -288727837
		return o
	},
	106019213: func() TLObject { // 0x651b98d
		o := MakeTLLangPackLanguage(nil)
		o.Data2.Constructor = 106019213
		return o
	},
	292985073: func() TLObject { // 0x117698f1
		o := MakeTLLangPackLanguage(nil)
		o.Data2.Constructor = 292985073
		return o
	},
	1064139624: func() TLObject { // 0x3f6d7b68
		o := MakeTLJsonNull(nil)
		o.Data2.Constructor = 1064139624
		return o
	},
	-952869270: func() TLObject { // 0xc7345e6a
		o := MakeTLJsonBool(nil)
		o.Data2.Constructor = -952869270
		return o
	},
	736157604: func() TLObject { // 0x2be0dfa4
		o := MakeTLJsonNumber(nil)
		o.Data2.Constructor = 736157604
		return o
	},
	-1222740358: func() TLObject { // 0xb71e767a
		o := MakeTLJsonString(nil)
		o.Data2.Constructor = -1222740358
		return o
	},
	-146520221: func() TLObject { // 0xf7444763
		o := MakeTLJsonArray(nil)
		o.Data2.Constructor = -146520221
		return o
	},
	-1715350371: func() TLObject { // 0x99c1d49d
		o := MakeTLJsonObject(nil)
		o.Data2.Constructor = -1715350371
		return o
	},
	1202287072: func() TLObject { // 0x47a971e0
		o := MakeTLStatsURL(nil)
		o.Data2.Constructor = 1202287072
		return o
	},
	-709641735: func() TLObject { // 0xd5b3b9f9
		o := MakeTLEmojiKeyword(nil)
		o.Data2.Constructor = -709641735
		return o
	},
	594408994: func() TLObject { // 0x236df622
		o := MakeTLEmojiKeywordDeleted(nil)
		o.Data2.Constructor = 594408994
		return o
	},
	-875679776: func() TLObject { // 0xcbce2fe0
		o := MakeTLStatsPercentValue(nil)
		o.Data2.Constructor = -875679776
		return o
	},
	590459437: func() TLObject { // 0x2331b22d
		o := MakeTLPhotoEmpty(nil)
		o.Data2.Constructor = 590459437
		return o
	},
	-82216347: func() TLObject { // 0xfb197a65
		o := MakeTLPhoto(nil)
		o.Data2.Constructor = -82216347
		return o
	},
	-797637467: func() TLObject { // 0xd07504a5
		o := MakeTLPhoto(nil)
		o.Data2.Constructor = -797637467
		return o
	},
	-1673036328: func() TLObject { // 0x9c477dd8
		o := MakeTLPhoto(nil)
		o.Data2.Constructor = -1673036328
		return o
	},
	-1836524247: func() TLObject { // 0x9288dd29
		o := MakeTLPhoto(nil)
		o.Data2.Constructor = -1836524247
		return o
	},
	-1504076211: func() TLObject { // 0xa6599e4d
		o := MakeTLMicroBlogs(nil)
		o.Data2.Constructor = -1504076211
		return o
	},
	695905689: func() TLObject { // 0x297aad99
		o := MakeTLMicroBlogs(nil)
		o.Data2.Constructor = 695905689
		return o
	},
	-120717989: func() TLObject { // 0xf8cdfd5b
		o := MakeTLMicroBlogsSlice(nil)
		o.Data2.Constructor = -120717989
		return o
	},
	451418809: func() TLObject { // 0x1ae81ab9
		o := MakeTLMicroBlogsNotModified(nil)
		o.Data2.Constructor = 451418809
		return o
	},
	1335282456: func() TLObject { // 0x4f96cb18
		o := MakeTLInputPrivacyKeyStatusTimestamp(nil)
		o.Data2.Constructor = 1335282456
		return o
	},
	-1107622874: func() TLObject { // 0xbdfb0426
		o := MakeTLInputPrivacyKeyChatInvite(nil)
		o.Data2.Constructor = -1107622874
		return o
	},
	-88417185: func() TLObject { // 0xfabadc5f
		o := MakeTLInputPrivacyKeyPhoneCall(nil)
		o.Data2.Constructor = -88417185
		return o
	},
	-610373422: func() TLObject { // 0xdb9e70d2
		o := MakeTLInputPrivacyKeyPhoneP2P(nil)
		o.Data2.Constructor = -610373422
		return o
	},
	-1529000952: func() TLObject { // 0xa4dd4c08
		o := MakeTLInputPrivacyKeyForwards(nil)
		o.Data2.Constructor = -1529000952
		return o
	},
	1461304012: func() TLObject { // 0x5719bacc
		o := MakeTLInputPrivacyKeyProfilePhoto(nil)
		o.Data2.Constructor = 1461304012
		return o
	},
	55761658: func() TLObject { // 0x352dafa
		o := MakeTLInputPrivacyKeyPhoneNumber(nil)
		o.Data2.Constructor = 55761658
		return o
	},
	-786326563: func() TLObject { // 0xd1219bdd
		o := MakeTLInputPrivacyKeyAddedByPhone(nil)
		o.Data2.Constructor = -786326563
		return o
	},
	-359970525: func() TLObject { // 0xea8b4923
		o := MakeTLInputPrivacyKeyAddedByUsername(nil)
		o.Data2.Constructor = -359970525
		return o
	},
	230398871: func() TLObject { // 0xdbb9b97
		o := MakeTLInputPrivacyKeySendMessage(nil)
		o.Data2.Constructor = 230398871
		return o
	},
	-1707344487: func() TLObject { // 0x9a3bfd99
		o := MakeTLMessagesHighScores(nil)
		o.Data2.Constructor = -1707344487
		return o
	},
	-1387279939: func() TLObject { // 0xad4fc9bd
		o := MakeTLMessageInteractionCounters(nil)
		o.Data2.Constructor = -1387279939
		return o
	},
	-116880147: func() TLObject { // 0xf9088ced
		o := MakeTLPredefinedUser(nil)
		o.Data2.Constructor = -116880147
		return o
	},
	-181407105: func() TLObject { // 0xf52ff27f
		o := MakeTLInputFile(nil)
		o.Data2.Constructor = -181407105
		return o
	},
	-95482955: func() TLObject { // 0xfa4f0bb5
		o := MakeTLInputFileBig(nil)
		o.Data2.Constructor = -95482955
		return o
	},
	-791039645: func() TLObject { // 0xd0d9b163
		o := MakeTLChannelsChannelParticipant(nil)
		o.Data2.Constructor = -791039645
		return o
	},
	1722786150: func() TLObject { // 0x66afa166
		o := MakeTLHelpDeepLinkInfoEmpty(nil)
		o.Data2.Constructor = 1722786150
		return o
	},
	1783556146: func() TLObject { // 0x6a4ee832
		o := MakeTLHelpDeepLinkInfo(nil)
		o.Data2.Constructor = 1783556146
		return o
	},
	1577484359: func() TLObject { // 0x5e068047
		o := MakeTLPageListOrderedItemText(nil)
		o.Data2.Constructor = 1577484359
		return o
	},
	-1730311882: func() TLObject { // 0x98dd8936
		o := MakeTLPageListOrderedItemBlocks(nil)
		o.Data2.Constructor = -1730311882
		return o
	},
	-1945767479: func() TLObject { // 0x8c05f1c9
		o := MakeTLHelpSupportName(nil)
		o.Data2.Constructor = -1945767479
		return o
	},
	497489295: func() TLObject { // 0x1da7158f
		o := MakeTLHelpAppUpdate(nil)
		o.Data2.Constructor = 497489295
		return o
	},
	-1987579119: func() TLObject { // 0x8987f311
		o := MakeTLHelpAppUpdate(nil)
		o.Data2.Constructor = -1987579119
		return o
	},
	-1000708810: func() TLObject { // 0xc45a6536
		o := MakeTLHelpNoAppUpdate(nil)
		o.Data2.Constructor = -1000708810
		return o
	},
	1821452115: func() TLObject { // 0x6c912753
		o := MakeTLInputBlogPhotos(nil)
		o.Data2.Constructor = 1821452115
		return o
	},
	162942657: func() TLObject { // 0x9b64ec1
		o := MakeTLInputBlogUploadVideo(nil)
		o.Data2.Constructor = 162942657
		return o
	},
	479366645: func() TLObject { // 0x1c928df5
		o := MakeTLInputBlogVideo(nil)
		o.Data2.Constructor = 479366645
		return o
	},
	-713800393: func() TLObject { // 0xd5744537
		o := MakeTLBlogsUnreadEmpty(nil)
		o.Data2.Constructor = -713800393
		return o
	},
	-899622183: func() TLObject { // 0xca60dad9
		o := MakeTLBlogsUnreadTooLong(nil)
		o.Data2.Constructor = -899622183
		return o
	},
	-1865474047: func() TLObject { // 0x90cf2001
		o := MakeTLBlogsUnreads(nil)
		o.Data2.Constructor = -1865474047
		return o
	},
	-1159937629: func() TLObject { // 0xbadcc1a3
		o := MakeTLPollResults(nil)
		o.Data2.Constructor = -1159937629
		return o
	},
	-932174686: func() TLObject { // 0xc87024a2
		o := MakeTLPollResults(nil)
		o.Data2.Constructor = -932174686
		return o
	},
	1465219162: func() TLObject { // 0x5755785a
		o := MakeTLPollResults(nil)
		o.Data2.Constructor = 1465219162
		return o
	},
	483901197: func() TLObject { // 0x1cd7bf0d
		o := MakeTLInputPhotoEmpty(nil)
		o.Data2.Constructor = 483901197
		return o
	},
	1001634122: func() TLObject { // 0x3bb3b94a
		o := MakeTLInputPhoto(nil)
		o.Data2.Constructor = 1001634122
		return o
	},
	-74070332: func() TLObject { // 0xfb95c6c4
		o := MakeTLInputPhoto(nil)
		o.Data2.Constructor = -74070332
		return o
	},
	1599050311: func() TLObject { // 0x5f4f9247
		o := MakeTLContactLinkUnknown(nil)
		o.Data2.Constructor = 1599050311
		return o
	},
	-17968211: func() TLObject { // 0xfeedd3ad
		o := MakeTLContactLinkNone(nil)
		o.Data2.Constructor = -17968211
		return o
	},
	-721239344: func() TLObject { // 0xd502c2d0
		o := MakeTLContactLinkContact(nil)
		o.Data2.Constructor = -721239344
		return o
	},
	646922073: func() TLObject { // 0x268f3f59
		o := MakeTLContactLinkHasPhone(nil)
		o.Data2.Constructor = 646922073
		return o
	},
	1571494644: func() TLObject { // 0x5dab1af4
		o := MakeTLExportedMessageLink(nil)
		o.Data2.Constructor = 1571494644
		return o
	},
	-317144808: func() TLObject { // 0xed18c118
		o := MakeTLEncryptedMessage(nil)
		o.Data2.Constructor = -317144808
		return o
	},
	594758406: func() TLObject { // 0x23734b06
		o := MakeTLEncryptedMessageService(nil)
		o.Data2.Constructor = 594758406
		return o
	},
	-567906571: func() TLObject { // 0xde266ef5
		o := MakeTLContactsTopPeersNotModified(nil)
		o.Data2.Constructor = -567906571
		return o
	},
	1891070632: func() TLObject { // 0x70b772a8
		o := MakeTLContactsTopPeers(nil)
		o.Data2.Constructor = 1891070632
		return o
	},
	-1255369827: func() TLObject { // 0xb52c939d
		o := MakeTLContactsTopPeersDisabled(nil)
		o.Data2.Constructor = -1255369827
		return o
	},
	-1678949555: func() TLObject { // 0x9bed434d
		o := MakeTLInputWebDocument(nil)
		o.Data2.Constructor = -1678949555
		return o
	},
	-526508104: func() TLObject { // 0xe09e1fb8
		o := MakeTLHelpProxyDataEmpty(nil)
		o.Data2.Constructor = -526508104
		return o
	},
	737668643: func() TLObject { // 0x2bf7ee23
		o := MakeTLHelpProxyDataPromo(nil)
		o.Data2.Constructor = 737668643
		return o
	},
	-568988681: func() TLObject { // 0xde15ebf7
		o := MakeTLBlogsTopic(nil)
		o.Data2.Constructor = -568988681
		return o
	},
	-1658158621: func() TLObject { // 0x9d2a81e3
		o := MakeTLSecureValueTypePersonalDetails(nil)
		o.Data2.Constructor = -1658158621
		return o
	},
	1034709504: func() TLObject { // 0x3dac6a00
		o := MakeTLSecureValueTypePassport(nil)
		o.Data2.Constructor = 1034709504
		return o
	},
	115615172: func() TLObject { // 0x6e425c4
		o := MakeTLSecureValueTypeDriverLicense(nil)
		o.Data2.Constructor = 115615172
		return o
	},
	-1596951477: func() TLObject { // 0xa0d0744b
		o := MakeTLSecureValueTypeIdentityCard(nil)
		o.Data2.Constructor = -1596951477
		return o
	},
	-1717268701: func() TLObject { // 0x99a48f23
		o := MakeTLSecureValueTypeInternalPassport(nil)
		o.Data2.Constructor = -1717268701
		return o
	},
	-874308058: func() TLObject { // 0xcbe31e26
		o := MakeTLSecureValueTypeAddress(nil)
		o.Data2.Constructor = -874308058
		return o
	},
	-63531698: func() TLObject { // 0xfc36954e
		o := MakeTLSecureValueTypeUtilityBill(nil)
		o.Data2.Constructor = -63531698
		return o
	},
	-1995211763: func() TLObject { // 0x89137c0d
		o := MakeTLSecureValueTypeBankStatement(nil)
		o.Data2.Constructor = -1995211763
		return o
	},
	-1954007928: func() TLObject { // 0x8b883488
		o := MakeTLSecureValueTypeRentalAgreement(nil)
		o.Data2.Constructor = -1954007928
		return o
	},
	-1713143702: func() TLObject { // 0x99e3806a
		o := MakeTLSecureValueTypePassportRegistration(nil)
		o.Data2.Constructor = -1713143702
		return o
	},
	-368907213: func() TLObject { // 0xea02ec33
		o := MakeTLSecureValueTypeTemporaryRegistration(nil)
		o.Data2.Constructor = -368907213
		return o
	},
	-1289704741: func() TLObject { // 0xb320aadb
		o := MakeTLSecureValueTypePhone(nil)
		o.Data2.Constructor = -1289704741
		return o
	},
	-1908627474: func() TLObject { // 0x8e3ca7ee
		o := MakeTLSecureValueTypeEmail(nil)
		o.Data2.Constructor = -1908627474
		return o
	},
	1338747336: func() TLObject { // 0x4fcba9c8
		o := MakeTLMessagesArchivedStickers(nil)
		o.Data2.Constructor = 1338747336
		return o
	},
	871426631: func() TLObject { // 0x33f0ea47
		o := MakeTLSecureCredentialsEncrypted(nil)
		o.Data2.Constructor = 871426631
		return o
	},
	-398136321: func() TLObject { // 0xe844ebff
		o := MakeTLMessagesSearchCounter(nil)
		o.Data2.Constructor = -398136321
		return o
	},
	1443858741: func() TLObject { // 0x560f8935
		o := MakeTLMessagesSentEncryptedMessage(nil)
		o.Data2.Constructor = 1443858741
		return o
	},
	-1802240206: func() TLObject { // 0x9493ff32
		o := MakeTLMessagesSentEncryptedFile(nil)
		o.Data2.Constructor = -1802240206
		return o
	},
	1038967584: func() TLObject { // 0x3ded6320
		o := MakeTLMessageMediaEmpty(nil)
		o.Data2.Constructor = 1038967584
		return o
	},
	1766936791: func() TLObject { // 0x695150d7
		o := MakeTLMessageMediaPhoto(nil)
		o.Data2.Constructor = 1766936791
		return o
	},
	-1256047857: func() TLObject { // 0xb5223b0f
		o := MakeTLMessageMediaPhoto(nil)
		o.Data2.Constructor = -1256047857
		return o
	},
	1457575028: func() TLObject { // 0x56e0d474
		o := MakeTLMessageMediaGeo(nil)
		o.Data2.Constructor = 1457575028
		return o
	},
	-873313984: func() TLObject { // 0xcbf24940
		o := MakeTLMessageMediaContact(nil)
		o.Data2.Constructor = -873313984
		return o
	},
	1585262393: func() TLObject { // 0x5e7d2f39
		o := MakeTLMessageMediaContact(nil)
		o.Data2.Constructor = 1585262393
		return o
	},
	-1618676578: func() TLObject { // 0x9f84f49e
		o := MakeTLMessageMediaUnsupported(nil)
		o.Data2.Constructor = -1618676578
		return o
	},
	-1666158377: func() TLObject { // 0x9cb070d7
		o := MakeTLMessageMediaDocument(nil)
		o.Data2.Constructor = -1666158377
		return o
	},
	2084836563: func() TLObject { // 0x7c4414d3
		o := MakeTLMessageMediaDocument(nil)
		o.Data2.Constructor = 2084836563
		return o
	},
	-1557277184: func() TLObject { // 0xa32dd600
		o := MakeTLMessageMediaWebPage(nil)
		o.Data2.Constructor = -1557277184
		return o
	},
	784356159: func() TLObject { // 0x2ec0533f
		o := MakeTLMessageMediaVenue(nil)
		o.Data2.Constructor = 784356159
		return o
	},
	-38694904: func() TLObject { // 0xfdb19008
		o := MakeTLMessageMediaGame(nil)
		o.Data2.Constructor = -38694904
		return o
	},
	-2074799289: func() TLObject { // 0x84551347
		o := MakeTLMessageMediaInvoice(nil)
		o.Data2.Constructor = -2074799289
		return o
	},
	-1186937242: func() TLObject { // 0xb940c666
		o := MakeTLMessageMediaGeoLive(nil)
		o.Data2.Constructor = -1186937242
		return o
	},
	2084316681: func() TLObject { // 0x7c3c2609
		o := MakeTLMessageMediaGeoLive(nil)
		o.Data2.Constructor = 2084316681
		return o
	},
	1272375192: func() TLObject { // 0x4bd6e798
		o := MakeTLMessageMediaPoll(nil)
		o.Data2.Constructor = 1272375192
		return o
	},
	1065280907: func() TLObject { // 0x3f7ee58b
		o := MakeTLMessageMediaDice(nil)
		o.Data2.Constructor = 1065280907
		return o
	},
	1670374507: func() TLObject { // 0x638fe46b
		o := MakeTLMessageMediaDice(nil)
		o.Data2.Constructor = 1670374507
		return o
	},
	2124445994: func() TLObject { // 0x7ea0792a
		o := MakeTLMessageMediaBizDataRaw(nil)
		o.Data2.Constructor = 2124445994
		return o
	},
	414687501: func() TLObject { // 0x18b7a10d
		o := MakeTLDcOption(nil)
		o.Data2.Constructor = 414687501
		return o
	},
	98092748: func() TLObject { // 0x5d8c6cc
		o := MakeTLDcOption(nil)
		o.Data2.Constructor = 98092748
		return o
	},
	-368917890: func() TLObject { // 0xea02c27e
		o := MakeTLPaymentCharge(nil)
		o.Data2.Constructor = -368917890
		return o
	},
	-1239335713: func() TLObject { // 0xb6213cdf
		o := MakeTLShippingOption(nil)
		o.Data2.Constructor = -1239335713
		return o
	},
	-732254058: func() TLObject { // 0xd45ab096
		o := MakeTLPasswordKdfAlgoUnknown(nil)
		o.Data2.Constructor = -732254058
		return o
	},
	982592842: func() TLObject { // 0x3a912d4a
		o := MakeTLPasswordKdfAlgoModPow(nil)
		o.Data2.Constructor = 982592842
		return o
	},
	-1237164374: func() TLObject { // 0xb6425eaa
		o := MakeTLPasswordKdfAlgo10000(nil)
		o.Data2.Constructor = -1237164374
		return o
	},
	-1132882121: func() TLObject { // 0xbc799737
		o := MakeTLBoolFalse(nil)
		o.Data2.Constructor = -1132882121
		return o
	},
	-1720552011: func() TLObject { // 0x997275b5
		o := MakeTLBoolTrue(nil)
		o.Data2.Constructor = -1720552011
		return o
	},
	2010127419: func() TLObject { // 0x77d01c3b
		o := MakeTLContactsImportedContacts(nil)
		o.Data2.Constructor = 2010127419
		return o
	},
	-1395872411: func() TLObject { // 0xacccad65
		o := MakeTLBlogsIdTypeBlog(nil)
		o.Data2.Constructor = -1395872411
		return o
	},
	1539461112: func() TLObject { // 0x5bc24ff8
		o := MakeTLBlogsIdTypeComment(nil)
		o.Data2.Constructor = 1539461112
		return o
	},
	218751099: func() TLObject { // 0xd09e07b
		o := MakeTLInputPrivacyValueAllowContacts(nil)
		o.Data2.Constructor = 218751099
		return o
	},
	407582158: func() TLObject { // 0x184b35ce
		o := MakeTLInputPrivacyValueAllowAll(nil)
		o.Data2.Constructor = 407582158
		return o
	},
	320652927: func() TLObject { // 0x131cc67f
		o := MakeTLInputPrivacyValueAllowUsers(nil)
		o.Data2.Constructor = 320652927
		return o
	},
	195371015: func() TLObject { // 0xba52007
		o := MakeTLInputPrivacyValueDisallowContacts(nil)
		o.Data2.Constructor = 195371015
		return o
	},
	-697604407: func() TLObject { // 0xd66b66c9
		o := MakeTLInputPrivacyValueDisallowAll(nil)
		o.Data2.Constructor = -697604407
		return o
	},
	-1877932953: func() TLObject { // 0x90110467
		o := MakeTLInputPrivacyValueDisallowUsers(nil)
		o.Data2.Constructor = -1877932953
		return o
	},
	1283572154: func() TLObject { // 0x4c81c1ba
		o := MakeTLInputPrivacyValueAllowChatParticipants(nil)
		o.Data2.Constructor = 1283572154
		return o
	},
	-668769361: func() TLObject { // 0xd82363af
		o := MakeTLInputPrivacyValueDisallowChatParticipants(nil)
		o.Data2.Constructor = -668769361
		return o
	},
	182649427: func() TLObject { // 0xae30253
		o := MakeTLMessageRange(nil)
		o.Data2.Constructor = 182649427
		return o
	},
	-892779534: func() TLObject { // 0xcac943f2
		o := MakeTLWebAuthorization(nil)
		o.Data2.Constructor = -892779534
		return o
	},
	471437699: func() TLObject { // 0x1c199183
		o := MakeTLAccountWallPapersNotModified(nil)
		o.Data2.Constructor = 471437699
		return o
	},
	1881892265: func() TLObject { // 0x702b65a9
		o := MakeTLAccountWallPapers(nil)
		o.Data2.Constructor = 1881892265
		return o
	},
	-557924733: func() TLObject { // 0xdebebe83
		o := MakeTLCodeSettings(nil)
		o.Data2.Constructor = -557924733
		return o
	},
	808409587: func() TLObject { // 0x302f59f3
		o := MakeTLCodeSettings(nil)
		o.Data2.Constructor = 808409587
		return o
	},
	1251247737: func() TLObject { // 0x4a948679
		o := MakeTLChatFull(nil)
		o.Data2.Constructor = 1251247737
		return o
	},
	461151667: func() TLObject { // 0x1b7c9db3
		o := MakeTLChatFull(nil)
		o.Data2.Constructor = 461151667
		return o
	},
	581055962: func() TLObject { // 0x22a235da
		o := MakeTLChatFull(nil)
		o.Data2.Constructor = 581055962
		return o
	},
	-304961647: func() TLObject { // 0xedd2a791
		o := MakeTLChatFull(nil)
		o.Data2.Constructor = -304961647
		return o
	},
	771925524: func() TLObject { // 0x2e02a614
		o := MakeTLChatFull(nil)
		o.Data2.Constructor = 771925524
		return o
	},
	-1569432445: func() TLObject { // 0xa2745c83
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = -1569432445
		return o
	},
	-253335766: func() TLObject { // 0xf0e6672a
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = -253335766
		return o
	},
	763976820: func() TLObject { // 0x2d895c74
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = 763976820
		return o
	},
	277964371: func() TLObject { // 0x10916653
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = 277964371
		return o
	},
	-1736252138: func() TLObject { // 0x9882e516
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = -1736252138
		return o
	},
	56920439: func() TLObject { // 0x3648977
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = 56920439
		return o
	},
	478652186: func() TLObject { // 0x1c87a71a
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = 478652186
		return o
	},
	1991201921: func() TLObject { // 0x76af5481
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = 1991201921
		return o
	},
	641506392: func() TLObject { // 0x263c9c58
		o := MakeTLSchemeNotModified(nil)
		o.Data2.Constructor = 641506392
		return o
	},
	1315894878: func() TLObject { // 0x4e6ef65e
		o := MakeTLScheme(nil)
		o.Data2.Constructor = 1315894878
		return o
	},
	-1986399595: func() TLObject { // 0x8999f295
		o := MakeTLStatsMessageStats(nil)
		o.Data2.Constructor = -1986399595
		return o
	},
	381645902: func() TLObject { // 0x16bf744e
		o := MakeTLSendMessageTypingAction(nil)
		o.Data2.Constructor = 381645902
		return o
	},
	-44119819: func() TLObject { // 0xfd5ec8f5
		o := MakeTLSendMessageCancelAction(nil)
		o.Data2.Constructor = -44119819
		return o
	},
	-1584933265: func() TLObject { // 0xa187d66f
		o := MakeTLSendMessageRecordVideoAction(nil)
		o.Data2.Constructor = -1584933265
		return o
	},
	-378127636: func() TLObject { // 0xe9763aec
		o := MakeTLSendMessageUploadVideoAction(nil)
		o.Data2.Constructor = -378127636
		return o
	},
	-718310409: func() TLObject { // 0xd52f73f7
		o := MakeTLSendMessageRecordAudioAction(nil)
		o.Data2.Constructor = -718310409
		return o
	},
	-212740181: func() TLObject { // 0xf351d7ab
		o := MakeTLSendMessageUploadAudioAction(nil)
		o.Data2.Constructor = -212740181
		return o
	},
	-774682074: func() TLObject { // 0xd1d34a26
		o := MakeTLSendMessageUploadPhotoAction(nil)
		o.Data2.Constructor = -774682074
		return o
	},
	-1441998364: func() TLObject { // 0xaa0cd9e4
		o := MakeTLSendMessageUploadDocumentAction(nil)
		o.Data2.Constructor = -1441998364
		return o
	},
	393186209: func() TLObject { // 0x176f8ba1
		o := MakeTLSendMessageGeoLocationAction(nil)
		o.Data2.Constructor = 393186209
		return o
	},
	1653390447: func() TLObject { // 0x628cbc6f
		o := MakeTLSendMessageChooseContactAction(nil)
		o.Data2.Constructor = 1653390447
		return o
	},
	-580219064: func() TLObject { // 0xdd6a8f48
		o := MakeTLSendMessageGamePlayAction(nil)
		o.Data2.Constructor = -580219064
		return o
	},
	-1997373508: func() TLObject { // 0x88f27fbc
		o := MakeTLSendMessageRecordRoundAction(nil)
		o.Data2.Constructor = -1997373508
		return o
	},
	608050278: func() TLObject { // 0x243e1c66
		o := MakeTLSendMessageUploadRoundAction(nil)
		o.Data2.Constructor = 608050278
		return o
	},
	911761060: func() TLObject { // 0x36585ea4
		o := MakeTLMessagesBotCallbackAnswer(nil)
		o.Data2.Constructor = 911761060
		return o
	},
	53231223: func() TLObject { // 0x32c3e77
		o := MakeTLInputGameID(nil)
		o.Data2.Constructor = 53231223
		return o
	},
	-1020139510: func() TLObject { // 0xc331e80a
		o := MakeTLInputGameShortName(nil)
		o.Data2.Constructor = -1020139510
		return o
	},
	878078826: func() TLObject { // 0x34566b6a
		o := MakeTLPageTableCell(nil)
		o.Data2.Constructor = 878078826
		return o
	},
	1348066419: func() TLObject { // 0x5059dc73
		o := MakeTLFeedPosition(nil)
		o.Data2.Constructor = 1348066419
		return o
	},
	1182322895: func() TLObject { // 0x4678d0cf
		o := MakeTLMessagesFeedMessagesNotModified(nil)
		o.Data2.Constructor = 1182322895
		return o
	},
	1438884273: func() TLObject { // 0x55c3a1b1
		o := MakeTLMessagesFeedMessages(nil)
		o.Data2.Constructor = 1438884273
		return o
	},
	164646985: func() TLObject { // 0x9d05049
		o := MakeTLUserStatusEmpty(nil)
		o.Data2.Constructor = 164646985
		return o
	},
	-306628279: func() TLObject { // 0xedb93949
		o := MakeTLUserStatusOnline(nil)
		o.Data2.Constructor = -306628279
		return o
	},
	9203775: func() TLObject { // 0x8c703f
		o := MakeTLUserStatusOffline(nil)
		o.Data2.Constructor = 9203775
		return o
	},
	-496024847: func() TLObject { // 0xe26f42f1
		o := MakeTLUserStatusRecently(nil)
		o.Data2.Constructor = -496024847
		return o
	},
	129960444: func() TLObject { // 0x7bf09fc
		o := MakeTLUserStatusLastWeek(nil)
		o.Data2.Constructor = 129960444
		return o
	},
	2011940674: func() TLObject { // 0x77ebc742
		o := MakeTLUserStatusLastMonth(nil)
		o.Data2.Constructor = 2011940674
		return o
	},
	-1390001672: func() TLObject { // 0xad2641f8
		o := MakeTLAccountPassword(nil)
		o.Data2.Constructor = -1390001672
		return o
	},
	1753693093: func() TLObject { // 0x68873ba5
		o := MakeTLAccountPassword(nil)
		o.Data2.Constructor = 1753693093
		return o
	},
	-902187961: func() TLObject { // 0xca39b447
		o := MakeTLAccountPassword(nil)
		o.Data2.Constructor = -902187961
		return o
	},
	2081952796: func() TLObject { // 0x7c18141c
		o := MakeTLAccountPassword(nil)
		o.Data2.Constructor = 2081952796
		return o
	},
	1587643126: func() TLObject { // 0x5ea182f6
		o := MakeTLAccountNoPassword(nil)
		o.Data2.Constructor = 1587643126
		return o
	},
	-1764049896: func() TLObject { // 0x96dabc18
		o := MakeTLAccountNoPassword(nil)
		o.Data2.Constructor = -1764049896
		return o
	},
	-1606526075: func() TLObject { // 0xa03e5b85
		o := MakeTLReplyKeyboardHide(nil)
		o.Data2.Constructor = -1606526075
		return o
	},
	-200242528: func() TLObject { // 0xf4108aa0
		o := MakeTLReplyKeyboardForceReply(nil)
		o.Data2.Constructor = -200242528
		return o
	},
	889353612: func() TLObject { // 0x3502758c
		o := MakeTLReplyKeyboardMarkup(nil)
		o.Data2.Constructor = 889353612
		return o
	},
	1218642516: func() TLObject { // 0x48a30254
		o := MakeTLReplyInlineMarkup(nil)
		o.Data2.Constructor = 1218642516
		return o
	},
	-292807034: func() TLObject { // 0xee8c1e86
		o := MakeTLInputChannelEmpty(nil)
		o.Data2.Constructor = -292807034
		return o
	},
	-1343524562: func() TLObject { // 0xafeb712e
		o := MakeTLInputChannel(nil)
		o.Data2.Constructor = -1343524562
		return o
	},
	707290417: func() TLObject { // 0x2a286531
		o := MakeTLInputChannelFromMessage(nil)
		o.Data2.Constructor = 707290417
		return o
	},
	-177282392: func() TLObject { // 0xf56ee2a8
		o := MakeTLChannelsChannelParticipants(nil)
		o.Data2.Constructor = -177282392
		return o
	},
	-266911767: func() TLObject { // 0xf0173fe9
		o := MakeTLChannelsChannelParticipantsNotModified(nil)
		o.Data2.Constructor = -266911767
		return o
	},
	-2048646399: func() TLObject { // 0x85e42301
		o := MakeTLPhoneCallDiscardReasonMissed(nil)
		o.Data2.Constructor = -2048646399
		return o
	},
	-527056480: func() TLObject { // 0xe095c1a0
		o := MakeTLPhoneCallDiscardReasonDisconnect(nil)
		o.Data2.Constructor = -527056480
		return o
	},
	1471006352: func() TLObject { // 0x57adc690
		o := MakeTLPhoneCallDiscardReasonHangup(nil)
		o.Data2.Constructor = 1471006352
		return o
	},
	-84416311: func() TLObject { // 0xfaf7e8c9
		o := MakeTLPhoneCallDiscardReasonBusy(nil)
		o.Data2.Constructor = -84416311
		return o
	},
	-614138572: func() TLObject { // 0xdb64fd34
		o := MakeTLAccountTmpPassword(nil)
		o.Data2.Constructor = -614138572
		return o
	},
	1462101002: func() TLObject { // 0x5725e40a
		o := MakeTLCdnConfig(nil)
		o.Data2.Constructor = 1462101002
		return o
	},
	1417103303: func() TLObject { // 0x547747c7
		o := MakeTLWalletRecord(nil)
		o.Data2.Constructor = 1417103303
		return o
	},
	-399391402: func() TLObject { // 0xe831c556
		o := MakeTLVideoSize(nil)
		o.Data2.Constructor = -399391402
		return o
	},
	1130084743: func() TLObject { // 0x435bb987
		o := MakeTLVideoSize(nil)
		o.Data2.Constructor = 1130084743
		return o
	},
	-1964327229: func() TLObject { // 0x8aeabec3
		o := MakeTLSecureData(nil)
		o.Data2.Constructor = -1964327229
		return o
	},
	1679398724: func() TLObject { // 0x64199744
		o := MakeTLSecureFileEmpty(nil)
		o.Data2.Constructor = 1679398724
		return o
	},
	-534283678: func() TLObject { // 0xe0277a62
		o := MakeTLSecureFile(nil)
		o.Data2.Constructor = -534283678
		return o
	},
	-1831650802: func() TLObject { // 0x92d33a0e
		o := MakeTLUrlAuthResultRequest(nil)
		o.Data2.Constructor = -1831650802
		return o
	},
	-1886646706: func() TLObject { // 0x8f8c0e4e
		o := MakeTLUrlAuthResultAccepted(nil)
		o.Data2.Constructor = -1886646706
		return o
	},
	-1445536993: func() TLObject { // 0xa9d6db1f
		o := MakeTLUrlAuthResultDefault(nil)
		o.Data2.Constructor = -1445536993
		return o
	},
	-1014526429: func() TLObject { // 0xc3878e23
		o := MakeTLHelpCountry(nil)
		o.Data2.Constructor = -1014526429
		return o
	},
	1093204652: func() TLObject { // 0x4128faac
		o := MakeTLMessageReplies(nil)
		o.Data2.Constructor = 1093204652
		return o
	},
	859091184: func() TLObject { // 0x3334b0f0
		o := MakeTLInputSecureFileUploaded(nil)
		o.Data2.Constructor = 859091184
		return o
	},
	1399317950: func() TLObject { // 0x5367e5be
		o := MakeTLInputSecureFile(nil)
		o.Data2.Constructor = 1399317950
		return o
	},
	537022650: func() TLObject { // 0x200250ba
		o := MakeTLUserEmpty(nil)
		o.Data2.Constructor = 537022650
		return o
	},
	-1820043071: func() TLObject { // 0x938458c1
		o := MakeTLUser(nil)
		o.Data2.Constructor = -1820043071
		return o
	},
	773059779: func() TLObject { // 0x2e13f4c3
		o := MakeTLUser(nil)
		o.Data2.Constructor = 773059779
		return o
	},
	-158703678: func() TLObject { // 0xf68a5fc2
		o := MakeTLBlogGeoPoint(nil)
		o.Data2.Constructor = -158703678
		return o
	},
	1567990072: func() TLObject { // 0x5d75a138
		o := MakeTLUpdatesDifferenceEmpty(nil)
		o.Data2.Constructor = 1567990072
		return o
	},
	16030880: func() TLObject { // 0xf49ca0
		o := MakeTLUpdatesDifference(nil)
		o.Data2.Constructor = 16030880
		return o
	},
	-1459938943: func() TLObject { // 0xa8fb1981
		o := MakeTLUpdatesDifferenceSlice(nil)
		o.Data2.Constructor = -1459938943
		return o
	},
	1258196845: func() TLObject { // 0x4afe8f6d
		o := MakeTLUpdatesDifferenceTooLong(nil)
		o.Data2.Constructor = 1258196845
		return o
	},
	-247351839: func() TLObject { // 0xf141b5e1
		o := MakeTLInputEncryptedChat(nil)
		o.Data2.Constructor = -247351839
		return o
	},
	-532532493: func() TLObject { // 0xe04232f3
		o := MakeTLAutoDownloadSettings(nil)
		o.Data2.Constructor = -532532493
		return o
	},
	-767099577: func() TLObject { // 0xd246fd47
		o := MakeTLAutoDownloadSettings(nil)
		o.Data2.Constructor = -767099577
		return o
	},
	-994444869: func() TLObject { // 0xc4b9f9bb
		o := MakeTLError(nil)
		o.Data2.Constructor = -994444869
		return o
	},
	295067450: func() TLObject { // 0x11965f3a
		o := MakeTLBotInlineResult(nil)
		o.Data2.Constructor = 295067450
		return o
	},
	-1679053127: func() TLObject { // 0x9bebaeb9
		o := MakeTLBotInlineResult(nil)
		o.Data2.Constructor = -1679053127
		return o
	},
	400266251: func() TLObject { // 0x17db940b
		o := MakeTLBotInlineMediaResult(nil)
		o.Data2.Constructor = 400266251
		return o
	},
	-1022713000: func() TLObject { // 0xc30aa358
		o := MakeTLInvoice(nil)
		o.Data2.Constructor = -1022713000
		return o
	},
	1304052993: func() TLObject { // 0x4dba4501
		o := MakeTLAccountTakeout(nil)
		o.Data2.Constructor = 1304052993
		return o
	},
	1012306921: func() TLObject { // 0x3c5693e9
		o := MakeTLInputTheme(nil)
		o.Data2.Constructor = 1012306921
		return o
	},
	-175567375: func() TLObject { // 0xf5890df1
		o := MakeTLInputThemeSlug(nil)
		o.Data2.Constructor = -175567375
		return o
	},
	-1107852396: func() TLObject { // 0xbdf78394
		o := MakeTLStatsBroadcastStats(nil)
		o.Data2.Constructor = -1107852396
		return o
	},
	-1495959709: func() TLObject { // 0xa6d57763
		o := MakeTLMessageReplyHeader(nil)
		o.Data2.Constructor = -1495959709
		return o
	},
	1041346555: func() TLObject { // 0x3e11affb
		o := MakeTLUpdatesChannelDifferenceEmpty(nil)
		o.Data2.Constructor = 1041346555
		return o
	},
	-1531132162: func() TLObject { // 0xa4bcc6fe
		o := MakeTLUpdatesChannelDifferenceTooLong(nil)
		o.Data2.Constructor = -1531132162
		return o
	},
	1788705589: func() TLObject { // 0x6a9d7b35
		o := MakeTLUpdatesChannelDifferenceTooLong(nil)
		o.Data2.Constructor = 1788705589
		return o
	},
	543450958: func() TLObject { // 0x2064674e
		o := MakeTLUpdatesChannelDifference(nil)
		o.Data2.Constructor = 543450958
		return o
	},
	-1240849242: func() TLObject { // 0xb60a24a6
		o := MakeTLMessagesStickerSet(nil)
		o.Data2.Constructor = -1240849242
		return o
	},
	1601666510: func() TLObject { // 0x5f777dce
		o := MakeTLMessageFwdHeader(nil)
		o.Data2.Constructor = 1601666510
		return o
	},
	893020267: func() TLObject { // 0x353a686b
		o := MakeTLMessageFwdHeader(nil)
		o.Data2.Constructor = 893020267
		return o
	},
	-332168592: func() TLObject { // 0xec338270
		o := MakeTLMessageFwdHeader(nil)
		o.Data2.Constructor = -332168592
		return o
	},
	1436466797: func() TLObject { // 0x559ebe6d
		o := MakeTLMessageFwdHeader(nil)
		o.Data2.Constructor = 1436466797
		return o
	},
	1008755359: func() TLObject { // 0x3c20629f
		o := MakeTLInlineBotSwitchPM(nil)
		o.Data2.Constructor = 1008755359
		return o
	},
	324435594: func() TLObject { // 0x13567e8a
		o := MakeTLPageBlockUnsupported(nil)
		o.Data2.Constructor = 324435594
		return o
	},
	1890305021: func() TLObject { // 0x70abc3fd
		o := MakeTLPageBlockTitle(nil)
		o.Data2.Constructor = 1890305021
		return o
	},
	-1879401953: func() TLObject { // 0x8ffa9a1f
		o := MakeTLPageBlockSubtitle(nil)
		o.Data2.Constructor = -1879401953
		return o
	},
	-1162877472: func() TLObject { // 0xbaafe5e0
		o := MakeTLPageBlockAuthorDate(nil)
		o.Data2.Constructor = -1162877472
		return o
	},
	-1076861716: func() TLObject { // 0xbfd064ec
		o := MakeTLPageBlockHeader(nil)
		o.Data2.Constructor = -1076861716
		return o
	},
	-248793375: func() TLObject { // 0xf12bb6e1
		o := MakeTLPageBlockSubheader(nil)
		o.Data2.Constructor = -248793375
		return o
	},
	1182402406: func() TLObject { // 0x467a0766
		o := MakeTLPageBlockParagraph(nil)
		o.Data2.Constructor = 1182402406
		return o
	},
	-1066346178: func() TLObject { // 0xc070d93e
		o := MakeTLPageBlockPreformatted(nil)
		o.Data2.Constructor = -1066346178
		return o
	},
	1216809369: func() TLObject { // 0x48870999
		o := MakeTLPageBlockFooter(nil)
		o.Data2.Constructor = 1216809369
		return o
	},
	-618614392: func() TLObject { // 0xdb20b188
		o := MakeTLPageBlockDivider(nil)
		o.Data2.Constructor = -618614392
		return o
	},
	-837994576: func() TLObject { // 0xce0d37b0
		o := MakeTLPageBlockAnchor(nil)
		o.Data2.Constructor = -837994576
		return o
	},
	-454524911: func() TLObject { // 0xe4e88011
		o := MakeTLPageBlockList(nil)
		o.Data2.Constructor = -454524911
		return o
	},
	978896884: func() TLObject { // 0x3a58c7f4
		o := MakeTLPageBlockList(nil)
		o.Data2.Constructor = 978896884
		return o
	},
	641563686: func() TLObject { // 0x263d7c26
		o := MakeTLPageBlockBlockquote(nil)
		o.Data2.Constructor = 641563686
		return o
	},
	1329878739: func() TLObject { // 0x4f4456d3
		o := MakeTLPageBlockPullquote(nil)
		o.Data2.Constructor = 1329878739
		return o
	},
	391759200: func() TLObject { // 0x1759c560
		o := MakeTLPageBlockPhoto(nil)
		o.Data2.Constructor = 391759200
		return o
	},
	-372860542: func() TLObject { // 0xe9c69982
		o := MakeTLPageBlockPhoto(nil)
		o.Data2.Constructor = -372860542
		return o
	},
	2089805750: func() TLObject { // 0x7c8fe7b6
		o := MakeTLPageBlockVideo(nil)
		o.Data2.Constructor = 2089805750
		return o
	},
	-640214938: func() TLObject { // 0xd9d71866
		o := MakeTLPageBlockVideo(nil)
		o.Data2.Constructor = -640214938
		return o
	},
	972174080: func() TLObject { // 0x39f23300
		o := MakeTLPageBlockCover(nil)
		o.Data2.Constructor = 972174080
		return o
	},
	-1468953147: func() TLObject { // 0xa8718dc5
		o := MakeTLPageBlockEmbed(nil)
		o.Data2.Constructor = -1468953147
		return o
	},
	-840826671: func() TLObject { // 0xcde200d1
		o := MakeTLPageBlockEmbed(nil)
		o.Data2.Constructor = -840826671
		return o
	},
	-229005301: func() TLObject { // 0xf259a80b
		o := MakeTLPageBlockEmbedPost(nil)
		o.Data2.Constructor = -229005301
		return o
	},
	690781161: func() TLObject { // 0x292c7be9
		o := MakeTLPageBlockEmbedPost(nil)
		o.Data2.Constructor = 690781161
		return o
	},
	1705048653: func() TLObject { // 0x65a0fa4d
		o := MakeTLPageBlockCollage(nil)
		o.Data2.Constructor = 1705048653
		return o
	},
	145955919: func() TLObject { // 0x8b31c4f
		o := MakeTLPageBlockCollage(nil)
		o.Data2.Constructor = 145955919
		return o
	},
	52401552: func() TLObject { // 0x31f9590
		o := MakeTLPageBlockSlideshow(nil)
		o.Data2.Constructor = 52401552
		return o
	},
	319588707: func() TLObject { // 0x130c8963
		o := MakeTLPageBlockSlideshow(nil)
		o.Data2.Constructor = 319588707
		return o
	},
	-283684427: func() TLObject { // 0xef1751b5
		o := MakeTLPageBlockChannel(nil)
		o.Data2.Constructor = -283684427
		return o
	},
	-2143067670: func() TLObject { // 0x804361ea
		o := MakeTLPageBlockAudio(nil)
		o.Data2.Constructor = -2143067670
		return o
	},
	834148991: func() TLObject { // 0x31b81a7f
		o := MakeTLPageBlockAudio(nil)
		o.Data2.Constructor = 834148991
		return o
	},
	504660880: func() TLObject { // 0x1e148390
		o := MakeTLPageBlockKicker(nil)
		o.Data2.Constructor = 504660880
		return o
	},
	-1085412734: func() TLObject { // 0xbf4dea82
		o := MakeTLPageBlockTable(nil)
		o.Data2.Constructor = -1085412734
		return o
	},
	-1702174239: func() TLObject { // 0x9a8ae1e1
		o := MakeTLPageBlockOrderedList(nil)
		o.Data2.Constructor = -1702174239
		return o
	},
	1987480557: func() TLObject { // 0x76768bed
		o := MakeTLPageBlockDetails(nil)
		o.Data2.Constructor = 1987480557
		return o
	},
	370236054: func() TLObject { // 0x16115a96
		o := MakeTLPageBlockRelatedArticles(nil)
		o.Data2.Constructor = 370236054
		return o
	},
	-1538310410: func() TLObject { // 0xa44f3ef6
		o := MakeTLPageBlockMap(nil)
		o.Data2.Constructor = -1538310410
		return o
	},
	-421545947: func() TLObject { // 0xe6dfb825
		o := MakeTLChannelAdminLogEventActionChangeTitle(nil)
		o.Data2.Constructor = -421545947
		return o
	},
	1427671598: func() TLObject { // 0x55188a2e
		o := MakeTLChannelAdminLogEventActionChangeAbout(nil)
		o.Data2.Constructor = 1427671598
		return o
	},
	1783299128: func() TLObject { // 0x6a4afc38
		o := MakeTLChannelAdminLogEventActionChangeUsername(nil)
		o.Data2.Constructor = 1783299128
		return o
	},
	1129042607: func() TLObject { // 0x434bd2af
		o := MakeTLChannelAdminLogEventActionChangePhoto(nil)
		o.Data2.Constructor = 1129042607
		return o
	},
	-1204857405: func() TLObject { // 0xb82f55c3
		o := MakeTLChannelAdminLogEventActionChangePhoto(nil)
		o.Data2.Constructor = -1204857405
		return o
	},
	460916654: func() TLObject { // 0x1b7907ae
		o := MakeTLChannelAdminLogEventActionToggleInvites(nil)
		o.Data2.Constructor = 460916654
		return o
	},
	648939889: func() TLObject { // 0x26ae0971
		o := MakeTLChannelAdminLogEventActionToggleSignatures(nil)
		o.Data2.Constructor = 648939889
		return o
	},
	-370660328: func() TLObject { // 0xe9e82c18
		o := MakeTLChannelAdminLogEventActionUpdatePinned(nil)
		o.Data2.Constructor = -370660328
		return o
	},
	1889215493: func() TLObject { // 0x709b2405
		o := MakeTLChannelAdminLogEventActionEditMessage(nil)
		o.Data2.Constructor = 1889215493
		return o
	},
	1121994683: func() TLObject { // 0x42e047bb
		o := MakeTLChannelAdminLogEventActionDeleteMessage(nil)
		o.Data2.Constructor = 1121994683
		return o
	},
	405815507: func() TLObject { // 0x183040d3
		o := MakeTLChannelAdminLogEventActionParticipantJoin(nil)
		o.Data2.Constructor = 405815507
		return o
	},
	-124291086: func() TLObject { // 0xf89777f2
		o := MakeTLChannelAdminLogEventActionParticipantLeave(nil)
		o.Data2.Constructor = -124291086
		return o
	},
	-484690728: func() TLObject { // 0xe31c34d8
		o := MakeTLChannelAdminLogEventActionParticipantInvite(nil)
		o.Data2.Constructor = -484690728
		return o
	},
	-422036098: func() TLObject { // 0xe6d83d7e
		o := MakeTLChannelAdminLogEventActionParticipantToggleBan(nil)
		o.Data2.Constructor = -422036098
		return o
	},
	-714643696: func() TLObject { // 0xd5676710
		o := MakeTLChannelAdminLogEventActionParticipantToggleAdmin(nil)
		o.Data2.Constructor = -714643696
		return o
	},
	-1312568665: func() TLObject { // 0xb1c3caa7
		o := MakeTLChannelAdminLogEventActionChangeStickerSet(nil)
		o.Data2.Constructor = -1312568665
		return o
	},
	1599903217: func() TLObject { // 0x5f5c95f1
		o := MakeTLChannelAdminLogEventActionTogglePreHistoryHidden(nil)
		o.Data2.Constructor = 1599903217
		return o
	},
	771095562: func() TLObject { // 0x2df5fc0a
		o := MakeTLChannelAdminLogEventActionDefaultBannedRights(nil)
		o.Data2.Constructor = 771095562
		return o
	},
	-1895328189: func() TLObject { // 0x8f079643
		o := MakeTLChannelAdminLogEventActionStopPoll(nil)
		o.Data2.Constructor = -1895328189
		return o
	},
	-1569748965: func() TLObject { // 0xa26f881b
		o := MakeTLChannelAdminLogEventActionChangeLinkedChat(nil)
		o.Data2.Constructor = -1569748965
		return o
	},
	241923758: func() TLObject { // 0xe6b76ae
		o := MakeTLChannelAdminLogEventActionChangeLocation(nil)
		o.Data2.Constructor = 241923758
		return o
	},
	1401984889: func() TLObject { // 0x53909779
		o := MakeTLChannelAdminLogEventActionToggleSlowMode(nil)
		o.Data2.Constructor = 1401984889
		return o
	},
	-483352705: func() TLObject { // 0xe3309f7f
		o := MakeTLHelpTermsOfServiceUpdateEmpty(nil)
		o.Data2.Constructor = -483352705
		return o
	},
	686618977: func() TLObject { // 0x28ecf961
		o := MakeTLHelpTermsOfServiceUpdate(nil)
		o.Data2.Constructor = 686618977
		return o
	},
	831924812: func() TLObject { // 0x31962a4c
		o := MakeTLStatsGroupTopInviter(nil)
		o.Data2.Constructor = 831924812
		return o
	},
	-1142971181: func() TLObject { // 0xbbdfa4d3
		o := MakeTLVisibleTypePublic(nil)
		o.Data2.Constructor = -1142971181
		return o
	},
	73337637: func() TLObject { // 0x45f0b25
		o := MakeTLVisibleTypePrivate(nil)
		o.Data2.Constructor = 73337637
		return o
	},
	1872566708: func() TLObject { // 0x6f9d19b4
		o := MakeTLVisibleTypeFriend(nil)
		o.Data2.Constructor = 1872566708
		return o
	},
	902771484: func() TLObject { // 0x35cf331c
		o := MakeTLVisibleTypeFollow(nil)
		o.Data2.Constructor = 902771484
		return o
	},
	-302209623: func() TLObject { // 0xedfca5a9
		o := MakeTLVisibleTypeFans(nil)
		o.Data2.Constructor = -302209623
		return o
	},
	599409422: func() TLObject { // 0x23ba430e
		o := MakeTLVisibleTypeUser(nil)
		o.Data2.Constructor = 599409422
		return o
	},
	-812171549: func() TLObject { // 0xcf973ee3
		o := MakeTLVisibleTypeAllow(nil)
		o.Data2.Constructor = -812171549
		return o
	},
	1482668631: func() TLObject { // 0x585fba57
		o := MakeTLVisibleTypeNotAllow(nil)
		o.Data2.Constructor = 1482668631
		return o
	},
	-1095378172: func() TLObject { // 0xbeb5db04
		o := MakeTLVisibleTypeTopic(nil)
		o.Data2.Constructor = -1095378172
		return o
	},
	-1078612597: func() TLObject { // 0xbfb5ad8b
		o := MakeTLChannelLocationEmpty(nil)
		o.Data2.Constructor = -1078612597
		return o
	},
	547062491: func() TLObject { // 0x209b82db
		o := MakeTLChannelLocation(nil)
		o.Data2.Constructor = 547062491
		return o
	},
	2012136335: func() TLObject { // 0x77eec38f
		o := MakeTLCdnFileHash(nil)
		o.Data2.Constructor = 2012136335
		return o
	},
	-74456004: func() TLObject { // 0xfb8fe43c
		o := MakeTLPaymentsSavedInfo(nil)
		o.Data2.Constructor = -74456004
		return o
	},
	-842892769: func() TLObject { // 0xcdc27a1f
		o := MakeTLPaymentSavedCredentialsCard(nil)
		o.Data2.Constructor = -842892769
		return o
	},
	-1132476723: func() TLObject { // 0xbc7fc6cd
		o := MakeTLFileLocationToBeDeprecated(nil)
		o.Data2.Constructor = -1132476723
		return o
	},
	2086234950: func() TLObject { // 0x7c596b46
		o := MakeTLFileLocationUnavailable(nil)
		o.Data2.Constructor = 2086234950
		return o
	},
	152900075: func() TLObject { // 0x91d11eb
		o := MakeTLFileLocation(nil)
		o.Data2.Constructor = 152900075
		return o
	},
	1406570614: func() TLObject { // 0x53d69076
		o := MakeTLFileLocation(nil)
		o.Data2.Constructor = 1406570614
		return o
	},
	1984136919: func() TLObject { // 0x764386d7
		o := MakeTLWalletLiteResponse(nil)
		o.Data2.Constructor = 1984136919
		return o
	},
	1251549527: func() TLObject { // 0x4a992157
		o := MakeTLInputStickeredMediaPhoto(nil)
		o.Data2.Constructor = 1251549527
		return o
	},
	70813275: func() TLObject { // 0x438865b
		o := MakeTLInputStickeredMediaDocument(nil)
		o.Data2.Constructor = 70813275
		return o
	},
	512535275: func() TLObject { // 0x1e8caaeb
		o := MakeTLPostAddress(nil)
		o.Data2.Constructor = 512535275
		return o
	},
	-1361650766: func() TLObject { // 0xaed6dbb2
		o := MakeTLMaskCoords(nil)
		o.Data2.Constructor = -1361650766
		return o
	},
	568808380: func() TLObject { // 0x21e753bc
		o := MakeTLUploadWebFile(nil)
		o.Data2.Constructor = 568808380
		return o
	},
	-313079300: func() TLObject { // 0xed56c9fc
		o := MakeTLAccountWebAuthorizations(nil)
		o.Data2.Constructor = -313079300
		return o
	},
	-1275374751: func() TLObject { // 0xb3fb5361
		o := MakeTLEmojiLanguage(nil)
		o.Data2.Constructor = -1275374751
		return o
	},
	-539317279: func() TLObject { // 0xdfdaabe1
		o := MakeTLInputFileLocation(nil)
		o.Data2.Constructor = -539317279
		return o
	},
	342061462: func() TLObject { // 0x14637196
		o := MakeTLInputFileLocation(nil)
		o.Data2.Constructor = 342061462
		return o
	},
	-182231723: func() TLObject { // 0xf5235d55
		o := MakeTLInputEncryptedFileLocation(nil)
		o.Data2.Constructor = -182231723
		return o
	},
	-1160743548: func() TLObject { // 0xbad07584
		o := MakeTLInputDocumentFileLocation(nil)
		o.Data2.Constructor = -1160743548
		return o
	},
	426148825: func() TLObject { // 0x196683d9
		o := MakeTLInputDocumentFileLocation(nil)
		o.Data2.Constructor = 426148825
		return o
	},
	1125058340: func() TLObject { // 0x430f0724
		o := MakeTLInputDocumentFileLocation(nil)
		o.Data2.Constructor = 1125058340
		return o
	},
	1313188841: func() TLObject { // 0x4e45abe9
		o := MakeTLInputDocumentFileLocation(nil)
		o.Data2.Constructor = 1313188841
		return o
	},
	-876089816: func() TLObject { // 0xcbc7ee28
		o := MakeTLInputSecureFileLocation(nil)
		o.Data2.Constructor = -876089816
		return o
	},
	700340377: func() TLObject { // 0x29be5899
		o := MakeTLInputTakeoutFileLocation(nil)
		o.Data2.Constructor = 700340377
		return o
	},
	1075322878: func() TLObject { // 0x40181ffe
		o := MakeTLInputPhotoFileLocation(nil)
		o.Data2.Constructor = 1075322878
		return o
	},
	-667654413: func() TLObject { // 0xd83466f3
		o := MakeTLInputPhotoLegacyFileLocation(nil)
		o.Data2.Constructor = -667654413
		return o
	},
	668375447: func() TLObject { // 0x27d69997
		o := MakeTLInputPeerPhotoFileLocation(nil)
		o.Data2.Constructor = 668375447
		return o
	},
	230353641: func() TLObject { // 0xdbaeae9
		o := MakeTLInputStickerSetThumb(nil)
		o.Data2.Constructor = 230353641
		return o
	},
	-1705233435: func() TLObject { // 0x9a5c33e5
		o := MakeTLAccountPasswordSettings(nil)
		o.Data2.Constructor = -1705233435
		return o
	},
	2077869041: func() TLObject { // 0x7bd9c3f1
		o := MakeTLAccountPasswordSettings(nil)
		o.Data2.Constructor = 2077869041
		return o
	},
	-1212732749: func() TLObject { // 0xb7b72ab3
		o := MakeTLAccountPasswordSettings(nil)
		o.Data2.Constructor = -1212732749
		return o
	},
	186120336: func() TLObject { // 0xb17f890
		o := MakeTLMessagesRecentStickersNotModified(nil)
		o.Data2.Constructor = 186120336
		return o
	},
	586395571: func() TLObject { // 0x22f3afb3
		o := MakeTLMessagesRecentStickers(nil)
		o.Data2.Constructor = 586395571
		return o
	},
	1558317424: func() TLObject { // 0x5ce20970
		o := MakeTLMessagesRecentStickers(nil)
		o.Data2.Constructor = 1558317424
		return o
	},
	-6249322: func() TLObject { // 0xffa0a496
		o := MakeTLInputStickerSetItem(nil)
		o.Data2.Constructor = -6249322
		return o
	},
	-1389486888: func() TLObject { // 0xad2e1cd8
		o := MakeTLAccountAuthorizationForm(nil)
		o.Data2.Constructor = -1389486888
		return o
	},
	-879268525: func() TLObject { // 0xcb976d53
		o := MakeTLAccountAuthorizationForm(nil)
		o.Data2.Constructor = -879268525
		return o
	},
	1869903447: func() TLObject { // 0x6f747657
		o := MakeTLPageCaption(nil)
		o.Data2.Constructor = 1869903447
		return o
	},
	-264117680: func() TLObject { // 0xf041e250
		o := MakeTLChatOnlines(nil)
		o.Data2.Constructor = -264117680
		return o
	},
	1674235686: func() TLObject { // 0x63cacf26
		o := MakeTLAccountAutoDownloadSettings(nil)
		o.Data2.Constructor = 1674235686
		return o
	},
	-976602448: func() TLObject { // 0xc5ca3ab0
		o := MakeTLBlogsTopics(nil)
		o.Data2.Constructor = -976602448
		return o
	},
	1158290442: func() TLObject { // 0x450a1c0a
		o := MakeTLMessagesFoundGifs(nil)
		o.Data2.Constructor = 1158290442
		return o
	},
	240579230: func() TLObject { // 0xe56f29e
		o := MakeTLBlogsComment(nil)
		o.Data2.Constructor = 240579230
		return o
	},
	539045032: func() TLObject { // 0x20212ca8
		o := MakeTLPhotosPhoto(nil)
		o.Data2.Constructor = 539045032
		return o
	},
	-1995686519: func() TLObject { // 0x890c3d89
		o := MakeTLInputBotInlineMessageID(nil)
		o.Data2.Constructor = -1995686519
		return o
	},
	1107543535: func() TLObject { // 0x4203c5ef
		o := MakeTLHelpCountryCode(nil)
		o.Data2.Constructor = 1107543535
		return o
	},
	-582464156: func() TLObject { // 0xdd484d64
		o := MakeTLWalletSecretSalt(nil)
		o.Data2.Constructor = -582464156
		return o
	},
	-1377390588: func() TLObject { // 0xade6b004
		o := MakeTLInputPhotoCropAuto(nil)
		o.Data2.Constructor = -1377390588
		return o
	},
	-644787419: func() TLObject { // 0xd9915325
		o := MakeTLInputPhotoCrop(nil)
		o.Data2.Constructor = -644787419
		return o
	},
	-1182234929: func() TLObject { // 0xb98886cf
		o := MakeTLInputUserEmpty(nil)
		o.Data2.Constructor = -1182234929
		return o
	},
	-138301121: func() TLObject { // 0xf7c1b13f
		o := MakeTLInputUserSelf(nil)
		o.Data2.Constructor = -138301121
		return o
	},
	-668391402: func() TLObject { // 0xd8292816
		o := MakeTLInputUser(nil)
		o.Data2.Constructor = -668391402
		return o
	},
	756118935: func() TLObject { // 0x2d117597
		o := MakeTLInputUserFromMessage(nil)
		o.Data2.Constructor = 756118935
		return o
	},
	415997816: func() TLObject { // 0x18cb9f78
		o := MakeTLHelpInviteText(nil)
		o.Data2.Constructor = 415997816
		return o
	},
	-1634752813: func() TLObject { // 0x9e8fa6d3
		o := MakeTLMessagesFavedStickersNotModified(nil)
		o.Data2.Constructor = -1634752813
		return o
	},
	-209768682: func() TLObject { // 0xf37f2f16
		o := MakeTLMessagesFavedStickers(nil)
		o.Data2.Constructor = -209768682
		return o
	},
	-391902247: func() TLObject { // 0xe8a40bd9
		o := MakeTLSecureValueErrorData(nil)
		o.Data2.Constructor = -391902247
		return o
	},
	12467706: func() TLObject { // 0xbe3dfa
		o := MakeTLSecureValueErrorFrontSide(nil)
		o.Data2.Constructor = 12467706
		return o
	},
	-2037765467: func() TLObject { // 0x868a2aa5
		o := MakeTLSecureValueErrorReverseSide(nil)
		o.Data2.Constructor = -2037765467
		return o
	},
	-449327402: func() TLObject { // 0xe537ced6
		o := MakeTLSecureValueErrorSelfie(nil)
		o.Data2.Constructor = -449327402
		return o
	},
	2054162547: func() TLObject { // 0x7a700873
		o := MakeTLSecureValueErrorFile(nil)
		o.Data2.Constructor = 2054162547
		return o
	},
	1717706985: func() TLObject { // 0x666220e9
		o := MakeTLSecureValueErrorFiles(nil)
		o.Data2.Constructor = 1717706985
		return o
	},
	-2036501105: func() TLObject { // 0x869d758f
		o := MakeTLSecureValueError(nil)
		o.Data2.Constructor = -2036501105
		return o
	},
	-1592506512: func() TLObject { // 0xa1144770
		o := MakeTLSecureValueErrorTranslationFile(nil)
		o.Data2.Constructor = -1592506512
		return o
	},
	878931416: func() TLObject { // 0x34636dd8
		o := MakeTLSecureValueErrorTranslationFiles(nil)
		o.Data2.Constructor = 878931416
		return o
	},
	-1012849566: func() TLObject { // 0xc3a12462
		o := MakeTLBaseThemeClassic(nil)
		o.Data2.Constructor = -1012849566
		return o
	},
	-69724536: func() TLObject { // 0xfbd81688
		o := MakeTLBaseThemeDay(nil)
		o.Data2.Constructor = -69724536
		return o
	},
	-1212997976: func() TLObject { // 0xb7b31ea8
		o := MakeTLBaseThemeNight(nil)
		o.Data2.Constructor = -1212997976
		return o
	},
	1834973166: func() TLObject { // 0x6d5f77ee
		o := MakeTLBaseThemeTinted(nil)
		o.Data2.Constructor = 1834973166
		return o
	},
	1527845466: func() TLObject { // 0x5b11125a
		o := MakeTLBaseThemeArctic(nil)
		o.Data2.Constructor = 1527845466
		return o
	},
	-1461589623: func() TLObject { // 0xa8e1e989
		o := MakeTLSchemeType(nil)
		o.Data2.Constructor = -1461589623
		return o
	},
	-805141448: func() TLObject { // 0xd0028438
		o := MakeTLImportedContact(nil)
		o.Data2.Constructor = -805141448
		return o
	},
	-445792507: func() TLObject { // 0xe56dbf05
		o := MakeTLDialogPeer(nil)
		o.Data2.Constructor = -445792507
		return o
	},
	1363483106: func() TLObject { // 0x514519e2
		o := MakeTLDialogPeerFolder(nil)
		o.Data2.Constructor = 1363483106
		return o
	},
	-633170927: func() TLObject { // 0xda429411
		o := MakeTLDialogPeerFeed(nil)
		o.Data2.Constructor = -633170927
		return o
	},
	-1456996667: func() TLObject { // 0xa927fec5
		o := MakeTLMessagesInactiveChats(nil)
		o.Data2.Constructor = -1456996667
		return o
	},
	-1815339214: func() TLObject { // 0x93cc1f32
		o := MakeTLHelpCountriesListNotModified(nil)
		o.Data2.Constructor = -1815339214
		return o
	},
	-2016381538: func() TLObject { // 0x87d0759e
		o := MakeTLHelpCountriesList(nil)
		o.Data2.Constructor = -2016381538
		return o
	},
	1352683077: func() TLObject { // 0x50a04e45
		o := MakeTLAccountPrivacyRules(nil)
		o.Data2.Constructor = 1352683077
		return o
	},
	1430961007: func() TLObject { // 0x554abb6f
		o := MakeTLAccountPrivacyRules(nil)
		o.Data2.Constructor = 1430961007
		return o
	},
	-1194283041: func() TLObject { // 0xb8d0afdf
		o := MakeTLAccountDaysTTL(nil)
		o.Data2.Constructor = -1194283041
		return o
	},
	-1107729093: func() TLObject { // 0xbdf9653b
		o := MakeTLGame(nil)
		o.Data2.Constructor = -1107729093
		return o
	},
	-290921362: func() TLObject { // 0xeea8e46e
		o := MakeTLUploadCdnFileReuploadNeeded(nil)
		o.Data2.Constructor = -290921362
		return o
	},
	-1449145777: func() TLObject { // 0xa99fca4f
		o := MakeTLUploadCdnFile(nil)
		o.Data2.Constructor = -1449145777
		return o
	},
	-206688531: func() TLObject { // 0xf3ae2eed
		o := MakeTLHelpUserInfoEmpty(nil)
		o.Data2.Constructor = -206688531
		return o
	},
	32192344: func() TLObject { // 0x1eb3758
		o := MakeTLHelpUserInfo(nil)
		o.Data2.Constructor = 32192344
		return o
	},
	-70073706: func() TLObject { // 0xfbd2c296
		o := MakeTLInputFolderPeer(nil)
		o.Data2.Constructor = -70073706
		return o
	},
	-1519637954: func() TLObject { // 0xa56c2a3e
		o := MakeTLUpdatesState(nil)
		o.Data2.Constructor = -1519637954
		return o
	},
	-1230047312: func() TLObject { // 0xb6aef7b0
		o := MakeTLMessageActionEmpty(nil)
		o.Data2.Constructor = -1230047312
		return o
	},
	-1503425638: func() TLObject { // 0xa6638b9a
		o := MakeTLMessageActionChatCreate(nil)
		o.Data2.Constructor = -1503425638
		return o
	},
	-1247687078: func() TLObject { // 0xb5a1ce5a
		o := MakeTLMessageActionChatEditTitle(nil)
		o.Data2.Constructor = -1247687078
		return o
	},
	2144015272: func() TLObject { // 0x7fcb13a8
		o := MakeTLMessageActionChatEditPhoto(nil)
		o.Data2.Constructor = 2144015272
		return o
	},
	-1780220945: func() TLObject { // 0x95e3fbef
		o := MakeTLMessageActionChatDeletePhoto(nil)
		o.Data2.Constructor = -1780220945
		return o
	},
	1217033015: func() TLObject { // 0x488a7337
		o := MakeTLMessageActionChatAddUser(nil)
		o.Data2.Constructor = 1217033015
		return o
	},
	-1297179892: func() TLObject { // 0xb2ae9b0c
		o := MakeTLMessageActionChatDeleteUser(nil)
		o.Data2.Constructor = -1297179892
		return o
	},
	-123931160: func() TLObject { // 0xf89cf5e8
		o := MakeTLMessageActionChatJoinedByLink(nil)
		o.Data2.Constructor = -123931160
		return o
	},
	-1781355374: func() TLObject { // 0x95d2ac92
		o := MakeTLMessageActionChannelCreate(nil)
		o.Data2.Constructor = -1781355374
		return o
	},
	1371385889: func() TLObject { // 0x51bdb021
		o := MakeTLMessageActionChatMigrateTo(nil)
		o.Data2.Constructor = 1371385889
		return o
	},
	-1336546578: func() TLObject { // 0xb055eaee
		o := MakeTLMessageActionChannelMigrateFrom(nil)
		o.Data2.Constructor = -1336546578
		return o
	},
	-1799538451: func() TLObject { // 0x94bd38ed
		o := MakeTLMessageActionPinMessage(nil)
		o.Data2.Constructor = -1799538451
		return o
	},
	-1615153660: func() TLObject { // 0x9fbab604
		o := MakeTLMessageActionHistoryClear(nil)
		o.Data2.Constructor = -1615153660
		return o
	},
	-1834538890: func() TLObject { // 0x92a72876
		o := MakeTLMessageActionGameScore(nil)
		o.Data2.Constructor = -1834538890
		return o
	},
	-1892568281: func() TLObject { // 0x8f31b327
		o := MakeTLMessageActionPaymentSentMe(nil)
		o.Data2.Constructor = -1892568281
		return o
	},
	1080663248: func() TLObject { // 0x40699cd0
		o := MakeTLMessageActionPaymentSent(nil)
		o.Data2.Constructor = 1080663248
		return o
	},
	-2132731265: func() TLObject { // 0x80e11a7f
		o := MakeTLMessageActionPhoneCall(nil)
		o.Data2.Constructor = -2132731265
		return o
	},
	1200788123: func() TLObject { // 0x4792929b
		o := MakeTLMessageActionScreenshotTaken(nil)
		o.Data2.Constructor = 1200788123
		return o
	},
	-85549226: func() TLObject { // 0xfae69f56
		o := MakeTLMessageActionCustomAction(nil)
		o.Data2.Constructor = -85549226
		return o
	},
	-1410748418: func() TLObject { // 0xabe9affe
		o := MakeTLMessageActionBotAllowed(nil)
		o.Data2.Constructor = -1410748418
		return o
	},
	455635795: func() TLObject { // 0x1b287353
		o := MakeTLMessageActionSecureValuesSentMe(nil)
		o.Data2.Constructor = 455635795
		return o
	},
	-648257196: func() TLObject { // 0xd95c6154
		o := MakeTLMessageActionSecureValuesSent(nil)
		o.Data2.Constructor = -648257196
		return o
	},
	-202219658: func() TLObject { // 0xf3f25f76
		o := MakeTLMessageActionContactSignUp(nil)
		o.Data2.Constructor = -202219658
		return o
	},
	1894744724: func() TLObject { // 0x70ef8294
		o := MakeTLMessageActionContactSignUp(nil)
		o.Data2.Constructor = 1894744724
		return o
	},
	-1730095465: func() TLObject { // 0x98e0d697
		o := MakeTLMessageActionGeoProximityReached(nil)
		o.Data2.Constructor = -1730095465
		return o
	},
	805171639: func() TLObject { // 0x2ffdf1b7
		o := MakeTLMessageActionBizDataRaw(nil)
		o.Data2.Constructor = 805171639
		return o
	},
	-994830336: func() TLObject { // 0xc4b41800
		o := MakeTLBlogsUserDate(nil)
		o.Data2.Constructor = -994830336
		return o
	},
	-1738178803: func() TLObject { // 0x98657f0d
		o := MakeTLPage(nil)
		o.Data2.Constructor = -1738178803
		return o
	},
	-1366746132: func() TLObject { // 0xae891bec
		o := MakeTLPage(nil)
		o.Data2.Constructor = -1366746132
		return o
	},
	-241590104: func() TLObject { // 0xf199a0a8
		o := MakeTLPage(nil)
		o.Data2.Constructor = -241590104
		return o
	},
	-1908433218: func() TLObject { // 0x8e3f9ebe
		o := MakeTLPagePart(nil)
		o.Data2.Constructor = -1908433218
		return o
	},
	1433323434: func() TLObject { // 0x556ec7aa
		o := MakeTLPageFull(nil)
		o.Data2.Constructor = 1433323434
		return o
	},
	-457104426: func() TLObject { // 0xe4c123d6
		o := MakeTLInputGeoPointEmpty(nil)
		o.Data2.Constructor = -457104426
		return o
	},
	1210199983: func() TLObject { // 0x48222faf
		o := MakeTLInputGeoPoint(nil)
		o.Data2.Constructor = 1210199983
		return o
	},
	-206066487: func() TLObject { // 0xf3b7acc9
		o := MakeTLInputGeoPoint(nil)
		o.Data2.Constructor = -206066487
		return o
	},
	-566281095: func() TLObject { // 0xde3f3c79
		o := MakeTLChannelParticipantsRecent(nil)
		o.Data2.Constructor = -566281095
		return o
	},
	-1268741783: func() TLObject { // 0xb4608969
		o := MakeTLChannelParticipantsAdmins(nil)
		o.Data2.Constructor = -1268741783
		return o
	},
	-1548400251: func() TLObject { // 0xa3b54985
		o := MakeTLChannelParticipantsKicked(nil)
		o.Data2.Constructor = -1548400251
		return o
	},
	-1328445861: func() TLObject { // 0xb0d1865b
		o := MakeTLChannelParticipantsBots(nil)
		o.Data2.Constructor = -1328445861
		return o
	},
	338142689: func() TLObject { // 0x1427a5e1
		o := MakeTLChannelParticipantsBanned(nil)
		o.Data2.Constructor = 338142689
		return o
	},
	106343499: func() TLObject { // 0x656ac4b
		o := MakeTLChannelParticipantsSearch(nil)
		o.Data2.Constructor = 106343499
		return o
	},
	-1150621555: func() TLObject { // 0xbb6ae88d
		o := MakeTLChannelParticipantsContacts(nil)
		o.Data2.Constructor = -1150621555
		return o
	},
	-531931925: func() TLObject { // 0xe04b5ceb
		o := MakeTLChannelParticipantsMentions(nil)
		o.Data2.Constructor = -531931925
		return o
	},
	1035688326: func() TLObject { // 0x3dbb5986
		o := MakeTLAuthSentCodeTypeApp(nil)
		o.Data2.Constructor = 1035688326
		return o
	},
	-1073693790: func() TLObject { // 0xc000bba2
		o := MakeTLAuthSentCodeTypeSms(nil)
		o.Data2.Constructor = -1073693790
		return o
	},
	1398007207: func() TLObject { // 0x5353e5a7
		o := MakeTLAuthSentCodeTypeCall(nil)
		o.Data2.Constructor = 1398007207
		return o
	},
	-1425815847: func() TLObject { // 0xab03c6d9
		o := MakeTLAuthSentCodeTypeFlashCall(nil)
		o.Data2.Constructor = -1425815847
		return o
	},
	1431132616: func() TLObject { // 0x554d59c8
		o := MakeTLDouble(nil)
		o.Data2.Constructor = 1431132616
		return o
	},
	-395967805: func() TLObject { // 0xe86602c3
		o := MakeTLMessagesAllStickersNotModified(nil)
		o.Data2.Constructor = -395967805
		return o
	},
	-302170017: func() TLObject { // 0xedfd405f
		o := MakeTLMessagesAllStickers(nil)
		o.Data2.Constructor = -302170017
		return o
	},
	1342771681: func() TLObject { // 0x500911e1
		o := MakeTLPaymentsPaymentReceipt(nil)
		o.Data2.Constructor = 1342771681
		return o
	},
	411017418: func() TLObject { // 0x187fa0ca
		o := MakeTLSecureValue(nil)
		o.Data2.Constructor = 411017418
		return o
	},
	-1263225191: func() TLObject { // 0xb4b4b699
		o := MakeTLSecureValue(nil)
		o.Data2.Constructor = -1263225191
		return o
	},
	-331270968: func() TLObject { // 0xec4134c8
		o := MakeTLSecureValue(nil)
		o.Data2.Constructor = -331270968
		return o
	},
	-1736378792: func() TLObject { // 0x9880f658
		o := MakeTLInputCheckPasswordEmpty(nil)
		o.Data2.Constructor = -1736378792
		return o
	},
	-763367294: func() TLObject { // 0xd27ff082
		o := MakeTLInputCheckPasswordSRP(nil)
		o.Data2.Constructor = -763367294
		return o
	},
	-199313886: func() TLObject { // 0xf41eb622
		o := MakeTLAccountThemesNotModified(nil)
		o.Data2.Constructor = -199313886
		return o
	},
	2137482273: func() TLObject { // 0x7f676421
		o := MakeTLAccountThemes(nil)
		o.Data2.Constructor = 2137482273
		return o
	},
	-1728664459: func() TLObject { // 0x98f6ac75
		o := MakeTLHelpPromoDataEmpty(nil)
		o.Data2.Constructor = -1728664459
		return o
	},
	-1942390465: func() TLObject { // 0x8c39793f
		o := MakeTLHelpPromoData(nil)
		o.Data2.Constructor = -1942390465
		return o
	},
	418631927: func() TLObject { // 0x18f3d0f7
		o := MakeTLStatsGroupTopPoster(nil)
		o.Data2.Constructor = 418631927
		return o
	},
	1568467877: func() TLObject { // 0x5d7ceba5
		o := MakeTLChannelAdminRights(nil)
		o.Data2.Constructor = 1568467877
		return o
	},
	-1353671392: func() TLObject { // 0xaf509d20
		o := MakeTLPeerNotifySettings(nil)
		o.Data2.Constructor = -1353671392
		return o
	},
	-1697798976: func() TLObject { // 0x9acda4c0
		o := MakeTLPeerNotifySettings(nil)
		o.Data2.Constructor = -1697798976
		return o
	},
	1889961234: func() TLObject { // 0x70a68512
		o := MakeTLPeerNotifySettingsEmpty(nil)
		o.Data2.Constructor = 1889961234
		return o
	},
	-57668565: func() TLObject { // 0xfc900c2b
		o := MakeTLChatParticipantsForbidden(nil)
		o.Data2.Constructor = -57668565
		return o
	},
	1061556205: func() TLObject { // 0x3f460fed
		o := MakeTLChatParticipants(nil)
		o.Data2.Constructor = 1061556205
		return o
	},
	286776671: func() TLObject { // 0x1117dd5f
		o := MakeTLGeoPointEmpty(nil)
		o.Data2.Constructor = 286776671
		return o
	},
	-1297942941: func() TLObject { // 0xb2a2f663
		o := MakeTLGeoPoint(nil)
		o.Data2.Constructor = -1297942941
		return o
	},
	43446532: func() TLObject { // 0x296f104
		o := MakeTLGeoPoint(nil)
		o.Data2.Constructor = 43446532
		return o
	},
	541710092: func() TLObject { // 0x2049d70c
		o := MakeTLGeoPoint(nil)
		o.Data2.Constructor = 541710092
		return o
	},
	-1916114267: func() TLObject { // 0x8dca6aa5
		o := MakeTLPhotosPhotos(nil)
		o.Data2.Constructor = -1916114267
		return o
	},
	352657236: func() TLObject { // 0x15051f54
		o := MakeTLPhotosPhotosSlice(nil)
		o.Data2.Constructor = 352657236
		return o
	},
	313694676: func() TLObject { // 0x12b299d4
		o := MakeTLStickerPack(nil)
		o.Data2.Constructor = 313694676
		return o
	},
	2103482845: func() TLObject { // 0x7d6099dd
		o := MakeTLSecurePlainPhone(nil)
		o.Data2.Constructor = 2103482845
		return o
	},
	569137759: func() TLObject { // 0x21ec5a5f
		o := MakeTLSecurePlainEmail(nil)
		o.Data2.Constructor = 569137759
		return o
	},
	-524237339: func() TLObject { // 0xe0c0c5e5
		o := MakeTLPageTableRow(nil)
		o.Data2.Constructor = -524237339
		return o
	},
	997055186: func() TLObject { // 0x3b6ddad2
		o := MakeTLPollAnswerVoters(nil)
		o.Data2.Constructor = 997055186
		return o
	},
	-1771768449: func() TLObject { // 0x9664f57f
		o := MakeTLInputMediaEmpty(nil)
		o.Data2.Constructor = -1771768449
		return o
	},
	505969924: func() TLObject { // 0x1e287d04
		o := MakeTLInputMediaUploadedPhoto(nil)
		o.Data2.Constructor = 505969924
		return o
	},
	792191537: func() TLObject { // 0x2f37e231
		o := MakeTLInputMediaUploadedPhoto(nil)
		o.Data2.Constructor = 792191537
		return o
	},
	-1279654347: func() TLObject { // 0xb3ba0635
		o := MakeTLInputMediaPhoto(nil)
		o.Data2.Constructor = -1279654347
		return o
	},
	-2114308294: func() TLObject { // 0x81fa373a
		o := MakeTLInputMediaPhoto(nil)
		o.Data2.Constructor = -2114308294
		return o
	},
	-373312269: func() TLObject { // 0xe9bfb4f3
		o := MakeTLInputMediaPhoto(nil)
		o.Data2.Constructor = -373312269
		return o
	},
	-104578748: func() TLObject { // 0xf9c44144
		o := MakeTLInputMediaGeoPoint(nil)
		o.Data2.Constructor = -104578748
		return o
	},
	-122978821: func() TLObject { // 0xf8ab7dfb
		o := MakeTLInputMediaContact(nil)
		o.Data2.Constructor = -122978821
		return o
	},
	-1494984313: func() TLObject { // 0xa6e45987
		o := MakeTLInputMediaContact(nil)
		o.Data2.Constructor = -1494984313
		return o
	},
	1530447553: func() TLObject { // 0x5b38c6c1
		o := MakeTLInputMediaUploadedDocument(nil)
		o.Data2.Constructor = 1530447553
		return o
	},
	-476700163: func() TLObject { // 0xe39621fd
		o := MakeTLInputMediaUploadedDocument(nil)
		o.Data2.Constructor = -476700163
		return o
	},
	598418386: func() TLObject { // 0x23ab23d2
		o := MakeTLInputMediaDocument(nil)
		o.Data2.Constructor = 598418386
		return o
	},
	1523279502: func() TLObject { // 0x5acb668e
		o := MakeTLInputMediaDocument(nil)
		o.Data2.Constructor = 1523279502
		return o
	},
	-1052959727: func() TLObject { // 0xc13d1c11
		o := MakeTLInputMediaVenue(nil)
		o.Data2.Constructor = -1052959727
		return o
	},
	-440664550: func() TLObject { // 0xe5bbfe1a
		o := MakeTLInputMediaPhotoExternal(nil)
		o.Data2.Constructor = -440664550
		return o
	},
	153267905: func() TLObject { // 0x922aec1
		o := MakeTLInputMediaPhotoExternal(nil)
		o.Data2.Constructor = 153267905
		return o
	},
	-78455655: func() TLObject { // 0xfb52dc99
		o := MakeTLInputMediaDocumentExternal(nil)
		o.Data2.Constructor = -78455655
		return o
	},
	-1225309387: func() TLObject { // 0xb6f74335
		o := MakeTLInputMediaDocumentExternal(nil)
		o.Data2.Constructor = -1225309387
		return o
	},
	-750828557: func() TLObject { // 0xd33f43f3
		o := MakeTLInputMediaGame(nil)
		o.Data2.Constructor = -750828557
		return o
	},
	-186607933: func() TLObject { // 0xf4e096c3
		o := MakeTLInputMediaInvoice(nil)
		o.Data2.Constructor = -186607933
		return o
	},
	-1759532989: func() TLObject { // 0x971fa843
		o := MakeTLInputMediaGeoLive(nil)
		o.Data2.Constructor = -1759532989
		return o
	},
	-833715459: func() TLObject { // 0xce4e82fd
		o := MakeTLInputMediaGeoLive(nil)
		o.Data2.Constructor = -833715459
		return o
	},
	2065305999: func() TLObject { // 0x7b1a118f
		o := MakeTLInputMediaGeoLive(nil)
		o.Data2.Constructor = 2065305999
		return o
	},
	261416433: func() TLObject { // 0xf94e5f1
		o := MakeTLInputMediaPoll(nil)
		o.Data2.Constructor = 261416433
		return o
	},
	-1410741723: func() TLObject { // 0xabe9ca25
		o := MakeTLInputMediaPoll(nil)
		o.Data2.Constructor = -1410741723
		return o
	},
	112424539: func() TLObject { // 0x6b3765b
		o := MakeTLInputMediaPoll(nil)
		o.Data2.Constructor = 112424539
		return o
	},
	-428884101: func() TLObject { // 0xe66fbf7b
		o := MakeTLInputMediaDice(nil)
		o.Data2.Constructor = -428884101
		return o
	},
	-1358977017: func() TLObject { // 0xaeffa807
		o := MakeTLInputMediaDice(nil)
		o.Data2.Constructor = -1358977017
		return o
	},
	1212395773: func() TLObject { // 0x4843b0fd
		o := MakeTLInputMediaGifExternal(nil)
		o.Data2.Constructor = 1212395773
		return o
	},
	-1097470438: func() TLObject { // 0xbe95ee1a
		o := MakeTLInputMediaBizDataRaw(nil)
		o.Data2.Constructor = -1097470438
		return o
	},
	1244130093: func() TLObject { // 0x4a27eb2d
		o := MakeTLStatsGraphAsync(nil)
		o.Data2.Constructor = 1244130093
		return o
	},
	-1092839390: func() TLObject { // 0xbedc9822
		o := MakeTLStatsGraphError(nil)
		o.Data2.Constructor = -1092839390
		return o
	},
	-1901828938: func() TLObject { // 0x8ea464b6
		o := MakeTLStatsGraph(nil)
		o.Data2.Constructor = -1901828938
		return o
	},
	-1626209256: func() TLObject { // 0x9f120418
		o := MakeTLChatBannedRights(nil)
		o.Data2.Constructor = -1626209256
		return o
	},
	-1142327219: func() TLObject { // 0xbbe9784d
		o := MakeTLInputBlogPhoto(nil)
		o.Data2.Constructor = -1142327219
		return o
	},
	2134011769: func() TLObject { // 0x7f326f79
		o := MakeTLInputBlogPhotoFile(nil)
		o.Data2.Constructor = 2134011769
		return o
	},
	-1036396922: func() TLObject { // 0xc239d686
		o := MakeTLInputWebFileLocation(nil)
		o.Data2.Constructor = -1036396922
		return o
	},
	-1625153079: func() TLObject { // 0x9f2221c9
		o := MakeTLInputWebFileGeoPointLocation(nil)
		o.Data2.Constructor = -1625153079
		return o
	},
	1713855074: func() TLObject { // 0x66275a62
		o := MakeTLInputWebFileGeoPointLocation(nil)
		o.Data2.Constructor = 1713855074
		return o
	},
	1430205163: func() TLObject { // 0x553f32eb
		o := MakeTLInputWebFileGeoMessageLocation(nil)
		o.Data2.Constructor = 1430205163
		return o
	},
	-1502174430: func() TLObject { // 0xa676a322
		o := MakeTLInputMessageID(nil)
		o.Data2.Constructor = -1502174430
		return o
	},
	-1160215659: func() TLObject { // 0xbad88395
		o := MakeTLInputMessageReplyTo(nil)
		o.Data2.Constructor = -1160215659
		return o
	},
	-2037963464: func() TLObject { // 0x86872538
		o := MakeTLInputMessagePinned(nil)
		o.Data2.Constructor = -2037963464
		return o
	},
	-1392895362: func() TLObject { // 0xacfa1a7e
		o := MakeTLInputMessageCallbackQuery(nil)
		o.Data2.Constructor = -1392895362
		return o
	},
	1072550713: func() TLObject { // 0x3fedd339
		o := MakeTLTrue(nil)
		o.Data2.Constructor = 1072550713
		return o
	},
	864077702: func() TLObject { // 0x3380c786
		o := MakeTLInputBotInlineMessageMediaAuto(nil)
		o.Data2.Constructor = 864077702
		return o
	},
	691006739: func() TLObject { // 0x292fed13
		o := MakeTLInputBotInlineMessageMediaAuto(nil)
		o.Data2.Constructor = 691006739
		return o
	},
	1036876423: func() TLObject { // 0x3dcd7a87
		o := MakeTLInputBotInlineMessageText(nil)
		o.Data2.Constructor = 1036876423
		return o
	},
	-1768777083: func() TLObject { // 0x96929a85
		o := MakeTLInputBotInlineMessageMediaGeo(nil)
		o.Data2.Constructor = -1768777083
		return o
	},
	-1045340827: func() TLObject { // 0xc1b15d65
		o := MakeTLInputBotInlineMessageMediaGeo(nil)
		o.Data2.Constructor = -1045340827
		return o
	},
	1098628881: func() TLObject { // 0x417bbf11
		o := MakeTLInputBotInlineMessageMediaVenue(nil)
		o.Data2.Constructor = 1098628881
		return o
	},
	-1431327288: func() TLObject { // 0xaaafadc8
		o := MakeTLInputBotInlineMessageMediaVenue(nil)
		o.Data2.Constructor = -1431327288
		return o
	},
	-1494368259: func() TLObject { // 0xa6edbffd
		o := MakeTLInputBotInlineMessageMediaContact(nil)
		o.Data2.Constructor = -1494368259
		return o
	},
	766443943: func() TLObject { // 0x2daf01a7
		o := MakeTLInputBotInlineMessageMediaContact(nil)
		o.Data2.Constructor = 766443943
		return o
	},
	1262639204: func() TLObject { // 0x4b425864
		o := MakeTLInputBotInlineMessageGame(nil)
		o.Data2.Constructor = 1262639204
		return o
	},
	-1803769784: func() TLObject { // 0x947ca848
		o := MakeTLMessagesBotResults(nil)
		o.Data2.Constructor = -1803769784
		return o
	},
	-797791052: func() TLObject { // 0xd072acb4
		o := MakeTLRestrictionReason(nil)
		o.Data2.Constructor = -797791052
		return o
	},
	1928391342: func() TLObject { // 0x72f0eaae
		o := MakeTLInputDocumentEmpty(nil)
		o.Data2.Constructor = 1928391342
		return o
	},
	448771445: func() TLObject { // 0x1abfb575
		o := MakeTLInputDocument(nil)
		o.Data2.Constructor = 448771445
		return o
	},
	410618194: func() TLObject { // 0x18798952
		o := MakeTLInputDocument(nil)
		o.Data2.Constructor = 410618194
		return o
	},
	-1539849235: func() TLObject { // 0xa437c3ed
		o := MakeTLWallPaper(nil)
		o.Data2.Constructor = -1539849235
		return o
	},
	-263220756: func() TLObject { // 0xf04f91ec
		o := MakeTLWallPaper(nil)
		o.Data2.Constructor = -263220756
		return o
	},
	-860866985: func() TLObject { // 0xccb03657
		o := MakeTLWallPaper(nil)
		o.Data2.Constructor = -860866985
		return o
	},
	-1963717851: func() TLObject { // 0x8af40b25
		o := MakeTLWallPaperNoFile(nil)
		o.Data2.Constructor = -1963717851
		return o
	},
	1662091044: func() TLObject { // 0x63117f24
		o := MakeTLWallPaperSolid(nil)
		o.Data2.Constructor = 1662091044
		return o
	},
	1194918371: func() TLObject { // 0x473901e3
		o := MakeTLBlogContentPhotos(nil)
		o.Data2.Constructor = 1194918371
		return o
	},
	1439069091: func() TLObject { // 0x55c673a3
		o := MakeTLBlogContentDocument(nil)
		o.Data2.Constructor = 1439069091
		return o
	},
	32795759: func() TLObject { // 0x1f46c6f
		o := MakeTLBlogsState(nil)
		o.Data2.Constructor = 32795759
		return o
	},
	-75283823: func() TLObject { // 0xfb834291
		o := MakeTLTopPeerCategoryPeers(nil)
		o.Data2.Constructor = -75283823
		return o
	},
	372165663: func() TLObject { // 0x162ecc1f
		o := MakeTLFoundGif(nil)
		o.Data2.Constructor = 372165663
		return o
	},
	-1670052855: func() TLObject { // 0x9c750409
		o := MakeTLFoundGifCached(nil)
		o.Data2.Constructor = -1670052855
		return o
	},
	-2128698738: func() TLObject { // 0x811ea28e
		o := MakeTLAuthCheckedPhone(nil)
		o.Data2.Constructor = -2128698738
		return o
	},
	1200838592: func() TLObject { // 0x479357c0
		o := MakeTLSchemeMethod(nil)
		o.Data2.Constructor = 1200838592
		return o
	},
	1933519201: func() TLObject { // 0x733f2961
		o := MakeTLPeerSettings(nil)
		o.Data2.Constructor = 1933519201
		return o
	},
	-2122045747: func() TLObject { // 0x818426cd
		o := MakeTLPeerSettings(nil)
		o.Data2.Constructor = -2122045747
		return o
	},
	223655517: func() TLObject { // 0xd54b65d
		o := MakeTLMessagesFoundStickerSetsNotModified(nil)
		o.Data2.Constructor = 223655517
		return o
	},
	1359533640: func() TLObject { // 0x5108d648
		o := MakeTLMessagesFoundStickerSets(nil)
		o.Data2.Constructor = 1359533640
		return o
	},
	1968737087: func() TLObject { // 0x75588b3f
		o := MakeTLInputClientProxy(nil)
		o.Data2.Constructor = 1968737087
		return o
	},
	1949890536: func() TLObject { // 0x7438f7e8
		o := MakeTLDialogFilter(nil)
		o.Data2.Constructor = 1949890536
		return o
	},
	-1269012015: func() TLObject { // 0xb45c69d1
		o := MakeTLMessagesAffectedHistory(nil)
		o.Data2.Constructor = -1269012015
		return o
	},
	-886477832: func() TLObject { // 0xcb296bf8
		o := MakeTLLabeledPrice(nil)
		o.Data2.Constructor = -886477832
		return o
	},
	-618540889: func() TLObject { // 0xdb21d0a7
		o := MakeTLInputSecureValue(nil)
		o.Data2.Constructor = -618540889
		return o
	},
	108557032: func() TLObject { // 0x67872e8
		o := MakeTLInputSecureValue(nil)
		o.Data2.Constructor = 108557032
		return o
	},
	-1059442448: func() TLObject { // 0xc0da30f0
		o := MakeTLInputSecureValue(nil)
		o.Data2.Constructor = -1059442448
		return o
	},
	-1392388579: func() TLObject { // 0xad01d61d
		o := MakeTLAuthorization(nil)
		o.Data2.Constructor = -1392388579
		return o
	},
	2079516406: func() TLObject { // 0x7bf2e6f6
		o := MakeTLAuthorization(nil)
		o.Data2.Constructor = 2079516406
		return o
	},
	1326562017: func() TLObject { // 0x4f11bae1
		o := MakeTLUserProfilePhotoEmpty(nil)
		o.Data2.Constructor = 1326562017
		return o
	},
	1775479590: func() TLObject { // 0x69d3ab26
		o := MakeTLUserProfilePhoto(nil)
		o.Data2.Constructor = 1775479590
		return o
	},
	-321430132: func() TLObject { // 0xecd75d8c
		o := MakeTLUserProfilePhoto(nil)
		o.Data2.Constructor = -321430132
		return o
	},
	-715532088: func() TLObject { // 0xd559d8c8
		o := MakeTLUserProfilePhoto(nil)
		o.Data2.Constructor = -715532088
		return o
	},
	-1683826688: func() TLObject { // 0x9ba2d800
		o := MakeTLChatEmpty(nil)
		o.Data2.Constructor = -1683826688
		return o
	},
	1004149726: func() TLObject { // 0x3bda1bde
		o := MakeTLChat(nil)
		o.Data2.Constructor = 1004149726
		return o
	},
	-652419756: func() TLObject { // 0xd91cdd54
		o := MakeTLChat(nil)
		o.Data2.Constructor = -652419756
		return o
	},
	120753115: func() TLObject { // 0x7328bdb
		o := MakeTLChatForbidden(nil)
		o.Data2.Constructor = 120753115
		return o
	},
	-753232354: func() TLObject { // 0xd31a961e
		o := MakeTLChannel(nil)
		o.Data2.Constructor = -753232354
		return o
	},
	1307772980: func() TLObject { // 0x4df30834
		o := MakeTLChannel(nil)
		o.Data2.Constructor = 1307772980
		return o
	},
	-930515796: func() TLObject { // 0xc88974ac
		o := MakeTLChannel(nil)
		o.Data2.Constructor = -930515796
		return o
	},
	1158377749: func() TLObject { // 0x450b7115
		o := MakeTLChannel(nil)
		o.Data2.Constructor = 1158377749
		return o
	},
	681420594: func() TLObject { // 0x289da732
		o := MakeTLChannelForbidden(nil)
		o.Data2.Constructor = 681420594
		return o
	},
	922273905: func() TLObject { // 0x36f8c871
		o := MakeTLDocumentEmpty(nil)
		o.Data2.Constructor = 922273905
		return o
	},
	512177195: func() TLObject { // 0x1e87342b
		o := MakeTLDocument(nil)
		o.Data2.Constructor = 512177195
		return o
	},
	-1683841855: func() TLObject { // 0x9ba29cc1
		o := MakeTLDocument(nil)
		o.Data2.Constructor = -1683841855
		return o
	},
	1498631756: func() TLObject { // 0x59534e4c
		o := MakeTLDocument(nil)
		o.Data2.Constructor = 1498631756
		return o
	},
	-2027738169: func() TLObject { // 0x87232bc7
		o := MakeTLDocument(nil)
		o.Data2.Constructor = -2027738169
		return o
	},
	-1137792208: func() TLObject { // 0xbc2eab30
		o := MakeTLPrivacyKeyStatusTimestamp(nil)
		o.Data2.Constructor = -1137792208
		return o
	},
	1343122938: func() TLObject { // 0x500e6dfa
		o := MakeTLPrivacyKeyChatInvite(nil)
		o.Data2.Constructor = 1343122938
		return o
	},
	1030105979: func() TLObject { // 0x3d662b7b
		o := MakeTLPrivacyKeyPhoneCall(nil)
		o.Data2.Constructor = 1030105979
		return o
	},
	961092808: func() TLObject { // 0x39491cc8
		o := MakeTLPrivacyKeyPhoneP2P(nil)
		o.Data2.Constructor = 961092808
		return o
	},
	1777096355: func() TLObject { // 0x69ec56a3
		o := MakeTLPrivacyKeyForwards(nil)
		o.Data2.Constructor = 1777096355
		return o
	},
	-1777000467: func() TLObject { // 0x96151fed
		o := MakeTLPrivacyKeyProfilePhoto(nil)
		o.Data2.Constructor = -1777000467
		return o
	},
	-778378131: func() TLObject { // 0xd19ae46d
		o := MakeTLPrivacyKeyPhoneNumber(nil)
		o.Data2.Constructor = -778378131
		return o
	},
	1124062251: func() TLObject { // 0x42ffd42b
		o := MakeTLPrivacyKeyAddedByPhone(nil)
		o.Data2.Constructor = 1124062251
		return o
	},
	-324801088: func() TLObject { // 0xeca3edc0
		o := MakeTLPrivacyKeyAddedByUsername(nil)
		o.Data2.Constructor = -324801088
		return o
	},
	358616271: func() TLObject { // 0x15600ccf
		o := MakeTLPrivacyKeySendMessage(nil)
		o.Data2.Constructor = 358616271
		return o
	},
	-1551583367: func() TLObject { // 0xa384b779
		o := MakeTLReceivedNotifyMessage(nil)
		o.Data2.Constructor = -1551583367
		return o
	},
	4883767: func() TLObject { // 0x4a8537
		o := MakeTLSecurePasswordKdfAlgoUnknown(nil)
		o.Data2.Constructor = 4883767
		return o
	},
	-1141711456: func() TLObject { // 0xbbf2dda0
		o := MakeTLSecurePasswordKdfAlgoPBKDF2(nil)
		o.Data2.Constructor = -1141711456
		return o
	},
	-2042159726: func() TLObject { // 0x86471d92
		o := MakeTLSecurePasswordKdfAlgoSHA512(nil)
		o.Data2.Constructor = -2042159726
		return o
	},
	136574537: func() TLObject { // 0x823f649
		o := MakeTLMessagesVotesList(nil)
		o.Data2.Constructor = 136574537
		return o
	},
	-208488460: func() TLObject { // 0xf392b7f4
		o := MakeTLInputPhoneContact(nil)
		o.Data2.Constructor = -208488460
		return o
	},
	-242530370: func() TLObject { // 0xf18b47be
		o := MakeTLChatParticipant(nil)
		o.Data2.Constructor = -242530370
		return o
	},
	-925415106: func() TLObject { // 0xc8d7493e
		o := MakeTLChatParticipant(nil)
		o.Data2.Constructor = -925415106
		return o
	},
	-1542470162: func() TLObject { // 0xa40fc5ee
		o := MakeTLChatParticipantCreator(nil)
		o.Data2.Constructor = -1542470162
		return o
	},
	-636267638: func() TLObject { // 0xda13538a
		o := MakeTLChatParticipantCreator(nil)
		o.Data2.Constructor = -636267638
		return o
	},
	-1392859011: func() TLObject { // 0xacfaa87d
		o := MakeTLChatParticipantAdmin(nil)
		o.Data2.Constructor = -1392859011
		return o
	},
	-489233354: func() TLObject { // 0xe2d6e436
		o := MakeTLChatParticipantAdmin(nil)
		o.Data2.Constructor = -489233354
		return o
	},
	649453030: func() TLObject { // 0x26b5dde6
		o := MakeTLMessagesMessageEditData(nil)
		o.Data2.Constructor = 649453030
		return o
	},
	995769920: func() TLObject { // 0x3b5a3e40
		o := MakeTLChannelAdminLogEvent(nil)
		o.Data2.Constructor = 995769920
		return o
	},
	1474462241: func() TLObject { // 0x57e28221
		o := MakeTLAccountContentSettings(nil)
		o.Data2.Constructor = 1474462241
		return o
	},
	2004110666: func() TLObject { // 0x77744d4a
		o := MakeTLDialogFilterSuggested(nil)
		o.Data2.Constructor = 2004110666
		return o
	},
	1840491641: func() TLObject { // 0x6db3ac79
		o := MakeTLBizDataRaw(nil)
		o.Data2.Constructor = 1840491641
		return o
	},
	480546647: func() TLObject { // 0x1ca48f57
		o := MakeTLInputChatPhotoEmpty(nil)
		o.Data2.Constructor = 480546647
		return o
	},
	-968723890: func() TLObject { // 0xc642724e
		o := MakeTLInputChatUploadedPhoto(nil)
		o.Data2.Constructor = -968723890
		return o
	},
	-1837345356: func() TLObject { // 0x927c55b4
		o := MakeTLInputChatUploadedPhoto(nil)
		o.Data2.Constructor = -1837345356
		return o
	},
	-1991004873: func() TLObject { // 0x8953ad37
		o := MakeTLInputChatPhoto(nil)
		o.Data2.Constructor = -1991004873
		return o
	},

	// Method
	1615239032: func() TLObject { // 0x60469778
		return &TLReqPq{
			Constructor: 1615239032,
		}
	},
	-1099002127: func() TLObject { // 0xbe7e8ef1
		return &TLReqPqMulti{
			Constructor: -1099002127,
		}
	},
	-686627650: func() TLObject { // 0xd712e4be
		return &TLReq_DHParams{
			Constructor: -686627650,
		}
	},
	-184262881: func() TLObject { // 0xf5045f1f
		return &TLSetClient_DHParams{
			Constructor: -184262881,
		}
	},
	-784117408: func() TLObject { // 0xd1435160
		return &TLDestroyAuthKey{
			Constructor: -784117408,
		}
	},
	1491380032: func() TLObject { // 0x58e4a740
		return &TLRpcDropAnswer{
			Constructor: 1491380032,
		}
	},
	-1188971260: func() TLObject { // 0xb921bd04
		return &TLGetFutureSalts{
			Constructor: -1188971260,
		}
	},
	2059302892: func() TLObject { // 0x7abe77ec
		return &TLPing{
			Constructor: 2059302892,
		}
	},
	-213746804: func() TLObject { // 0xf3427b8c
		return &TLPingDelayDisconnect{
			Constructor: -213746804,
		}
	},
	-414113498: func() TLObject { // 0xe7512126
		return &TLDestroySession{
			Constructor: -414113498,
		}
	},
	-1705021803: func() TLObject { // 0x9a5f6e95
		return &TLContestSaveDeveloperInfo{
			Constructor: -1705021803,
		}
	},
	-878758099: func() TLObject { // 0xcb9f372d
		return &TLInvokeAfterMsg{
			Constructor: -878758099,
		}
	},
	1036301552: func() TLObject { // 0x3dc4b4f0
		return &TLInvokeAfterMsgs{
			Constructor: 1036301552,
		}
	},
	-1043505495: func() TLObject { // 0xc1cd5ea9
		return &TLInitConnection{
			Constructor: -1043505495,
		}
	},
	2018609336: func() TLObject { // 0x785188b8
		return &TLInitConnection{
			Constructor: 2018609336,
		}
	},
	-951575130: func() TLObject { // 0xc7481da6
		return &TLInitConnection{
			Constructor: -951575130,
		}
	},
	-627372787: func() TLObject { // 0xda9b0d0d
		return &TLInvokeWithLayer{
			Constructor: -627372787,
		}
	},
	-1080796745: func() TLObject { // 0xbf9459b7
		return &TLInvokeWithoutUpdates{
			Constructor: -1080796745,
		}
	},
	911373810: func() TLObject { // 0x365275f2
		return &TLInvokeWithMessagesRange{
			Constructor: 911373810,
		}
	},
	-1398145746: func() TLObject { // 0xaca9fd2e
		return &TLInvokeWithTakeout{
			Constructor: -1398145746,
		}
	},
	-1502141361: func() TLObject { // 0xa677244f
		return &TLAuthSendCode{
			Constructor: -1502141361,
		}
	},
	-2035355412: func() TLObject { // 0x86aef0ec
		return &TLAuthSendCode{
			Constructor: -2035355412,
		}
	},
	-855805745: func() TLObject { // 0xccfd70cf
		return &TLAuthSendCode{
			Constructor: -855805745,
		}
	},
	-2131827673: func() TLObject { // 0x80eee427
		return &TLAuthSignUp{
			Constructor: -2131827673,
		}
	},
	453408308: func() TLObject { // 0x1b067634
		return &TLAuthSignUp{
			Constructor: 453408308,
		}
	},
	-1126886015: func() TLObject { // 0xbcd51581
		return &TLAuthSignIn{
			Constructor: -1126886015,
		}
	},
	1461180992: func() TLObject { // 0x5717da40
		return &TLAuthLogOut{
			Constructor: 1461180992,
		}
	},
	-1616179942: func() TLObject { // 0x9fab0d1a
		return &TLAuthResetAuthorizations{
			Constructor: -1616179942,
		}
	},
	-440401971: func() TLObject { // 0xe5bfffcd
		return &TLAuthExportAuthorization{
			Constructor: -440401971,
		}
	},
	-470837741: func() TLObject { // 0xe3ef9613
		return &TLAuthImportAuthorization{
			Constructor: -470837741,
		}
	},
	-841733627: func() TLObject { // 0xcdd42a05
		return &TLAuthBindTempAuthKey{
			Constructor: -841733627,
		}
	},
	1738800940: func() TLObject { // 0x67a3ff2c
		return &TLAuthImportBotAuthorization{
			Constructor: 1738800940,
		}
	},
	-779399914: func() TLObject { // 0xd18b4d16
		return &TLAuthCheckPassword{
			Constructor: -779399914,
		}
	},
	174260510: func() TLObject { // 0xa63011e
		return &TLAuthCheckPassword{
			Constructor: 174260510,
		}
	},
	-661144474: func() TLObject { // 0xd897bc66
		return &TLAuthRequestPasswordRecovery{
			Constructor: -661144474,
		}
	},
	1319464594: func() TLObject { // 0x4ea56e92
		return &TLAuthRecoverPassword{
			Constructor: 1319464594,
		}
	},
	1056025023: func() TLObject { // 0x3ef1a9bf
		return &TLAuthResendCode{
			Constructor: 1056025023,
		}
	},
	520357240: func() TLObject { // 0x1f040578
		return &TLAuthCancelCode{
			Constructor: 520357240,
		}
	},
	-1907842680: func() TLObject { // 0x8e48a188
		return &TLAuthDropTempAuthKeys{
			Constructor: -1907842680,
		}
	},
	-1313598185: func() TLObject { // 0xb1b41517
		return &TLAuthExportLoginToken{
			Constructor: -1313598185,
		}
	},
	-1783866140: func() TLObject { // 0x95ac5ce4
		return &TLAuthImportLoginToken{
			Constructor: -1783866140,
		}
	},
	-392909491: func() TLObject { // 0xe894ad4d
		return &TLAuthAcceptLoginToken{
			Constructor: -392909491,
		}
	},
	1754754159: func() TLObject { // 0x68976c6f
		return &TLAccountRegisterDevice{
			Constructor: 1754754159,
		}
	},
	1555998096: func() TLObject { // 0x5cbea590
		return &TLAccountRegisterDevice{
			Constructor: 1555998096,
		}
	},
	1280460: func() TLObject { // 0x1389cc
		return &TLAccountRegisterDevice{
			Constructor: 1280460,
		}
	},
	1669245048: func() TLObject { // 0x637ea878
		return &TLAccountRegisterDevice{
			Constructor: 1669245048,
		}
	},
	813089983: func() TLObject { // 0x3076c4bf
		return &TLAccountUnregisterDevice{
			Constructor: 813089983,
		}
	},
	1707432768: func() TLObject { // 0x65c55b40
		return &TLAccountUnregisterDevice{
			Constructor: 1707432768,
		}
	},
	-2067899501: func() TLObject { // 0x84be5b93
		return &TLAccountUpdateNotifySettings{
			Constructor: -2067899501,
		}
	},
	313765169: func() TLObject { // 0x12b3ad31
		return &TLAccountGetNotifySettings{
			Constructor: 313765169,
		}
	},
	-612493497: func() TLObject { // 0xdb7e1747
		return &TLAccountResetNotifySettings{
			Constructor: -612493497,
		}
	},
	2018596725: func() TLObject { // 0x78515775
		return &TLAccountUpdateProfile{
			Constructor: 2018596725,
		}
	},
	1713919532: func() TLObject { // 0x6628562c
		return &TLAccountUpdateStatus{
			Constructor: 1713919532,
		}
	},
	-1430579357: func() TLObject { // 0xaabb1763
		return &TLAccountGetWallPapersAABB1763{
			Constructor: -1430579357,
		}
	},
	-1374118561: func() TLObject { // 0xae189d5f
		return &TLAccountReportPeer{
			Constructor: -1374118561,
		}
	},
	655677548: func() TLObject { // 0x2714d86c
		return &TLAccountCheckUsername{
			Constructor: 655677548,
		}
	},
	1040964988: func() TLObject { // 0x3e0bdd7c
		return &TLAccountUpdateUsername{
			Constructor: 1040964988,
		}
	},
	-623130288: func() TLObject { // 0xdadbc950
		return &TLAccountGetPrivacy{
			Constructor: -623130288,
		}
	},
	-906486552: func() TLObject { // 0xc9f81ce8
		return &TLAccountSetPrivacy{
			Constructor: -906486552,
		}
	},
	1099779595: func() TLObject { // 0x418d4e0b
		return &TLAccountDeleteAccount{
			Constructor: 1099779595,
		}
	},
	150761757: func() TLObject { // 0x8fc711d
		return &TLAccountGetAccountTTL{
			Constructor: 150761757,
		}
	},
	608323678: func() TLObject { // 0x2442485e
		return &TLAccountSetAccountTTL{
			Constructor: 608323678,
		}
	},
	-2108208411: func() TLObject { // 0x82574ae5
		return &TLAccountSendChangePhoneCode{
			Constructor: -2108208411,
		}
	},
	149257707: func() TLObject { // 0x8e57deb
		return &TLAccountSendChangePhoneCode{
			Constructor: 149257707,
		}
	},
	1891839707: func() TLObject { // 0x70c32edb
		return &TLAccountChangePhone{
			Constructor: 1891839707,
		}
	},
	954152242: func() TLObject { // 0x38df3532
		return &TLAccountUpdateDeviceLocked{
			Constructor: 954152242,
		}
	},
	-484392616: func() TLObject { // 0xe320c158
		return &TLAccountGetAuthorizations{
			Constructor: -484392616,
		}
	},
	-545786948: func() TLObject { // 0xdf77f3bc
		return &TLAccountResetAuthorization{
			Constructor: -545786948,
		}
	},
	1418342645: func() TLObject { // 0x548a30f5
		return &TLAccountGetPassword{
			Constructor: 1418342645,
		}
	},
	-1663767815: func() TLObject { // 0x9cd4eaf9
		return &TLAccountGetPasswordSettings{
			Constructor: -1663767815,
		}
	},
	-1131605573: func() TLObject { // 0xbc8d11bb
		return &TLAccountGetPasswordSettings{
			Constructor: -1131605573,
		}
	},
	-1516564433: func() TLObject { // 0xa59b102f
		return &TLAccountUpdatePasswordSettings{
			Constructor: -1516564433,
		}
	},
	-92517498: func() TLObject { // 0xfa7c4b86
		return &TLAccountUpdatePasswordSettings{
			Constructor: -92517498,
		}
	},
	457157256: func() TLObject { // 0x1b3faa88
		return &TLAccountSendConfirmPhoneCode{
			Constructor: 457157256,
		}
	},
	353818557: func() TLObject { // 0x1516d7bd
		return &TLAccountSendConfirmPhoneCode{
			Constructor: 353818557,
		}
	},
	1596029123: func() TLObject { // 0x5f2178c3
		return &TLAccountConfirmPhone{
			Constructor: 1596029123,
		}
	},
	1151208273: func() TLObject { // 0x449e0b51
		return &TLAccountGetTmpPassword{
			Constructor: 1151208273,
		}
	},
	1250046590: func() TLObject { // 0x4a82327e
		return &TLAccountGetTmpPassword{
			Constructor: 1250046590,
		}
	},
	405695855: func() TLObject { // 0x182e6d6f
		return &TLAccountGetWebAuthorizations{
			Constructor: 405695855,
		}
	},
	755087855: func() TLObject { // 0x2d01b9ef
		return &TLAccountResetWebAuthorization{
			Constructor: 755087855,
		}
	},
	1747789204: func() TLObject { // 0x682d2594
		return &TLAccountResetWebAuthorizations{
			Constructor: 1747789204,
		}
	},
	-1299661699: func() TLObject { // 0xb288bc7d
		return &TLAccountGetAllSecureValues{
			Constructor: -1299661699,
		}
	},
	1936088002: func() TLObject { // 0x73665bc2
		return &TLAccountGetSecureValue{
			Constructor: 1936088002,
		}
	},
	-1986010339: func() TLObject { // 0x899fe31d
		return &TLAccountSaveSecureValue{
			Constructor: -1986010339,
		}
	},
	-1199522741: func() TLObject { // 0xb880bc4b
		return &TLAccountDeleteSecureValue{
			Constructor: -1199522741,
		}
	},
	-1200903967: func() TLObject { // 0xb86ba8e1
		return &TLAccountGetAuthorizationForm{
			Constructor: -1200903967,
		}
	},
	-419267436: func() TLObject { // 0xe7027c94
		return &TLAccountAcceptAuthorization{
			Constructor: -419267436,
		}
	},
	-1516022023: func() TLObject { // 0xa5a356f9
		return &TLAccountSendVerifyPhoneCode{
			Constructor: -1516022023,
		}
	},
	-2110553932: func() TLObject { // 0x823380b4
		return &TLAccountSendVerifyPhoneCode{
			Constructor: -2110553932,
		}
	},
	1305716726: func() TLObject { // 0x4dd3a7f6
		return &TLAccountVerifyPhone{
			Constructor: 1305716726,
		}
	},
	1880182943: func() TLObject { // 0x7011509f
		return &TLAccountSendVerifyEmailCode{
			Constructor: 1880182943,
		}
	},
	-323339813: func() TLObject { // 0xecba39db
		return &TLAccountVerifyEmail{
			Constructor: -323339813,
		}
	},
	-262453244: func() TLObject { // 0xf05b4804
		return &TLAccountInitTakeoutSession{
			Constructor: -262453244,
		}
	},
	489050862: func() TLObject { // 0x1d2652ee
		return &TLAccountFinishTakeoutSession{
			Constructor: 489050862,
		}
	},
	-1881204448: func() TLObject { // 0x8fdf1920
		return &TLAccountConfirmPasswordEmail{
			Constructor: -1881204448,
		}
	},
	2055154197: func() TLObject { // 0x7a7f2a15
		return &TLAccountResendPasswordEmail{
			Constructor: 2055154197,
		}
	},
	-1043606090: func() TLObject { // 0xc1cbd5b6
		return &TLAccountCancelPasswordEmail{
			Constructor: -1043606090,
		}
	},
	-1626880216: func() TLObject { // 0x9f07c728
		return &TLAccountGetContactSignUpNotification{
			Constructor: -1626880216,
		}
	},
	-806076575: func() TLObject { // 0xcff43f61
		return &TLAccountSetContactSignUpNotification{
			Constructor: -806076575,
		}
	},
	1398240377: func() TLObject { // 0x53577479
		return &TLAccountGetNotifyExceptions{
			Constructor: 1398240377,
		}
	},
	-57811990: func() TLObject { // 0xfc8ddbea
		return &TLAccountGetWallPaper{
			Constructor: -57811990,
		}
	},
	-578472351: func() TLObject { // 0xdd853661
		return &TLAccountUploadWallPaper{
			Constructor: -578472351,
		}
	},
	-944071859: func() TLObject { // 0xc7ba9b4d
		return &TLAccountUploadWallPaper{
			Constructor: -944071859,
		}
	},
	1817860919: func() TLObject { // 0x6c5a5b37
		return &TLAccountSaveWallPaper{
			Constructor: 1817860919,
		}
	},
	412451251: func() TLObject { // 0x189581b3
		return &TLAccountSaveWallPaper{
			Constructor: 412451251,
		}
	},
	-18000023: func() TLObject { // 0xfeed5769
		return &TLAccountInstallWallPaper{
			Constructor: -18000023,
		}
	},
	1241741518: func() TLObject { // 0x4a0378ce
		return &TLAccountInstallWallPaper{
			Constructor: 1241741518,
		}
	},
	-1153722364: func() TLObject { // 0xbb3b9804
		return &TLAccountResetWallPapers{
			Constructor: -1153722364,
		}
	},
	1457130303: func() TLObject { // 0x56da0b3f
		return &TLAccountGetAutoDownloadSettings{
			Constructor: 1457130303,
		}
	},
	1995661875: func() TLObject { // 0x76f36233
		return &TLAccountSaveAutoDownloadSettings{
			Constructor: 1995661875,
		}
	},
	473805619: func() TLObject { // 0x1c3db333
		return &TLAccountUploadTheme{
			Constructor: 473805619,
		}
	},
	-2077048289: func() TLObject { // 0x8432c21f
		return &TLAccountCreateTheme{
			Constructor: -2077048289,
		}
	},
	729808255: func() TLObject { // 0x2b7ffd7f
		return &TLAccountCreateTheme{
			Constructor: 729808255,
		}
	},
	1555261397: func() TLObject { // 0x5cb367d5
		return &TLAccountUpdateTheme{
			Constructor: 1555261397,
		}
	},
	999203330: func() TLObject { // 0x3b8ea202
		return &TLAccountUpdateTheme{
			Constructor: 999203330,
		}
	},
	-229175188: func() TLObject { // 0xf257106c
		return &TLAccountSaveTheme{
			Constructor: -229175188,
		}
	},
	2061776695: func() TLObject { // 0x7ae43737
		return &TLAccountInstallTheme{
			Constructor: 2061776695,
		}
	},
	-1919060949: func() TLObject { // 0x8d9d742b
		return &TLAccountGetTheme{
			Constructor: -1919060949,
		}
	},
	676939512: func() TLObject { // 0x285946f8
		return &TLAccountGetThemes{
			Constructor: 676939512,
		}
	},
	-1250643605: func() TLObject { // 0xb574b16b
		return &TLAccountSetContentSettings{
			Constructor: -1250643605,
		}
	},
	-1952756306: func() TLObject { // 0x8b9b4dae
		return &TLAccountGetContentSettings{
			Constructor: -1952756306,
		}
	},
	1705865692: func() TLObject { // 0x65ad71dc
		return &TLAccountGetMultiWallPapers{
			Constructor: 1705865692,
		}
	},
	-349483786: func() TLObject { // 0xeb2b4cf6
		return &TLAccountGetGlobalPrivacySettings{
			Constructor: -349483786,
		}
	},
	517647042: func() TLObject { // 0x1edaaac2
		return &TLAccountSetGlobalPrivacySettings{
			Constructor: 517647042,
		}
	},
	227648840: func() TLObject { // 0xd91a548
		return &TLUsersGetUsers{
			Constructor: 227648840,
		}
	},
	-902781519: func() TLObject { // 0xca30a5b1
		return &TLUsersGetFullUser{
			Constructor: -902781519,
		}
	},
	-1865902923: func() TLObject { // 0x90c894b5
		return &TLUsersSetSecureValueErrors{
			Constructor: -1865902923,
		}
	},
	749357634: func() TLObject { // 0x2caa4a42
		return &TLContactsGetContactIDs{
			Constructor: 749357634,
		}
	},
	-995929106: func() TLObject { // 0xc4a353ee
		return &TLContactsGetStatuses{
			Constructor: -995929106,
		}
	},
	-1071414113: func() TLObject { // 0xc023849f
		return &TLContactsGetContacts{
			Constructor: -1071414113,
		}
	},
	746589157: func() TLObject { // 0x2c800be5
		return &TLContactsImportContacts{
			Constructor: 746589157,
		}
	},
	157945344: func() TLObject { // 0x96a0e00
		return &TLContactsDeleteContacts96A0E00{
			Constructor: 157945344,
		}
	},
	269745566: func() TLObject { // 0x1013fd9e
		return &TLContactsDeleteByPhones{
			Constructor: 269745566,
		}
	},
	1758204945: func() TLObject { // 0x68cc1411
		return &TLContactsBlock{
			Constructor: 1758204945,
		}
	},
	858475004: func() TLObject { // 0x332b49fc
		return &TLContactsBlock{
			Constructor: 858475004,
		}
	},
	-1096393392: func() TLObject { // 0xbea65d50
		return &TLContactsUnblock{
			Constructor: -1096393392,
		}
	},
	-448724803: func() TLObject { // 0xe54100bd
		return &TLContactsUnblock{
			Constructor: -448724803,
		}
	},
	-176409329: func() TLObject { // 0xf57c350f
		return &TLContactsGetBlocked{
			Constructor: -176409329,
		}
	},
	301470424: func() TLObject { // 0x11f812d8
		return &TLContactsSearch{
			Constructor: 301470424,
		}
	},
	-113456221: func() TLObject { // 0xf93ccba3
		return &TLContactsResolveUsername{
			Constructor: -113456221,
		}
	},
	-728224331: func() TLObject { // 0xd4982db5
		return &TLContactsGetTopPeers{
			Constructor: -728224331,
		}
	},
	451113900: func() TLObject { // 0x1ae373ac
		return &TLContactsResetTopPeerRating{
			Constructor: 451113900,
		}
	},
	-2020263951: func() TLObject { // 0x879537f1
		return &TLContactsResetSaved{
			Constructor: -2020263951,
		}
	},
	-2098076769: func() TLObject { // 0x82f1e39f
		return &TLContactsGetSaved{
			Constructor: -2098076769,
		}
	},
	-2062238246: func() TLObject { // 0x8514bdda
		return &TLContactsToggleTopPeers{
			Constructor: -2062238246,
		}
	},
	-386636848: func() TLObject { // 0xe8f463d0
		return &TLContactsAddContact{
			Constructor: -386636848,
		}
	},
	-130964977: func() TLObject { // 0xf831a20f
		return &TLContactsAcceptContact{
			Constructor: -130964977,
		}
	},
	-750207932: func() TLObject { // 0xd348bc44
		return &TLContactsGetLocated{
			Constructor: -750207932,
		}
	},
	171270230: func() TLObject { // 0xa356056
		return &TLContactsGetLocated{
			Constructor: 171270230,
		}
	},
	698914348: func() TLObject { // 0x29a8962c
		return &TLContactsBlockFromReplies{
			Constructor: 698914348,
		}
	},
	1673946374: func() TLObject { // 0x63c66506
		return &TLMessagesGetMessages{
			Constructor: 1673946374,
		}
	},
	1109588596: func() TLObject { // 0x4222fa74
		return &TLMessagesGetMessages{
			Constructor: 1109588596,
		}
	},
	-1594999949: func() TLObject { // 0xa0ee3b73
		return &TLMessagesGetDialogs{
			Constructor: -1594999949,
		}
	},
	-1332171034: func() TLObject { // 0xb098aee6
		return &TLMessagesGetDialogs{
			Constructor: -1332171034,
		}
	},
	421243333: func() TLObject { // 0x191ba9c5
		return &TLMessagesGetDialogs{
			Constructor: 421243333,
		}
	},
	96533218: func() TLObject { // 0x5c0fae2
		return &TLMessagesGetDialogs{
			Constructor: 96533218,
		}
	},
	-591691168: func() TLObject { // 0xdcbb8260
		return &TLMessagesGetHistory{
			Constructor: -591691168,
		}
	},
	-1347868602: func() TLObject { // 0xafa92846
		return &TLMessagesGetHistory{
			Constructor: -1347868602,
		}
	},
	204812012: func() TLObject { // 0xc352eec
		return &TLMessagesSearch{
			Constructor: 204812012,
		}
	},
	1310163211: func() TLObject { // 0x4e17810b
		return &TLMessagesSearch{
			Constructor: 1310163211,
		}
	},
	-2045448344: func() TLObject { // 0x8614ef68
		return &TLMessagesSearch{
			Constructor: -2045448344,
		}
	},
	60726944: func() TLObject { // 0x39e9ea0
		return &TLMessagesSearch{
			Constructor: 60726944,
		}
	},
	-225926539: func() TLObject { // 0xf288a275
		return &TLMessagesSearch{
			Constructor: -225926539,
		}
	},
	238054714: func() TLObject { // 0xe306d3a
		return &TLMessagesReadHistoryE306D3A{
			Constructor: 238054714,
		}
	},
	469850889: func() TLObject { // 0x1c015b09
		return &TLMessagesDeleteHistory{
			Constructor: 469850889,
		}
	},
	-443640366: func() TLObject { // 0xe58e95d2
		return &TLMessagesDeleteMessages{
			Constructor: -443640366,
		}
	},
	94983360: func() TLObject { // 0x5a954c0
		return &TLMessagesReceivedMessages{
			Constructor: 94983360,
		}
	},
	1486110434: func() TLObject { // 0x58943ee2
		return &TLMessagesSetTyping{
			Constructor: 1486110434,
		}
	},
	-1551737264: func() TLObject { // 0xa3825e50
		return &TLMessagesSetTyping{
			Constructor: -1551737264,
		}
	},
	-991292016: func() TLObject { // 0xc4ea1590
		return &TLMessagesSendMessage{
			Constructor: -991292016,
		}
	},
	1376532592: func() TLObject { // 0x520c3870
		return &TLMessagesSendMessage{
			Constructor: 1376532592,
		}
	},
	-91733382: func() TLObject { // 0xfa88427a
		return &TLMessagesSendMessage{
			Constructor: -91733382,
		}
	},
	-298183410: func() TLObject { // 0xee3a150e
		return &TLMessagesSendMedia{
			Constructor: -298183410,
		}
	},
	881978281: func() TLObject { // 0x3491eba9
		return &TLMessagesSendMedia{
			Constructor: 881978281,
		}
	},
	-1194252757: func() TLObject { // 0xb8d1262b
		return &TLMessagesSendMedia{
			Constructor: -1194252757,
		}
	},
	-923703407: func() TLObject { // 0xc8f16791
		return &TLMessagesSendMedia{
			Constructor: -923703407,
		}
	},
	-637606386: func() TLObject { // 0xd9fee60e
		return &TLMessagesForwardMessages{
			Constructor: -637606386,
		}
	},
	1888354709: func() TLObject { // 0x708e0195
		return &TLMessagesForwardMessages{
			Constructor: 1888354709,
		}
	},
	-820669733: func() TLObject { // 0xcf1592db
		return &TLMessagesReportSpam{
			Constructor: -820669733,
		}
	},
	913498268: func() TLObject { // 0x3672e09c
		return &TLMessagesGetPeerSettings{
			Constructor: 913498268,
		}
	},
	-1115507112: func() TLObject { // 0xbd82b658
		return &TLMessagesReport{
			Constructor: -1115507112,
		}
	},
	1013621127: func() TLObject { // 0x3c6aa187
		return &TLMessagesGetChats{
			Constructor: 1013621127,
		}
	},
	998448230: func() TLObject { // 0x3b831c66
		return &TLMessagesGetFullChat{
			Constructor: 998448230,
		}
	},
	-599447467: func() TLObject { // 0xdc452855
		return &TLMessagesEditChatTitle{
			Constructor: -599447467,
		}
	},
	-900957736: func() TLObject { // 0xca4c79d8
		return &TLMessagesEditChatPhoto{
			Constructor: -900957736,
		}
	},
	-106911223: func() TLObject { // 0xf9a0aa09
		return &TLMessagesAddChatUser{
			Constructor: -106911223,
		}
	},
	-530505962: func() TLObject { // 0xe0611f16
		return &TLMessagesDeleteChatUser{
			Constructor: -530505962,
		}
	},
	164303470: func() TLObject { // 0x9cb126e
		return &TLMessagesCreateChat{
			Constructor: 164303470,
		}
	},
	651135312: func() TLObject { // 0x26cf8950
		return &TLMessagesGetDhConfig{
			Constructor: 651135312,
		}
	},
	-162681021: func() TLObject { // 0xf64daf43
		return &TLMessagesRequestEncryption{
			Constructor: -162681021,
		}
	},
	1035731989: func() TLObject { // 0x3dbc0415
		return &TLMessagesAcceptEncryption{
			Constructor: 1035731989,
		}
	},
	-304536635: func() TLObject { // 0xedd923c5
		return &TLMessagesDiscardEncryption{
			Constructor: -304536635,
		}
	},
	2031374829: func() TLObject { // 0x791451ed
		return &TLMessagesSetEncryptedTyping{
			Constructor: 2031374829,
		}
	},
	2135648522: func() TLObject { // 0x7f4b690a
		return &TLMessagesReadEncryptedHistory{
			Constructor: 2135648522,
		}
	},
	1157265941: func() TLObject { // 0x44fa7a15
		return &TLMessagesSendEncrypted{
			Constructor: 1157265941,
		}
	},
	-1451792525: func() TLObject { // 0xa9776773
		return &TLMessagesSendEncrypted{
			Constructor: -1451792525,
		}
	},
	1431914525: func() TLObject { // 0x5559481d
		return &TLMessagesSendEncryptedFile{
			Constructor: 1431914525,
		}
	},
	-1701831834: func() TLObject { // 0x9a901b66
		return &TLMessagesSendEncryptedFile{
			Constructor: -1701831834,
		}
	},
	852769188: func() TLObject { // 0x32d439a4
		return &TLMessagesSendEncryptedService{
			Constructor: 852769188,
		}
	},
	1436924774: func() TLObject { // 0x55a5bb66
		return &TLMessagesReceivedQueue{
			Constructor: 1436924774,
		}
	},
	1259113487: func() TLObject { // 0x4b0c8c0f
		return &TLMessagesReportEncryptedSpam{
			Constructor: 1259113487,
		}
	},
	916930423: func() TLObject { // 0x36a73f77
		return &TLMessagesReadMessageContents{
			Constructor: 916930423,
		}
	},
	71126828: func() TLObject { // 0x43d4f2c
		return &TLMessagesGetStickers{
			Constructor: 71126828,
		}
	},
	-2050272894: func() TLObject { // 0x85cb5182
		return &TLMessagesGetStickers{
			Constructor: -2050272894,
		}
	},
	479598769: func() TLObject { // 0x1c9618b1
		return &TLMessagesGetAllStickers{
			Constructor: 479598769,
		}
	},
	-1956073268: func() TLObject { // 0x8b68b0cc
		return &TLMessagesGetWebPagePreview{
			Constructor: -1956073268,
		}
	},
	623001124: func() TLObject { // 0x25223e24
		return &TLMessagesGetWebPagePreview{
			Constructor: 623001124,
		}
	},
	234312524: func() TLObject { // 0xdf7534c
		return &TLMessagesExportChatInvite{
			Constructor: 234312524,
		}
	},
	2106086025: func() TLObject { // 0x7d885289
		return &TLMessagesExportChatInvite{
			Constructor: 2106086025,
		}
	},
	1051570619: func() TLObject { // 0x3eadb1bb
		return &TLMessagesCheckChatInvite{
			Constructor: 1051570619,
		}
	},
	1817183516: func() TLObject { // 0x6c50051c
		return &TLMessagesImportChatInvite{
			Constructor: 1817183516,
		}
	},
	639215886: func() TLObject { // 0x2619a90e
		return &TLMessagesGetStickerSet{
			Constructor: 639215886,
		}
	},
	-946871200: func() TLObject { // 0xc78fe460
		return &TLMessagesInstallStickerSet{
			Constructor: -946871200,
		}
	},
	-110209570: func() TLObject { // 0xf96e55de
		return &TLMessagesUninstallStickerSet{
			Constructor: -110209570,
		}
	},
	-421563528: func() TLObject { // 0xe6df7378
		return &TLMessagesStartBot{
			Constructor: -421563528,
		}
	},
	1468322785: func() TLObject { // 0x5784d3e1
		return &TLMessagesGetMessagesViews5784D3E1{
			Constructor: 1468322785,
		}
	},
	-1444503762: func() TLObject { // 0xa9e69f2e
		return &TLMessagesEditChatAdmin{
			Constructor: -1444503762,
		}
	},
	363051235: func() TLObject { // 0x15a3b8e3
		return &TLMessagesMigrateChat{
			Constructor: 363051235,
		}
	},
	1271290010: func() TLObject { // 0x4bc6589a
		return &TLMessagesSearchGlobal{
			Constructor: 1271290010,
		}
	},
	-1083038300: func() TLObject { // 0xbf7225a4
		return &TLMessagesSearchGlobal{
			Constructor: -1083038300,
		}
	},
	259638801: func() TLObject { // 0xf79c611
		return &TLMessagesSearchGlobal{
			Constructor: 259638801,
		}
	},
	-1640190800: func() TLObject { // 0x9e3cacb0
		return &TLMessagesSearchGlobal{
			Constructor: -1640190800,
		}
	},
	2016638777: func() TLObject { // 0x78337739
		return &TLMessagesReorderStickerSets{
			Constructor: 2016638777,
		}
	},
	864953444: func() TLObject { // 0x338e2464
		return &TLMessagesGetDocumentByHash{
			Constructor: 864953444,
		}
	},
	-2084618926: func() TLObject { // 0x83bf3d52
		return &TLMessagesGetSavedGifs{
			Constructor: -2084618926,
		}
	},
	846868683: func() TLObject { // 0x327a30cb
		return &TLMessagesSaveGif{
			Constructor: 846868683,
		}
	},
	1364105629: func() TLObject { // 0x514e999d
		return &TLMessagesGetInlineBotResults{
			Constructor: 1364105629,
		}
	},
	-346119674: func() TLObject { // 0xeb5ea206
		return &TLMessagesSetInlineBotResults{
			Constructor: -346119674,
		}
	},
	570955184: func() TLObject { // 0x220815b0
		return &TLMessagesSendInlineBotResult{
			Constructor: 570955184,
		}
	},
	-1318189314: func() TLObject { // 0xb16e06fe
		return &TLMessagesSendInlineBotResult{
			Constructor: -1318189314,
		}
	},
	-39416522: func() TLObject { // 0xfda68d36
		return &TLMessagesGetMessageEditData{
			Constructor: -39416522,
		}
	},
	1224152952: func() TLObject { // 0x48f71778
		return &TLMessagesEditMessage{
			Constructor: 1224152952,
		}
	},
	-787025122: func() TLObject { // 0xd116f31e
		return &TLMessagesEditMessage{
			Constructor: -787025122,
		}
	},
	-1073683256: func() TLObject { // 0xc000e4c8
		return &TLMessagesEditMessage{
			Constructor: -1073683256,
		}
	},
	97630429: func() TLObject { // 0x5d1b8dd
		return &TLMessagesEditMessage{
			Constructor: 97630429,
		}
	},
	-2091549254: func() TLObject { // 0x83557dba
		return &TLMessagesEditInlineBotMessage{
			Constructor: -2091549254,
		}
	},
	-1379669976: func() TLObject { // 0xadc3e828
		return &TLMessagesEditInlineBotMessage{
			Constructor: -1379669976,
		}
	},
	-1327463869: func() TLObject { // 0xb0e08243
		return &TLMessagesEditInlineBotMessage{
			Constructor: -1327463869,
		}
	},
	-1824339449: func() TLObject { // 0x9342ca07
		return &TLMessagesGetBotCallbackAnswer{
			Constructor: -1824339449,
		}
	},
	-2130010132: func() TLObject { // 0x810a9fec
		return &TLMessagesGetBotCallbackAnswer{
			Constructor: -2130010132,
		}
	},
	-712043766: func() TLObject { // 0xd58f130a
		return &TLMessagesSetBotCallbackAnswer{
			Constructor: -712043766,
		}
	},
	-462373635: func() TLObject { // 0xe470bcfd
		return &TLMessagesGetPeerDialogs{
			Constructor: -462373635,
		}
	},
	764901049: func() TLObject { // 0x2d9776b9
		return &TLMessagesGetPeerDialogs{
			Constructor: 764901049,
		}
	},
	-1137057461: func() TLObject { // 0xbc39e14b
		return &TLMessagesSaveDraft{
			Constructor: -1137057461,
		}
	},
	1782549861: func() TLObject { // 0x6a3f8d65
		return &TLMessagesGetAllDrafts{
			Constructor: 1782549861,
		}
	},
	766298703: func() TLObject { // 0x2dacca4f
		return &TLMessagesGetFeaturedStickers{
			Constructor: 766298703,
		}
	},
	1527873830: func() TLObject { // 0x5b118126
		return &TLMessagesReadFeaturedStickers{
			Constructor: 1527873830,
		}
	},
	1587647177: func() TLObject { // 0x5ea192c9
		return &TLMessagesGetRecentStickers{
			Constructor: 1587647177,
		}
	},
	958863608: func() TLObject { // 0x392718f8
		return &TLMessagesSaveRecentSticker{
			Constructor: 958863608,
		}
	},
	881736127: func() TLObject { // 0x348e39bf
		return &TLMessagesSaveRecentSticker{
			Constructor: 881736127,
		}
	},
	-1986437075: func() TLObject { // 0x8999602d
		return &TLMessagesClearRecentStickers{
			Constructor: -1986437075,
		}
	},
	1475442322: func() TLObject { // 0x57f17692
		return &TLMessagesGetArchivedStickers{
			Constructor: 1475442322,
		}
	},
	1706608543: func() TLObject { // 0x65b8c79f
		return &TLMessagesGetMaskStickers{
			Constructor: 1706608543,
		}
	},
	-866424884: func() TLObject { // 0xcc5b67cc
		return &TLMessagesGetAttachedStickers{
			Constructor: -866424884,
		}
	},
	-1896289088: func() TLObject { // 0x8ef8ecc0
		return &TLMessagesSetGameScore{
			Constructor: -1896289088,
		}
	},
	363700068: func() TLObject { // 0x15ad9f64
		return &TLMessagesSetInlineGameScore{
			Constructor: 363700068,
		}
	},
	-400399203: func() TLObject { // 0xe822649d
		return &TLMessagesGetGameHighScores{
			Constructor: -400399203,
		}
	},
	258170395: func() TLObject { // 0xf635e1b
		return &TLMessagesGetInlineGameHighScores{
			Constructor: 258170395,
		}
	},
	218777796: func() TLObject { // 0xd0a48c4
		return &TLMessagesGetCommonChats{
			Constructor: 218777796,
		}
	},
	-341307408: func() TLObject { // 0xeba80ff0
		return &TLMessagesGetAllChats{
			Constructor: -341307408,
		}
	},
	852135825: func() TLObject { // 0x32ca8f91
		return &TLMessagesGetWebPage{
			Constructor: 852135825,
		}
	},
	-1489903017: func() TLObject { // 0xa731e257
		return &TLMessagesToggleDialogPin{
			Constructor: -1489903017,
		}
	},
	847887978: func() TLObject { // 0x3289be6a
		return &TLMessagesToggleDialogPin{
			Constructor: 847887978,
		}
	},
	991616823: func() TLObject { // 0x3b1adf37
		return &TLMessagesReorderPinnedDialogs{
			Constructor: 991616823,
		}
	},
	1532089919: func() TLObject { // 0x5b51d63f
		return &TLMessagesReorderPinnedDialogs{
			Constructor: 1532089919,
		}
	},
	-1784678844: func() TLObject { // 0x959ff644
		return &TLMessagesReorderPinnedDialogs{
			Constructor: -1784678844,
		}
	},
	-692498958: func() TLObject { // 0xd6b94df2
		return &TLMessagesGetPinnedDialogs{
			Constructor: -692498958,
		}
	},
	-497756594: func() TLObject { // 0xe254d64e
		return &TLMessagesGetPinnedDialogs{
			Constructor: -497756594,
		}
	},
	-436833542: func() TLObject { // 0xe5f672fa
		return &TLMessagesSetBotShippingResults{
			Constructor: -436833542,
		}
	},
	163765653: func() TLObject { // 0x9c2dd95
		return &TLMessagesSetBotPrecheckoutResults{
			Constructor: 163765653,
		}
	},
	1369162417: func() TLObject { // 0x519bc2b1
		return &TLMessagesUploadMedia{
			Constructor: 1369162417,
		}
	},
	-914493408: func() TLObject { // 0xc97df020
		return &TLMessagesSendScreenshotNotification{
			Constructor: -914493408,
		}
	},
	567151374: func() TLObject { // 0x21ce0b0e
		return &TLMessagesGetFavedStickers{
			Constructor: 567151374,
		}
	},
	-1174420133: func() TLObject { // 0xb9ffc55b
		return &TLMessagesFaveSticker{
			Constructor: -1174420133,
		}
	},
	1180140658: func() TLObject { // 0x46578472
		return &TLMessagesGetUnreadMentions{
			Constructor: 1180140658,
		}
	},
	251759059: func() TLObject { // 0xf0189d3
		return &TLMessagesReadMentions{
			Constructor: 251759059,
		}
	},
	-1144759543: func() TLObject { // 0xbbc45b09
		return &TLMessagesGetRecentLocations{
			Constructor: -1144759543,
		}
	},
	613691874: func() TLObject { // 0x249431e2
		return &TLMessagesGetRecentLocations{
			Constructor: 613691874,
		}
	},
	49993182: func() TLObject { // 0x2fad5de
		return &TLMessagesSendMultiMedia{
			Constructor: 49993182,
		}
	},
	-872345397: func() TLObject { // 0xcc0110cb
		return &TLMessagesSendMultiMedia{
			Constructor: -872345397,
		}
	},
	546656559: func() TLObject { // 0x2095512f
		return &TLMessagesSendMultiMedia{
			Constructor: 546656559,
		}
	},
	1347929239: func() TLObject { // 0x5057c497
		return &TLMessagesUploadEncryptedFile{
			Constructor: 1347929239,
		}
	},
	-1028140917: func() TLObject { // 0xc2b7d08b
		return &TLMessagesSearchStickerSets{
			Constructor: -1028140917,
		}
	},
	486505992: func() TLObject { // 0x1cff7e08
		return &TLMessagesGetSplitRanges{
			Constructor: 486505992,
		}
	},
	-1031349873: func() TLObject { // 0xc286d98f
		return &TLMessagesMarkDialogUnread{
			Constructor: -1031349873,
		}
	},
	585256482: func() TLObject { // 0x22e24e22
		return &TLMessagesGetDialogUnreadMarks{
			Constructor: 585256482,
		}
	},
	2119757468: func() TLObject { // 0x7e58ee9c
		return &TLMessagesClearAllDrafts{
			Constructor: 2119757468,
		}
	},
	-760547348: func() TLObject { // 0xd2aaf7ec
		return &TLMessagesUpdatePinnedMessage{
			Constructor: -760547348,
		}
	},
	283795844: func() TLObject { // 0x10ea6184
		return &TLMessagesSendVote{
			Constructor: 283795844,
		}
	},
	1941660731: func() TLObject { // 0x73bb643b
		return &TLMessagesGetPollResults{
			Constructor: 1941660731,
		}
	},
	1848369232: func() TLObject { // 0x6e2be050
		return &TLMessagesGetOnlines{
			Constructor: 1848369232,
		}
	},
	-2127811866: func() TLObject { // 0x812c2ae6
		return &TLMessagesGetStatsURL{
			Constructor: -2127811866,
		}
	},
	-2080980787: func() TLObject { // 0x83f6c0cd
		return &TLMessagesGetStatsURL{
			Constructor: -2080980787,
		}
	},
	-554301545: func() TLObject { // 0xdef60797
		return &TLMessagesEditChatAbout{
			Constructor: -554301545,
		}
	},
	-1935616490: func() TLObject { // 0x8ca0d616
		return &TLMessagesEditChatNotice{
			Constructor: -1935616490,
		}
	},
	-1517917375: func() TLObject { // 0xa5866b41
		return &TLMessagesEditChatDefaultBannedRights{
			Constructor: -1517917375,
		}
	},
	899735650: func() TLObject { // 0x35a0e062
		return &TLMessagesGetEmojiKeywords{
			Constructor: 899735650,
		}
	},
	352892591: func() TLObject { // 0x1508b6af
		return &TLMessagesGetEmojiKeywordsDifference{
			Constructor: 352892591,
		}
	},
	1318675378: func() TLObject { // 0x4e9963b2
		return &TLMessagesGetEmojiKeywordsLanguages{
			Constructor: 1318675378,
		}
	},
	-709817306: func() TLObject { // 0xd5b10c26
		return &TLMessagesGetEmojiURL{
			Constructor: -709817306,
		}
	},
	1932455680: func() TLObject { // 0x732eef00
		return &TLMessagesGetSearchCounters{
			Constructor: 1932455680,
		}
	},
	-482388461: func() TLObject { // 0xe33f5613
		return &TLMessagesRequestUrlAuth{
			Constructor: -482388461,
		}
	},
	-148247912: func() TLObject { // 0xf729ea98
		return &TLMessagesAcceptUrlAuth{
			Constructor: -148247912,
		}
	},
	1336717624: func() TLObject { // 0x4facb138
		return &TLMessagesHidePeerSettingsBar{
			Constructor: 1336717624,
		}
	},
	-490575781: func() TLObject { // 0xe2c2685b
		return &TLMessagesGetScheduledHistory{
			Constructor: -490575781,
		}
	},
	-1111817116: func() TLObject { // 0xbdbb0464
		return &TLMessagesGetScheduledMessages{
			Constructor: -1111817116,
		}
	},
	-1120369398: func() TLObject { // 0xbd38850a
		return &TLMessagesSendScheduledMessages{
			Constructor: -1120369398,
		}
	},
	1504586518: func() TLObject { // 0x59ae2b16
		return &TLMessagesDeleteScheduledMessages{
			Constructor: 1504586518,
		}
	},
	-1200736242: func() TLObject { // 0xb86e380e
		return &TLMessagesGetPollVotes{
			Constructor: -1200736242,
		}
	},
	-1257951254: func() TLObject { // 0xb5052fea
		return &TLMessagesToggleStickerSets{
			Constructor: -1257951254,
		}
	},
	-241247891: func() TLObject { // 0xf19ed96d
		return &TLMessagesGetDialogFilters{
			Constructor: -241247891,
		}
	},
	-1566780372: func() TLObject { // 0xa29cd42c
		return &TLMessagesGetSuggestedDialogFilters{
			Constructor: -1566780372,
		}
	},
	450142282: func() TLObject { // 0x1ad4a04a
		return &TLMessagesUpdateDialogFilter{
			Constructor: 450142282,
		}
	},
	-983318044: func() TLObject { // 0xc563c1e4
		return &TLMessagesUpdateDialogFiltersOrder{
			Constructor: -983318044,
		}
	},
	1608974939: func() TLObject { // 0x5fe7025b
		return &TLMessagesGetOldFeaturedStickers{
			Constructor: 1608974939,
		}
	},
	615875002: func() TLObject { // 0x24b581ba
		return &TLMessagesGetReplies{
			Constructor: 615875002,
		}
	},
	1147761405: func() TLObject { // 0x446972fd
		return &TLMessagesGetDiscussionMessage{
			Constructor: 1147761405,
		}
	},
	-147740172: func() TLObject { // 0xf731a9f4
		return &TLMessagesReadDiscussion{
			Constructor: -147740172,
		}
	},
	-265962357: func() TLObject { // 0xf025bc8b
		return &TLMessagesUnpinAllMessages{
			Constructor: -265962357,
		}
	},
	-304838614: func() TLObject { // 0xedd4882a
		return &TLUpdatesGetState{
			Constructor: -304838614,
		}
	},
	630429265: func() TLObject { // 0x25939651
		return &TLUpdatesGetDifference{
			Constructor: 630429265,
		}
	},
	51854712: func() TLObject { // 0x3173d78
		return &TLUpdatesGetChannelDifference{
			Constructor: 51854712,
		}
	},
	-1154295872: func() TLObject { // 0xbb32d7c0
		return &TLUpdatesGetChannelDifference{
			Constructor: -1154295872,
		}
	},
	1926525996: func() TLObject { // 0x72d4742c
		return &TLPhotosUpdateProfilePhoto72D4742C{
			Constructor: 1926525996,
		}
	},
	-1980559511: func() TLObject { // 0x89f30f69
		return &TLPhotosUploadProfilePhoto{
			Constructor: -1980559511,
		}
	},
	1328726168: func() TLObject { // 0x4f32c098
		return &TLPhotosUploadProfilePhoto{
			Constructor: 1328726168,
		}
	},
	-720397176: func() TLObject { // 0xd50f9c88
		return &TLPhotosUploadProfilePhoto{
			Constructor: -720397176,
		}
	},
	-2016444625: func() TLObject { // 0x87cf7f2f
		return &TLPhotosDeletePhotos{
			Constructor: -2016444625,
		}
	},
	-1848823128: func() TLObject { // 0x91cd32a8
		return &TLPhotosGetUserPhotos{
			Constructor: -1848823128,
		}
	},
	-1291540959: func() TLObject { // 0xb304a621
		return &TLUploadSaveFilePart{
			Constructor: -1291540959,
		}
	},
	-1319462148: func() TLObject { // 0xb15a9afc
		return &TLUploadGetFile{
			Constructor: -1319462148,
		}
	},
	-475607115: func() TLObject { // 0xe3a6cfb5
		return &TLUploadGetFile{
			Constructor: -475607115,
		}
	},
	-562337987: func() TLObject { // 0xde7b673d
		return &TLUploadSaveBigFilePart{
			Constructor: -562337987,
		}
	},
	619086221: func() TLObject { // 0x24e6818d
		return &TLUploadGetWebFile{
			Constructor: 619086221,
		}
	},
	536919235: func() TLObject { // 0x2000bcc3
		return &TLUploadGetCdnFile{
			Constructor: 536919235,
		}
	},
	-1691921240: func() TLObject { // 0x9b2754a8
		return &TLUploadReuploadCdnFile9B2754A8{
			Constructor: -1691921240,
		}
	},
	1302676017: func() TLObject { // 0x4da54231
		return &TLUploadGetCdnFileHashes4DA54231{
			Constructor: 1302676017,
		}
	},
	-956147407: func() TLObject { // 0xc7025931
		return &TLUploadGetFileHashes{
			Constructor: -956147407,
		}
	},
	-990308245: func() TLObject { // 0xc4f9186b
		return &TLHelpGetConfig{
			Constructor: -990308245,
		}
	},
	531836966: func() TLObject { // 0x1fb33026
		return &TLHelpGetNearestDc{
			Constructor: 531836966,
		}
	},
	1378703997: func() TLObject { // 0x522d5a7d
		return &TLHelpGetAppUpdate{
			Constructor: 1378703997,
		}
	},
	-1372724842: func() TLObject { // 0xae2de196
		return &TLHelpGetAppUpdate{
			Constructor: -1372724842,
		}
	},
	-938300290: func() TLObject { // 0xc812ac7e
		return &TLHelpGetAppUpdate{
			Constructor: -938300290,
		}
	},
	1295590211: func() TLObject { // 0x4d392343
		return &TLHelpGetInviteText{
			Constructor: 1295590211,
		}
	},
	-1532407418: func() TLObject { // 0xa4a95186
		return &TLHelpGetInviteText{
			Constructor: -1532407418,
		}
	},
	-1663104819: func() TLObject { // 0x9cdf08cd
		return &TLHelpGetSupport{
			Constructor: -1663104819,
		}
	},
	-1877938321: func() TLObject { // 0x9010ef6f
		return &TLHelpGetAppChangelog{
			Constructor: -1877938321,
		}
	},
	-333262899: func() TLObject { // 0xec22cfcd
		return &TLHelpSetBotUpdatesStatus{
			Constructor: -333262899,
		}
	},
	1375900482: func() TLObject { // 0x52029342
		return &TLHelpGetCdnConfig{
			Constructor: 1375900482,
		}
	},
	1036054804: func() TLObject { // 0x3dc0f114
		return &TLHelpGetRecentMeUrls{
			Constructor: 1036054804,
		}
	},
	749019089: func() TLObject { // 0x2ca51fd1
		return &TLHelpGetTermsOfServiceUpdate{
			Constructor: 749019089,
		}
	},
	-294455398: func() TLObject { // 0xee72f79a
		return &TLHelpAcceptTermsOfService{
			Constructor: -294455398,
		}
	},
	1072547679: func() TLObject { // 0x3fedc75f
		return &TLHelpGetDeepLinkInfo{
			Constructor: 1072547679,
		}
	},
	-1735311088: func() TLObject { // 0x98914110
		return &TLHelpGetAppConfig{
			Constructor: -1735311088,
		}
	},
	1862465352: func() TLObject { // 0x6f02f748
		return &TLHelpSaveAppLog{
			Constructor: 1862465352,
		}
	},
	-966677240: func() TLObject { // 0xc661ad08
		return &TLHelpGetPassportConfig{
			Constructor: -966677240,
		}
	},
	-748624084: func() TLObject { // 0xd360e72c
		return &TLHelpGetSupportName{
			Constructor: -748624084,
		}
	},
	59377875: func() TLObject { // 0x38a08d3
		return &TLHelpGetUserInfo{
			Constructor: 59377875,
		}
	},
	1723407216: func() TLObject { // 0x66b91b70
		return &TLHelpEditUserInfo{
			Constructor: 1723407216,
		}
	},
	-1063816159: func() TLObject { // 0xc0977421
		return &TLHelpGetPromoData{
			Constructor: -1063816159,
		}
	},
	505748629: func() TLObject { // 0x1e251c95
		return &TLHelpHidePromoData{
			Constructor: 505748629,
		}
	},
	125807007: func() TLObject { // 0x77fa99f
		return &TLHelpDismissSuggestion{
			Constructor: 125807007,
		}
	},
	1935116200: func() TLObject { // 0x735787a8
		return &TLHelpGetCountriesList{
			Constructor: 1935116200,
		}
	},
	-871347913: func() TLObject { // 0xcc104937
		return &TLChannelsReadHistory{
			Constructor: -871347913,
		}
	},
	-2067661490: func() TLObject { // 0x84c1fd4e
		return &TLChannelsDeleteMessages{
			Constructor: -2067661490,
		}
	},
	-787622117: func() TLObject { // 0xd10dd71b
		return &TLChannelsDeleteUserHistory{
			Constructor: -787622117,
		}
	},
	-32999408: func() TLObject { // 0xfe087810
		return &TLChannelsReportSpam{
			Constructor: -32999408,
		}
	},
	-1383294429: func() TLObject { // 0xad8c9a23
		return &TLChannelsGetMessages{
			Constructor: -1383294429,
		}
	},
	-1814580409: func() TLObject { // 0x93d7b347
		return &TLChannelsGetMessages{
			Constructor: -1814580409,
		}
	},
	306054633: func() TLObject { // 0x123e05e9
		return &TLChannelsGetParticipants{
			Constructor: 306054633,
		}
	},
	1416484774: func() TLObject { // 0x546dd7a6
		return &TLChannelsGetParticipant{
			Constructor: 1416484774,
		}
	},
	176122811: func() TLObject { // 0xa7f6bbb
		return &TLChannelsGetChannels{
			Constructor: 176122811,
		}
	},
	141781513: func() TLObject { // 0x8736a09
		return &TLChannelsGetFullChannel{
			Constructor: 141781513,
		}
	},
	-114523545: func() TLObject { // 0xf92c8267
		return &TLChannelsCreateChannel{
			Constructor: -114523545,
		}
	},
	1029681423: func() TLObject { // 0x3d5fb10f
		return &TLChannelsCreateChannel{
			Constructor: 1029681423,
		}
	},
	-192332417: func() TLObject { // 0xf4893d7f
		return &TLChannelsCreateChannel{
			Constructor: -192332417,
		}
	},
	-751007486: func() TLObject { // 0xd33c8902
		return &TLChannelsEditAdmin{
			Constructor: -751007486,
		}
	},
	1895338938: func() TLObject { // 0x70f893ba
		return &TLChannelsEditAdmin{
			Constructor: 1895338938,
		}
	},
	548962836: func() TLObject { // 0x20b88214
		return &TLChannelsEditAdmin{
			Constructor: 548962836,
		}
	},
	1450044624: func() TLObject { // 0x566decd0
		return &TLChannelsEditTitle{
			Constructor: 1450044624,
		}
	},
	-248621111: func() TLObject { // 0xf12e57c9
		return &TLChannelsEditPhoto{
			Constructor: -248621111,
		}
	},
	283557164: func() TLObject { // 0x10e6bd2c
		return &TLChannelsCheckUsername{
			Constructor: 283557164,
		}
	},
	890549214: func() TLObject { // 0x3514b3de
		return &TLChannelsUpdateUsername{
			Constructor: 890549214,
		}
	},
	615851205: func() TLObject { // 0x24b524c5
		return &TLChannelsJoinChannel{
			Constructor: 615851205,
		}
	},
	-130635115: func() TLObject { // 0xf836aa95
		return &TLChannelsLeaveChannel{
			Constructor: -130635115,
		}
	},
	429865580: func() TLObject { // 0x199f3a6c
		return &TLChannelsInviteToChannel{
			Constructor: 429865580,
		}
	},
	-1072619549: func() TLObject { // 0xc0111fe3
		return &TLChannelsDeleteChannel{
			Constructor: -1072619549,
		}
	},
	-432034325: func() TLObject { // 0xe63fadeb
		return &TLChannelsExportMessageLink{
			Constructor: -432034325,
		}
	},
	-826838685: func() TLObject { // 0xceb77163
		return &TLChannelsExportMessageLink{
			Constructor: -826838685,
		}
	},
	-934882771: func() TLObject { // 0xc846d22d
		return &TLChannelsExportMessageLink{
			Constructor: -934882771,
		}
	},
	527021574: func() TLObject { // 0x1f69b606
		return &TLChannelsToggleSignatures{
			Constructor: 527021574,
		}
	},
	-122669393: func() TLObject { // 0xf8b036af
		return &TLChannelsGetAdminedPublicChannels{
			Constructor: -122669393,
		}
	},
	-1920105769: func() TLObject { // 0x8d8d82d7
		return &TLChannelsGetAdminedPublicChannels{
			Constructor: -1920105769,
		}
	},
	1920559378: func() TLObject { // 0x72796912
		return &TLChannelsEditBanned{
			Constructor: 1920559378,
		}
	},
	-1076292147: func() TLObject { // 0xbfd915cd
		return &TLChannelsEditBanned{
			Constructor: -1076292147,
		}
	},
	870184064: func() TLObject { // 0x33ddf480
		return &TLChannelsGetAdminLog{
			Constructor: 870184064,
		}
	},
	-359881479: func() TLObject { // 0xea8ca4f9
		return &TLChannelsSetStickers{
			Constructor: -359881479,
		}
	},
	-357180360: func() TLObject { // 0xeab5dc38
		return &TLChannelsReadMessageContents{
			Constructor: -357180360,
		}
	},
	-1355375294: func() TLObject { // 0xaf369d42
		return &TLChannelsDeleteHistory{
			Constructor: -1355375294,
		}
	},
	-356796084: func() TLObject { // 0xeabbb94c
		return &TLChannelsTogglePreHistoryHidden{
			Constructor: -356796084,
		}
	},
	-2092831552: func() TLObject { // 0x8341ecc0
		return &TLChannelsGetLeftChannels{
			Constructor: -2092831552,
		}
	},
	-170208392: func() TLObject { // 0xf5dad378
		return &TLChannelsGetGroupsForDiscussion{
			Constructor: -170208392,
		}
	},
	1079520178: func() TLObject { // 0x40582bb2
		return &TLChannelsSetDiscussionGroup{
			Constructor: 1079520178,
		}
	},
	-1892102881: func() TLObject { // 0x8f38cd1f
		return &TLChannelsEditCreator{
			Constructor: -1892102881,
		}
	},
	1491484525: func() TLObject { // 0x58e63f6d
		return &TLChannelsEditLocation{
			Constructor: 1491484525,
		}
	},
	-304832784: func() TLObject { // 0xedd49ef0
		return &TLChannelsToggleSlowMode{
			Constructor: -304832784,
		}
	},
	300429806: func() TLObject { // 0x11e831ee
		return &TLChannelsGetInactiveChannels{
			Constructor: 300429806,
		}
	},
	-1440257555: func() TLObject { // 0xaa2769ed
		return &TLBotsSendCustomRequest{
			Constructor: -1440257555,
		}
	},
	-434028723: func() TLObject { // 0xe6213f4d
		return &TLBotsAnswerWebhookJSONQuery{
			Constructor: -434028723,
		}
	},
	-2141370634: func() TLObject { // 0x805d46f6
		return &TLBotsSetBotCommands{
			Constructor: -2141370634,
		}
	},
	-1712285883: func() TLObject { // 0x99f09745
		return &TLPaymentsGetPaymentForm{
			Constructor: -1712285883,
		}
	},
	-1601001088: func() TLObject { // 0xa092a980
		return &TLPaymentsGetPaymentReceipt{
			Constructor: -1601001088,
		}
	},
	1997180532: func() TLObject { // 0x770a8e74
		return &TLPaymentsValidateRequestedInfo{
			Constructor: 1997180532,
		}
	},
	730364339: func() TLObject { // 0x2b8879b3
		return &TLPaymentsSendPaymentForm{
			Constructor: 730364339,
		}
	},
	578650699: func() TLObject { // 0x227d824b
		return &TLPaymentsGetSavedInfo{
			Constructor: 578650699,
		}
	},
	-667062079: func() TLObject { // 0xd83d70c1
		return &TLPaymentsClearSavedInfo{
			Constructor: -667062079,
		}
	},
	779736953: func() TLObject { // 0x2e79d779
		return &TLPaymentsGetBankCardData{
			Constructor: 779736953,
		}
	},
	-251435136: func() TLObject { // 0xf1036780
		return &TLStickersCreateStickerSet{
			Constructor: -251435136,
		}
	},
	-1680314774: func() TLObject { // 0x9bd86e6a
		return &TLStickersCreateStickerSet{
			Constructor: -1680314774,
		}
	},
	-143257775: func() TLObject { // 0xf7760f51
		return &TLStickersRemoveStickerFromSet{
			Constructor: -143257775,
		}
	},
	-4795190: func() TLObject { // 0xffb6d4ca
		return &TLStickersChangeStickerPosition{
			Constructor: -4795190,
		}
	},
	-2041315650: func() TLObject { // 0x8653febe
		return &TLStickersAddStickerToSet{
			Constructor: -2041315650,
		}
	},
	-1707717072: func() TLObject { // 0x9a364e30
		return &TLStickersSetStickerSetThumb{
			Constructor: -1707717072,
		}
	},
	1430593449: func() TLObject { // 0x55451fa9
		return &TLPhoneGetCallConfig{
			Constructor: 1430593449,
		}
	},
	1124046573: func() TLObject { // 0x42ff96ed
		return &TLPhoneRequestCall{
			Constructor: 1124046573,
		}
	},
	1536537556: func() TLObject { // 0x5b95b3d4
		return &TLPhoneRequestCall{
			Constructor: 1536537556,
		}
	},
	1003664544: func() TLObject { // 0x3bd2b4a0
		return &TLPhoneAcceptCall{
			Constructor: 1003664544,
		}
	},
	788404002: func() TLObject { // 0x2efe1722
		return &TLPhoneConfirmCall{
			Constructor: 788404002,
		}
	},
	399855457: func() TLObject { // 0x17d54f61
		return &TLPhoneReceivedCall{
			Constructor: 399855457,
		}
	},
	-1295269440: func() TLObject { // 0xb2cbc1c0
		return &TLPhoneDiscardCall{
			Constructor: -1295269440,
		}
	},
	2027164582: func() TLObject { // 0x78d413a6
		return &TLPhoneDiscardCall{
			Constructor: 2027164582,
		}
	},
	1508562471: func() TLObject { // 0x59ead627
		return &TLPhoneSetCallRating{
			Constructor: 1508562471,
		}
	},
	475228724: func() TLObject { // 0x1c536a34
		return &TLPhoneSetCallRating{
			Constructor: 475228724,
		}
	},
	662363518: func() TLObject { // 0x277add7e
		return &TLPhoneSaveCallDebug{
			Constructor: 662363518,
		}
	},
	-8744061: func() TLObject { // 0xff7a9383
		return &TLPhoneSendSignalingData{
			Constructor: -8744061,
		}
	},
	-219008246: func() TLObject { // 0xf2f2330a
		return &TLLangpackGetLangPack{
			Constructor: -219008246,
		}
	},
	-1699363442: func() TLObject { // 0x9ab5c58e
		return &TLLangpackGetLangPack{
			Constructor: -1699363442,
		}
	},
	-269862909: func() TLObject { // 0xefea3803
		return &TLLangpackGetStrings{
			Constructor: -269862909,
		}
	},
	773776152: func() TLObject { // 0x2e1ee318
		return &TLLangpackGetStrings{
			Constructor: 773776152,
		}
	},
	-845657435: func() TLObject { // 0xcd984aa5
		return &TLLangpackGetDifference{
			Constructor: -845657435,
		}
	},
	-1655576556: func() TLObject { // 0x9d51e814
		return &TLLangpackGetDifference{
			Constructor: -1655576556,
		}
	},
	187583869: func() TLObject { // 0xb2e4d7d
		return &TLLangpackGetDifference{
			Constructor: 187583869,
		}
	},
	1120311183: func() TLObject { // 0x42c6978f
		return &TLLangpackGetLanguages{
			Constructor: 1120311183,
		}
	},
	-2146445955: func() TLObject { // 0x800fd57d
		return &TLLangpackGetLanguages{
			Constructor: -2146445955,
		}
	},
	1784243458: func() TLObject { // 0x6a596502
		return &TLLangpackGetLanguage{
			Constructor: 1784243458,
		}
	},
	1749536939: func() TLObject { // 0x6847d0ab
		return &TLFoldersEditPeerFolders{
			Constructor: 1749536939,
		}
	},
	472471681: func() TLObject { // 0x1c295881
		return &TLFoldersDeleteFolder{
			Constructor: 472471681,
		}
	},
	-1421720550: func() TLObject { // 0xab42441a
		return &TLStatsGetBroadcastStats{
			Constructor: -1421720550,
		}
	},
	-433058374: func() TLObject { // 0xe6300dba
		return &TLStatsGetBroadcastStats{
			Constructor: -433058374,
		}
	},
	1646092192: func() TLObject { // 0x621d5fa0
		return &TLStatsLoadAsyncGraph{
			Constructor: 1646092192,
		}
	},
	-589330937: func() TLObject { // 0xdcdf8607
		return &TLStatsGetMegagroupStats{
			Constructor: -589330937,
		}
	},
	1445996571: func() TLObject { // 0x5630281b
		return &TLStatsGetMessagePublicForwards{
			Constructor: 1445996571,
		}
	},
	-1226791947: func() TLObject { // 0xb6e0a3f5
		return &TLStatsGetMessageStats{
			Constructor: -1226791947,
		}
	},
	-2030965169: func() TLObject { // 0x86f1ee4f
		return &TLWalletGetInfo{
			Constructor: -2030965169,
		}
	},
	1121001551: func() TLObject { // 0x42d1204f
		return &TLWalletGetRecords{
			Constructor: 1121001551,
		}
	},
	-579032072: func() TLObject { // 0xdd7cabf8
		return &TLBlogsGetUser{
			Constructor: -579032072,
		}
	},
	1171744961: func() TLObject { // 0x45d768c1
		return &TLBlogsFollow{
			Constructor: 1171744961,
		}
	},
	1283676963: func() TLObject { // 0x4c835b23
		return &TLBlogsLike{
			Constructor: 1283676963,
		}
	},
	-434256691: func() TLObject { // 0xe61dc4cd
		return &TLBlogsSendComment{
			Constructor: -434256691,
		}
	},
	-1582875955: func() TLObject { // 0xa1a73acd
		return &TLBlogsReward{
			Constructor: -1582875955,
		}
	},
	-1240125656: func() TLObject { // 0xb6152f28
		return &TLBlogsGetFollows{
			Constructor: -1240125656,
		}
	},
	-782458473: func() TLObject { // 0xd15ca197
		return &TLBlogsGetFans{
			Constructor: -782458473,
		}
	},
	1689428782: func() TLObject { // 0x64b2a32e
		return &TLBlogsCreateGroupTag{
			Constructor: 1689428782,
		}
	},
	568326946: func() TLObject { // 0x21dffb22
		return &TLBlogsDeleteGroupTag{
			Constructor: 568326946,
		}
	},
	316490375: func() TLObject { // 0x12dd4287
		return &TLBlogsEditGroupTag{
			Constructor: 316490375,
		}
	},
	1457176378: func() TLObject { // 0x56dabf3a
		return &TLBlogsAddGroupTagMember{
			Constructor: 1457176378,
		}
	},
	-40812416: func() TLObject { // 0xfd914080
		return &TLBlogsDeleteGroupTagMember{
			Constructor: -40812416,
		}
	},
	-969478218: func() TLObject { // 0xc636efb6
		return &TLBlogsGetGroupTags{
			Constructor: -969478218,
		}
	},
	-618794919: func() TLObject { // 0xdb1df059
		return &TLBlogsSendBlog{
			Constructor: -618794919,
		}
	},
	1070871648: func() TLObject { // 0x3fd43460
		return &TLBlogsSendBlog{
			Constructor: 1070871648,
		}
	},
	-1008379326: func() TLObject { // 0xc3e55a42
		return &TLBlogsDeleteBlog{
			Constructor: -1008379326,
		}
	},
	2135218055: func() TLObject { // 0x7f44d787
		return &TLBlogsGetBlogs{
			Constructor: 2135218055,
		}
	},
	1794466639: func() TLObject { // 0x6af5634f
		return &TLBlogsGetCommentList{
			Constructor: 1794466639,
		}
	},
	-1547563864: func() TLObject { // 0xa3c20ca8
		return &TLBlogsReadHistory{
			Constructor: -1547563864,
		}
	},
	-2088311748: func() TLObject { // 0x8386e43c
		return &TLBlogsGetHistory{
			Constructor: -2088311748,
		}
	},
	1565320413: func() TLObject { // 0x5d4ce4dd
		return &TLBlogsGetComments{
			Constructor: 1565320413,
		}
	},
	613461553: func() TLObject { // 0x2490ae31
		return &TLBlogsGetLikes{
			Constructor: 613461553,
		}
	},
	-268169030: func() TLObject { // 0xf00410ba
		return &TLBlogsGetUnreads{
			Constructor: -268169030,
		}
	},
	426044861: func() TLObject { // 0x1964edbd
		return &TLBlogsGetTopics{
			Constructor: 426044861,
		}
	},
	122645497: func() TLObject { // 0x74f6bf9
		return &TLBlogsGetHotTopics{
			Constructor: 122645497,
		}
	},
	-993483427: func() TLObject { // 0xc4c8a55d
		return &TLMessagesGetMessagesViewsC4C8A55D{
			Constructor: -993483427,
		}
	},
	-256159406: func() TLObject { // 0xf0bb5152
		return &TLPhotosUpdateProfilePhotoF0BB5152{
			Constructor: -256159406,
		}
	},
	-1080395925: func() TLObject { // 0xbf9a776b
		return &TLMessagesSearchGifs{
			Constructor: -1080395925,
		}
	},
	1031231713: func() TLObject { // 0x3d7758e1
		return &TLHelpGetProxyData{
			Constructor: 1031231713,
		}
	},
	-490089666: func() TLObject { // 0xe2c9d33e
		return &TLWalletSendLiteRequest{
			Constructor: -490089666,
		}
	},
	190313286: func() TLObject { // 0xb57f346
		return &TLWalletGetKeySecretSalt{
			Constructor: 190313286,
		}
	},
	-1902823612: func() TLObject { // 0x8e953744
		return &TLContactsDeleteContact{
			Constructor: -1902823612,
		}
	},
	1504393374: func() TLObject { // 0x59ab389e
		return &TLContactsDeleteContacts59AB389E{
			Constructor: 1504393374,
		}
	},
	-1460572005: func() TLObject { // 0xa8f1709b
		return &TLMessagesHideReportSpam{
			Constructor: -1460572005,
		}
	},
	445117188: func() TLObject { // 0x1a87f304
		return &TLChannelsGetBroadcastsForDiscussion{
			Constructor: 445117188,
		}
	},
	-1068696894: func() TLObject { // 0xc04cfac2
		return &TLAccountGetWallPapersC04CFAC2{
			Constructor: -1068696894,
		}
	},
	-326379039: func() TLObject { // 0xec8bd9e1
		return &TLMessagesToggleChatAdmins{
			Constructor: -326379039,
		}
	},
	333610782: func() TLObject { // 0x13e27f1e
		return &TLChannelsEditAbout{
			Constructor: 333610782,
		}
	},
	-950663035: func() TLObject { // 0xc7560885
		return &TLChannelsExportInvite{
			Constructor: -950663035,
		}
	},
	1231065863: func() TLObject { // 0x49609307
		return &TLChannelsToggleInvites{
			Constructor: 1231065863,
		}
	},
	-1490162350: func() TLObject { // 0xa72ded52
		return &TLChannelsUpdatePinnedMessage{
			Constructor: -1490162350,
		}
	},
	-2065352905: func() TLObject { // 0x84e53737
		return &TLContactsExportCard{
			Constructor: -2065352905,
		}
	},
	1340184318: func() TLObject { // 0x4fe196fe
		return &TLContactsImportCard{
			Constructor: 1340184318,
		}
	},
	1998331287: func() TLObject { // 0x771c1d97
		return &TLAuthSendInvites{
			Constructor: 1998331287,
		}
	},
	-1906722841: func() TLObject { // 0x8e59b7e7
		return &TLHelpGetTermsOfService{
			Constructor: -1906722841,
		}
	},
	889286899: func() TLObject { // 0x350170f3
		return &TLHelpGetTermsOfService{
			Constructor: 889286899,
		}
	},
	1877286395: func() TLObject { // 0x6fe51dfb
		return &TLAuthCheckPhone{
			Constructor: 1877286395,
		}
	},
	194049104: func() TLObject { // 0xb90f450
		return &TLChannelsGetFeed{
			Constructor: 194049104,
		}
	},
	403799538: func() TLObject { // 0x18117df2
		return &TLChannelsGetFeed{
			Constructor: 403799538,
		}
	},
	-2009967767: func() TLObject { // 0x88325369
		return &TLChannelsSearchFeed{
			Constructor: -2009967767,
		}
	},
	-657579154: func() TLObject { // 0xd8ce236e
		return &TLChannelsGetFeedSources{
			Constructor: -657579154,
		}
	},
	-5016303: func() TLObject { // 0xffb37511
		return &TLChannelsChangeFeedBroadcastFFB37511{
			Constructor: -5016303,
		}
	},
	-360661074: func() TLObject { // 0xea80bfae
		return &TLChannelsSetFeedBroadcastsEA80BFAE{
			Constructor: -360661074,
		}
	},
	163774749: func() TLObject { // 0x9c3011d
		return &TLChannelsReadFeed{
			Constructor: 163774749,
		}
	},
	865483769: func() TLObject { // 0x33963bf9
		return &TLMessagesForwardMessage{
			Constructor: 865483769,
		}
	},
	452533257: func() TLObject { // 0x1af91c09
		return &TLUploadReuploadCdnFile1AF91C09{
			Constructor: 452533257,
		}
	},
	-149567365: func() TLObject { // 0xf715c87b
		return &TLUploadGetCdnFileHashesF715C87B{
			Constructor: -149567365,
		}
	},
	623413022: func() TLObject { // 0x2528871e
		return &TLChannelsChangeFeedBroadcast2528871E{
			Constructor: 623413022,
		}
	},
	2123479282: func() TLObject { // 0x7e91b8f2
		return &TLChannelsSetFeedBroadcasts7E91B8F2{
			Constructor: 2123479282,
		}
	},
	-1336990448: func() TLObject { // 0xb04f2510
		return &TLMessagesReadHistoryB04F2510{
			Constructor: -1336990448,
		}
	},
	-608789858: func() TLObject { // 0xdbb69a9e
		return &TLHelpGetScheme{
			Constructor: -608789858,
		}
	},
	-1058929929: func() TLObject { // 0xc0e202f7
		return &TLHelpTest{
			Constructor: -1058929929,
		}
	},
	1270152629: func() TLObject { // 0x4bb4fdb5
		return &TLAccountCreatePredefinedUser{
			Constructor: 1270152629,
		}
	},
	1595162427: func() TLObject { // 0x5f143f3b
		return &TLAccountUpdatePredefinedUsername{
			Constructor: 1595162427,
		}
	},
	2091625414: func() TLObject { // 0x7cababc6
		return &TLAccountUpdatePredefinedProfile{
			Constructor: 2091625414,
		}
	},
	1361800622: func() TLObject { // 0x512b6dae
		return &TLAccountUpdateVerified{
			Constructor: 1361800622,
		}
	},
	1178012274: func() TLObject { // 0x46370a72
		return &TLAccountUpdatePredefinedVerified{
			Constructor: 1178012274,
		}
	},
	169499657: func() TLObject { // 0xa1a5c09
		return &TLAccountUpdatePredefinedCode{
			Constructor: 169499657,
		}
	},
	1664018599: func() TLObject { // 0x632ee8a7
		return &TLAuthToggleBan{
			Constructor: 1664018599,
		}
	},
	1749863255: func() TLObject { // 0x684ccb57
		return &TLUsersGetPredefinedUser{
			Constructor: 1749863255,
		}
	},
	1795489223: func() TLObject { // 0x6b04fdc7
		return &TLUsersGetPredefinedUsers{
			Constructor: 1795489223,
		}
	},
	201236394: func() TLObject { // 0xbfe9faa
		return &TLUsersGetMe{
			Constructor: 201236394,
		}
	},
	1511592262: func() TLObject { // 0x5a191146
		return &TLBizInvokeBizDataRaw{
			Constructor: 1511592262,
		}
	},
}

func NewTLObjectByClassID(classId int32) TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}
