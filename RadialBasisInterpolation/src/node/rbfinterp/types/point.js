/**
 * @typedef {Object} Point
 * @property {number} dimensions
 * @property {number[]} coordinates
 */

class Point {
    /**
     * @param {number[]} coordinates
     */
    constructor(coordinates) {
        this.coordinates = coordinates;
        this.dimensions = coordinates.length;
    }

}

// Export
module.exports = Point;