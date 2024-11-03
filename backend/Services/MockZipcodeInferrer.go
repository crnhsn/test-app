package services

import (
	"context"
)

type MockZipcodeInferrer struct {
	DefaultZipcodes []string
}

func NewMockZipcodeInferrer(defaultZipcodes []string) *MockZipcodeInferrer {

	repo := MockZipcodeInferrer{defaultZipcodes}
	return &repo

}

func (repo *MockZipcodeInferrer) InferLocation(ctx context.Context) ([]string, error) {

	// for mock purposes, just return the default zip codes - but ideally
	// use the context to infer the location based on incoming request params
	return repo.DefaultZipcodes, nil

}
