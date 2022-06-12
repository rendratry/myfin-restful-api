package service

import "C"
import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"gopkg.in/gomail.v2"
	"log"
	"myfin/entity/domain"
	"myfin/entity/web"
	"myfin/helper"
	"myfin/repository"
)

type AkunNasabahServiceImpl struct {
	AkunNasabahRepository repository.AkunNasabahRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewAkunNasabahService(akunnasabahRepository repository.AkunNasabahRepository, DB *sql.DB, validate *validator.Validate) *AkunNasabahServiceImpl {
	return &AkunNasabahServiceImpl{
		AkunNasabahRepository: akunnasabahRepository,
		DB:                    DB,
		Validate:              validate}
}

func (service *AkunNasabahServiceImpl) CreateAkun(ctx context.Context, request web.AkunNasabahCreateRequest) web.AkunNasabahResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	gen_otp := helper.EncodeToString(4)

	akunnasabah := domain.AkunNasabah{
		Email:              request.Email,
		KodeOtp:            gen_otp,
		PinUser:            "0",
		Saldo:              "0",
		Nik:                "0",
		NamaLengkap:        "0",
		NoHp:               "0",
		Alamat:             "0",
		TglLahir:           "0",
		KotaLahir:          "0",
		KtpFile:            "0",
		SwaFile:            "0",
		Ava:                "0",
		PertanyaanKeamanan: "0",
		JawabanKeamanan:    "0",
	}

	emailtemplate := "<!DOCTYPE html>\n<html xmlns:v=\"urn:schemas-microsoft-com:vml\" xmlns:o=\"urn:schemas-microsoft-com:office:office\" lang=\"en\">\n\n<head>\n    <title></title>\n    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <!--[if mso]><xml><o:OfficeDocumentSettings><o:PixelsPerInch>96</o:PixelsPerInch><o:AllowPNG/></o:OfficeDocumentSettings></xml><![endif]-->\n    <style>\n        * {\n            box-sizing: border-box;\n        }\n\n        body {\n            margin: 0;\n            padding: 0;\n        }\n\n        a[x-apple-data-detectors] {\n            color: inherit !important;\n            text-decoration: inherit !important;\n        }\n\n        #MessageViewBody a {\n            color: inherit;\n            text-decoration: none;\n        }\n\n        p {\n            line-height: inherit\n        }\n\n        .desktop_hide,\n        .desktop_hide table {\n            mso-hide: all;\n            display: none;\n            max-height: 0px;\n            overflow: hidden;\n        }\n\n        @media (max-width:520px) {\n            .row-content {\n                width: 100% !important;\n            }\n\n            .column .border,\n            .mobile_hide {\n                display: none;\n            }\n\n            table {\n                table-layout: fixed !important;\n            }\n\n            .stack .column {\n                width: 100%;\n                display: block;\n            }\n\n            .mobile_hide {\n                min-height: 0;\n                max-height: 0;\n                max-width: 0;\n                overflow: hidden;\n                font-size: 0px;\n            }\n\n            .desktop_hide,\n            .desktop_hide table {\n                display: table !important;\n                max-height: none !important;\n            }\n        }\n    </style>\n</head>\n\n<body style=\"background-color: #FFFFFF; margin: 0; padding: 0; -webkit-text-size-adjust: none; text-size-adjust: none;\">\n<table class=\"nl-container\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #FFFFFF;\">\n    <tbody>\n    <tr>\n        <td>\n            <table class=\"row row-1\" align=\"center\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                <tbody>\n                <tr>\n                    <td>\n                        <table class=\"row-content stack\" align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000000; width: 500px;\" width=\"500\">\n                            <tbody>\n                            <tr>\n                                <td class=\"column column-1\" width=\"100%\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-top: 5px; padding-bottom: 5px; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;\">\n                                    <table class=\"image_block\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                        <tr>\n                                            <td style=\"width:100%;padding-right:0px;padding-left:0px;\">\n                                                <div align=\"center\" style=\"line-height:10px\"><img src=\"https://d15k2d11r6t6rl.cloudfront.net/public/users/Integrators/BeeProAgency/805691_789567/Myfin%20logo.png\" style=\"display: block; height: auto; border: 0; width: 125px; max-width: 100%;\" width=\"125\"></div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"divider_block\" width=\"100%\" border=\"0\" cellpadding=\"10\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                        <tr>\n                                            <td>\n                                                <div align=\"center\">\n                                                    <table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" width=\"100%\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                                        <tr>\n                                                            <td class=\"divider_inner\" style=\"font-size: 1px; line-height: 1px; border-top: 1px solid #BBBBBB;\"><span>&#8202;</span></td>\n                                                        </tr>\n                                                    </table>\n                                                </div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"heading_block\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                        <tr>\n                                            <td style=\"width:100%;text-align:center;\">\n                                                <h1 style=\"margin: 0; color: #393d47; font-size: 23px; font-family: Arial, Helvetica Neue, Helvetica, sans-serif; line-height: 120%; text-align: center; direction: ltr; font-weight: 700; letter-spacing: normal; margin-top: 0; margin-bottom: 0;\"><span class=\"tinyMce-placeholder\">Verifikasi Email</span></h1>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                </td>\n                            </tr>\n                            </tbody>\n                        </table>\n                    </td>\n                </tr>\n                </tbody>\n            </table>\n            <table class=\"row row-2\" align=\"center\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                <tbody>\n                <tr>\n                    <td>\n                        <table class=\"row-content stack\" align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000000; width: 500px;\" width=\"500\">\n                            <tbody>\n                            <tr>\n                                <td class=\"column column-1\" width=\"100%\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-top: 5px; padding-bottom: 5px; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;\">\n                                    <table class=\"paragraph_block\" width=\"100%\" border=\"0\" cellpadding=\"10\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;\">\n                                        <tr>\n                                            <td>\n                                                <div style=\"color:#393d47;font-size:14px;font-family:Arial, Helvetica Neue, Helvetica, sans-serif;font-weight:400;line-height:120%;text-align:center;direction:ltr;letter-spacing:0px;\">\n                                                    <p style=\"margin: 0;\">Anda mendaftar ke layanan MyFin, masukkan kode verifikasi untuk melanjutkan berikut kode verifikasi anda</p>\n                                                </div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"heading_block\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                        <tr>\n                                            <td style=\"width:100%;text-align:center;\">\n                                                <h1 style=\"margin: 0; color: #2b66ff; font-size: 29px; font-family: Arial, Helvetica Neue, Helvetica, sans-serif; line-height: 120%; text-align: center; direction: ltr; font-weight: 700; letter-spacing: normal; margin-top: 0; margin-bottom: 0;\"><span class=\"tinyMce-placeholder\">{{.otp}}</span></h1>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"paragraph_block\" width=\"100%\" border=\"0\" cellpadding=\"10\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;\">\n                                        <tr>\n                                            <td>\n                                                <div style=\"color:#393d47;font-size:14px;font-family:Arial, Helvetica Neue, Helvetica, sans-serif;font-weight:400;line-height:120%;text-align:center;direction:ltr;letter-spacing:0px;\"></div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"paragraph_block\" width=\"100%\" border=\"0\" cellpadding=\"10\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;\">\n                                        <tr>\n                                            <td>\n                                                <div style=\"color:#393d47;font-size:14px;font-family:Arial, Helvetica Neue, Helvetica, sans-serif;font-weight:400;line-height:120%;text-align:center;direction:ltr;letter-spacing:0px;\">\n                                                    <p style=\"margin: 0; margin-bottom: 16px;\">Jika Butuh Bantuan silahkan hubungi</p>\n                                                    <p style=\"margin: 0;\"><strong>Email : help@hellomyfin.engineer</strong></p>\n                                                </div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                </td>\n                            </tr>\n                            </tbody>\n                        </table>\n                    </td>\n                </tr>\n                </tbody>\n            </table>\n        </td>\n    </tr>\n    </tbody>\n</table><!-- End -->\n</body>\n\n</html>"

	//t := template.Must(template.New("Email").Parse(emailtemplate))
	//t.ExecuteTemplate()

	const CONFIG_SMTP_HOST = "smtp.titan.email"
	const CONFIG_SMTP_PORT = 465
	const CONFIG_SENDER_NAME = "MyFin <help@hellomyfin.engineer>"
	const CONFIG_AUTH_EMAIL = "help@hellomyfin.engineer"
	const CONFIG_AUTH_PASSWORD = "adminmyfin"

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", request.Email)
	mailer.SetHeader("Subject", "Verifikasi Email")
	mailer.SetBody("text/html", emailtemplate+gen_otp)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	akunnasabah = service.AkunNasabahRepository.CreateAkun(ctx, tx, akunnasabah)

	return helper.ToAkunNasabahResponse(akunnasabah)
}
