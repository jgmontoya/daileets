# [542. 01 Matrix](https://leetcode.com/problems/01-matrix/)

Given an `m x n` binary matrix `mat`, return the _distance of the nearest `0` for each cell._

The distance between two adjacent cells is `1`.


Example 1:

![Example 1](01-1-grid.jpg)
```
Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]
```
Example 2:

![Example 2](01-2-grid.jpg)
```
Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]
```

Constraints:
* `m == mat.length`
* `n == mat[i].length`
* `1 <= m, n <= 10^4`
* `1 <= m * n <= 10^4`
* `mat[i][j]` is either `0` or `1`.
* There is at least one `0` in mat.
