# Regex Go
#### Use the NFA to implement which inspire by Russ Cox

# Usage
```go
reg := NewRegexGo("a*b") //Create new regex pattern 
result := reg.CheckIsMatch("aaaab") // return boolean type to check the string is match the pattern or not
```



# Reference
https://swtch.com/~rsc/regexp/regexp1.html