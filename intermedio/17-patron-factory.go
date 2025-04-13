package main 
import "fmt"


// 1. definimos la interfaz comun se va hacer para SMS y Email 

type INotificationFactory interface{
	sendNotification()
	GetSender() ISender
}

type ISender interface{
	GetSenderMethod() string
	GetSenderChannel() string
}

// 2. Creamos dos structs que la implementen 

	//---- primera creacion de objectos para  SMS ----
	type SMSNotification struct{

	}

	func (SMSNotification) sendNotification(){
		fmt.Println("Sending Notificacion via SMS")
	}

	func (SMSNotification) GetSender()ISender{
		return SMSNotificationSender{}
	}

	type SMSNotificationSender struct{

	}

	func (SMSNotificationSender) GetSenderMethod() string{
		return "SMS"

	}

	func (SMSNotificationSender) GetSenderChannel() string{
		return "Twilio"

	}
	//---- segunda creacion de objectos para  SMS ----

	type EmailNotification struct{

	}
	func (EmailNotification) sendNotification(){
		fmt.Println("Sending Notification via email")
	}

	func (EmailNotification) GetSender()ISender{
		return EmailNotificationSender{}
	}
	type EmailNotificationSender struct{

	}

	func (EmailNotificationSender) GetSenderMethod()string{
		return "Email"
	}

	func (EmailNotificationSender) GetSenderChannel()string{
		return "SES"
	}


// 3. creamos la funcion factory 
	func getNotificationFactory(notificationType string)(INotificationFactory, error){
		if notificationType == "SMS"{
			return SMSNotification{},nil
		}
		if notificationType == "Email"{
			return EmailNotification{},nil

		}
		return nil, fmt.Errorf("no notification type")

	}

	func sendNotification(f INotificationFactory) {
		f.sendNotification()
	}

	func getMethod(f INotificationFactory) {
		fmt.Println(f.GetSender().GetSenderMethod())
	}

func main(){

// 4. implementamos la función factory (AnimalFactory) 
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)

}

