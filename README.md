# TSIG
> Typescript Index Generator written in pure Go

## Install
```bash
go install github.com/daarxwalker/tsig@latest
```

## Config
- default config filename tsig.json
- subdir can have different option
```json
{
    "root": "../",
    "options": [
        {
            "dir": "src/components",
            "export": "single",
            "recursive": true
        },
        {
            "dir": "src/components/style",
            "export": "all",
            "recursive": false
        }
    ]
}
```

## Usage
### Default
```bash
tsig
```
### With a custom config path
```bash
tsig -config=tsig.json
```