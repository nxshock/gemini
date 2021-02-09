package gemini

const (
	StatusInput          = 10 // As per definition of single-digit code 1 in 3.2.
	StatusSensitiveInput = 11 // As per status code 10, but for use with sensitive input such as passwords. Clients should present the prompt as per status code 10, but the user's input should not be echoed to the screen to prevent it being read by "shoulder surfers".

	// The request was handled successfully and a response body will follow the response header.
	StatusSuccess = 20

	// The server is redirecting the client to a new location for the requested resource.
	StatusRedirectTemporary = 30

	StatusRedirectPermanent = 31 // The requested resource should be consistently requested from the new URL provided in future. Tools like search engine indexers or content aggregators should update their configurations to avoid requesting the old URL, and end-user clients may automatically update bookmarks, etc. Note that clients which only pay attention to the initial digit of status codes will treat this as a temporary redirect. They will still end up at the right place, they just won't be able to make use of the knowledge that this redirect is permanent, so they'll pay a small performance penalty by having to follow the redirect each time.
	StatusTemporaryFailure  = 40 // As per definition of single-digit code 4 in 3.2.
	StatusServerUnavailable = 41 // The server is unavailable due to overload or maintenance. (cf HTTP 503)
	StatusCgiError          = 42 // A CGI process, or similar system for generating dynamic content, died unexpectedly or timed out.
	StatusProxyError        = 43 // A proxy request failed because the server was unable to successfully complete a transaction with the remote host. (cf HTTP 502, 504)
	StatusSlowDown          = 44 // Rate limiting is in effect. <META> is an integer number of seconds which the client must wait before another request is made to this server. (cf HTTP 429)

	// The request has failed.
	StatusPermanentFailure = 50

	StatusNotFound            = 51 // The requested resource could not be found but may be available in the future. (cf HTTP 404) (struggling to remember this important status code? Easy: you can't find things hidden at Area 51!)
	StatusGone                = 52 // The resource requested is no longer available and will not be available again. Search engines and similar tools should remove this resource from their indices. Content aggregators should stop requesting the resource and convey to their human users that the subscribed resource is gone. (cf HTTP 410)
	StatusProxyRequestRefused = 53 // The request was for a resource at a domain not served by the server and the server does not accept proxy requests.
	StatusBadRequest          = 59 // The server was unable to parse the client's request, presumably due to a malformed request. (cf HTTP 400)

	// The requested resource requires a client certificate to access.
	StatusClientCertificateRequired = 60

	StatusCertificateNotAuthorised = 61 // The supplied client certificate is not authorised for accessing the particular requested resource. The problem is not with the certificate itself, which may be authorised for other resources.
	StatusCertificateNotValid      = 62 // The supplied client certificate was not accepted because it is not valid. This indicates a problem with the certificate in and of itself, with no consideration of the particular requested resource. The most likely cause is that the certificate's validity start date is in the future or its expiry date has passed, but this code may also indicate an invalid signature, or a violation of a X509 standard requirements. The <META> should provide more information about the exact error.)
)

var (
	scheme               = "gemini"
	newLine              = []byte("\r\n")
	defaultListenAddress = ":1965"
)
