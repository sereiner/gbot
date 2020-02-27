package gbot

type Option func(*option)

type option struct {
	path     string
	server   ServerConf
	download DownloadConf
	console  ConsoleConf
	session  string
}

type ServerConf struct {
	status bool
	ip     string
	port   int
}

type DownloadConf struct {
	Image        bool
	Voice        bool
	Video        bool
	Emoticon     bool
	File         bool
	EmoticonPath string //' => $path.'emoticons',
}

type ConsoleConf struct {
	output  bool // 是否输出
	message bool // 是否输出接收消息 （若上面为 false 此处无效）
}

func WithPath(path string) Option {
	return func(o *option) {
		o.path = path
	}
}

func WithServer(ip string, port int, status bool) Option {
	return func(o *option) {
		o.server = ServerConf{
			status: status,
			ip:     ip,
			port:   port,
		}
	}
}

func WithDownload(image, voice, video, emoticon, file bool, EmoticonPath string) Option {
	return func(o *option) {
		o.download = DownloadConf{
			Image:        image,
			Voice:        voice,
			Video:        video,
			Emoticon:     emoticon,
			File:         file,
			EmoticonPath: o.path + EmoticonPath,
		}
	}
}

func WithConsole(output, message bool) Option {
	return func(o *option) {
		o.console = ConsoleConf{
			output:  output,
			message: message,
		}
	}
}
