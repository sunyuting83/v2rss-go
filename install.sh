#! /bin/bash

success() {
    echo -e "\033[42;37m 成功 \033[0m $1"
}

start_step() {
  if [ -n "$STEP" ]; then success "$STEP"; fi
  STEP="$1"
  echo -e "\033[34m------------------------------------------------------------------\033[0m"
  echo -e "\033[34m$STEP\033[0m"
  echo -e "\033[34m------------------------------------------------------------------\033[0m"
}

tmp_dir="/tmp/v2rss"
v2rss_file="$tmp_dir/v2rss-linux-x86_64.tar.gz"
v2rss_download="https://github.com/sunyuting83/v2rss-go/releases/download/v1.0.2/v2rss-linux-x86_64.tar.gz"

read -p "Please enter root password: " -s rootPassword
start_step '1. 安装zip'
echo "$rootPassword" | sudo apt install -y zip unzip

start_step '2. 下载v2rss'

mkdir -p $tmp_dir
if [ ! -f "$v2rss_file" ]; then
  echo "================================="
  echo "[注意]需要下载v2rss.请耐心等待下载完成"
  echo "$v2rss_download"
  echo "================================="
  wget "$v2rss_download" -O "$v2rss_file"
fi

tar -xf "$v2rss_file" -C $tmp_dir
success '解压v2rss'

start_step '3. 安装v2rss服务'
start_dir="/etc/v2rss"
if [ ! -d $start_dir ]; then
  echo "$rootPassword" | sudo -S mkdir -p -m 755 $start_dir
fi

echo "$rootPassword" | sudo -S install -Dm755 $tmp_dir/v2rss /usr/bin/v2rss &&
echo "$rootPassword" | sudo -S install -Dm644 $tmp_dir/v2rss.service /usr/lib/systemd/system/v2rss.service &&
echo "$rootPassword" | sudo -S install -Dm777 $tmp_dir/v2rss.sh /etc/v2rss/v2rss.sh &&
echo "$rootPassword" | sudo -S systemctl enable v2rss &&
echo "$rootPassword" | sudo -S systemctl daemon-reload &&
echo "$rootPassword" | sudo -S systemctl start v2rss &&

start_step '4. 服务安装成功。端口：5500'

success "v2rss success"