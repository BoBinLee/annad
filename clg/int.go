package clg

import (
	"math"
)

func (i *index) DivideInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 / i2

	return []interface{}{s}, nil
}

func (i *index) GreaterThanInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	greater := i1 > i2

	return []interface{}{greater}, nil
}

func (i *index) LesserThanInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	lesser := i1 < i2

	return []interface{}{lesser}, nil
}

func (i *index) MultiplyInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 * i2

	return []interface{}{s}, nil
}

func (i *index) PowInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := math.Pow(float64(i1), float64(i2))

	return []interface{}{s}, nil
}

func (i *index) SqrtInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	s := math.Sqrt(float64(i1))

	return []interface{}{s}, nil
}

func (i *index) SubtractInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 - i2

	return []interface{}{s}, nil
}

func (i *index) SumInt(args ...interface{}) ([]interface{}, error) {
	i1, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	i2, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	s := i1 + i2

	return []interface{}{s}, nil
}
