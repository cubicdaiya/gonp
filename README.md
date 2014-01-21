# gonp

[![Build Status](https://travis-ci.org/cubicdaiya/gonp.png?branch=master)](https://travis-ci.org/cubicdaiya/gonp)

gonp is a diff algorithm implementation in Go.

# How To

```go
diff := gonp.New("abc", "abd")
diff.Compose()
ed := diff.Editdistance() // ed is 2
lcs := diff.Lcs() // lcs is "ab"

ses := diff.Ses()
// ses is []SesElem{
//        {c: 'a', t: Common},
//        {c: 'b', t: Common},
//        {c: 'c', t: Delete},
//        {c: 'd', t: Add},
//        }
```
