    const fastify = require('fastify')({ logger: true });

    fastify.register(require('@fastify/rate-limit'), {
    max: 100,           // max requests
    timeWindow: '1 minute'
    });

    // Register auth plugin first
    fastify.register(require('./plugins/authplugin'));

    //acts as a proxy
    fastify.register(require('@fastify/reply-from'),{
    base: 'http://localhost:8080/' // Base URL for forwarding requests
    });

    // Now register routes that depend on the auth plugin

    fastify.register(require('./routes/auth'), {prefix: '/api/v1'});

    fastify.register(require('./routes/incidents'), {prefix: '/api/v1'});


        const start = async () => {
        try {
            await fastify.listen({ port: 3000, host: '0.0.0.0' });
            fastify.log.info('Server listening on http://localhost:3000');
        } catch (err) {
            fastify.log.error(err);
            process.exit(1);
        }
        };

    start();
