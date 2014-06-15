sublime_go_build
================

# How to use

## shell版本配置
1. 将/config/kill_go.sh文件保存到本地。
2. 将/build/Go.sublime-build文件部署到~/Library/Application Support/Sublime Text 3/Packages/User下。修改"cmd": ["sh", "**<your_path_to_kill_go.sh_file\>**", "$file_name", "$file_base_name"]. 
** 注：**受Sublime Text限制，目前**<your_path_to_kill_go.sh_file\>**只支持**绝对路径**。
3. 在Sublime Text的Preferences->Key Bindings - User中设置key binding. 绑定到alt+c参考设置如下：
		`{ "keys": ["alt+c"], "command": "build", "args": {"variant": "Stop-sh"} }`

## Go版本配置
1. 将go build kill_go.go，得到kill_go可执行文件；
2. 将/build/Go.sublime-build文件部署到~/Library/Application Support/Sublime Text 3/Packages/User下。修改"cmd": ["sh", "-c", "**<your_path_to_kill_go>** $file_name exe/$file_base_name"];
** 注：**受Sublime Text限制，目前**<your_path_to_kill_go\>**只支持**绝对路径**。
3. 在Sublime Text的Preferences->Key Bindings - User中设置key binding. 绑定到shift+alt+c参考设置如下：
		` { "keys": ["ctrl+alt+c"], "command": "build", "args": {"variant": "Stop-Go"} }`
