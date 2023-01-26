"""
Radial Basis Function Interpolation
"""

from typing import List, Tuple

import numpy as np

from rbfinterp.kernels.factory import make_kernel


class RBFInterpolator:
    """
    Radial Basis Function Interpolator
    """

    def __init__(self, points: List[Tuple[float, ...]], values: List[Tuple[float, ...]], function: str = 'gaussian', sigma: float = 0.5):
        """
        :param points: 2D array of points
        :param values: 1D array of values
        :param function: radial basis function
        :param epsilon: shape parameter
        """
        self.points = points
        self.values = values
        self.kernel_matrix_loaded = False
        parameters = {"sigma": sigma, "distance": "euclidean"}
        self.kernel = make_kernel(function, parameters)

    def make_kernel_matrix(self) -> None:
        """
        Make the kernel matrix
        """
        #self.kernel_matrix = np.zeros((len(self.points), len(self.points)))
        self.kernel_matrix = [[self.kernel.similarity(point1, point2) for point2 in self.points] for point1 in self.points]
        self.kernel_matrix_loaded = True

    def calculate_weights(self) -> None:
        """
        Calculate the weights
        """
        if not self.kernel_matrix_loaded:
            self.make_kernel_matrix()

        self.weights = np.linalg.solve(self.kernel_matrix, self.values)

    def interpolate(self, point: Tuple[float, ...]) -> float:
        """
        Interpolate a point
        """
        if not self.kernel_matrix_loaded:
            self.make_kernel_matrix()

        if not hasattr(self, "weights"):
            self.calculate_weights()

        return sum([self.weights[i] * self.kernel.similarity(point, self.points[i]) for i in range(len(self.points))])