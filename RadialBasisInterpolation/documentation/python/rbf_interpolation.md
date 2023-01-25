# Radial Basis Interpolation in Python

TODO: Add description

## Architecture
RBF Interpolation Library in Python `rbfinterp` is a Python library that provides the functionality to interpolate the value of a d-dimensional point in a (n+d) dimensional space using Radial Basis Functions (RBF). 

To use this library, you need to instantiate an object of the class `RBFInterpolation` and feed it with _supporting points_ and their _values_. You also need to provide it with the _kernel function_ that you want to use for similarity measurement, in addition to the _distance function_ that you want to use to calculate the distance between two points.

