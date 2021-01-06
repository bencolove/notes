# Manage Multiple `nodejs`
By [nvm][nvm]
* on Linux

---
>1. Install `nvm` on Linux  
```sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.34.0/install.sh | bash
# OR
wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.34.0/install.sh | bash
```
Structure:
* root folder: `$HOME/.nvm`
* installed node: `$HOME/.nvm/versions/node/<VERSION>`

>2. Verify  

`nvm`  
`nvm ls-remote`  
`nvm ls`  

>3. Install Specific `nodejs` Version  

`nvm install <VERSION>`  
`nvm ls`  
`nvm which <VERSION>`  

>4. Change Current `nodejs`  

`nvm use <VERSION>`  

>5. Run `nodejs`  

`nvm run <VERSION>`  


[nvm]: https://github.com/nvm-sh/nvm