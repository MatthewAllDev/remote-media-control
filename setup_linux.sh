#!/bin/bash

if [ "$(id -u)" -ne 0 ]; then
  echo "This script must be run as root."
  exit 1
fi

install_dir="/opt/RemoteMediaControl"
service_name="remote_media_control.service"

case "$1" in
  install)
    if [ ! -f "./RemoteMediaControl" ]; then
      echo "Error: Binary file not found."
      exit 1
    fi

    if [ ! -d "./static" ]; then
      echo "Error: 'static' directory not found."
      exit 1
    fi

    read -p "Would you like to install the application as a service with auto-start? (y/n): " install_service

    mkdir -p "$install_dir"
    cp ./RemoteMediaControl "$install_dir/"
    cp -r ./static "$install_dir/"

    echo "$install_dir" >> /etc/environment
    echo "Added $install_dir to system-wide PATH."

    if [ "$install_service" == "y" ]; then
      cat > /etc/systemd/system/$service_name << EOF
[Unit]
Description=Remote Media Control Application
After=network.target

[Service]
ExecStart=$install_dir/RemoteMediaControl
WorkingDirectory=$install_dir
Restart=always
User=root
Group=root
Environment=PATH=/usr/bin:/usr/local/bin

[Install]
WantedBy=multi-user.target
EOF

      systemctl daemon-reload
      systemctl enable "$service_name"
      systemctl start "$service_name"
      echo "Service installed and started."
    fi

    if command -v ufw &> /dev/null; then
      ufw allow 80/tcp
      echo "Port 80 opened in UFW firewall."
    elif command -v firewall-cmd &> /dev/null; then
      firewall-cmd --zone=public --add-port=80/tcp --permanent
      firewall-cmd --reload
      echo "Port 80 opened in firewalld."
    elif command -v iptables &> /dev/null; then
      iptables -A INPUT -p tcp --dport 80 -j ACCEPT
      echo "Port 80 opened in iptables."
    else
      echo "Firewall tool not found, unable to open port 80."
    fi

    echo "Installation completed successfully."
    ;;
  
  uninstall)
    echo "Stopping and disabling service if installed..."
    systemctl stop "$service_name" 2>/dev/null || true
    systemctl disable "$service_name" 2>/dev/null || true
    rm -f "/etc/systemd/system/$service_name"
    systemctl daemon-reload

    echo "Removing $install_dir from system-wide PATH."
    sed -i "/$install_dir/d" /etc/environment

    echo "Removing application files..."
    rm -rf "$install_dir"

    echo "Removing firewall rules..."

    if command -v ufw &> /dev/null; then
      ufw delete allow 80/tcp
      echo "Port 80 removed from UFW firewall."
    elif command -v firewall-cmd &> /dev/null; then
      firewall-cmd --zone=public --remove-port=80/tcp --permanent
      firewall-cmd --reload
      echo "Port 80 removed from firewalld."
    elif command -v iptables &> /dev/null; then
      iptables -D INPUT -p tcp --dport 80 -j ACCEPT
      echo "Port 80 removed from iptables."
    else
      echo "Firewall tool not found, unable to remove port 80."
    fi

    echo "Uninstallation completed."
    ;;
  
  *)
    echo "Usage: $0 {install|uninstall}"
    exit 1
    ;;
esac
