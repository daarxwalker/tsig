# TSIG
> Typescript Index Generator written in pure Go

## Install
```bash
    go install github.com/daarxwalker/tsig@latest
```

## Config
> default config filename tsig.json
```json
{
	"options": [
		{
			"dir": "src/components",
			"export": "single",
			"recursive": true
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