# quiz


Q: Given a list of words like https://github.com/NodePrime/quiz/blob/master/word.list find the longest compound-word in the list, which is also a concatenation of other sub-words that exist in the list. The program should allow the user to input different data. The finished solution shouldn't take more than one hour. Any programming language can be used, but Go is preferred.

## Installation
Download: `go get github.com/thellimist/quiz`  
Build:    `go build`  
Run:      
`./quiz -f word.list`         
`./quiz -f ./inputs/words.txt`  
`./quiz -f ./inputs/words2.txt`  

## How it works
1) We read the file  
2) We sort the file (for faster search)  
3) We iterate over the sorted word list  
4) Recursively check if the iterated word is a compound word   

## Notes
- Word list is in memory  
- Didn't check if it supports unicode.  

## Speed
Fast enough?



