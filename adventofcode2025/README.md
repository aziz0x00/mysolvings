# Advent of Code 2025

My Go solutions for [Advent of Code 2025](https://adventofcode.com/2025).

Each day lives in its own folder and is written as a small standalone Go program.

## Live streams

I streamed the event on YouTube: [AoC 2025 playlist](https://www.youtube.com/playlist?list=PLWqo_ha9ecV_VLDVQhaB-BBNryAJDhRzS).

| Day  | Topics | Stream |
| ---- | ------ | ------ |
| 1    | `circular arithmetic` `zero-crossing count` | [Day 1](https://www.youtube.com/watch?v=hZYAxSl-mBA) |
| 2    | `range scan` `repeated-pattern strings` | [Day 2](https://www.youtube.com/watch?v=eQCPLDcxVJg) |
| 3    | `greedy selection` `digit subsequences` | [Day 3](https://www.youtube.com/watch?v=hS7blNSfURU) |
| 4    | `2D grid` `8-neighbor simulation` | [Day 4](https://www.youtube.com/watch?v=W4LYZtcYjo8) |
| 5    | `interval membership` `sort and merge` | [Day 5 - will the real struggle start????](https://www.youtube.com/watch?v=AgRjivYDEwg) |
| 6    | `column parsing` `operator folds` | [Day 6 - the end of the first half of the event???](https://www.youtube.com/watch?v=F3Dr5O7jriQ) |
| 7    | `DFS traversal` `memoized path counting` | [easy is over ??](https://www.youtube.com/watch?v=HDNnUQNoMyw) / [stream crash th ?!](https://www.youtube.com/watch?v=aFxcnlAyflg) |
| 8    | `pairwise distances` `Kruskal-style components` | [Day 8 - it's 2/3 of the event](https://www.youtube.com/watch?v=3e8sJnqXK4Q) |
| 9    | `O(n²) rectangle pairs` `point-in-polygon parity` | [Day 9 - enter last 1/3](https://www.youtube.com/watch?v=mGQJTcsk5iE) |
| 10   | `bitmasks` `combination search` `integer linear systems` | [i will miss this...](https://www.youtube.com/watch?v=N8rPAOEdkEM) / [stream cashed.., i will miss this...](https://www.youtube.com/watch?v=VgEx04UNnwE) / [part 2 the right way! late to day 11 (a)](https://www.youtube.com/watch?v=YMCodC__Fi8) / [part 2 the right way! late to day 11 (b)](https://www.youtube.com/watch?v=bE3VxGI0Mfc) |
| 11   | `DAG traversal` `memoized path counting` | [finishing day 10 part 2 implementation](https://www.youtube.com/watch?v=xNpASxjaz2U) / [We do it for the learning!](https://www.youtube.com/watch?v=kowWKa3Rr4s) |
| 12   | `bitmask shapes` `rotations and flips` `backtracking tiling` | [THE FINALE](https://www.youtube.com/watch?v=7IczG4aeDTc) |

## Layout

```
.
├── Day01/
├── Day02/
├── ...
└── Day12/
```

Each day is a self-contained folder with:

- `partOne.go` — solution for part 1
- `partTwo.go` — solution for part 2
- `input.txt` — problem input
- `input-test.txt` — example input

Some days include extra experiment files, such as `Day10/partTwo_slow.go` and
`Day10/partTwo_linalg.go`.

## Running

Run each solution from its day directory, because the programs read
`./input.txt`:

```sh
cd Day01
go run partOne.go
go run partTwo.go
```
