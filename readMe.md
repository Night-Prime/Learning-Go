# Go Program Documentation

## Overview

This Go program demonstrates the use of generics and concurrency. It includes functions that utilize generics to create reusable code and goroutines to handle concurrent operations.

## Functions

### `genericSample[T any](s []T) ([]T, []T)`

This function takes a slice of any type `T` and splits it into two slices. It returns the first half of the slice twice.

#### Parameters:

- `s`: A slice of type `T`.

#### Returns:

- Two slices of type `T`, both containing the first half of the input slice.

### `sendEmail(message string)`

This function sends an email asynchronously using a goroutine. It prints a message indicating that the email has been sent and another message when the email is received after a short delay.

#### Parameters:

- `message`: A string containing the email message.

### `filterOldEmails(emails []email)`

This function checks if emails are old by comparing their dates to a specified date. It uses a goroutine and a channel to handle the checks concurrently.

#### Parameters:

- `emails`: A slice of `email` objects (the `email` type should be defined elsewhere in the program).

## Notes

- The program demonstrates the use of generics in Go, allowing for reusable functions that can operate on different data types.
- Concurrency is handled using goroutines and channels, enabling asynchronous operations and communication between goroutines.
