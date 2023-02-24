"""Optimizer Protocol"""
from typing import Protocol

import numpy as np


class Optimizer(Protocol):
    """Optimizer Protocol"""

    def optimize(self, A: np.ndarray, b: np.ndarray) -> np.ndarray:
        """Optimize the given system of linear equations"""
        ...
