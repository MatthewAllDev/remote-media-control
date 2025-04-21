#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -L "$SOURCE" ]; do
  DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"
  SOURCE="$(readlink "$SOURCE")"
  [[ "$SOURCE" != /* ]] && SOURCE="$DIR/$SOURCE"
done
install_dir="$(cd -P "$(dirname "$SOURCE")" && pwd)"
service_name="remote_media_control.service"

show_help() {
  echo "Usage: $0 [command]"
  echo
  echo "Available commands:"
  echo "  install-service   Install RemoteMediaControl as a systemd service"
  echo "  enable-service    Enable the service to start on boot"
  echo "  disable-service   Disable the service from starting on boot"
  echo "  start-service     Start the systemd service"
  echo "  stop-service      Stop the systemd service"
  echo "  uninstall         Uninstall RemoteMediaControl (requires sudo)"
  echo "  help              Show this help message"
  echo
  echo "Without arguments, starts RemoteMediaControl with sudo"
}

check_service_installed() {
  if ! systemctl list-unit-files | grep -q "^$service_name"; then
    echo "Error: $service_name service is not installed."
    echo "Hint: Run '$(basename "$0") install-service' to install it."
    exit 1
  fi
}

case "$1" in
  uninstall)
    echo "Superuser privileges are required to uninstall RemoteMediaControl"
    sudo "$install_dir/setup.sh" uninstall
    exit $?
    ;;
  install-service)
    echo "Installing RemoteMediaControl as a systemd service"
    sudo "$install_dir/setup.sh" install-service
    exit $?
    ;;
  enable-service)
    check_service_installed
    echo "Enabling the RemoteMediaControl service"
    sudo systemctl enable $service_name
    exit $?
    ;;
  disable-service)
    check_service_installed
    echo "Disabling the RemoteMediaControl service"
    sudo systemctl disable $service_name
    exit $?
    ;;
  start-service)
    check_service_installed
    echo "Starting the RemoteMediaControl service"
    sudo systemctl start $service_name
    exit $?
    ;;
  stop-service)
    check_service_installed
    echo "Stopping the RemoteMediaControl service"
    sudo systemctl stop $service_name
    exit $?
    ;;
  help)
    show_help
    exit 0
    ;;
esac

echo "Superuser privileges are required to start RemoteMediaControl"
sudo "$install_dir/RemoteMediaControl" --static-path "$install_dir/static" "$@"
