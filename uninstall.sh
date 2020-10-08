#!/bin/sh
success() {
  echo -e "\033[42;37m 成功 \033[0m $1"
}
# uninstall Kodi Live
read -p "Please enter root password: " -s rootPassword
echo "Start..."
echo "$rootPassword" | sudo -S systemctl stop v2rss &&
echo "$rootPassword" | sudo -S systemctl disable v2rss &&
echo "$rootPassword" | sudo -S systemctl daemon-reload &&
echo "$rootPassword" | sudo -S rm /usr/bin/v2rss &&
echo "$rootPassword" | sudo -S rm /usr/lib/systemd/system/v2rss.service &&
echo "$rootPassword" | sudo -S rm -rf /etc/v2rss &&
success "uninstall complete"