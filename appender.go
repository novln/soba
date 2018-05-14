package soba

type Appender interface {
	flush()
}

type ConsoleAppender struct {
}

func (ConsoleAppender) flush() {

}

type FileAppender struct {
}

func (FileAppender) flush() {

}
