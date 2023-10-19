package utils

const (
	ParseUUIDErr       = "Error parsing UUID, err:%s"
	ParseFormErr       = "Unable to parse form, please verify"
	UploadedFileErr    = "Unable to get file, please verify"
	InvalidFileErr     = "Invalid file type. Only .csv files are allowed."
	CsvIdErr           = "csv err: Invalid data in CSV: ID must be an integer, err: %s"
	CsvDateErr         = "csv err: Invalid data in CSV: Date must be an date with format MM/DD, err: %s"
	CsvTransactionErr  = "csv err: Invalid data in CSV: Transaction must be a number, err: %s"
	ParsingTemplateErr = "Error parsing template, err: %s"
	ExecuteTemplateErr = "Error executing template, err: %s"
	EmailErr           = "Error sending email, err:%s"
	NoUserErr          = "there are no results for this user, please validate"
)

const EmailTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Monthly Financial Report</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0 auto;
            max-width: 600px;
            padding: 20px;
            background-color: #f9f9f9;
            color: #333;
        }

        .container {
            background-color: #fff;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            text-align: center;
            margin: 0 auto;
        }

        h1, h2 {
            color: #007BFF;
        }

        p, li {
            font-size: 18px;
            text-align: center;
        }

        h2 {
            margin-top: 18px;
        }

        ul {
            list-style-type: none;
            padding: 0;
            text-align: center;
        }

        li {
            font-size: 15px;
            margin-bottom: 10px;
            text-align: center;
        }

        .highlight {
            font-weight: bold;
            color: #28a745;
        }

        .footer {
            text-align: center;
            font-size: 14px;
            color: #888;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <div class="container">
        <img src="https://play-lh.googleusercontent.com/acrduPY3FNGodTtdx0e6WGOz_EEqJ7KeRJMJARWnETH-5j2oxV2_Tb4hwIWraQKymd4" alt="Stori Logo" width="200" height="200">
        <h1>Financial Report</h1>
        <p class="highlight">Account Summary:</p>
        <ul>
            <li>Total Balance: {{.Summary.TotalBalance}}</li>
        </ul>

        {{range $month, $stats := .Summary.MonthStats}}
            <h2>Transactions for Month {{$month}}</h2>
            <ul>
                <li>Average Credit Amount: {{$stats.AverageCredit}}</li>
                <li>Average Debit Amount: {{$stats.AverageDebit}}</li>
                <li>Number of transactions: {{$stats.TransactionCount}}</li>
            </ul>
        {{end}}
    </div>
    <div class="footer">
        This is an automated email. Please do not reply.
    </div>

`
