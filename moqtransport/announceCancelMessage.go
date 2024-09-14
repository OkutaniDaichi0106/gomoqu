package moqtransport

import (
	"github.com/quic-go/quic-go/quicvarint"
)

type AnnounceCancelMessage struct {
	TrackNamespace TrackNamespace
}

func (ac AnnounceCancelMessage) serialize() []byte {
	/*
	 * Serialize as following formatt
	 *
	 * ANNOUNCE_CANCEL Payload {
	 *   Track Namespace ([]byte),
	 * }
	 */

	// TODO?: Check track namespace exists

	// TODO: Tune the length of the "b"
	b := make([]byte, 0, 1<<10) /* Byte slice storing whole data */

	// Append the type of the message
	b = quicvarint.Append(b, uint64(ANNOUNCE_CANCEL))

	// Append the supported versions
	b = ac.TrackNamespace.append(b)

	return b
}

// func (ac *AnnounceCancelMessage) deserialize(r quicvarint.Reader) error {
// 	// Get Message ID and check it
// 	id, err := deserializeHeader(r)
// 	if err != nil {
// 		return err
// 	}
// 	if id != ANNOUNCE_CANCEL {
// 		return errors.New("unexpected message")
// 	}

// 	return ac.deserializeBody(r)
// }

func (ac *AnnounceCancelMessage) deserializeBody(r quicvarint.Reader) error {
	// Get Track Namespace
	if ac.TrackNamespace == nil {
		ac.TrackNamespace = make(TrackNamespace, 0, 1)
	}

	err := ac.TrackNamespace.deserialize(r)

	return err
}
