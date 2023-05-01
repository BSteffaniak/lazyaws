/*

本翻译文件中的词语的翻译参考了 https://github.com/progit/progit2-zh/blob/master/TRANSLATION_NOTES.asc。
下方的术语对照表是对其的补充。

Translation in this file refer to https://github.com/progit/progit2-zh/blob/master/TRANSLATION_NOTES.asc.
Glossary below is a supplement of that documentation.

Glossary 术语对照表

change   更改
fixup    修正
reset    重置

*/

package i18n

const chineseIntroPopupMessage = `
感谢使用 lazyaws！你真的太棒了。下面几点你可能会感兴趣：

 1) 观看此视频，快速了解 lazyaws 的功能：
      https://youtu.be/CPLdltN7wgE

 2) 记得看看最新发行说明：
      https://github.com/BSteffaniak/lazyaws/releases

 3) 使用 git 说明你是一位程序员！你可以和我们一起让 lazyaws 变得更好。
    考虑为本项目做些贡献吧：
      https://github.com/BSteffaniak/lazyaws
    你也可以直接赞助，并告诉我哪里需要改进，点右下角的捐赠按钮就好了。
    哪怕只是给仓库点个星星也很棒！
`

// exporting this so we can use it in tests
func chineseTranslationSet() TranslationSet {
	return TranslationSet{
		NotEnoughSpace:                      "没有足够的空间来渲染面板",
		DiffTitle:                           "差异",
		FilesTitle:                          "文件",
		BranchesTitle:                       "分支",
		CommitsTitle:                        "提交",
		StashTitle:                          "贮藏",
		UnstagedChanges:                     `未暂存更改`,
		StagedChanges:                       `已暂存更改`,
		MainTitle:                           "主要",
		StagingTitle:                        "正在暂存",
		MergingTitle:                        "正在合并",
		NormalTitle:                         "正常",
		CommitMessage:                       "提交信息",
		CredentialsUsername:                 "用户名",
		CredentialsPassword:                 "密码",
		CredentialsPassphrase:               "输入 SSH 密钥的密码",
		PassUnameWrong:                      "密码 和/或 用户名错误",
		CommitChanges:                       "提交更改",
		AmendLastCommit:                     "修补最后一次提交",
		AmendLastCommitTitle:                "修补最后一次提交",
		SureToAmend:                         "您确定要修补上一次提交吗？之后您可以从提交面板更改提交消息。",
		NoCommitToAmend:                     "没有需要提交的修补。",
		CommitChangesWithEditor:             "提交更改（使用编辑器编辑提交信息）",
		StatusTitle:                         "状态",
		LcNavigate:                          "导航",
		LcMenu:                              "菜单",
		LcExecute:                           "执行",
		LcToggleStaged:                      "切换暂存状态",
		LcToggleStagedAll:                   "切换所有文件的暂存状态",
		LcToggleTreeView:                    "切换文件树视图",
		LcOpenMergeTool:                     "打开外部合并工具 (git mergetool)",
		LcRefresh:                           "刷新",
		LcPush:                              "推送",
		LcPull:                              "拉取",
		LcScroll:                            "滚动",
		MergeConflictsTitle:                 "合并冲突",
		LcCheckout:                          "检出",
		NoChangedFiles:                      "没有更改过文件",
		PullWait:                            "正在拉取…",
		PushWait:                            "正在推送…",
		FetchWait:                           "正在抓取…",
		LcSoftReset:                         "软重置",
		AlreadyCheckedOutBranch:             "您已经检出至此分支",
		SureForceCheckout:                   "您确定要强制检出吗？您将丢失所有本地更改",
		ForceCheckoutBranch:                 "强制检出分支",
		BranchName:                          "分支名称",
		NewBranchNameBranchOff:              "新分支名称（基于 {{.branchName}}）",
		CantDeleteCheckOutBranch:            "您不能删除已检出的分支！",
		DeleteBranch:                        "删除分支",
		DeleteBranchMessage:                 "您确定要删除分支 {{.selectedBranchName}} 吗？",
		ForceDeleteBranchMessage:            "{{.selectedBranchName}} 还没有被完全合并。您确定要删除它吗？",
		LcRebaseBranch:                      "将已检出的分支变基到该分支",
		CantRebaseOntoSelf:                  "您不能将分支变基到其自身",
		CantMergeBranchIntoItself:           "您不能将分支合并到其自身",
		LcForceCheckout:                     "强制检出",
		LcCheckoutByName:                    "按名称检出",
		LcNewBranch:                         "新分支",
		LcDeleteBranch:                      "删除分支",
		NoBranchesThisRepo:                  "此仓库中没有分支",
		CommitMessageConfirm:                "{{.keyBindClose}}：关闭，{{.keyBindNewLine}}：新行，{{.keyBindConfirm}}：确认",
		CommitWithoutMessageErr:             "您必须编写提交消息才能进行提交",
		CloseConfirm:                        "{{.keyBindClose}}：关闭，{{.keyBindConfirm}}：确认",
		LcClose:                             "关闭",
		LcQuit:                              "退出",
		LcSquashDown:                        "向下压缩",
		LcFixupCommit:                       "修正提交（fixup）",
		NoCommitsThisBranch:                 "该分支没有提交",
		CannotSquashOrFixupFirstCommit:      "There's no commit below to squash into",
		Fixup:                               "修正（fixup）",
		SureFixupThisCommit:                 "您确定要“修正”此提交吗？它将合并到下面的提交中",
		SureSquashThisCommit:                "您确定要将这个提交压缩到下面的提交中吗？",
		Squash:                              "压缩",
		LcPickCommit:                        "选择提交（变基过程中）",
		LcRevertCommit:                      "还原提交",
		LcRewordCommit:                      "改写提交",
		LcDeleteCommit:                      "删除提交",
		LcMoveDownCommit:                    "下移提交",
		LcMoveUpCommit:                      "上移提交",
		LcEditCommit:                        "编辑提交",
		LcAmendToCommit:                     "用已暂存的更改来修补提交",
		LcRenameCommitEditor:                "使用编辑器重命名提交",
		Error:                               "错误",
		LcSelectHunk:                        "切换区块",
		LcNavigateConflicts:                 "浏览冲突",
		LcPickHunk:                          "选中区块",
		LcPickAllHunks:                      "选中所有区块",
		LcUndo:                              "撤销",
		LcUndoReflog:                        "（通过 reflog）撤销「实验功能」",
		LcRedoReflog:                        "（通过 reflog）重做「实验功能」",
		LcPop:                               "应用并删除",
		LcDrop:                              "删除",
		LcApply:                             "应用",
		NoStashEntries:                      "没有贮藏条目",
		StashDrop:                           "删除贮藏",
		SureDropStashEntry:                  "您确定要删除此贮藏条目吗？",
		StashPop:                            "应用并删除贮藏",
		SurePopStashEntry:                   "您确定要应用并删除此贮藏条目吗？",
		StashApply:                          "应用贮藏",
		SureApplyStashEntry:                 "您确定要应用此贮藏条目？",
		NoTrackedStagedFilesStash:           "没有可以贮藏的已跟踪/暂存文件",
		StashChanges:                        "贮藏更改",
		LcRenameStash:                       "rename stash",
		RenameStashPrompt:                   "Rename stash: {{.stashName}}",
		OpenConfig:                          "打开配置文件",
		EditConfig:                          "编辑配置文件",
		ForcePush:                           "强制推送",
		ForcePushPrompt:                     "您的分支已与远程分支不同。按‘esc’取消，或‘enter’强制推送.",
		ForcePushDisabled:                   "您的分支已与远程分支不同, 并且您已经禁用了强行推送",
		UpdatesRejectedAndForcePushDisabled: "更新被拒绝，您已禁用强制推送",
		LcCheckForUpdate:                    "检查更新",
		CheckingForUpdates:                  "正在检查更新…",
		OnLatestVersionErr:                  "已是最新版本",
		MajorVersionErr:                     "新版本 ({{.newVersion}}) 与当前版本 ({{.currentVersion}}) 相比，具有非向后兼容的更改",
		CouldNotFindBinaryErr:               "在 {{.url}} 处找不到任何二进制文件",
		MergeToolTitle:                      "合并工具",
		MergeToolPrompt:                     "确定要打开 `git mergetool` 吗?",
		IntroPopupMessage:                   chineseIntroPopupMessage,
		GitconfigParseErr:                   `由于存在未加引号的'\'字符，因此 Gogit 无法解析您的 gitconfig 文件。删除它们应该可以解决问题。`,
		LcEditFile:                          `编辑文件`,
		LcOpenFile:                          `打开文件`,
		LcIgnoreFile:                        `添加到 .gitignore`,
		LcRefreshFiles:                      `刷新文件`,
		LcMergeIntoCurrentBranch:            `合并到当前检出的分支`,
		ConfirmQuit:                         `您确定要退出吗？`,
		SwitchRepo:                          `切换到最近的仓库`,
		LcAllBranchesLogGraph:               `显示所有分支的日志`,
		UnsupportedGitService:               `不支持的 git 服务`,
		LcCreatePullRequest:                 `创建抓取请求`,
		LcCopyPullRequestURL:                `将抓取请求 URL 复制到剪贴板`,
		NoBranchOnRemote:                    `该分支在远程上不存在. 您需要先将其推送到远程.`,
		LcFetch:                             `抓取`,
		NoAutomaticGitFetchTitle:            `无法自动进行 "git fetch"`,
		NoAutomaticGitFetchBody:             `Lazyaws 不能在私人仓库中使用 "git fetch"; 请在文件面板中使用 'f' 手动运行 "git fetch"`,
		FileEnter:                           `暂存单个 块/行 用于文件, 或 折叠/展开 目录`,
		FileStagingRequirements:             `只能暂存跟踪文件的单独行`,
		StageSelection:                      `切换行暂存状态`,
		ResetSelection:                      `取消变更 (git reset)`,
		ToggleDragSelect:                    `切换拖动选择`,
		ToggleSelectHunk:                    `切换选择区块`,
		ToggleSelectionForPatch:             `添加/移除 行到补丁`,
		ToggleStagingPanel:                  `切换到其他面板`,
		ReturnToFilesPanel:                  `返回文件面板`,
		FastForward:                         `从上游快进此分支`,
		Fetching:                            "抓取并快进 {{.from}} -> {{.to}} ...",
		FoundConflicts:                      "冲突！需要中止操作请按‘esc’, 否则按‘enter’",
		FoundConflictsTitle:                 "自动合并失败",
		PickHunk:                            "选中区块",
		PickAllHunks:                        "选中所有区块",
		ViewMergeRebaseOptions:              "查看 合并/变基 选项",
		NotMergingOrRebasing:                "您目前既不进行变基也不进行合并",
		RecentRepos:                         "最近的仓库",
		MergeOptionsTitle:                   "合并选项",
		RebaseOptionsTitle:                  "变基选项",
		CommitMessageTitle:                  "提交讯息",
		LocalBranchesTitle:                  "分支页面",
		SearchTitle:                         "搜索",
		TagsTitle:                           "标签页面",
		MenuTitle:                           "菜单",
		RemotesTitle:                        "远程页面",
		RemoteBranchesTitle:                 "远程分支",
		PatchBuildingTitle:                  "构建补丁中",
		InformationTitle:                    "信息",
		SecondaryTitle:                      "次要",
		ReflogCommitsTitle:                  "Reflog 页面",
		GlobalTitle:                         "全局键绑定",
		ConflictsResolved:                   "已解决所有冲突。是否继续？",
		ConfirmMerge:                        "您确定要将分支 {{.selectedBranch}} 合并到 {{.checkedOutBranch}} 吗？",
		FwdNoUpstream:                       "此分支没有上游，无法快进",
		FwdNoLocalUpstream:                  "此分支的远程未在本地注册，无法快进",
		FwdCommitsToPush:                    "此分支带有尚未推送的提交，无法快进",
		ErrorOccurred:                       "发生错误！请在以下位置创建 issue",
		NoRoom:                              "空间不足",
		YouAreHere:                          "您在这里",
		LcRewordNotSupported:                "当前不支持交互式重新基准化时的重新措词提交",
		LcCherryPickCopy:                    "复制提交（拣选）",
		LcCherryPickCopyRange:               "复制提交范围（拣选）",
		LcPasteCommits:                      "粘贴提交（拣选）",
		SureCherryPick:                      "您确定要将选中的提交进行拣选到这个分支吗？",
		CherryPick:                          "拣选 (Cherry-Pick)",
		Donate:                              "捐助",
		AskQuestion:                         "提问咨询",
		PrevLine:                            "选择上一行",
		NextLine:                            "选择下一行",
		PrevHunk:                            "选择上一个区块",
		NextHunk:                            "选择下一个区块",
		PrevConflict:                        "选择上一个冲突",
		NextConflict:                        "选择下一个冲突",
		SelectPrevHunk:                      "选择顶部块",
		SelectNextHunk:                      "选择底部块",
		ScrollDown:                          "向下滚动",
		ScrollUp:                            "向上滚动",
		LcScrollUpMainPanel:                 "向上滚动主面板",
		LcScrollDownMainPanel:               "向下滚动主面板",
		AmendCommitTitle:                    "修改提交",
		AmendCommitPrompt:                   "您确定要使用暂存文件来修改此提交吗？",
		DeleteCommitTitle:                   "删除提交",
		DeleteCommitPrompt:                  "您确定要删除此提交吗？",
		SquashingStatus:                     "正在压缩",
		FixingStatus:                        "正在修正",
		DeletingStatus:                      "正在删除",
		MovingStatus:                        "正在移动",
		RebasingStatus:                      "正在变基",
		AmendingStatus:                      "正在修改",
		CherryPickingStatus:                 "正在拣选",
		UndoingStatus:                       "正在撤销",
		RedoingStatus:                       "正在重做",
		CheckingOutStatus:                   "长子检出",
		CommittingStatus:                    "正在提交",
		CommitFiles:                         "提交文件",
		LcViewItemFiles:                     "查看提交的文件",
		CommitFilesTitle:                    "提交文件",
		LcCheckoutCommitFile:                "检出文件",
		LcDiscardOldFileChange:              "放弃对此文件的提交更改",
		DiscardFileChangesTitle:             "放弃文件更改",
		DiscardFileChangesPrompt:            "您确定要舍弃此提交对该文件的更改吗？如果此文件是在此提交中创建的，它将被删除",
		DisabledForGPG:                      "该功能不适用于使用 GPG 的用户",
		CreateRepo:                          "当前目录不在 git 仓库中。是否在此目录创建一个新的 git 仓库？(y/n): ",
		AutoStashTitle:                      "自动存储？",
		AutoStashPrompt:                     "您必须隐藏并弹出更改以使更改生效。自动执行？(enter/esc)",
		StashPrefix:                         "自动隐藏更改 ",
		LcViewDiscardOptions:                "查看'放弃更改'选项",
		LcCancel:                            "取消",
		LcDiscardAllChanges:                 "放弃所有更改",
		LcDiscardUnstagedChanges:            "放弃未暂存的变更",
		LcDiscardAllChangesToAllFiles:       "清空工作区",
		LcDiscardAnyUnstagedChanges:         "丢弃未暂存的变更",
		LcDiscardUntrackedFiles:             "丢弃未跟踪的文件",
		LcHardReset:                         "硬重置",
		LcViewResetOptions:                  `查看重置选项`,
		LcCreateFixupCommit:                 `为此提交创建修正`,
		LcSquashAboveCommits:                `压缩在所选提交之上的所有“fixup!”提交（自动压缩）`,
		SquashAboveCommits:                  `压缩在所选提交之上的所有“fixup!”提交（自动压缩）`,
		SureSquashAboveCommits:              `您确定要压缩在 {{.commit}} 之上的所有“fixup!”提交吗?`,
		CreateFixupCommit:                   `创建修正提交`,
		SureCreateFixupCommit:               `您确定要对 {{.commit}} 创建修正提交吗？`,
		LcExecuteCustomCommand:              "执行自定义命令",
		CustomCommand:                       "自定义命令：",
		LcCommitChangesWithoutHook:          "提交更改而无需预先提交钩子",
		SkipHookPrefixNotConfigured:         "您尚未配置用于跳过钩子的提交消息前缀。请在您的配置中设置 `git.skipHookPrefix ='WIP'`",
		LcResetTo:                           `重置为`,
		PressEnterToReturn:                  "按下 Enter 键返回 lazyaws",
		LcViewStashOptions:                  "查看贮藏选项",
		LcStashAllChanges:                   "将所有更改加入贮藏",
		LcStashAllChangesKeepIndex:          "将已暂存的更改加入贮藏",
		LcStashOptions:                      "贮藏选项",
		NotARepository:                      "错误：必须在 git 仓库中运行",
		LcJump:                              "跳到面板",
		LcScrollLeftRight:                   "左右滚动",
		LcScrollLeft:                        "向左滚动",
		LcScrollRight:                       "向右滚动",
		DiscardPatch:                        "丢弃补丁",
		DiscardPatchConfirm:                 "您一次只能通过一个提交或贮藏条目构建补丁。需要放弃当前补丁吗？",
		CantPatchWhileRebasingError:         "处于合并或变基状态时，您无法构建修补程序或运行修补程序命令",
		LcToggleAddToPatch:                  "补丁中包含的切换文件",
		ViewPatchOptions:                    "查看自定义补丁选项",
		PatchOptionsTitle:                   "补丁选项",
		NoPatchError:                        "尚未创建补丁。你可以在提交中的文件上按下“空格”或使用“回车”添加其中的特定行以开始构建补丁",
		LcEnterFile:                         "输入文件以将所选行添加到补丁中（或切换目录折叠）",
		ExitCustomPatchBuilder:              `退出逐行模式`,
		EnterUpstream:                       `以这种格式输入上游：'<远程仓库> <分支名称>'`,
		InvalidUpstream:                     "上游格式无效，格式应当为：'<remote> <branchname>'",
		ReturnToRemotesList:                 `返回远程仓库列表`,
		LcAddNewRemote:                      `添加新的远程仓库`,
		LcNewRemoteName:                     `新远程仓库名称:`,
		LcNewRemoteUrl:                      `新远程仓库 URL:`,
		LcEditRemoteName:                    `输入远程仓库 {{.remoteName}} 的新名称：`,
		LcEditRemoteUrl:                     `输入远程仓库 {{.remoteName}} 的新 URL：`,
		LcRemoveRemote:                      `删除远程`,
		LcRemoveRemotePrompt:                "您确定要删除远程仓库吗？",
		DeleteRemoteBranch:                  "删除远程分支",
		DeleteRemoteBranchMessage:           "您确定要删除远程分支吗？",
		LcSetUpstream:                       "设置为检出分支的上游",
		LcSetAsUpstream:                     "设置为检出分支的上游",
		SetUpstreamTitle:                    "设置上游分支",
		SetUpstreamMessage:                  "您确定要将 {{.checkedOut}} 的上游分支设置为 {{.selected}} 吗？",
		LcEditRemote:                        "编辑远程仓库",
		LcTagCommit:                         "标签提交",
		TagMenuTitle:                        "创建标签",
		TagNameTitle:                        "标签名称：",
		TagMessageTitle:                     "标签消息：",
		LcAnnotatedTag:                      "附注标签",
		LcLightweightTag:                    "轻量标签",
		LcDeleteTag:                         "删除标签",
		DeleteTagTitle:                      "删除标签",
		DeleteTagPrompt:                     "您确定要删除标签 {{.tagName}} 吗？",
		PushTagTitle:                        "将 {{.tagName}} 推送到远程仓库：",
		LcPushTag:                           "推送标签",
		LcCreateTag:                         "创建标签",
		CreateTagTitle:                      "标签名称：",
		LcFetchRemote:                       "抓取远程仓库",
		FetchingRemoteStatus:                "抓取远程仓库中",
		LcCheckoutCommit:                    "检出提交",
		SureCheckoutThisCommit:              "您确定要检出此提交吗?",
		LcGitFlowOptions:                    "显示 git-flow 选项",
		NotAGitFlowBranch:                   "这似乎不是 git flow 分支",
		NewGitFlowBranchPrompt:              "新的 {{.branchType}} 名称：",
		IgnoreTracked:                       "忽略跟踪文件",
		IgnoreTrackedPrompt:                 "您确定要忽略已跟踪的文件吗？",
		LcViewResetToUpstreamOptions:        "查看上游重置选项",
		LcNextScreenMode:                    "下一屏模式（正常/半屏/全屏）",
		LcPrevScreenMode:                    "上一屏模式",
		LcStartSearch:                       "开始搜索",
		Panel:                               "面板",
		Keybindings:                         "按键绑定",
		LcRenameBranch:                      "重命名分支",
		NewBranchNamePrompt:                 "输入分支的新名称",
		RenameBranchWarning:                 "该分支正在跟踪远程仓库。此操作将仅会重命名本地分支名称，而不会重命名远程分支的名称。确定继续？",
		LcOpenMenu:                          "打开菜单",
		LcResetCherryPick:                   "重置已拣选（复制）的提交",
		LcNextTab:                           "下一个标签",
		LcPrevTab:                           "上一个标签",
		LcCantUndoWhileRebasing:             "进行基础调整时无法撤消",
		LcCantRedoWhileRebasing:             "变基时无法重做",
		MustStashWarning:                    "将补丁拉出到索引中需要存储和取消存储所做的更改。如果出现问题，您将可以从存储中访问文件。继续？",
		MustStashTitle:                      "必须保存进度",
		ConfirmationTitle:                   "确认面板",
		LcPrevPage:                          "上一页",
		LcNextPage:                          "下一页",
		LcGotoTop:                           "滚动到顶部",
		LcGotoBottom:                        "滚动到底部",
		LcFilteringBy:                       "过滤依据",
		ResetInParentheses:                  "（重置）",
		LcOpenFilteringMenu:                 "查看按路径过滤选项",
		LcFilterBy:                          "过滤",
		LcExitFilterMode:                    "停止按路径过滤",
		LcFilterPathOption:                  "输入要过滤的路径",
		EnterFileName:                       "输入路径：",
		FilteringMenuTitle:                  "正在过滤",
		MustExitFilterModeTitle:             "命令不可用",
		MustExitFilterModePrompt:            "命令在过滤模式下不可用。退出过滤模式？",
		LcDiff:                              "差异",
		LcEnterRefToDiff:                    "输入 ref 以 diff",
		LcEnteRefName:                       "输入 ref：",
		LcExitDiffMode:                      "退出差异模式",
		DiffingMenuTitle:                    "正在 diff",
		LcSwapDiff:                          "反向 diff",
		LcOpenDiffingMenu:                   "打开 diff 菜单",
		// 实际视图 (actual view) 是附加视图 (extras view)，未来,我打算为附加视图提供更多选项卡，但现在，上面的文本只需要提及“命令日志”这个部分
		LcOpenExtrasMenu:                    "打开命令日志菜单",
		LcShowingGitDiff:                    "显示输出：",
		LcCopyCommitShaToClipboard:          "将提交的 SHA 复制到剪贴板",
		LcCopyCommitMessageToClipboard:      "将提交消息复制到剪贴板",
		LcCopyBranchNameToClipboard:         "将分支名称复制到剪贴板",
		LcCopyFileNameToClipboard:           "将文件名复制到剪贴板",
		LcCopyCommitFileNameToClipboard:     "将提交的文件名复制到剪贴板",
		LcCopySelectedTexToClipboard:        "将选中文本复制到剪贴板",
		LcCommitPrefixPatternError:          "提交前缀模式错误",
		NoFilesStagedTitle:                  "没有暂存文件",
		NoFilesStagedPrompt:                 "您尚未暂存任何文件。提交所有文件？",
		BranchNotFoundTitle:                 "找不到分支",
		BranchNotFoundPrompt:                "找不到分支。创建一个新分支命名为：",
		UnstageLinesTitle:                   "取消暂存选中的行",
		UnstageLinesPrompt:                  "您确定要删除所选的行（git reset）吗？这是不可逆的。\n要禁用此对话框，请将 'gui.skipUnstageLineWarning' 的配置键设置为 true",
		LcCreateNewBranchFromCommit:         "从提交创建新分支",
		LcBuildingPatch:                     "正在构建补丁",
		LcViewCommits:                       "查看提交",
		MinGitVersionError:                  "Git 版本必须至少为 2.20（即从 2018 年开始的版本）。请更新 git。或者在 https://github.com/BSteffaniak/lazyaws/issues 上提出一个问题，以使 lazyaws 更加向后兼容。",
		LcRunningCustomCommandStatus:        "正在运行自定义命令",
		LcSubmoduleStashAndReset:            "存放未提交的子模块更改和更新",
		LcAndResetSubmodules:                "和重置子模块",
		LcEnterSubmodule:                    "输入子模块",
		LcCopySubmoduleNameToClipboard:      "将子模块名称复制到剪贴板",
		RemoveSubmodule:                     "删除子模块",
		LcRemoveSubmodule:                   "删除子模块",
		RemoveSubmodulePrompt:               "您确定要删除子模块 '%s' 及其对应的目录吗？这是不可逆的。",
		LcResettingSubmoduleStatus:          "正在重置子模块",
		LcNewSubmoduleName:                  "新的子模块名称：",
		LcNewSubmoduleUrl:                   "新的子模块 URL：",
		LcNewSubmodulePath:                  "新的子模块路径：",
		LcAddSubmodule:                      "添加新的子模块",
		LcAddingSubmoduleStatus:             "添加子模块",
		LcUpdateSubmoduleUrl:                "更新子模块 '%s' 的 URL",
		LcUpdatingSubmoduleUrlStatus:        "更新 URL 中",
		LcEditSubmoduleUrl:                  "更新子模块 URL",
		LcInitializingSubmoduleStatus:       "正在初始化子模块",
		LcInitSubmodule:                     "初始化子模块",
		LcSubmoduleUpdate:                   "更新子模块",
		LcUpdatingSubmoduleStatus:           "正在更新子模块",
		LcBulkInitSubmodules:                "批量初始化子模块",
		LcBulkUpdateSubmodules:              "批量更新子模块",
		LcBulkDeinitSubmodules:              "批量反初始化子模块",
		LcViewBulkSubmoduleOptions:          "查看批量子模块选项",
		LcBulkSubmoduleOptions:              "批量子模块选项",
		LcRunningCommand:                    "运行命令",
		SubCommitsTitle:                     "子提交",
		SubmodulesTitle:                     "子模块",
		NavigationTitle:                     "列表面板导航",
		SuggestionsCheatsheetTitle:          "意见建议",
		SuggestionsTitle:                    "意见建议 (点击 %s 以聚焦)",
		ExtrasTitle:                         "附加",
		PushingTagStatus:                    "推送标签",
		PullRequestURLCopiedToClipboard:     "抓取请求网址已复制到剪贴板",
		CommitMessageCopiedToClipboard:      "提交消息复制到剪贴板",
		LcCopiedToClipboard:                 "复制到剪贴板",
		ErrCannotEditDirectory:              "无法编辑目录：您只能编辑单个文件",
		ErrStageDirWithInlineMergeConflicts: "无法 暂存/取消暂存 包含具有内联合并冲突的文件的目录。请先解决合并冲突",
		ErrRepositoryMovedOrDeleted:         "找不到仓库。它可能已被移动或删除 ¯\\_(ツ)_/¯",
		CommandLog:                          "命令日志",
		ToggleShowCommandLog:                "切换 显示/隐藏 命令日志",
		FocusCommandLog:                     "焦点命令日志",
		CommandLogHeader:                    "您可以通过按 '%s' 隐藏或集中显示该面板，或使用 `gui.showCommandLog: false`\n将其永久隐藏在您的配置中",
		RandomTip:                           "随机小提示",
		SelectParentCommitForMerge:          "选择父提交进行合并",
		ToggleWhitespaceInDiffView:          "切换是否在差异视图中显示空白字符差异",
		IgnoringWhitespaceInDiffView:        "将会在差异视图中忽略空格字符差异",
		ShowingWhitespaceInDiffView:         "将会在差异视图中显示空白字符差异",
		IncreaseContextInDiffView:           "扩大差异视图中显示的上下文范围",
		DecreaseContextInDiffView:           "缩小差异视图中显示的上下文范围",
		CreatePullRequest:                   "创建抓取请求",
		CreatePullRequestOptions:            "创建抓取请求选项",
		LcCreatePullRequestOptions:          "创建抓取请求选项",
		LcDefaultBranch:                     "默认分支",
		LcSelectBranch:                      "选择分支",
		SelectConfigFile:                    "选择配置文件",
		NoConfigFileFoundErr:                "找不到配置文件",
		LcLoadingFileSuggestions:            "正在加载文件建议",
		LcLoadingCommits:                    "正在加载提交",
		MustSpecifyOriginError:              "指定分支时，必须同时指定远程",
		GitOutput:                           "Git 输出：",
		GitCommandFailed:                    "Git 命令执行失败。查看命令日志了解详情 (使用 %s 打开)",
		AbortTitle:                          "放弃 %s",
		AbortPrompt:                         "您确定要放弃当前 %s 吗？",
		LcOpenLogMenu:                       "打开日志菜单",
		LogMenuTitle:                        "提交日志选项",
		ToggleShowGitGraphAll:               "切换显示完整 git 分支图 (向 `git log` 命令传入 `--all` 选项)",
		ShowGitGraph:                        "显示 git 分支图",
		SortCommits:                         "提交排序",
		CantChangeContextSizeError:          "无法在补丁构建模式下更改上下文，因为我们在发布该功能时懒得支持它。 如果你真的想要这么做，请告诉我们！",
		LcOpenCommitInBrowser:               "在浏览器中打开提交",
		LcViewBisectOptions:                 "查看二分查找选项",
		Actions: Actions{
			// TODO: combine this with the original keybinding descriptions (those are all in lowercase atm)
			CheckoutCommit:                    "检出提交",
			CheckoutTag:                       "检出标签",
			CheckoutBranch:                    "检出分支",
			ForceCheckoutBranch:               "强制检出分支",
			DeleteBranch:                      "删除分支",
			Merge:                             "合并",
			RebaseBranch:                      "变基分支",
			RenameBranch:                      "重命名分支",
			CreateBranch:                      "建立分支",
			CherryPick:                        "(拣选) 粘贴提交",
			CheckoutFile:                      "检出文件",
			DiscardOldFileChange:              "放弃旧文件更改",
			SquashCommitDown:                  "向下压缩提交",
			FixupCommit:                       "修正提交",
			RewordCommit:                      "改写提交",
			DropCommit:                        "删除提交",
			EditCommit:                        "编辑提交",
			AmendCommit:                       "修改提交",
			RevertCommit:                      "还原提交",
			CreateFixupCommit:                 "创建修正提交",
			SquashAllAboveFixupCommits:        "压缩以上所有的修正提交",
			CreateLightweightTag:              "创建轻量标签",
			CreateAnnotatedTag:                "创建附注标签",
			CopyCommitMessageToClipboard:      "将提交消息复制到剪贴板",
			MoveCommitUp:                      "上移提交",
			MoveCommitDown:                    "下移提交",
			CustomCommand:                     "自定义命令",
			DiscardAllChangesInDirectory:      "丢弃目录中的所有更改",
			DiscardUnstagedChangesInDirectory: "丢弃目录中未暂存的更改",
			DiscardAllChangesInFile:           "丢弃文件中的所有更改",
			DiscardAllUnstagedChangesInFile:   "丢弃文件中所有未暂存的更改",
			StageFile:                         "暂存文件",
			UnstageFile:                       "取消暂存文件",
			UnstageAllFiles:                   "取消暂存所有文件",
			StageAllFiles:                     "暂存所有文件",
			LcIgnoreExcludeFile:               "忽略文件",
			Commit:                            "提交 (Commit)",
			EditFile:                          "编辑文件",
			Push:                              "推送 (Push)",
			Pull:                              "拉取 (Pull)",
			OpenFile:                          "打开文件",
			StashAllChanges:                   "贮藏所有更改",
			StashStagedChanges:                "贮藏暂存的更改",
			GitFlowFinish:                     "Git flow 结果",
			GitFlowStart:                      "Git Flow 开始",
			CopyToClipboard:                   "复制到剪贴板",
			CopySelectedTextToClipboard:       "将选中文本复制到剪贴板",
			RemovePatchFromCommit:             "从提交中删除补丁",
			MovePatchToSelectedCommit:         "将补丁移动到选定的提交",
			MovePatchIntoIndex:                "将补丁移到索引",
			MovePatchIntoNewCommit:            "将补丁移到新提交中",
			DeleteRemoteBranch:                "删除远程分支",
			SetBranchUpstream:                 "设置分支上游",
			AddRemote:                         "添加远程",
			RemoveRemote:                      "移除远程",
			UpdateRemote:                      "更新远程",
			ApplyPatch:                        "应用补丁",
			Stash:                             "贮藏 (Stash)",
			RenameStash:                       "Rename stash",
			RemoveSubmodule:                   "删除子模块",
			ResetSubmodule:                    "重置子模块",
			AddSubmodule:                      "添加子模块",
			UpdateSubmoduleUrl:                "更新子模块 URL",
			InitialiseSubmodule:               "初始化子模块",
			BulkInitialiseSubmodules:          "批量初始化子模块",
			BulkUpdateSubmodules:              "批量更新子模块",
			BulkDeinitialiseSubmodules:        "批量取消初始化子模块",
			UpdateSubmodule:                   "更新子模块",
			DeleteTag:                         "删除标签",
			PushTag:                           "推送标签",
			NukeWorkingTree:                   "Nuke 工作树",
			DiscardUnstagedFileChanges:        "放弃未暂存的文件更改",
			RemoveUntrackedFiles:              "删除未跟踪的文件",
			SoftReset:                         "软重置",
			MixedReset:                        "混合重置",
			HardReset:                         "硬重置",
			FastForwardBranch:                 "快进分支",
			Undo:                              "撤销",
			Redo:                              "重做",
			CopyPullRequestURL:                "复制拉取请求 URL",
			OpenMergeTool:                     "打开合并工具",
			OpenCommitInBrowser:               "在浏览器中打开提交",
			OpenPullRequest:                   "在浏览器中打开拉取请求",
			StartBisect:                       "开始二分查找 (Bisect)",
			ResetBisect:                       "重置二分查找",
			BisectSkip:                        "二分查找跳过",
			BisectMark:                        "二分查找标记",
		},
		Bisect: Bisect{
			Mark:                        "将 %s 标记为 %s",
			MarkStart:                   "将 %s 标记为 %s (start bisect)",
			Skip:                        "跳过 %s",
			ResetTitle:                  "重置 'git bisect'",
			ResetPrompt:                 "您确定要重置 'git bisect' 吗？",
			ResetOption:                 "重置二分查找",
			BisectMenuTitle:             "二分查找",
			CompleteTitle:               "二分查找完成",
			CompletePrompt:              "二分查找完成！以下提交引入了此变更：\n\n%s\n\n您现在要重置 'git bisect' 吗？",
			CompletePromptIndeterminate: "二分查找完成！一些提交被跳过了，所以下列提交中的任何一个都可能引入了此变更：\n\n%s\n\n您现在要重置 'git bisect' 吗？",
		},
	}
}
