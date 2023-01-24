"""
Gaussian RBF Kernel
Implements the RBFKernel protocol
"""

from typing import Any, Dict, List, Optional, Tuple

import numpy as np


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

    def similarity(self, point1: Tuple[float, ...], point2: Tuple[float, ...]) -> float:
        """
        Compute the similarity between two points
        """
        # TODO 
        pass 

    def __call__(self, point1: Tuple[float, ...], point2: Tuple[float, ...], params: Any = None) -> float:
        """
        Compute the similarity between two points
        """
        # TODO
        pass

    def bulk_similarity(self, point1: List[Tuple[float, ...]], point2: List[Tuple[float, ...]]) -> List[float]:
        """
        Compute the similarity between two points
        """
        # TODO 
        pass

    @property
    def params(self) -> Any:
        """
        Get the parameters of the RBF kernel
        """
        # TODO
        pass

    @params.setter
    def params(self, params: Any) -> None:
        """
        Set the parameters of the RBF kernel
        """
        # TODO
        pass

    @property
    def name(self) -> str:
        """
        Get the name of the RBF kernel
        """
        return "gaussian"
