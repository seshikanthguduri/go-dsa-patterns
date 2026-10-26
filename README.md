# Go DSA Patterns

Data structures & algorithms in Go, organized **by pattern, not by problem list**.

Interview problems repeat a small set of ideas. This repo builds one directory
per pattern: the pattern's gold-mine problem first (the one that teaches the
idea), then variations. Every solution ships with table-driven tests and a
short note on the mental model, time, and space.

Built as part of a deliberate Go + DSA study track. Progress below is honest:
checked means written, tested, and passing `go test -race`.

## Patterns

- [x] **01 · Hashing** — `01_hashing/` — Two Sum (LC 1)
- [x] **02 · Two Pointers** — `02_two_pointers/` — Valid Palindrome (LC 125)
- [x] **03 · Sliding Window** — `03_sliding_window/` — Best Time to Buy and Sell Stock (LC 121)
- [x] **04 · Binary Search** — `04_binary_search/` — Search in Rotated Sorted Array (LC 33)
- [ ] **05 · Stack** — Valid Parentheses (LC 20)
- [ ] **06 · Linked List** — Reverse Linked List (LC 206)
- [ ] **07 · Trees (DFS)** — Maximum Depth of Binary Tree (LC 104)
- [ ] **08 · Trees (BFS)** — Binary Tree Level Order Traversal (LC 102)
- [ ] **09 · Heap / Priority Queue** — Kth Largest Element (LC 215)
- [ ] **10 · Backtracking** — Subsets (LC 78)
- [ ] **11 · Graphs (BFS/DFS)** — Number of Islands (LC 200)
- [ ] **12 · Dynamic Programming (1D)** — Climbing Stairs (LC 70)
- [ ] **13 · Dynamic Programming (2D)** — Unique Paths (LC 62)
- [ ] **14 · Greedy** — Jump Game (LC 55)
- [ ] **15 · Intervals** — Merge Intervals (LC 56)
- [ ] **16 · Bit Manipulation** — Single Number (LC 136)

Target: a 60-problem ladder across these 16 patterns.

## Run

```bash
go vet ./...
go test -race ./...
```

## Layout

Each pattern directory is its own package: one file per problem
(`problem.go`), one table-driven test file (`problem_test.go`), and a header
comment in the solution stating the idea in two or three lines.
