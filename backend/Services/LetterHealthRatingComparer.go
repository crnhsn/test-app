package services

import (
	"errors"
)

var healthRatingValues = map[string]int{
	"A": 4,
	"B": 3,
	"C": 2,
	"D": 1,
}

type LetterHealthRatingComparer struct{}

func NewLetterHealthRatingComparer() *LetterHealthRatingComparer {
	return &LetterHealthRatingComparer{}
}

func (comparer *LetterHealthRatingComparer) IsBetterOrEqual(toCompare string, baseline string) (bool, error) {
	toCompareValue, toCompareExists := healthRatingValues[toCompare]
	baselineValue, baselineExists := healthRatingValues[baseline]

	if !(toCompareExists || baselineExists) {
		return false, errors.New("invalid health rating")
	}

	return toCompareValue >= baselineValue, nil

}
