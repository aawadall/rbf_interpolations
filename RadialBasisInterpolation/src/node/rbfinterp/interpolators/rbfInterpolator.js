// rbf interpolator class

class RBFInterpolator {
    // constructor
    constructor(points, values, kernel, epsilon, optimizer) {
        // TODO
        this.points = points;
        this.values = values;
        this.kernel = kernel;
        this.epsilon = epsilon;
        this.optimizer = optimizer;
    }

    // train 
    train() {
        // TODO - compute K matrix 
        var K = this.points.map(p1 => this.points.map(p2 => this.kernel.similarity(p1, p2)));

        // TODO - compute weights
        this.weights = this.optimizer.optimize(K, this.values, this.epsilon);
    }

    // predict
    predict(point) {
        // TODO
        // compute K matrix
        var K = this.points.map(p => this.kernel.similarity(p, point));

        // compute prediction
        var prediction = 0;
        for (let i = 0; i < this.weights.length; i++) {
            prediction += this.weights[i] * K[i];
        }

        return prediction;
    }

}

// Export RBFInterpolator
module.exports = RBFInterpolator;