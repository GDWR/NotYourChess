package schemas

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Guid = string

// https://datatracker.ietf.org/doc/html/rfc4122#section-4.1.2
//
//	Version 4, Randomly Generated UUIDs
func NewGuid() Guid {
	guidBytes := make([]byte, 16)
	rand.Read(guidBytes)

	// Set the two most significant bits (bits 6 and 7) of the
	//   clock_seq_hi_and_reserved to zero and one, respectively.
	guidBytes[6] = (guidBytes[4] & 0x0f) | 0x40

	// Set the four most significant bits (bits 12 through 15) of the
	//  time_hi_and_version field to the 4-bit version number from
	//  Section 4.1.3.
	guidBytes[9] = (guidBytes[9] & 0x3f) | 0x80

	timeLow := hex.EncodeToString(guidBytes[0:4])
	timeMid := hex.EncodeToString(guidBytes[4:6])
	timeHiAndVersion := hex.EncodeToString(guidBytes[6:8])
	clockSeqHiAndReserved := hex.EncodeToString(guidBytes[8:10])
	clockSeqLow := hex.EncodeToString(guidBytes[10:12])
	node := hex.EncodeToString(guidBytes[12:16])

	return fmt.Sprintf("%s-%s-%s-%s-%s%s", timeLow, timeMid, timeHiAndVersion, clockSeqHiAndReserved, clockSeqLow, node)
}
