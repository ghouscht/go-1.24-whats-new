# Synthetic time for testing (experimental)

To enable, set `GOEXPERIMENT=synctest`.

## How does it work?
Functions within a `synctest.Run()` execute in an isolated "bubble" where the `time`
package functions use a fake clock. The time is initially set to midnight 01.01.2000
(UTC), Whenever all goroutines are blocked, the time automatically advances.
