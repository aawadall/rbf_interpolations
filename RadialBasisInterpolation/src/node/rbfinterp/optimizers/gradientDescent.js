// import linear algebra functions
const { dot, add, negative, scalarMult, l2Norm, transpose } = require('../utils/linearAlgebra');
// Gradient Descent Optimizer

class GradientDescent {
    constructor(learningRate = 0.01, epsilon) {
        this.learningRate = learningRate;
        this.epsilon = epsilon;
    }

    optimize(A, y) {
        // TODO - implement 
        console.log('Optimizing (using gradient descent)...');
        // initialize weights
        var w = Array(A[0].length).fill(0);

        // console.log('y: ' , y);
        // console.log('w: ' , w);
        // calculate yHat as A dot w
        var yHat = dot(A, w);

        // calculate error as  y - yHat
        var error = add(negative(yHat), y);

        // calculate gradient as 2 * A transpose dot error
        var gradient = scalarMult(2, dot(transpose(A), error));

        console.log('J: ' + l2Norm(gradient), 'epsilon: ' + this.epsilon);
        // optimization loop 
        while (l2Norm(gradient) > this.epsilon) {
            console.log('J: ' + l2Norm(gradient));
            // update weights
            w = add(w, scalarMult(this.learningRate, gradient));

            // calculate yHat as A dot w
            yHat = dot(A, w);

            // calculate error as  y - yHat
            var error = add(negative(yHat), y);

            // calculate gradient as 2 * A transpose dot error
            gradient = scalarMult(2, dot(transpose(A), error));
        }

        // return weights
        return w;
    }
}

// Export
module.exports = GradientDescent;