# Email Query

EmailQuery is an interface designed for efficiently searching information in an email database. This tool provides an intuitive user experience to explore and retrieve relevant data stored in the database, facilitating management and search.

## How to Run the App

### Uploading Enron Data to AWS S3

Upload the Enron data available at [http://www.cs.cmu.edu/~enron](http://www.cs.cmu.edu/~enron) to S3 in AWS. These are the data that will be subsequently indexed in the ZincSearch search engine [https://zincsearch-docs.zinc.dev/](https://zincsearch-docs.zinc.dev/).

**Note:** You must download the special binary for your machine, the binaries that you see in the .zinc file are the version
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
     REQUEST_ORIGIN=""
     ```

   - Front - .env:
     ```
     VITE_API_URL=""
     ```

2. Afterward, execute the `docker-compose up` command.

### Deploying

If you want to deploy, the `main.ts` file is configured to work with AWS:

1. First, it's desirable to install Terraform.
   
2. Configure the key associated with your account and the names you want
   
3. Then execute the following commands:

terraform init
terraform plan
terraform apply


**Note:** Ensure that you have the necessary permissions and configurations set up for AWS and Terraform before running the deployment commands.

4. Once the instance is running, connect to it and execute the following commands:

- git clone https://github.com/KerenBermeo/CorreoQuery.git
- sudo apt install awscli
- aws configure
- aws s3 cp s3://bucket_name/compressed_file_name /tmp/compressed_file_name
- Create a folder and unzip the file into that folder.
- Delete the compressed file.
- cd CorreoQuery, enter the cloned repository, and configure the environment variables at the root - of both the backend and frontend.
- Finally, execute the command: docker-compose up -d.

**Note:** To view the program online, remember to use HTTP.