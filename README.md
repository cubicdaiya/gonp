# gonp

gonp is a diff algorithm implementation in Go.

# Algorithm

The algorithm `gonp` uses is based on "An O(NP) Sequence Comparison Algorithm" by described by Sun Wu, Udi Manber and Gene Myers.
An O(NP) Sequence Comparison Algorithm(following, Wu's O(NP) Algorithm) is the efficient algorithm for comparing two sequences.

## Computational complexity

The computational complexity of Wu's O(NP) Algorithm is averagely O(N+PD), in the worst case, is O(NP).

# Getting started

## strdiff

```go
diff := gonp.New[rune]([]rune("abc"), []rune("abd"))
diff.Compose()
ed := diff.Editdistance() // ed is 2
lcs := diff.Lcs() // lcs is "ab"

ses := diff.Ses()
// ses is []SesElem{
//        {e: 'a', t: Common},
//        {e: 'b', t: Common},
//        {e: 'c', t: Delete},
//        {e: 'd', t: Add},
//        }
```

## intdiff

```go
diff := gonp.New[int]([]int{1,2,3}, []int{1,5,3})
diff.Compose()
ed := diff.Editdistance() // ed is 2
lcs := diff.Lcs() // lcs is [1,3]

ses := diff.Ses()
// ses is []SesElem{
//        {e: 1, t: Common},
//        {e: 2, t: Delete},
//        {e: 5, t: Add},
//        {e: 3, t: Common},
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
