package simple_log

import (
	"os"
	"time"

	"github.com/arthurkiller/rollingwriter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel  = zap.DebugLevel
	InfoLevel   = zap.InfoLevel
	WarnLevel   = zap.WarnLevel
	ErrorLevel  = zap.ErrorLevel
	DPanicLevel = zap.DPanicLevel
	PanicLevel  = zap.PanicLevel
	FatalLevel  = zap.FatalLevel
)

func init() {
	// 初始化默认值
	logger = &Logger{}
	logger.logger, _ = zap.NewDevelopment()
	logger.level = zap.NewAtomicLevel()
}

type Logger struct {
	logger *zap.Logger
	level  zap.AtomicLevel

	writer  zapcore.WriteSyncer
	encoder zapcore.Encoder
}

type Option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(logger *Logger) {
	f(logger)
}

// SetLevel 设置日志等级
func SetLevel(level zap.AtomicLevel) Option {
	return optionFunc(func(l *Logger) {
		logger.level = level
	})
}

// SetFileWriter 设置日志存放路径
func SetFileWriter(filePath string) Option {
	return optionFunc(func(l *Logger) {
		if filePath == "" {
			// 如果未配置日志路径，使用默认路径
			filePath = "./log"
		}
		rollingWriter, err := rollingwriter.NewWriterFromConfig(&rollingwriter.Config{
			TimeTagFormat:          "060102150405", // 时间格式
			LogPath:                filePath,
			FileName:               "logger",
			MaxRemain:              10,
			RollingPolicy:          rollingwriter.TimeRolling,
			RollingTimePattern:     "0 */10 * * * *", // 滚动时间：每十分钟
			RollingVolumeSize:      "500M",
			WriterMode:             "lock",
			BufferWriterThershould: 8 * 1024 * 1024,
			Compress:               false,
			FilterEmptyBackup:      false,
		})
		if err != nil {
			return
		}
		logger.writer = zapcore.AddSync(rollingWriter)
	})
}

// InitLogger 初始化 Logger，其中可以选配日志等级、指定日志路径
func InitLogger(opts ...Option) (err error) {
	config := zapcore.EncoderConfig{
		MessageKey:     "name",
		LevelKey:       "Level",
		TimeKey:        "log-time",
		NameKey:        "msg",
		CallerKey:      "caller",
		FunctionKey:    "function",
		StacktraceKey:  "stack",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Local().Format("2006-01-02 15:04:05"))
		},
		EncodeDuration:      zapcore.StringDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          zapcore.FullNameEncoder,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "",
	}

	// set default encoder
	logger.encoder = zapcore.NewJSONEncoder(config)

	// set default writer
	logger.writer = os.Stdout

	// 生效配置
	for _, opt := range opts {
		opt.apply(logger)
	}

	// init zap core
	core := zapcore.NewCore(logger.encoder, logger.writer, logger.level)

	// 附加部分选项功能
	zapOpts := []zap.Option{zap.AddCaller(), zap.AddStacktrace(ErrorLevel), zap.Development()}

	// 创建 Logger
	logger.logger = zap.New(core, zapOpts...)
	return
}
