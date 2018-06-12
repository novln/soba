package soba

type Appender interface {
	flush()
}

type ConsoleAppender struct {
}

func (ConsoleAppender) flush() {

}

// TODO (novln): Add a rolling system to FileAppender

type FileAppender struct {
}

func (FileAppender) flush() {

}
