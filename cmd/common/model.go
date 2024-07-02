package common

type GosterConfig struct {
	ConfigName string `json:"ConfigName"`
}

type GostConfig struct {
	Services []Services `json:"services"`
	Chains   []Chains   `json:"chains"`
}
type Handler struct {
	Type     string   `json:"type"`
	Chain    string   `json:"chain"`
	Auth     Auth     `json:"auth"`
	Metadata Metadata `json:"metadata"`
}
type Listener struct {
	Type     string   `json:"type"`
	Chain    string   `json:"chain"`
	Metadata Metadata `json:"metadata"`
}
type Forwarder struct {
	Nodes []Nodes `json:"nodes"`
}
type Services struct {
	Name      string    `json:"name"`
	Addr      string    `json:"addr"`
	Handler   Handler   `json:"handler"`
	Listener  Listener  `json:"listener"`
	Forwarder Forwarder `json:"forwarder"`
	Metadata  Metadata  `json:"metadata"`
}
type Connector struct {
	Type string `json:"type"`
	Auth Auth   `json:"auth"`
}
type TLS struct {
	ServerName string `json:"serverName"`
}
type Dialer struct {
	Type string `json:"type"`
	TLS  TLS    `json:"tls"`
}
type Nodes struct {
	Name      string    `json:"name"`
	Addr      string    `json:"addr"`
	Connector Connector `json:"connector"`
	Dialer    Dialer    `json:"dialer"`
}
type Hops struct {
	Name  string  `json:"name"`
	Nodes []Nodes `json:"nodes"`
}
type Chains struct {
	Name string `json:"name"`
	Hops []Hops `json:"hops"`
}

type Metadata struct {
	Bind string `json:"bind"`
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//
//type GostConfig struct {
//	Admissions []struct {
//		File struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Matchers []string `json:"matchers"`
//		Name     string   `json:"name"`
//		Plugin   struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload    int  `json:"reload"`
//		Reverse   bool `json:"reverse"`
//		Whitelist bool `json:"whitelist"`
//	} `json:"admissions"`
//	API struct {
//		Accesslog bool   `json:"accesslog"`
//		Addr      string `json:"addr"`
//		Auth      struct {
//			Password string `json:"password"`
//			Username string `json:"username"`
//		} `json:"auth"`
//		Auther     string `json:"auther"`
//		PathPrefix string `json:"pathPrefix"`
//	} `json:"api"`
//	Authers []struct {
//		Auths []struct {
//			Password string `json:"password"`
//			Username string `json:"username"`
//		} `json:"auths"`
//		File struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Name   string `json:"name"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload int `json:"reload"`
//	} `json:"authers"`
//	Bypasses []struct {
//		File struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Matchers []string `json:"matchers"`
//		Name     string   `json:"name"`
//		Plugin   struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload    int  `json:"reload"`
//		Reverse   bool `json:"reverse"`
//		Whitelist bool `json:"whitelist"`
//	} `json:"bypasses"`
//	Chains    []Chain `json:"chains"`
//	Climiters []struct {
//		File struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Limits []string `json:"limits"`
//		Name   string   `json:"name"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload int `json:"reload"`
//	} `json:"climiters"`
//	Hops []struct {
//		Bypass   string   `json:"bypass"`
//		Bypasses []string `json:"bypasses"`
//		File     struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		Hosts string `json:"hosts"`
//		HTTP  struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Interface string `json:"interface"`
//		Name      string `json:"name"`
//		Nodes     []struct {
//			Addr string `json:"addr"`
//			Auth struct {
//				Password string `json:"password"`
//				Username string `json:"username"`
//			} `json:"auth"`
//			Bypass    string   `json:"bypass"`
//			Bypasses  []string `json:"bypasses"`
//			Connector struct {
//				Auth struct {
//					Password string `json:"password"`
//					Username string `json:"username"`
//				} `json:"auth"`
//				Metadata struct {
//					AdditionalProp1 string `json:"additionalProp1"`
//					AdditionalProp2 string `json:"additionalProp2"`
//					AdditionalProp3 string `json:"additionalProp3"`
//				} `json:"metadata"`
//				TLS struct {
//					CaFile     string `json:"caFile"`
//					CertFile   string `json:"certFile"`
//					CommonName string `json:"commonName"`
//					KeyFile    string `json:"keyFile"`
//					Options    struct {
//						CipherSuites []string `json:"cipherSuites"`
//						MaxVersion   string   `json:"maxVersion"`
//						MinVersion   string   `json:"minVersion"`
//					} `json:"options"`
//					Organization string `json:"organization"`
//					Secure       bool   `json:"secure"`
//					ServerName   string `json:"serverName"`
//					Validity     int    `json:"validity"`
//				} `json:"tls"`
//				Type string `json:"type"`
//			} `json:"connector"`
//			Dialer struct {
//				Auth struct {
//					Password string `json:"password"`
//					Username string `json:"username"`
//				} `json:"auth"`
//				Metadata struct {
//					AdditionalProp1 string `json:"additionalProp1"`
//					AdditionalProp2 string `json:"additionalProp2"`
//					AdditionalProp3 string `json:"additionalProp3"`
//				} `json:"metadata"`
//				TLS struct {
//					CaFile     string `json:"caFile"`
//					CertFile   string `json:"certFile"`
//					CommonName string `json:"commonName"`
//					KeyFile    string `json:"keyFile"`
//					Options    struct {
//						CipherSuites []string `json:"cipherSuites"`
//						MaxVersion   string   `json:"maxVersion"`
//						MinVersion   string   `json:"minVersion"`
//					} `json:"options"`
//					Organization string `json:"organization"`
//					Secure       bool   `json:"secure"`
//					ServerName   string `json:"serverName"`
//					Validity     int    `json:"validity"`
//				} `json:"tls"`
//				Type string `json:"type"`
//			} `json:"dialer"`
//			Host  string `json:"host"`
//			Hosts string `json:"hosts"`
//			HTTP  struct {
//				Header struct {
//					AdditionalProp1 string `json:"additionalProp1"`
//					AdditionalProp2 string `json:"additionalProp2"`
//					AdditionalProp3 string `json:"additionalProp3"`
//				} `json:"header"`
//				Host string `json:"host"`
//			} `json:"http"`
//			Interface string `json:"interface"`
//			Metadata  struct {
//				AdditionalProp1 string `json:"additionalProp1"`
//				AdditionalProp2 string `json:"additionalProp2"`
//				AdditionalProp3 string `json:"additionalProp3"`
//			} `json:"metadata"`
//			Name     string `json:"name"`
//			Network  string `json:"network"`
//			Path     string `json:"path"`
//			Protocol string `json:"protocol"`
//			Resolver string `json:"resolver"`
//			Sockopts struct {
//				Mark int `json:"mark"`
//			} `json:"sockopts"`
//			TLS struct {
//				Options struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Secure     bool   `json:"secure"`
//				ServerName string `json:"serverName"`
//			} `json:"tls"`
//		} `json:"nodes"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload   int    `json:"reload"`
//		Resolver string `json:"resolver"`
//		Selector struct {
//			FailTimeout int    `json:"failTimeout"`
//			MaxFails    int    `json:"maxFails"`
//			Strategy    string `json:"strategy"`
//		} `json:"selector"`
//		Sockopts struct {
//			Mark int `json:"mark"`
//		} `json:"sockopts"`
//	} `json:"hops"`
//	Hosts []struct {
//		File struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Mappings []struct {
//			Aliases  []string `json:"aliases"`
//			Hostname string   `json:"hostname"`
//			IP       string   `json:"ip"`
//		} `json:"mappings"`
//		Name   string `json:"name"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload int `json:"reload"`
//	} `json:"hosts"`
//	Ingresses []struct {
//		File struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Name   string `json:"name"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload int `json:"reload"`
//		Rules  []struct {
//			Endpoint string `json:"endpoint"`
//			Hostname string `json:"hostname"`
//		} `json:"rules"`
//	} `json:"ingresses"`
//	Limiters []struct {
//		File struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Limits []string `json:"limits"`
//		Name   string   `json:"name"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload int `json:"reload"`
//	} `json:"limiters"`
//	Log struct {
//		Format   string `json:"format"`
//		Level    string `json:"level"`
//		Output   string `json:"output"`
//		Rotation struct {
//			Compress   bool `json:"compress"`
//			LocalTime  bool `json:"localTime"`
//			MaxAge     int  `json:"maxAge"`
//			MaxBackups int  `json:"maxBackups"`
//			MaxSize    int  `json:"maxSize"`
//		} `json:"rotation"`
//	} `json:"log"`
//	Metrics struct {
//		Addr string `json:"addr"`
//		Auth struct {
//			Password string `json:"password"`
//			Username string `json:"username"`
//		} `json:"auth"`
//		Auther string `json:"auther"`
//		Path   string `json:"path"`
//	} `json:"metrics"`
//	Profiling struct {
//		Addr string `json:"addr"`
//	} `json:"profiling"`
//	Recorders []struct {
//		File struct {
//			Path string `json:"path"`
//			Sep  string `json:"sep"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Name   string `json:"name"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		TCP struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//		} `json:"tcp"`
//	} `json:"recorders"`
//	Resolvers []struct {
//		Name        string `json:"name"`
//		Nameservers []struct {
//			Addr     string `json:"addr"`
//			Async    bool   `json:"async"`
//			Chain    string `json:"chain"`
//			ClientIP string `json:"clientIP"`
//			Hostname string `json:"hostname"`
//			Only     string `json:"only"`
//			Prefer   string `json:"prefer"`
//			Timeout  int    `json:"timeout"`
//			TTL      int    `json:"ttl"`
//		} `json:"nameservers"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//	} `json:"resolvers"`
//	Rlimiters []struct {
//		File struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Limits []string `json:"limits"`
//		Name   string   `json:"name"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload int `json:"reload"`
//	} `json:"rlimiters"`
//	Routers []struct {
//		File struct {
//			Path string `json:"path"`
//		} `json:"file"`
//		HTTP struct {
//			Timeout int    `json:"timeout"`
//			URL     string `json:"url"`
//		} `json:"http"`
//		Name   string `json:"name"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//		Redis struct {
//			Addr     string `json:"addr"`
//			Db       int    `json:"db"`
//			Key      string `json:"key"`
//			Password string `json:"password"`
//			Type     string `json:"type"`
//		} `json:"redis"`
//		Reload int `json:"reload"`
//		Routes []struct {
//			Gateway string `json:"gateway"`
//			Net     string `json:"net"`
//		} `json:"routes"`
//	} `json:"routers"`
//	Sds []struct {
//		Name   string `json:"name"`
//		Plugin struct {
//			Addr    string `json:"addr"`
//			Timeout int    `json:"timeout"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Token string `json:"token"`
//			Type  string `json:"type"`
//		} `json:"plugin"`
//	} `json:"sds"`
//	Services []struct {
//		Addr       string   `json:"addr"`
//		Admission  string   `json:"admission"`
//		Admissions []string `json:"admissions"`
//		Bypass     string   `json:"bypass"`
//		Bypasses   []string `json:"bypasses"`
//		Climiter   string   `json:"climiter"`
//		Forwarder  struct {
//			Name  string `json:"name"`
//			Nodes []struct {
//				Addr string `json:"addr"`
//				Auth struct {
//					Password string `json:"password"`
//					Username string `json:"username"`
//				} `json:"auth"`
//				Bypass   string   `json:"bypass"`
//				Bypasses []string `json:"bypasses"`
//				Host     string   `json:"host"`
//				HTTP     struct {
//					Header struct {
//						AdditionalProp1 string `json:"additionalProp1"`
//						AdditionalProp2 string `json:"additionalProp2"`
//						AdditionalProp3 string `json:"additionalProp3"`
//					} `json:"header"`
//					Host string `json:"host"`
//				} `json:"http"`
//				Name     string `json:"name"`
//				Network  string `json:"network"`
//				Path     string `json:"path"`
//				Protocol string `json:"protocol"`
//				TLS      struct {
//					Options struct {
//						CipherSuites []string `json:"cipherSuites"`
//						MaxVersion   string   `json:"maxVersion"`
//						MinVersion   string   `json:"minVersion"`
//					} `json:"options"`
//					Secure     bool   `json:"secure"`
//					ServerName string `json:"serverName"`
//				} `json:"tls"`
//			} `json:"nodes"`
//			Selector struct {
//				FailTimeout int    `json:"failTimeout"`
//				MaxFails    int    `json:"maxFails"`
//				Strategy    string `json:"strategy"`
//			} `json:"selector"`
//		} `json:"forwarder"`
//		Handler struct {
//			Auth struct {
//				Password string `json:"password"`
//				Username string `json:"username"`
//			} `json:"auth"`
//			Auther     string   `json:"auther"`
//			Authers    []string `json:"authers"`
//			Chain      string   `json:"chain"`
//			ChainGroup struct {
//				Chains   []string `json:"chains"`
//				Selector struct {
//					FailTimeout int    `json:"failTimeout"`
//					MaxFails    int    `json:"maxFails"`
//					Strategy    string `json:"strategy"`
//				} `json:"selector"`
//			} `json:"chainGroup"`
//			Limiter  string `json:"limiter"`
//			Metadata struct {
//				AdditionalProp1 string `json:"additionalProp1"`
//				AdditionalProp2 string `json:"additionalProp2"`
//				AdditionalProp3 string `json:"additionalProp3"`
//			} `json:"metadata"`
//			Retries int `json:"retries"`
//			TLS     struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Type string `json:"type"`
//		} `json:"handler"`
//		Hosts     string `json:"hosts"`
//		Interface string `json:"interface"`
//		Limiter   string `json:"limiter"`
//		Listener  struct {
//			Auth struct {
//				Password string `json:"password"`
//				Username string `json:"username"`
//			} `json:"auth"`
//			Auther     string   `json:"auther"`
//			Authers    []string `json:"authers"`
//			Chain      string   `json:"chain"`
//			ChainGroup struct {
//				Chains   []string `json:"chains"`
//				Selector struct {
//					FailTimeout int    `json:"failTimeout"`
//					MaxFails    int    `json:"maxFails"`
//					Strategy    string `json:"strategy"`
//				} `json:"selector"`
//			} `json:"chainGroup"`
//			Metadata struct {
//				AdditionalProp1 string `json:"additionalProp1"`
//				AdditionalProp2 string `json:"additionalProp2"`
//				AdditionalProp3 string `json:"additionalProp3"`
//			} `json:"metadata"`
//			TLS struct {
//				CaFile     string `json:"caFile"`
//				CertFile   string `json:"certFile"`
//				CommonName string `json:"commonName"`
//				KeyFile    string `json:"keyFile"`
//				Options    struct {
//					CipherSuites []string `json:"cipherSuites"`
//					MaxVersion   string   `json:"maxVersion"`
//					MinVersion   string   `json:"minVersion"`
//				} `json:"options"`
//				Organization string `json:"organization"`
//				Secure       bool   `json:"secure"`
//				ServerName   string `json:"serverName"`
//				Validity     int    `json:"validity"`
//			} `json:"tls"`
//			Type string `json:"type"`
//		} `json:"listener"`
//		Metadata struct {
//			AdditionalProp1 string `json:"additionalProp1"`
//			AdditionalProp2 string `json:"additionalProp2"`
//			AdditionalProp3 string `json:"additionalProp3"`
//		} `json:"metadata"`
//		Name      string `json:"name"`
//		Recorders []struct {
//			Metadata struct {
//				AdditionalProp1 string `json:"additionalProp1"`
//				AdditionalProp2 string `json:"additionalProp2"`
//				AdditionalProp3 string `json:"additionalProp3"`
//			} `json:"Metadata"`
//			Name   string `json:"name"`
//			Record string `json:"record"`
//		} `json:"recorders"`
//		Resolver string `json:"resolver"`
//		Rlimiter string `json:"rlimiter"`
//		Sockopts struct {
//			Mark int `json:"mark"`
//		} `json:"sockopts"`
//	} `json:"services"`
//	TLS struct {
//		CaFile     string `json:"caFile"`
//		CertFile   string `json:"certFile"`
//		CommonName string `json:"commonName"`
//		KeyFile    string `json:"keyFile"`
//		Options    struct {
//			CipherSuites []string `json:"cipherSuites"`
//			MaxVersion   string   `json:"maxVersion"`
//			MinVersion   string   `json:"minVersion"`
//		} `json:"options"`
//		Organization string `json:"organization"`
//		Secure       bool   `json:"secure"`
//		ServerName   string `json:"serverName"`
//		Validity     int    `json:"validity"`
//	} `json:"tls"`
//}
//
//type Chain struct {
//	Hops     []Hop `json:"hops"`
//	Metadata struct {
//		AdditionalProp1 string `json:"additionalProp1"`
//		AdditionalProp2 string `json:"additionalProp2"`
//		AdditionalProp3 string `json:"additionalProp3"`
//	} `json:"metadata"`
//	Name string `json:"name"`
//}
//
//type Hop struct {
//	Bypass   string   `json:"bypass"`
//	Bypasses []string `json:"bypasses"`
//	File     struct {
//		Path string `json:"path"`
//	} `json:"file"`
//	Hosts string `json:"hosts"`
//	HTTP  struct {
//		Timeout int    `json:"timeout"`
//		URL     string `json:"url"`
//	} `json:"http"`
//	Interface string `json:"interface"`
//	Name      string `json:"name"`
//	Nodes     []Node `json:"nodes"`
//	Plugin    struct {
//		Addr    string `json:"addr"`
//		Timeout int    `json:"timeout"`
//		TLS     struct {
//			CaFile     string `json:"caFile"`
//			CertFile   string `json:"certFile"`
//			CommonName string `json:"commonName"`
//			KeyFile    string `json:"keyFile"`
//			Options    struct {
//				CipherSuites []string `json:"cipherSuites"`
//				MaxVersion   string   `json:"maxVersion"`
//				MinVersion   string   `json:"minVersion"`
//			} `json:"options"`
//			Organization string `json:"organization"`
//			Secure       bool   `json:"secure"`
//			ServerName   string `json:"serverName"`
//			Validity     int    `json:"validity"`
//		} `json:"tls"`
//		Token string `json:"token"`
//		Type  string `json:"type"`
//	} `json:"plugin"`
//	Redis struct {
//		Addr     string `json:"addr"`
//		Db       int    `json:"db"`
//		Key      string `json:"key"`
//		Password string `json:"password"`
//		Type     string `json:"type"`
//	} `json:"redis"`
//	Reload   int    `json:"reload"`
//	Resolver string `json:"resolver"`
//	Selector struct {
//		FailTimeout int    `json:"failTimeout"`
//		MaxFails    int    `json:"maxFails"`
//		Strategy    string `json:"strategy"`
//	} `json:"selector"`
//	Sockopts struct {
//		Mark int `json:"mark"`
//	} `json:"sockopts"`
//}
//
//type Node struct {
//	Addr string `json:"addr"`
//	Auth struct {
//		Password string `json:"password"`
//		Username string `json:"username"`
//	} `json:"auth"`
//	Bypass    string    `json:"bypass"`
//	Bypasses  []string  `json:"bypasses"`
//	Connector Connector `json:"connector"`
//	Dialer    Dialer    `json:"dialer"`
//	Host      string    `json:"host"`
//	Hosts     string    `json:"hosts"`
//	HTTP      struct {
//		Header struct {
//			AdditionalProp1 string `json:"additionalProp1"`
//			AdditionalProp2 string `json:"additionalProp2"`
//			AdditionalProp3 string `json:"additionalProp3"`
//		} `json:"header"`
//		Host string `json:"host"`
//	} `json:"http"`
//	Interface string `json:"interface"`
//	Metadata  struct {
//		AdditionalProp1 string `json:"additionalProp1"`
//		AdditionalProp2 string `json:"additionalProp2"`
//		AdditionalProp3 string `json:"additionalProp3"`
//	} `json:"metadata"`
//	Name     string `json:"name"`
//	Network  string `json:"network"`
//	Path     string `json:"path"`
//	Protocol string `json:"protocol"`
//	Resolver string `json:"resolver"`
//	Sockopts struct {
//		Mark int `json:"mark"`
//	} `json:"sockopts"`
//	TLS struct {
//		Options struct {
//			CipherSuites []string `json:"cipherSuites"`
//			MaxVersion   string   `json:"maxVersion"`
//			MinVersion   string   `json:"minVersion"`
//		} `json:"options"`
//		Secure     bool   `json:"secure"`
//		ServerName string `json:"serverName"`
//	} `json:"tls"`
//}
//
//type Connector struct {
//	Auth struct {
//		Password string `json:"password"`
//		Username string `json:"username"`
//	} `json:"auth"`
//	Metadata struct {
//		AdditionalProp1 string `json:"additionalProp1"`
//		AdditionalProp2 string `json:"additionalProp2"`
//		AdditionalProp3 string `json:"additionalProp3"`
//	} `json:"metadata"`
//	TLS struct {
//		CaFile     string `json:"caFile"`
//		CertFile   string `json:"certFile"`
//		CommonName string `json:"commonName"`
//		KeyFile    string `json:"keyFile"`
//		Options    struct {
//			CipherSuites []string `json:"cipherSuites"`
//			MaxVersion   string   `json:"maxVersion"`
//			MinVersion   string   `json:"minVersion"`
//		} `json:"options"`
//		Organization string `json:"organization"`
//		Secure       bool   `json:"secure"`
//		ServerName   string `json:"serverName"`
//		Validity     int    `json:"validity"`
//	} `json:"tls"`
//	Type string `json:"type"`
//}
//
//type Dialer struct {
//	Auth struct {
//		Password string `json:"password"`
//		Username string `json:"username"`
//	} `json:"auth"`
//	Metadata struct {
//		AdditionalProp1 string `json:"additionalProp1"`
//		AdditionalProp2 string `json:"additionalProp2"`
//		AdditionalProp3 string `json:"additionalProp3"`
//	} `json:"metadata"`
//	TLS struct {
//		CaFile     string `json:"caFile"`
//		CertFile   string `json:"certFile"`
//		CommonName string `json:"commonName"`
//		KeyFile    string `json:"keyFile"`
//		Options    struct {
//			CipherSuites []string `json:"cipherSuites"`
//			MaxVersion   string   `json:"maxVersion"`
//			MinVersion   string   `json:"minVersion"`
//		} `json:"options"`
//		Organization string `json:"organization"`
//		Secure       bool   `json:"secure"`
//		ServerName   string `json:"serverName"`
//		Validity     int    `json:"validity"`
//	} `json:"tls"`
//	Type string `json:"type"`
//}
