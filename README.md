# gonp

[![Build Status](https://travis-ci.org/cubicdaiya/gonp.png?branch=master)](https://travis-ci.org/cubicdaiya/gonp)

gonp is a diff algorithm implementation in Go.

# Algorithm

The algorithm `gonp` uses is based on "An O(NP) Sequence Comparison Algorithm" by described by Sun Wu, Udi Manber and Gene Myers.
An O(NP) Sequence Comparison Algorithm(following, Wu's O(NP) Algorithm) is the efficient algorithm for comparing two sequences.

## Computational complexity

The computational complexity of Wu's O(NP) Algorithm is averagely O(N+PD), in the worst case, is O(NP).

# How to

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

# Example

```
$ make strdiff
go build -o strdiff examples/strdiff.go
$ ./strdiff abc abd
Editdistance: 2
LCS: ab
SES:
  a
  b
- c
+ d
```
