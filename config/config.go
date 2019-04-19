package config

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"io/ioutil"
	"jspring.top/pushmess/log"
	"math/big"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Config 配置
type Config struct {
	ConfigFile *ExplicitString `short:"C" long:"configfile" description:"Path to configuration file"`
	// KeyPair      tls.Certificate `no-flag:"ltx"`
	Listen       string `long:"listen" description:"Listen for RPC connections on this interface/port"`
	GetuiAppid   string `long:"getuiappid" description:"app getuiappid."`
	GetuiAppkey  string `long:"getuiappkey" description:"app getuiappkey."`
	GetuiAppms   string `long:"getuiappms" description:"app getuiappms."`
	HuaweiAppid  string `long:"huaweiappid" description:"app huaweiappid."`
	HuaweiAppsec string `long:"huaweiappsec" description:"app huaweiappsec."`
	PkgName      string `long:"pkgName" description:"app pkgName."`
}

// ExplicitString 是否配置过
type ExplicitString struct {
	Value         string
	explicitlySet bool
}

const (
	defaultConfigFile = "./pushmess.conf"
)

var (
	// Cfg 配置对象
	Cfg *Config
)

// LoadConfig 初始化配置
func LoadConfig() {
	conf := &Config{
		ConfigFile:   NewExplicitString(defaultConfigFile),
		Listen:       os.Getenv("pushmess_listen"),
		GetuiAppid:   os.Getenv("GetuiAppid"),
		GetuiAppkey:  os.Getenv("GetuiAppkey"),
		GetuiAppms:   os.Getenv("GetuiAppms"),
		HuaweiAppid:  os.Getenv("HuaweiAppid"),
		HuaweiAppsec: os.Getenv("HuaweiAppsec"),
		PkgName:      os.Getenv("push_pkgname"),
	}

	preParser := flags.NewParser(conf, flags.Default)
	_, err := preParser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			preParser.WriteHelp(os.Stderr)
		}
		log.Log.Fatalln(err)
	}
	// Load additional config from file.
	configFilePath := conf.ConfigFile.Value
	if conf.ConfigFile.ExplicitlySet() {
		configFilePath = CleanAndExpandPath(configFilePath)
	}
	parser := flags.NewParser(conf, flags.Default)
	err = flags.NewIniParser(parser).ParseFile(configFilePath)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok {
			parser.WriteHelp(os.Stderr)
		}
		log.Log.Fatalln(err)
	}

	// Parse command line options again to ensure they take precedence.
	_, err = parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			parser.WriteHelp(os.Stderr)
		}
		log.Log.Fatalln(err)
	}

	// keyPair, err := OpenRPCKeyPair("./pushmess.key",
	// 	"./pushmess.cert", false, "pushmess")
	// if err != nil {
	// 	log.Log.Fatalln("证书读取或创建失败")
	// }
	// conf.KeyPair = keyPair

	if conf.Listen == "" {
		log.Log.Fatalln("listens获取失败")
	}
	Cfg = conf
}

// NewExplicitString creates a string flag with the provided default value.
func NewExplicitString(defaultValue string) *ExplicitString {
	return &ExplicitString{Value: defaultValue, explicitlySet: false}
}

// ExplicitlySet returns whether the flag was explicitly set through the
// flags.Unmarshaler interface.
func (e *ExplicitString) ExplicitlySet() bool { return e.explicitlySet }

// MarshalFlag implements the flags.Marshaler interface.
func (e *ExplicitString) MarshalFlag() (string, error) { return e.Value, nil }

// UnmarshalFlag implements the flags.Unmarshaler interface.
func (e *ExplicitString) UnmarshalFlag(value string) error {
	e.Value = value
	e.explicitlySet = true
	return nil
}

// CleanAndExpandPath clear 配置
func CleanAndExpandPath(path string) string {
	// NOTE: The os.ExpandEnv doesn't work with Windows cmd.exe-style
	// %VARIABLE%, but they variables can still be expanded via POSIX-style
	// $VARIABLE.
	path = os.ExpandEnv(path)

	if !strings.HasPrefix(path, "~") {
		return filepath.Clean(path)
	}

	// Expand initial ~ to the current user's home directory, or ~otheruser
	// to otheruser's home directory.  On Windows, both forward and backward
	// slashes can be used.
	path = path[1:]

	var pathSeparators string
	if runtime.GOOS == "windows" {
		pathSeparators = string(os.PathSeparator) + "/"
	} else {
		pathSeparators = string(os.PathSeparator)
	}

	userName := ""
	if i := strings.IndexAny(path, pathSeparators); i != -1 {
		userName = path[:i]
		path = path[i:]
	}

	homeDir := ""
	var u *user.User
	var err error
	if userName == "" {
		u, err = user.Current()
	} else {
		u, err = user.Lookup(userName)
	}
	if err == nil {
		homeDir = u.HomeDir
	}
	// Fallback to CWD if user lookup fails or user has no home directory.
	if homeDir == "" {
		homeDir = "."
	}

	return filepath.Join(homeDir, path)
}

// OpenRPCKeyPair 打开或创建RPCkey
func OpenRPCKeyPair(rpckey, rpccert string, onetime bool, org string) (tls.Certificate, error) {
	// Check for existence of the TLS key file.  If one time TLS keys are
	// enabled but a key already exists, this function should error since
	// it's possible that a persistent certificate was copied to a remote
	// machine.  Otherwise, generate a new keypair when the key is missing.
	// When generating new persistent keys, overwriting an existing cert is
	// acceptable if the previous execution used a one time TLS key.
	// Otherwise, both the cert and key should be read from disk.  If the
	// cert is missing, the read error will occur in LoadX509KeyPair.
	_, e := os.Stat(rpckey)
	keyExists := !os.IsNotExist(e)
	switch {
	case onetime && keyExists:
		err := fmt.Errorf("one time TLS keys are enabled, but TLS key "+
			"`%s` already exists", rpckey)
		return tls.Certificate{}, err
	case onetime:
		return generateRPCKeyPair(rpckey, rpccert, false, org)
	case !keyExists:
		return generateRPCKeyPair(rpckey, rpccert, true, org)
	default:
		return tls.LoadX509KeyPair(rpccert, rpckey)
	}
}

// generateRPCKeyPair generates a new RPC TLS keypair and writes the cert and
// possibly also the key in PEM format to the paths specified by the config.  If
// successful, the new keypair is returned.
func generateRPCKeyPair(rpckey, rpccert string, writeKey bool, org string) (tls.Certificate, error) {
	log.Log.Infof("Generating TLS certificates...")

	// Create directories for cert and key files if they do not yet exist.
	certDir, _ := filepath.Split(rpccert)
	keyDir, _ := filepath.Split(rpckey)
	err := os.MkdirAll(certDir, 0700)
	if err != nil {
		return tls.Certificate{}, err
	}
	err = os.MkdirAll(keyDir, 0700)
	if err != nil {
		return tls.Certificate{}, err
	}

	validUntil := time.Now().Add(time.Hour * 24 * 365 * 10)
	cert, key, err := NewTLSCertPair(org, validUntil, nil)
	if err != nil {
		return tls.Certificate{}, err
	}
	keyPair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return tls.Certificate{}, err
	}

	// Write cert and (potentially) the key files.
	err = ioutil.WriteFile(rpccert, cert, 0600)
	if err != nil {
		return tls.Certificate{}, err
	}
	if writeKey {
		err = ioutil.WriteFile(rpckey, key, 0600)
		if err != nil {
			rmErr := os.Remove(rpccert)
			if rmErr != nil {
				log.Log.Warnf("Cannot remove written certificates: %v",
					rmErr)
			}
			return tls.Certificate{}, err
		}
	}

	log.Log.Info("Done generating TLS certificates")
	return keyPair, nil
}

// NewTLSCertPair returns a new PEM-encoded x.509 certificate pair
// based on a 521-bit ECDSA private key.  The machine's local interface
// addresses and all variants of IPv4 and IPv6 localhost are included as
// valid IP addresses.
func NewTLSCertPair(organization string, validUntil time.Time, extraHosts []string) (cert, key []byte, err error) {
	now := time.Now()
	if validUntil.Before(now) {
		return nil, nil, errors.New("validUntil would create an already-expired certificate")
	}

	priv, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// end of ASN.1 time
	endOfTime := time.Date(2049, 12, 31, 23, 59, 59, 0, time.UTC)
	if validUntil.After(endOfTime) {
		validUntil = endOfTime
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate serial number: %s", err)
	}

	host, err := os.Hostname()
	if err != nil {
		return nil, nil, err
	}

	ipAddresses := []net.IP{net.ParseIP("127.0.0.1"), net.ParseIP("::1")}
	dnsNames := []string{host}
	if host != "localhost" {
		dnsNames = append(dnsNames, "localhost")
	}

	addIP := func(ipAddr net.IP) {
		for _, ip := range ipAddresses {
			if bytes.Equal(ip, ipAddr) {
				return
			}
		}
		ipAddresses = append(ipAddresses, ipAddr)
	}
	addHost := func(host string) {
		for _, dnsName := range dnsNames {
			if host == dnsName {
				return
			}
		}
		dnsNames = append(dnsNames, host)
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, nil, err
	}
	for _, a := range addrs {
		ipAddr, _, err := net.ParseCIDR(a.String())
		if err == nil {
			addIP(ipAddr)
		}
	}

	for _, hostStr := range extraHosts {
		host, _, err := net.SplitHostPort(hostStr)
		if err != nil {
			host = hostStr
		}
		if ip := net.ParseIP(host); ip != nil {
			addIP(ip)
		} else {
			addHost(host)
		}
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{organization},
			CommonName:   host,
		},
		NotBefore: now.Add(-time.Hour * 24),
		NotAfter:  validUntil,

		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature |
			x509.KeyUsageCertSign,
		IsCA:                  true, // so can sign self.
		BasicConstraintsValid: true,

		DNSNames:    dnsNames,
		IPAddresses: ipAddresses,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template,
		&template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create certificate: %v", err)
	}

	certBuf := &bytes.Buffer{}
	err = pem.Encode(certBuf, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode certificate: %v", err)
	}

	keybytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal private key: %v", err)
	}

	keyBuf := &bytes.Buffer{}
	err = pem.Encode(keyBuf, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keybytes})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode private key: %v", err)
	}

	return certBuf.Bytes(), keyBuf.Bytes(), nil
}
