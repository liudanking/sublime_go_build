sublime_go_build
================

# How to use

1. 将/config/kill_go.sh文件保存到本地。
2. 将/build/Go.sublime-build文件部署到~/Library/Application Support/Sublime Text 3/Packages/User下。修改"cmd": ["sh", "**<your_path_to_kill_go.sh_file\>**", "$file_name", "$file_base_name"]. 
** 注：**<your_path_to_kill_go.sh_file\>受Sublime Text限制，目前只支持**绝对路径**。
3. 在Sublime Text的Preferences->Key Bindings - User中设置key binding. 绑定到alt+c参考设置如下：
> { "keys": ["alt+c"], "command": "build", "args": {"variant": "Stop"} }
