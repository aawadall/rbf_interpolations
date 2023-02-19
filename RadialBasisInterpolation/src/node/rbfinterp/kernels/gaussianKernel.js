// Gaussian kernel
// implements the kernel.js interface

class GaussianKernel {
    // constructor
    constructor(distanceFunction, options) {
        // TODO
        this.sigma = options.sigma;
        this.distanceFunction = distanceFunction;
    }

    // similarity
    similarity(a, b) {
        // TODO
        var distance = this.distanceFunction(a, b);

        return Math.exp(-distance / (2 * Math.pow(this.sigma, 2)));
    }
}

// export Gaussian kernel
module.exports = GaussianKernel;