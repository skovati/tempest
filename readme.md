# tmpst

tmpst (pronounced tempest) is a lightweight terminal weather application written in Go. It uses the reliable [US Gov't National Weather Service's API](https://weather-gov.github.io/api/).

## installation
clone repo, run tests, and compile
```sh
git clone https://github.com/skovati/tmpst
cd tmpst
make test && make install
```

## usage
```sh
tmpst [lat] [long] # provide location
tmpst # no location will try to read from config file
```
configuration found at $XDG_CONFIG_HOME/tmpst/config.yml

## features
- [ ] config file
- [ ] ascii art

## removal
```sh
make uninstall && cd .. && rm -rf tmpst
```
