"""Radial Basis Function Kernel Protocol"""

from typing import Any, List, Protocol, Tuple

class RBFKernel(Protocol):
    """RBF Kernel Protocol"""

    def __init__(self, params: Any) -> None:
        """
        Initialize the RBF kernel with parameters
        """
        ...

    def similarity(self, point1: Tuple[float,...], point2: Tuple[float,...]) -> float:
        """
        Compute the similarity between two points
        """
        ...

    # def __call__(self, point1: Tuple[float,...], point2: Tuple[float,...], params:Any = None) -> float:
    #     """
    #     Compute the similarity between two points
    #     """
    #     ...

    def bulk_similarity(self, point1: List[Tuple[float,...]], point2: List[Tuple[float,...]]) -> List[float]:
        """
        Compute the similarity between two points
        """
        ...

    # @property
    # def params(self) -> Any:
    #     """
    #     Get the parameters of the RBF kernel
    #     """
    #     ...

    # @params.setter
    # def params(self, params: Any) -> None:
    #     """
    #     Set the parameters of the RBF kernel
    #     """
    #     ...

    @property
    def name(self) -> str:
        """
        Get the name of the RBF kernel
        """
        ...
