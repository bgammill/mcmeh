# mcmeh
Go script that searches repo commit log for specified word.


## Dependencies
* go-git

## Why?
I wanted to experiment with goroutines. This seemed like a good way to do it.

### Usage
```
Usage of ./mcmeh:
  -repo string
    	Path to repo (default os.Getwd())
  -type string
    	Type to search. 'message' or 'hash' (default "message")
  -word string
    	Word to find (default "commit")
```

### Examples
```
./mcmeh -type="message" -repo="/Users/default/Desktop/go-git" -word="fuck"

merkletrie: fix const action type fuck up (#268)

Action constants (Insert, Delete, Modify) have type int instead of
Action.  This patch make them Actions. contains fuck
```
```
./mcmeh -type="hash" -repo="/Users/bgammill/Desktop/go-git" -word="fabb"
ae788cfabbc02c2f836f5d8c3cc18021a97e9a88 contains fabb
7a428a915ce2b7bb0f4fc6dcee77932ebacfabbf contains fabb
d3f5eaf71b01b9f7f823470ac22958bd5fabbbcd contains fabb
aab2d9b912ad7eed53fabb424fb032e58bb4d2ca contains fabb
```
