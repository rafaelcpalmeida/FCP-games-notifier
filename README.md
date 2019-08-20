# FCP Games Notifier

FCP Games Notifier is a small project, which consists of a scrapper integrating with Twilio to send an SMS whenever [F.C. Porto](https://www.fcporto.pt/pt) is playing on its stadium. As I've noticed that during football day games, subway gets a little too much overcrowded. As I'm not a football fan, I never know when it's going to be game day. I felt the need to find a way to know when Porto was playing at home without having to check daily.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

- Golang installed on your computer
- Valid [Twilio](https://www.twilio.com/) account

### Building

In order to compile the code into one binary file just run:

```bash
go build -o fcp-games-notifier bin/main.go
```

## Deployment

To use this in a live system you need to setup three environment variables:

- *FCP_GAMES_NOTIFIER_RECIPIENT_NUMBER* - The SMS recipient phone number, [E.164 phone number format](https://en.wikipedia.org/wiki/E.164).
- *TWILIO_ACCOUNT_SID* - Twilio specific account identifier.
- *TWILIO_ACCOUNT_AUTH_TOKEN* - Twilio specific account authentication token.

To have the binary being executed daily you can use:

```bash
00 9 * * 1-5 . /root/.bashrc; /root/fcp-games-notifier
```

The above will execute every working day at 9:00. You can costumize the cron per your needs / desires.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/rafaelcpalmeida/FCP-games-notifier/releases).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* [PurpleBooth](https://github.com/PurpleBooth) for this README template.
