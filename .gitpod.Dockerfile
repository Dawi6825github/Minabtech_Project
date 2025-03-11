FROM gitpod/workspace-full

# Install PostgreSQL
RUN sudo apt update && sudo apt install -y postgresql postgresql-client

# Install Golang 1.22
RUN curl -fsSL https://go.dev/dl/go1.22.0.linux-amd64.tar.gz | sudo tar -xz -C /usr/local
ENV PATH="/usr/local/go/bin:${PATH}"

# Install Node.js (for Vue, Nuxt, Apollo)
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash - \
    && sudo apt install -y nodejs \
    && npm install -g @vue/cli vite nuxt vue-apollo vee-validate tailwindcss

# Install Hasura CLI
RUN curl -L https://github.com/hasura/graphql-engine/releases/latest/download/cli_linux_amd64 -o hasura-cli \
    && chmod +x hasura-cli \
    && mv hasura-cli /usr/local/bin/

# Install Docker & Docker Compose
RUN sudo apt install -y docker.io docker-compose

# Set working directory
WORKDIR /workspace
