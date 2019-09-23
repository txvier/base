package txlogger

import (
	"fmt"
	mate "github.com/heralight/logrus_mate"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/txvier/base/txlogger/hook"
	"io"
	"os"
)

func NewRotatelogsWriter(options mate.Options) (writer io.Writer, err error) {
	fmt.Println(options)
	var conf hook.FileLogConfig
	if err = options.ToObject(&conf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return rotatelogs.New(conf.GlobPattern,
		rotatelogs.WithLinkName(conf.LinkName),              // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(conf.GetMaxAge()),             // 文件最大保存时间
		rotatelogs.WithRotationTime(conf.GetRotationTime()), // 日志切割时间间隔
		rotatelogs.WithClock(conf.GetClock()),
		rotatelogs.WithRotationCount(uint(conf.RotationCount)),
	)
}
