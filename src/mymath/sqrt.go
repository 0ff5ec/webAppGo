package mymath

func Sqrt(f float64) float64 {
    z := 0.0
    for i := 0; i < 1000; i++ {
        z -= (z*z - f)/(2*f)
    }
    return z
}
