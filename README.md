# Radial Basis Functions (RBF) Interpolation
The _Radial Basis Functions_ (RBF) Interpolation is a technique used to approximate a continuous function from a set of discrete data points.

In this repository, we present a comprehensive library of different implementations of RBF Interpolation in various programming languages, including Python, Go, Scala, Java, and C#.

The library provides a clear and concise implementation of the mathematical theory of RBF Interpolation, including the use of different kernel types.

The library also includes a detailed description of the implementation architecture, making it easy for users to understand and adapt the code to their specific needs.

The library is intended for researchers, engineers, and students who are interested in using RBF interpolation for various applications such as geostatisics, machine learning, and computer graphics.

The library is open-source and freely available for anyone to use and contribute to under the terms of the MIT license.

## Mathematical Theory and Intuition

In this section, we will provide a comprehensive overview of the mathematical concepts and principals that underline RBF interpolation.

We will begin by explaining the intuition behind the technique and then delve into the mathematical details.

We will cover the main fomulas and concepts used in the library, including te yse of different kernel types.

By the end of this section, you will have a deep understanding of how RBF interpolation works and how to use in the library to approximate a function from a set of discrete data points.

### Intuition
We are trying to predict the values if an unknown function $\psi(\mathbf{x})$ for new inputs that fall in the same domain as a set of known inputs $\mathcal{X}$, but not explicitly mentioned in it.

We have some known data points $(\mathcal{X}, \mathcal{Y}) = \{(\mathbf{x}^{(i)}, \mathbf{y}^{(i)})\}$, where $\mathbf{x}^{(i)} \in \mathbb{R}^d$ and $\mathbf{y}^{(i)} \in \mathbb{R}^n$.

We aim to use this information to approximate the value ofs of $\psi(\mathbf{x}^{(j)})$ for new inputs $\mathbf{x} \in \mathbb{R}^d$.

__placeholder for intuition__

### Mathematical Theory
__placeholder for formulas__
## Implementation Architecture

### Python

### Go

### Scala

### Java
