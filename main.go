package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/smtp"
)

func main() {
	mxs, err := net.LookupMX("example.com")
	if err != nil {
		log.Fatal(err)
	}
	for _, mx := range mxs {
		log.Println(mx.Host)
	}

	to := []string{"rcpt@example.com"}
	msg := []byte("To: rcpt@example.com\r\n" +
		"Subject: Golang's SMTP Sample\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err = sendMail("mx1.example.com:25", "sender@otherdomain.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the remote SMTP server.
	// c, err := smtp.Dial("mx3.example.com:25")
	// c, err := smtp.Dial("mx2.example.com:25")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Set the sender and recipient first
	// if err := c.Mail("fasfs@example.org"); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := c.Rcpt("example@example.com"); err != nil {
	// 	log.Fatal(err)
	// }

	// // Send the email body.
	// wc, err := c.Data()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = fmt.Fprintf(wc, "This is the email body")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = wc.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Send the QUIT command and close the connection.
	// err = c.Quit()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func sendMail(addr string, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Hello("atlanssia.com"); err != nil {
		return err
	}
	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{InsecureSkipVerify: true}

		if err = c.StartTLS(config); err != nil {
			return err
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
