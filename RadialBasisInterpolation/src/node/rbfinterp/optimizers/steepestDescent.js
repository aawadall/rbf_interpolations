// import linear algebra functions
const { dot, add, negative, scalarMult, l2Norm, transpose } = require('../utils/linearAlgebra');

// Steepest Descent Optimizer
class SteepestDescent {
    constructor(learningRate = 0.01, epsilon, lambda = 0.5) {
        this.learningRate = learningRate;
        this.epsilon = epsilon;
        this.lambda = lambda;
    }

    optimize(A, y) {
        // TODO - implement
        console.log('Optimizing (using steepest descent)...');
        // initialize weights
        var w = Array(A[0].length).fill(0);

        // calculate gradient
        var gradient = this.calcGradient(A, y, w);

        var J = l2Norm(gradient);

        // optimization loop
        while (J > this.epsilon) {
            console.log('J: ' + J);
            // calculate gradient
            gradient = this.calcGradient(A, y, w);

            // line search
            var stepSize = this.lineSearch(A, y, w, gradient );

            // update weights
            w = add(w, scalarMult(stepSize, gradient));

            // update gradient 
            gradient = this.calcGradient(A, y, w);

            // calculate J
            J = l2Norm(gradient);

        }

        // return weights
        return w;

    }

    lineSearch(A, y, w, gradient ) {
        // TODO - implement
        console.log('Line Search...');
        var stepSize = 1.0;

        // iterate until stop converging 
        while (!this.converging(A, y, w, gradient, stepSize, )) {
            stepSize *= this.alpha;
        }

        return stepSize;
    }

    converging(A, y, w, gradient, stepSize) {
        var newW = add(w, scalarMult(stepSize, gradient));
        var newGradient = this.calcGradient(A, y, newW);
        return l2Norm(newGradient) < (1 - this.lambda * stepSize) * l2Norm(gradient);
    }

    calcGradient(A, y, w) {
        var yHat = dot(A, w);
        var error = add(negative(yHat), y);
        var gradient = scalarMult(-2, dot(transpose(A), error));
    }
}

// Export
module.exports = SteepestDescent;