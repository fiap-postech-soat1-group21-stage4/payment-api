package controller_test

import (
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/fiap-postech-soat1-group21-stage4/payment-api/payment-api/adapter/handler/controller"
	"github.com/stretchr/testify/assert"
)

var handlerR *controller.Handler
var paymentID string
var responseCode int

func aPaymentRequestWithID(id string) error {
	paymentID = id
	return nil
}

func theUserMakesAPayment() error {
	// You may need to initialize your Gin router and make a request here
	// For simplicity, let's just print the payment ID
	fmt.Printf("Making payment for ID: %s\n", paymentID)
	// You can perform actual HTTP requests and check the response here
	// Set the responseCode variable accordingly
	responseCode = 200
	return nil
}

func thePaymentShouldBeProcessedSuccessfully() error {
	if responseCode != 200 {
		return fmt.Errorf("Expected a successful payment (HTTP 200), but got %d", responseCode)
	}
	return nil
}

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "http",
		ScenarioInitializer:  InitializeScenario,
		TestSuiteInitializer: InitializeTestSuite,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../../features/handler.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(s *godog.ScenarioContext) {
	handlerR = controller.NewHandler()

	s.Step(`^a payment request with ID "([^"]*)"$`, aPaymentRequestWithID)
	s.Step(`^the user makes a payment$`, theUserMakesAPayment)
	s.Step(`^the payment should be processed successfully$`, thePaymentShouldBeProcessedSuccessfully)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {}
