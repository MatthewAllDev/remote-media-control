# Remote Media Control

## Description
**Remote Media Control** is an application for remotely controlling media playback on your PC. It supports Windows, Linux, and potentially macOS. Key features include:
- Media playback control using standard media keys.
- VLC mode, which uses VLC media player's hotkeys for control.

## Features
- **Cross-platform**: Works on Windows and Linux (macOS not tested).
- **Media Keys Mode**: Control playback using standard commands.
- **VLC Mode**: Control playback with VLC hotkeys.
- **Runs as a service**: The application can be installed and run as a system service.

## Installation
### Step 1: Download the Application
1. Go to the [Releases](https://github.com/MatthewAllDev/remote-media-control/releases) page or the appropriate hosting service where the application is stored.
2. Download the archive that matches your operating system:
   - For Windows: `remote-media-control-windows.zip`
   - For Linux: `remote-media-control-linux.tar.gz`
   - For macOS (if available): `remote-media-control-macos.tar.gz`
3. Extract the contents of the archive to a directory of your choice.

### Step 2: Run the Installation Script
#### Windows
1. Open a terminal with administrator privileges.
2. Run the installation script:
   ```powershell
   setup_windows.ps1 install
   ```

3. To uninstall, run:
   ```powershell
   setup_windows.ps1 uninstall
   ```

#### Linux
1. Open a terminal and switch to the root user (or use `sudo`).
2. Run the installation script:
   ```bash
   sudo ./setup_linux.sh install
   ```

3. To uninstall, run:
   ```bash
   sudo ./setup_linux.sh uninstall
   ```

#### macOS
The application may work on macOS, but it has not been tested.

## Usage
Once installed, the application will run as a service. To start using it:
1. Open a browser on any device (your PC, smartphone, tablet, etc.).
2. Navigate to `http://<your-IP>` (port `80` is used by default, so no port input is required).
3. Use the web interface to control media playback on your PC:
   - Standard playback commands (Play, Pause, Stop, Next, Previous).
   - VLC mode allows control using VLC hotkeys.

### Alternative Manual Start
If you do not want to install the application as a service, you can start it manually with the following command:
```bash
remote-media-control --port=8080 --debug
```

#### Command-Line Parameters
- `--port` (default: `80`): Specifies the port to run the server on.
- `--debug` (default: `false`): Enables debug mode.

Example:
```bash
remote-media-control --port=8080 --debug
```

Once started, the server will be accessible at `http://<your-IP>:<port>`.

## Contribution
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
3. Make your changes and commit:
   ```bash
   git commit -m "Added new feature"
   ```
4. Push your changes to GitHub:
   ```bash
   git push origin feature-name
   ```
5. Create a Pull Request.

## License
This project is distributed under the MIT License. See the [LICENSE](./LICENSE) file for details.