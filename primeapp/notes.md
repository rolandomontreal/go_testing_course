# Notes from section 3

Easiest to run multiple test cases is with table tests, a slice of structs with a name, test value and expected and run tests in a loop.

Some commands for testing and what they do:

```bash
go test # Run tests in current folder
go test -cover # Get very simple coverage report of the tests
go test -coverprofile=coverage.out # Create a simple test coverage file called `coverage.out`
go tool cover -html=coverage.out # Open `coverage.out` in a browser to see what statements are tested or not
```