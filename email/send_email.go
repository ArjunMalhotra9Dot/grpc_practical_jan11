package email

import (
	"context"
	"log"

	pb "github.com/ArjunMalhotra07/proto"
	"gopkg.in/gomail.v2"
)

type EmailServiceServer struct {
	pb.UnimplementedEmailServiceServer
}

func (s *EmailServiceServer) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "arjun03.malhotra@gmail.com")
	m.SetHeader("To", req.GetTo())
	m.SetHeader("Subject", req.GetSubject())
	m.SetBody("text/plain", req.GetBody())

	d := gomail.NewDialer("smtp.gmail.com", 587, "arjun03.malhotra@gmail.com", "msjm zhbb roht yudz")
	if err := d.DialAndSend(m); err != nil {
		return nil, err
	}

	log.Println("Email sent to:", req.GetTo())
	return &pb.SendEmailResponse{Status: "Success"}, nil
}
