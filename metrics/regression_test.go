package metrics

import "testing"
import "math"
import "gonum.org/v1/gonum/mat"

func TestR2Score(t *testing.T) {
	//1st example of sklearn metrics r2score
	yTrue := mat.NewDense(4, 1, []float64{3, -0.5, 2, 7})
	yPred := mat.NewDense(4, 1, []float64{2.5, 0.0, 2, 8})
	r2Score := R2Score(yTrue, yPred, nil, "")
	eps := 1e-3
	if math.Abs(0.948-r2Score.At(0, 0)) > eps {
		t.Error("expected 0.948")
	}
	yTrue = mat.NewDense(3, 2, []float64{0.5, 1, -1, 1, 7, -6})
	yPred = mat.NewDense(3, 2, []float64{0, 2, -1, 2, 8, -5})
	if math.Abs(0.938-R2Score(yTrue, yPred, nil, "variance_weighted").At(0, 0)) >= 1e-3 {
		t.Error("expected 0.938")
	}
	yTrue = mat.NewDense(3, 1, []float64{1, 2, 3})
	yPred = mat.NewDense(3, 1, []float64{1, 2, 3})
	if math.Abs(1.-R2Score(yTrue, yPred, nil, "").At(0, 0)) >= 1e-3 {
		t.Error("expected 1")
	}
	yTrue = mat.NewDense(3, 1, []float64{1, 2, 3})
	yPred = mat.NewDense(3, 1, []float64{2, 2, 2})
	if math.Abs(0.-R2Score(yTrue, yPred, nil, "").At(0, 0)) >= 1e-3 {
		t.Error("expected 0")

	}
	yTrue = mat.NewDense(3, 1, []float64{1, 2, 3})
	yPred = mat.NewDense(3, 1, []float64{3, 2, 1})
	if math.Abs(-3.-R2Score(yTrue, yPred, nil, "").At(0, 0)) >= 1e-3 {
		t.Error("expected -3")
	}
}

// >>> from sklearn.metrics import mean_squared_error
//     >>> y_true = [3, -0.5, 2, 7]
//     >>> y_pred = [2.5, 0.0, 2, 8]
//     >>> mean_squared_error(y_true, y_pred)
//     0.375
//     >>> y_true = [[0.5, 1],[-1, 1],[7, -6]]
//     >>> y_pred = [[0, 2],[-1, 2],[8, -5]]
//     >>> mean_squared_error(y_true, y_pred)  # doctest: +ELLIPSIS
//     0.708...
//     >>> mean_squared_error(y_true, y_pred, multioutput='raw_values')
//     ... # doctest: +ELLIPSIS
//     array([ 0.416...,  1.        ])
//     >>> mean_squared_error(y_true, y_pred, multioutput=[0.3, 0.7])
//     ... # doctest: +ELLIPSIS
//     0.824...
func TestMeanSquaredError(t *testing.T) {
	yTrue := mat.NewDense(1, 4, []float64{3, -0.5, 2, 7})
	yPred := mat.NewDense(1, 4, []float64{2.5, 0.0, 2, 8})
	var sampleWeight *mat.Dense
	multioutput := ""
	if 0.375 != MeanSquaredError(yTrue, yPred, sampleWeight, multioutput).At(0, 0) {
		t.Fail()
	}
	yTrue = mat.NewDense(3, 2, []float64{0.5, 1, -1, 1, 7, -6})
	yPred = mat.NewDense(3, 2, []float64{0, 2, -1, 2, 8, -5})
	if math.Abs(0.708-MeanSquaredError(yTrue, yPred, sampleWeight, multioutput).At(0, 0)) >= 1e-3 {
		t.Fail()
	}
	multioutput = "raw_values"
	if math.Abs(0.416-MeanSquaredError(yTrue, yPred, sampleWeight, multioutput).At(0, 0)) >= 1e-3 {
		t.Fail()
	}
	if math.Abs(1.-MeanSquaredError(yTrue, yPred, sampleWeight, multioutput).At(0, 1)) >= 1e-3 {
		t.Fail()
	}
}

func TestMeanAbsoluteError(t *testing.T) {
	yTrue := mat.NewDense(1, 4, []float64{3, -0.5, 2, 7})
	yPred := mat.NewDense(1, 4, []float64{2.5, 0.0, 2, 8})
	var sampleWeight *mat.Dense
	multioutput := ""
	if 0.5 != MeanAbsoluteError(yTrue, yPred, sampleWeight, multioutput).At(0, 0) {
		t.Fail()
	}

	yTrue = mat.NewDense(3, 2, []float64{0.5, 1, -1, 1, 7, -6})
	yPred = mat.NewDense(3, 2, []float64{0, 2, -1, 2, 8, -5})
	if math.Abs(0.75-MeanAbsoluteError(yTrue, yPred, sampleWeight, multioutput).At(0, 0)) >= 1e-3 {
		t.Fail()
	}
	multioutput = "raw_values"
	if math.Abs(0.5-MeanAbsoluteError(yTrue, yPred, sampleWeight, multioutput).At(0, 0)) >= 1e-3 {
		t.Fail()
	}
	if math.Abs(1.-MeanAbsoluteError(yTrue, yPred, sampleWeight, multioutput).At(0, 1)) >= 1e-3 {
		t.Fail()
	}
}
