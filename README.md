# passg
Generate secure passwords with given length and symbols

## Compile
```
go build passg.go
```
## Usage
```
passg --help
```
```
Usage: ./passg [-len 8] [-count 1] [-all]
Generate secure passwords with given length and symbols
 
Options:
  -len      Password length
  -count    Number of passwords to generate
  -all      Use all English letters (upper/lower), digits and special characters
  -up       Use upper case letters
  -low      Use lower case letters
  -digits   Use digits
  -spec     Use special characters
	
  -help     Print this help message

```
## Examples
`passg -all`
`mez$9Wo=`

`passg -all -len 10 -count 5`
```
xXcVo$3P#3
xPw8X6vn3A
WdVuSSMiWV
vpKP93m42p
3#i9mY@qm!
```
