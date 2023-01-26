"""Main file used for initial testing of the Radial Basis Interpolation"""

from typing import List, Tuple

import numpy as np
from sklearn.model_selection import train_test_split

from rbfinterp.interpolation.rbf_interpolator import RBFInterpolator


def mock_function(point: Tuple[float, ...], noise: float) -> float:
    """Mock function to interpolate"""
    # make random noise 
    applied_noise = np.random.normal(-noise, noise)

    # clean value of function
    x, y = point
    a = 1
    b = 1
    c = 1
    d = 1

    value = a * np.sin(x/ b) + c * np.cos(y/ d)
    return value + applied_noise

def make_points(num_points: int, function_range: Tuple[float,...], dimensions: int) -> List[Tuple[float, ...]]:
    """Make points to interpolate"""
    # make random points
    points = []

    for _ in range(num_points):
        point = []
        for _ in range(dimensions):
            point.append(np.random.uniform(function_range[0], function_range[1]))
        
        # convert to tuple
        point = tuple(point)
        points.append(point)

    return points


def main():
    """Main function"""
    # make mock function to interpolate
    num_points = 500
    function_range = (-2 * np.pi, 2 * np.pi)
    dimensions = 2
    noise = 0.01

    points = make_points(num_points, function_range, dimensions)
    values = [mock_function(point, noise) for point in points]

    # split points into train and test sets 
    train_points, test_points, train_values, test_values = train_test_split(points, values, test_size=0.2)


    # instantiate interpolator
    interpolator = RBFInterpolator(train_points, train_values, function="gaussian", sigma=0.71)

    # interpolate test points
    interpolated_values = [interpolator.interpolate(point) for point in test_points]

    # calculate error
    error = np.mean(np.abs(np.array(interpolated_values) - np.array(test_values)))
    print(f"Mean Error: {error}")

    for actual, interpolated in zip(test_values, interpolated_values):
        print(f"Actual: {actual}, Interpolated: {interpolated}, Relative Error: {100 * np.abs((actual - interpolated)/actual)} %")
    
if __name__ == "__main__":
    main()