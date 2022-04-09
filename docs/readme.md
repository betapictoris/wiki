# Wiki CLI 
View Wikipedia articles through the CLI

## Installation
### From release
```bash
curl -LO https://github.com/BetaPictoris/wiki/releases/latest/download/wcli    # Download the latest binary.
sudo install -Dt /usr/local/bin -m 755 wiki                                    # Install Wiki CLI to "/usr/local/bin" with the mode "755"
```

### Build from source 
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
