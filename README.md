# gonp

[![Build Status](https://travis-ci.org/cubicdaiya/gonp.png?branch=master)](https://travis-ci.org/cubicdaiya/gonp)

gonp is the diff library in Go.

# How To

```go
diff := New("abc", "abd")
diff.Compose()
ed := diff.Editdistance() // ed is 2
lcs := diff.Lcs() // lcs is "ab"

// ses is []SesElem{
//        {c: 'a', t: Common},
//        {c: 'b', t: Common},
//        {c: 'c', t: Delete},
//        {c: 'd', t: Add},
//        }
ses := diff.Ses()
```
