package monitor

type Request struct {
    Id              string   `sql:"type:text PRIMARY KEY"`
    Ip              string
    Route           string
    StatusCode      int32
    ResponseSize    int64
    StartedAt       int64
    FinishedAt      int64
}

type Trace struct {
    Id          string `sql:"type:text PRIMARY KEY"`
    RequestId   string
    ParentId    string
    Method      string
    Service     string
    StartedAt   int64
    FinishedAt  int64
}

type Log struct {
    Id          int64   `sql:"type:serial PRIMARY KEY"`
    TraceId     string
    Message     string
    Level       string
    CreatedAt   int64
}
