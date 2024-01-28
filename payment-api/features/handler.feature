Feature: Payment API

  Scenario: Making a Payment
    Given a payment request with ID "123"
    When the user makes a payment
    Then the payment should be processed successfully
