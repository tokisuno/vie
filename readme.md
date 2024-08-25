# vie -- Dead-simple journalling CLI tool (written in go!)
Vie (meaning life in French) grabs the current date, and checks to see if there's an already existing file of the same name in your journal directory. If there isn't it will create a new one for you. 

This is only meant to be a tool to make daily journalling easier. 

## Installation
* Method 1:
```bash
# 1. git clone into a repo
git clone https://github.com/tokisuno/vie.git
# 2. run go build inside project dir
go run
# 3. move binary into a folder in your path 
cp vie ~/.local/bin
```

## TODO
### [ ] User config 
* Options
    - name
    - dir
    - editor
### [ ] Template markdown files
* For now, just manually add to the buffer while creating it
