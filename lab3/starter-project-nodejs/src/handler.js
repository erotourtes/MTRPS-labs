const ips = new Set();

export default async function (request) {
  const ip = request.socket.remoteAddress;
  const userAgent = request.headers["user-agent"];
  const message = ips.has(ip) ? `You again!` : `I am watching you, ${ip}!`;
  ips.add(ip);

  return {
    message,
    ip,
    userAgent,
  };
}
