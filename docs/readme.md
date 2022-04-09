# Wiki CLI [![Go](https://github.com/BetaPictoris/wcli/actions/workflows/go.yml/badge.svg)](https://github.com/BetaPictoris/wcli/actions/workflows/go.yml)
View Wikipedia articles through the CLI

## Installation
### From release
```bash
curl -LO https://github.com/BetaPictoris/wiki/releases/latest/download/wcli    # Download the latest binary.
sudo install -Dt /usr/local/bin -m 755 wiki                                    # Install Wiki CLI to "/usr/local/bin" with the mode "755"
```

### Build from source 

#### Dependencies

You need go installed to build this program. You can install it from your distro's repository using one of the following commands:

```bash
# Arch/Manjaro (and derivatives)
sudo pacman -S go

# Debian/Ubuntu (and derivatives)
sudo apt install golang-go
```

Alternatively, you can install it from go's official website: https://go.dev/doc/install

```bash
git clone git@github.com:BetaPictoris/wcli.git      # Clone the repository
cd wiki                                             # Change into the repository's directory
bash build.sh                                       # Build Wiki CLI
sudo install -Dt /usr/local/bin -m 755 wiki         # Install Wiki CLI to "/usr/local/bin" with the mode "755"
```

### User install
If you don't have access to `sudo` on your system you can install to your user's `~/.local/bin` directory with this command: 
```bash
install -Dt ~/.local/bin -m 755 wiki
```
