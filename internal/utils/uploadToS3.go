package utils

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/example/internal/entity"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePayrollPDF(payroll entity.Payroll, employeeName string) (*bytes.Buffer, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Payroll Summary")

	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)

	data := []struct {
		Label string
		Value interface{}
	}{
		{"Employee Name", employeeName},
		{"Month", payroll.Month.Format("January 2006")},
		{"Base Salary", payroll.BaseSalary},
		{"Total Hours", payroll.TotalHours},
		{"Overtime Hours", payroll.OverTimeHours},
		{"Paid Leaves", payroll.PaidLeaves},
		{"Unpaid Leaves", payroll.UnpaidLeaves},
		{"Deductions", payroll.Deductions},
		{"Bonuses", payroll.Bonuses},
		{"Final Salary", payroll.FinalSalary},
	}

	for _, item := range data {
		pdf.CellFormat(60, 10, fmt.Sprintf("%s:", item.Label), "", 0, "L", false, 0, "")
		pdf.CellFormat(60, 10, fmt.Sprintf("%v", item.Value), "", 1, "L", false, 0, "")
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func UploadToS3(ctx context.Context, pdfBuffer *bytes.Buffer, filename string) (string, error) {
	// Load from environment variables
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")

	if accessKey == "" || secretKey == "" || region == "" || bucket == "" {
		return "", fmt.Errorf("missing one or more AWS environment variables")
	}

	// Load AWS config with static credentials
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	key := "payrolls/" + filename
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(pdfBuffer.Bytes()),
		ACL:         "bucket-owner-full-control", // Change to "private" if needed
		ContentType: aws.String("application/pdf"),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload pdf: %w", err)
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)
	return url, nil
}
