![workflow status](https://github.com/cubicdaiya/gonp/actions/workflows/go.yml/badge.svg)

# gonp

gonp is a diff algorithm implementation in Go.

# Algorithm

The algorithm `gonp` uses is based on "An O(NP) Sequence Comparison Algorithm" by described by Sun Wu, Udi Manber and Gene Myers.
An O(NP) Sequence Comparison Algorithm(following, Wu's O(NP) Algorithm) is the efficient algorithm for comparing two sequences.

## Computational complexity

The computational complexity of Wu's O(NP) Algorithm is averagely O(N+PD), in the worst case, is O(NP).

# Getting started

## string difference

```go
diff := gonp.New([]rune("abc"), []rune("abd"))
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

## int array difference

```go
diff := gonp.New([]int{1,2,3}, []int{1,5,3})
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

## unified format difference

```go
diff := gonp.New([]rune("abc"), []rune("abd"))
diff.Compose()

uniHunks := diff.UnifiedHunks()
diff.PrintUniHunks(uniHunks)
// @@ -1,3 +1,3 @@
//  a
//  b
// -c
// +d
```



# Example

## strdiff

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

## intdiff

```
$ make intdiff
go build -o intdiff examples/intdiff.go
$ ./intdiff
diff [1 2 3 4 5] [1 2 9 4 5]
Editdistance: 2
LCS: [1 2 4 5]
SES:
  1
  2
- 3
+ 9
  4
  5
```

## unistrdiff

```
$ make unistrdiff
go build -o unistrdiff examples/unistrdiff.go
$ ./unistrdiff abc abd
Editdistance:2
LCS:ab
Unified format difference:
@@ -1,3 +1,3 @@
 a
 b
-c
+d
```

## uniintdiff

```
$ make uniintdiff
go build -o uniintdiff examples/uniintdiff.go
$ ./uniintdiff
diff [1 2 3 4 5] [1 2 9 4 5]
Editdistance: 2
LCS: [1 2 4 5]
Unified format difference:
@@ -1,5 +1,5 @@
 1
 2
-3
+9
 4
 5
```

## unifilediff

```
$ make unifilediff
go build -o unifilediff examples/unifilediff.go
$ cat a.txt
a
b
c
$ cat b.txt
a
b
d
$ ./unifilediff a.txt b.txt
@@ -1,3 +1,3 @@
 a
 b
-c
+d
```
