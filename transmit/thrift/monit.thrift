namespace go monit

/*
struct Example {
    1: required string Message
}
service ExampleService {
    Example Echo(1: Example request)
}
*/
enum Type {
    REQ,
    RES,
}

struct Header {
    1: required Type type,
    2: required string time,
    3: required string code,
    4: required string msg,
}

struct Response {
    1: required Type type,
    2: required string source,
    3: required string code,
    4: required string msg
}

struct RegInfo {
    1: required Header header,
    2: required string addr,
    3: required string host,
    4: optional string mac
}

struct Info {
    1: required Header header,
    2: required string info
}

service MonitService {
    void Ping(),
    Response Regist(1: RegInfo request),
    Response GetBaseInfo(1: Info request),
    Response LoadCurrentInfo(1: Info request)

}