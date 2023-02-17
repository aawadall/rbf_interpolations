/**
 * Distance functions
 */

const Point = require('../types/point');

// Euclidean distance
const euclidean = (point1, point2) => {
    if (point1.dimensions !== point2.dimensions) {
        throw new Error('Points must have the same dimensions');
    }

    let sum = 0;
    for (let i = 0; i < point1.dimensions; i++) {
        sum += Math.pow(point1.coordinates[i] - point2.coordinates[i], 2);
    }

    return Math.sqrt(sum);
}

// Manhattan distance
const manhattan = (point1, point2) => {
    if (point1.dimensions !== point2.dimensions) {
        throw new Error('Points must have the same dimensions');
    }

    let sum = 0;
    for (let i = 0; i < point1.dimensions; i++) {
        sum += Math.abs(point1.coordinates[i] - point2.coordinates[i]);
    }

    return sum;
}

// Chebyshev distance
const chebyshev = (point1, point2) => {
    if (point1.dimensions !== point2.dimensions) {
        throw new Error('Points must have the same dimensions');
    }

    let max = 0;
    for (let i = 0; i < point1.dimensions; i++) {
        max = Math.max(max, Math.abs(point1.coordinates[i] - point2.coordinates[i]));
    }

    return max;
}

module.exports = {
    euclidean,
    manhattan,
    chebyshev
};