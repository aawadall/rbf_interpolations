// imports 
const GradientDescent = require('./optimizers/gradientDescent');
const GaussianKernel = require('./kernels/gaussianKernel');
const RBFInterpolator = require('./interpolators/rbfInterpolator');
const Point = require('./types/point');
const { euclidean } = require('./utils/distance');
// TODO

main = function() {
    // TODO
    // define dataset 
    const points = pointsMaker([-1, 1], [-1, 1], 100, 100);
    const values = points.map(p => noiseWrapper(hiddenFunction(p), 0.5));

    // define calculators 
    const optimizer = new GradientDescent(0.1, 0.0001);
    const distance = euclidean;
    const kernel = new GaussianKernel(distance, { sigma: 0.1 });

    // define interpolator
    const interpolator = new RBFInterpolator(points, values, kernel, 0.1, optimizer);

    // train interpolator
    interpolator.train();

    // define test points
    const testPoints = pointsMaker([-1.1, 1.1], [-1.1, 1.1], 10, 10);

    // predict values
    const predictions = testPoints.map(p => interpolator.predict(p));

    // report results 
    report(testPoints, predictions, hiddenFunction)

}

// define report function
report = (points, predictions, hiddenFunction) => {
    // TODO
    // compute error
    var error = 0;
    for (let i = 0; i < points.length; i++) {
        error += Math.pow(predictions[i] - hiddenFunction(points[i].x, points[i].y), 2);
    }
    error /= points.length;

    // report error
    console.log('Normalized Error: ' + error);
}
// define hidden function 
hiddenFunction = (p) => {
    const x = p.coordinates[0];
    const y = p.coordinates[1];
    const result = x * x + y * y;
    // console.log(`x: ${x}, y: ${y}, result: ${result}`);
    // TODO
    return result;
}

// define noise wrapper 
noiseWrapper = (fn, signalToNoise) => {
    const result = fn * signalToNoise + (1 - signalToNoise) * Math.random();
    // console.log(result);
    // Add noise to the function
    return result;
}

// define points maker 
pointsMaker = (xInterval, yInterval, numX, numY) => {
    var points = [];
    // xInterval = [xMin, xMax]
    const xStep = (xInterval[1] - xInterval[0]) / numX;
    // yInterval = [yMin, yMax]
    const yStep = (yInterval[1] - yInterval[0]) / numY;

    // define points grid
    for (let i = 0; i < numX; i++) {
        const x = xInterval[0] + i * xStep;
        for (let j = 0; j < numY; j++) {
            const y = yInterval[0] + j * yStep;
            points.push(new Point([x, y]));
        }
    }
    // return points
    return points;
}
// Run the main function
main();