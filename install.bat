@echo off
NET SESSION >nul 2>nul
IF %ERRORLEVEL% NEQ 0 (
    echo This script must be run as Administrator.
    echo Restarting with elevated privileges...
    start "" powershell -Command "Start-Process cmd -ArgumentList '/c \"%~dp0%~nx0\"' -Verb RunAs"
    exit /b
)

powershell -Command "Set-ExecutionPolicy Bypass -Scope Process -Force"

IF NOT EXIST "%~dp0RemoteMediaControl.exe" (
    echo Error: Binary file not found.
    exit /b
)

IF NOT EXIST "%~dp0static" (
    echo Error: 'static' directory not found.
    exit /b
)

IF NOT EXIST "%~dp0uninstall.bat" (
    echo Error: 'uninstall.bat' file not found.
    exit /b
)

set "APP_NAME=RemoteMediaControl"
set "install_dir=C:\Program Files\RemoteMediaControl"
set "EXE_PATH=%install_dir%\RemoteMediaControl.exe"
set "START_MENU_FOLDER=%ProgramData%\Microsoft\Windows\Start Menu\Programs\%APP_NAME%"

mkdir "%install_dir%"
copy /Y "%~dp0RemoteMediaControl.exe" "%install_dir%"
xcopy /E /I /Y "%~dp0static" "%install_dir%"
copy /Y "%~dp0uninstall.bat" "%install_dir%"

if not exist "%START_MENU_FOLDER%" mkdir "%START_MENU_FOLDER%"
powershell -command "$WshShell = New-Object -ComObject WScript.Shell; $Shortcut = $WshShell.CreateShortcut('%START_MENU_FOLDER%\%APP_NAME%.lnk'); $Shortcut.TargetPath = '%EXE_PATH%'; $Shortcut.WorkingDirectory = '%install_dir%'; $Shortcut.IconLocation = '%EXE_PATH%,0'; $Shortcut.Save()"
powershell -command "$WshShell = New-Object -ComObject WScript.Shell; $Shortcut = $WshShell.CreateShortcut('%START_MENU_FOLDER%\Uninstall %APP_NAME%.lnk'); $Shortcut.TargetPath = '%~dp0uninstall.bat'; $Shortcut.WorkingDirectory = '%~dp0'; $Shortcut.IconLocation = '%SystemRoot%\System32\shell32.dll,31'; $Shortcut.Save()"

reg add "HKCU\Software\Microsoft\Windows\CurrentVersion\Run" /v "RemoteMediaControl" /t REG_SZ /d "%install_dir%\RemoteMediaControl.exe --static-path %install_dir%\static" /f

powershell -Command "New-NetFirewallRule -DisplayName 'RemoteMediaControl_Firewall_Rule' -Direction Inbound -Protocol TCP -LocalPort 80 -Action Allow -Profile Any"
echo Port 80 opened in Windows Firewall.

echo Installation completed successfully.
pause
