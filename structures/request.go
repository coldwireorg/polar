package structures

type Request struct {
	Proof   string      // Signed request with sender private key
	Content interface{} // content
}
