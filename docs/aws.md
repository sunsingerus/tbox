# Required components
1. `aws`
1. `aws-iam-authenticator`

## How to install `aws`
Download into `~/tmp`
```bash
cd ~/tmp
```
Latest version
```bash
curl https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip -o aws-cli.zip
```
specific version
```bash
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64-2.0.30.zip" -o "awscliv2.zip"
```

```bash
unzip aws-cli.zip
./aws/install --install-dir ~/local/aws-cli --bin-dir ~/bin
aws --version
```

## How to install `aws-iam-authenticator`
```bash
cd ~/tmp
curl https://amazon-eks.s3.us-west-2.amazonaws.com/1.21.2/2021-07-05/bin/linux/amd64/aws-iam-authenticator -o aws-iam-authenticator
chmod +x ./aws-iam-authenticator
mv aws-iam-authenticator ~/bin
aws-iam-authenticator version
```
