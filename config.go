package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config represents the configuration of the server application.
type Config struct {
	Address string
	Port    string
	Prefix  string
	Dir     string
	TLS     *TLS
	Log     Logging
	Realm   string
	Users   map[string]*UserInfo
	Cors    Cors
}

// Logging allows definition for logging each CRUD method.
type Logging struct {
	Error  bool
	Create bool
	Read   bool
	Update bool
	Delete bool
}

// TLS allows specification of a certificate and private key file.
type TLS struct {
	CertFile string
	KeyFile  string
}

// UserInfo allows storing of a password and user directory.
type UserInfo struct {
	Password string
	Subdir   *string
}

// Cors contains settings related to Cross-Origin Resource Sharing (CORS)
type Cors struct {
	Origin      string
	Credentials bool
}

func ValueOrDefault(envKey string, defaultValue string) string {
	valueFromEnv := os.Getenv(envKey)
	if valueFromEnv == "" {
		return defaultValue
	}
	return valueFromEnv
}

// ParseConfig parses the application configuration and sets defaults.
func ParseConfig() *Config {
	tlsConfig := getTLSConfig()
	userConfig := getUserConfig()
	var cfg = &Config{
		Address: ValueOrDefault("DR_BIND_TO_IP", "127.0.0.1"),
		Port:    ValueOrDefault("DR_BIND_TO_PORT", "8000"),
		Prefix:  ValueOrDefault("DR_URL_PREFIX", "/"),
		Dir:     ValueOrDefault("DR_ROOT", "/tmp"),
		Users:   userConfig,
		TLS:     tlsConfig,
		Realm:   ValueOrDefault("DR_AUTH_REALM", "dav-rage"),
		Cors:    Cors{Credentials: false},
		Log:     Logging{Error: true},
	}

	return cfg
}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func getUserConfig() map[string]*UserInfo {
	userConfigFile := os.Getenv("DR_AUTH_FILE")
	if userConfigFile == "" {
		return nil
	}
	authFileContent := readFile(userConfigFile)
	users := make(map[string]*UserInfo)
	for _, line := range authFileContent {
		username, passhash := splitUserPass(line)
		if username != "" && passhash != "" {
			continue
		}
		subdir := "/" + username
		users[username] = &UserInfo{Password: passhash, Subdir: &subdir}
	}

	return users
}

func splitUserPass(line string) (string, string) {
	// split line at ":"
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

func getTLSConfig() *TLS {
	certFile := os.Getenv("DR_TLS_CERT")
	keyFile := os.Getenv("DR_TLS_KEY")
	if certFile == "" || keyFile == "" {
		return nil
	}
	return &TLS{CertFile: certFile, KeyFile: keyFile}
}

func (cfg *Config) AuthenticationNeeded() bool {
	return cfg.Users != nil && len(cfg.Users) != 0
}

func (cfg *Config) ensureUserDirs() {
	if _, err := os.Stat(cfg.Dir); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(cfg.Dir, os.ModePerm)
		if mkdirErr != nil {
			fmt.Println(mkdirErr)
			fmt.Println("Could not create directory: " + cfg.Dir)
			return
		}
		fmt.Println("Created base dir: " + cfg.Dir)
	}

	for _, user := range cfg.Users {
		if user.Subdir != nil {
			path := filepath.Join(cfg.Dir, *user.Subdir)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, os.ModePerm)
				fmt.Println("Created user dir: " + path)
			}
		}
	}
}
