package log

type Logger interface {
	MultiHandler
	Leveler
	DebugLogger
	PrintLogger
	InfoLogger
	WarnLogger
	FatalLogger
}

type RPCLogger interface {
	Logger
	// RPC APIs
	SetRPCID(rpcID string)
	SetRequestID(requestID string)
}

type DebugLogger interface {
	// Debug APIs
	Debug(a ...interface{})
	Debugf(format string, a ...interface{})
}

type PrintLogger interface {
	// Print APIs
	Print(a ...interface{})
	Println(a ...interface{})
	Printf(f string, a ...interface{})
}

type InfoLogger interface {
	// Info APIs
	Info(a ...interface{})
	Infof(f string, a ...interface{})
}

type WarnLogger interface {
	// Warn APIs
	Warn(a ...interface{})
	Warnf(f string, a ...interface{})
}

type FatalLogger interface {
	// Fatal APIs
	Fatal(a ...interface{})
	Fatalf(f string, a ...interface{})
}

type MultiHandler interface {
	AddHandler(h Handler)
	RemoveHandler(h Handler)
}

type Leveler interface {
	Level() LevelType
	SetLevel(lv LevelType)
}
