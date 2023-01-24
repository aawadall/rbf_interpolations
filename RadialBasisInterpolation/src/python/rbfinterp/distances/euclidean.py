"""
Eucledian distance
"""

from typing import Tuple

import numpy as np


def euclidean_distance(point_1: Tuple[float,...], point_2: Tuple[float,...]) -> float:
    """
    Euclidean distance between two points
    """
    # convert to numpy arrays
    x1 = np.array(point_1)
    x2 = np.array(point_2)
    
    return np.sqrt(np.sum((x1 - x2)**2))