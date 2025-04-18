@echo off
NET SESSION >nul 2>nul
IF %ERRORLEVEL% NEQ 0 (
    echo This script must be run as Administrator.
    echo Restarting with elevated privileges...
    start "" powershell -Command "Start-Process cmd -ArgumentList '/c \"%~f0\"' -Verb RunAs"
    exit /b
)

set "APP_NAME=RemoteMediaControl"
set "install_dir=C:\Program Files\RemoteMediaControl"
set "START_MENU_FOLDER=%ProgramData%\Microsoft\Windows\Start Menu\Programs\RemoteMediaControl"

tasklist /FI "IMAGENAME eq RemoteMediaControl.exe" | find /I "RemoteMediaControl.exe" >nul
if %ERRORLEVEL% == 0 (
    echo Closing RemoteMediaControl.exe...
    taskkill /IM "RemoteMediaControl.exe" /F
    timeout /t 2 /nobreak >nul
)

if exist "%install_dir%" (
    if /I not "%~dp0" == "%install_dir%\" (
        echo Removing program files...
        rmdir /S /Q "%install_dir%"
    ) else (
        echo Skipping self-deletion - will be removed on reboot...
        reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnce" /v "Delete_%APP_NAME%" /d "cmd /c rmdir /S /Q \"%install_dir%\"" /f
    )
)

if exist "%START_MENU_FOLDER%" (
    echo Removing Start Menu shortcuts...
    del "%START_MENU_FOLDER%\%APP_NAME%.lnk" 2>nul
    del "%START_MENU_FOLDER%\Uninstall %APP_NAME%.lnk" 2>nul
    rmdir /Q "%START_MENU_FOLDER%" 2>nul
)

reg delete "HKCU\Software\Microsoft\Windows\CurrentVersion\Run" /v "%APP_NAME%" /f 2>nul

powershell -Command "Remove-NetFirewallRule -DisplayName '%APP_NAME%_Firewall_Rule' -ErrorAction SilentlyContinue"

echo Uninstallation completed successfully.
if exist "%install_dir%" (
    echo Note: Some files will be removed after system reboot.
)
pause