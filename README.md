# Radial Basis Functions (RBF) Interpolation (DRAFT)
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
RBF Interpolation is a technique that uses a set of known data points to approximate the value of an unknown function for new points that fall in the same domain. 

Let us consider the example of weather forecasting, where we are attempting to forecast the weather conditions in a certain location, given latitude, longitude, and elevation, in addition to date and time. We are interested in the following numbers, temperature, pressure, humidity, and wind speed. We happen to have weather stations scattered across the globe, gathering data points on some random time intervals, storing them in a large matrix. Would it be possible to utilize this matrix for predicting the weather conditions in a new location, given the same parameters?

RBF interpolation can help us with this problem by accounting for the influence of neighboring locations on a value of a given location, while giving less weight to locations that are further away. This technique uses _Radial Basis Functions_ (RBF) to calculate the similarity between an unknown point and each of the known points based on the distance between them. Points that are close to the unknown point are given more weight, while points that are farther away are given less weight. The final approximation for the unknown point is then calculated by taking a weighted average of the known values, where weights are calculated using the RBF.

This approach allowes the technique to take into account the local behavior of the function and it is particularly useful when the function is smooth and has a limited number of local variations.


### Mathematical Foundations
The mathematicl foundations of RBF interpolation are based on the use of a radial basis function (RBF) $\phi(.)$ to calculate the similarity between two points in our input space. 
The RBF assigns a weight to each known point based on its distance from the _query point_ $x_q$. 
Points that are close to the query point are given more weight, while points that are farther away are given less weight. The final approximation for the query point is then calculated by taking a weighted average of the known values, where weights are calculated using the RBF.

Assuming we have a set of known data points in input space $\mathbf{x}^{(i)} \in \mathbb{R}^d$ we call that set $\mathcal{X}$. Each point $\mathbf{x}^{(i)}$ has a vector of values $\mathbf{y}^{(i)} \in \mathbb{R}^n$ associated with it. We call this set $\mathcal{Y}$.

We are interested in approximating the value of a function $\psi(\mathbf{x})$ for a new point $\mathbf{x}_q \in \mathbb{R}^d$ that falls in the same domain as the set of known points $\mathcal{X}$.

In this context, we denote the actual function $\psi(\mathbf{x}) : \mathbb{R}^d \mapsto \mathbb{R}^n$ as the _true function_ and the approximation $\hat{\psi}(\mathbf{x})$ of it as the _interpolant_.

To approximate the value of the true function for a new point $\mathbf{x}_q$, we use the following formula:
$$
    \hat{\mathbf{y}}_q = \sum_{i=1}^N \phi(||\mathbf{x}_q - \mathbf{x}^{(i)}||) \cdot \mathbf{w}^{(i)}
$$

Such that $\mathbf{w}^{(i)}$ is the vector of weights associated with the known point (also known as support points) $\mathbf{x}^{(i)}$ and $\phi(||\mathbf{x}_q - \mathbf{x}^{(i)}||)$ is the RBF that calculates the similarity between the query point $\mathbf{x}_q$ and the point $\mathbf{x}^{(i)}$.

The RBF function is an n-dimensional function that takes a scalar value as input and returns an n-dimensional vector as output. The RBF function can be one of many possible functions.

TODO
__placeholder for formulas__
## Implementation Architecture

### Python

### Go

### Scala

### Java
