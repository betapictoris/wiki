# Wiki CLI [![Go](https://github.com/BetaPictoris/wcli/actions/workflows/go.yml/badge.svg)](https://github.com/BetaPictoris/wcli/actions/workflows/go.yml)
View Wikipedia articles through the CLI

![wiki-md](https://user-images.githubusercontent.com/65696362/162582942-8d0c2f14-c689-4438-94fb-cbe482090c51.png)


## Installation
### From release
```bash
curl -LO https://github.com/BetaPictoris/wiki/releases/latest/download/wiki    # Download the latest binary.
sudo install -Dt /usr/local/bin -m 755 wiki                                    # Install Wiki CLI to "/usr/local/bin" with the mode "755"
```

### Build from source 

#### Dependencies

You need Go installed to build this program. You can install it from your distro's repository using one of the following commands:

```bash
# Arch/Manjaro (and derivatives)
sudo pacman -S go

# Debian/Ubuntu (and derivatives)
sudo apt install golang-go
```

Alternatively, you can install it from go's official website: https://go.dev/doc/install

```bash
git clone git@github.com:BetaPictoris/wiki.git      # Clone the repository
cd wiki                                             # Change into the repository's directory
make                                                # Build Wiki CLI
sudo make install                                   # Install Wiki CLI to "/usr/local/bin" with the mode "755"
```

### User install
If you don't have access to `sudo` on your system you can install to your user's `~/.local/bin` directory with this command: 
```bash
install -Dt ~/.local/bin -m 755 wiki
```
