# ga-qr
`ga-qr` is a tool to print qr code of a google auth key. You can scan generated qr code to add on Google Authanticator app. 

## Installation
`% go get -u github.com/oguzhane/ga-qr`
## Usage
`% ga-qr -s <secretkey> -i <issuer> -l <label>`
### Flags
    -i string
        Issuer
    -l string
        Label of account
    -s string
        Auth secret key
    -t string
        Format type: [totp|hotp] (default "totp")