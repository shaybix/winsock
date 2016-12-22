// +build windows
package winsock

// Windows Sockets Error Codes - https://msdn.microsoft.com/en-us/library/windows/desktop/ms740668(v=vs.85).aspx
const (
	WSA_INVALID_HANDLE = 6 // Specified event object handle is invalid

	WSA_NOT_ENOUGH_MEMORY = 8 // Insufficient memory available

	WSA_INVALID_PARAMETER = 87 // One or more parameters are invalid

	WSA_OPERATION_ABORTED = 995 // Overlapped operation aborted
	WSA_IO_INCOMPLETE     = 996 // Overlapped I/O event object not in signaled state
	WSA_IO_PENDING        = 997 // Overlapped operations will complete later

	WSAEINTR = 10004 // Interrupted function call

	WSAEBADF = 10009 // File handle is not valid

	WSAEACCES = 10013 // Permission denied
	WSAEFAULT = 10014 // Bad address

	WSAEINVAL = 10022 // Invalid Argument

	WSAEMFIL = 10024 // Too many open files

	WSAEWOULDBLOCK     = 10035 // Resource temporarily unavailable
	WSAEINPROGRESS     = 10036 // Operation now in progress
	WSAEALREADY        = 10037 // Operation already in progress
	WSAENOTSOCK        = 10038 // Socket operation on nonsocket
	WSADESTADDRREQ     = 10039 // Destination address required
	WSAEMSGSIZE        = 10040 // Message too long
	WSAEPROTOTYPE      = 10041 // Protocol wrong type for socket
	WSAENOPROTOOPT     = 10042 // Bad protocol option
	WSAEPROTONOSUPPORT = 10043 // Protocol not supported
	WSAESOCKTNOSUPPORT = 10044 // Socket type not supported
	WSAEOPNOTSUPP      = 10045 // Operation not supported
	WSAEPFNOSUPPORT    = 10046 // Protocol family not supported
	WSAEAFNOSUPPORT    = 10047 // Address family not supported by protocol family
	WSAEADDRINUSE      = 10048 // Address already in use
	WSAEADDRNOTAVAIL   = 10049 // Cannot assign requested address
	WSAENETDOWN        = 10050 // Network is down
	WSAENETUNREACH     = 10051 // Network is unreachable
	WSAENETRESET       = 10052 // Network dropped connection on reset
	WSAECONNABORTED    = 10053 // Software caused connection abort
	WSAECONNRESET      = 10054 // Connection reset by peer
	WSAENOBUFS         = 10055 // No buffer space available
	WSAEISCONN         = 10056 // Socket is already connected
	WSAENOTCONN        = 10057 // Socket is not connected
	WSAESHUTDOWN       = 10058 // Cannot send after socket shutdown
	WSAETOOMANYREFS    = 10059 // Too many references
	WSAETIMEDOUT       = 10060 // Connection timed out
	WSAECONNREFUSED    = 10061 // Connection refused
	WSAELOOP           = 10062 // Cannot translate name
	WSAENAMETOOLONG    = 10063 // Name too long
	WSAEHOSTDOWN       = 10064 // Host is down
	WSAEHOSTUNREACH    = 10065 // No route to host
	WSAENOTEMPTY       = 10066 // Directory not empty
	WSAEPROCLIM        = 10067 // Too many processes
	WSAEUSERS          = 10068 // User quota exceeded
	WSAEDQUOT          = 10069 // Disk quota exceeded
	WSAESTALE          = 10070 // Stale file handle reference
	WSAEREMOTE         = 10071 // Item is remote

	WSASYSNOTREADY     = 10091 // Network subsystem is unavailable
	WSAVERNOTSUPPORTED = 10092 // Winsock.dll version out of range
	WSANOTINITIALISED  = 10093 // Successful WSAStartup not yet performed

	WSAEDISCON             = 10101 // Graceful shutdown in progress
	WSAENOMORE             = 10102 // No more results
	WSAECANCELLED          = 10103 // Call has been canceled
	WSAEINVALIDPROCTABLE   = 10104 // Procedure call table is invalid
	WSAEINVALIDPROVIDER    = 10105 // Service provider is invalid
	WSAEPROVIDERFAILEDINIT = 10106 // Service provider failed to initialize
	WSASYSCALLFAILURE      = 10107 // System call failure
	WSASERVICE_NOT_FOUND   = 10108 // Service not found
	WSATYPE_NOT_FOUND      = 10109 // Class type not found
	WSA_E_NO_MORE          = 10110 // No more results
	WSA_E_CANCELLED        = 10111 // Call was canceled
	WSAEREFUSED            = 10112 // Database query was refused

	WSAHOST_NOT_FOUND          = 11001 // Host not found
	WSATRY_AGAIN               = 11002 // Nonauthoritative host not found
	WSANO_RECOVERY             = 11003 // This is a nonrecoverable error
	WSANO_DATA                 = 11004 // Valid name, no data record of requested type
	WSA_QOS_RECEIVERS          = 11005 // QoS receivers
	WSA_QOS_SENDERS            = 11006 // QoS senders
	WSA_WOS_NO_SENDERS         = 11007 // No QoS senders
	WSA_QOS_NO_RECEIVERS       = 11008 // QoS no receivers
	WSA_QOS_REQUEST_CONFIRMED  = 11009 // QoS request confirmed
	WSA_QOS_ADMISSION_FAILURE  = 11010 // QoS admission error
	WSA_QOS_POLICY_FAILURE     = 11011 // QoS policy failure
	WSA_QOS_BAD_STYLE          = 11012 // QoS bad style
	WSA_QOS_BAD_OBJECT         = 11013 // QoS bad object
	WSA_QOS_TRAFFIC_CTRL_ERROR = 11014 // QoS traffic control error
	WSA_QOS_GENERIC_ERROR      = 11015 // QoS generic error
	WSA_QOS_ESERVICETYPE       = 11016 // QoS service type error
	WSA_QOS_EFLOWSPEC          = 11017 // QoS flowspec error
	WSA_QOS_EPROVSPECBUF       = 11018 // Invalid QoS provider buffer
	WSA_QOS_EFILTERSTYLE       = 11019 // Invalid QoS filter style
	WSA_QOS_EFILTERTYPE        = 11020 // Invalid QoS filter type
	WSA_QOS_EFILTERCOUNT       = 11021 // Invalid QoS filter count
	WSA_QOS_EOBJLENGTH         = 11022 // Invalid QoS object length
	WSA_QOS_EFLOWCOUNT         = 11023 // Invalid QoS flow count
	WSA_QOS_EUNKNOWNPSOBJ      = 11024 // Unrecognized QoS object
	WSA_QOS_EPOLICYOBJ         = 11025 // Invalid QoS policy object
	WSA_QOS_EFLOWDESC          = 11026 // Invalid QoS flow descriptor
	WSA_QOS_EPSFLOWSPEC        = 11027 // Invalid QoS provider-specific flowspec
	WSA_QOS_EPSFILTERSPEC      = 11028 // Invalid QoS provider-specific filterspec
	WSA_QOS_ESDMODEOBJ         = 11029 // Invalid QoS shape discard mode object
	WSA_QOS_ESHAPERATEOBJ      = 11030 // Invalid QoS shaping rate object
	WSA_QOS_RESERVED_PETYPE    = 11031 // Reserved policy QoS element type
)

// WSAGetLastError returns the last error code that occurred for the calling thread
func WSAGetLastError() int {}
