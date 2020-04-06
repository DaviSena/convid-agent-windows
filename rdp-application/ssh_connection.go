package main

import (
	"math"
	"math/rand"
	"time"

	client "github.com/jairsjunior/go-ssh-client-tunnel/clientv2"
	"github.com/sirupsen/logrus"
)

var successConfirmedTimer *time.Timer

func connect(sshServerHost string, sshServerPort int, user string, password string, localRDPHost string, localRDPPort int, tunneltoHost string, tunneltoPort int) {
	localRDPEndpoint := client.Endpoint{
		Host: localRDPHost,
		Port: localRDPPort,
	}
	tunneltoEndpoint := client.Endpoint{
		Host: tunneltoHost,
		Port: tunneltoPort,
	}
	sshServerEndpoint := client.Endpoint{
		Host: sshServerHost,
		Port: 2222,
	}
	logrus.Infof("Connecting to remote server: SSHServerHost: %s | SSHServerPort: %d | user: %s | password: %s", sshServerHost, 2222, user, password)
	logrus.Infof("Local routing to: Host: %s | Port: %d", localRDPHost, localRDPPort)
	logrus.Infof("Tunneling to: Host: %s | Port: %d", tunneltoHost, tunneltoPort)

	isConnected := make(chan bool)
	at := 0.0

	for {
		go client.CreateConnectionLocalV2(user, password, localRDPEndpoint, tunneltoEndpoint, sshServerEndpoint, isConnected)
		// successConfirmedTimer = time.AfterFunc(time.Duration(3)*time.Second, func() {
		// 	agentInstance.runtime.Events.Emit("ConnectionSucceed")
		// })
		// err := client.CreateConnectionLocalV2(user, password, localRDPEndpoint, tunneltoEndpoint, sshServerEndpoint)
		// if err != nil {
		// 	successConfirmedTimer.Stop()
		// 	agentInstance.runtime.Events.Emit("ConnectionError")
		// }

		v := <-isConnected

		if !v {

			if at < 6 {
				r := rand.Intn(10) + int(math.Exp2(at))
				time.Sleep(time.Duration(r) * time.Second)
				logrus.Warningf("Error connecting to SSH... retrying in %d seconds.", r)
				at++
				continue
			} else {
				agentInstance.runtime.Events.Emit("ConnectionError")
				break
			}

		} else {
			agentInstance.runtime.Events.Emit("ConnectionSucceed")
			logrus.Infof("===>>>> CONNECTED")
			break
		}

	}
}
