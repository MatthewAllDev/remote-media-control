@echo off
NET SESSION >nul 2>nul
IF %ERRORLEVEL% NEQ 0 (
    echo This script must be run as Administrator.
    echo Restarting with elevated privileges...
    start "" powershell -Command "Start-Process cmd -ArgumentList '/c \"%~dp0%~nx0\"' -Verb RunAs"
    exit /b
)

set "APP_NAME=RemoteMediaControl"
set "install_dir=C:\Program Files\RemoteMediaControl"
set "START_MENU_FOLDER=%ProgramData%\Microsoft\Windows\Start Menu\Programs\RemoteMediaControl"

if exist "%install_dir%" (
    rmdir /S /Q "%install_dir%"
)

if exist "%START_MENU_FOLDER%" (
    echo Removing shortcuts from Start Menu...
    del "%START_MENU_FOLDER%\%APP_NAME%.lnk"
    del "%START_MENU_FOLDER%\Uninstall %APP_NAME%.lnk"
    rmdir /Q "%START_MENU_FOLDER%"
)

reg delete "HKCU\Software\Microsoft\Windows\CurrentVersion\Run" /v "RemoteMediaControl" /f

powershell -Command "Remove-NetFirewallRule -DisplayName 'RemoteMediaControl_Firewall_Rule'"

echo Uninstallation completed successfully.
pause
