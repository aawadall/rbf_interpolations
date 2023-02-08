# Go Implementation of Radial Basis Interpolation

## Challenges
Need to figure out how to calculate inverse of a matrix in Go.

need to check out Gauss-Jordan elimination algorithm as follows

* Assuming a square matrix $\mathbf{A} \in \mathbb{R}^{m \times m}$
* Let $\mathbf{I} \in \mathbb{R}^{m \times m}$ be the identity matrix
* Let $\mathbf{B} = [\mathbf{A} \quad \mathbf{I}] \in \mathbb{R}^{m \times 2m}$
* Let $\mathbf{C} = \mathbf{B}$
* For $i = 1, 2, \dots, m$
    * Let $p = i$
    * For $j = i + 1, i + 2, \dots, m$
        * If $|\mathbf{C}_{j,i}| > |\mathbf{C}_{p,i}|$
            * Let $p = j$
    * Swap rows $i$ and $p$ of $\mathbf{C}$
    * For $j = i + 1, i + 2, \dots, m$
        * Let $\mathbf{C}_{j,i} = \mathbf{C}_{j,i} / \mathbf{C}_{i,i}$
        * For $k = i + 1, i + 2, \dots, 2m$
            * Let $\mathbf{C}_{j,k} = \mathbf{C}_{j,k} - \mathbf{C}_{j,i} \cdot \mathbf{C}_{i,k}$
* Let $\mathbf{D} = \mathbf{C}$
* For $i = m, m - 1, \dots, 1$
    * For $j = i - 1, i - 2, \dots, 1$
        * Let $\mathbf{D}_{j,i} = \mathbf{D}_{j,i} / \mathbf{D}_{i,i}$
        * For $k = i - 1, i - 2, \dots, 1$
            * Let $\mathbf{D}_{j,k} = \mathbf{D}_{j,k} - \mathbf{D}_{j,i} \cdot \mathbf{D}_{i,k}$
* Let $\mathbf{E} = \mathbf{D}$
* Let $\mathbf{A}^{-1} = \mathbf{E}[:,m+1:2m]$

