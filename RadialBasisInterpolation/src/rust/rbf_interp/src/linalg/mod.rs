// linalg/mod.rs

pub mod matrix;
pub mod vector;

/// Computes the dot product of two vectors 
pub fn dot_product(x: &Vector, y: &Vector) -> f64 {
    // if dimensions don't match, return error 
    if x.len() != y.len() {
        panic!("Dimensions don't match");
    }

    // compute dot product
    x.iter().zip(y.iter()).map(|(a, b)| a * b).sum()
}

/// Adds two vectors
pub fn add(x: &Vector, y: &Vector) -> Result<Vector, VectorError> {
    // if dimensions don't match, return error 
    if x.len() != y.len() {
        return Err(VectorError::DimensionMismatch);
    }

    // compute sum
    Ok(x.iter().zip(y.iter()).map(|(a, b)| a + b).collect())
}

/// Multiplies a vector by a scalar
pub fn scalar_multiply(x: &Vector, a: f64) -> Vector {
    x.iter().map(|&b| a * b).collect()
}

/// Computes L2 norm of a vector
pub fn norm(x: &Vector) -> f64 {
    x.iter().map(|&a| a * a).sum::<f64>().sqrt()
}

/// Computes dot product of two matrices
pub fn dot_product(x: &Matrix, y: &Matrix) -> Result<Matrix, MatrixError> {
    // if dimensions don't match, return error 
    if x.cols() != y.rows() {
        return Err(MatrixError::DimensionMismatch);
    }

    // compute dot product
    let mut result = Matrix::new(x.rows(), y.cols());
    for i in 0..x.rows() {
        for j in 0..y.cols() {
            let mut sum = 0.0;
            for k in 0..x.cols() {
                sum += x[(i, k)] * y[(k, j)];
            }
            result[(i, j)] = sum;
        }
    }
    Ok(result)
}
    