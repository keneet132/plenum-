# Plenum


<div align="center">
  <img src="https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version" />
  <img src="https://img.shields.io/badge/Platform-Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black" alt="Platform" />
  <img src="https://img.shields.io/badge/License-MIT-blue?style=for-the-badge" alt="MIT License" />

</div>


  <br>
  <div>
    <b>Plenum</b> is a lightweight, zero-dependency peer-to-peer virtual LAN bridge written in Go. It establishes direct connections between machines behind restrictive CGNAT firewalls using <a href="https://github.com/WireGuard/wireguard-go">wireguard-go</a> and <a href="https://github.com/pion/stun">pion/stun</a>, without requiring an external VPS or central relay server.
  </div>
  
currently plenum only supports 1 to 1 connection.

  <br>

</div>

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

---

## Installation

### Prerequisites

- Go 1.21 or higher
- Linux OS with `iproute2` installed
- Windows:`wintun.dll` (from [wintun.net](https://www.wintun.net/)) placed in the same directory as the compiled binary

### Pre-built Binaries (Windows)
Download `plenum-windows-amd64.zip` from the [Releases](https://github.com/keneet132/plenum-/releases) page and extract it. Ensure `plenum.exe` and `wintun.dll` remain in the same folder.

### Build from source

```bash
git clone https://github.com/keneet132/plenum-.git
cd plenum-
go build -o plenum .
```

## Usage

since plenum creates a virtual TUN interface, all commands must be run with `sudo`and administrator powershell on windows

### Start the host

The host starts first and listens for incoming peer handshakes.
Linux:
```bash
sudo ./plenum --mode host 
```
Windos: 
```bash
.\plenum.exe --mode host
```
The host will output its public IP, public port, and Base64-encoded WireGuard public key.

Share these details securely with your peer.

### Connect the client

Once you have the host's endpoint, connect from the client machine:

Linux:
```bash
sudo ./plenum --mode client --peer <HOST_IP>:<HOST_PORT>
```

Windows:
```bash
.\plenum.exe --mode client --peer <HOST_IP>:<HOST_PORT>
```

When prompted with `Enter Host's Public Key:`, paste the host's Base64-encoded public key.

This triggers the UDP hole punching process and cryptographic handshake.

### Assign tunnel IP addresses

To enable packet routing across the bridge, assign local IP addresses to the `plenum0` interface on both machines.

**On the host machine:**
Linux:
```bash
sudo ip addr add 10.0.0.1/24 dev plenum0
sudo ip link set dev plenum0 up
```
Windows:
```bash
New-NetIPAddress -InterfaceAlias "plenum0" -IPAddress 10.0.0.1 -PrefixLength 24
```

**On the client machine:**
Linux:
```bash
sudo ip addr add 10.0.0.2/24 dev plenum0
sudo ip link set dev plenum0 up
```
Windows:
```bash
New-NetIPAddress -InterfaceAlias "plenum0" -IPAddress 10.0.0.2 -PrefixLength 24
```
### Verify connectivity

Test the tunnel using `ping`.

**From the host:**

```bash
ping 10.0.0.2
```

**From the client:**

```bash
ping 10.0.0.1
```

## Contributing

Contributions, issues, and feature requests are welcome!

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
