async function routes(fastify, options) {
  fastify.get('/incidents',{
    preHandler: [fastify.authenticate]
  },async (request, reply) => {
    // Placeholder for fetching incidents from a database or external API
    console.log('authenticate is:', fastify.authenticate); 
    const incidents = [
      { id: 1, title: 'Incident 1', description: 'Description of incident 1' },
      { id: 2, title: 'Incident 2', description: 'Description of incident 2' },
    ];
    return incidents;
  });

  fastify.post('/incidents', {
    preHandler: [fastify.authenticate]
  }, async (request, reply) => {
    const { title, description } = request.body;
    // Placeholder for saving the new incident to a database
    const newIncident = { id: Date.now(), title, description };
    return newIncident;
  });
}

module.exports = routes;
