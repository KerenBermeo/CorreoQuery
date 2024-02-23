# Email Query

EmailQuery is an interface designed for efficiently searching information in an email database. This tool provides an intuitive user experience to explore and retrieve relevant data stored in the database, facilitating management and search.

## How to Run the App

### Uploading Enron Data to AWS S3

Upload the Enron data available at [http://www.cs.cmu.edu/~enron](http://www.cs.cmu.edu/~enron) to S3 in AWS. These are the data that will be subsequently indexed in the ZincSearch search engine [https://zincsearch-docs.zinc.dev/](https://zincsearch-docs.zinc.dev/).

**Note:** You must download the special binary for your machine, the binaries that you see in the .zinc file are the version
- zincsearch_0.4.10_windows_arm64.tar.gz
- zincsearch_0.4.10_linux_x86_64.tar.gz
See which one is right for you at: https://github.com/zincsearch/zincsearch/releases

### Running Locally

If you wish to run it locally:

1. Clone the repository and set up your environment variables in .env files:

   - Back - vars.env:
     ```
     ZINC_FIRST_ADMIN_USER=admin
     ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123
     ZINC_SEARCH_URL=""
     ZINC_SERVER_NAME_INDEX=desired name for your index
     BACK_ROOT_DIRECTORY=here you should put the data direction
     BACK_LISTEN_SERVER=""
     ```

   - Front - .env:
     ```
     VITE_API_URL=""
     ```

2. Afterward, execute the `docker-compose up` command.

### Deploying

If you want to deploy, the `main.ts` file is configured to work with AWS:

1. First, it's desirable to install Terraform.
   
2. Configure the keys and names you desire.
   
3. Set up the necessary variables in a `vars.ts` file.
   
4. Then execute the following commands:

terraform init
terraform plan
terraform apply


**Note:** Ensure that you have the necessary permissions and configurations set up for AWS and Terraform before running the deployment commands.

