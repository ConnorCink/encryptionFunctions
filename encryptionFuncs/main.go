package encryptionfuncs

import (
	b64 "encoding/base64"
	"fmt"
)

func encryptTheThing(unencryptedThing string) (encryptedThing string) {
	fmt.Println("Encrypting unencrypted thing")
	encryptedThing = b64.StdEncoding.EncodeToString([]byte(unencryptedThing))
	fmt.Println("Encrypted the thing")
	return
}
