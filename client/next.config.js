module.exports = {
  async rewrites() {
    // When running Next.js via Node.js (e.g. `dev` mode), proxy API requests
    // to the Go server.

    let port = process.env.SERVER_PORT;
    if (!port || port.length === 0) {
      port = '8080';
    }

    return [
      {
        source: "/api",
        destination: "http://localhost:" + port + "/api",
      },
    ];
  },
  webpack5: true,
  trailingSlash: true,
};
