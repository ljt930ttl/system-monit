namespace go example
struct Example {
    1: required string Message
}
service ExampleService {
    Example Echo(1: Example request)
}

