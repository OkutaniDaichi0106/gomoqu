package moqtmessage

import "github.com/quic-go/quic-go/quicvarint"

type SubscribeNamespaceOkMessage struct {
	TrackNamespacePrefix TrackNamespacePrefix
}

func (sno SubscribeNamespaceOkMessage) Serialize() []byte {
	/*
	 * Serialize the message in the following formatt
	 *
	 * SUBSCRIBE_NAMESPACE_OK Message {
	 *   Type (varint) = 0x12,
	 *   Length (varint),
	 *   Track Namespace Prefix (tuple),
	 * }
	 */

	/*
	 * Serialize the payload
	 */
	p := make([]byte, 0, 1<<8)

	// Append the Track Namespace Prefix
	p = sno.TrackNamespacePrefix.Append(p)

	/*
	 * Serialize the whole message
	 */
	b := make([]byte, 0, len(p)+1<<4)

	// Append the message ID
	b = quicvarint.Append(b, uint64(SUBSCRIBE_NAMESPACE_OK))

	// Append the payload
	b = quicvarint.Append(b, uint64(len(p)))
	b = append(b, p...)

	return b
}

func (sno *SubscribeNamespaceOkMessage) DeserializePayload(r quicvarint.Reader) error {
	var tnsp TrackNamespacePrefix

	err := tnsp.Deserialize(r)
	if err != nil {
		return err
	}
	sno.TrackNamespacePrefix = tnsp

	return nil
}
