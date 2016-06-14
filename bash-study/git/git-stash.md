### git stash 命令使用说明

#### 命令说明
  stash -- 储藏可以获取你工作目录中的中间状态--也就是你修改过的被追踪的文件和暂存的变更--并将它保存到一个未完结变更的堆栈中, 随时可以重新应用

#### SYNOPSIS(概要)
```
    git stash list [<options>]
    git stash show [<stash>]
    git stash drop [-q|--quiet] [<stash>]
    git stash ( pop | apply ) [--index] [-q|--quiet] [<stash>]
    git stash branch <branchname> [<stash>]
    git stash [save [-p|--patch] [-k|--[no-]keep-index] [-q|--quiet]
             [-u|--include-untracked] [-a|--all] [<message>]]
    git stash clear
    git stash create [<message>]
    git stash store [-m|--message <message>] [-q|--quiet] <commit>
```

#### OPTIONS
##### save [-p|--patch] [-k|--[no-]keep-index] [-u|--include-untracked] [-a|--all] [-q|--quiet] [<message>]
    保存本地修改到一个新的stash中, 可以使用 * git reset --hard * 来恢复它们. *<message>* 部分是可选的, 用来给 stash 状态添加描述, 对于快速生成一个 stash 可以同时省略 *save* 和 *<message>*, 但是如果只添加 *<message>* 则会当成是一个拼写错误而忽略操作.
    使用-k或者--keep-index参数，在保存进度后不会将暂存区重置。默认会将暂存区和工作区强制重置。
    --include-untracked 和 -u 选项可以让Git储藏任何创建的未跟踪文件. (默认情况下，git stash 只会储藏已经在索引中的文件)
    --all 参数可以将.gitignore 中的文件也添加到储藏中
    --patch 标记，Git 不会储藏所有修改过的任何东西，但是会交互式地提示哪些改动想要储藏、哪些改动需要保存在工作目录中
    
