"""
RBF Kernel Factory Method
"""

from typing import Dict, Optional

from rbfinterp.kernels.rbf_kernel import RBFKernel
from rbfinterp.kernels.gaussian_kernel import GaussianKernel


def make_kernel(function: str, parameters: Optional[Dict[str, float]] = None) -> RBFKernel:
    """
    Factory method for RBF kernels
    """

    if function == "gaussian":
        return GaussianKernel(parameters)
    else:
        raise ValueError("Unknown RBF kernel function: {}".format(function))