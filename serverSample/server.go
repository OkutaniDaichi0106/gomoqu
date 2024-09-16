package main

import (
	"context"
	"go-moq/moqtransport"
	"go-moq/moqtransport/moqtversion"
	"log"
	"net/http"
	"time"

	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/webtransport-go"
)

func main() {
	ws := webtransport.Server{
		H3: http3.Server{
			Addr: "0.0.0.0:8443",
			//TLSConfig: tlsConfig, //TODO: set appropriate tls config
		},
		//CheckOrigin: func(r *http.Request) bool {},
	}
	ms := moqtransport.Server{
		WebTransportServer: &ws,
		Versions:           []moqtversion.Version{moqtversion.LATEST},
	}

	// Handle the Publisher
	ms.OnPublisher(func(sess moqtransport.PublisherSession) {
		_, err := sess.ReceiveAnnounce()
		if err != nil {
			log.Fatal(err)
		}

		err = sess.SendAnnounceOk()
		if err != nil {
			log.Fatal(err)
		}

		err = sess.ReceiveObjects(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	})

	// Handle the Subscriber
	ms.OnSubscriber(func(sess moqtransport.SubscriberSession) {
		_, err := sess.ReceiveSubscribe()
		if err != nil {
			log.Fatal(err)
		}

		err = sess.SendSubscribeOk(30 * time.Minute)
		if err != nil {
			log.Fatal(err)
		}

		sess.DeliverObjects()
	})

	http.HandleFunc("/setup", func(w http.ResponseWriter, r *http.Request) {

		// Establish WebTransport connection after receive EXTEND CONNECT message
		sess, err := ms.Upgrade(w, r)
		if err != nil {
			log.Println(err)
			return
		}

		sess.HandleRole()

		//Handle the session. Here goes the application logic
		//sess.CloseWithError(1234, "stop connection!!!")
	})

	ms.ListenAndServeTLS("localhost.pem", "localhost-key.pem")
}
