package moqtmessage

import (
	"github.com/quic-go/quic-go/quicvarint"
)

type UnannounceMessage struct {
	TrackNamespace TrackNamespace
}

func (ua UnannounceMessage) Serialize() []byte {
	/*
	 * Serialize the message in the following formatt
	 *
	 * UNANNOUNCE Payload {
	 *   Track Namespace (tuple),
	 * }
	 */

	/*
	 * Serialize the payload
	 */
	p := make([]byte, 0, 1<<8)

	// Append the Track Namespace
	p = ua.TrackNamespace.Append(p)

	/*
	 * Serialize the whole message
	 */
	b := make([]byte, 0, len(p)+1<<4)

	// Append the message type
	b = quicvarint.Append(b, uint64(UNANNOUNCE))

	// Append the payload
	b = quicvarint.Append(b, uint64(len(p)))
	b = append(b, p...)

	return b
}

func (ua *UnannounceMessage) DeserializePayload(r quicvarint.Reader) error {
	var tns TrackNamespace
	err := tns.Deserialize(r)
	if err != nil {
		return err
	}

	ua.TrackNamespace = tns

	return nil
}
