<p align="center">
  <img src="https://user-images.githubusercontent.com/44043159/176008017-3d3e51d6-5d7e-49a9-b8a2-67f9ee0ffa84.png" width="350" height="450">
</p>

    The "Local Administrator Password Solution" (LAPS) provides management of local account passwords
    of domain joined computers. Passwords are stored in Active Directory (AD) and protected by ACL, 
    so only eligible users can read it or request its reset.
---
## GUIDE
```
> to read the LAPS password (you need to have the ACL for this):

  >> ./laps -u username -p 'pass123' -d labvuln.local -h 192.168.0.31
```

</p>
<p align="center">
  <img src="https://user-images.githubusercontent.com/44043159/176015271-bcb45387-61f8-42bf-961b-e9b6f5ff9c72.png">
</p>

```bash
git clone https://github.com/march0s1as/pwn-laps/
cd pwn-laps
chmod +x install.sh
./install.sh

happy hacking!
```
