# give-help

É pra sua vida

## <a name="quick-start"></a>Quick Start

Follow the steps below to create and deploy this service on aws Lambda.

1. **Install Serverless Framework via npm:**

```bash
npm install -g serverless
```

2. **Set-up your [Provider Credentials](https://github.com/serverless/serverless/blob/master/docs/providers/aws/guide/credentials.md)**. [Watch the video on setting up credentials](https://www.youtube.com/watch?v=HSd9uYj2LJA)

3. **Configure Database:**

```bash
# Open file serverless.yml
vim serverless.yml
# Put your database variables here:
    environment:
      DATABASE_DBNAME: "*****YOUR DATABASE NAME*****"
      DATABASE_HOST: "*****YOUR DATABASE HOST*****"
      DATABASE_PASS: "*****YOUR DATABASE PASS*****"
      DATABASE_USER: "*****YOUR DATABASE USER*****"
# Ps: scripts to create a test database are here: (./scripts/db)

```

4. **Configure AWS Subnets and Security Group:**

```bash
# Open file serverless.yml
vim serverless.yml
# Put your VPC and Security Goup variables here:
    vpc:
      securityGroupIds:
        - "*****YOUR DATABASE SECURITY GROUP ID*****"
      subnetIds:
        - "*****YOUR DATABASE SUBNET 1 ID*****"
        - "*****YOUR DATABASE SUBNET 2 ID*****"
```

5. **Configure Firebase:**
```bash
# To generate a private key file for your service account see: (https://firebase.google.com/docs/admin/setup#initialize_the_sdk)

# Copy your json file in this path:
./bin/etc/serviceAccountKey.json
```

6. **Deploy a Service:**

Use this when you have made changes to your Functions, Events or Resources in `serverless.yml` or you simply want to deploy all changes within your Service at the same time.

```bash
make deploy
```

7. **Remove the Service:**

Removes all Functions, Events and Resources from your AWS account.

```bash
serverless remove
```

## Database ER
![Alt text](doc/db-models.jpg "DB ER")