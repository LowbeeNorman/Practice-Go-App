### Building and running your application

When you're ready, start your application by running:
`docker compose up --build`.

Your application listens on TCP port 38759. You can connect using netcat (nc) or any TCP client (in another terminal of course):

### On Windows, WSL, or macOS (the command Caleb is running using WSL):
nc -v host.docker.internal 38759

### On Linux without Docker Desktop, you can try:
nc -v <your-host-ip> 38759

### Deploying your application to the cloud

First, build your image, e.g.: `docker build -t myapp .`.
If your cloud uses a different CPU architecture than your development
machine (e.g., you are on a Mac M1 and your cloud provider is amd64),
you'll want to build the image for that platform, e.g.:
`docker build --platform=linux/amd64 -t myapp .`.

Then, push it to your registry, e.g. `docker push myregistry.com/myapp`.

Consult Docker's [getting started](https://docs.docker.com/go/get-started-sharing/)
docs for more detail on building and pushing.

### References
* [Docker's Go guide](https://docs.docker.com/language/golang/)
