const baseUrl = "api/v1/incidents";

async function routes(fastify, options) {
  fastify.get('/incidents',{
    preHandler: [fastify.authenticate]
  },async (request, reply) => {
    // Placeholder for fetching incidents from a database or external API
    //console.log('authenticate is:', fastify.authenticate); 
    return reply.from(baseUrl);
  });

  fastify.post('/incidents', {
    preHandler: [fastify.authenticate]
  }, async (request, reply) => {
    // Placeholder for saving the new incident to a database
    console.log('authenticate is:', fastify.authenticate);
    return reply.from(baseUrl);
  });

  fastify.get('/incidents/:id', {
    //preHandler: [fastify.authenticate]
  }, async (request, reply) => {
    const { id } = request.params;
    // Placeholder for fetching a specific incident by ID from a database or external API
    console.log('authenticate is:', fastify.authenticate);
    return reply.from(`${baseUrl}/${id}`);
  });
}

module.exports = routes;
