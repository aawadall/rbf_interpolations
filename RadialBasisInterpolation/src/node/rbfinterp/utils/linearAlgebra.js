// Linear Algebra Utilities

// Dot Product
function dot(a, b) {
    // if a and b are vectors invoke vectorDot 
    if (isVector(a) && isVector(b)) {
        return vectorDot(a, b);
    }

    // if a and b are matrices invoke matrixDot
    else if (isMatrix(a) && isMatrix(b)) {
        return matrixDot(a, b);
    }

    // if a is matrix and b is vector invoke matrixVectorDot
    if (isMatrix(a) && isVector(b)) {
        return matrixVectorDot(a, b);
    }

    // if a is vector and b is matrix invoke vectorMatrixDot
    if (isVector(a) && isMatrix(b)) {
        return vectorMatrixDot(a, b);
    }

    // else throw error
    throw new Error('Invalid arguments');
}

// tensor dimensionality 
function isVector(a) {
    return a.length !== undefined && a[0].length === undefined;
}

function isMatrix(a) {
    return a.length !== undefined && a[0].length !== undefined;
}

// vector dot product
function vectorDot(a, b) {
    if (a.length !== b.length) {
        throw new Error('Vectors must have the same length');
    }

    // result in map reduce 
    return a
        .map((val, i) => val * b[i])
        .reduce((acc, val) => acc + val, 0);
}

// matrix dot product
function matrixDot(a, b) {
    if (a[0].length !== b.length) {
        throw new Error('Matrices must have compatible dimensions');
    }

    // result in map reduce 
    return a
        .map(row => b[0].map((_, i) => // simply a series of vector dot products
            vectorDot(row, b.map(col => col[i]))));
}

// matrix vector dot product
function matrixVectorDot(a, b) {
    if (a[0].length !== b.length) {
        throw new Error('Matrix and vector must have compatible dimensions');
    }

    // result in map reduce 
    return a
        .map(row => vectorDot(row, b));
}

// vector matrix dot product
function vectorMatrixDot(a, b) {
    if (a.length !== b.length) {
        throw new Error('Vector and matrix must have compatible dimensions');
    }

    // result in map reduce 
    return b[0]
        .map((_, i) => vectorDot(a, b.map(col => col[i])));
}

// Transpose
function transpose(a) {
    // if a is a vector invoke vectorTranspose
    if (isVector(a)) {
        return a;
    }

    // if a is a matrix invoke matrixTranspose
    else if (isMatrix(a)) {
        return matrixTranspose(a);
    }

    // else throw error
    throw new Error('Invalid arguments');
}

// matrix transpose
function matrixTranspose(a) {
    // return a matrix with rows and columns swapped
    var result = [];

    // using map reduce to create a new matrix
    for (var i = 0; i < a[0].length; i++) {
        result.push(a.map(row => row[i]));
    }

    return result;

}

// Add 
function add(a, b) {
    // if a and b are vectors invoke vectorAdd
    if (isVector(a) && isVector(b)) {
        return vectorAdd(a, b);
    }

    // if a and b are matrices invoke matrixAdd
    else if (isMatrix(a) && isMatrix(b)) {
        return matrixAdd(a, b);
    }

    // else throw error
    throw new Error('Invalid arguments');
}

// vector add
function vectorAdd(a, b) {
    if (a.length !== b.length) {
        throw new Error('Vectors must have the same length');
    }

    // result in map reduce
    return a.map((val, i) => val + b[i]);
}

// matrix add
function matrixAdd(a, b) {
    if (a.length !== b.length || a[0].length !== b[0].length) {
        throw new Error('Matrices must have the same dimensions');
    }

    // result in map reduce
    return a.map((row, i) => row.map((val, j) => val + b[i][j]));
}

// negative 
function negative(a) {
    // if a is a vector invoke vectorNegative
    if (isVector(a)) {
        return vectorNegative(a);
    }

    // if a is a matrix invoke matrixNegative
    else if (isMatrix(a)) {
        return matrixNegative(a);
    }

    // else throw error
    throw new Error('Invalid arguments');
}

// vector negative
function vectorNegative(a) {
    // result in map reduce
    return a.map(val => -val);
}

// matrix negative
function matrixNegative(a) {
    // result in map reduce
    return a.map(row => row.map(val => -val));
}

// scalar multiplication
function scalarMult(scalar, a) {
    // if a is a vector invoke vectorScalarMult
    if (isVector(a)) {
        return vectorScalarMult(scalar, a);
    }

    // if a is a matrix invoke matrixScalarMult
    else if (isMatrix(a)) {
        return matrixScalarMult(scalar, a);
    }

    // if a is a scalar invoke scalarScalarMult
    else if (typeof a === 'number') {
        return scalar * a;
    }

    // else throw error
    throw new Error('Invalid arguments');
}

// vector scalar multiplication
function vectorScalarMult(scalar, a) {
    // result in map reduce
    return a.map(val => scalar * val);
}

// matrix scalar multiplication
function matrixScalarMult(scalar, a) {
    // result in map reduce
    return a.map(row => row.map(val => scalar * val));
}

// L2 norm 
function l2Norm(a) {
    // if a is a vector invoke vectorNorm
    if (isVector(a)) {
        return vectorL2Norm(a);
    }

    // if a is a matrix invoke matrixNorm
    else if (isMatrix(a)) {
        return matrixL2Norm(a);
    }

    // else throw error
    throw new Error('Invalid arguments');
}

// vector L2 norm
function vectorL2Norm(a) {
    // result in map reduce
    return Math.sqrt(
        a.map(val => val * val)
            .reduce((acc, val) => acc + val
                , 0)
    );
}

// matrix L2 norm
function matrixL2Norm(a) {
    // result in map reduce
    return Math.sqrt(
        a.map(row => row.map(val => val * val))
            .reduce((acc, row) => acc + row
                .reduce((acc, val) => acc + val
                    , 0)
                , 0)
    );
}

// Export
module.exports = {
    dot,
    transpose,
    add,
    negative,
    scalarMult,
    l2Norm
};