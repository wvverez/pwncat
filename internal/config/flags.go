package config

import (
	"flag"
)

func (cfg *Config) ParseFlags() {
	flag.StringVar(&cfg.URL, "url", cfg.URL, "URL objetivo con PWN")
	flag.StringVar(&cfg.URL, "u", cfg.URL, "URL objetivo con PWN (shorthand)")
	flag.StringVar(&cfg.Wordlist, "wordlist", cfg.Wordlist, "Ruta a la wordlist")
	flag.StringVar(&cfg.Wordlist, "w", cfg.Wordlist, "Ruta a la wordlist (shorthand)")
	flag.StringVar(&cfg.Method, "method", cfg.Method, "Método HTTP")
	flag.StringVar(&cfg.Method, "X", cfg.Method, "Método HTTP (shorthand)")

	flag.IntVar(&cfg.Threads, "threads", cfg.Threads, "Número de workers")
	flag.IntVar(&cfg.Threads, "t", cfg.Threads, "Número de workers (shorthand)")
	flag.IntVar(&cfg.Rate, "rate", cfg.Rate, "Peticiones/segundo (0 = sin límite)")
	flag.IntVar(&cfg.Rate, "r", cfg.Rate, "Peticiones/segundo (shorthand)")
	flag.DurationVar(&cfg.Timeout, "timeout", cfg.Timeout, "Timeout por petición")
	flag.DurationVar(&cfg.Timeout, "to", cfg.Timeout, "Timeout por petición (shorthand)")
	flag.IntVar(&cfg.Retry, "retry", cfg.Retry, "Número de reintentos")
	flag.IntVar(&cfg.Retry, "re", cfg.Retry, "Número de reintentos (shorthand)")
	flag.DurationVar(&cfg.Delay, "delay", cfg.Delay, "Delay entre peticiones")
	flag.DurationVar(&cfg.Delay, "dl", cfg.Delay, "Delay entre peticiones (shorthand)")

	flag.StringVar(&cfg.MatchStatus, "status", cfg.MatchStatus, "Códigos a matchear (ej: 200,301)")
	flag.StringVar(&cfg.MatchStatus, "s", cfg.MatchStatus, "Códigos a matchear (shorthand)")
	flag.StringVar(&cfg.MatchSize, "size-match", cfg.MatchSize, "Tamaño a matchear (min-max)")
	flag.StringVar(&cfg.MatchSize, "sm", cfg.MatchSize, "Tamaño a matchear (shorthand)")
	flag.StringVar(&cfg.MatchRegex, "regex", cfg.MatchRegex, "Regex a matchear")
	flag.StringVar(&cfg.MatchRegex, "rg", cfg.MatchRegex, "Regex a matchear (shorthand)")

	flag.StringVar(&cfg.ExcludeStatus, "exclude", cfg.ExcludeStatus, "Códigos a excluir (ej: 404,500)")
	flag.StringVar(&cfg.ExcludeStatus, "e", cfg.ExcludeStatus, "Códigos a excluir (shorthand)")
	flag.StringVar(&cfg.ExcludeSize, "exclude-size", cfg.ExcludeSize, "Tamaño a excluir (min-max)")
	flag.StringVar(&cfg.ExcludeSize, "ex", cfg.ExcludeSize, "Tamaño a excluir (shorthand)")
	flag.StringVar(&cfg.ExcludeRegex, "regex-exclude", cfg.ExcludeRegex, "Regex a excluir")
	flag.StringVar(&cfg.ExcludeRegex, "rx", cfg.ExcludeRegex, "Regex a excluir (shorthand)")

	flag.StringVar(&cfg.Output, "output", cfg.Output, "Archivo de salida")
	flag.StringVar(&cfg.Output, "o", cfg.Output, "Archivo de salida (shorthand)")
	flag.StringVar(&cfg.Format, "format", cfg.Format, "Formato: json|csv|html")
	flag.StringVar(&cfg.Format, "f", cfg.Format, "Formato (shorthand)")
	flag.BoolVar(&cfg.NoColor, "no-color", cfg.NoColor, "Desactivar colores")
	flag.BoolVar(&cfg.NoColor, "nc", cfg.NoColor, "Desactivar colores (shorthand)")
	flag.BoolVar(&cfg.Verbose, "verbose", cfg.Verbose, "Modo verboso")
	flag.BoolVar(&cfg.Verbose, "v", cfg.Verbose, "Modo verboso (shorthand)")
	flag.StringVar(&cfg.DebugLog, "debug-log", cfg.DebugLog, "Archivo de depuración")
	flag.StringVar(&cfg.DebugLog, "log", cfg.DebugLog, "Archivo de depuración (shorthand)")

	flag.StringVar(&cfg.ConfigFile, "config", cfg.ConfigFile, "Archivo de configuración")
	flag.StringVar(&cfg.ConfigFile, "cfg", cfg.ConfigFile, "Archivo de configuración (shorthand)")
	flag.StringVar(&cfg.Replay, "replay", cfg.Replay, "URL para replay")
	flag.StringVar(&cfg.Replay, "rp", cfg.Replay, "URL para replay (shorthand)")
	flag.StringVar(&cfg.Cert, "cert", cfg.Cert, "Certificado TLS")
	flag.StringVar(&cfg.Key, "key", cfg.Key, "Clave TLS")
	flag.BoolVar(&cfg.Insecure, "insecure", cfg.Insecure, "Ignorar TLS")
	flag.BoolVar(&cfg.Insecure, "k", cfg.Insecure, "Ignorar TLS (shorthand)")
	flag.BoolVar(&cfg.ShowVersion, "version", cfg.ShowVersion, "Mostrar versión")

	flag.Parse()
}
