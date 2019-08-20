package monitor

type Request struct {
    Ip          string
    Route       string
    CreatedAt   int64 `sql:"default extract(epoch from now())"`
}
