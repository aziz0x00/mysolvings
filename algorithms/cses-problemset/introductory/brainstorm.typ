#set page(margin: 1.5em)
#set text(size: 1.3em)

"A game of $n$ cards" $<=>$ a permutation $sigma$ of ${1, ..., n}$, defined by
$sigma(k) =$ "the adversary's play \#"

we're solving the problem of determining the permutation $sigma$ given
$
  a & := \# {k : sigma(k) < k} \
  b & := \# {k : sigma(k) > k}.
$
#line(length: 100%)
#set math.mat(delim: none)
// #show math.equation: set text(font: "Iosevka Extended")
$
  overbrace(
    mat(
      1, 0, 0, 0, 0;
      0, 0, 0, 1, 0;
      0, 1, 0, 0, 0;
      0, 0, 0, 0, 1;
      0, 0, 1, 0, 0;
    ), (a=2, b=2)
  )
  quad
  ->
  quad
  overbrace(
    mat(
      1, 0, 0, 0, 0;
      0, 0, 0, 1, 0;
      0, 0, 0, 0, 1;
      0, 1, 0, 0, 0;
      0, 0, 1, 0, 0;
    ), (a=2, b=2)
  )
  quad;
  quad
  overbrace(
    mat(
      1, 0, 0, 0, 0;
      0, 0, 1, 0, 0;
      0, 0, 0, 1, 0;
      0, 1, 0, 0, 0;
      0, 0, 0, 0, 1;
    ), (a=1, b=2)
  )
  quad
  ->
  quad
  overbrace(
    mat(
      1, 0, 0, 0, 0;
      0, 1, 0, 0, 0;
      0, 0, 0, 1, 0;
      0, 0, 0, 0, 1;
      0, 0, 1, 0, 0;
    ), (a=1, b=2)
  )
$

for $sigma : {1,...,n} -> {1, ..., n}$, associate: $A_sigma = (delta_(i, sigma(i)))$ where $delta_(i j) = 1 "if" i = j "otherwise" 0$.

(a=0, b=5)

```
12 34567
12 37654
	// (n-a-b)      (a)              (b)
	// [DRAWs] [PlayerA wins] [PlayerB wins]

```

$k$, with a rotation of $p$, is compared to $ "rot"_p (k) := (k+p) mod n = cases(k+p &"if" k < n-p, k - underbrace((n - p), >0) quad &"otherwise") $
so $k < "rot"_p (k)$ if $k < n-p$, otherwise, $k > "rot"_p (k)$

#pagebreak()
= Week\#3

Conjectures:
- The matrix is symmetric
- For each even row, the next is disjoint transposes of it.

#let AA = $mono("A")$

$ AA[i][j] = min { AA[0:i-1, j-1] union AA[i-1, 0:j-1] }^c $

