"""
Gaussian RBF Kernel
Implements the RBFKernel protocol
"""

from typing import Any, Dict, List, Optional, Tuple

import numpy as np

from rbfinterp.distances.euclidean import euclidean_distance


class GaussianKernel:
    """
    Gaussian RBF Kernel
    """

    def __init__(self, params: Optional[Dict[str, float]]) -> None:
        """
        Initialize the RBF kernel with parameters
        """
        self.params = params

        # Take sigma from the parameters
        if "sigma" in self.params:
            self.sigma = self.params["sigma"]
        else:
            self.sigma = 0.5

        if "distance" in self.params:
            self.distance = self.params["distance"]
        else:
            self.distance = "euclidean"

    def similarity(self, point1: Tuple[float, ...], point2: Tuple[float, ...]) -> float:
        """
        Compute the similarity between two points
        """
        distance = self._distance(point1, point2)
        return np.exp(-distance / (2 * self.sigma ** 2)) 

    # def __call__(self, point1: Tuple[float, ...], point2: Tuple[float, ...], params: Any = None) -> float:
    #     """
    #     Compute the similarity between two points
    #     """
    #     distance = euclidean_distance(point1, point2)
        
    #     # in case of other methods 
    #     if "distance" in params:
    #         if params["distance"] == "euclidean":
    #             distance = euclidean_distance(point1, point2)

    #     if "sigma" in params:
    #         sigma = params["sigma"]
    #     else:
    #         sigma = 0.5

    #     return np.exp(-distance / (2 * sigma ** 2))

    def bulk_similarity(self, point1: List[Tuple[float, ...]], point2: List[Tuple[float, ...]]) -> List[float]:
        """
        Compute the similarity between two points
        """
        # TODO 
        pass

    # @property
    # def params(self) -> Any:
    #     """
    #     Get the parameters of the RBF kernel
    #     """
        
    #     return self.params

    # @params.setter
    # def params(self, params: Any) -> None:
    #     """
    #     Set the parameters of the RBF kernel
    #     """
    #     self.params = params

    @property
    def name(self) -> str:
        """
        Get the name of the RBF kernel
        """
        return "gaussian"

    def _distance(self, point1: Tuple[float, ...], point2: Tuple[float, ...]) -> float:
        """
        Compute the distance between two points
        """
        if self.distance == "euclidean":
            return euclidean_distance(point1, point2)

        return euclidean_distance(point1, point2)