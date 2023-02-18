// TODO

main = function() {
    // TODO
    // define dataset 
    const points = pointsMaker([-1, 1], [-1, 1], 100, 100);
    const values = points.map(p => noiseWrapper(hiddenFunction(p.x, p.y), 0.5));

}

// define hidden function 
hiddenFunction = (x, y) => {
    // TODO
    return x*x + y*y;
}

// define noise wrapper 
noiseWrapper = (fn, signalToNoise) => {
    // Add noise to the function
    return fn * signalToNoise + (1 - signalToNoise) * Math.random();
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