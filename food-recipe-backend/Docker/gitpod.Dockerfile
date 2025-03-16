FROM gitpod/workspace-full

# Set working directory
WORKDIR /workspace

# Install dependencies
RUN apt-get update && apt-get install -y postgresql-client

# Install Go
RUN curl -fsSL https://golang.org/dl/go1.21.linux-amd64.tar.gz | tar -xz -C /usr/local
ENV PATH="/usr/local/go/bin:$PATH"

# Set up environment variables
ENV DATABASE_URL=postgres://postgres:postgres@db:5432/recipesdb?sslmode=disable

# Install Hasura CLI
RUN curl -L https://github.com/hasura/graphql-engine/raw/stable/cli/get.sh | bash

# Expose backend port
EXPOSE 8080
