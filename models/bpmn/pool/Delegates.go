package pool

<<<<<<< HEAD
import "strings"

var (
	structPool = "pool"
)

func IsPool(field string) bool {
	return strings.ToLower(field) == structPool
=======
// e *Participant, typ, typProcessRef, name string, hash ...string
func SetParticipant(p DelegateParameter) {
	p.PPT.SetID(p.T, p.H[0])
	p.PPT.SetName(p.N)
	p.PPT.SetProcessRef(p.PR, p.H[1])
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}
