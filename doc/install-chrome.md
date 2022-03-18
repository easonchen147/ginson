## 安装chromedp 依赖环境

### Ubuntu 安装步骤

```shell

# 下载源加入软件列表
wget http://www.linuxidc.com/files/repo/google-chrome.list -P /etc/apt/sources.list.d/

# 导入google软件公钥
wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add -

# 更新列表
apt-get update

# 安装chrome
apt-get install google-chrome-stable

# 查看版本
google-chrome --version

```

### CentOs7 安装步骤

```shell

# 进入源管理文件夹
cd /ect/yum.repos.d/

# 编辑chrome源
vim google-chrome.repo

# 保存到google-chrome.repo
[google-chrome]
name=google-chrome
baseurl=http://dl.google.com/linux/chrome/rpm/stable/$basearch
enabled=1
gpgcheck=1
gpgkey=https://dl-ssl.google.com/linux/linux_signing_key.pub

# 安装chrome
yum -y install google-chrome-stable --nogpgcheck

# 查看版本
google-chrome --version

```