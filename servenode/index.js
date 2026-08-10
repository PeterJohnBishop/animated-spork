import http from 'node:http';

const hostname = '0.0.0.0'; // Important: must be 0.0.0.0 inside Docker, not localhost
const port = process.env.PORT || 3000;
const environment = process.env.NODE_ENV || 'development';

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end(`Hello from Node.js REST API! Running in ${environment} mode.\n`);
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/ in ${environment} mode`);
});
