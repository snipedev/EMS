const fp = require('fastify-plugin');

async function AuthPlugin(fastify,options) {
        await fastify.register(require('@fastify/jwt'), {
                secret: process.env.JWT_SECRET || 'your-secret-key-here'
        })

        fastify.decorate('authenticate', async function(request, reply) {
            try {
                await request.jwtVerify()
            } catch (err) {
                reply.code(401).send(err)
            }
        })
}

module.exports = fp(AuthPlugin);