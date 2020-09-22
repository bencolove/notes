# Debug Golang on Vscode
[Delve debugger][go-debug-dlv]
1. Setup tools
2. Configure environments

## Config Vscode
>Must set  
`go.gopath` := $GOPATH  
`go.goroot` := $GOROOT  
`terminal.integrated.env.linux.PATH` := $HOME/go/bin:${env:PATH}

--_`.vscode/settings.json`_--
```json
{
    // go executable path on Windows
    // "go.gopath.windows": "//wsl$/deb-go/usr/local/go",

    // $GOROOT -> /usr/local/ : go binary -> $GOROOT/bin/go 
    // $GOPATH -> ~/.go : go modules -> $GOPATH/pkg
    "go.goroot": "/usr/local/go",
    "go.gopath": "/home/roger/go",

    // shell binary with arguments on Windows
    "terminal.integrated.shell.windows": "wsl",
    "terminal.integrated.shellArgs.windows": ["-d", "deb-go", "-u", "roger"],
    // shell binary with args on Linux
    "terminal.integrated.shell.linux": "/usr/bin/bash",

    // editor settings for golang
    "go.autocompleteUnimportedPackages": true,
    "[go]": {
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true
        }
    },

    // add any path to $PATH when shell launched
    // like dlv debugger '$HOME/go/bin/dlv'
    "terminal.integrated.env.linux": {
        "PATH": "/home/roger/go/bin:${env:PATH}"
    },
}
```

## Config `dlv` Debugger
1. install `dlv`  
`$go get github.com/go-delve/delve/cmd/dlv`
>NOTE
The `dlv` executable will be by default installed in `$GOPATH/bin`. The path can be included in environment variables when shell launched as configured as above in `settings.json`.

>[Basic Usage][dlv-get-started]  
`$dlv debug app.go`  

2. config `vscode` to use `dlv`  
--_`.vscode/launch.json`_--
```json
{
    // visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            // go extension knows what to do
            "mode": "debug",
            // find main.main in any go files from the folder
            "program": "${workspaceRoot}",
            "env": {},
            "args": []
        }
    ],
}
```

[go-debug-dlv]: https://github.com/go-delve/delve
[dlv-get-started]: https://github.com/go-delve/delve/blob/master/Documentation/cli/getting_started.md