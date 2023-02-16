# mango
SSH Brute Force Tool

## Quick Start
Download the latest release or compile from source
```
usage: mango [-h|--help] -i|--ip "<value>" [-p|--port "<value>"] -u|--username
             "<value>" --password-file "<value>"

             ssh brute force

Arguments:

  -h  --help           Print help information
  -i  --ip             IP to target
  -p  --port           Port. Default: 22
  -u  --username       Username to bruteforce
      --password-file  Path to file of passwords
```

## Examples
```
# Brute Force user root on 192.168.64.5 on port 2220 with rockyou
mango -u root --password-file ./password-list.txt -i 192.168.64.5 -p 2220
```
