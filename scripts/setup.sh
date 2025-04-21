#!/bin/bash

install_dir="/opt/RemoteMediaControl"
binary_name="RemoteMediaControl"
entrypoint_script="$install_dir/start.sh"
bin_link_remotemediacontrol="/usr/local/bin/remotemediacontrol"
bin_link_rmc="/usr/local/bin/rmc"
service_name="remote_media_control.service"
icon_path="$install_dir/app.svg"

check_root() {
  if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root."
    exit 1
  fi
}

copy_files() {
  if [ ! -f "./RemoteMediaControl" ]; then
    echo "Error: Binary file not found."
    exit 1
  fi
  if [ ! -d "./static" ]; then
    echo "Error: 'static' directory not found."
    exit 1
  fi
  mkdir -p "$install_dir"
  cp ./RemoteMediaControl "$install_dir/$binary_name"
  cp ./app.svg "$icon_path"
  cp -r ./static "$install_dir/"
  cp ./setup.sh "$install_dir/"
}

create_symbolic_links() {
  if [ ! -L "$bin_link_remotemediacontrol" ]; then
    ln -s "$entrypoint_script" "$bin_link_remotemediacontrol"
    echo "Symbolic link 'remotemediacontrol' created in /usr/local/bin."
  else
    if [ "$(readlink -f "$bin_link_remotemediacontrol")" != "$entrypoint_script" ]; then
        echo "Error: symbolic link 'remotemediacontrol' exists but does not point to the correct binary."
    fi
  fi
  if [ ! -L "$bin_link_rmc" ]; then
    ln -s "$entrypoint_script" "$bin_link_rmc"
    echo "Symbolic link 'rmc' created in /usr/local/bin."
  else
    if [ "$(readlink -f "$bin_link_rmc")" != "$entrypoint_script" ]; then
        echo "Error: symbolic link 'rmc' exists but does not point to the correct binary."
    fi
  fi
}

remove_symbolic_links() {
  echo "Removing symbolic links..."
  if [ -L "$bin_link_remotemediacontrol" ] && [ "$(readlink -f "$bin_link_remotemediacontrol")" == "$entrypoint_script" ]; then
    rm "$bin_link_remotemediacontrol"
  fi

  if [ -L "$bin_link_rmc" ] && [ "$(readlink -f "$bin_link_rmc")" == "$entrypoint_script" ]; then
    rm "$bin_link_rmc"
  fi
}

create_start_script() {
  cat > "$install_dir/start.sh" <<EOL
#!/bin/bash

if [ "\$1" == "uninstall" ]; then
  echo "Superuser privileges are required to uninstall RemoteMediaControl"
  sudo "$install_dir/setup.sh" uninstall
  exit \$?
fi

echo "Superuser privileges are required to start RemoteMediaControl"
sudo "$install_dir/RemoteMediaControl" --static-path "$install_dir/static" "\$@"
EOL
  chmod +x "$install_dir/start.sh"
  echo "Start script created at $install_dir/start.sh"
}

install_service() {
  if [ ! -f "$install_dir/RemoteMediaControl" ]; then
    echo "Error: Binary file not found in $install_dir."
    exit 1
  fi
  cat > /etc/systemd/system/$service_name << EOF
[Unit]
Description=Remote Media Control Application
After=network.target

[Service]
ExecStart=$install_dir/$binary_name
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
}

create_desktop_file() {
  desktop_file="/usr/share/applications/remote-media-control.desktop"
  cat > "$desktop_file" << EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=Remote Media Control
Exec=bash -c "$entrypoint_script"
Icon=$icon_path
Terminal=true
StartupNotify=true
StartupWMClass=RemoteMediaControl
Categories=Utility;AudioVideo;Network;
EOF
  echo "Application shortcut added to system menu."
}

add_firewall_rules() {
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
}

remove_firewall_rules() {
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
}

install() {
  copy_files
  create_start_script
  read -p "Would you like to install the application as a service with auto-start? (y/n): " install_service
  create_symbolic_links

  if [ "$install_service" == "y" ]; then
    install_service
  else
    create_desktop_file
  fi
  add_firewall_rules
  echo "Installation completed successfully."
}

uninstall() {
  echo "Stopping and disabling service if installed..."
    systemctl stop "$service_name" 2>/dev/null || true
    systemctl disable "$service_name" 2>/dev/null || true
    rm -f "/etc/systemd/system/$service_name"
    systemctl daemon-reload

    echo "Removing application files..."
    rm -rf "$install_dir"

    remove_firewall_rules

    remove_symbolic_links

    echo "Removing application shortcut..."
    rm -f /usr/share/applications/remote-media-control.desktop

    echo "Uninstallation completed."
  
}

check_root

case "$1" in
  install)
  install
    ;;
  
  uninstall)
  uninstall
    ;;
  install-service)
  install_service
    ;;  
  *)
    echo "Usage: $0 {install|uninstall|install-service}"
    exit 1
esac
