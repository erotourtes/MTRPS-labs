import Fastify from "fastify";
import handler from "./src/handler.js";

const host = "0.0.0.0";
const port = 8080;

const fastify = Fastify({ logger: true });

fastify.get("/", handler);

fastify.listen({ port, host }).catch((err) => {
  fastify.log.error(err);
  process.exit(1);
});
