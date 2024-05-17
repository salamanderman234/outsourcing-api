package tests

import (
	"testing"

	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/jobs"
	"github.com/salamanderman234/outsourcing-api/app/mails"
)

func TestJobDispatchNow(t *testing.T) {
	t.Logf("Testing helpers.JobManager.DispatchNow()...")
	to := "tresnasaputra9@gmail.com"
	mail := mails.NewResetPasswordMail(map[string]any{
		"email": to,
		"token": "fsdf",
	})
	t.Logf("Test Case #1: {to: '%s'}", to)
	job := jobs.NewSendMailJob(mail, []string{to})
	if err := helpers.JobManager.DispatchNow(job); err != nil {
		t.Logf("Failing Test Case #1: %s", err.Error())
	}

}
