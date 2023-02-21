// import linear algebra functions
const { dot, add, negative, scalarMult, l2Norm, transpose } = require('../utils/linearAlgebra');
// Gradient Descent Optimizer

class GradientDescent {
    constructor(learningRate = 0.1, epsilon) {
        this.learningRate = learningRate;
    }

    optimize(A, y) {
        // TODO - implement 
        console.log('Optimizing...');
        // initialize weights
        var w = Array(A[0].length).fill(0);

        // console.log('y: ' , y);
        // console.log('w: ' , w);
        // calculate yHat as A dot w
        var yHat = dot(A, w);

        // calculate error as yHat - y
        var error = add(yHat, negative(y));

        // calculate gradient as 2 * A transpose dot error
        var gradient = scalarMult(2, dot(transpose(A), error));

        // optimization loop 
        while (l2Norm(gradient) > this.epsilon) {
            console.log('J: ' + l2Norm(gradient));
            // update weights
            w = add(w, scalarMult(-this.learningRate, gradient));

            // calculate yHat as A dot w
            yHat = dot(A, w);

            // calculate error as yHat - y
            error = add(yHat, negative(y));

            // calculate gradient as 2 * A transpose dot error
            gradient = scalarMult(2, dot(transpose(A), error));
        }

        // return weights
        return w;
    }   
}

// Export
module.exports = GradientDescent;