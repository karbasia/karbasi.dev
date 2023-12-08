# My Personal Blog Site
This is the code repository for `karbasi.dev`. Feel free to use it and modify it to fit your needs! I quickly built my site with SvelteKit, Skeleton and Pocketbase.

# Development Quick Start

Start the Pocketbase instance by running `go run main.go serve`.

Run `npm install` to install the SvelteKit modules

Rename `.env.dev.local` to `.env`

Finally, run `npm run dev` to start the dev instance.

The default dev URLs are:

- Pocketbase admin: http://127.0.0.1:8090/_/
- SvelteKit frontend: http://localhost:5173

Please ensure to log into the Pocketbase admin panel and set up the first account. From here, you can enter posts and site details.

# Production Deployment

I'm currently running my site on an old laptop with Cloudflare Tunnels. To build the production files, run the following commands:

Pocketbase: `go build main.go`
SvelteKit: `npm run build`

From here, you will have a `main` executable for the backend and a `build` folder for the frontend node app.

Set up a local environment variable `POCKETBASE_URL` that's pointing to the Pocketbase endpoint. This can be an internal URL since SvelteKit will perform SSR on all requests.

I've set up these two as services on my computer abd utilize an nginx proxy to expose certain paths to the Cloudflare Tunnel. Further instructions TBD!