# gitme-shelter
It's a Go utility that backs up git repos to an S3 bucket. Backing up your repos can be seen as shelter from the storm of a stupid move or attack. Also, Gimmie Shelter by the Rolling Stones was playing when I created this repo. 🤷  
Also an excuse to play around with Go a little for fun and knowledge.

# Use
## Running Locally
1. Restore packages: `go mod download && go mod verify`
2. Build: `go build`
3. Run: `bin/gitme-shelter -cfg YOUR_CFG_FILE`

## Running In Docker
1. Restore packages: `go mod download && go mod verify`
2. Build Docker image: `make docker-build`
3. Run: `docker run gitmeshelter`

## Environment Variables
If you're running locally and you have aws-cli configured and git credentials configured things _may_ work with your configuration right away. However, I expect the typical use case will be firing up the docker container via a cron job. For that to work with the base Docker image, the following env vars are required.  
**AWS_ACCESS_KEY**: The access key used to access the AWS S3 bucket in the config.  
**AWS_SECRET_ACCESS_KEY**: Secret key for AWS access.  
**GIT_USERNAME**: Your git username (Github email address). Only required if targeting private repos  
**GIT_PASSWORD**: Your git password (Github token). Only required if targeting private repos  
**LOG_LEVEL**: Sets verbosity of the logs with this. Default is `debug`. Accepts: `trace,debug,warn,error`

## Config File
The app needs a configuration file which, for lack of a better format, is YAML. You can see an example in `default-test.yaml` which is copied to the Docker image by default.  
Your config **must** contain 1 value for `s3Bucket` and 1 value for `awsRegion`.  
Your config must also contain 1 or more entries under `githubRepo` with a `name` and `uri`. The name is arbitrary and is used in creating the backup file. The uri should be a Github repo link (tested only with https links at this time)

## Putting It All Together
The most common use case is likely to be running this in a k8s cronjob or something like that. If so, you know how to use env vars via configmaps and secrets. If just running the container it might look something like this:
```
docker run -e AWS_ACCESS_KEY_ID=your_key 
    -e AWS_SECRET_ACCESS_KEY="your_secret" 
    -e GIT_USERNAME="your_username" 
    -e GIT_PASSWORD="your_token" 
    -e LOG_LEVEL=info 
    --mount type=bind,source=YOUR_CONFIG_DIR,destination=/gms/cfg 
    gitmeshelter /gms/gitme-shelter -cfg /gms/cfg/my-secret-cfg.yaml
```

🚨 **WARNING!** Whatever you do, DO NOT EVER check in AWS or Github secrets! 🚨

### ToDo
- My interfaces are bad and I feel bad. I should do something about that.
- Finish tests
- Put a limit on simultaneous uploads
- Actually handle errors/output from upload routines
- Clean up dependencies
- Clean up code to allow for better substitution and extension down the road (ex. storage options)
- Test & fix to work with git SSH access
