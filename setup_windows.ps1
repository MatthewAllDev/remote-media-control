if (-not ([Security.Principal.WindowsPrincipal[System.Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator))) {
    Write-Host "This script must be run as Administrator."
    exit
}

$install_dir = "C:\Program Files\RemoteMediaControl"
$service_name = "RemoteMediaControl"
$firewall_rule = "RemoteMediaControl_Firewall_Rule"

if ($args[0] -eq "install") {
    if (-not (Test-Path ".\RemoteMediaControl.exe")) {
        Write-Host "Error: Binary file not found."
        exit
    }

    if (-not (Test-Path ".\static")) {
        Write-Host "Error: 'static' directory not found."
        exit
    }

    $install_service = Read-Host "Would you like to install the application as a service with auto-start? (y/n)"

    New-Item -ItemType Directory -Force -Path $install_dir

    Copy-Item .\RemoteMediaControl.exe -Destination $install_dir
    Copy-Item .\static -Destination $install_dir -Recurse

    if ($install_service -eq "y") {
        New-Service -Name $service_name -Binary "C:\Windows\System32\cmd.exe" -Argument "/C `"$install_dir\RemoteMediaControl.exe`"" -StartupType Automatic
        Start-Service -Name $service_name
        Write-Host "Service installed and started."
    }

    New-NetFirewallRule -DisplayName $firewall_rule -Direction Inbound -Protocol TCP -LocalPort 80 -Action Allow -Profile Any
    Write-Host "Port 80 opened in Windows Firewall."

    Write-Host "Installation completed successfully."
}

elseif ($args[0] -eq "uninstall") {
    Write-Host "Stopping and removing service..."
    Stop-Service -Name $service_name -ErrorAction SilentlyContinue
    Remove-Service -Name $service_name -ErrorAction SilentlyContinue

    Write-Host "Removing application files..."
    Remove-Item -Path $install_dir -Recurse -Force

    Write-Host "Removing firewall rule..."
    Remove-NetFirewallRule -DisplayName $firewall_rule

    Write-Host "Uninstallation completed."
}

else {
    Write-Host "Usage: .\setup_windows.ps1 {install|uninstall}"
    exit
}
